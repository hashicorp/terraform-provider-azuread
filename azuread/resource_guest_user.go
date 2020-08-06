package azuread

import (
	"fmt"
	"log"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/graph/1.0/graph"
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
			"id": {
				Type:         schema.TypeString,
				Computed:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},
		},
	}
}
func resourceGuestUserCreate(d *schema.ResourceData, meta interface{}) error {
	guestUserClient := meta.(*ArmClient).guestUsersClient
	ctx := meta.(*ArmClient).StopContext
	// send invitation : guestUserClient.Create

	//get back user to set resource guestUserClient.Get mybe use resourceGuestUserRead
}

func resourceGuestUserUpdate(d *schema.ResourceData, meta interface{}) error {
	guestUserClient := meta.(*ArmClient).guestUsersClient
	ctx := meta.(*ArmClient).StopContext
}

func resourceGuestUserRead(d *schema.ResourceData, meta interface{}) error {
	guestUserClient := meta.(*ArmClient).guestUsersClient
	ctx := meta.(*ArmClient).StopContext
}

func resourceGuestUserDelete(d *schema.ResourceData, meta interface{}) error {
	guestUserClient := meta.(*ArmClient).guestUsersClient
	ctx := meta.(*ArmClient).StopContext
}
