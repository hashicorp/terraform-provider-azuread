// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package conditionalaccess

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
	"github.com/manicminer/hamilton/msgraph"
)

func authenticationStrengthPolicyResource() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		CreateContext: authenticationStrengthPolicyCreate,
		ReadContext:   authenticationStrengthPolicyRead,
		UpdateContext: authenticationStrengthPolicyUpdate,
		DeleteContext: authenticationStrengthPolicyDelete,

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
			"display_name": {
				Description:  "The display name for the authentication strength policy",
				Type:         pluginsdk.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"description": {
				Description: "The description for the authentication strength policy",
				Type:        pluginsdk.TypeString,
				Optional:    true,
			},

			"allowed_combinations": {
				Description: "The allowed MFA methods for this policy",
				Type:        pluginsdk.TypeSet,
				Required:    true,
				Elem: &pluginsdk.Schema{
					Type: pluginsdk.TypeString,
				},
			},
		},
	}
}

func authenticationStrengthPolicyCreate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ConditionalAccess.AuthenticationStrengthPoliciesClient

	properties := msgraph.AuthenticationStrengthPolicy{
		DisplayName:         pointer.To(d.Get("display_name").(string)),
		Description:         pointer.To(d.Get("description").(string)),
		AllowedCombinations: tf.ExpandStringSlicePtr(d.Get("allowed_combinations").(*pluginsdk.Set).List()),
	}

	authenticationStrengthPolicy, _, err := client.Create(ctx, properties)
	if err != nil {
		return tf.ErrorDiagF(err, "Could not create authentication strength policy")
	}

	d.SetId(*authenticationStrengthPolicy.ID)

	return authenticationStrengthPolicyRead(ctx, d, meta)
}

func authenticationStrengthPolicyUpdate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ConditionalAccess.AuthenticationStrengthPoliciesClient

	properties := msgraph.AuthenticationStrengthPolicy{
		ID:          pointer.To(d.Id()),
		DisplayName: pointer.To(d.Get("display_name").(string)),
		Description: pointer.To(d.Get("description").(string)),
		// AllowedCombinations: tf.ExpandStringSlicePtr(d.Get("allowed_combinations").(*pluginsdk.Set).List()),
	}

	_, err := client.Update(ctx, properties)
	if err != nil {
		return tf.ErrorDiagF(err, "Could not update authentication strength policy")
	}

	if d.HasChange("allowed_combinations") {
		properties.AllowedCombinations = tf.ExpandStringSlicePtr(d.Get("allowed_combinations").(*pluginsdk.Set).List())
		_, err := client.UpdateAllowedCombinations(ctx, properties)
		if err != nil {
			return tf.ErrorDiagF(err, "Could not update authentication strength policy allowed combinations")
		}
	}

	return authenticationStrengthPolicyRead(ctx, d, meta)
}

func authenticationStrengthPolicyRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ConditionalAccess.AuthenticationStrengthPoliciesClient

	authenticationStrengthPolicy, status, err := client.Get(ctx, d.Id(), odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Authentication Strength Policy with Object ID %q was not found - removing from state", d.Id())
			d.SetId("")
			return nil
		}
	}
	if authenticationStrengthPolicy == nil {
		return tf.ErrorDiagF(errors.New("Bad API response"), "Result is nil")
	}

	d.SetId(*authenticationStrengthPolicy.ID)
	tf.Set(d, "display_name", authenticationStrengthPolicy.DisplayName)
	tf.Set(d, "description", authenticationStrengthPolicy.Description)
	tf.Set(d, "allowed_combinations", tf.FlattenStringSlicePtr(authenticationStrengthPolicy.AllowedCombinations))

	return nil
}

func authenticationStrengthPolicyDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ConditionalAccess.AuthenticationStrengthPoliciesClient
	authenticationStrengthPolicyId := d.Id()

	if _, status, err := client.Get(ctx, authenticationStrengthPolicyId, odata.Query{}); err != nil {
		if status == http.StatusNotFound {
			log.Printf("[DEBUG] Authentication Strength Policy with ID %q already deleted", authenticationStrengthPolicyId)
			return nil
		}

		return tf.ErrorDiagPathF(err, "id", "Retrieving Authentication Strength Policy with ID %q", authenticationStrengthPolicyId)
	}

	status, err := client.Delete(ctx, authenticationStrengthPolicyId)
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Deleting Authentication Strength Policy with ID %q, got status %d", authenticationStrengthPolicyId, status)
	}

	if err := helpers.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
		defer func() { client.BaseClient.DisableRetries = false }()
		client.BaseClient.DisableRetries = true
		if _, status, err := client.Get(ctx, authenticationStrengthPolicyId, odata.Query{}); err != nil {
			if status == http.StatusNotFound {
				return pointer.To(false), nil
			}
			return nil, err
		}
		return pointer.To(true), nil
	}); err != nil {
		return tf.ErrorDiagF(err, "waiting for deletion of Authentication Strength Policy with ID %q", authenticationStrengthPolicyId)
	}

	return nil
}
