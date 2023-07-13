// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package administrativeunits_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
)

type AdministrativeUnitDataSource struct{}

func TestAccAdministrativeUnitDataSource_byDisplayName(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_administrative_unit", "test")

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: AdministrativeUnitDataSource{}.displayName(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctestAdministrativeUnit-%d", data.RandomInteger)),
			),
		},
	})
}

func TestAccAdministrativeUnitDataSource_byObjectId(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_administrative_unit", "test")

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: AdministrativeUnitDataSource{}.objectId(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctestAdministrativeUnit-%d", data.RandomInteger)),
			),
		},
	})
}

func TestAccAdministrativeUnitDataSource_members(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_administrative_unit", "test")

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: AdministrativeUnitDataSource{}.members(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctestAdministrativeUnit-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("members.#").HasValue("4"),
			),
		},
	})
}

func (AdministrativeUnitDataSource) displayName(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_administrative_unit" "test" {
  display_name = azuread_administrative_unit.test.display_name
}
`, AdministrativeUnitResource{}.basic(data))
}

func (AdministrativeUnitDataSource) objectId(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_administrative_unit" "test" {
  object_id = azuread_administrative_unit.test.object_id
}
`, AdministrativeUnitResource{}.basic(data))
}

func (AdministrativeUnitDataSource) members(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_administrative_unit" "test" {
  object_id = azuread_administrative_unit.test.object_id
}
`, AdministrativeUnitResource{}.withMembers(data))
}
