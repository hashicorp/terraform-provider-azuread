// Copyright IBM Corp. 2019, 2026
// SPDX-License-Identifier: MPL-2.0

package applications

import (
	"context"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/environments"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
)

func applicationPublishedAppIdsDataSource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		Description: "Use this data source to discover application IDs for APIs published by Microsoft.\n\n" +
			"This data source uses an [unofficial source of application IDs](https://github.com/hashicorp/go-azure-sdk/blob/main/sdk/environments/application_ids.go), as there is currently no available official indexed source for applications or APIs published by Microsoft.\n\n" +
			"The app IDs returned by this data source are sourced from the Azure Global (Public) Cloud, however some of them are known to work in government and national clouds.",
		ReadContext: func(_ context.Context, d *pluginsdk.ResourceData, _ interface{}) pluginsdk.Diagnostics {
			tf.Set(d, "result", environments.PublishedApis)
			d.SetId("appIds")
			return nil
		},

		Timeouts: &pluginsdk.ResourceTimeout{
			Read: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*pluginsdk.Schema{
			"result": {
				Description: "A map of application names to application IDs.",
				Type:        pluginsdk.TypeMap,
				Computed:    true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},
		},
	}
}
