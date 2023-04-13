package applications_test

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"testing"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
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
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctest-APP-%d", data.RandomInteger)),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApplication_basicFromTemplate(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application", "test")
	r := ApplicationResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basicFromTemplate(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
				check.That(data.ResourceName).Key("object_id").Exists(),
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctest-APP-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("template_id").HasValue(testApplicationTemplateId),
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

func TestAccApplication_completeFromTemplate(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application", "test")
	r := ApplicationResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.completeFromTemplate(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
				check.That(data.ResourceName).Key("object_id").Exists(),
				check.That(data.ResourceName).Key("template_id").HasValue(testApplicationTemplateId),
				check.That(data.ResourceName).Key("app_role.#").HasValue("1"),
				check.That(data.ResourceName).Key("app_role.0.id").HasValue(testApplicationTemplateAppRoleId),
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
			Config: r.basic(data),
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
	roleIDs := []string{
		data.UUID(),
		data.UUID(),
	}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("app_role.#").HasValue("0"),
				check.That(data.ResourceName).Key("app_role_ids.%").HasValue("0"),
			),
		},
		data.ImportStep(),
		{
			Config: r.appRoleNoValue(data, roleIDs),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("app_role.#").HasValue("1"),
				check.That(data.ResourceName).Key("app_role_ids.%").HasValue("0"),
			),
		},
		data.ImportStep(),
		{
			Config: r.appRole(data, roleIDs),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("app_role.#").HasValue("1"),
				check.That(data.ResourceName).Key("app_role_ids.%").HasValue("1"),
			),
		},
		data.ImportStep(),
		{
			Config: r.appRolesUpdate(data, roleIDs),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("app_role.#").HasValue("2"),
				check.That(data.ResourceName).Key("app_role_ids.%").HasValue("2"),
			),
		},
		data.ImportStep(),
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("app_role.#").HasValue("0"),
				check.That(data.ResourceName).Key("app_role_ids.%").HasValue("0"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApplication_duplicateAppRolesOauth2PermissionsIdsUnknown(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application", "test")
	r := ApplicationResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.duplicateAppRolesOauth2PermissionsIdsUnknown(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("app_role.#").HasValue("1"),
				check.That(data.ResourceName).Key("app_role_ids.%").HasValue("1"),
				check.That(data.ResourceName).Key("api.0.oauth2_permission_scope.#").HasValue("1"),
				check.That(data.ResourceName).Key("oauth2_permission_scope_ids.%").HasValue("1"),
			),
		},
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

func TestAccApplication_duplicateAppRolesOauth2PermissionsMatchingIdAndValueWithCommonMetadata(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application", "test")
	r := ApplicationResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.duplicateAppRolesOauth2PermissionsMatchingIdAndValueWithCommonMetadata(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("app_role.#").HasValue("1"),
				check.That(data.ResourceName).Key("app_role_ids.%").HasValue("1"),
				check.That(data.ResourceName).Key("api.0.oauth2_permission_scope.#").HasValue("1"),
				check.That(data.ResourceName).Key("oauth2_permission_scope_ids.%").HasValue("1"),
			),
		},
	})
}

func TestAccApplication_duplicateAppRolesOauth2PermissionsMatchingIdAndValueWithMismatchingMetadata(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application", "test")
	r := ApplicationResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config:      r.duplicateAppRolesOauth2PermissionsMatchingIdAndValueWithMismatchingMetadata(data),
			ExpectError: regexp.MustCompile("validation failed: The following values must match for the"),
		},
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
			Config: r.withGroupMembershipClaims(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApplication_oauth2PermissionScopes(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application", "test")
	r := ApplicationResource{}
	scopeIDs := []string{
		data.UUID(),
		data.UUID(),
		data.UUID(),
	}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("api.0.oauth2_permission_scope.#").HasValue("0"),
				check.That(data.ResourceName).Key("oauth2_permission_scope_ids.%").HasValue("0"),
			),
		},
		data.ImportStep(),
		{
			Config: r.oauth2PermissionScopes(data, scopeIDs),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("api.0.oauth2_permission_scope.#").HasValue("2"),
				check.That(data.ResourceName).Key("oauth2_permission_scope_ids.%").HasValue("2"),
			),
		},
		data.ImportStep(),
		{
			Config: r.oauth2PermissionScopesUpdate(data, scopeIDs),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("api.0.oauth2_permission_scope.#").HasValue("3"),
				check.That(data.ResourceName).Key("oauth2_permission_scope_ids.%").HasValue("3"),
			),
		},
		data.ImportStep(),
		{
			Config: r.oauth2PermissionScopes(data, scopeIDs),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("api.0.oauth2_permission_scope.#").HasValue("2"),
				check.That(data.ResourceName).Key("oauth2_permission_scope_ids.%").HasValue("2"),
			),
		},
		data.ImportStep(),
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("api.0.oauth2_permission_scope.#").HasValue("0"),
				check.That(data.ResourceName).Key("oauth2_permission_scope_ids.%").HasValue("0"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApplication_owners(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application", "test")
	r := ApplicationResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
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
			Config: r.noOwners(data),
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
			Config: r.noOwners(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("owners.#").HasValue("0"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApplication_createWithNoOwners(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application", "test")
	r := ApplicationResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.noOwners(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("owners.#").HasValue("0"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApplication_manyOwners(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application", "test")
	r := ApplicationResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.manyOwners(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("owners.#").HasValue("45"),
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

func TestAccApplication_related(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application", "test")
	r := ApplicationResource{}
	uuids := []string{
		data.UUID(),
		data.UUID(),
		data.UUID(),
		data.UUID(),
	}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.related(data, uuids),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.relatedUpdate(data, uuids),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApplication_featureTags(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application", "test")
	r := ApplicationResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.featureTags(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApplication_featureTagsUpdate(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application", "test")
	r := ApplicationResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.noFeatureTags(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.featureTags(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.featureTags(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.tags(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.featureTags(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.noFeatureTags(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.featureTags(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApplication_logo(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application", "test")
	r := ApplicationResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.logo(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("logo"),
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.logo(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("logo"),
	})
}

func (r ApplicationResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.Applications.ApplicationsClient
	client.BaseClient.DisableRetries = true
	app, status, err := client.Get(ctx, state.ID, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return nil, fmt.Errorf("Application with object ID %q does not exist", state.ID)
		}
		return nil, fmt.Errorf("failed to retrieve Application with object ID %q: %+v", state.ID, err)
	}
	return utils.Bool(app.ID() != nil && *app.ID() == state.ID), nil
}

func (ApplicationResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application" "test" {
  display_name = "acctest-APP-%[1]d"
}
`, data.RandomInteger)
}

func (ApplicationResource) basicFromTemplate(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

data "azuread_client_config" "test" {}

resource "azuread_application" "test" {
  display_name = "acctest-APP-%[1]d"
  owners       = [data.azuread_client_config.test.object_id]
  template_id  = "%[2]s"
}
`, data.RandomInteger, testApplicationTemplateId)
}

func (ApplicationResource) withGroupMembershipClaims(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application" "test" {
  display_name            = "acctest-APP-%[1]d"
  group_membership_claims = ["DirectoryRole", "SecurityGroup", "ApplicationGroup"]
}
`, data.RandomInteger)
}

func (ApplicationResource) complete(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

data "azuread_domains" "test" {
  only_initial = true
}

resource "azuread_user" "test" {
  user_principal_name = "acctestUser.%[1]d@${data.azuread_domains.test.domains.0.domain_name}"
  display_name        = "acctestUser-%[1]d"
  password            = "%[2]s"
}

resource "azuread_application" "known1" {
  display_name = "acctest-APP-known1-%[1]d"
}

resource "azuread_application" "known2" {
  display_name = "acctest-APP-known2-%[1]d"
}

resource "azuread_application" "test" {
  display_name            = "acctest-APP-complete-%[1]d"
  group_membership_claims = ["All"]
  sign_in_audience        = "AzureADandPersonalMicrosoftAccount"

  identifier_uris = [
    "api://hashicorptestapp-%[1]d",
    "api://acctest-APP-complete-%[1]d",
  ]

  device_only_auth_enabled       = true
  fallback_public_client_enabled = true
  oauth2_post_response_required  = true

  description                  = "Acceptance testing application"
  notes                        = "Testing application"
  service_management_reference = "app-for-testing"

  marketing_url         = "https://hashitown-%[1]d.com/"
  privacy_statement_url = "https://hashitown-%[1]d.com/privacy"
  support_url           = "https://support.hashitown-%[1]d.com/"
  terms_of_service_url  = "https://hashitown-%[1]d.com/terms"

  api {
    mapped_claims_enabled          = true
    requested_access_token_version = 2

    known_client_applications = [
      azuread_application.known1.application_id,
      azuread_application.known2.application_id,
    ]

    oauth2_permission_scope {
      admin_consent_description  = "Administer the application"
      admin_consent_display_name = "Administer"
      enabled                    = true
      id                         = "%[3]s"
      type                       = "Admin"
      value                      = "administer"
    }

    oauth2_permission_scope {
      admin_consent_description  = "Allow the application to access acctest-APP-%[1]d on behalf of the signed-in user."
      admin_consent_display_name = "Access acctest-APP-%[1]d"
      enabled                    = true
      id                         = "%[4]s"
      type                       = "User"
      user_consent_description   = "Allow the application to access acctest-APP-%[1]d on your behalf."
      user_consent_display_name  = "Access acctest-APP-%[1]d"
      value                      = "user_impersonation"
    }
  }

  app_role {
    allowed_member_types = ["User"]
    description          = "Admins can manage roles and perform all task actions"
    display_name         = "Admin"
    enabled              = true
    id                   = "%[5]s"
    value                = "admin"
  }

  app_role {
    allowed_member_types = ["User"]
    description          = "ReadOnly roles have limited query access"
    display_name         = "ReadOnly"
    enabled              = true
    id                   = "%[6]s"
    value                = "user"
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

    saml2_token {
      name = "samlexample"
    }
  }

  public_client {
    redirect_uris = [
      "myapp://auth",
      "sample.mobile.app.bundie.id://auth",
      "https://login.microsoftonline.com/common/oauth2/nativeclient",
      "https://login.live.com/oauth20_desktop.srf",
      "ms-appx-web://Microsoft.AAD.BrokerPlugin/00000000-1111-1111-1111-222222222222",
      "urn:ietf:wg:oauth:2.0:foo",
    ]
  }

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

  single_page_application {
    redirect_uris = [
      "https://beta.hashitown-%[1]d.com/",
    ]
  }

  tags = [
    "HideApp",
    "WindowsAzureActiveDirectoryCustomSingleSignOnApplication",
    "WindowsAzureActiveDirectoryIntegratedApp",
    "WindowsAzureActiveDirectoryGalleryApplicationNonPrimaryV1",
  ]

  web {
    homepage_url = "https://app.hashitown-%[1]d.com/"
    logout_url   = "https://app.hashitown-%[1]d.com/logout"

    redirect_uris = [
      "https://app.hashitown-%[1]d.com/",
      "https://classic.hashitown-%[1]d.com/",
      "urn:ietf:wg:oauth:2.0:oob",
    ]

    implicit_grant {
      access_token_issuance_enabled = true
      id_token_issuance_enabled     = true
    }
  }

  owners = [azuread_user.test.object_id]
}
`, data.RandomInteger, data.RandomPassword, data.UUID(), data.UUID(), data.UUID(), data.UUID())
}

func (ApplicationResource) completeFromTemplate(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

data "azuread_domains" "test" {
  only_initial = true
}

resource "azuread_user" "test" {
  user_principal_name = "acctestUser.%[1]d@${data.azuread_domains.test.domains.0.domain_name}"
  display_name        = "acctestUser-%[1]d"
  password            = "%[2]s"
}

resource "azuread_application" "known1" {
  display_name = "acctest-APP-known1-%[1]d"
}

resource "azuread_application" "known2" {
  display_name = "acctest-APP-known2-%[1]d"
}

resource "azuread_application" "test" {
  display_name            = "acctest-APP-complete-%[1]d"
  template_id             = "%[3]s"
  group_membership_claims = ["All"]
  sign_in_audience        = "AzureADandPersonalMicrosoftAccount"

  identifier_uris = [
    "api://hashicorptestapp-%[1]d",
    "api://acctest-APP-complete-%[1]d",
  ]

  device_only_auth_enabled       = true
  fallback_public_client_enabled = true
  oauth2_post_response_required  = true

  description           = "Acceptance testing application"
  marketing_url         = "https://templatetown-%[1]d.com/"
  privacy_statement_url = "https://templatetown-%[1]d.com/privacy"
  support_url           = "https://support.templatetown-%[1]d.com/"
  terms_of_service_url  = "https://templatetown-%[1]d.com/terms"

  api {
    mapped_claims_enabled          = true
    requested_access_token_version = 2

    known_client_applications = [
      azuread_application.known1.application_id,
      azuread_application.known2.application_id,
    ]

    oauth2_permission_scope {
      admin_consent_description  = "Administer the application"
      admin_consent_display_name = "Administer"
      enabled                    = true
      id                         = "%[5]s"
      type                       = "Admin"
      value                      = "administer"
    }

    oauth2_permission_scope {
      admin_consent_description  = "Allow the application to access acctest-APP-%[1]d on behalf of the signed-in user."
      admin_consent_display_name = "Access acctest-APP-%[1]d"
      enabled                    = true
      id                         = "%[6]s"
      type                       = "User"
      user_consent_description   = "Allow the application to access acctest-APP-%[1]d on your behalf."
      user_consent_display_name  = "Access acctest-APP-%[1]d"
      value                      = "user_impersonation"
    }
  }

  app_role {
    allowed_member_types = ["User"]
    description          = "msiam_access"
    display_name         = "msiam_access"
    enabled              = true
    id                   = "%[4]s"
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

    saml2_token {
      name = "samlexample"
    }
  }

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

  single_page_application {
    redirect_uris = [
      "https://beta.templatetown-%[1]d.com/",
    ]
  }

  tags = [
    "WindowsAzureActiveDirectoryCustomSingleSignOnApplication",
    "WindowsAzureActiveDirectoryIntegratedApp",
    "WindowsAzureActiveDirectoryGalleryApplicationNonPrimaryV1",
  ]

  web {
    homepage_url = "https://app.templatetown-%[1]d.com/"
    logout_url   = "https://app.templatetown-%[1]d.com/logout"

    redirect_uris = [
      "https://app.templatetown-%[1]d.com/",
      "https://classic.templatetown-%[1]d.com/",
    ]

    implicit_grant {
      access_token_issuance_enabled = true
      id_token_issuance_enabled     = true
    }
  }

  owners = [azuread_user.test.object_id]
}
`, data.RandomInteger, data.RandomPassword, testApplicationTemplateId, testApplicationTemplateAppRoleId, data.UUID(), data.UUID())
}

func (ApplicationResource) appRole(data acceptance.TestData, roleIDs []string) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application" "test" {
  display_name = "acctest-APP-%[1]d"

  app_role {
    allowed_member_types = ["User", "Application"]
    description          = "Admins can manage roles and perform all task actions"
    display_name         = "Admin"
    enabled              = true
    id                   = "%[2]s"
    value                = "admin"
  }
}
`, data.RandomInteger, roleIDs[0])
}

func (ApplicationResource) appRoleNoValue(data acceptance.TestData, roleIDs []string) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application" "test" {
  display_name = "acctest-APP-%[1]d"

  app_role {
    allowed_member_types = ["User"]
    description          = "Admins can manage roles and perform all task actions"
    display_name         = "Admin"
    id                   = "%[2]s"
    enabled              = true
  }
}
`, data.RandomInteger, roleIDs[0])
}

func (ApplicationResource) appRolesUpdate(data acceptance.TestData, roleIDs []string) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application" "test" {
  display_name = "acctestApp-%[1]d"

  app_role {
    allowed_member_types = ["User"]
    description          = "Admins can manage roles and perform all task actions"
    display_name         = "Admin"
    enabled              = true
    id                   = "%[2]s"
    value                = "admin"
  }

  app_role {
    allowed_member_types = ["User"]
    description          = "ReadOnly roles have limited query access"
    display_name         = "ReadOnly"
    enabled              = true
    id                   = "%[3]s"
    value                = "user"
  }
}
`, data.RandomInteger, roleIDs[0], roleIDs[1])
}

func (ApplicationResource) oauth2PermissionScopes(data acceptance.TestData, scopeIDs []string) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application" "test" {
  display_name = "acctest-APP-%[1]d"

  api {
    oauth2_permission_scope {
      admin_consent_description  = "Allow the application to access acctest-APP-%[1]d on behalf of the signed-in user."
      admin_consent_display_name = "Access acctest-APP-%[1]d"
      enabled                    = true
      id                         = "%[2]s"
      type                       = "User"
      user_consent_description   = "Allow the application to access acctest-APP-%[1]d on your behalf."
      user_consent_display_name  = "Access acctest-APP-%[1]d"
      value                      = "user_impersonation"
    }

    oauth2_permission_scope {
      admin_consent_description  = "Administer the application"
      admin_consent_display_name = "Administer"
      enabled                    = true
      id                         = "%[3]s"
      type                       = "Admin"
      value                      = "administer"
    }
  }
}
`, data.RandomInteger, scopeIDs[0], scopeIDs[1])
}

func (ApplicationResource) oauth2PermissionScopesUpdate(data acceptance.TestData, scopeIDs []string) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application" "test" {
  display_name = "acctest-APP-%[1]d"

  api {
    oauth2_permission_scope {
      admin_consent_description  = "Allow the application to access on behalf blah... this changed"
      admin_consent_display_name = "Access acctest-APP-%[1]d"
      enabled                    = true
      id                         = "%[2]s"
      type                       = "User"
      user_consent_description   = "Allow the application to access acctest-APP-%[1]d on your behalf."
      user_consent_display_name  = "Access acctest-APP-%[1]d"
      value                      = "user_impersonation"
    }

    oauth2_permission_scope {
      admin_consent_description  = "Administer the application"
      admin_consent_display_name = "Administer"
      enabled                    = true
      id                         = "%[3]s"
      type                       = "Admin"
      value                      = "administer"
    }

    oauth2_permission_scope {
      admin_consent_description  = "Audit the application"
      admin_consent_display_name = "Audit"
      enabled                    = true
      id                         = "%[4]s"
      type                       = "Admin"
      value                      = "audit"
    }
  }
}
`, data.RandomInteger, scopeIDs[0], scopeIDs[1], scopeIDs[2])
}

func (ApplicationResource) preventDuplicateNamesPass(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

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
  display_name            = azuread_application.test.display_name
  prevent_duplicate_names = true
}
`, r.basic(data))
}

func (ApplicationResource) related(data acceptance.TestData, uuids []string) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application" "service" {
  display_name = "acctest-APP-service-%[1]d"

  api {
    oauth2_permission_scope {
      admin_consent_description  = "Allow the application to access acctest-APP-%[1]d on behalf of the signed-in user."
      admin_consent_display_name = "Access acctest-APP-%[1]d"
      enabled                    = true
      id                         = "%[2]s"
      type                       = "User"
      user_consent_description   = "Allow the application to access acctest-APP-%[1]d on your behalf."
      user_consent_display_name  = "Access acctest-APP-%[1]d"
      value                      = "user_impersonation"
    }
  }

  app_role {
    allowed_member_types = ["User"]
    description          = "ReadOnly roles have limited query access"
    display_name         = "ReadOnly"
    enabled              = true
    id                   = "%[3]s"
    value                = "user"
  }
}

resource "azuread_application" "test" {
  display_name = "acctest-APP-related-%[1]d"

  required_resource_access {
    resource_app_id = azuread_application.service.application_id

    resource_access {
      id   = azuread_application.service.app_role_ids["user"]
      type = "Role"
    }

    resource_access {
      id   = azuread_application.service.oauth2_permission_scope_ids["user_impersonation"]
      type = "Scope"
    }
  }
}
`, data.RandomInteger, uuids[0], uuids[1], uuids[2], uuids[3])
}

func (ApplicationResource) relatedUpdate(data acceptance.TestData, uuids []string) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application" "service" {
  display_name = "acctest-APP-service-%[1]d"

  api {
    oauth2_permission_scope {
      admin_consent_description  = "Allow the application to access acctest-APP-%[1]d on behalf of the signed-in user."
      admin_consent_display_name = "Access acctest-APP-%[1]d"
      enabled                    = true
      id                         = "%[2]s"
      type                       = "User"
      user_consent_description   = "Allow the application to access acctest-APP-%[1]d on your behalf."
      user_consent_display_name  = "Access acctest-APP-%[1]d"
      value                      = "user_impersonation"
    }

    oauth2_permission_scope {
      admin_consent_description  = "Administer the application"
      admin_consent_display_name = "Administer"
      enabled                    = true
      id                         = "%[4]s"
      type                       = "Admin"
      value                      = "administer"
    }
  }

  app_role {
    allowed_member_types = ["User"]
    description          = "Admins can manage roles and perform all task actions"
    display_name         = "Admin"
    enabled              = true
    id                   = "%[5]s"
    value                = "admin"
  }

  app_role {
    allowed_member_types = ["User"]
    description          = "ReadOnly roles have limited query access"
    display_name         = "ReadOnly"
    enabled              = true
    id                   = "%[3]s"
    value                = "user"
  }
}

resource "azuread_application" "test" {
  display_name = "acctest-APP-related-%[1]d"

  required_resource_access {
    resource_app_id = azuread_application.service.application_id

    resource_access {
      id   = azuread_application.service.app_role_ids["admin"]
      type = "Role"
    }

    resource_access {
      id   = azuread_application.service.app_role_ids["user"]
      type = "Role"
    }

    resource_access {
      id   = azuread_application.service.oauth2_permission_scope_ids["administer"]
      type = "Scope"
    }

    resource_access {
      id   = azuread_application.service.oauth2_permission_scope_ids["user_impersonation"]
      type = "Scope"
    }
  }
}
`, data.RandomInteger, uuids[0], uuids[1], uuids[2], uuids[3])
}

func (ApplicationResource) duplicateAppRolesOauth2PermissionsIdsUnknown(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application" "test" {
  display_name = "acctest-APP-%[1]d"

  api {
    oauth2_permission_scope {
      admin_consent_description  = "Administer the application"
      admin_consent_display_name = "Administer"
      enabled                    = true
      id                         = "%[2]s"
      type                       = "Admin"
      value                      = "administer"
    }
  }

  app_role {
    allowed_member_types = ["User"]
    description          = "Admins can manage roles and perform all task actions"
    display_name         = "Admin"
    enabled              = true
    id                   = "%[3]s"
    value                = "administrate"
  }
}
`, data.RandomInteger, data.UUID(), data.UUID())
}

func (ApplicationResource) duplicateAppRolesOauth2PermissionsValues(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application" "test" {
  display_name = "acctest-APP-%[1]d"

  api {
    oauth2_permission_scope {
      admin_consent_description  = "Administer the application"
      admin_consent_display_name = "Administer"
      enabled                    = true
      id                         = "%[2]s"
      type                       = "Admin"
      value                      = "administer"
    }
  }

  app_role {
    allowed_member_types = ["User"]
    description          = "Admins can manage roles and perform all task actions"
    display_name         = "Admin"
    enabled              = true
    id                   = "%[3]s"
    value                = "administer"
  }
}
`, data.RandomInteger, data.UUID(), data.UUID())
}

func (ApplicationResource) duplicateAppRolesOauth2PermissionsMatchingIdAndValueWithCommonMetadata(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application" "test" {
  display_name = "acctest-APP-%[1]d"

  api {
    oauth2_permission_scope {
      admin_consent_description  = "Administer the application"
      admin_consent_display_name = "Administer"
      enabled                    = true
      id                         = "%[2]s"
      type                       = "Admin"
      value                      = "administer"
    }
  }

  app_role {
    allowed_member_types = ["User"]
    description          = "Administer the application"
    display_name         = "Administer"
    enabled              = true
    id                   = "%[2]s"
    value                = "administer"
  }
}
`, data.RandomInteger, data.UUID())
}

func (ApplicationResource) duplicateAppRolesOauth2PermissionsMatchingIdAndValueWithMismatchingMetadata(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application" "test" {
  display_name = "acctest-APP-%[1]d"

  api {
    oauth2_permission_scope {
      admin_consent_description  = "This should (see app_role[0].description"
      admin_consent_display_name = "Administer"
      enabled                    = true
      id                         = "%[2]s"
      type                       = "Admin"
      value                      = "administer"
    }
  }

  app_role {
    allowed_member_types = ["User"]
    description          = "Not work"
    display_name         = "Administer"
    enabled              = true
    id                   = "%[2]s"
    value                = "administer"
  }
}
`, data.RandomInteger, data.UUID())
}

func (ApplicationResource) templateThreeUsers(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

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

func (r ApplicationResource) noOwners(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_application" "test" {
  display_name = "acctest-APP-%[2]d"
  owners       = []
}
`, r.templateThreeUsers(data), data.RandomInteger, data.RandomInteger)
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

func (r ApplicationResource) manyOwners(data acceptance.TestData) string {
	return fmt.Sprintf(`
data "azuread_client_config" "test" {}

data "azuread_domains" "test" {
  only_initial = true
}

resource "azuread_application" "owner" {
  count        = 27
  display_name = "acctestApplicationOwner${count.index}-%[1]d"
}

resource "azuread_service_principal" "owner" {
  count          = 27
  application_id = azuread_application.owner[count.index].application_id
}

resource "azuread_user" "owner" {
  count               = 17
  user_principal_name = "acctestApplicationOwner${count.index}-%[1]d@${data.azuread_domains.test.domains.0.domain_name}"
  display_name        = "acctestApplicationOwner${count.index}-%[1]d"
  password            = "Qwer5678!@#"
}

resource "azuread_application" "test" {
  display_name = "acctest-APP-%[1]d"

  owners = flatten([
    data.azuread_client_config.test.object_id,
    azuread_service_principal.owner.*.object_id,
    azuread_user.owner.*.object_id,
  ])
}
`, data.RandomInteger)
}

func (r ApplicationResource) featureTags(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application" "test" {
  display_name = "acctest-APP-%[1]d"

  feature_tags {
    custom_single_sign_on = true
    enterprise            = true
    gallery               = true
    hide                  = true
  }
}
`, data.RandomInteger)
}

func (r ApplicationResource) noFeatureTags(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application" "test" {
  display_name = "acctest-APP-%[1]d"

  feature_tags {
    custom_single_sign_on = false
    enterprise            = false
    gallery               = false
    hide                  = false
  }
}
`, data.RandomInteger)
}

func (r ApplicationResource) tags(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application" "test" {
  display_name = "acctest-APP-%[1]d"

  tags = [
    "WindowsAzureActiveDirectoryCustomSingleSignOnApplication",
    "WindowsAzureActiveDirectoryIntegratedApp",
    "WindowsAzureActiveDirectoryGalleryApplicationNonPrimaryV1",
  ]
}
`, data.RandomInteger)
}

func (r ApplicationResource) logo(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application" "test" {
  display_name = "acctest-APP-%[1]d"
  logo_image   = "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" // thisisfine
}
`, data.RandomInteger)
}
