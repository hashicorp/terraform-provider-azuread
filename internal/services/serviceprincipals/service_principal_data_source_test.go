package serviceprincipals_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
)

type ServicePrincipalDataSource struct{}

func TestAccServicePrincipalDataSource_byApplicationId(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_service_principal", "test")
	r := ServicePrincipalDataSource{}

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: r.byApplicationId(data),
			Check:  r.testCheckFunc(data),
		},
	})
}

func TestAccServicePrincipalDataSource_byDisplayName(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_service_principal", "test")
	r := ServicePrincipalDataSource{}

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: r.byDisplayName(data),
			Check:  r.testCheckFunc(data),
		},
	})
}

func TestAccServicePrincipalDataSource_byObjectId(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_service_principal", "test")
	r := ServicePrincipalDataSource{}

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: r.byObjectId(data),
			Check:  r.testCheckFunc(data),
		},
	})
}

func (ServicePrincipalDataSource) testCheckFunc(data acceptance.TestData) resource.TestCheckFunc {
	tenantId := os.Getenv("ARM_TENANT_ID")
	return resource.ComposeTestCheckFunc(
		check.That(data.ResourceName).Key("account_enabled").HasValue("false"),
		check.That(data.ResourceName).Key("alternative_names.#").HasValue("2"),
		check.That(data.ResourceName).Key("app_role_assignment_required").HasValue("true"),
		check.That(data.ResourceName).Key("app_role_ids.%").HasValue("2"),
		check.That(data.ResourceName).Key("app_roles.#").HasValue("2"),
		check.That(data.ResourceName).Key("application_id").IsUuid(),
		check.That(data.ResourceName).Key("application_tenant_id").HasValue(tenantId),
		check.That(data.ResourceName).Key("description").HasValue("An internal app for testing"),
		check.That(data.ResourceName).Key("display_name").Exists(),
		check.That(data.ResourceName).Key("feature_tags.#").HasValue("1"),
		check.That(data.ResourceName).Key("feature_tags.0.custom_single_sign_on").HasValue("true"),
		check.That(data.ResourceName).Key("feature_tags.0.enterprise").HasValue("true"),
		check.That(data.ResourceName).Key("feature_tags.0.gallery").HasValue("true"),
		check.That(data.ResourceName).Key("feature_tags.0.hide").HasValue("true"),
		check.That(data.ResourceName).Key("homepage_url").HasValue(fmt.Sprintf("https://test-%d.internal", data.RandomInteger)),
		check.That(data.ResourceName).Key("login_url").HasValue(fmt.Sprintf("https://test-%d.internal/login", data.RandomInteger)),
		check.That(data.ResourceName).Key("logout_url").HasValue(fmt.Sprintf("https://test-%d.internal/logout", data.RandomInteger)),
		check.That(data.ResourceName).Key("notes").HasValue("Just testing something"),
		check.That(data.ResourceName).Key("notification_email_addresses.#").HasValue("2"),
		check.That(data.ResourceName).Key("oauth2_permission_scope_ids.%").HasValue("2"),
		check.That(data.ResourceName).Key("oauth2_permission_scopes.#").HasValue("2"),
		check.That(data.ResourceName).Key("object_id").IsUuid(),
		check.That(data.ResourceName).Key("redirect_uris.#").HasValue("2"),
		check.That(data.ResourceName).Key("saml_single_sign_on.#").HasValue("1"),
		check.That(data.ResourceName).Key("saml_single_sign_on.0.relay_state").HasValue("/samlHome"),
		check.That(data.ResourceName).Key("service_principal_names.#").HasValue("2"),
		check.That(data.ResourceName).Key("sign_in_audience").HasValue("AzureADMyOrg"),
		check.That(data.ResourceName).Key("tags.#").HasValue("4"),
		check.That(data.ResourceName).Key("type").HasValue("Application"),
	)
}

func (ServicePrincipalDataSource) byApplicationId(data acceptance.TestData) string {
	endDate := time.Now().AddDate(0, 3, 27).UTC().Format(time.RFC3339)
	return fmt.Sprintf(`
%[1]s

data "azuread_service_principal" "test" {
  application_id = azuread_service_principal.test.application_id
}
`, ServicePrincipalResource{}.complete(data, endDate))
}

func (ServicePrincipalDataSource) byDisplayName(data acceptance.TestData) string {
	endDate := time.Now().AddDate(0, 3, 27).UTC().Format(time.RFC3339)
	return fmt.Sprintf(`
%[1]s

data "azuread_service_principal" "test" {
  display_name = azuread_service_principal.test.display_name
}
`, ServicePrincipalResource{}.complete(data, endDate))
}

func (ServicePrincipalDataSource) byObjectId(data acceptance.TestData) string {
	endDate := time.Now().AddDate(0, 3, 27).UTC().Format(time.RFC3339)
	return fmt.Sprintf(`
%[1]s

data "azuread_service_principal" "test" {
  object_id = azuread_service_principal.test.object_id
}
`, ServicePrincipalResource{}.complete(data, endDate))
}
