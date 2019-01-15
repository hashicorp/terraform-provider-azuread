package azuread

import (
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/ar"
)

func dataSourceArmActiveDirectoryGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceArmActiveDirectoryGroupRead,

		Schema: map[string]*schema.Schema{
			"object_id": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"display_name", "display_name"},
			},

			"display_name": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"object_id", "object_id"},
			},

			"mail": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func dataSourceArmActiveDirectoryGroupRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient).groupsClient
	ctx := meta.(*ArmClient).StopContext

	var group *graphrbac.ADGroup

	if v, ok := d.GetOk("object_id"); ok {
		objectId := v.(string)
		result, err := client.Get(ctx, objectId)
		if err != nil {
			if ar.ResponseWasNotFound(result.Response) {
				return fmt.Errorf("Group with Object ID %q was not found!", objectId)
			}

			return fmt.Errorf("Error retrieving Group with Object ID %q: %+v", objectId, err)
		}

		group = &result

	} else if v, ok := d.GetOk("display_name"); ok {
		displayName := v.(string)
		filter := fmt.Sprintf("displayName eq '%s'", displayName)
		log.Printf("[DEBUG] [data_source_azuread_group] Using filter %q", filter)

		results, err := client.ListComplete(ctx, filter)
		if err != nil {
			return fmt.Errorf("Error listing Groups: %+v", err)
		}

		for _, result := range *results.Response().Value {
			if result.DisplayName == nil {
				continue
			}

			if *result.DisplayName == displayName {
				group = &result
				break
			}
		}

		if group == nil {
			return fmt.Errorf("A Group with the Display Name %q was not found", displayName)
		}
	}

	d.SetId(*group.ObjectID)
	d.Set("display_name", group.DisplayName)
	d.Set("object_id", group.ObjectID)
	d.Set("mail", group.Mail)

	return nil
}
