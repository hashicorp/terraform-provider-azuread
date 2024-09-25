// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package policies

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/claimsmappingpolicy"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/policies/migrations"
)

func claimsMappingPolicyResource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		CreateContext: claimsMappingPolicyResourceCreate,
		ReadContext:   claimsMappingPolicyResourceRead,
		UpdateContext: claimsMappingPolicyResourceUpdate,
		DeleteContext: claimsMappingPolicyResourceDelete,

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(5 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Update: pluginsdk.DefaultTimeout(5 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(5 * time.Minute),
		},

		Importer: pluginsdk.ImporterValidatingResourceId(func(id string) error {
			if _, errs := stable.ValidatePolicyClaimsMappingPolicyID(id, "id"); len(errs) > 0 {
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
				Type:    migrations.ResourceClaimsMappingPolicyInstanceResourceV0().CoreConfigSchema().ImpliedType(),
				Upgrade: migrations.ResourceClaimsMappingPolicyInstanceStateUpgradeV0,
				Version: 0,
			},
		},

		Schema: map[string]*pluginsdk.Schema{
			"definition": {
				Description: "A string collection containing a JSON string that defines the rules and settings for this policy",
				Type:        pluginsdk.TypeList,
				Required:    true,
				Elem: &pluginsdk.Schema{
					Type:         pluginsdk.TypeString,
					ValidateFunc: validation.StringIsNotEmpty,
				},
			},

			"display_name": {
				Description:  "Display name for this policy",
				Type:         pluginsdk.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
		},
	}
}

func claimsMappingPolicyResourceCreate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Policies.ClaimsMappingPolicyClient

	properties := stable.ClaimsMappingPolicy{
		Definition:  tf.ExpandStringSlice(d.Get("definition").([]interface{})),
		DisplayName: nullable.Value(d.Get("display_name").(string)),
	}

	resp, err := client.CreateClaimsMappingPolicy(ctx, properties, claimsmappingpolicy.DefaultCreateClaimsMappingPolicyOperationOptions())
	if err != nil {
		return tf.ErrorDiagF(err, "Could not create Claims Mapping Policy")
	}

	claimsMappingPolicy := resp.Model
	if claimsMappingPolicy == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "Could not create Claims Mapping Policy")
	}
	if claimsMappingPolicy.Id == nil {
		return tf.ErrorDiagF(errors.New("model return with nil ID"), "Could not create Claims Mapping Policy")
	}

	id := stable.NewPolicyClaimsMappingPolicyID(*claimsMappingPolicy.Id)
	d.SetId(id.ID())

	return claimsMappingPolicyResourceRead(ctx, d, meta)
}

func claimsMappingPolicyResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Policies.ClaimsMappingPolicyClient

	id, err := stable.ParsePolicyClaimsMappingPolicyID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing ID")
	}

	resp, err := client.GetClaimsMappingPolicy(ctx, *id, claimsmappingpolicy.DefaultGetClaimsMappingPolicyOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			log.Printf("[DEBUG] %s - removing from state!", id)
			d.SetId("")
			return nil
		}

		return tf.ErrorDiagF(err, "retrieving %s", id)
	}

	claimsMappingPolicy := resp.Model
	if claimsMappingPolicy == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "Retrieving %s", id)
	}

	tf.Set(d, "definition", tf.FlattenStringSlice(claimsMappingPolicy.Definition))
	tf.Set(d, "display_name", claimsMappingPolicy.DisplayName.GetOrZero())

	return nil
}

func claimsMappingPolicyResourceUpdate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Policies.ClaimsMappingPolicyClient

	id, err := stable.ParsePolicyClaimsMappingPolicyID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing ID")
	}

	properties := stable.ClaimsMappingPolicy{
		Definition:  tf.ExpandStringSlice(d.Get("definition").([]interface{})),
		DisplayName: nullable.Value(d.Get("display_name").(string)),
	}

	if _, err := client.UpdateClaimsMappingPolicy(ctx, *id, properties, claimsmappingpolicy.DefaultUpdateClaimsMappingPolicyOperationOptions()); err != nil {
		return tf.ErrorDiagF(err, "Could not update %s", id)
	}

	return claimsMappingPolicyResourceRead(ctx, d, meta)
}

func claimsMappingPolicyResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Policies.ClaimsMappingPolicyClient

	id, err := stable.ParsePolicyClaimsMappingPolicyID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing ID")
	}

	if _, err := client.DeleteClaimsMappingPolicy(ctx, *id, claimsmappingpolicy.DefaultDeleteClaimsMappingPolicyOperationOptions()); err != nil {
		return tf.ErrorDiagF(err, "Deleting %s", id)
	}

	return nil
}
