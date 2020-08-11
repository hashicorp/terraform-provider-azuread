package azuread

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/services/graph/1.0/graph"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/ar"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/validate"
	"log"
	"strings"
)

func guestUserFields() []string {
	return []string{"userPrincipalName", "id", "mail", "displayName"}
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
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			},
			"display_name": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},
			"mail": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.StringIsEmailAddress,
			},
		},
	}
}
func resourceGuestUserCreate(d *schema.ResourceData, meta interface{}) error {
	guestUserClient := meta.(*ArmClient).guestUsersClient
	ctx := meta.(*ArmClient).StopContext
	// send invitation : guestUserClient.Create
	email := d.Get("mail").(string)
	username := d.Get("display_name").(string)
	invitationObj := graph.InvitationSended{InvitedUserEmailAddress: to.StringPtr(email), InviteRedirectURL: to.StringPtr("https://portal.azure.com"), SendInvitationMessage: to.BoolPtr(true), InvitedUserDisplayName: to.StringPtr(username)}
	invitation, err := guestUserClient.Create(ctx, invitationObj)
	if err != nil {
		log.Printf("[DEBUG] Error Creating Guest User : %+v", err)
		return nil
	}
	d.SetId(*invitation.InvitedUser.ID)
	return resourceGuestUserRead(d, meta)
}

func resourceGuestUserUpdate(d *schema.ResourceData, meta interface{}) error {
	guestUserClient := meta.(*ArmClient).guestUsersClient
	ctx := meta.(*ArmClient).StopContext

	guestUser := graph.GuestUser{
		DisplayName:       to.StringPtr(d.Get("display_name").(string)),
		UserPrincipalName: to.StringPtr(d.Get("user_principal_name").(string)),
		Mail:              to.StringPtr(d.Get("mail").(string)),
	}

	resp, err := guestUserClient.Update(ctx, guestUser, d.Id())
	if err != nil {
		log.Printf("[DEBUG] Error updating Guest User with ID %q: %+v \n %+v", d.Id(), err, resp)
		return nil
	}
	return resourceGuestUserRead(d, meta)
}

func resourceGuestUserRead(d *schema.ResourceData, meta interface{}) error {
	guestUserClient := meta.(*ArmClient).guestUsersClient
	ctx := meta.(*ArmClient).StopContext
	objectID := d.Id()
	user, err := guestUserClient.Get(ctx, objectID, strings.Join(guestUserFields(), ","))
	if err != nil {
		if ar.ResponseWasNotFound(user.Response) {
			log.Printf("[DEBUG] Guest User with Object ID %q was not found - removing from state!", objectID)
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error retrieving Guest User with ID %q: %+v", objectID, err)
	}
	d.Set("user_principal_name", *user.UserPrincipalName)
	d.Set("display_name", *user.DisplayName)
	d.Set("mail", *user.Mail)
	return nil
}

func resourceGuestUserDelete(d *schema.ResourceData, meta interface{}) error {
	guestUserClient := meta.(*ArmClient).guestUsersClient
	ctx := meta.(*ArmClient).StopContext
	resp, err := guestUserClient.Delete(ctx, d.Id())
	if err != nil {
		log.Printf("[DEBUG] Error deleting Guest User with ID %q: %+v \n %+v", d.Id(), err, resp)
		return nil
	}
	return nil
}
