// enable_ms_graph_api

package directory_roles_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
)

type DirectoryRoleResource struct{}

func TestAccDirectoryRole_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_directory_role", "test")
	r := DirectoryRoleResource{}

	// activated directory role cannot be deactivated or deleted.
	// The `CheckDestroy` function is designed to check if directory role exists after `tf destroy`.
	// In the case of directory roles it will always exist so we don't use this check for Directory Role resources.
	data.ResourceTestWithoutCheckDestroy(t, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("display_name").HasValue(data.DirectoryRole.DisplayName),
				check.That(data.ResourceName).Key("description").HasValue(data.DirectoryRole.Description),
				check.That(data.ResourceName).Key("role_template_id").HasValue(data.DirectoryRole.RoleTemplateId),
			),
		},
		data.ImportStep(),
	})
}

func (r DirectoryRoleResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	var id *string

	dirRole, status, err := clients.DirectoryRoles.MsClient.Get(ctx, state.ID)
	if err != nil {
		if status == http.StatusNotFound {
			return nil, fmt.Errorf("directory role with object ID %q does not exist", state.ID)
		}
		return nil, fmt.Errorf("failed to retrieve directory role with object ID %q: %+v", state.ID, err)
	}
	id = dirRole.ID

	return utils.Bool(id != nil && *id == state.ID), nil
}

func (DirectoryRoleResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_directory_role" "test" {
  display_name = "%s"
}
`, data.DirectoryRole.DisplayName)
}
