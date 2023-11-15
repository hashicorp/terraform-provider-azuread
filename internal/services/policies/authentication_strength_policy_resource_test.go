// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package policies_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"
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

func (r AuthenticationStrengthPolicyResource) Exists(ctx context.Context, client *clients.Client, state *pluginsdk.InstanceState) (*bool, error) {
	var id *string

	authstrengthpolicy, status, err := client.Policies.AuthenticationStrengthPoliciesClient.Get(ctx, state.ID, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return nil, fmt.Errorf("Authentication Strength Policy with ID %q does not exist", state.ID)
		}
		return nil, fmt.Errorf("failed to retrieve Authentication Strength Policy with ID %q: %+v", state.ID, err)
	}
	id = authstrengthpolicy.ID

	return pointer.To(id != nil && *id == state.ID), nil
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
    "fido2",
    "password",
    "deviceBasedPush",
    "temporaryAccessPassOneTime",
    "federatedMultiFactor",
    "federatedSingleFactor",
    "hardwareOath,federatedSingleFactor",
    "microsoftAuthenticatorPush,federatedSingleFactor",
    "password,hardwareOath",
    "password,microsoftAuthenticatorPush",
    "password,sms",
    "password,softwareOath",
    "password,voice",
    "sms",
    "sms,federatedSingleFactor",
    "softwareOath,federatedSingleFactor",
    "temporaryAccessPassMultiUse",
    "voice,federatedSingleFactor",
    "windowsHelloForBusiness",
    "x509CertificateMultiFactor",
    "x509CertificateSingleFactor",
  ]
}
`, data.RandomInteger)
}
