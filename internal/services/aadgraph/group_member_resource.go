package aadgraph

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/terraform-providers/terraform-provider-azuread/internal/clients"
	"github.com/terraform-providers/terraform-provider-azuread/internal/services/aadgraph/graph"
	"github.com/terraform-providers/terraform-provider-azuread/internal/tf"
	"github.com/terraform-providers/terraform-provider-azuread/internal/validate"
)

const groupMemberResourceName = "azuread_group_member"

func groupMemberResource() *schema.Resource {
	return &schema.Resource{
		Create: groupMemberResourceCreate,
		Read:   groupMemberResourceRead,
		Delete: groupMemberResourceDelete,

		Importer: tf.ValidateResourceIDPriorToImport(func(id string) error {
			_, err := graph.ParseGroupMemberId(id)
			return err
		}),

		Schema: map[string]*schema.Schema{
			"group_object_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.UUID,
			},

			"member_object_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.UUID,
			},
		},
	}
}

func groupMemberResourceCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AadClient).AadGraph.GroupsClient
	ctx := meta.(*clients.AadClient).StopContext

	groupID := d.Get("group_object_id").(string)
	memberID := d.Get("member_object_id").(string)

	id := graph.GroupMemberIdFrom(groupID, memberID)

	tf.LockByName(groupMemberResourceName, groupID)
	defer tf.UnlockByName(groupMemberResourceName, groupID)

	existingMembers, err := graph.GroupAllMembers(ctx, client, groupID)
	if err != nil {
		return fmt.Errorf("listing existing members for group with ID %q", groupID)
	}
	if len(existingMembers) > 0 {
		for _, v := range existingMembers {
			if strings.EqualFold(v, memberID) {
				return tf.ImportAsExistsError("azuread_group_member", id.String())
			}
		}
	}

	if err := graph.GroupAddMember(ctx, client, groupID, memberID); err != nil {
		return err
	}

	d.SetId(id.String())
	return groupMemberResourceRead(d, meta)
}

func groupMemberResourceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AadClient).AadGraph.GroupsClient
	ctx := meta.(*clients.AadClient).StopContext

	id, err := graph.ParseGroupMemberId(d.Id())
	if err != nil {
		return fmt.Errorf("unable to parse ID: %v", err)
	}

	members, err := graph.GroupAllMembers(ctx, client, id.GroupId)
	if err != nil {
		return fmt.Errorf("retrieving Group members (groupObjectId: %q): %+v", id.GroupId, err)
	}

	var memberObjectID string
	for _, objectID := range members {
		if strings.EqualFold(objectID, id.MemberId) {
			memberObjectID = objectID
			break
		}
	}

	if memberObjectID == "" {
		d.SetId("")
		return nil
	}

	d.Set("group_object_id", id.GroupId)
	d.Set("member_object_id", memberObjectID)

	return nil
}

func groupMemberResourceDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AadClient).AadGraph.GroupsClient
	ctx := meta.(*clients.AadClient).StopContext

	id, err := graph.ParseGroupMemberId(d.Id())
	if err != nil {
		return fmt.Errorf("Unable to parse ID: %v", err)
	}

	tf.LockByName(groupMemberResourceName, id.GroupId)
	defer tf.UnlockByName(groupMemberResourceName, id.GroupId)

	if err := graph.GroupRemoveMember(ctx, client, d.Timeout(schema.TimeoutDelete), id.GroupId, id.MemberId); err != nil {
		return err
	}

	if _, err := graph.WaitForListRemove(id.MemberId, func() ([]string, error) {
		return graph.GroupAllMembers(ctx, client, id.GroupId)
	}); err != nil {
		return fmt.Errorf("waiting for group membership removal: %+v", err)
	}

	return nil
}
