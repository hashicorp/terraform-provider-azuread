package serviceprincipals_test

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/manicminer/hamilton/odata"

	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
)

type ServicePrincipalResource struct{}

const testApplicationTemplateId = "4601ed45-8ff3-4599-8377-b6649007e876" // Marketo

func TestAccServicePrincipal_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_service_principal", "test")
	r := ServicePrincipalResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("use_existing"),
	})
}

func TestAccServicePrincipal_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_service_principal", "test")
	r := ServicePrincipalResource{}
	tenantId := os.Getenv("ARM_TENANT_ID")

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.complete(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("app_role_ids.%").HasValue("2"),
				check.That(data.ResourceName).Key("app_roles.#").HasValue("2"),
				check.That(data.ResourceName).Key("application_tenant_id").HasValue(tenantId),
				check.That(data.ResourceName).Key("homepage_url").HasValue(fmt.Sprintf("https://test-%d.internal", data.RandomInteger)),
				check.That(data.ResourceName).Key("logout_url").HasValue(fmt.Sprintf("https://test-%d.internal/logout", data.RandomInteger)),
				check.That(data.ResourceName).Key("oauth2_permission_scope_ids.%").HasValue("2"),
				check.That(data.ResourceName).Key("oauth2_permission_scopes.#").HasValue("2"),
				check.That(data.ResourceName).Key("service_principal_names.#").HasValue("2"),
				check.That(data.ResourceName).Key("redirect_uris.#").HasValue("2"),
				check.That(data.ResourceName).Key("sign_in_audience").HasValue("AzureADMyOrg"),
				check.That(data.ResourceName).Key("type").HasValue("Application"),
			),
		},
		data.ImportStep("use_existing"),
	})
}

func TestAccServicePrincipal_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_service_principal", "test")
	r := ServicePrincipalResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("app_roles.#").HasValue("0"),
				check.That(data.ResourceName).Key("app_role_ids.%").HasValue("0"),
				check.That(data.ResourceName).Key("oauth2_permission_scope_ids.%").HasValue("0"),
				check.That(data.ResourceName).Key("oauth2_permission_scopes.#").HasValue("0"),
			),
		},
		data.ImportStep("use_existing"),
		{
			Config: r.complete(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("app_roles.#").HasValue("2"),
				check.That(data.ResourceName).Key("app_role_ids.%").HasValue("2"),
				check.That(data.ResourceName).Key("oauth2_permission_scope_ids.%").HasValue("2"),
				check.That(data.ResourceName).Key("oauth2_permission_scopes.#").HasValue("2"),
			),
		},
		data.ImportStep("use_existing"),
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("app_roles.#").HasValue("0"),
				check.That(data.ResourceName).Key("app_role_ids.%").HasValue("0"),
				check.That(data.ResourceName).Key("oauth2_permission_scope_ids.%").HasValue("0"),
				check.That(data.ResourceName).Key("oauth2_permission_scopes.#").HasValue("0"),
			),
		},
		data.ImportStep("use_existing"),
	})
}

func TestAccServicePrincipal_owners(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_service_principal", "test")
	r := ServicePrincipalResource{}

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
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("owners.#").HasValue("0"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApplication_createWithNoOwners(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_service_principal", "test")
	r := ServicePrincipalResource{}

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

func TestAccServicePrincipal_manyOwners(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_service_principal", "test")
	r := ServicePrincipalResource{}

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

func TestAccServicePrincipal_useExisting(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_service_principal", "msgraph")
	r := ServicePrincipalResource{}

	data.ResourceTestIgnoreDangling(t, r, []resource.TestStep{
		{
			Config: r.useExisting(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("app_roles.#").Exists(),
				check.That(data.ResourceName).Key("app_role_ids.%").Exists(),
				check.That(data.ResourceName).Key("oauth2_permission_scope_ids.%").Exists(),
				check.That(data.ResourceName).Key("oauth2_permission_scopes.#").Exists(),
			),
		},
		data.ImportStep("use_existing"),
	})
}

func TestAccServicePrincipal_fromApplicationTemplate(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_service_principal", "test")
	r := ServicePrincipalResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.fromApplicationTemplate(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("use_existing"),
	})
}

func (r ServicePrincipalResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.ServicePrincipals.ServicePrincipalsClient
	client.BaseClient.DisableRetries = true

	servicePrincipal, status, err := client.Get(ctx, state.ID, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return nil, fmt.Errorf("Service Principal with object ID %q does not exist", state.ID)
		}
		return nil, fmt.Errorf("failed to retrieve Service Principal with object ID %q: %+v", state.ID, err)
	}
	return utils.Bool(servicePrincipal.ID != nil && *servicePrincipal.ID == state.ID), nil
}

func (ServicePrincipalResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application" "test" {
  display_name = "acctestServicePrincipal-%[1]d"
}

resource "azuread_service_principal" "test" {
  application_id = azuread_application.test.application_id
}
`, data.RandomInteger)
}

func (ServicePrincipalResource) complete(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application" "test" {
  display_name     = "acctestServicePrincipal-%[1]d"
  sign_in_audience = "AzureADMyOrg"

  identifier_uris = [
    "api://acctestServicePrincipal-%[1]d",
    "https://acctestServicePrincipal-%[1]d.net",
  ]

  api {
    oauth2_permission_scope {
      admin_consent_description  = "Administer the application"
      admin_consent_display_name = "Administer"
      enabled                    = true
      id                         = "%[2]s"
      type                       = "Admin"
      value                      = "administer"
    }

    oauth2_permission_scope {
      admin_consent_description  = "Allow the application to access acctest-APP-%[1]d on behalf of the signed-in user."
      admin_consent_display_name = "Access acctest-APP-%[1]d"
      enabled                    = true
      id                         = "%[3]s"
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
    id                   = "%[4]s"
    value                = "superAdmin"
  }

  app_role {
    allowed_member_types = ["User"]
    description          = "ReadOnly roles have limited query access"
    display_name         = "ReadOnly"
    enabled              = true
    id                   = "%[5]s"
    value                = "readOnlyUser"
  }

  web {
    homepage_url = "https://test-%[1]d.internal"
    logout_url   = "https://test-%[1]d.internal/logout"

    redirect_uris = [
      "https://test-%[1]d.internal/dashboard",
      "https://test-%[1]d.internal/account",
    ]
  }
}

resource "azuread_service_principal" "test" {
  application_id = azuread_application.test.application_id

  account_enabled               = false
  app_role_assignment_required  = true
  description                   = "An internal app for testing"
  login_url                     = "https://test-%[1]d.internal/login"
  notes                         = "Just testing something"
  preferred_single_sign_on_mode = "saml"

  notification_email_addresses = [
    "alerts.internal@hashitown.net",
    "cto@hashitown.net",
  ]

  saml_single_sign_on {
    relay_state = "/samlHome"
  }

  alternative_names = ["foo", "bar"]
  tags              = ["test", "multiple", "CapitalS"]
}
`, data.RandomInteger, data.UUID(), data.UUID(), data.UUID(), data.UUID())
}

func (ServicePrincipalResource) templateThreeUsers(data acceptance.TestData) string {
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

func (ServicePrincipalResource) noOwners(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application" "test" {
  display_name = "acctestServicePrincipal-%[1]d"
}

resource "azuread_service_principal" "test" {
  application_id = azuread_application.test.application_id
  owners         = []
}
`, data.RandomInteger)
}

func (r ServicePrincipalResource) singleOwner(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_application" "test" {
  display_name = "acctestServicePrincipal-%[2]d"
}

resource "azuread_service_principal" "test" {
  application_id = azuread_application.test.application_id
  owners = [
    azuread_user.testA.object_id,
  ]
}
`, r.templateThreeUsers(data), data.RandomInteger)
}

func (r ServicePrincipalResource) threeOwners(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_application" "test" {
  display_name = "acctestServicePrincipal-%[2]d"
}

resource "azuread_service_principal" "test" {
  application_id = azuread_application.test.application_id
  owners = [
    azuread_user.testA.object_id,
    azuread_user.testB.object_id,
    azuread_user.testC.object_id,
  ]
}
`, r.templateThreeUsers(data), data.RandomInteger)
}

func (r ServicePrincipalResource) manyOwners(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

data "azuread_client_config" "test" {}

data "azuread_domains" "test" {
  only_initial = true
}

resource "azuread_application" "owner" {
  count        = 27
  display_name = "acctestServicePrincipalOwner${count.index}-%[1]d"
}

resource "azuread_service_principal" "owner" {
  count          = 27
  application_id = azuread_application.owner[count.index].application_id
}

resource "azuread_user" "owner" {
  count               = 17
  user_principal_name = "acctestServicePrincipalOwner${count.index}-%[1]d@${data.azuread_domains.test.domains.0.domain_name}"
  display_name        = "acctestServicePrincipalOwner${count.index}-%[1]d"
  password            = "Qwer5678!@#"
}

resource "azuread_application" "test" {
  display_name = "acctestServicePrincipal-%[1]d"
}

resource "azuread_service_principal" "test" {
  application_id = azuread_application.test.application_id

  owners = flatten([
    data.azuread_client_config.test.object_id,
    azuread_service_principal.owner.*.object_id,
    azuread_user.owner.*.object_id,
  ])
}
`, data.RandomInteger)
}

func (ServicePrincipalResource) useExisting(_ acceptance.TestData) string {
	return `
provider "azuread" {}

resource "azuread_service_principal" "msgraph" {
  application_id = "00000003-0000-0000-c000-000000000000" # Microsoft Graph
  use_existing   = true
}
`
}

func (ServicePrincipalResource) fromApplicationTemplate(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application" "test" {
  display_name = "acctest-APP-%[1]d"
  template_id  = "%[2]s"
}

resource "azuread_service_principal" "test" {
  application_id = azuread_application.test.application_id
  use_existing   = true
}
`, data.RandomInteger, testApplicationTemplateId)
}

func (ServicePrincipalResource) threeServicePrincipalsABC(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application" "testA" {
  display_name = "acctestServicePrincipalA-%[1]d"
}

resource "azuread_application" "testB" {
  display_name = "acctestServicePrincipalB-%[1]d"
}

resource "azuread_application" "testC" {
  display_name = "acctestServicePrincipalC-%[1]d"
}

resource "azuread_service_principal" "testA" {
  application_id = azuread_application.testA.application_id
}

resource "azuread_service_principal" "testB" {
  application_id = azuread_application.testB.application_id
}

resource "azuread_service_principal" "testC" {
  application_id = azuread_application.testC.application_id
}
`, data.RandomInteger)
}
