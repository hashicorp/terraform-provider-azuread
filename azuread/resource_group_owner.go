package azuread

import (
	"fmt"

	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/graph"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/tf"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/ar"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/validate"
)

const groupOwnberResourceName = "azuread_group_owner"

func resourceGroupOwner() *schema.Resource {
	return &schema.Resource{
		Create: resourceGroupOwnerCreate,
		Read:   resourceGroupOwnerRead,
		Delete: resourceGroupOwnerDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"group_object_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.UUID,
			},

			"owner_object_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.UUID,
			},
		},
	}
}

func resourceGroupOwnerCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient).groupsClient
	ctx := meta.(*ArmClient).StopContext

	groupID := d.Get("group_object_id").(string)
	ownerID := d.Get("owner_object_id").(string)

	tf.LockByName(groupOwnberResourceName, groupID)
	defer tf.UnlockByName(groupOwnberResourceName, groupID)

	if err := graph.GroupAddOwner(client, ctx, groupID, ownerID); err != nil {
		return err
	}

	d.SetId(graph.GroupOwnerIdFrom(groupID, ownerID).String())
	return resourceGroupOwnerRead(d, meta)
}

func resourceGroupOwnerRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient).groupsClient
	ctx := meta.(*ArmClient).StopContext

	id, err := graph.ParseGroupOwnerId(d.Id())
	if err != nil {
		return fmt.Errorf("Unable to parse ID: %v", err)
	}

	owners, err := graph.GroupAllOwners(client, ctx, id.GroupId)
	if err != nil {
		return fmt.Errorf("Error retrieving Azure AD Group owners (groupObjectId: %q): %+v", id.GroupId, err)
	}

	var ownerObjectID string
	for _, objectID := range owners {
		if objectID == id.OwnerId {
			ownerObjectID = objectID
		}
	}

	if ownerObjectID == "" {
		d.SetId("")
		return fmt.Errorf("Azure AD Group Owner not found - groupObjectId:%q / ownerObjectId:%q", id.GroupId, id.OwnerId)
	}

	d.Set("group_object_id", id.GroupId)
	d.Set("owner_object_id", id.OwnerId)

	return nil
}

func resourceGroupOwnerDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient).groupsClient
	ctx := meta.(*ArmClient).StopContext

	id, err := graph.ParseGroupOwnerId(d.Id())
	if err != nil {
		return fmt.Errorf("Unable to parse ID: %v", err)
	}

	tf.LockByName(groupOwnberResourceName, id.GroupId)
	defer tf.UnlockByName(groupOwnberResourceName, id.GroupId)

	resp, err := client.RemoveOwner(ctx, id.GroupId, id.OwnerId)
	if err != nil {
		if !ar.ResponseWasNotFound(resp) {
			return fmt.Errorf("Error removing Owner (ownerObjectId: %q) from Azure AD Group (groupObjectId: %q): %+v", id.OwnerId, id.GroupId, err)
		}
	}

	return nil
}
