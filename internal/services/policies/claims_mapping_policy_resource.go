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
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
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
			if _, err := uuid.ParseUUID(id); err != nil {
				return fmt.Errorf("specified ID (%q) is not valid: %s", id, err)
			}
			return nil
		}),

		Schema: map[string]*pluginsdk.Schema{
			"definition": {
				Description: "A string collection containing a JSON string that defines the rules and settings for this policy",
				Type:        pluginsdk.TypeList,
				Required:    true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},

			"display_name": {
				Description: "Display name for this policy",
				Type:        pluginsdk.TypeString,
				Required:    true,
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

	resp, err := client.CreateClaimsMappingPolicy(ctx, properties)
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
	d.SetId(id.ClaimsMappingPolicyId)

	return claimsMappingPolicyResourceRead(ctx, d, meta)
}

func claimsMappingPolicyResourceRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Policies.ClaimsMappingPolicyClient
	id := stable.NewPolicyClaimsMappingPolicyID(d.Id())

	resp, err := client.GetClaimsMappingPolicy(ctx, id, claimsmappingpolicy.DefaultGetClaimsMappingPolicyOperationOptions())
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
	id := stable.NewPolicyClaimsMappingPolicyID(d.Id())

	properties := stable.ClaimsMappingPolicy{
		Definition:  tf.ExpandStringSlice(d.Get("definition").([]interface{})),
		DisplayName: nullable.Value(d.Get("display_name").(string)),
	}

	if _, err := client.UpdateClaimsMappingPolicy(ctx, id, properties); err != nil {
		return tf.ErrorDiagF(err, "Could not update %s", id)
	}

	return claimsMappingPolicyResourceRead(ctx, d, meta)
}

func claimsMappingPolicyResourceDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) pluginsdk.Diagnostics {
	client := meta.(*clients.Client).Policies.ClaimsMappingPolicyClient
	id := stable.NewPolicyClaimsMappingPolicyID(d.Id())

	if _, err := client.DeleteClaimsMappingPolicy(ctx, id, claimsmappingpolicy.DefaultDeleteClaimsMappingPolicyOperationOptions()); err != nil {
		return tf.ErrorDiagF(err, "Deleting %s", id)
	}

	return nil
}
