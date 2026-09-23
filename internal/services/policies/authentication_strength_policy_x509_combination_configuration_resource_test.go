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

type AuthenticationStrengthPolicyX509CombinationConfigurationResource struct{}

func TestAccAuthenticationStrengthPolicyX509CombinationConfiguration_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_authentication_strength_policy_x509_combination_configuration", "test")
	r := AuthenticationStrengthPolicyX509CombinationConfigurationResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("applies_to_combinations.#").HasValue("1"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccAuthenticationStrengthPolicyX509CombinationConfiguration_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_authentication_strength_policy_x509_combination_configuration", "test")
	r := AuthenticationStrengthPolicyX509CombinationConfigurationResource{}

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
				check.That(data.ResourceName).Key("allowed_policy_oids.#").HasValue("1"),
			),
		},
		data.ImportStep(),
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("allowed_policy_oids.#").HasValue("0"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccAuthenticationStrengthPolicyX509CombinationConfiguration_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_authentication_strength_policy_x509_combination_configuration", "test")
	r := AuthenticationStrengthPolicyX509CombinationConfigurationResource{}

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

// A policy may carry a separate configuration for each x509 combination.
func TestAccAuthenticationStrengthPolicyX509CombinationConfiguration_multiple(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_authentication_strength_policy_x509_combination_configuration", "test")
	r := AuthenticationStrengthPolicyX509CombinationConfigurationResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.multiple(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That("azuread_authentication_strength_policy_x509_combination_configuration.multi").ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

// Removing a combination configuration alongside its allowed combination relies on Terraform destroying the
// configuration before the policy is updated.
func TestAccAuthenticationStrengthPolicyX509CombinationConfiguration_removeWithCombination(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_authentication_strength_policy_x509_combination_configuration", "test")
	r := AuthenticationStrengthPolicyX509CombinationConfigurationResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.multiple(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That("azuread_authentication_strength_policy_x509_combination_configuration.multi").ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func (r AuthenticationStrengthPolicyX509CombinationConfigurationResource) Exists(ctx context.Context, clients *clients.Client, state *pluginsdk.InstanceState) (*bool, error) {
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

func (AuthenticationStrengthPolicyX509CombinationConfigurationResource) template(data acceptance.TestData, allowedCombinations string) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_authentication_strength_policy" "test" {
  display_name         = "acctestASP-%[1]d"
  description          = "test"
  allowed_combinations = %[2]s
}
`, data.RandomInteger, allowedCombinations)
}

func (r AuthenticationStrengthPolicyX509CombinationConfigurationResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_authentication_strength_policy_x509_combination_configuration" "test" {
  authentication_strength_policy_id = azuread_authentication_strength_policy.test.object_id
  applies_to_combinations           = ["x509CertificateSingleFactor"]
  allowed_issuer_skis               = ["9A4248C6AC8C2931AB2A86537818E92E7B6C97B6"]
}
`, r.template(data, `["x509CertificateSingleFactor"]`))
}

func (r AuthenticationStrengthPolicyX509CombinationConfigurationResource) complete(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_authentication_strength_policy_x509_combination_configuration" "test" {
  authentication_strength_policy_id = azuread_authentication_strength_policy.test.object_id
  applies_to_combinations           = ["x509CertificateSingleFactor"]
  allowed_issuer_skis               = ["9A4248C6AC8C2931AB2A86537818E92E7B6C97B6"]
  allowed_policy_oids               = ["1.2.3.4.5"]
}
`, r.template(data, `["x509CertificateSingleFactor"]`))
}

func (r AuthenticationStrengthPolicyX509CombinationConfigurationResource) multiple(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_authentication_strength_policy_x509_combination_configuration" "test" {
  authentication_strength_policy_id = azuread_authentication_strength_policy.test.object_id
  applies_to_combinations           = ["x509CertificateSingleFactor"]
  allowed_issuer_skis               = ["9A4248C6AC8C2931AB2A86537818E92E7B6C97B6"]
}

resource "azuread_authentication_strength_policy_x509_combination_configuration" "multi" {
  authentication_strength_policy_id = azuread_authentication_strength_policy.test.object_id
  applies_to_combinations           = ["x509CertificateMultiFactor"]
  allowed_policy_oids               = ["1.2.3.4.5"]
}
`, r.template(data, `["x509CertificateSingleFactor", "x509CertificateMultiFactor"]`))
}

func (r AuthenticationStrengthPolicyX509CombinationConfigurationResource) requiresImport(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_authentication_strength_policy_x509_combination_configuration" "import" {
  authentication_strength_policy_id = azuread_authentication_strength_policy_x509_combination_configuration.test.authentication_strength_policy_id
  applies_to_combinations           = azuread_authentication_strength_policy_x509_combination_configuration.test.applies_to_combinations
  allowed_issuer_skis               = azuread_authentication_strength_policy_x509_combination_configuration.test.allowed_issuer_skis
}
`, r.basic(data))
}
