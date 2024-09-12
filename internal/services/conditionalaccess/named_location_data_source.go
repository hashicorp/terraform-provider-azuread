// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package conditionalaccess

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/conditionalaccessnamedlocation"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
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

func namedLocationDataSourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).ConditionalAccess.NamedLocationClient

	displayName := d.Get("display_name").(string)
	options := conditionalaccessnamedlocation.ListConditionalAccessNamedLocationsOperationOptions{
		Filter: pointer.To(fmt.Sprintf("displayName eq '%s'", displayName)),
	}
	resp, err := client.ListConditionalAccessNamedLocations(ctx, options)
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return tf.ErrorDiagPathF(nil, "display_name", "Named Location with display name %q was not found", displayName)
		}
	}

	namedLocations := resp.Model
	if namedLocations == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "Bad API Response")
	}
	if len(*namedLocations) == 0 {
		return tf.ErrorDiagPathF(nil, "display_name", "No Named Location was found with display name %q", displayName)
	}
	if len(*namedLocations) > 1 {
		return tf.ErrorDiagPathF(nil, "display_name", "More than one Named Location was found with display name %q", displayName)
	}

	item := (*namedLocations)[0]

	if item == nil {
		return tf.ErrorDiagF(errors.New("NamedLocation was nil"), "Bad API Response")
	}

	switch namedLocation := item.(type) {
	case stable.IPNamedLocation:
		if namedLocation.Id == nil {
			return tf.ErrorDiagF(errors.New("ID is nil for returned IP Named Location"), "Bad API response")
		}
		d.SetId(*namedLocation.Id)
		tf.Set(d, "display_name", namedLocation.DisplayName)
		tf.Set(d, "ip", flattenIPNamedLocation(&namedLocation))

	case stable.CountryNamedLocation:
		if namedLocation.Id == nil {
			return tf.ErrorDiagF(errors.New("ID is nil for returned Country Named Location"), "Bad API response")
		}
		d.SetId(*namedLocation.Id)
		tf.Set(d, "display_name", namedLocation.DisplayName)
		tf.Set(d, "country", flattenCountryNamedLocation(&namedLocation))
	}

	return nil
}
