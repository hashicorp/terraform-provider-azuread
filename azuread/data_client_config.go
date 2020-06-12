package azuread

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataClientConfig() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceArmClientConfigRead,

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

			"subscription_id": {
				Type:       schema.TypeString,
				Computed:   true,
				Deprecated: fmt.Sprintf("The %q attribute will be removed in version 1.0 of the provider. If you are using this attribute, you should instead use the %q data source from the AzureRM provider", "subscription_id", "azurerm_client_config"),
			},

			"object_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceArmClientConfigRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*ArmClient)
	ctx := meta.(*ArmClient).StopContext

	if client.authenticatedAsAServicePrincipal {
		spClient := client.servicePrincipalsClient
		// Application & Service Principal is 1:1 per tenant. Since we know the appId (client_id)
		// here, we can query for the Service Principal whose appId matches.
		filter := fmt.Sprintf("appId eq '%s'", client.clientID)
		listResult, listErr := spClient.List(ctx, filter)

		if listErr != nil {
			return fmt.Errorf("Error listing Service Principals: %#v", listErr)
		}

		if listResult.Values() == nil || len(listResult.Values()) != 1 {
			return fmt.Errorf("Unexpected Service Principal query result: %#v", listResult.Values())
		}
	}

	d.SetId(time.Now().UTC().String())
	d.Set("client_id", client.clientID)
	d.Set("object_id", client.objectID)
	d.Set("tenant_id", client.tenantID)

	// TODO: remove in v1.0
	if client.subscriptionID != client.tenantID {
		d.Set("subscription_id", client.subscriptionID)
	}

	return nil
}
