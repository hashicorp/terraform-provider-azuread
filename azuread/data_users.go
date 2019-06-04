package azuread

import (
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/ar"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/validate"
)

func dataSourceUsers() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceUsersRead,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"object_ids": {
				Type:          schema.TypeList,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"user_principal_names"},
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validate.UUID,
				},
			},

			"user_principal_names": {
				Type:          schema.TypeList,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"object_ids"},
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validate.NoEmptyStrings,
				},
			},
		},
	}
}

func dataSourceUsersRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient).usersClient
	ctx := meta.(*ArmClient).StopContext

	var user graphrbac.User

	if oId, ok := d.GetOk("user_principal_name"); ok {
		// use the object_id to find the Azure AD application
		resp, err := client.Get(ctx, oId.(string))
		if err != nil {
			if ar.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("Error: AzureAD User with ID %q was not found", oId.(string))
			}

			return fmt.Errorf("Error making Read request on AzureAD User with ID %q: %+v", oId.(string), err)
		}

		user = resp
	} else if name, ok := d.Get("object_id").(string); ok {
		filter := fmt.Sprintf("objectId eq '%s'", name)

		resp, err := client.ListComplete(ctx, filter)
		if err != nil {
			return fmt.Errorf("Error listing Azure AD Users for filter %q: %+v", filter, err)
		}

		values := resp.Response().Value
		if values == nil {
			return fmt.Errorf("nil values for AD Users matching %q", filter)
		}
		if len(*values) == 0 {
			return fmt.Errorf("Found no AD Users matching %q", filter)
		}
		if len(*values) > 2 {
			return fmt.Errorf("Found multiple AD Users matching %q", filter)
		}

		user = (*values)[0]
		if user.DisplayName == nil {
			return fmt.Errorf("nil DisplayName for AD Users matching %q", filter)
		}
		if *user.DisplayName != name {
			return fmt.Errorf("displayname for AD Users matching %q does is does not match(%q!=%q)", filter, *user.DisplayName, name)
		}
	} else {
		return fmt.Errorf("one of `object_id` or `user_principal_name` must be supplied")
	}

	if user.ObjectID == nil {
		return fmt.Errorf("Group objectId is nil")
	}

	d.SetId(*user.ObjectID)
	d.SetId(*user.ObjectID)
	d.Set("object_id", user.ObjectID)


	return nil
}
