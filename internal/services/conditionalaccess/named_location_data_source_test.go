// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package conditionalaccess_test

import (
	"fmt"
	"testing"

	"github.com/glueckkanja/terraform-provider-azuread/internal/acceptance"
	"github.com/glueckkanja/terraform-provider-azuread/internal/acceptance/check"
)

type NamedLocationDataSource struct{}

func TestAccNamedLocationDataSource_country(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_named_location", "test")

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: NamedLocationDataSource{}.country(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("country.#").HasValue("1"),
				check.That(data.ResourceName).Key("country.0.countries_and_regions.#").HasValue("3"),
				check.That(data.ResourceName).Key("country.0.include_unknown_countries_and_regions").HasValue("true"),
				check.That(data.ResourceName).Key("country.0.country_lookup_method").HasValue("clientIpAddress"),
			),
		},
	})
}

func TestAccNamedLocationDataSource_countryByGps(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_named_location", "test")

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: NamedLocationDataSource{}.countryByGps(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("country.#").HasValue("1"),
				check.That(data.ResourceName).Key("country.0.countries_and_regions.#").HasValue("3"),
				check.That(data.ResourceName).Key("country.0.include_unknown_countries_and_regions").HasValue("true"),
				check.That(data.ResourceName).Key("country.0.country_lookup_method").HasValue("authenticatorAppGps"),
			),
		},
	})
}

func TestAccNamedLocationDataSource_ip(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_named_location", "test")

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: NamedLocationDataSource{}.ip(data),
			Check: acceptance.ComposeTestCheckFunc(
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

func (NamedLocationDataSource) countryByGps(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_named_location" "test" {
  display_name = azuread_named_location.test.display_name
}
`, NamedLocationResource{}.completeCountryByGps(data))
}

func (NamedLocationDataSource) ip(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_named_location" "test" {
  display_name = azuread_named_location.test.display_name
}
`, NamedLocationResource{}.completeIP(data))
}
