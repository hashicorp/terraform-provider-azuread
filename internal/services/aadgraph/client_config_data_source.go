package aadgraph

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/terraform-providers/terraform-provider-azuread/internal/clients"
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
			return diag.Diagnostics{diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Listing Service Principals",
				Detail:   err.Error(),
			}}
		}

		if result.Values() == nil || len(result.Values()) != 1 {
			return diag.Diagnostics{diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unexpected Service Principal query result",
				Detail:   fmt.Sprintf("%#v", result.Values()),
			}}
		}
	}

	d.SetId(fmt.Sprintf("%s-%s-%s", client.TenantID, client.ObjectID, client.ClientID))

	if err := d.Set("client_id", client.ClientID); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Could not set attribute",
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "client_id"}},
		}}
	}

	if err := d.Set("object_id", client.ObjectID); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Could not set attribute",
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "object_id"}},
		}}
	}

	if err := d.Set("tenant_id", client.TenantID); err != nil {
		return diag.Diagnostics{diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "Could not set attribute",
			Detail:        err.Error(),
			AttributePath: cty.Path{cty.GetAttrStep{Name: "tenant_id"}},
		}}
	}

	return nil
}
