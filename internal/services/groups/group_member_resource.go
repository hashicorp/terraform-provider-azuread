// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package groups

import (
	"context"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	groupBeta "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/group"
	memberBeta "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/member"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/groups/parse"
)

func groupMemberResource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		CreateContext: groupMemberResourceCreate,
		ReadContext:   groupMemberResourceRead,
		DeleteContext: groupMemberResourceDelete,

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(5 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Importer: pluginsdk.ImporterValidatingResourceId(func(id string) error {
			_, err := parse.GroupMemberID(id)
			return err
		}),

		Schema: map[string]*pluginsdk.Schema{
			"group_object_id": {
				Description:  "The object ID of the group you want to add the member to",
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsUUID,
			},

			"member_object_id": {
				Description:  "The object ID of the principal you want to add as a member to the group. Supported object types are Users, Groups or Service Principals",
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsUUID,
			},
		},
	}
}

func groupMemberResourceCreate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Groups.GroupClientBeta
	memberClient := meta.(*clients.Client).Groups.GroupMemberClientBeta

	id := beta.NewGroupIdMemberID(d.Get("group_object_id").(string), d.Get("member_object_id").(string))
	groupId := beta.NewGroupID(id.GroupId)
	resourceId := parse.NewGroupMemberID(id.GroupId, id.DirectoryObjectId)

	tf.LockByName(groupResourceName, id.GroupId)
	defer tf.UnlockByName(groupResourceName, id.GroupId)

	if resp, err := client.GetGroup(ctx, groupId, groupBeta.DefaultGetGroupOperationOptions()); err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return tf.ErrorDiagPathF(nil, "object_id", "%s was not found", groupId)
		}
		return tf.ErrorDiagPathF(err, "object_id", "Retrieving %s", groupId)
	}

	resp, err := memberClient.ListMembers(ctx, groupId, memberBeta.DefaultListMembersOperationOptions())
	if err != nil {
		return tf.ErrorDiagF(err, "Listing existing members for %s", groupId)
	}

	existingMembers := resp.Model
	if existingMembers != nil {
		for _, v := range *existingMembers {
			if strings.EqualFold(pointer.From(v.DirectoryObject().Id), id.DirectoryObjectId) {
				return tf.ImportAsExistsDiag("azuread_group_member", resourceId.String())
			}
		}
	}

	memberId := beta.NewDirectoryObjectID(id.DirectoryObjectId)

	memberRef := beta.ReferenceCreate{
		ODataId: pointer.To(client.Client.BaseUri + memberId.ID()),
	}

	if _, err = memberClient.AddMemberRef(ctx, groupId, memberRef, memberBeta.DefaultAddMemberRefOperationOptions()); err != nil {
		return tf.ErrorDiagF(err, "Adding %s", id)
	}

	d.SetId(resourceId.String())

	return groupMemberResourceRead(ctx, d, meta)
}

func groupMemberResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	//client := meta.(*clients.Client).Groups.GroupMemberClientBeta

	resourceId, err := parse.GroupMemberID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing Group Member ID %q", d.Id())
	}
	id := beta.NewGroupIdMemberID(resourceId.GroupId, resourceId.MemberId)

	// if member, err := groupGetMember(ctx, client, id); err != nil {
	// 	return tf.ErrorDiagF(err, "Retrieving member %q for group with object ID: %q", id.DirectoryObjectId, id.GroupId)
	// } else if member == nil {
	// 	log.Printf("[DEBUG] %s - removing from state", id)
	// 	d.SetId("")
	// 	return nil
	// }

	tf.Set(d, "group_object_id", id.GroupId)
	tf.Set(d, "member_object_id", id.DirectoryObjectId)

	return nil
}

func groupMemberResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Groups.GroupMemberClientBeta

	resourceId, err := parse.GroupMemberID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing Group Member ID %q", d.Id())
	}
	id := beta.NewGroupIdMemberID(resourceId.GroupId, resourceId.MemberId)

	tf.LockByName(groupResourceName, id.GroupId)
	defer tf.UnlockByName(groupResourceName, id.GroupId)

	if _, err := client.RemoveMemberRef(ctx, id, memberBeta.DefaultRemoveMemberRefOperationOptions()); err != nil {
		return tf.ErrorDiagF(err, "Removing %s", id)
	}

	// Wait for membership link to be deleted
	// if err := consistency.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
	// 	if member, err := groupGetMember(ctx, client, id); err != nil {
	// 		return nil, err
	// 	} else if member == nil {
	// 		return pointer.To(false), nil
	// 	}
	// 	return pointer.To(true), nil
	// }); err != nil {
	// 	return tf.ErrorDiagF(err, "Waiting for removal of %s", id)
	// }

	return nil
}
