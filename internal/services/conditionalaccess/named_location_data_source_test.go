// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package conditionalaccess_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
)

type NamedLocationDataSource struct{}

func TestAccNamedLocationDataSource_country(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_named_location", "test")

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: NamedLocationDataSource{}.country(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("country.#").HasValue("1"),
				check.That(data.ResourceName).Key("country.0.countries_and_regions.#").HasValue("3"),
				check.That(data.ResourceName).Key("country.0.include_unknown_countries_and_regions").HasValue("true"),
			),
		},
	})
}

func TestAccNamedLocationDataSource_ip(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_named_location", "test")

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: NamedLocationDataSource{}.ip(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("ip.#").HasValue("1"),
				check.That(data.ResourceName).Key("ip.0.ip_ranges.#").HasValue("4"),
				check.That(data.ResourceName).Key("ip.0.trusted").HasValue("true"),
			),
		},
	})
}

func (NamedLocationDataSource) country(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_named_location" "test" {
  display_name = azuread_named_location.test.display_name
}
`, NamedLocationResource{}.completeCountry(data))
}

func (NamedLocationDataSource) ip(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_named_location" "test" {
  display_name = azuread_named_location.test.display_name
}
`, NamedLocationResource{}.completeIP(data))
}
