// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package policies

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/authenticationstrengthpolicy"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/sdk"
)

var _ sdk.DataSource = AuthenticationStrengthPolicyDataSource{}

type AuthenticationStrengthPolicyDataSourceModel struct {
	DisplayName         string   `tfschema:"display_name"`
	Description         string   `tfschema:"description"`
	AllowedCombinations []string `tfschema:"allowed_combinations"`
}

type AuthenticationStrengthPolicyDataSource struct{}

func (r AuthenticationStrengthPolicyDataSource) Arguments() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"display_name": {
			Description:  "The display name for the authentication strength policy",
			Type:         pluginsdk.TypeString,
			Required:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},
	}
}

func (r AuthenticationStrengthPolicyDataSource) Attributes() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": {
			Description: "The description for the authentication strength policy",
			Type:        pluginsdk.TypeString,
			Computed:    true,
		},

		"allowed_combinations": {
			Description: "The allowed MFA methods for this policy",
			Type:        pluginsdk.TypeSet,
			Computed:    true,
			Elem: &pluginsdk.Schema{
				Type: pluginsdk.TypeString,
			},
		},
	}
}

func (r AuthenticationStrengthPolicyDataSource) ModelObject() interface{} {
	return &AuthenticationStrengthPolicyDataSourceModel{}
}

func (r AuthenticationStrengthPolicyDataSource) ResourceType() string {
	return "azuread_authentication_strength_policy"
}

func (r AuthenticationStrengthPolicyDataSource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Policies.AuthenticationStrengthPolicyClient

			var model AuthenticationStrengthPolicyDataSourceModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			resp, err := client.ListAuthenticationStrengthPolicies(ctx, authenticationstrengthpolicy.DefaultListAuthenticationStrengthPoliciesOperationOptions())
			if err != nil {
				return fmt.Errorf("listing authentication strength policies: %+v", err)
			}
			if resp.Model == nil {
				return fmt.Errorf("listing authentication strength policies: API error, result was nil")
			}

			var matches []stable.AuthenticationStrengthPolicy
			for _, policy := range *resp.Model {
				if pointer.From(policy.DisplayName) == model.DisplayName {
					matches = append(matches, policy)
				}
			}

			switch len(matches) {
			case 0:
				return fmt.Errorf("no authentication strength policy found with display name %q", model.DisplayName)
			case 1:
				// Expected, continue below
			default:
				return fmt.Errorf("multiple authentication strength policies found with display name %q, please ensure the display name is unique", model.DisplayName)
			}

			policy := matches[0]
			if policy.Id == nil {
				return fmt.Errorf("retrieving authentication strength policy with display name %q: API error, ID was nil", model.DisplayName)
			}

			id := stable.NewPolicyAuthenticationStrengthPolicyID(*policy.Id)

			allowedCombinations := make([]string, 0)
			for _, v := range pointer.From(policy.AllowedCombinations) {
				allowedCombinations = append(allowedCombinations, string(v))
			}

			state := AuthenticationStrengthPolicyDataSourceModel{
				DisplayName:         pointer.From(policy.DisplayName),
				Description:         policy.Description.GetOrZero(),
				AllowedCombinations: allowedCombinations,
			}

			metadata.ResourceData.SetId(id.ID())
			return metadata.Encode(&state)
		},
	}
}
