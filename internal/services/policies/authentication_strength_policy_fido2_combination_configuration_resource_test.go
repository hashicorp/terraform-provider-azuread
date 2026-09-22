// Copyright IBM Corp. 2019, 2026
// SPDX-License-Identifier: MPL-2.0

package policies_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/authenticationstrengthpolicycombinationconfiguration"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
)

type AuthenticationStrengthPolicyFido2CombinationConfigurationResource struct{}

func TestAccAuthenticationStrengthPolicyFido2CombinationConfiguration_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_authentication_strength_policy_fido2_combination_configuration", "test")
	r := AuthenticationStrengthPolicyFido2CombinationConfigurationResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("allowed_aaguids.#").HasValue("1"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccAuthenticationStrengthPolicyFido2CombinationConfiguration_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_authentication_strength_policy_fido2_combination_configuration", "test")
	r := AuthenticationStrengthPolicyFido2CombinationConfigurationResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.complete(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("allowed_aaguids.#").HasValue("2"),
			),
		},
		data.ImportStep(),
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("allowed_aaguids.#").HasValue("1"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccAuthenticationStrengthPolicyFido2CombinationConfiguration_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_authentication_strength_policy_fido2_combination_configuration", "test")
	r := AuthenticationStrengthPolicyFido2CombinationConfigurationResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.RequiresImportErrorStep(r.requiresImport(data)),
	})
}

func (r AuthenticationStrengthPolicyFido2CombinationConfigurationResource) Exists(ctx context.Context, clients *clients.Client, state *pluginsdk.InstanceState) (*bool, error) {
	client := clients.Policies.AuthenticationStrengthPolicyCombinationConfigurationClient

	id, err := stable.ParsePolicyAuthenticationStrengthPolicyIdCombinationConfigurationID(state.ID)
	if err != nil {
		return nil, err
	}

	resp, err := client.GetAuthenticationStrengthPolicyCombinationConfiguration(ctx, *id, authenticationstrengthpolicycombinationconfiguration.DefaultGetAuthenticationStrengthPolicyCombinationConfigurationOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return pointer.To(false), nil
		}
		return nil, fmt.Errorf("failed to retrieve %s: %v", id, err)
	}

	return pointer.To(true), nil
}

func (AuthenticationStrengthPolicyFido2CombinationConfigurationResource) template(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_authentication_strength_policy" "test" {
  display_name         = "acctestASP-%[1]d"
  description          = "test"
  allowed_combinations = ["fido2"]
}
`, data.RandomInteger)
}

func (r AuthenticationStrengthPolicyFido2CombinationConfigurationResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_authentication_strength_policy_fido2_combination_configuration" "test" {
  authentication_strength_policy_id = azuread_authentication_strength_policy.test.object_id
  allowed_aaguids                   = ["de1e552d-db1d-4423-a619-566b625cdc84"]
}
`, r.template(data))
}

func (r AuthenticationStrengthPolicyFido2CombinationConfigurationResource) complete(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_authentication_strength_policy_fido2_combination_configuration" "test" {
  authentication_strength_policy_id = azuread_authentication_strength_policy.test.object_id

  allowed_aaguids = [
    "de1e552d-db1d-4423-a619-566b625cdc84",
    "90a3ccdf-635c-4729-a248-9b709135078f",
  ]
}
`, r.template(data))
}

func (r AuthenticationStrengthPolicyFido2CombinationConfigurationResource) requiresImport(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_authentication_strength_policy_fido2_combination_configuration" "import" {
  authentication_strength_policy_id = azuread_authentication_strength_policy_fido2_combination_configuration.test.authentication_strength_policy_id
  allowed_aaguids                   = azuread_authentication_strength_policy_fido2_combination_configuration.test.allowed_aaguids
}
`, r.basic(data))
}
