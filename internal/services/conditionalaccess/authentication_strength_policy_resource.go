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

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/hashicorp/terraform-provider-azuread/internal/validate"
	"github.com/manicminer/hamilton/msgraph"
)

func authenticationStrengthPolicyResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: authenticationStrengthPolicyCreate,
		ReadContext:   authenticationStrengthPolicyRead,
		UpdateContext: authenticationStrengthPolicyUpdate,
		DeleteContext: authenticationStrengthPolicyDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Importer: tf.ValidateResourceIDPriorToImport(func(id string) error {
			if _, err := uuid.ParseUUID(id); err != nil {
				return fmt.Errorf("specified ID (%q) is not valid: %s", id, err)
			}
			return nil
		}),

		Schema: map[string]*schema.Schema{
			"display_name": {
				Description:      "The display name for the authentication strength policy",
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validate.NoEmptyStrings,
			},

			"description": {
				Description: "The description for the authentication strength policy",
				Type:        schema.TypeString,
				Optional:    true,
			},

			"allowed_combinations": {
				Description: "The allowed MFA methods for this policy",
				Type:        schema.TypeSet,
				Required:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func authenticationStrengthPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ConditionalAccess.AuthenticationStrengthPoliciesClient

	properties := msgraph.AuthenticationStrengthPolicy{
		DisplayName:         utils.String(d.Get("display_name").(string)),
		Description:         utils.String(d.Get("description").(string)),
		AllowedCombinations: tf.ExpandStringSlicePtr(d.Get("allowed_combinations").(*schema.Set).List()),
	}

	authenticationStrengthPolicy, _, err := client.Create(ctx, properties)
	if err != nil {
		return tf.ErrorDiagF(err, "Could not create authentication strength policy")
	}

	d.SetId(*authenticationStrengthPolicy.ID)

	return authenticationStrengthPolicyRead(ctx, d, meta)
}

func authenticationStrengthPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).ConditionalAccess.AuthenticationStrengthPoliciesClient

	properties := msgraph.AuthenticationStrengthPolicy{
		ID:          utils.String(d.Id()),
		DisplayName: utils.String(d.Get("display_name").(string)),
		Description: utils.String(d.Get("description").(string)),
		// AllowedCombinations: tf.ExpandStringSlicePtr(d.Get("allowed_combinations").(*schema.Set).List()),
	}

	_, err := client.Update(ctx, properties)
	if err != nil {
		return tf.ErrorDiagF(err, "Could not update authentication strength policy")
	}

	if d.HasChange("allowed_combinations") {
		properties.AllowedCombinations = tf.ExpandStringSlicePtr(d.Get("allowed_combinations").(*schema.Set).List())
		_, err := client.UpdateAllowedCombinations(ctx, properties)
		if err != nil {
			return tf.ErrorDiagF(err, "Could not update authentication strength policy allowed combinations")
		}
	}

	return authenticationStrengthPolicyRead(ctx, d, meta)
}

func authenticationStrengthPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func authenticationStrengthPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
				return utils.Bool(false), nil
			}
			return nil, err
		}
		return utils.Bool(true), nil
	}); err != nil {
		return tf.ErrorDiagF(err, "waiting for deletion of Authentication Strength Policy with ID %q", authenticationStrengthPolicyId)
	}

	return nil
}
