package applications_test

import (
	"context"
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
)

type ApplicationResource struct{}

func TestAccApplication_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application", "test")
	r := ApplicationResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
				check.That(data.ResourceName).Key("object_id").Exists(),
				check.That(data.ResourceName).Key("name").HasValue(fmt.Sprintf("acctest-APP-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctest-APP-%d", data.RandomInteger)),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApplication_basicDeprecated(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application", "test")
	r := ApplicationResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basicDeprecated(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
				check.That(data.ResourceName).Key("object_id").Exists(),
				check.That(data.ResourceName).Key("name").HasValue(fmt.Sprintf("acctest-APP-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctest-APP-%d", data.RandomInteger)),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApplication_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application", "test")
	r := ApplicationResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.complete(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
				check.That(data.ResourceName).Key("object_id").Exists(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApplication_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application", "test")
	r := ApplicationResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
				check.That(data.ResourceName).Key("object_id").Exists(),
			),
		},
		data.ImportStep(),
		{
			Config: r.complete(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
				check.That(data.ResourceName).Key("object_id").Exists(),
			),
		},
		data.ImportStep(),
		{
			Config: r.basicEmpty(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
				check.That(data.ResourceName).Key("object_id").Exists(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApplication_publicClient(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application", "test")
	r := ApplicationResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.publicClient(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
				check.That(data.ResourceName).Key("object_id").Exists(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApplication_appRoles(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application", "test")
	r := ApplicationResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.appRoles(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("app_role.#").HasValue("1"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApplication_appRolesNoValue(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application", "test")
	r := ApplicationResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.appRolesNoValue(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("app_role.#").HasValue("1"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApplication_appRolesUpdate(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application", "test")
	r := ApplicationResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("app_role.#").HasValue("0"),
			),
		},
		data.ImportStep(),
		{
			Config: r.appRoles(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("app_role.#").HasValue("1"),
			),
		},
		data.ImportStep(),
		{
			Config: r.appRolesUpdate(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("app_role.#").HasValue("2"),
			),
		},
		data.ImportStep(),
		{
			Config: r.appRolesEmpty(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("app_role.#").HasValue("0"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApplication_groupMembershipClaimsUpdate(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application", "test")
	r := ApplicationResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.withGroupMembershipClaimsDirectoryRole(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("group_membership_claims").HasValue("DirectoryRole"),
			),
		},
		data.ImportStep(),
		{
			Config: r.withGroupMembershipClaimsSecurityGroup(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("group_membership_claims").HasValue("SecurityGroup"),
			),
		},
		data.ImportStep(),
		{
			Config: r.withGroupMembershipClaimsApplicationGroup(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("group_membership_claims").HasValue("ApplicationGroup"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApplication_native(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application", "test")
	r := ApplicationResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.native(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("homepage").HasValue(""),
				check.That(data.ResourceName).Key("identifier_uris.#").HasValue("0"),
				check.That(data.ResourceName).Key("type").HasValue("native"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApplication_nativeReplyUrls(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application", "test")
	r := ApplicationResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.nativeReplyUrls(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("homepage").HasValue(""),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApplication_nativeUpdate(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application", "test")
	r := ApplicationResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("identifier_uris.#").HasValue("0"),
				check.That(data.ResourceName).Key("type").HasValue("webapp/api"),
			),
		},
		data.ImportStep(),
		{
			Config: r.native(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("identifier_uris.#").HasValue("0"),
				check.That(data.ResourceName).Key("type").HasValue("native"),
			),
		},
		data.ImportStep(),
		{
			Config: r.complete(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("identifier_uris.#").HasValue("1"),
				check.That(data.ResourceName).Key("type").HasValue("webapp/api"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApplication_nativeAppDoesNotAllowIdentifierUris(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application", "test")
	r := ApplicationResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config:      r.nativeAppDoesNotAllowIdentifierUris(data),
			ExpectError: regexp.MustCompile("not required for a native application"),
		},
	})
}

func TestAccApplication_oauth2PermissionsUpdate(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application", "test")
	r := ApplicationResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("oauth2_permissions.#").HasValue("1"),
			),
		},
		data.ImportStep(),
		{
			Config: r.oauth2Permissions(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("oauth2_permissions.#").HasValue("2"),
			),
		},
		data.ImportStep(),
		{
			Config: r.oauth2PermissionsEmpty(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("oauth2_permissions.#").HasValue("0"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApplication_preventDuplicateNamesPass(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application", "test")
	r := ApplicationResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.preventDuplicateNamesPass(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("prevent_duplicate_names"),
	})
}

func TestAccApplication_preventDuplicateNamesFail(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application", "test")
	r := ApplicationResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		data.RequiresImportErrorStep(r.preventDuplicateNamesFail(data)),
	})
}

func TestAccApplication_preventDuplicateNamesPassDeprecated(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application", "test")
	r := ApplicationResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.preventDuplicateNamesPassDeprecated(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("prevent_duplicate_names"),
	})
}

func TestAccApplication_preventDuplicateNamesFailDeprecated(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application", "test")
	r := ApplicationResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		data.RequiresImportErrorStep(r.preventDuplicateNamesFailDeprecated(data)),
	})
}

func TestAccApplication_duplicateAppRolesOauth2PermissionsValues(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application", "test")
	r := ApplicationResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config:      r.duplicateAppRolesOauth2PermissionsValues(data),
			ExpectError: regexp.MustCompile("validation failed: duplicate value found:"),
		},
	})
}

func TestAccApplication_ownersUpdate(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application", "test")
	r := ApplicationResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("owners.#").HasValue("1"),
			),
		},
		data.ImportStep(),
		{
			Config: r.removeOwners(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("owners.#").HasValue("0"),
			),
		},
		data.ImportStep(),
		{
			Config: r.singleOwner(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("owners.#").HasValue("1"),
			),
		},
		data.ImportStep(),
		{
			Config: r.threeOwners(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("owners.#").HasValue("3"),
			),
		},
		data.ImportStep(),
		{
			Config: r.removeOwners(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("owners.#").HasValue("0"),
			),
		},
		data.ImportStep(),
	})
}

func (r ApplicationResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	resp, err := clients.Applications.AadClient.Get(ctx, state.ID)

	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			return nil, fmt.Errorf("Application with object ID %q does not exist", state.ID)
		}
		return nil, fmt.Errorf("failed to retrieve Application with object ID %q: %+v", state.ID, err)
	}
	id := resp.ObjectID

	return utils.Bool(id != nil && *id == state.ID), nil
}

func (ApplicationResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  display_name = "acctest-APP-%[1]d"
}
`, data.RandomInteger)
}

func (ApplicationResource) basicDeprecated(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name = "acctest-APP-%[1]d"
}
`, data.RandomInteger)
}

func (ApplicationResource) basicEmpty(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  display_name            = "acctest-APP-%[1]s"
  identifier_uris         = []
  oauth2_permissions      = []
  reply_urls              = []
  owners                  = []
  group_membership_claims = "None"
}
`, data.RandomString)
}

func (ApplicationResource) publicClient(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  display_name  = "acctest-APP-%[1]d"
  type          = "native"
  public_client = true
}
`, data.RandomInteger)
}

func (ApplicationResource) withGroupMembershipClaimsDirectoryRole(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  display_name            = "acctest-APP-%[1]d"
  group_membership_claims = "DirectoryRole"
}
`, data.RandomInteger)
}

func (ApplicationResource) withGroupMembershipClaimsSecurityGroup(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  display_name            = "acctest-APP-%[1]d"
  group_membership_claims = "SecurityGroup"
}
`, data.RandomInteger)
}

func (ApplicationResource) withGroupMembershipClaimsApplicationGroup(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  display_name            = "acctest-APP-%[1]d"
  group_membership_claims = "ApplicationGroup"
}
`, data.RandomInteger)
}

func (ApplicationResource) complete(data acceptance.TestData) string {
	return fmt.Sprintf(`
data "azuread_domains" "test" {
  only_initial = true
}

resource "azuread_user" "test" {
  user_principal_name = "acctestUser.%[1]d@${data.azuread_domains.test.domains.0.domain_name}"
  display_name        = "acctestUser-%[1]d"
  password            = "%[2]s"
}

resource "azuread_application" "test" {
  display_name               = "acctest-APP-%[1]d"
  homepage                   = "https://homepage-%[1]d"
  identifier_uris            = ["api://hashicorptestapp-%[1]d"]
  reply_urls                 = ["https://unittest.hashicorptest.com"]
  logout_url                 = "https://log.me.out"
  available_to_other_tenants = true
  group_membership_claims    = "All"
  oauth2_allow_implicit_flow = true
  type                       = "webapp/api"

  required_resource_access {
    resource_app_id = "00000003-0000-0000-c000-000000000000"

    resource_access {
      id   = "7ab1d382-f21e-4acd-a863-ba3e13f7da61"
      type = "Role"
    }

    resource_access {
      id   = "e1fe6dd8-ba31-4d61-89e7-88639da4683d"
      type = "Scope"
    }

    resource_access {
      id   = "06da0dbc-49e2-44d2-8312-53f166ab848a"
      type = "Scope"
    }
  }

  required_resource_access {
    resource_app_id = "00000002-0000-0000-c000-000000000000"

    resource_access {
      id   = "311a71cc-e848-46a1-bdf8-97ff7156d8e6"
      type = "Scope"
    }
  }

  oauth2_permissions {
    admin_consent_description  = "Administer the application"
    admin_consent_display_name = "Administer"
    is_enabled                 = true
    type                       = "Admin"
    value                      = "administer"
  }

  oauth2_permissions {
    admin_consent_description  = "Allow the application to access acctest-APP-%[1]d on behalf of the signed-in user."
    admin_consent_display_name = "Access acctest-APP-%[1]d"
    is_enabled                 = true
    type                       = "User"
    user_consent_description   = "Allow the application to access acctest-APP-%[1]d on your behalf."
    user_consent_display_name  = "Access acctest-APP-%[1]d"
    value                      = "user_impersonation"
  }

  optional_claims {
    access_token {
      name = "myclaim"
    }

    access_token {
      name = "otherclaim"
    }

    id_token {
      name                  = "userclaim"
      source                = "user"
      essential             = true
      additional_properties = ["emit_as_roles"]
    }
  }

  owners = [azuread_user.test.object_id]
}
`, data.RandomInteger, data.RandomPassword)
}

func (ApplicationResource) appRoles(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  display_name = "acctest-APP-%[1]d"

  app_role {
    allowed_member_types = [
      "User",
      "Application",
    ]

    description  = "Admins can manage roles and perform all task actions"
    display_name = "Admin"
    is_enabled   = true
    value        = "Admin"
  }
}
`, data.RandomInteger)
}

func (ApplicationResource) appRolesNoValue(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  display_name = "acctest-APP-%[1]d"

  app_role {
    allowed_member_types = ["User"]
    description          = "Admins can manage roles and perform all task actions"
    display_name         = "Admin"
    is_enabled           = true
  }
}
`, data.RandomInteger)
}

func (ApplicationResource) appRolesUpdate(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  display_name = "acctestApp-%[1]d"

  app_role {
    allowed_member_types = ["User"]
    description          = "Admins can manage roles and perform all task actions"
    display_name         = "Admin"
    is_enabled           = true
    value                = ""
  }

  app_role {
    allowed_member_types = ["User"]
    description          = "ReadOnly roles have limited query access"
    display_name         = "ReadOnly"
    is_enabled           = true
    value                = "User"
  }
}
`, data.RandomInteger)
}

func (ApplicationResource) appRolesEmpty(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  display_name = "acctestApp-%[1]d"
  app_role     = []
}
`, data.RandomInteger)
}

func (ApplicationResource) oauth2Permissions(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  display_name = "acctest-APP-%[1]d"

  oauth2_permissions {
    admin_consent_description  = "Allow the application to access acctest-APP-%[1]d on behalf of the signed-in user."
    admin_consent_display_name = "Access acctest-APP-%[1]d"
    is_enabled                 = true
    type                       = "User"
    user_consent_description   = "Allow the application to access acctest-APP-%[1]d on your behalf."
    user_consent_display_name  = "Access acctest-APP-%[1]d"
    value                      = "user_impersonation"
  }

  oauth2_permissions {
    admin_consent_description  = "Administer the application"
    admin_consent_display_name = "Administer"
    is_enabled                 = true
    type                       = "Admin"
    value                      = "administer"
  }
}
`, data.RandomInteger)
}

func (ApplicationResource) oauth2PermissionsEmpty(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  display_name       = "acctest-APP-%[1]d"
  oauth2_permissions = []
}
`, data.RandomInteger)
}

func (ApplicationResource) native(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  display_name = "acctest-APP-%[1]d"
  type         = "native"
}
`, data.RandomInteger)
}

func (ApplicationResource) nativeReplyUrls(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  display_name = "acctest-APP-%[1]d"
  type         = "native"
  reply_urls   = ["urn:ietf:wg:oauth:2.0:oob"]
}
`, data.RandomInteger)
}

func (ApplicationResource) nativeAppDoesNotAllowIdentifierUris(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  display_name    = "acctest-APP-%[1]d"
  identifier_uris = ["http://%[1]d.hashicorptest.com"]
  type            = "native"
}
`, data.RandomInteger)
}

func (ApplicationResource) preventDuplicateNamesPass(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  display_name            = "acctest-APP-%[1]d"
  prevent_duplicate_names = true
}
`, data.RandomInteger)
}

func (r ApplicationResource) preventDuplicateNamesFail(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_application" "duplicate" {
  display_name            = azuread_application.test.name
  prevent_duplicate_names = true
}
`, r.basic(data))
}

func (ApplicationResource) preventDuplicateNamesPassDeprecated(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name                    = "acctest-APP-%[1]d"
  prevent_duplicate_names = true
}
`, data.RandomInteger)
}

func (r ApplicationResource) preventDuplicateNamesFailDeprecated(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_application" "duplicate" {
  name                    = azuread_application.test.name
  prevent_duplicate_names = true
}
`, r.basic(data))
}

func (ApplicationResource) duplicateAppRolesOauth2PermissionsValues(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  display_name = "acctest-APP-%[1]d"

  app_role {
    allowed_member_types = ["User"]
    description          = "Admins can manage roles and perform all task actions"
    display_name         = "Admin"
    is_enabled           = true
    value                = "administer"
  }

  oauth2_permissions {
    admin_consent_description  = "Administer the application"
    admin_consent_display_name = "Administer"
    is_enabled                 = true
    type                       = "Admin"
    value                      = "administer"
  }
}
`, data.RandomInteger)
}

func (ApplicationResource) templateThreeUsers(data acceptance.TestData) string {
	return fmt.Sprintf(`
data "azuread_domains" "test" {
  only_initial = true
}

resource "azuread_user" "testA" {
  user_principal_name = "acctestUser.%[1]d.A@${data.azuread_domains.test.domains.0.domain_name}"
  display_name        = "acctestUser-%[1]d-A"
  password            = "%[2]s"
}

resource "azuread_user" "testB" {
  user_principal_name = "acctestUser.%[1]d.B@${data.azuread_domains.test.domains.0.domain_name}"
  display_name        = "acctestUser-%[1]d-B"
  mail_nickname       = "acctestUser-%[1]d-B"
  password            = "%[2]s"
}

resource "azuread_user" "testC" {
  user_principal_name = "acctestUser.%[1]d.C@${data.azuread_domains.test.domains.0.domain_name}"
  display_name        = "acctestUser-%[1]d-C"
  password            = "%[2]s"
}
`, data.RandomInteger, data.RandomPassword)
}

func (r ApplicationResource) singleOwner(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_application" "test" {
  display_name = "acctest-APP-%[2]d"
  owners = [
    azuread_user.testA.object_id,
  ]
}
`, r.templateThreeUsers(data), data.RandomInteger)
}

func (r ApplicationResource) threeOwners(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_application" "test" {
  display_name = "acctest-APP-%[2]d"
  owners = [
    azuread_user.testA.object_id,
    azuread_user.testB.object_id,
    azuread_user.testC.object_id,
  ]
}
`, r.templateThreeUsers(data), data.RandomInteger)
}

func (r ApplicationResource) removeOwners(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_application" "test" {
  display_name = "acctest-APP-%[2]d"
  owners       = []
}
`, r.templateThreeUsers(data), data.RandomInteger)
}
