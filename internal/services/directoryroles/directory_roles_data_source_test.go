// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package directoryroles_test

import (
	"testing"

	"github.com/glueckkanja/terraform-provider-azuread/internal/acceptance"
	"github.com/glueckkanja/terraform-provider-azuread/internal/acceptance/check"
)

type DirectoryRolesDataSource struct{}

func TestAccDirectoryRolesDataSource_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_directory_roles", "test")
	r := DirectoryRolesDataSource{}

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: r.basic(),
			Check:  r.testCheckFunc(data),
		},
	})
}

func (DirectoryRolesDataSource) testCheckFunc(data acceptance.TestData, additionalChecks ...acceptance.TestCheckFunc) acceptance.TestCheckFunc {
	checks := []acceptance.TestCheckFunc{
		check.That(data.ResourceName).Key("roles.0.description").Exists(),
		check.That(data.ResourceName).Key("roles.0.display_name").Exists(),
		check.That(data.ResourceName).Key("roles.0.object_id").Exists(),
		check.That(data.ResourceName).Key("roles.0.template_id").Exists(),
		check.That(data.ResourceName).Key("object_ids.#").Exists(),
		check.That(data.ResourceName).Key("template_ids.#").Exists(),
	}
	checks = append(checks, additionalChecks...)
	return acceptance.ComposeTestCheckFunc(checks...)
}

func (DirectoryRolesDataSource) basic() string {
	return `data "azuread_directory_roles" "test" {}`
}
