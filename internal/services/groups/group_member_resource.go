// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package groups

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/groups/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/manicminer/hamilton/msgraph"
)

func groupMemberResource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		CreateContext: groupMemberResourceCreate,
		ReadContext:   groupMemberResourceRead,
		DeleteContext: groupMemberResourceDelete,

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(5 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Update: pluginsdk.DefaultTimeout(5 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Importer: pluginsdk.ImporterValidatingResourceId(func(id string) error {
			_, err := parse.GroupMemberID(id)
			return err
		}),

		Schema: map[string]*pluginsdk.Schema{
			"group_object_id": {
				Description:      "The object ID of the group you want to add the member to",
				Type:             pluginsdk.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validation.ValidateDiag(validation.IsUUID),
			},

			"member_object_id": {
				Description:      "The object ID of the principal you want to add as a member to the group. Supported object types are Users, Groups or Service Principals",
				Type:             pluginsdk.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validation.ValidateDiag(validation.IsUUID),
			},
		},
	}
}

func groupMemberResourceCreate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Groups.GroupsClient
	directoryObjectsClient := meta.(*clients.Client).Groups.DirectoryObjectsClient
	tenantId := meta.(*clients.Client).TenantID
	groupId := d.Get("group_object_id").(string)
	memberId := d.Get("member_object_id").(string)

	id := parse.NewGroupMemberID(groupId, memberId)

	tf.LockByName(groupResourceName, id.GroupId)
	defer tf.UnlockByName(groupResourceName, id.GroupId)

	group, status, err := client.Get(ctx, groupId, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return tf.ErrorDiagPathF(nil, "object_id", "Group with object ID %q was not found", groupId)
		}
		return tf.ErrorDiagPathF(err, "object_id", "Retrieving group with object ID: %q", groupId)
	}

	existingMembers, _, err := client.ListMembers(ctx, id.GroupId)
	if err != nil {
		return tf.ErrorDiagF(err, "Listing existing members for group with object ID: %q", id.GroupId)
	}
	if existingMembers != nil {
		for _, v := range *existingMembers {
			if strings.EqualFold(v, memberId) {
				return tf.ImportAsExistsDiag("azuread_group_member", id.String())
			}
		}
	}

	memberObject, _, err := directoryObjectsClient.Get(ctx, memberId, odata.Query{})
	if err != nil {
		return tf.ErrorDiagF(err, "Could not retrieve principal object %q", memberId)
	}
	if memberObject == nil {
		return tf.ErrorDiagF(errors.New("returned memberObject was nil"), "Could not retrieve member principal object %q", memberId)
	}
	memberObject.ODataId = (*odata.Id)(utils.String(fmt.Sprintf("%s/v1.0/%s/directoryObjects/%s",
		client.BaseClient.Endpoint, tenantId, memberId)))

	group.Members = &msgraph.Members{*memberObject}

	if _, err := client.AddMembers(ctx, group); err != nil {
		return tf.ErrorDiagF(err, "Adding group member %q to group %q", memberId, groupId)
	}

	d.SetId(id.String())
	return groupMemberResourceRead(ctx, d, meta)
}

func groupMemberResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Groups.GroupsClient

	id, err := parse.GroupMemberID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing Group Member ID %q", d.Id())
	}

	members, status, err := client.ListMembers(ctx, id.GroupId)
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Group with ID %q was not found - removing group member with ID %q from state", id.GroupId, d.Id())
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagF(err, "Retrieving members for group with object ID: %q", id.GroupId)
	}

	var memberObjectId string
	if members != nil {
		for _, objectId := range *members {
			if strings.EqualFold(objectId, id.MemberId) {
				memberObjectId = objectId
				break
			}
		}
	}

	if memberObjectId == "" {
		log.Printf("[DEBUG] Member with ID %q was not found in Group %q - removing from state", id.MemberId, id.GroupId)
		d.SetId("")
		return nil
	}

	tf.Set(d, "group_object_id", id.GroupId)
	tf.Set(d, "member_object_id", memberObjectId)

	return nil
}

func groupMemberResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Groups.GroupsClient

	id, err := parse.GroupMemberID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing Group Member ID %q", d.Id())
	}

	tf.LockByName(groupResourceName, id.GroupId)
	defer tf.UnlockByName(groupResourceName, id.GroupId)

	if _, err := client.RemoveMembers(ctx, id.GroupId, &[]string{id.MemberId}); err != nil {
		return tf.ErrorDiagF(err, "Removing member %q from group with object ID: %q", id.MemberId, id.GroupId)
	}

	// Wait for membership link to be deleted
	if err := helpers.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
		defer func() { client.BaseClient.DisableRetries = false }()
		client.BaseClient.DisableRetries = true
		if _, status, err := client.GetMember(ctx, id.GroupId, id.MemberId); err != nil {
			if status == http.StatusNotFound {
				return utils.Bool(false), nil
			}
			return nil, err
		}
		return utils.Bool(true), nil
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for removal of member %q from group with object ID %q", id.MemberId, id.GroupId)
	}

	return nil
}
