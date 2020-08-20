package aadgraph_test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-azuread/internal/acceptance"
	"github.com/terraform-providers/terraform-provider-azuread/internal/clients"
	"github.com/terraform-providers/terraform-provider-azuread/internal/services/aadgraph/graph"
	"github.com/terraform-providers/terraform-provider-azuread/internal/utils"
)

func testCheckAppRoleExists(name string) resource.TestCheckFunc { //nolint unparam
	return func(s *terraform.State) error {
		client := acceptance.AzureADProvider.Meta().(*clients.AadClient).AadGraph.ApplicationsClient
		ctx := acceptance.AzureADProvider.Meta().(*clients.AadClient).StopContext

		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %q", name)
		}

		id, err := graph.ParseAppRoleId(rs.Primary.ID)
		if err != nil {
			return fmt.Errorf("parsing App Role ID: %v", err)
		}
		resp, err := client.Get(ctx, id.ObjectId)
		if err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("Bad: Application %q does not exist", id.ObjectId)
			}
			return fmt.Errorf("Bad: Get on applicationsClient: %+v", err)
		}

		role := graph.AppRoleFindByRoleId(resp, id.RoleId)
		if role != nil {
			return nil
		}

		return fmt.Errorf("App Role %q was not found in Application %q", id.RoleId, id.ObjectId)
	}
}

func testCheckAppRoleCheckDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		client := acceptance.AzureADProvider.Meta().(*clients.AadClient).AadGraph.ApplicationsClient
		ctx := acceptance.AzureADProvider.Meta().(*clients.AadClient).StopContext

		if rs.Type != "azuread_application_app_role" {
			continue
		}

		id, err := graph.ParseAppRoleId(rs.Primary.ID)
		if err != nil {
			return fmt.Errorf("parsing App Role ID: %v", err)
		}

		resp, err := client.Get(ctx, id.ObjectId)
		if err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
				return nil
			}

			return err
		}

		role := graph.AppRoleFindByRoleId(resp, id.RoleId)
		if role == nil {
			return nil
		}

		return fmt.Errorf("App Role still exists:\n%#v", resp)
	}

	return nil
}

func TestAccApplicationAppRole_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_app_role", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAppRoleCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccApplicationAppRole_basic(data.RandomInteger),
				Check: resource.ComposeTestCheckFunc(
					testCheckAppRoleExists(data.ResourceName),
				),
			},
			data.ImportStep(),
		},
	})
}

func TestAccApplicationAppRole_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_app_role", "test")
	id := uuid.New().String()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAppRoleCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccApplicationAppRole_complete(data.RandomInteger, id),
				Check: resource.ComposeTestCheckFunc(
					testCheckAppRoleExists(data.ResourceName),
				),
			},
			data.ImportStep(),
		},
	})
}

func TestAccApplicationAppRole_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_app_role", "test")
	id := uuid.New().String()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAppRoleCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccApplicationAppRole_complete(data.RandomInteger, id),
				Check: resource.ComposeTestCheckFunc(
					testCheckAppRoleExists(data.ResourceName),
				),
			},
			data.ImportStep(),
			{
				Config: testAccApplicationAppRole_update(data.RandomInteger, id),
				Check: resource.ComposeTestCheckFunc(
					testCheckAppRoleExists(data.ResourceName),
				),
			},
			data.ImportStep(),
			{
				Config: testAccApplicationAppRole_complete(data.RandomInteger, id),
				Check: resource.ComposeTestCheckFunc(
					testCheckAppRoleExists(data.ResourceName),
				),
			},
			data.ImportStep(),
		},
	})
}

func TestAccApplicationAppRole_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_app_role", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAppRoleCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccApplicationAppRole_basic(data.RandomInteger),
				Check: resource.ComposeTestCheckFunc(
					testCheckAppRoleExists(data.ResourceName),
				),
			},
			data.RequiresImportErrorStep(testAccApplicationAppRole_requiresImport(data.RandomInteger)),
		},
	})
}

func testAccApplicationAppRole_template(ri int) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name = "acctestApp-%d"
}
`, ri)
}

func testAccApplicationAppRole_basic(ri int) string {
	return fmt.Sprintf(`
%s

resource "azuread_application_app_role" "test" {
  application_object_id = azuread_application.test.id
  allowed_member_types = ["User"]
  description          = "Admins can manage roles and perform all task actions"
  display_name         = "Admin"
  is_enabled           = true
}
`, testAccApplicationAppRole_template(ri))
}

func testAccApplicationAppRole_requiresImport(ri int) string {
	return fmt.Sprintf(`
%s

resource "azuread_application_app_role" "import" {
  application_object_id = azuread_application_app_role.test.application_object_id
  allowed_member_types  = azuread_application_app_role.test.allowed_member_types
  description           = azuread_application_app_role.test.description
  display_name          = azuread_application_app_role.test.display_name
  is_enabled            = azuread_application_app_role.test.is_enabled
  role_id               = azuread_application_app_role.test.role_id
}
`, testAccApplicationAppRole_basic(ri))
}

func testAccApplicationAppRole_complete(ri int, id string) string {
	return fmt.Sprintf(`
%s

resource "azuread_application_app_role" "test" {
  application_object_id = azuread_application.test.id
  allowed_member_types = ["User"]
  description          = "Admins can manage roles and perform all task actions"
  display_name         = "Admin"
  is_enabled           = true
  role_id              = "%s"
  value                = "administer"
}
`, testAccApplicationAppRole_template(ri), id)
}

func testAccApplicationAppRole_update(ri int, id string) string {
	return fmt.Sprintf(`
%s

resource "azuread_application_app_role" "test" {
  application_object_id = azuread_application.test.id
  allowed_member_types = ["Application", "User"]
  description          = "Administrators can administrate all the things"
  display_name         = "Administrate"
  is_enabled           = true
  role_id              = "%s"
  value                = "administrate"
}
`, testAccApplicationAppRole_template(ri), id)
}
