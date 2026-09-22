// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package policies_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
)

type AuthenticationStrengthPolicyDataSource struct{}

func TestAccAuthenticationStrengthPolicyDataSource_displayName(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_authentication_strength_policy", "test")

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: AuthenticationStrengthPolicyDataSource{}.displayName(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("object_id").Exists(),
				check.That(data.ResourceName).Key("display_name").Exists(),
				check.That(data.ResourceName).Key("description").Exists(),
				check.That(data.ResourceName).Key("allowed_combinations.#").HasValue("1"),
				check.That(data.ResourceName).Key("combination_configurations.#").HasValue("0"),
			),
		},
	})
}

func TestAccAuthenticationStrengthPolicyDataSource_objectId(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_authentication_strength_policy", "test")

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: AuthenticationStrengthPolicyDataSource{}.objectId(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("object_id").Exists(),
				check.That(data.ResourceName).Key("display_name").Exists(),
				check.That(data.ResourceName).Key("description").Exists(),
				check.That(data.ResourceName).Key("allowed_combinations.#").HasValue("1"),
				check.That(data.ResourceName).Key("combination_configurations.#").HasValue("0"),
			),
		},
	})
}

func TestAccAuthenticationStrengthPolicyDataSource_builtIn(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_authentication_strength_policy", "test")

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: AuthenticationStrengthPolicyDataSource{}.builtIn(),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("display_name").HasValue("Multifactor authentication"),
				check.That(data.ResourceName).Key("object_id").Exists(),
				check.That(data.ResourceName).Key("allowed_combinations.#").Exists(),
			),
		},
	})
}

func TestAccAuthenticationStrengthPolicyDataSource_combinationConfigurations(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_authentication_strength_policy", "test")

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: AuthenticationStrengthPolicyDataSource{}.combinationConfigurations(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("combination_configurations.#").HasValue("1"),
				check.That(data.ResourceName).Key("combination_configurations.0.object_id").Exists(),
				check.That(data.ResourceName).Key("combination_configurations.0.type").HasValue("fido2"),
				check.That(data.ResourceName).Key("combination_configurations.0.applies_to_combinations.#").HasValue("1"),
				check.That(data.ResourceName).Key("combination_configurations.0.applies_to_combinations.0").HasValue("fido2"),
				check.That(data.ResourceName).Key("combination_configurations.0.allowed_aaguids.#").HasValue("1"),
				check.That(data.ResourceName).Key("combination_configurations.0.allowed_aaguids.0").HasValue("de1e552d-db1d-4423-a619-566b625cdc84"),
				check.That(data.ResourceName).Key("combination_configurations.0.allowed_issuer_skis.#").HasValue("0"),
				check.That(data.ResourceName).Key("combination_configurations.0.allowed_policy_oids.#").HasValue("0"),
			),
		},
	})
}

func (AuthenticationStrengthPolicyDataSource) displayName(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_authentication_strength_policy" "test" {
  display_name = azuread_authentication_strength_policy.test.display_name
}
`, AuthenticationStrengthPolicyResource{}.basic(data))
}

func (AuthenticationStrengthPolicyDataSource) objectId(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_authentication_strength_policy" "test" {
  object_id = azuread_authentication_strength_policy.test.object_id
}
`, AuthenticationStrengthPolicyResource{}.basic(data))
}

func (AuthenticationStrengthPolicyDataSource) combinationConfigurations(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_authentication_strength_policy" "test" {
  display_name         = "acctestASP-%[1]d"
  description          = "test"
  allowed_combinations = ["fido2"]
}

resource "azuread_authentication_strength_policy_fido2_combination_configuration" "test" {
  authentication_strength_policy_id = azuread_authentication_strength_policy.test.object_id
  allowed_aaguids                   = ["de1e552d-db1d-4423-a619-566b625cdc84"]
}

data "azuread_authentication_strength_policy" "test" {
  object_id = azuread_authentication_strength_policy.test.object_id

  depends_on = [azuread_authentication_strength_policy_fido2_combination_configuration.test]
}
`, data.RandomInteger)
}

func (AuthenticationStrengthPolicyDataSource) builtIn() string {
	return `
provider "azuread" {}

data "azuread_authentication_strength_policy" "test" {
  display_name = "Multifactor authentication"
}
`
}
