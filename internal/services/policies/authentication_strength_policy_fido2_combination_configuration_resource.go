// Copyright IBM Corp. 2023, 2026
// SPDX-License-Identifier: MPL-2.0

package policies

import (
	"context"
	"fmt"
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

type AuthenticationStrengthPolicyFido2CombinationConfigurationResourceModel struct {
	AuthenticationStrengthPolicyId string   `tfschema:"authentication_strength_policy_id"`
	AllowedAAGUIDs                 []string `tfschema:"allowed_aaguids"`
}

var (
	_ sdk.Resource           = AuthenticationStrengthPolicyFido2CombinationConfigurationResource{}
	_ sdk.ResourceWithUpdate = AuthenticationStrengthPolicyFido2CombinationConfigurationResource{}
)

type AuthenticationStrengthPolicyFido2CombinationConfigurationResource struct{}

func (r AuthenticationStrengthPolicyFido2CombinationConfigurationResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return stable.ValidatePolicyAuthenticationStrengthPolicyIdCombinationConfigurationID
}

func (r AuthenticationStrengthPolicyFido2CombinationConfigurationResource) ResourceType() string {
	return "azuread_authentication_strength_policy_fido2_combination_configuration"
}

func (r AuthenticationStrengthPolicyFido2CombinationConfigurationResource) ModelObject() interface{} {
	return &AuthenticationStrengthPolicyFido2CombinationConfigurationResourceModel{}
}

func (r AuthenticationStrengthPolicyFido2CombinationConfigurationResource) Arguments() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"authentication_strength_policy_id": {
			Description:  "The object ID of the authentication strength policy to which this combination configuration applies",
			Type:         pluginsdk.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.IsUUID,
		},

		"allowed_aaguids": {
			Description: "A set of AAGUIDs allowed to be used as part of the `fido2` combination",
			Type:        pluginsdk.TypeSet,
			Required:    true,
			MinItems:    1,
			Elem: &pluginsdk.Schema{
				Type:         pluginsdk.TypeString,
				ValidateFunc: validation.IsUUID,
			},
		},
	}
}

func (r AuthenticationStrengthPolicyFido2CombinationConfigurationResource) Attributes() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{}
}

func (r AuthenticationStrengthPolicyFido2CombinationConfigurationResource) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Policies.AuthenticationStrengthPolicyCombinationConfigurationClient

			var model AuthenticationStrengthPolicyFido2CombinationConfigurationResourceModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			policyId := stable.NewPolicyAuthenticationStrengthPolicyID(model.AuthenticationStrengthPolicyId)

			tf.LockByName(authenticationStrengthPolicyResourceName, policyId.AuthenticationStrengthPolicyId)
			defer tf.UnlockByName(authenticationStrengthPolicyResourceName, policyId.AuthenticationStrengthPolicyId)

			existingId, err := findCombinationConfiguration(ctx, client, policyId, func(config stable.AuthenticationCombinationConfiguration) bool {
				_, ok := config.(stable.Fido2CombinationConfiguration)
				return ok
			})
			if err != nil {
				return fmt.Errorf("checking for existing fido2 combination configuration for %s: %+v", policyId, err)
			}
			if existingId != "" {
				return metadata.ResourceRequiresImport(r.ResourceType(), stable.NewPolicyAuthenticationStrengthPolicyIdCombinationConfigurationID(policyId.AuthenticationStrengthPolicyId, existingId))
			}

			properties := stable.Fido2CombinationConfiguration{
				AppliesToCombinations: &[]stable.AuthenticationMethodModes{stable.AuthenticationMethodModes_Fido2},
				AllowedAAGUIDs:        pointer.To(model.AllowedAAGUIDs),
			}

			resp, err := client.CreateAuthenticationStrengthPolicyCombinationConfiguration(ctx, policyId, properties, authenticationstrengthpolicycombinationconfiguration.DefaultCreateAuthenticationStrengthPolicyCombinationConfigurationOperationOptions())
			if err != nil {
				return fmt.Errorf("creating fido2 combination configuration for %s: %+v", policyId, err)
			}

			if resp.Model == nil {
				return fmt.Errorf("creating fido2 combination configuration for %s: model was nil", policyId)
			}

			combinationConfigurationId := resp.Model.AuthenticationCombinationConfiguration().Id
			if combinationConfigurationId == nil {
				return fmt.Errorf("creating fido2 combination configuration for %s: model returned with nil ID", policyId)
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

func (r AuthenticationStrengthPolicyFido2CombinationConfigurationResource) Read() sdk.ResourceFunc {
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

			combinationConfiguration, ok := resp.Model.(stable.Fido2CombinationConfiguration)
			if !ok {
				return fmt.Errorf("retrieving %s: not a fido2 combination configuration", id)
			}

			model := AuthenticationStrengthPolicyFido2CombinationConfigurationResourceModel{
				AuthenticationStrengthPolicyId: id.AuthenticationStrengthPolicyId,
				AllowedAAGUIDs:                 pointer.From(combinationConfiguration.AllowedAAGUIDs),
			}

			return metadata.Encode(&model)
		},
	}
}

func (r AuthenticationStrengthPolicyFido2CombinationConfigurationResource) Update() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Policies.AuthenticationStrengthPolicyCombinationConfigurationClient

			id, err := stable.ParsePolicyAuthenticationStrengthPolicyIdCombinationConfigurationID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			var model AuthenticationStrengthPolicyFido2CombinationConfigurationResourceModel
			if err = metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			tf.LockByName(authenticationStrengthPolicyResourceName, id.AuthenticationStrengthPolicyId)
			defer tf.UnlockByName(authenticationStrengthPolicyResourceName, id.AuthenticationStrengthPolicyId)

			properties := stable.Fido2CombinationConfiguration{
				AppliesToCombinations: &[]stable.AuthenticationMethodModes{stable.AuthenticationMethodModes_Fido2},
				AllowedAAGUIDs:        pointer.To(model.AllowedAAGUIDs),
			}

			if _, err = client.UpdateAuthenticationStrengthPolicyCombinationConfiguration(ctx, *id, properties, authenticationstrengthpolicycombinationconfiguration.DefaultUpdateAuthenticationStrengthPolicyCombinationConfigurationOperationOptions()); err != nil {
				return fmt.Errorf("updating %s: %+v", id, err)
			}

			return nil
		},
	}
}

func (r AuthenticationStrengthPolicyFido2CombinationConfigurationResource) Delete() sdk.ResourceFunc {
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
