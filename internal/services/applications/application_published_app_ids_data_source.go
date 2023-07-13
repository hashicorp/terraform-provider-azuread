// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applications

import (
	"context"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/environments"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
)

func applicationPublishedAppIdsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: func(_ context.Context, d *schema.ResourceData, _ interface{}) diag.Diagnostics {
			tf.Set(d, "result", environments.PublishedApis)
			d.SetId("appIds")
			return nil
		},

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"result": {
				Description: "A mapping of application names and application IDs",
				Type:        schema.TypeMap,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}
