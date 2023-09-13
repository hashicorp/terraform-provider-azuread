// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package conditionalaccess

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
	"github.com/manicminer/hamilton/msgraph"
)

func namedLocationDataSource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		ReadContext: namedLocationDataSourceRead,

		Timeouts: &pluginsdk.ResourceTimeout{
			Read: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*pluginsdk.Schema{
			"display_name": {
				Type:             pluginsdk.TypeString,
				Required:         true,
				ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
			},

			"ip": {
				Type:     pluginsdk.TypeList,
				Computed: true,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"ip_ranges": {
							Type:     pluginsdk.TypeList,
							Computed: true,
							Elem: &pluginsdk.Schema{
								Type: pluginsdk.TypeString,
							},
						},

						"trusted": {
							Type:     pluginsdk.TypeBool,
							Computed: true,
						},
					},
				},
			},

			"country": {
				Type:     pluginsdk.TypeList,
				Computed: true,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"countries_and_regions": {
							Type:     pluginsdk.TypeList,
							Computed: true,
							Elem: &pluginsdk.Schema{
								Type: pluginsdk.TypeString,
							},
						},

						"include_unknown_countries_and_regions": {
							Type:     pluginsdk.TypeBool,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func namedLocationDataSourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ConditionalAccess.NamedLocationsClient

	displayName := d.Get("display_name").(string)
	query := odata.Query{Filter: fmt.Sprintf("displayName eq '%s'", displayName)}
	result, status, err := client.List(ctx, query)
	if err != nil {
		if status == http.StatusNotFound {
			return tf.ErrorDiagPathF(nil, "display_name", "Named Location with display name %q was not found", displayName)
		}
	}
	if result == nil {
		return tf.ErrorDiagF(errors.New("Bad API response"), "Result is nil")
	}
	if len(*result) > 1 {
		return tf.ErrorDiagPathF(nil, "display_name", "More than one Named Location was found with display name %q", displayName)
	}

	location := (*result)[0]

	if ipnl, ok := location.(msgraph.IPNamedLocation); ok {
		if ipnl.ID == nil {
			return tf.ErrorDiagF(errors.New("Bad API response"), "ID is nil for returned IP Named Location")
		}
		d.SetId(*ipnl.ID)
		tf.Set(d, "display_name", ipnl.DisplayName)
		tf.Set(d, "ip", flattenIPNamedLocation(&ipnl))
	}

	if cnl, ok := location.(msgraph.CountryNamedLocation); ok {
		if cnl.ID == nil {
			return tf.ErrorDiagF(errors.New("Bad API response"), "ID is nil for returned Country Named Location")
		}
		d.SetId(*cnl.ID)
		tf.Set(d, "display_name", cnl.DisplayName)
		tf.Set(d, "country", flattenCountryNamedLocation(&cnl))
	}

	return nil
}
