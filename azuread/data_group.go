package azuread

import (
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/validate"
)

func dataGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceActiveDirectoryGroupRead,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.NoEmptyStrings,
			},
		},
	}
}

func dataSourceActiveDirectoryGroupRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient).groupsClient
	ctx := meta.(*ArmClient).StopContext

	var adgroup graphrbac.ADGroup
	var groupObj *graphrbac.ADGroup

	// use the name to find the Azure AD group
	name := d.Get("name").(string)
	filter := fmt.Sprintf("displayName eq '%s'", name)
	log.Printf("[DEBUG] Using filter %q", filter)

	resp, err := client.ListComplete(ctx, filter)
	if err != nil {
		return fmt.Errorf("Error listing Azure AD groups: %+v", err)
	}

	for _, v := range *resp.Response().Value {
		if v.DisplayName == nil {
			continue
		}

		if *v.DisplayName == name {
			log.Printf("[DEBUG] %q (API result) matches %q (given value)", *v.DisplayName, name)
			groupObj = &v
			break
		} else {
			log.Printf("[DEBUG] %q (API result) does not match %q (given value)", *v.DisplayName, name)
		}
	}

	if groupObj == nil {
		return fmt.Errorf("Couldn't locate a Azure AD group with a name of %q", name)
	}

	adgroup = *groupObj

	d.SetId(*adgroup.ObjectID)

	return nil
}
