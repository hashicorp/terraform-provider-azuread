package serviceprincipals

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
)

func clientConfigDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: clientConfigDataSourceRead,

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

func clientConfigDataSourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client)

	if client.AuthenticatedAsAServicePrincipal {
		spClient := client.ServicePrincipals.AadClient
		// Application & Service Principal is 1:1 per tenant. Since we know the appId (client_id)
		// here, we can query for the Service Principal whose appId matches.
		filter := fmt.Sprintf("appId eq '%s'", client.ClientID)
		result, err := spClient.List(ctx, filter)

		if err != nil {
			return tf.ErrorDiagF(err, "Listing Service Principals")
		}

		if result.Values() == nil || len(result.Values()) != 1 {
			return tf.ErrorDiagF(fmt.Errorf("%#v", result.Values()), "Unexpected Service Principal query result")
		}
	}

	d.SetId(fmt.Sprintf("%s-%s-%s", client.TenantID, client.ObjectID, client.ClientID))

	tf.Set(d, "client_id", client.ClientID)
	tf.Set(d, "object_id", client.ObjectID)
	tf.Set(d, "tenant_id", client.TenantID)

	return nil
}
