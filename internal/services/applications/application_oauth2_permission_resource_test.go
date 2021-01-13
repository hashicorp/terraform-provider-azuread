package applications_test

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
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/aadgraph"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/msgraph"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/applications/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
)

type ApplicationOAuth2PermissionResource struct{}

func TestAccApplicationOAuth2Permission_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_oauth2_permission", "test")
	r := ApplicationOAuth2PermissionResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("permission_id").Exists(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApplicationOAuth2Permission_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_oauth2_permission", "test")
	r := ApplicationOAuth2PermissionResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.complete(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("permission_id").Exists(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApplicationOAuth2Permission_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_oauth2_permission", "test")
	r := ApplicationOAuth2PermissionResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.complete(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("permission_id").Exists(),
			),
		},
		data.ImportStep(),
		{
			Config: r.update(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("permission_id").Exists(),
			),
		},
		data.ImportStep(),
		{
			Config: r.complete(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("permission_id").Exists(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApplicationOAuth2Permission_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_oauth2_permission", "test")
	r := ApplicationOAuth2PermissionResource{}

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

func (r ApplicationOAuth2PermissionResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	id, err := parse.OAuth2PermissionID(state.ID)
	if err != nil {
		return nil, fmt.Errorf("parsing OAuth2 Permission ID: %v", err)
	}

	switch clients.EnableMsGraphBeta {
	case true:
		app, status, err := clients.Applications.MsClient.Get(ctx, id.ObjectId)
		if err != nil {
			if status == http.StatusNotFound {
				return nil, fmt.Errorf("Application with object ID %q does not exist", id.ObjectId)
			}
			return nil, fmt.Errorf("failed to retrieve Application with object ID %q: %+v", id.ObjectId, err)
		}

		role, err := msgraph.OAuth2PermissionFindById(app, id.PermissionId)
		if err != nil {
			return nil, fmt.Errorf("failed to identity OAuth2 Permission: %s", err)
		} else if role != nil {
			return utils.Bool(true), nil
		}

	case false:
		resp, err := clients.Applications.AadClient.Get(ctx, id.ObjectId)
		if err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
				return nil, fmt.Errorf("Application with object ID %q does not exist", id.ObjectId)
			}
			return nil, fmt.Errorf("failed to retrieve Application with object ID %q: %+v", id.ObjectId, err)
		}

		scope, err := aadgraph.OAuth2PermissionFindById(resp, id.PermissionId)
		if err != nil {
			return nil, fmt.Errorf("failed to identity OAuth2 Permission: %s", err)
		} else if scope != nil {
			return utils.Bool(true), nil
		}
	}

	return nil, fmt.Errorf("OAuth2 Permission %q was not found in Application %q", id.PermissionId, id.ObjectId)
}

func (ApplicationOAuth2PermissionResource) template(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name = "acctestApp-%[1]d"
}
`, data.RandomInteger)
}

func (r ApplicationOAuth2PermissionResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

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
`, r.template(data))
}

func (r ApplicationOAuth2PermissionResource) complete(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_application_oauth2_permission" "test" {
  application_object_id      = azuread_application.test.id
  admin_consent_description  = "Administer the application"
  admin_consent_display_name = "Administer"
  is_enabled                 = true
  permission_id              = "%[2]s"
  type                       = "Admin"
  user_consent_description   = "Administer the application"
  user_consent_display_name  = "Administer"
  value                      = "administer"
}
`, r.template(data), data.RandomID)
}

func (r ApplicationOAuth2PermissionResource) update(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_application_oauth2_permission" "test" {
  application_object_id      = azuread_application.test.id
  admin_consent_description  = "Administrators can administrate all the things"
  admin_consent_display_name = "Administrate"
  is_enabled                 = true
  permission_id              = "%[2]s"
  type                       = "User"
  user_consent_description   = "Administrators can administrate all the things"
  user_consent_display_name  = "Administrate"
  value                      = "administrate"
}
`, r.template(data), data.RandomID)
}

func (r ApplicationOAuth2PermissionResource) requiresImport(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

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
`, r.basic(data))
}
