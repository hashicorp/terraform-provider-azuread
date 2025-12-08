// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package directoryroles_test

import (
	"testing"

	"github.com/glueckkanja/terraform-provider-azuread/internal/acceptance"
	"github.com/glueckkanja/terraform-provider-azuread/internal/acceptance/check"
)

type DirectoryRoleTemplatesDataSource struct{}

func TestAccDirectoryRoleTemplatesDataSource_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_directory_role_templates", "test")
	r := DirectoryRoleTemplatesDataSource{}

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: r.basic(),
			Check:  r.testCheckFunc(data),
		},
	})
}

func (DirectoryRoleTemplatesDataSource) testCheckFunc(data acceptance.TestData, additionalChecks ...acceptance.TestCheckFunc) acceptance.TestCheckFunc {
	checks := []acceptance.TestCheckFunc{
		check.That(data.ResourceName).Key("role_templates.0.description").Exists(),
		check.That(data.ResourceName).Key("role_templates.0.display_name").Exists(),
		check.That(data.ResourceName).Key("role_templates.0.object_id").Exists(),
		check.That(data.ResourceName).Key("object_ids.#").Exists(),
	}
	checks = append(checks, additionalChecks...)
	return acceptance.ComposeTestCheckFunc(checks...)
}

func (DirectoryRoleTemplatesDataSource) basic() string {
	return `data "azuread_directory_role_templates" "test" {}`
}
