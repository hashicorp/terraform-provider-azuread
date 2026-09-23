// Copyright IBM Corp. 2019, 2026
// SPDX-License-Identifier: MPL-2.0

package policies

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/authenticationstrengthpolicy"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/sdk"
)

var _ sdk.DataSource = AuthenticationStrengthPolicyDataSource{}

const (
	combinationConfigurationTypeFido2           = "fido2CombinationConfiguration"
	combinationConfigurationTypeX509Certificate = "x509CertificateCombinationConfiguration"
)

type AuthenticationStrengthPolicyDataSourceModel struct {
	ObjectId                  string                                                      `tfschema:"object_id"`
	DisplayName               string                                                      `tfschema:"display_name"`
	Description               string                                                      `tfschema:"description"`
	AllowedCombinations       []string                                                    `tfschema:"allowed_combinations"`
	CombinationConfigurations []AuthenticationStrengthPolicyCombinationConfigurationModel `tfschema:"combination_configurations"`
}

type AuthenticationStrengthPolicyCombinationConfigurationModel struct {
	ObjectId              string   `tfschema:"object_id"`
	Type                  string   `tfschema:"type"`
	AppliesToCombinations []string `tfschema:"applies_to_combinations"`
	AllowedAAGUIDs        []string `tfschema:"allowed_aaguids"`
	AllowedIssuerSkis     []string `tfschema:"allowed_issuer_skis"`
	AllowedPolicyOIDs     []string `tfschema:"allowed_policy_oids"`
}

type AuthenticationStrengthPolicyDataSource struct{}

func (r AuthenticationStrengthPolicyDataSource) Arguments() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"object_id": {
			Description:  "The object ID of the authentication strength policy",
			Type:         pluginsdk.TypeString,
			Optional:     true,
			Computed:     true,
			ExactlyOneOf: []string{"display_name", "object_id"},
			ValidateFunc: validation.IsUUID,
		},

		"display_name": {
			Description:  "The display name for the authentication strength policy",
			Type:         pluginsdk.TypeString,
			Optional:     true,
			Computed:     true,
			ExactlyOneOf: []string{"display_name", "object_id"},
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
			Type:        pluginsdk.TypeList,
			Computed:    true,
			Elem: &pluginsdk.Schema{
				Type: pluginsdk.TypeString,
			},
		},

		"combination_configurations": {
			Description: "The combination configurations which further constrain the allowed combinations for this policy",
			Type:        pluginsdk.TypeList,
			Computed:    true,
			Elem: &pluginsdk.Resource{
				Schema: map[string]*pluginsdk.Schema{
					"object_id": {
						Description: "The object ID of the combination configuration",
						Type:        pluginsdk.TypeString,
						Computed:    true,
					},

					"type": {
						Description: "The type of the combination configuration",
						Type:        pluginsdk.TypeString,
						Computed:    true,
					},

					"applies_to_combinations": {
						Description: "The authentication method combinations this configuration applies to",
						Type:        pluginsdk.TypeList,
						Computed:    true,
						Elem: &pluginsdk.Schema{
							Type: pluginsdk.TypeString,
						},
					},

					"allowed_aaguids": {
						Description: "The AAGUIDs allowed by this configuration, for the `fido2CombinationConfiguration` type",
						Type:        pluginsdk.TypeList,
						Computed:    true,
						Elem: &pluginsdk.Schema{
							Type: pluginsdk.TypeString,
						},
					},

					"allowed_issuer_skis": {
						Description: "The certificate issuer subject key identifiers allowed by this configuration, for the `x509CertificateCombinationConfiguration` type",
						Type:        pluginsdk.TypeList,
						Computed:    true,
						Elem: &pluginsdk.Schema{
							Type: pluginsdk.TypeString,
						},
					},

					"allowed_policy_oids": {
						Description: "The certificate policy OIDs allowed by this configuration, for the `x509CertificateCombinationConfiguration` type",
						Type:        pluginsdk.TypeList,
						Computed:    true,
						Elem: &pluginsdk.Schema{
							Type: pluginsdk.TypeString,
						},
					},
				},
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

			var policy *stable.AuthenticationStrengthPolicy

			if model.ObjectId != "" {
				id := stable.NewPolicyAuthenticationStrengthPolicyID(model.ObjectId)

				resp, err := client.GetAuthenticationStrengthPolicy(ctx, id, authenticationstrengthpolicy.DefaultGetAuthenticationStrengthPolicyOperationOptions())
				if err != nil {
					if response.WasNotFound(resp.HttpResponse) {
						return fmt.Errorf("no authentication strength policy found with object ID %q", model.ObjectId)
					}
					return fmt.Errorf("retrieving %s: %+v", id, err)
				}
				if resp.Model == nil {
					return fmt.Errorf("retrieving %s: API error, model was nil", id)
				}

				policy = resp.Model
			} else {
				resp, err := client.ListAuthenticationStrengthPoliciesComplete(ctx, authenticationstrengthpolicy.DefaultListAuthenticationStrengthPoliciesOperationOptions())
				if err != nil {
					return fmt.Errorf("listing authentication strength policies: %+v", err)
				}

				var matches []stable.AuthenticationStrengthPolicy
				for _, item := range resp.Items {
					if pointer.From(item.DisplayName) == model.DisplayName {
						matches = append(matches, item)
					}
				}

				switch len(matches) {
				case 0:
					return fmt.Errorf("no authentication strength policy found with display name %q", model.DisplayName)
				case 1:
					policy = &matches[0]
				default:
					return fmt.Errorf("multiple authentication strength policies found with display name %q, please specify `object_id` instead", model.DisplayName)
				}
			}

			if policy.Id == nil {
				return errors.New("retrieving authentication strength policy: API error, ID was nil")
			}

			id := stable.NewPolicyAuthenticationStrengthPolicyID(*policy.Id)

			state := AuthenticationStrengthPolicyDataSourceModel{
				ObjectId:                  pointer.From(policy.Id),
				DisplayName:               pointer.From(policy.DisplayName),
				Description:               policy.Description.GetOrZero(),
				AllowedCombinations:       flattenAuthenticationMethodModes(policy.AllowedCombinations),
				CombinationConfigurations: flattenCombinationConfigurations(policy.CombinationConfigurations),
			}

			metadata.ResourceData.SetId(id.ID())
			return metadata.Encode(&state)
		},
	}
}

func flattenAuthenticationMethodModes(input *[]stable.AuthenticationMethodModes) []string {
	output := make([]string, 0)
	for _, v := range pointer.From(input) {
		output = append(output, string(v))
	}
	return output
}

func flattenCombinationConfigurations(input *[]stable.AuthenticationCombinationConfiguration) []AuthenticationStrengthPolicyCombinationConfigurationModel {
	output := make([]AuthenticationStrengthPolicyCombinationConfigurationModel, 0)
	for _, item := range pointer.From(input) {
		base := item.AuthenticationCombinationConfiguration()

		combinationConfiguration := AuthenticationStrengthPolicyCombinationConfigurationModel{
			ObjectId:              pointer.From(base.Id),
			AppliesToCombinations: flattenAuthenticationMethodModes(base.AppliesToCombinations),
			AllowedAAGUIDs:        make([]string, 0),
			AllowedIssuerSkis:     make([]string, 0),
			AllowedPolicyOIDs:     make([]string, 0),
		}

		switch v := item.(type) {
		case stable.Fido2CombinationConfiguration:
			combinationConfiguration.Type = combinationConfigurationTypeFido2
			if v.AllowedAAGUIDs != nil {
				combinationConfiguration.AllowedAAGUIDs = *v.AllowedAAGUIDs
			}
		case stable.X509CertificateCombinationConfiguration:
			combinationConfiguration.Type = combinationConfigurationTypeX509Certificate
			if v.AllowedIssuerSkis != nil {
				combinationConfiguration.AllowedIssuerSkis = *v.AllowedIssuerSkis
			}
			if v.AllowedPolicyOIDs != nil {
				combinationConfiguration.AllowedPolicyOIDs = *v.AllowedPolicyOIDs
			}
		default:
			// The SDK returns a raw type for discriminators it doesn't model, so fall back to the raw @odata.type
			combinationConfiguration.Type = strings.TrimPrefix(pointer.From(base.ODataType), "#microsoft.graph.")
		}

		output = append(output, combinationConfiguration)
	}

	return output
}
