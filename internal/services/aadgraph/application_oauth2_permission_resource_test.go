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

func testCheckOAuth2PermissionExists(name string) resource.TestCheckFunc { //nolint unparam
	return func(s *terraform.State) error {
		client := acceptance.AzureADProvider.Meta().(*clients.AadClient).AadGraph.ApplicationsClient
		ctx := acceptance.AzureADProvider.Meta().(*clients.AadClient).StopContext

		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %q", name)
		}

		id, err := graph.ParseOAuth2PermissionId(rs.Primary.ID)
		if err != nil {
			return fmt.Errorf("parsing OAuth2 Permission ID: %v", err)
		}
		resp, err := client.Get(ctx, id.ObjectId)
		if err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("Bad: Application %q does not exist", id.ObjectId)
			}
			return fmt.Errorf("Bad: Get on applicationsClient: %+v", err)
		}

		scope, err := graph.OAuth2PermissionFindById(resp, id.PermissionId)
		if err != nil {
			return fmt.Errorf("failed to identity OAuth2 Permission: %s", err)
		} else if scope != nil {
			return nil
		}

		return fmt.Errorf("OAuth2 Permission %q was not found in Application %q", id.PermissionId, id.ObjectId)
	}
}

func testCheckOAuth2PermissionCheckDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		client := acceptance.AzureADProvider.Meta().(*clients.AadClient).AadGraph.ApplicationsClient
		ctx := acceptance.AzureADProvider.Meta().(*clients.AadClient).StopContext

		if rs.Type != "azuread_application_oauth2_permission" {
			continue
		}

		id, err := graph.ParseOAuth2PermissionId(rs.Primary.ID)
		if err != nil {
			return fmt.Errorf("parsing OAuth2 Permission ID: %v", err)
		}

		resp, err := client.Get(ctx, id.ObjectId)
		if err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
				return nil
			}

			return err
		}

		scope, err := graph.OAuth2PermissionFindById(resp, id.PermissionId)
		if err != nil {
			return fmt.Errorf("failed to identity OAuth2 Permission: %s", err)
		} else if scope == nil {
			return nil
		}

		return fmt.Errorf("OAuth2 Permission still exists:\n%#v", resp)
	}

	return nil
}

func TestAccApplicationOAuth2Permission_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_oauth2_permission", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckOAuth2PermissionCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccApplicationOAuth2Permission_basic(data.RandomInteger),
				Check: resource.ComposeTestCheckFunc(
					testCheckOAuth2PermissionExists(data.ResourceName),
				),
			},
			data.ImportStep(),
		},
	})
}

func TestAccApplicationOAuth2Permission_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_oauth2_permission", "test")
	id := uuid.New().String()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckOAuth2PermissionCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccApplicationOAuth2Permission_complete(data.RandomInteger, id),
				Check: resource.ComposeTestCheckFunc(
					testCheckOAuth2PermissionExists(data.ResourceName),
				),
			},
			data.ImportStep(),
		},
	})
}

func TestAccApplicationOAuth2Permission_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_oauth2_permission", "test")
	id := uuid.New().String()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckOAuth2PermissionCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccApplicationOAuth2Permission_complete(data.RandomInteger, id),
				Check: resource.ComposeTestCheckFunc(
					testCheckOAuth2PermissionExists(data.ResourceName),
				),
			},
			data.ImportStep(),
			{
				Config: testAccApplicationOAuth2Permission_update(data.RandomInteger, id),
				Check: resource.ComposeTestCheckFunc(
					testCheckOAuth2PermissionExists(data.ResourceName),
				),
			},
			data.ImportStep(),
			{
				Config: testAccApplicationOAuth2Permission_complete(data.RandomInteger, id),
				Check: resource.ComposeTestCheckFunc(
					testCheckOAuth2PermissionExists(data.ResourceName),
				),
			},
			data.ImportStep(),
		},
	})
}

func TestAccApplicationOAuth2Permission_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_oauth2_permission", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckOAuth2PermissionCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccApplicationOAuth2Permission_basic(data.RandomInteger),
				Check: resource.ComposeTestCheckFunc(
					testCheckOAuth2PermissionExists(data.ResourceName),
				),
			},
			data.RequiresImportErrorStep(testAccApplicationOAuth2Permission_requiresImport(data.RandomInteger)),
		},
	})
}

func testAccApplicationOAuth2Permission_template(ri int) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name = "acctestApp-%d"
}
`, ri)
}

func testAccApplicationOAuth2Permission_basic(ri int) string {
	return fmt.Sprintf(`
%s

resource "azuread_application_oauth2_permission" "test" {
  application_object_id      = azuread_application.test.id
  admin_consent_description  = "Administer the application"
  admin_consent_display_name = "Administer"
  is_enabled                 = true
  type                       = "Admin"
  user_consent_description   = "Administer the application"
  user_consent_display_name  = "Administer"
  value                      = "administer"
}
`, testAccApplicationOAuth2Permission_template(ri))
}

func testAccApplicationOAuth2Permission_requiresImport(ri int) string {
	return fmt.Sprintf(`
%s

resource "azuread_application_oauth2_permission" "import" {
  application_object_id      = azuread_application_oauth2_permission.test.application_object_id
  admin_consent_description  = azuread_application_oauth2_permission.test.admin_consent_description
  admin_consent_display_name = azuread_application_oauth2_permission.test.admin_consent_display_name
  is_enabled                 = azuread_application_oauth2_permission.test.is_enabled
  permission_id              = azuread_application_oauth2_permission.test.permission_id
  type                       = azuread_application_oauth2_permission.test.type
  value                      = azuread_application_oauth2_permission.test.value
  user_consent_description   = azuread_application_oauth2_permission.test.user_consent_description
  user_consent_display_name  = azuread_application_oauth2_permission.test.user_consent_display_name
}
`, testAccApplicationOAuth2Permission_basic(ri))
}

func testAccApplicationOAuth2Permission_complete(ri int, id string) string {
	return fmt.Sprintf(`
%s

resource "azuread_application_oauth2_permission" "test" {
  application_object_id      = azuread_application.test.id
  admin_consent_description  = "Administer the application"
  admin_consent_display_name = "Administer"
  is_enabled                 = true
  permission_id              = "%s"
  type                       = "Admin"
  user_consent_description   = "Administer the application"
  user_consent_display_name  = "Administer"
  value                      = "administer"
}
`, testAccApplicationOAuth2Permission_template(ri), id)
}

func testAccApplicationOAuth2Permission_update(ri int, id string) string {
	return fmt.Sprintf(`
%s

resource "azuread_application_oauth2_permission" "test" {
  application_object_id      = azuread_application.test.id
  admin_consent_description  = "Administrators can administrate all the things"
  admin_consent_display_name = "Administrate"
  is_enabled                 = true
  permission_id              = "%s"
  type                       = "User"
  user_consent_description   = "Administrators can administrate all the things"
  user_consent_display_name  = "Administrate"
  value                      = "administrate"
}
`, testAccApplicationOAuth2Permission_template(ri), id)
}
