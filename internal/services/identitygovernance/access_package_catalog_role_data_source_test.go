// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identitygovernance_test

import (
	"testing"

	"github.com/glueckkanja/terraform-provider-azuread/internal/acceptance"
	"github.com/glueckkanja/terraform-provider-azuread/internal/acceptance/check"
)

type AccessPackageCatalogRoleDataSource struct{}

func TestAccAccessPackageCatalogRoleDataSource_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_access_package_catalog_role", "test")
	r := AccessPackageCatalogRoleDataSource{}

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("display_name").Exists(),
				check.That(data.ResourceName).Key("template_id").Exists(),
				check.That(data.ResourceName).Key("object_id").Exists(),
			),
		},
	})
}

func (AccessPackageCatalogRoleDataSource) basic(_ acceptance.TestData) string {
	return `provider azuread {}
data "azuread_access_package_catalog_role" "test" {
  display_name = "Catalog owner"
}`
}
