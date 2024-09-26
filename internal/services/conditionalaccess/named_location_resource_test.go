// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package conditionalaccess_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/conditionalaccessnamedlocation"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
)

type NamedLocationResource struct{}

func TestAccNamedLocation_basicIP(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_named_location", "test")
	r := NamedLocationResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basicIP(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccNamedLocation_completeIP(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_named_location", "test")
	r := NamedLocationResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.completeIP(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccNamedLocation_updateIP(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_named_location", "test")
	r := NamedLocationResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basicIP(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.completeIP(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.basicIP(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccNamedLocation_basicCountry(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_named_location", "test")
	r := NamedLocationResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basicCountry(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccNamedLocation_completeCountry(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_named_location", "test")
	r := NamedLocationResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.completeCountry(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccNamedLocation_updateCountry(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_named_location", "test")
	r := NamedLocationResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basicCountry(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.completeCountry(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.basicCountry(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func (r NamedLocationResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.ConditionalAccess.NamedLocationClient

	id, err := stable.ParseIdentityConditionalAccessNamedLocationID(state.ID)
	if err != nil {
		return nil, err
	}

	resp, err := client.GetConditionalAccessNamedLocation(ctx, *id, conditionalaccessnamedlocation.DefaultGetConditionalAccessNamedLocationOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return pointer.To(false), nil
		}
		return nil, fmt.Errorf("failed to retrieve Named Location with object ID %q: %+v", state.ID, err)
	}

	return pointer.To(true), nil
}

func (NamedLocationResource) basicIP(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_named_location" "test" {
  display_name = "acctestNLIP-%[1]d"
  ip {
    ip_ranges = [
      "1.1.1.1/32",
      "2.2.2.2/32",
    ]
  }
}
`, data.RandomInteger)
}

func (NamedLocationResource) completeIP(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_named_location" "test" {
  display_name = "acctestNLIP-%[1]d"
  ip {
    ip_ranges = [
      "1.1.1.1/32",
      "2.2.2.2/32",
      "3.3.3.3/32",
      "64:ff9b::/96",
    ]
    trusted = true
  }
}
`, data.RandomInteger)
}

func (NamedLocationResource) basicCountry(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_named_location" "test" {
  display_name = "acctestNLC-%[1]d"
  country {
    countries_and_regions = [
      "GB",
      "US",
    ]
  }
}
`, data.RandomInteger)
}

func (NamedLocationResource) completeCountry(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_named_location" "test" {
  display_name = "acctestNLC-%[1]d"
  country {
    countries_and_regions = [
      "GB",
      "US",
      "JP",
    ]
    include_unknown_countries_and_regions = true
  }
}
`, data.RandomInteger)
}
