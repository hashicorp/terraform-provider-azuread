package directoryroles_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
)

type DirectoryRolesDataSource struct{}

func TestAccDirectoryRolesDataSource_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_directory_roles", "test")
	r := DirectoryRolesDataSource{}

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: r.basic(),
			Check:  r.testCheckFunc(data),
		},
	})
}

func (DirectoryRolesDataSource) testCheckFunc(data acceptance.TestData, additionalChecks ...resource.TestCheckFunc) resource.TestCheckFunc {
	checks := []resource.TestCheckFunc{
		check.That(data.ResourceName).Key("roles.0.description").Exists(),
		check.That(data.ResourceName).Key("roles.0.display_name").Exists(),
		check.That(data.ResourceName).Key("roles.0.object_id").Exists(),
		check.That(data.ResourceName).Key("roles.0.template_id").Exists(),
		check.That(data.ResourceName).Key("object_ids.#").Exists(),
		check.That(data.ResourceName).Key("template_ids.#").Exists(),
	}
	checks = append(checks, additionalChecks...)
	return resource.ComposeTestCheckFunc(checks...)
}

func (DirectoryRolesDataSource) basic() string {
	return `data "azuread_directory_roles" "test" {}`
}
