//enable_ms_graph_api
package directory_roles_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
)

type DirectoryRoleDataSource struct{}

func TestAccDirectoryRoleDataSource_byDisplayName(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_directory_role", "test")

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: DirectoryRoleDataSource{}.displayName(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("display_name").HasValue(data.DirectoryRole.DisplayName),
				check.That(data.ResourceName).Key("description").HasValue(data.DirectoryRole.Description),
				check.That(data.ResourceName).Key("role_template_id").HasValue(data.DirectoryRole.RoleTemplateId),
			),
		},
	})
}

func (DirectoryRoleDataSource) displayName(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_directory_role" "test" {
  display_name = azuread_directory_role.test.display_name
}
`, DirectoryRoleResource{}.basic(data))
}
