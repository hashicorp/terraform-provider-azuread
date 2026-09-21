// Copyright IBM Corp. 2019, 2026
// SPDX-License-Identifier: MPL-2.0

package policies

import (
	"context"
	"fmt"
	"slices"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/authenticationstrengthpolicycombinationconfiguration"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/sdk"
)

type AuthenticationStrengthPolicyX509CombinationConfigurationResourceModel struct {
	AuthenticationStrengthPolicyId string   `tfschema:"authentication_strength_policy_id"`
	AppliesToCombinations          []string `tfschema:"applies_to_combinations"`
	AllowedIssuerSkis              []string `tfschema:"allowed_issuer_skis"`
	AllowedPolicyOIDs              []string `tfschema:"allowed_policy_oids"`
}

var (
	_ sdk.Resource           = AuthenticationStrengthPolicyX509CombinationConfigurationResource{}
	_ sdk.ResourceWithUpdate = AuthenticationStrengthPolicyX509CombinationConfigurationResource{}
)

type AuthenticationStrengthPolicyX509CombinationConfigurationResource struct{}

func (r AuthenticationStrengthPolicyX509CombinationConfigurationResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return stable.ValidatePolicyAuthenticationStrengthPolicyIdCombinationConfigurationID
}

func (r AuthenticationStrengthPolicyX509CombinationConfigurationResource) ResourceType() string {
	return "azuread_authentication_strength_policy_x509_combination_configuration"
}

func (r AuthenticationStrengthPolicyX509CombinationConfigurationResource) ModelObject() interface{} {
	return &AuthenticationStrengthPolicyX509CombinationConfigurationResourceModel{}
}

func (r AuthenticationStrengthPolicyX509CombinationConfigurationResource) Arguments() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"authentication_strength_policy_id": {
			Description:  "The object ID of the authentication strength policy to which this combination configuration applies",
			Type:         pluginsdk.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.IsUUID,
		},

		"applies_to_combinations": {
			Description: "The x509 certificate authentication method combinations this configuration applies to",
			Type:        pluginsdk.TypeSet,
			Required:    true,
			MinItems:    1,
			Elem: &pluginsdk.Schema{
				Type:         pluginsdk.TypeString,
				ValidateFunc: validation.StringInSlice(stable.PossibleValuesForX509CertificateAuthenticationMode(), false),
			},
		},

		"allowed_issuer_skis": {
			Description:  "A set of allowed certificate issuer subject key identifier values",
			Type:         pluginsdk.TypeSet,
			Optional:     true,
			AtLeastOneOf: []string{"allowed_issuer_skis", "allowed_policy_oids"},
			Elem: &pluginsdk.Schema{
				Type:         pluginsdk.TypeString,
				ValidateFunc: validation.StringIsNotEmpty,
			},
		},

		"allowed_policy_oids": {
			Description:  "A set of allowed certificate policy OIDs",
			Type:         pluginsdk.TypeSet,
			Optional:     true,
			AtLeastOneOf: []string{"allowed_issuer_skis", "allowed_policy_oids"},
			Elem: &pluginsdk.Schema{
				Type:         pluginsdk.TypeString,
				ValidateFunc: validation.StringIsNotEmpty,
			},
		},
	}
}

func (r AuthenticationStrengthPolicyX509CombinationConfigurationResource) Attributes() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{}
}

func (r AuthenticationStrengthPolicyX509CombinationConfigurationResource) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Policies.AuthenticationStrengthPolicyCombinationConfigurationClient

			var model AuthenticationStrengthPolicyX509CombinationConfigurationResourceModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			policyId := stable.NewPolicyAuthenticationStrengthPolicyID(model.AuthenticationStrengthPolicyId)

			tf.LockByName(authenticationStrengthPolicyResourceName, policyId.AuthenticationStrengthPolicyId)
			defer tf.UnlockByName(authenticationStrengthPolicyResourceName, policyId.AuthenticationStrengthPolicyId)

			existingId, err := findCombinationConfiguration(ctx, client, policyId, x509CombinationOverlaps(model.AppliesToCombinations, ""))
			if err != nil {
				return fmt.Errorf("checking for existing x509 combination configuration for %s: %+v", policyId, err)
			}
			if existingId != "" {
				return metadata.ResourceRequiresImport(r.ResourceType(), stable.NewPolicyAuthenticationStrengthPolicyIdCombinationConfigurationID(policyId.AuthenticationStrengthPolicyId, existingId))
			}

			resp, err := client.CreateAuthenticationStrengthPolicyCombinationConfiguration(ctx, policyId, expandX509CombinationConfiguration(model), authenticationstrengthpolicycombinationconfiguration.DefaultCreateAuthenticationStrengthPolicyCombinationConfigurationOperationOptions())
			if err != nil {
				return fmt.Errorf("creating x509 combination configuration for %s: %+v", policyId, err)
			}

			if resp.Model == nil {
				return fmt.Errorf("creating x509 combination configuration for %s: model was nil", policyId)
			}

			combinationConfigurationId := resp.Model.AuthenticationCombinationConfiguration().Id
			if combinationConfigurationId == nil {
				return fmt.Errorf("creating x509 combination configuration for %s: model returned with nil ID", policyId)
			}

			id := stable.NewPolicyAuthenticationStrengthPolicyIdCombinationConfigurationID(policyId.AuthenticationStrengthPolicyId, *combinationConfigurationId)

			if err = waitForCombinationConfiguration(ctx, client, id); err != nil {
				return fmt.Errorf("waiting for creation of %s: %+v", id, err)
			}

			metadata.SetID(id)
			return nil
		},
	}
}

func (r AuthenticationStrengthPolicyX509CombinationConfigurationResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Policies.AuthenticationStrengthPolicyCombinationConfigurationClient

			id, err := stable.ParsePolicyAuthenticationStrengthPolicyIdCombinationConfigurationID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			resp, err := client.GetAuthenticationStrengthPolicyCombinationConfiguration(ctx, *id, authenticationstrengthpolicycombinationconfiguration.DefaultGetAuthenticationStrengthPolicyCombinationConfigurationOperationOptions())
			if err != nil {
				if response.WasNotFound(resp.HttpResponse) {
					return metadata.MarkAsGone(id)
				}
				return fmt.Errorf("retrieving %s: %+v", id, err)
			}

			if resp.Model == nil {
				return fmt.Errorf("retrieving %s: model was nil", id)
			}

			combinationConfiguration, ok := resp.Model.(stable.X509CertificateCombinationConfiguration)
			if !ok {
				return fmt.Errorf("retrieving %s: not an x509 certificate combination configuration", id)
			}

			appliesToCombinations := make([]string, 0)
			for _, v := range pointer.From(combinationConfiguration.AppliesToCombinations) {
				appliesToCombinations = append(appliesToCombinations, string(v))
			}

			model := AuthenticationStrengthPolicyX509CombinationConfigurationResourceModel{
				AuthenticationStrengthPolicyId: id.AuthenticationStrengthPolicyId,
				AppliesToCombinations:          appliesToCombinations,
				AllowedIssuerSkis:              pointer.From(combinationConfiguration.AllowedIssuerSkis),
				AllowedPolicyOIDs:              pointer.From(combinationConfiguration.AllowedPolicyOIDs),
			}

			return metadata.Encode(&model)
		},
	}
}

func (r AuthenticationStrengthPolicyX509CombinationConfigurationResource) Update() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Policies.AuthenticationStrengthPolicyCombinationConfigurationClient

			id, err := stable.ParsePolicyAuthenticationStrengthPolicyIdCombinationConfigurationID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			var model AuthenticationStrengthPolicyX509CombinationConfigurationResourceModel
			if err = metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			policyId := stable.NewPolicyAuthenticationStrengthPolicyID(id.AuthenticationStrengthPolicyId)

			tf.LockByName(authenticationStrengthPolicyResourceName, id.AuthenticationStrengthPolicyId)
			defer tf.UnlockByName(authenticationStrengthPolicyResourceName, id.AuthenticationStrengthPolicyId)

			conflictingId, err := findCombinationConfiguration(ctx, client, policyId, x509CombinationOverlaps(model.AppliesToCombinations, id.AuthenticationCombinationConfigurationId))
			if err != nil {
				return fmt.Errorf("checking for conflicting x509 combination configuration for %s: %+v", policyId, err)
			}
			if conflictingId != "" {
				return fmt.Errorf("updating %s: `applies_to_combinations` overlaps with %s", id, stable.NewPolicyAuthenticationStrengthPolicyIdCombinationConfigurationID(id.AuthenticationStrengthPolicyId, conflictingId))
			}

			if _, err = client.UpdateAuthenticationStrengthPolicyCombinationConfiguration(ctx, *id, expandX509CombinationConfiguration(model), authenticationstrengthpolicycombinationconfiguration.DefaultUpdateAuthenticationStrengthPolicyCombinationConfigurationOperationOptions()); err != nil {
				return fmt.Errorf("updating %s: %+v", id, err)
			}

			return nil
		},
	}
}

func (r AuthenticationStrengthPolicyX509CombinationConfigurationResource) Delete() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Policies.AuthenticationStrengthPolicyCombinationConfigurationClient

			id, err := stable.ParsePolicyAuthenticationStrengthPolicyIdCombinationConfigurationID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			tf.LockByName(authenticationStrengthPolicyResourceName, id.AuthenticationStrengthPolicyId)
			defer tf.UnlockByName(authenticationStrengthPolicyResourceName, id.AuthenticationStrengthPolicyId)

			resp, err := client.DeleteAuthenticationStrengthPolicyCombinationConfiguration(ctx, *id, authenticationstrengthpolicycombinationconfiguration.DefaultDeleteAuthenticationStrengthPolicyCombinationConfigurationOperationOptions())
			if err != nil && !response.WasNotFound(resp.HttpResponse) {
				return fmt.Errorf("deleting %s: %+v", id, err)
			}

			if err = waitForCombinationConfigurationDeletion(ctx, client, *id); err != nil {
				return fmt.Errorf("waiting for deletion of %s: %+v", id, err)
			}

			return nil
		},
	}
}

// The API permits each combination to be configured by at most one configuration.
func x509CombinationOverlaps(appliesToCombinations []string, excludeId string) func(stable.AuthenticationCombinationConfiguration) bool {
	return func(config stable.AuthenticationCombinationConfiguration) bool {
		existing, ok := config.(stable.X509CertificateCombinationConfiguration)
		if !ok || pointer.From(existing.Id) == excludeId {
			return false
		}

		for _, combination := range pointer.From(existing.AppliesToCombinations) {
			if slices.Contains(appliesToCombinations, string(combination)) {
				return true
			}
		}

		return false
	}
}

func expandX509CombinationConfiguration(model AuthenticationStrengthPolicyX509CombinationConfigurationResourceModel) stable.X509CertificateCombinationConfiguration {
	appliesToCombinations := make([]stable.AuthenticationMethodModes, 0)
	for _, v := range model.AppliesToCombinations {
		appliesToCombinations = append(appliesToCombinations, stable.AuthenticationMethodModes(v))
	}

	allowedIssuerSkis := model.AllowedIssuerSkis
	if allowedIssuerSkis == nil {
		allowedIssuerSkis = []string{}
	}

	allowedPolicyOIDs := model.AllowedPolicyOIDs
	if allowedPolicyOIDs == nil {
		allowedPolicyOIDs = []string{}
	}

	return stable.X509CertificateCombinationConfiguration{
		AppliesToCombinations: pointer.To(appliesToCombinations),
		AllowedIssuerSkis:     pointer.To(allowedIssuerSkis),
		AllowedPolicyOIDs:     pointer.To(allowedPolicyOIDs),
	}
}
