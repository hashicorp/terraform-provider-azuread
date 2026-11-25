// Copyright IBM Corp. 2019, 2026
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
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
)

func namedLocationDataSource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		Description: "Gets information about a Named Location within Azure Active Directory.",
		ReadContext: namedLocationDataSourceRead,

		Timeouts: &pluginsdk.ResourceTimeout{
			Read: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*pluginsdk.Schema{
			"display_name": {
				Description:  "Specifies the display named of the named location to look up.",
				Type:         pluginsdk.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"ip": {
				Type:     pluginsdk.TypeList,
				Computed: true,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"ip_ranges": {
							Description: "List of IP address ranges in IPv4 CIDR format (e.g. `1.2.3.4/32`) or any allowable IPv6 format from IETF RFC596.",
							Type:        pluginsdk.TypeList,
							Computed:    true,
							Elem: &pluginsdk.Schema{
								Type: pluginsdk.TypeString,
							},
						},

						"trusted": {
							Description: "Whether the named location is trusted.",
							Type:        pluginsdk.TypeBool,
							Computed:    true,
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
							Description: "List of countries and/or regions in two-letter format specified by ISO 3166-2.",
							Type:        pluginsdk.TypeList,
							Computed:    true,
							Elem: &pluginsdk.Schema{
								Type: pluginsdk.TypeString,
							},
						},

						"include_unknown_countries_and_regions": {
							Description: "Whether IP addresses that don't map to a country or region are included in the named location.",
							Type:        pluginsdk.TypeBool,
							Computed:    true,
						},

						"country_lookup_method": {
							Description: "Method of detecting country the user is located in.",
							Type:        pluginsdk.TypeString,
							Computed:    true,
						},
					},
				},
			},

			"object_id": {
				Description: "The object ID of the named location.",
				Type:        pluginsdk.TypeString,
				Computed:    true,
			},
		},
	}
}

func namedLocationDataSourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).ConditionalAccess.NamedLocationClient

	displayName := d.Get("display_name").(string)
	options := conditionalaccessnamedlocation.ListConditionalAccessNamedLocationsOperationOptions{
		Filter: pointer.To(fmt.Sprintf("displayName eq '%s'", odata.EscapeSingleQuote(displayName))),
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

		tf.Set(d, "display_name", pointer.From(namedLocation.DisplayName))
		tf.Set(d, "object_id", pointer.From(namedLocation.Id))
		tf.Set(d, "ip", flattenIPNamedLocation(&namedLocation))

	case stable.CountryNamedLocation:
		if namedLocation.Id == nil {
			return tf.ErrorDiagF(errors.New("ID is nil for returned Country Named Location"), "Bad API response")
		}

		tf.Set(d, "display_name", pointer.From(namedLocation.DisplayName))
		tf.Set(d, "object_id", pointer.From(namedLocation.Id))
		tf.Set(d, "country", flattenCountryNamedLocation(&namedLocation))
	}

	id := stable.NewIdentityConditionalAccessNamedLocationID(pointer.From(item.NamedLocation().Id))
	d.SetId(id.ID())

	return nil
}
