package azuread

import (
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
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
	tenantID := client.TenantID

	memberGraphURL := fmt.Sprintf("https://graph.windows.net/%s/directoryObjects/%s", tenantID, memberID)

	properties := graphrbac.GroupAddMemberParameters{
		URL: &memberGraphURL,
	}

	if _, err := client.AddMember(ctx, groupID, properties); err != nil {
		return err
	}

	id := fmt.Sprintf("%s/%s", groupID, memberID)
	d.SetId(id)

	return resourceGroupMemberRead(d, meta)
}

func resourceGroupMemberRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient).groupsClient
	ctx := meta.(*ArmClient).StopContext

	id := strings.Split(d.Id(), "/")
	if len(id) != 2 {
		return fmt.Errorf("ID should be in the format {groupObjectId}/{memberObjectId} - but got %q", d.Id())
	}

	groupID := id[0]
	memberID := id[1]

	members, err := client.GetGroupMembersComplete(ctx, groupID)
	if err != nil {
		return fmt.Errorf("Error retrieving Azure AD Group members (groupObjectId: %q): %+v", groupID, err)
	}

	var memberObjectID string
	for members.NotDone() {
		// possible members are users, groups or service principals
		// we try to 'cast' each result as the corresponding type and diff
		// if we found the object we're looking for
		user, _ := members.Value().AsUser()
		if user != nil {
			if *user.ObjectID == memberID {
				memberObjectID = *user.ObjectID
				// we successfully found the directory object we're looking for, we can stop looping
				// through the results
				break
			}
		}

		group, _ := members.Value().AsADGroup()
		if group != nil {
			if *group.ObjectID == memberID {
				memberObjectID = *group.ObjectID
				// we successfully found the directory object we're looking for, we can stop looping
				// through the results
				break
			}
		}

		servicePrincipal, _ := members.Value().AsServicePrincipal()
		if servicePrincipal != nil {
			if *servicePrincipal.ObjectID == memberID {
				memberObjectID = *servicePrincipal.ObjectID
				// we successfully found the directory object we're looking for, we can stop looping
				// through the results
				break
			}
		}

		if err = members.NextWithContext(ctx); err != nil {
			return fmt.Errorf("Error listing Azure AD Group Members: %s", err)
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

	id := strings.Split(d.Id(), "/")
	if len(id) != 2 {
		return fmt.Errorf("ID should be in the format {groupObjectId}/{memberObjectId} - but got %q", d.Id())
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
