package azuread

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/services/graph/1.0/graph"
	"log"
	"strings"

	"github.com/Azure/go-autorest/autorest/to"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/ar"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/graph"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/p"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/validate"
)

func guestUserFields() []string {
	return [...]string{"userPrincipalName", "id", "mail", "employeeID", "displayName"}
}
func resourceGuestUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceGuestUserCreate,
		Read:   resourceGuestUserRead,
		Update: resourceGuestUserUpdate,
		Delete: resourceGuestUserDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"user_principal_name": {
				Type:         schema.TypeString,
				Computed:     true,
				ValidateFunc: validate.StringIsEmailAddress,
			},
			"display_name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},
			"mail": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.StringIsEmailAddress,
			},
			"employee_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			//"id": {
			//	Type:         schema.TypeString,
			//	Computed:     true,
			//	ValidateFunc: validate.NoEmptyStrings,
			//},
		},
	}
}
func resourceGuestUserCreate(d *schema.ResourceData, meta interface{}) error {
	guestUserClient := meta.(*ArmClient).guestUsersClient
	ctx := meta.(*ArmClient).StopContext
	// send invitation : guestUserClient.Create
	email := d.Get("mail").(string)
	invitationObj := graph.InvitationSended{InvitedUserEmailAddress: to.StringPtr(email), InviteRedirectURL: to.StringPtr("https://portal.azure.com"), SendInvitationMessage: to.BoolPtr(true)}
	invitation, err := guestUserClient.Create(ctx, invitationObj)
	d.SetId(*invitation.InvitedUser.ID)
	//get back user to set resource guestUserClient.Get mybe use resourceGuestUserRead
	return resourceGuestUserRead(d, meta)
}

func resourceGuestUserUpdate(d *schema.ResourceData, meta interface{}) error {
	guestUserClient := meta.(*ArmClient).guestUsersClient
	ctx := meta.(*ArmClient).StopContext

	guestUser := graph.GuestUser{
		DisplayName:       to.StringPtr(d.Get("display_name")),
		UserPrincipalName: to.StringPtr(d.Get("user_principal_name")),
		Mail:              to.StringPtr(d.Get("mail")),
		EmployeeID:        to.StringPtr(d.Get("employee_id")),
	}

	resp, err := guestUserClient.Update(ctx, guestUser, d.Id())
	if err != nil {
		return fmt.Errorf("Error updating Guest User with ID %q: %+v", d.Id(), err)
	}
	return nil
}

func resourceGuestUserRead(d *schema.ResourceData, meta interface{}) error {
	guestUserClient := meta.(*ArmClient).guestUsersClient
	ctx := meta.(*ArmClient).StopContext
	objectID := d.Id()
	user, err := guestUserClient.Get(context.Background(), objectID, strings.join(guestUserFields(), ','))
	if err != nil {
		d.SetId("")
		return nil
	}
	d.Set("user_principal_name", user.UserPrincipalName)
	d.Set("display_name", user.DisplayName)
	d.Set("mail", user.Mail)
	d.Set("employee_id", user.EmployeeID)

	return nil
}

func resourceGuestUserDelete(d *schema.ResourceData, meta interface{}) error {
	guestUserClient := meta.(*ArmClient).guestUsersClient
	ctx := meta.(*ArmClient).StopContext
	resp, err := guestUserClient.Delete(ctx, d.Id())
	if err != nill {
		return fmt.Errorf("Error deleting Guest User with ID %q: %+v", d.Id(), err)
	}
	return nil
}
