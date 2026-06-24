// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package policies_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/authenticationstrengthpolicy"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
)

type AuthenticationStrengthPolicyResource struct{}

func TestAccAuthenticationStrengthPolicy_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_authentication_strength_policy", "test")
	r := AuthenticationStrengthPolicyResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccAuthenticationStrengthPolicy_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_authentication_strength_policy", "test")
	r := AuthenticationStrengthPolicyResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.complete(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccAuthenticationStrengthPolicy_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_authentication_strength_policy", "test")
	r := AuthenticationStrengthPolicyResource{}

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

func TestAccAuthenticationStrengthPolicy_fido2CombinationConfiguration(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_authentication_strength_policy", "test")
	r := AuthenticationStrengthPolicyResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.fido2CombinationConfiguration(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("fido2_combination_configuration.0.id").Exists(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccAuthenticationStrengthPolicy_x509CombinationConfiguration(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_authentication_strength_policy", "test")
	r := AuthenticationStrengthPolicyResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.x509CombinationConfiguration(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("x509_certificate_combination_configuration.0.id").Exists(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccAuthenticationStrengthPolicy_combinationConfigurationUpdate(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_authentication_strength_policy", "test")
	r := AuthenticationStrengthPolicyResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.fido2(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.fido2CombinationConfiguration(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("fido2_combination_configuration.0.id").Exists(),
			),
		},
		data.ImportStep(),
		{
			Config: r.fido2CombinationConfigurationUpdated(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.fido2(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func (r AuthenticationStrengthPolicyResource) Exists(ctx context.Context, clients *clients.Client, state *pluginsdk.InstanceState) (*bool, error) {
	client := clients.Policies.AuthenticationStrengthPolicyClient

	id, err := stable.ParsePolicyAuthenticationStrengthPolicyID(state.ID)
	if err != nil {
		return nil, err
	}

	resp, err := client.GetAuthenticationStrengthPolicy(ctx, *id, authenticationstrengthpolicy.DefaultGetAuthenticationStrengthPolicyOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return pointer.To(false), nil
		}
		return nil, fmt.Errorf("failed to retrieve %s: %v", id, err)
	}

	return pointer.To(true), nil
}

func (AuthenticationStrengthPolicyResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_authentication_strength_policy" "test" {
  display_name         = "acctestASP-%[1]d"
  description          = "test"
  allowed_combinations = ["password"]
}
`, data.RandomInteger)
}

func (AuthenticationStrengthPolicyResource) complete(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_authentication_strength_policy" "test" {
  display_name = "acctestASP-%[1]d"
  description  = "test"
  allowed_combinations = [
    "deviceBasedPush",
    "federatedMultiFactor",
    "federatedSingleFactor",
    "fido2",
    "hardwareOath,federatedSingleFactor",
    "microsoftAuthenticatorPush,federatedSingleFactor",
    "password",
    "password,hardwareOath",
    "password,microsoftAuthenticatorPush",
    "password,sms",
    "password,softwareOath",
    "password,voice",
    "sms",
    "sms,federatedSingleFactor",
    "softwareOath,federatedSingleFactor",
    "temporaryAccessPassMultiUse",
    "temporaryAccessPassOneTime",
    "voice,federatedSingleFactor",
    "windowsHelloForBusiness",
    "x509CertificateMultiFactor",
    "x509CertificateSingleFactor",
  ]
}
`, data.RandomInteger)
}

func (AuthenticationStrengthPolicyResource) fido2(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_authentication_strength_policy" "test" {
  display_name         = "acctestASP-%[1]d"
  description          = "test"
  allowed_combinations = ["fido2"]
}
`, data.RandomInteger)
}

func (AuthenticationStrengthPolicyResource) fido2CombinationConfiguration(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_authentication_strength_policy" "test" {
  display_name         = "acctestASP-%[1]d"
  description          = "test"
  allowed_combinations = ["fido2"]

  fido2_combination_configuration {
    allowed_aaguids = [
      "de1e552d-db1d-4423-a619-566b625cdc84",
      "90a3ccdf-635c-4729-a248-9b709135078f",
    ]
  }
}
`, data.RandomInteger)
}

func (AuthenticationStrengthPolicyResource) fido2CombinationConfigurationUpdated(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_authentication_strength_policy" "test" {
  display_name         = "acctestASP-%[1]d"
  description          = "test"
  allowed_combinations = ["fido2"]

  fido2_combination_configuration {
    allowed_aaguids = [
      "de1e552d-db1d-4423-a619-566b625cdc84",
    ]
  }
}
`, data.RandomInteger)
}

func (AuthenticationStrengthPolicyResource) x509CombinationConfiguration(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_authentication_strength_policy" "test" {
  display_name = "acctestASP-%[1]d"
  description  = "test"
  allowed_combinations = [
    "x509CertificateSingleFactor",
    "x509CertificateMultiFactor",
  ]

  x509_certificate_combination_configuration {
    applies_to_combinations = [
      "x509CertificateSingleFactor",
      "x509CertificateMultiFactor",
    ]
    allowed_issuer_skis = ["9af52a26d8e4bd7d5e8f43e9c7c5e2f4a3b1c0d9"]
    allowed_policy_oids = ["1.2.3.4.5"]
  }
}
`, data.RandomInteger)
}
