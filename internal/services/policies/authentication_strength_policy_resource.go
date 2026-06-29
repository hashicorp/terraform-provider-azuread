// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package policies

import (
	"context"
	"errors"
	"fmt"
	"log"
	"slices"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/authenticationstrengthpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/authenticationstrengthpolicycombinationconfiguration"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/consistency"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/policies/migrations"
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
			if _, errs := stable.ValidatePolicyAuthenticationStrengthPolicyID(id, "id"); len(errs) > 0 {
				out := ""
				for _, err := range errs {
					out += err.Error()
				}
				return errors.New(out)
			}
			return nil
		}),

		SchemaVersion: 1,
		StateUpgraders: []pluginsdk.StateUpgrader{
			{
				Type:    migrations.ResourceAuthenticationStrengthPolicyInstanceResourceV0().CoreConfigSchema().ImpliedType(),
				Upgrade: migrations.ResourceAuthenticationStrengthPolicyInstanceStateUpgradeV0,
				Version: 0,
			},
		},

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
					ValidateFunc: func(in interface{}, k string) ([]string, []error) {
						val, ok := in.(string)
						if !ok {
							return nil, []error{fmt.Errorf("expected a string value for %q", k)}
						}
						split := strings.Split(val, ",")
						for _, s := range split {
							if !slices.Contains(stable.PossibleValuesForAuthenticationMethodModes(), strings.TrimSpace(s)) {
								return nil, []error{fmt.Errorf("unrecognized authentication method %q in %q", s, k)}
							}
						}
						return nil, nil
					},
				},
			},

			"fido2_combination_configuration": {
				Description: "Restrictions applied to the `fido2` authentication method combination",
				Type:        pluginsdk.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"allowed_aaguids": {
							Description: "A list of AAGUIDs allowed to be used as part of the `fido2` combination",
							Type:        pluginsdk.TypeSet,
							Required:    true,
							Elem: &pluginsdk.Schema{
								Type:             pluginsdk.TypeString,
								ValidateDiagFunc: validation.ValidateDiag(validation.IsUUID),
							},
						},

						"id": {
							Description: "The system-generated ID of the combination configuration",
							Type:        pluginsdk.TypeString,
							Computed:    true,
						},
					},
				},
			},

			"x509_certificate_combination_configuration": {
				Description: "Restrictions applied to the `x509Certificate` authentication method combinations",
				Type:        pluginsdk.TypeList,
				Optional:    true,
				MaxItems:    1,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"applies_to_combinations": {
							Description: "The x509 certificate authentication method combinations this configuration applies to",
							Type:        pluginsdk.TypeSet,
							Required:    true,
							Elem: &pluginsdk.Schema{
								Type: pluginsdk.TypeString,
								ValidateFunc: validation.StringInSlice([]string{
									string(stable.AuthenticationMethodModes_X509CertificateSingleFactor),
									string(stable.AuthenticationMethodModes_X509CertificateMultiFactor),
								}, false),
							},
						},

						"allowed_issuer_skis": {
							Description:  "A list of allowed subject key identifier values",
							Type:         pluginsdk.TypeSet,
							Optional:     true,
							AtLeastOneOf: []string{"x509_certificate_combination_configuration.0.allowed_issuer_skis", "x509_certificate_combination_configuration.0.allowed_policy_oids"},
							Elem: &pluginsdk.Schema{
								Type:         pluginsdk.TypeString,
								ValidateFunc: validation.StringIsNotEmpty,
							},
						},

						"allowed_policy_oids": {
							Description:  "A list of allowed certificate policy OIDs",
							Type:         pluginsdk.TypeSet,
							Optional:     true,
							AtLeastOneOf: []string{"x509_certificate_combination_configuration.0.allowed_issuer_skis", "x509_certificate_combination_configuration.0.allowed_policy_oids"},
							Elem: &pluginsdk.Schema{
								Type:         pluginsdk.TypeString,
								ValidateFunc: validation.StringIsNotEmpty,
							},
						},

						"id": {
							Description: "The system-generated ID of the combination configuration",
							Type:        pluginsdk.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func authenticationStrengthPolicyCreate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Policies.AuthenticationStrengthPolicyClient

	allowedCombinations := make([]stable.AuthenticationMethodModes, 0)
	for _, v := range d.Get("allowed_combinations").(*pluginsdk.Set).List() {
		allowedCombinations = append(allowedCombinations, stable.AuthenticationMethodModes(v.(string)))
	}

	properties := stable.AuthenticationStrengthPolicy{
		DisplayName:         pointer.To(d.Get("display_name").(string)),
		Description:         nullable.NoZero(d.Get("description").(string)),
		AllowedCombinations: pointer.To(allowedCombinations),
	}

	resp, err := client.CreateAuthenticationStrengthPolicy(ctx, properties, authenticationstrengthpolicy.DefaultCreateAuthenticationStrengthPolicyOperationOptions())
	if err != nil {
		return tf.ErrorDiagF(err, "Could not create authentication strength policy")
	}

	authenticationStrengthPolicy := resp.Model
	if authenticationStrengthPolicy == nil {
		return tf.ErrorDiagF(errors.New("model was nil"), "Could not create authentication strength policy")
	}
	if authenticationStrengthPolicy.Id == nil {
		return tf.ErrorDiagF(errors.New("model returned with nil ID"), "Could not create authentication strength policy")
	}

	id := stable.NewPolicyAuthenticationStrengthPolicyID(*authenticationStrengthPolicy.Id)

	// Wait for the policy to appear consistently
	if err = consistency.WaitForUpdate(ctx, func(ctx context.Context) (*bool, error) {
		resp, err := client.GetAuthenticationStrengthPolicy(ctx, id, authenticationstrengthpolicy.DefaultGetAuthenticationStrengthPolicyOperationOptions())
		if err != nil {
			if response.WasNotFound(resp.HttpResponse) {
				return pointer.To(false), nil
			}
			return pointer.To(false), fmt.Errorf("retrieving authentication strength policy")
		}
		return pointer.To(true), nil
	}); err != nil {
		return tf.ErrorDiagF(err, "Waiting for creation of %s", id)
	}

	d.SetId(id.ID())

	if diags := authenticationStrengthPolicyCombinationConfigurationsSync(ctx, meta, id, d); diags.HasError() {
		return diags
	}

	return authenticationStrengthPolicyRead(ctx, d, meta)
}

func authenticationStrengthPolicyUpdate(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Policies.AuthenticationStrengthPolicyClient

	id, err := stable.ParsePolicyAuthenticationStrengthPolicyID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing ID")
	}

	properties := stable.AuthenticationStrengthPolicy{
		DisplayName: pointer.To(d.Get("display_name").(string)),
		Description: nullable.NoZero(d.Get("description").(string)),
	}

	if _, err := client.UpdateAuthenticationStrengthPolicy(ctx, *id, properties, authenticationstrengthpolicy.DefaultUpdateAuthenticationStrengthPolicyOperationOptions()); err != nil {
		return tf.ErrorDiagF(err, "Could not update %s", id)
	}

	if d.HasChange("allowed_combinations") {
		allowedCombinations := make([]stable.AuthenticationMethodModes, 0)
		for _, v := range d.Get("allowed_combinations").(*pluginsdk.Set).List() {
			allowedCombinations = append(allowedCombinations, stable.AuthenticationMethodModes(v.(string)))
		}

		request := authenticationstrengthpolicy.UpdateAuthenticationStrengthPolicyAllowedCombinationsRequest{
			AllowedCombinations: pointer.To(allowedCombinations),
		}

		if _, err := client.UpdateAuthenticationStrengthPolicyAllowedCombinations(ctx, *id, request, authenticationstrengthpolicy.DefaultUpdateAuthenticationStrengthPolicyAllowedCombinationsOperationOptions()); err != nil {
			return tf.ErrorDiagF(err, "Could not update allowed combinations for %s", id)
		}
	}

	if diags := authenticationStrengthPolicyCombinationConfigurationsSync(ctx, meta, *id, d); diags.HasError() {
		return diags
	}

	return authenticationStrengthPolicyRead(ctx, d, meta)
}

func authenticationStrengthPolicyRead(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Policies.AuthenticationStrengthPolicyClient

	id, err := stable.ParsePolicyAuthenticationStrengthPolicyID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing ID")
	}

	resp, err := client.GetAuthenticationStrengthPolicy(ctx, *id, authenticationstrengthpolicy.DefaultGetAuthenticationStrengthPolicyOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			log.Printf("[DEBUG] Authentication Strength Policy with Object ID %q was not found - removing from state", d.Id())
			d.SetId("")
			return nil
		}
	}
	authenticationStrengthPolicy := resp.Model
	if authenticationStrengthPolicy == nil {
		return tf.ErrorDiagF(errors.New("Bad API response"), "Result is nil")
	}

	tf.Set(d, "display_name", pointer.From(authenticationStrengthPolicy.DisplayName))
	tf.Set(d, "description", authenticationStrengthPolicy.Description.GetOrZero())

	allowedCombinations := make([]string, 0)
	for _, v := range pointer.From(authenticationStrengthPolicy.AllowedCombinations) {
		allowedCombinations = append(allowedCombinations, string(v))
	}
	tf.Set(d, "allowed_combinations", tf.FlattenStringSlice(allowedCombinations))

	fido2, x509 := flattenCombinationConfigurations(authenticationStrengthPolicy.CombinationConfigurations)
	tf.Set(d, "fido2_combination_configuration", fido2)
	tf.Set(d, "x509_certificate_combination_configuration", x509)

	return nil
}

func authenticationStrengthPolicyDelete(ctx context.Context, d *pluginsdk.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*clients.Client).Policies.AuthenticationStrengthPolicyClient

	id, err := stable.ParsePolicyAuthenticationStrengthPolicyID(d.Id())
	if err != nil {
		return tf.ErrorDiagPathF(err, "id", "Parsing ID")
	}

	if _, err := client.DeleteAuthenticationStrengthPolicy(ctx, *id, authenticationstrengthpolicy.DefaultDeleteAuthenticationStrengthPolicyOperationOptions()); err != nil {
		return tf.ErrorDiagPathF(err, "id", "Deleting %s", id)
	}

	if err := consistency.WaitForDeletion(ctx, func(ctx context.Context) (*bool, error) {
		if resp, err := client.GetAuthenticationStrengthPolicy(ctx, *id, authenticationstrengthpolicy.DefaultGetAuthenticationStrengthPolicyOperationOptions()); err != nil {
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

func authenticationStrengthPolicyCombinationConfigurationsSync(ctx context.Context, meta interface{}, id stable.PolicyAuthenticationStrengthPolicyId, d *pluginsdk.ResourceData) diag.Diagnostics {
	client := meta.(*clients.Client).Policies.AuthenticationStrengthPolicyCombinationConfigurationClient

	if d.HasChange("fido2_combination_configuration") {
		if err := syncCombinationConfiguration(ctx, client, id, d, "fido2_combination_configuration", expandFido2CombinationConfiguration); err != nil {
			return tf.ErrorDiagF(err, "Could not sync fido2 combination configuration for %s", id)
		}
	}

	if d.HasChange("x509_certificate_combination_configuration") {
		if err := syncCombinationConfiguration(ctx, client, id, d, "x509_certificate_combination_configuration", expandX509CertificateCombinationConfiguration); err != nil {
			return tf.ErrorDiagF(err, "Could not sync x509 certificate combination configuration for %s", id)
		}
	}

	return nil
}

func syncCombinationConfiguration(ctx context.Context, client *authenticationstrengthpolicycombinationconfiguration.AuthenticationStrengthPolicyCombinationConfigurationClient, id stable.PolicyAuthenticationStrengthPolicyId, d *pluginsdk.ResourceData, key string, expand func([]interface{}) stable.AuthenticationCombinationConfiguration) error {
	oldRaw, newRaw := d.GetChange(key)
	oldList := oldRaw.([]interface{})
	newList := newRaw.([]interface{})

	var existingId string
	if len(oldList) > 0 && oldList[0] != nil {
		existingId = oldList[0].(map[string]interface{})["id"].(string)
	}

	switch {
	case len(newList) > 0 && existingId == "":
		if _, err := client.CreateAuthenticationStrengthPolicyCombinationConfiguration(ctx, id, expand(newList), authenticationstrengthpolicycombinationconfiguration.DefaultCreateAuthenticationStrengthPolicyCombinationConfigurationOperationOptions()); err != nil {
			return err
		}

	case len(newList) > 0 && existingId != "":
		configId := stable.NewPolicyAuthenticationStrengthPolicyIdCombinationConfigurationID(id.AuthenticationStrengthPolicyId, existingId)
		if _, err := client.UpdateAuthenticationStrengthPolicyCombinationConfiguration(ctx, configId, expand(newList), authenticationstrengthpolicycombinationconfiguration.DefaultUpdateAuthenticationStrengthPolicyCombinationConfigurationOperationOptions()); err != nil {
			return err
		}

	case len(newList) == 0 && existingId != "":
		configId := stable.NewPolicyAuthenticationStrengthPolicyIdCombinationConfigurationID(id.AuthenticationStrengthPolicyId, existingId)
		resp, err := client.DeleteAuthenticationStrengthPolicyCombinationConfiguration(ctx, configId, authenticationstrengthpolicycombinationconfiguration.DefaultDeleteAuthenticationStrengthPolicyCombinationConfigurationOperationOptions())
		if err != nil && !response.WasNotFound(resp.HttpResponse) {
			return err
		}
	}

	return nil
}

func expandFido2CombinationConfiguration(input []interface{}) stable.AuthenticationCombinationConfiguration {
	raw := input[0].(map[string]interface{})
	return stable.Fido2CombinationConfiguration{
		AppliesToCombinations: &[]stable.AuthenticationMethodModes{stable.AuthenticationMethodModes_Fido2},
		AllowedAAGUIDs:        tf.ExpandStringSlicePtr(raw["allowed_aaguids"].(*pluginsdk.Set).List()),
	}
}

func expandX509CertificateCombinationConfiguration(input []interface{}) stable.AuthenticationCombinationConfiguration {
	raw := input[0].(map[string]interface{})

	appliesToCombinations := make([]stable.AuthenticationMethodModes, 0)
	for _, v := range raw["applies_to_combinations"].(*pluginsdk.Set).List() {
		appliesToCombinations = append(appliesToCombinations, stable.AuthenticationMethodModes(v.(string)))
	}

	config := stable.X509CertificateCombinationConfiguration{
		AppliesToCombinations: pointer.To(appliesToCombinations),
	}

	if skis := raw["allowed_issuer_skis"].(*pluginsdk.Set).List(); len(skis) > 0 {
		config.AllowedIssuerSkis = tf.ExpandStringSlicePtr(skis)
	}
	if oids := raw["allowed_policy_oids"].(*pluginsdk.Set).List(); len(oids) > 0 {
		config.AllowedPolicyOIDs = tf.ExpandStringSlicePtr(oids)
	}

	return config
}

func flattenCombinationConfigurations(input *[]stable.AuthenticationCombinationConfiguration) (fido2 []interface{}, x509 []interface{}) {
	fido2 = make([]interface{}, 0)
	x509 = make([]interface{}, 0)
	if input == nil {
		return
	}

	for _, config := range *input {
		switch c := config.(type) {
		case stable.Fido2CombinationConfiguration:
			fido2 = append(fido2, map[string]interface{}{
				"allowed_aaguids": tf.FlattenStringSlicePtr(c.AllowedAAGUIDs),
				"id":              pointer.From(c.Id),
			})
		case stable.X509CertificateCombinationConfiguration:
			appliesToCombinations := make([]interface{}, 0)
			for _, v := range pointer.From(c.AppliesToCombinations) {
				appliesToCombinations = append(appliesToCombinations, string(v))
			}
			x509 = append(x509, map[string]interface{}{
				"applies_to_combinations": appliesToCombinations,
				"allowed_issuer_skis":     tf.FlattenStringSlicePtr(c.AllowedIssuerSkis),
				"allowed_policy_oids":     tf.FlattenStringSlicePtr(c.AllowedPolicyOIDs),
				"id":                      pointer.From(c.Id),
			})
		}
	}

	return
}
