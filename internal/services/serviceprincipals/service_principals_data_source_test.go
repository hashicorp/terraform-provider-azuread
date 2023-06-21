// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serviceprincipals_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
)

type ServicePrincipalsDataSource struct{}

func TestAccServicePrincipalsDataSource_byApplicationIds(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_service_principals", "test")

	data.DataSourceTest(t, []resource.TestStep{{
		Config: ServicePrincipalsDataSource{}.byApplicationIds(data),
		Check: resource.ComposeTestCheckFunc(
			check.That(data.ResourceName).Key("application_ids.#").HasValue("2"),
			check.That(data.ResourceName).Key("display_names.#").HasValue("2"),
			check.That(data.ResourceName).Key("object_ids.#").HasValue("2"),
			check.That(data.ResourceName).Key("service_principals.#").HasValue("2"),
		),
	}})
}

func TestAccServicePrincipalsDataSource_byApplicationIdsWithIgnoreMissing(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_service_principals", "test")

	data.DataSourceTest(t, []resource.TestStep{{
		Config: ServicePrincipalsDataSource{}.byApplicationIdsWithIgnoreMissing(data),
		Check: resource.ComposeTestCheckFunc(
			check.That(data.ResourceName).Key("application_ids.#").HasValue("2"),
			check.That(data.ResourceName).Key("display_names.#").HasValue("2"),
			check.That(data.ResourceName).Key("object_ids.#").HasValue("2"),
			check.That(data.ResourceName).Key("service_principals.#").HasValue("2"),
		),
	}})
}

func TestAccServicePrincipalsDataSource_byDisplayNames(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_service_principals", "test")

	data.DataSourceTest(t, []resource.TestStep{{
		Config: ServicePrincipalsDataSource{}.byDisplayNames(data),
		Check: resource.ComposeTestCheckFunc(
			check.That(data.ResourceName).Key("display_names.#").HasValue("2"),
			check.That(data.ResourceName).Key("object_ids.#").HasValue("2"),
			check.That(data.ResourceName).Key("service_principals.#").HasValue("2"),
		),
	}})
}

func TestAccServicePrincipalsDataSource_byDisplayNamesWithIgnoreMissing(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_service_principals", "test")

	data.DataSourceTest(t, []resource.TestStep{{
		Config: ServicePrincipalsDataSource{}.byDisplayNamesWithIgnoreMissing(data),
		Check: resource.ComposeTestCheckFunc(
			check.That(data.ResourceName).Key("display_names.#").HasValue("3"),
			check.That(data.ResourceName).Key("object_ids.#").HasValue("3"),
			check.That(data.ResourceName).Key("service_principals.#").HasValue("3"),
		),
	}})
}

func TestAccServicePrincipalsDataSource_byObjectIds(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_service_principals", "test")

	data.DataSourceTest(t, []resource.TestStep{{
		Config: ServicePrincipalsDataSource{}.byObjectIds(data),
		Check: resource.ComposeTestCheckFunc(
			check.That(data.ResourceName).Key("display_names.#").HasValue("2"),
			check.That(data.ResourceName).Key("object_ids.#").HasValue("2"),
			check.That(data.ResourceName).Key("service_principals.#").HasValue("2"),
		),
	}})
}

func TestAccServicePrincipalsDataSource_byObjectIdsWithIgnoreMissing(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_service_principals", "test")

	data.DataSourceTest(t, []resource.TestStep{{
		Config: ServicePrincipalsDataSource{}.byObjectIdsWithIgnoreMissing(data),
		Check: resource.ComposeTestCheckFunc(
			check.That(data.ResourceName).Key("display_names.#").HasValue("2"),
			check.That(data.ResourceName).Key("object_ids.#").HasValue("2"),
			check.That(data.ResourceName).Key("service_principals.#").HasValue("2"),
		),
	}})
}

func TestAccServicePrincipalsDataSource_noNames(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_service_principals", "test")

	data.DataSourceTest(t, []resource.TestStep{{
		Config: ServicePrincipalsDataSource{}.noNames(),
		Check: resource.ComposeTestCheckFunc(
			check.That(data.ResourceName).Key("application_ids.#").HasValue("0"),
			check.That(data.ResourceName).Key("display_names.#").HasValue("0"),
			check.That(data.ResourceName).Key("object_ids.#").HasValue("0"),
			check.That(data.ResourceName).Key("service_principals.#").HasValue("0"),
		),
	}})
}

func TestAccServicePrincipalsDataSource_returnAll(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_service_principals", "test")

	data.DataSourceTest(t, []resource.TestStep{{
		Config: ServicePrincipalsDataSource{}.returnAll(),
		Check: resource.ComposeTestCheckFunc(
			check.That(data.ResourceName).Key("application_ids.#").Exists(),
			check.That(data.ResourceName).Key("display_names.#").Exists(),
			check.That(data.ResourceName).Key("object_ids.#").Exists(),
			check.That(data.ResourceName).Key("service_principals.#").Exists(),
		),
	}})
}

func (ServicePrincipalsDataSource) byDisplayNames(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_service_principals" "test" {
  display_names = [
    azuread_service_principal.testA.display_name,
    azuread_service_principal.testB.display_name,
  ]
}
`, ServicePrincipalResource{}.threeServicePrincipalsABC(data))
}

func (ServicePrincipalsDataSource) byDisplayNamesWithIgnoreMissing(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_service_principals" "test" {
  ignore_missing = true

  display_names = [
    azuread_service_principal.testA.display_name,
    "not-a-real-service_principal-%[2]d-g1bb3r1sh",
    azuread_service_principal.testB.display_name,
    azuread_service_principal.testC.display_name,
  ]
}
`, ServicePrincipalResource{}.threeServicePrincipalsABC(data), data.RandomInteger)
}

func (ServicePrincipalsDataSource) byObjectIds(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_service_principals" "test" {
  object_ids = [
    azuread_service_principal.testA.object_id,
    azuread_service_principal.testB.object_id,
  ]
}
`, ServicePrincipalResource{}.threeServicePrincipalsABC(data))
}

func (ServicePrincipalsDataSource) byObjectIdsWithIgnoreMissing(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_service_principals" "test" {
  ignore_missing = true

  object_ids = [
    azuread_service_principal.testA.object_id,
    "f0000000-0000-0000-0000-000000000000",
    azuread_service_principal.testB.object_id,
  ]
}
`, ServicePrincipalResource{}.threeServicePrincipalsABC(data))
}

func (ServicePrincipalsDataSource) byApplicationIds(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_service_principals" "test" {
  application_ids = [
    azuread_service_principal.testA.application_id,
    azuread_service_principal.testB.application_id,
  ]
}
`, ServicePrincipalResource{}.threeServicePrincipalsABC(data))
}

func (ServicePrincipalsDataSource) byApplicationIdsWithIgnoreMissing(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_service_principals" "test" {
  ignore_missing = true

  application_ids = [
    azuread_service_principal.testA.application_id,
    "e0000000-0000-0000-0000-000000000000",
    azuread_service_principal.testB.application_id,
  ]
}
`, ServicePrincipalResource{}.threeServicePrincipalsABC(data), data.RandomInteger)
}

func (ServicePrincipalsDataSource) noNames() string {
	return `
data "azuread_service_principals" "test" {
  display_names = []
}
`
}

func (ServicePrincipalsDataSource) returnAll() string {
	return `
data "azuread_service_principals" "test" {
  return_all = true
}
`
}
