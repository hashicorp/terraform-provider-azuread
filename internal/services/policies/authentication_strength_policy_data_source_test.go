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

func TestAccAuthenticationStrengthPolicyDataSource_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_authentication_strength_policy", "test")

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: AuthenticationStrengthPolicyDataSource{}.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("display_name").Exists(),
				check.That(data.ResourceName).Key("description").Exists(),
				check.That(data.ResourceName).Key("allowed_combinations.#").HasValue("1"),
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
				check.That(data.ResourceName).Key("allowed_combinations.#").Exists(),
			),
		},
	})
}

func (AuthenticationStrengthPolicyDataSource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_authentication_strength_policy" "test" {
  display_name = azuread_authentication_strength_policy.test.display_name
}
`, AuthenticationStrengthPolicyResource{}.basic(data))
}

func (AuthenticationStrengthPolicyDataSource) builtIn() string {
	return `
provider "azuread" {}

data "azuread_authentication_strength_policy" "test" {
  display_name = "Multifactor authentication"
}
`
}
