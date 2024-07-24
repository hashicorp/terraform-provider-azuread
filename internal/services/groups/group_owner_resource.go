package groups

import (
	"context"
	"errors"
	"fmt"
	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/manicminer/hamilton/msgraph"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/terraform-provider-azuread/internal/services/groups/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
)

func groupOwnerResource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		CreateContext: groupOwnerResourceCreate,
		ReadContext:   groupOwnerResourceRead,
		DeleteContext: groupOwnerResourceDelete,

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
				Description:      "The object ID of the group you want to add the owner to",
				Type:             pluginsdk.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validation.ValidateDiag(validation.IsUUID),
			},

			"owner_object_id": {
				Description:      "The object ID of the principal you want to add as an owner to the group. Supported object types are Users, Groups or Service Principals",
				Type:             pluginsdk.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validation.ValidateDiag(validation.IsUUID),
			},
		},
	}
}

func groupOwnerResourceCreate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Groups.GroupsClient
	directoryObjectsClient := meta.(*clients.Client).Groups.DirectoryObjectsClient
	tenantId := meta.(*clients.Client).TenantID
	groupId := d.Get("group_object_id").(string)
	ownerId := d.Get("owner_object_id").(string)

	id := parse.NewGroupOwnerID(groupId, ownerId)

	tf.LockByName(groupResourceName, id.GroupId)
	defer tf.UnlockByName(groupResourceName, id.GroupId)

	group, status, err := client.Get(ctx, groupId, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return tf.ErrorDiagPathF(nil, "object_id", "Group with object ID %q not found", groupId)
		}
		return tf.ErrorDiagPathF(err, "object_id", "Error reading group with object ID %q", groupId)
	}

	existingOwners, _, err := client.ListOwners(ctx, id.GroupId)
	if err != nil {
		return tf.ErrorDiagF(err, "Listing existing owners for group with object ID: %q", id.GroupId)
	}
	if existingOwners != nil {
		for _, v := range *existingOwners {
			if strings.EqualFold(v, ownerId) {
				return tf.ImportAsExistsDiag("azuread_group_owner", id.String())
			}
		}
	}

	ownerObject, _, err := directoryObjectsClient.Get(ctx, ownerId, odata.Query{})
	if err != nil {
		return tf.ErrorDiagF(err, "Could not retrieve principal object %q", ownerId)
	}
	if ownerObject == nil {
		return tf.ErrorDiagF(errors.New("returned ownerObject was nil"), "Could not retrieve owner principal object %q", ownerId)
	}

	ownerObject.ODataId = (*odata.Id)(pointer.To(fmt.Sprintf("%s/v1.0/%s/directoryObjects/%s",
		client.BaseClient.Endpoint, tenantId, ownerId)))

	group.Owners = &msgraph.Owners{*ownerObject}

	if _, err := client.AddOwners(ctx, group); err != nil {
		return tf.ErrorDiagF(err, "Adding group owner %q to group %q", ownerId, groupId)
	}

	d.SetId(id.String())
	return groupOwnerResourceRead(ctx, d, meta)
}

func groupOwnerResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Groups.GroupsClient

	id, err := parse.GroupOwnerID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Error parsing Group Owner ID %q", d.Id())
	}

	owners, status, err := client.ListOwners(ctx, id.GroupId)
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Group with ID %q was not found - removing group owner with ID %q from state", id.GroupId, d.Id())
			d.SetId("")
			return nil
		}
		return tf.ErrorDiagF(err, "Retrieving owners for group with object ID: %q", id.GroupId)
	}

	var ownerObjectId string
	if owners != nil {
		for _, objectId := range *owners {
			if strings.EqualFold(objectId, id.OwnerId) {
				ownerObjectId = objectId
				break
			}
		}
	}

	if ownerObjectId == "" {
		log.Printf("[DEBUG] Owner with ID %q was not found in Group %q - removing from state", id.OwnerId, id.GroupId)
		d.SetId("")
		return nil
	}

	tf.Set(d, "group_object_id", id.GroupId)
	tf.Set(d, "owner_object_id", ownerObjectId)

	return nil
}

func groupOwnerResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(clients.Client).Groups.GroupsClient

	id, err := parse.GroupOwnerID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing Group Owner ID %q", d.Id())
	}

	tf.LockByName(groupResourceName, id.GroupId)
	defer tf.UnlockByName(groupResourceName, id.GroupId)

	if _, err := client.RemoveMembers(ctx, id.GroupId, &[]string{id.OwnerId}); err != nil {
		return tf.ErrorDiagF(err, "Removing group owner %q from group %q", id.OwnerId, id.GroupId)
	}

	// wait for owner link to be deleted
	if err := helpers.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
		defer func() { client.BaseClient.DisableRetries = false }()
		client.BaseClient.DisableRetries = true
		if _, status, err := client.GetOwner(ctx, id.GroupId, id.OwnerId); err != nil {
			if status == http.StatusNotFound {
				return pointer.To(false), nil
			}
			return nil, err
		}
		return pointer.To(true), nil
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for removal of Owner %q from group with object ID %q", id.OwnerId, id.GroupId)
	}

	return nil
}
