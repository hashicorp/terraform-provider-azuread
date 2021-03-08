package applications_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/aadgraph"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/applications/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
)

type ApplicationAppRoleResource struct{}

func TestAccApplicationAppRole_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_app_role", "test")
	r := ApplicationAppRoleResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("role_id").Exists(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApplicationAppRole_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_app_role", "test")
	r := ApplicationAppRoleResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.complete(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("role_id").Exists(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApplicationAppRole_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_app_role", "test")
	r := ApplicationAppRoleResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.complete(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("role_id").Exists(),
			),
		},
		data.ImportStep(),
		{
			Config: r.update(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("role_id").Exists(),
			),
		},
		data.ImportStep(),
		{
			Config: r.complete(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("role_id").Exists(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApplicationAppRole_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_app_role", "test")
	r := ApplicationAppRoleResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.RequiresImportErrorStep(r.requiresImport(data)),
	})
}

func (a ApplicationAppRoleResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	id, err := parse.AppRoleID(state.ID)
	if err != nil {
		return nil, fmt.Errorf("parsing App Role ID: %v", err)
	}

	resp, err := clients.Applications.AadClient.Get(ctx, id.ObjectId)
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			return nil, fmt.Errorf("Application with object ID %q does not exist", id.ObjectId)
		}
		return nil, fmt.Errorf("failed to retrieve Application with object ID %q: %+v", id.ObjectId, err)
	}

	role, err := aadgraph.AppRoleFindById(resp, id.RoleId)
	if err != nil {
		return nil, fmt.Errorf("failed to identity App Role: %s", err)
	} else if role != nil {
		return utils.Bool(true), nil
	}

	return nil, fmt.Errorf("App Role %q was not found in Application %q", id.RoleId, id.ObjectId)
}

func (ApplicationAppRoleResource) template(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name = "acctestApp-%[1]d"
}
`, data.RandomInteger)
}

func (r ApplicationAppRoleResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_application_app_role" "test" {
  application_object_id = azuread_application.test.id
  allowed_member_types  = ["User"]
  description           = "Admins can manage roles and perform all task actions"
  display_name          = "Admin"
  is_enabled            = true
}
`, r.template(data))
}

func (r ApplicationAppRoleResource) complete(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_application_app_role" "test" {
  application_object_id = azuread_application.test.id
  allowed_member_types  = ["User"]
  description           = "Admins can manage roles and perform all task actions"
  display_name          = "Admin"
  is_enabled            = true
  role_id               = "%[2]s"
  value                 = "administer"
}
`, r.template(data), data.RandomID)
}

func (r ApplicationAppRoleResource) update(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_application_app_role" "test" {
  application_object_id = azuread_application.test.id
  allowed_member_types  = ["Application", "User"]
  description           = "Administrators can administrate all the things"
  display_name          = "Administrate"
  is_enabled            = true
  role_id               = "%[2]s"
  value                 = "administrate"
}
`, r.template(data), data.RandomID)
}

func (r ApplicationAppRoleResource) requiresImport(data acceptance.TestData) string {
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
`, r.basic(data))
}
