package aadgraph

import (
	"fmt"
	"github.com/terraform-providers/terraform-provider-azuread/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/terraform-providers/terraform-provider-azuread/internal/clients"
	"github.com/terraform-providers/terraform-provider-azuread/internal/services/aadgraph/graph"
	"github.com/terraform-providers/terraform-provider-azuread/internal/tf"
	"github.com/terraform-providers/terraform-provider-azuread/internal/validate"
)

const groupMemberResourceName = "azuread_group_member"

func GroupMemberResource() *schema.Resource {
	return &schema.Resource{
		Create: groupMemberResourceCreate,
		Read:   groupMemberResourceRead,
		Delete: groupMemberResourceDelete,

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

	tf.LockByName(groupMemberResourceName, groupID)
	defer tf.UnlockByName(groupMemberResourceName, groupID)

	if err := graph.GroupAddMember(client, ctx, groupID, memberID); err != nil {
		return err
	}

	d.SetId(graph.GroupMemberIdFrom(groupID, memberID).String())
	return groupMemberResourceRead(d, meta)
}

func groupMemberResourceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AadClient).AadGraph.GroupsClient
	ctx := meta.(*clients.AadClient).StopContext

	id, err := graph.ParseGroupMemberId(d.Id())
	if err != nil {
		return fmt.Errorf("Unable to parse ID: %v", err)
	}

	members, err := graph.GroupAllMembers(client, ctx, id.GroupId)
	if err != nil {
		return fmt.Errorf("Error retrieving Azure AD Group members (groupObjectId: %q): %+v", id.GroupId, err)
	}

	var memberObjectID string
	for _, objectID := range members {
		if objectID == id.MemberId {
			memberObjectID = objectID
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

	resp, err := client.RemoveMember(ctx, id.GroupId, id.MemberId)
	if err != nil {
		if !utils.ResponseWasNotFound(resp) {
			return fmt.Errorf("Error removing Member (memberObjectId: %q) from Azure AD Group (groupObjectId: %q): %+v", id.MemberId, id.GroupId, err)
		}
	}

	return nil
}
