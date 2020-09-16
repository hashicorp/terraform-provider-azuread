package aadgraph

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/terraform-providers/terraform-provider-azuread/internal/clients"
	"github.com/terraform-providers/terraform-provider-azuread/internal/services/aadgraph/graph"
)

func groupAvailabilityData() *schema.Resource {
	return &schema.Resource{
		Read: groupAvailabilityDataRead,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.NoZeroValues,
			},
		},
	}
}

func groupAvailabilityDataRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AadClient).AadGraph.GroupsClient
	ctx := meta.(*clients.AadClient).StopContext

	name := d.Get("name").(string)

	existingGroup, err := graph.GroupFindByName(ctx, client, name)
	if err != nil {
		return err
	}
	if existingGroup != nil {
		return fmt.Errorf("existing Group with name %q (ID: %q) was found", name, *existingGroup.ObjectID)
	}

	d.SetId(fmt.Sprintf("%s:availability", name))

	return nil
}
