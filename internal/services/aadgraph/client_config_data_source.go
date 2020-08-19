package aadgraph

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/terraform-providers/terraform-provider-azuread/internal/clients"
)

func ClientConfigData() *schema.Resource {
	return &schema.Resource{
		Read: clientConfigDataRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"client_id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"tenant_id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"object_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func clientConfigDataRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AadClient)
	ctx := meta.(*clients.AadClient).StopContext

	if client.AuthenticatedAsAServicePrincipal {
		spClient := client.AadGraph.ServicePrincipalsClient
		// Application & Service Principal is 1:1 per tenant. Since we know the appId (client_id)
		// here, we can query for the Service Principal whose appId matches.
		filter := fmt.Sprintf("appId eq '%s'", client.ClientID)
		listResult, listErr := spClient.List(ctx, filter)

		if listErr != nil {
			return fmt.Errorf("listing Service Principals: %#v", listErr)
		}

		if listResult.Values() == nil || len(listResult.Values()) != 1 {
			return fmt.Errorf("unexpected Service Principal query result: %#v", listResult.Values())
		}
	}

	d.SetId(fmt.Sprintf("%s-%s-%s", client.TenantID, client.ObjectID, client.ClientID))
	d.Set("client_id", client.ClientID)
	d.Set("object_id", client.ObjectID)
	d.Set("tenant_id", client.TenantID)

	return nil
}
