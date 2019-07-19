package azuread

import (
	"fmt"
	"strings"

	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/graph"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/ar"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/validate"
)

func resourceGroupMember() *schema.Resource {
	return &schema.Resource{
		Create: resourceGroupMemberCreate,
		Read:   resourceGroupMemberRead,
		Delete: resourceGroupMemberDelete,
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

func resourceGroupMemberCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient).groupsClient
	ctx := meta.(*ArmClient).StopContext

	groupID := d.Get("group_object_id").(string)
	memberID := d.Get("member_object_id").(string)

	if err := graph.GroupAddMember(client, ctx, groupID, memberID); err != nil {
		return err
	}

	id := fmt.Sprintf("%s/member/%s", groupID, memberID)
	d.SetId(id)

	return resourceGroupMemberRead(d, meta)
}

func resourceGroupMemberRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient).groupsClient
	ctx := meta.(*ArmClient).StopContext

	id := strings.Split(d.Id(), "/member/")
	if len(id) != 2 {
		return fmt.Errorf("ID should be in the format {groupObjectId}/member/{memberObjectId} - but got %q", d.Id())
	}

	groupID := id[0]
	memberID := id[1]

	members, err := graph.GroupAllMembers(client, ctx, groupID)
	if err != nil {
		return fmt.Errorf("Error retrieving Azure AD Group members (groupObjectId: %q): %+v", groupID, err)
	}

	var memberObjectID string

	for _, objectID := range members {
		if objectID == memberID {
			memberObjectID = objectID
		}
	}

	if memberObjectID == "" {
		d.SetId("")
		return fmt.Errorf("Azure AD Group Member not found - groupObjectId:%q / memberObjectId:%q", groupID, memberID)
	}

	d.Set("group_object_id", groupID)
	d.Set("member_object_id", memberObjectID)

	return nil
}

func resourceGroupMemberDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient).groupsClient
	ctx := meta.(*ArmClient).StopContext

	id := strings.Split(d.Id(), "/member/")
	if len(id) != 2 {
		return fmt.Errorf("ID should be in the format {groupObjectId}/member/{memberObjectId} - but got %q", d.Id())
	}

	groupID := id[0]
	memberID := id[1]

	resp, err := client.RemoveMember(ctx, groupID, memberID)
	if err != nil {
		if !ar.ResponseWasNotFound(resp) {
			return fmt.Errorf("Error removing Member (memberObjectId: %q) from Azure AD Group (groupObjectId: %q): %+v", memberID, groupID, err)
		}
	}

	return nil
}
