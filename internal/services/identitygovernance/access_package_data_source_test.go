// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identitygovernance_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
)

type AccessPackageDataSource struct{}

func TestAccAccessPackageDataSource_byId(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_access_package", "test")
	r := AccessPackageDataSource{}

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: r.byId(data),
			Check:  r.testCheckFunc(data),
		},
	})
}

func TestAccAccessPackageDataSource_byDisplayName(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_access_package", "test")
	r := AccessPackageDataSource{}

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: r.byDisplayName(data),
			Check:  r.testCheckFunc(data),
		},
	})
}

func (AccessPackageDataSource) testCheckFunc(data acceptance.TestData) acceptance.TestCheckFunc {
	return acceptance.ComposeTestCheckFunc(
		check.That(data.ResourceName).Key("description").HasValue(fmt.Sprintf("Access Package %[1]d", data.RandomInteger)),
		check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("access-package-%[1]d", data.RandomInteger)),
		check.That(data.ResourceName).Key("hidden").HasValue("true"),
		check.That(data.ResourceName).Key("catalog_id").Exists(),
	)
}

func (AccessPackageDataSource) byId(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_access_package" "test" {
  object_id = azuread_access_package.test.id
}
`, AccessPackageResource{}.complete(data))
}

func (AccessPackageDataSource) byDisplayName(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_access_package" "test" {
  display_name = azuread_access_package.test.display_name
  catalog_id   = azuread_access_package_catalog.test_catalog.id
}
`, AccessPackageResource{}.complete(data))
}
