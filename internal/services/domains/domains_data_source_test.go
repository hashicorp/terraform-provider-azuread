// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package domains_test

import (
	"testing"

	"github.com/glueckkanja/terraform-provider-azuread/internal/acceptance"
	"github.com/glueckkanja/terraform-provider-azuread/internal/acceptance/check"
)

type DomainsDataSource struct{}

func TestAccDomainsDataSource_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_domains", "test")
	r := DomainsDataSource{}

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: r.basic(),
			Check:  r.testCheckFunc(data),
		},
	})
}

func TestAccDomainsDataSource_onlyDefault(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_domains", "test")
	r := DomainsDataSource{}

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: r.onlyDefault(),
			Check: r.testCheckFunc(data,
				check.That(data.ResourceName).Key("domains.0.default").HasValue("true"),
			),
		},
	})
}

func TestAccDomainsDataSource_onlyInitial(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_domains", "test")
	r := DomainsDataSource{}

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: r.onlyInitial(),
			Check: r.testCheckFunc(data,
				check.That(data.ResourceName).Key("domains.0.initial").HasValue("true"),
			),
		},
	})
}

func TestAccDomainsDataSource_onlyRoot(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_domains", "test")
	r := DomainsDataSource{}

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: r.onlyRoot(),
			Check: r.testCheckFunc(data,
				check.That(data.ResourceName).Key("domains.0.initial").HasValue("true"),
			),
		},
	})
}

func TestAccDomainsDataSource_supportsServices(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_domains", "test")
	r := DomainsDataSource{}

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: DomainsDataSource{}.supportsServices(),
			Check:  r.testCheckFunc(data),
		},
	})
}

func (DomainsDataSource) testCheckFunc(data acceptance.TestData, additionalChecks ...acceptance.TestCheckFunc) acceptance.TestCheckFunc {
	checks := []acceptance.TestCheckFunc{
		check.That(data.ResourceName).Key("domains.0.domain_name").Exists(),
		check.That(data.ResourceName).Key("domains.0.admin_managed").Exists(),
		check.That(data.ResourceName).Key("domains.0.default").Exists(),
		check.That(data.ResourceName).Key("domains.0.initial").Exists(),
		check.That(data.ResourceName).Key("domains.0.root").Exists(),
		check.That(data.ResourceName).Key("domains.0.supported_services.#").Exists(),
		check.That(data.ResourceName).Key("domains.0.verified").Exists(),
	}
	checks = append(checks, additionalChecks...)
	return acceptance.ComposeTestCheckFunc(checks...)
}

func (DomainsDataSource) basic() string {
	return `data "azuread_domains" "test" {}`
}

func (DomainsDataSource) onlyDefault() string {
	return `
data "azuread_domains" "test" {
  only_default = true
}
`
}

func (DomainsDataSource) onlyInitial() string {
	return `
data "azuread_domains" "test" {
  only_initial = true
}
`
}

func (DomainsDataSource) onlyRoot() string {
	return `
data "azuread_domains" "test" {
  only_root = true
}
`
}

func (DomainsDataSource) supportsServices() string {
	return `
data "azuread_domains" "test" {
  supports_services = ["Email", "OfficeCommunicationsOnline"]
}
`
}
