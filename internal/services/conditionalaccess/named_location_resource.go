// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package conditionalaccess

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/conditionalaccessnamedlocation"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/consistency"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/conditionalaccess/migrations"
)

func namedLocationResource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		CreateContext: namedLocationResourceCreate,
		ReadContext:   namedLocationResourceRead,
		UpdateContext: namedLocationResourceUpdate,
		DeleteContext: namedLocationResourceDelete,

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(5 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Update: pluginsdk.DefaultTimeout(5 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Importer: pluginsdk.ImporterValidatingResourceId(func(id string) error {
			if _, errs := stable.ValidateIdentityConditionalAccessNamedLocationID(id, "id"); len(errs) > 0 {
				out := ""
				for _, err := range errs {
					out += err.Error()
				}
				return fmt.Errorf(out)
			}
			return nil
		}),

		SchemaVersion: 1,
		StateUpgraders: []pluginsdk.StateUpgrader{
			{
				Type:    migrations.ResourceNamedLocationInstanceResourceV0().CoreConfigSchema().ImpliedType(),
				Upgrade: migrations.ResourceNamedLocationInstanceStateUpgradeV0,
				Version: 0,
			},
		},

		Schema: map[string]*pluginsdk.Schema{
			"display_name": {
				Type:         pluginsdk.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"ip": {
				Type:         pluginsdk.TypeList,
				Optional:     true,
				ForceNew:     true,
				MaxItems:     1,
				ExactlyOneOf: []string{"ip", "country"},
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"ip_ranges": {
							Type:     pluginsdk.TypeList,
							Required: true,
							Elem: &pluginsdk.Schema{
								Type:         pluginsdk.TypeString,
								ValidateFunc: validation.PrefixLengthAtLeast(8),
							},
						},

						"trusted": {
							Type:     pluginsdk.TypeBool,
							Optional: true,
						},
					},
				},
			},

			"country": {
				Type:         pluginsdk.TypeList,
				Optional:     true,
				ForceNew:     true,
				MaxItems:     1,
				ExactlyOneOf: []string{"ip", "country"},
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"countries_and_regions": {
							Type:     pluginsdk.TypeList,
							Required: true,
							Elem: &pluginsdk.Schema{
								Type:         pluginsdk.TypeString,
								ValidateFunc: validation.StringIsNotEmpty,
							},
						},

						"include_unknown_countries_and_regions": {
							Type:     pluginsdk.TypeBool,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func namedLocationResourceCreate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).ConditionalAccess.NamedLocationClient

	if v, ok := d.GetOk("ip"); ok {
		properties := expandIPNamedLocation(v.([]interface{}))
		properties.DisplayName = pointer.To(d.Get("display_name").(string))

		resp, err := client.CreateConditionalAccessNamedLocation(ctx, *properties, conditionalaccessnamedlocation.DefaultCreateConditionalAccessNamedLocationOperationOptions())
		if err != nil {
			return tf.ErrorDiagF(err, "Could not create named location")
		}

		if resp.Model == nil {
			return tf.ErrorDiagF(errors.New("returned model was nil"), "Bad API Response")
		}

		namedLocation, ok := resp.Model.(stable.IPNamedLocation)
		if !ok {
			return tf.ErrorDiagF(errors.New("returned model was not an IPNamedLocation"), "Bad API response")
		}

		if namedLocation.Id == nil {
			return tf.ErrorDiagF(errors.New("nil/empty object ID returned for named location"), "Bad API response")
		}

		id := stable.NewIdentityConditionalAccessNamedLocationID(*namedLocation.Id)
		d.SetId(id.ID())

	} else if v, ok = d.GetOk("country"); ok {
		properties := expandCountryNamedLocation(v.([]interface{}))
		properties.DisplayName = pointer.To(d.Get("display_name").(string))

		resp, err := client.CreateConditionalAccessNamedLocation(ctx, *properties, conditionalaccessnamedlocation.DefaultCreateConditionalAccessNamedLocationOperationOptions())
		if err != nil {
			return tf.ErrorDiagF(err, "Could not create named location")
		}

		if resp.Model == nil {
			return tf.ErrorDiagF(errors.New("returned model was nil"), "Bad API Response")
		}

		namedLocation, ok := resp.Model.(stable.CountryNamedLocation)
		if !ok {
			return tf.ErrorDiagF(errors.New("returned model was not a CountryNamedLocation"), "Bad API response")
		}

		if namedLocation.Id == nil {
			return tf.ErrorDiagF(errors.New("nil/empty object ID returned for named location"), "Bad API response")
		}

		id := stable.NewIdentityConditionalAccessNamedLocationID(*namedLocation.Id)
		d.SetId(id.ID())

	} else {
		return tf.ErrorDiagF(errors.New("one of `ip` or `country` must be specified"), "Unable to determine named location type")
	}

	return namedLocationResourceRead(ctx, d, meta)
}

func namedLocationResourceUpdate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).ConditionalAccess.NamedLocationClient

	id, err := stable.ParseIdentityConditionalAccessNamedLocationID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing Named Location ID")
	}

	if v, ok := d.GetOk("ip"); ok {
		properties := expandIPNamedLocation(v.([]interface{}))

		if d.HasChange("display_name") {
			properties.DisplayName = pointer.To(d.Get("display_name").(string))
		}

		if _, err := client.UpdateConditionalAccessNamedLocation(ctx, *id, *properties, conditionalaccessnamedlocation.DefaultUpdateConditionalAccessNamedLocationOperationOptions()); err != nil {
			return tf.ErrorDiagF(err, "Updating %s", id)
		}

		if err := consistency.WaitForUpdate(ctx, func(ctx context.Context) (*bool, error) {
			resp, err := client.GetConditionalAccessNamedLocation(ctx, *id, conditionalaccessnamedlocation.DefaultGetConditionalAccessNamedLocationOperationOptions())
			if err != nil {
				return nil, err
			}

			if resp.Model == nil {
				return nil, errors.New("returned model was nil")
			}

			namedLocation, ok := resp.Model.(stable.IPNamedLocation)
			if !ok {
				return nil, errors.New("returned model was not an IPNamedLocation")
			}

			if locationRaw := flattenIPNamedLocation(&namedLocation); len(locationRaw) > 0 {
				location := locationRaw[0].(map[string]interface{})
				ip := v.([]interface{})[0].(map[string]interface{})

				if !reflect.DeepEqual(location["ip_ranges"], ip["ip_ranges"]) {
					return pointer.To(false), nil
				}

				if location["trusted"].(bool) != ip["trusted"].(bool) {
					return pointer.To(false), nil
				}
			}

			return pointer.To(true), nil
		}); err != nil {
			return tf.ErrorDiagF(err, "waiting for update of %s", id)
		}

	} else if v, ok := d.GetOk("country"); ok {
		properties := expandCountryNamedLocation(v.([]interface{}))

		if d.HasChange("display_name") {
			properties.DisplayName = pointer.To(d.Get("display_name").(string))
		}

		if _, err := client.UpdateConditionalAccessNamedLocation(ctx, *id, *properties, conditionalaccessnamedlocation.DefaultUpdateConditionalAccessNamedLocationOperationOptions()); err != nil {
			return tf.ErrorDiagF(err, "Updating %s", id)
		}

		if err := consistency.WaitForUpdate(ctx, func(ctx context.Context) (*bool, error) {
			resp, err := client.GetConditionalAccessNamedLocation(ctx, *id, conditionalaccessnamedlocation.DefaultGetConditionalAccessNamedLocationOperationOptions())
			if err != nil {
				return nil, err
			}

			if resp.Model == nil {
				return nil, errors.New("returned model was nil")
			}

			namedLocation, ok := resp.Model.(stable.CountryNamedLocation)
			if !ok {
				return nil, errors.New("returned model was not a CountryNamedLocation")
			}

			if locationRaw := flattenCountryNamedLocation(&namedLocation); len(locationRaw) > 0 {
				location := locationRaw[0].(map[string]interface{})
				ip := v.([]interface{})[0].(map[string]interface{})

				if !reflect.DeepEqual(location["countries_and_regions"], ip["countries_and_regions"]) {
					return pointer.To(false), nil
				}

				if location["include_unknown_countries_and_regions"].(bool) != ip["include_unknown_countries_and_regions"].(bool) {
					return pointer.To(false), nil
				}
			}

			return pointer.To(true), nil
		}); err != nil {
			return tf.ErrorDiagF(err, "waiting for update of %s", id)
		}
	}

	return namedLocationResourceRead(ctx, d, meta)
}

func namedLocationResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).ConditionalAccess.NamedLocationClient

	id, err := stable.ParseIdentityConditionalAccessNamedLocationID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing Named Location ID")
	}

	resp, err := client.GetConditionalAccessNamedLocation(ctx, *id, conditionalaccessnamedlocation.DefaultGetConditionalAccessNamedLocationOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			log.Printf("[DEBUG] %s was not found - removing from state", id)
			d.SetId("")
			return nil
		}
	}

	if resp.Model == nil {
		return tf.ErrorDiagF(errors.New("returned model was nil"), "Bad API Response")
	}

	switch namedLocation := resp.Model.(type) {
	case stable.IPNamedLocation:
		if namedLocation.Id == nil {
			return tf.ErrorDiagF(errors.New("ID is nil for returned IP Named Location"), "Bad API response")
		}

		tf.Set(d, "display_name", pointer.From(namedLocation.DisplayName))
		tf.Set(d, "ip", flattenIPNamedLocation(&namedLocation))

	case stable.CountryNamedLocation:
		if namedLocation.Id == nil {
			return tf.ErrorDiagF(errors.New("ID is nil for returned Country Named Location"), "Bad API response")
		}

		tf.Set(d, "display_name", pointer.From(namedLocation.DisplayName))
		tf.Set(d, "country", flattenCountryNamedLocation(&namedLocation))
	}

	return nil
}

func namedLocationResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).ConditionalAccess.NamedLocationClient

	id, err := stable.ParseIdentityConditionalAccessNamedLocationID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing Named Location ID")
	}

	if v, ok := d.GetOk("ip"); ok {
		properties := expandIPNamedLocation(v.([]interface{}))
		properties.IsTrusted = pointer.To(false)

		if resp, err := client.UpdateConditionalAccessNamedLocation(ctx, *id, properties, conditionalaccessnamedlocation.DefaultUpdateConditionalAccessNamedLocationOperationOptions()); err != nil {
			if response.WasNotFound(resp.HttpResponse) {
				log.Printf("[DEBUG] %s already deleted", id)
				return nil
			}

			return tf.ErrorDiagF(err, "updating %s prior to deletion", id)
		}
	}

	resp, err := client.DeleteConditionalAccessNamedLocation(ctx, *id, conditionalaccessnamedlocation.DefaultDeleteConditionalAccessNamedLocationOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			log.Printf("[DEBUG] %s already deleted", id)
			return nil
		}

		return tf.ErrorDiagF(err, "deleting %s", id)
	}

	if err = consistency.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
		resp, err := client.GetConditionalAccessNamedLocation(ctx, *id, conditionalaccessnamedlocation.DefaultGetConditionalAccessNamedLocationOperationOptions())
		if err != nil {
			if response.WasNotFound(resp.HttpResponse) {
				return pointer.To(false), nil
			}

			return nil, err
		}

		return pointer.To(true), nil
	}); err != nil {
		return tf.ErrorDiagF(err, "waiting for deletion of %s", id)
	}

	return nil
}
