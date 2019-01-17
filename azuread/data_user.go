package azuread

import (
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/ar"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/validate"
)

func dataSourceUser() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceUserRead,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"user_principal_name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},

			"mail": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceUserRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient).usersClient
	ctx := meta.(*ArmClient).StopContext

	var user graphrbac.User
	var queryString string

	queryString = d.Get("user_principal_name").(string)

	log.Printf("[DEBUG] Using Get with the following query string: %q", queryString)
	resp, err := client.Get(ctx, queryString)
	if err != nil {
		if ar.ResponseWasNotFound(resp.Response) {
			return fmt.Errorf("Error: No AzureAD User found with the following query string: %q", queryString)
		}
		return fmt.Errorf("Error making Read request on AzureAD User the following query string: %q: %+v", queryString, err)
	}

	user = resp

	d.SetId(*user.ObjectID)
	d.Set("user_principal_name", user.UserPrincipalName)
	d.Set("mail", user.Mail)

	return nil
}
