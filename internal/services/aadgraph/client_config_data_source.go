package aadgraph

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/terraform-providers/terraform-provider-azuread/internal/clients"
	"github.com/terraform-providers/terraform-provider-azuread/internal/tf"
)

func clientConfigData() *schema.Resource {
	return &schema.Resource{
		ReadContext: clientConfigDataRead,

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

func clientConfigDataRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.AadClient)

	if client.AuthenticatedAsAServicePrincipal {
		spClient := client.AadGraph.ServicePrincipalsClient
		// Application & Service Principal is 1:1 per tenant. Since we know the appId (client_id)
		// here, we can query for the Service Principal whose appId matches.
		filter := fmt.Sprintf("appId eq '%s'", client.ClientID)
		result, err := spClient.List(ctx, filter)

		if err != nil {
			return tf.ErrorDiag("Listing Service Principals", err.Error(), "")
		}

		if result.Values() == nil || len(result.Values()) != 1 {
			return tf.ErrorDiag("Unexpected Service Principal query result", fmt.Sprintf("%#v", result.Values()), "")
		}
	}

	d.SetId(fmt.Sprintf("%s-%s-%s", client.TenantID, client.ObjectID, client.ClientID))

	if err := d.Set("client_id", client.ClientID); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "client_id")
	}

	if err := d.Set("object_id", client.ObjectID); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "object_id")
	}

	if err := d.Set("tenant_id", client.TenantID); err != nil {
		return tf.ErrorDiag("Could not set attribute", err.Error(), "tenant_id")
	}

	return nil
}
