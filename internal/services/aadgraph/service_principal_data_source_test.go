package aadgraph_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/terraform-providers/terraform-provider-azuread/internal/acceptance"
	"github.com/terraform-providers/terraform-provider-azuread/internal/acceptance/check"
)

type ServicePrincipalDataSource struct{}

func TestAccServicePrincipalDataSource_byApplicationId(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_service_principal", "test")
	r := ServicePrincipalDataSource{}

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: r.byApplicationId(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("application_id").Exists(),
				check.That(data.ResourceName).Key("object_id").Exists(),
				check.That(data.ResourceName).Key("display_name").Exists(),
				check.That(data.ResourceName).Key("app_roles.#").HasValue("0"),
				check.That(data.ResourceName).Key("oauth2_permissions.#").HasValue("1"),
				check.That(data.ResourceName).Key("oauth2_permissions.0.admin_consent_description").HasValue(
					fmt.Sprintf("Allow the application to access %s on behalf of the signed-in user.",
						fmt.Sprintf("acctestApp-%d", data.RandomInteger))),
			),
		},
	})
}

func TestAccServicePrincipalDataSource_byDisplayName(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_service_principal", "test")
	r := ServicePrincipalDataSource{}

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: r.byDisplayName(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("application_id").Exists(),
				check.That(data.ResourceName).Key("object_id").Exists(),
				check.That(data.ResourceName).Key("display_name").Exists(),
				check.That(data.ResourceName).Key("app_roles.#").HasValue("0"),
				check.That(data.ResourceName).Key("oauth2_permissions.#").HasValue("1"),
				check.That(data.ResourceName).Key("oauth2_permissions.0.admin_consent_description").HasValue(
					fmt.Sprintf("Allow the application to access %s on behalf of the signed-in user.",
						fmt.Sprintf("acctestApp-%d", data.RandomInteger))),
			),
		},
	})
}

func TestAccServicePrincipalDataSource_byObjectId(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_service_principal", "test")
	r := ServicePrincipalDataSource{}

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: r.byObjectId(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("application_id").Exists(),
				check.That(data.ResourceName).Key("object_id").Exists(),
				check.That(data.ResourceName).Key("display_name").Exists(),
				check.That(data.ResourceName).Key("app_roles.#").HasValue("0"),
				check.That(data.ResourceName).Key("oauth2_permissions.#").HasValue("1"),
				check.That(data.ResourceName).Key("oauth2_permissions.0.admin_consent_description").HasValue(
					fmt.Sprintf("Allow the application to access %s on behalf of the signed-in user.",
						fmt.Sprintf("acctestApp-%d", data.RandomInteger))),
			),
		},
	})
}

func (ServicePrincipalDataSource) byApplicationId(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_service_principal" "test" {
  application_id = azuread_service_principal.test.application_id
}
`, ServicePrincipalResource{}.basic(data))
}

func (ServicePrincipalDataSource) byDisplayName(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_service_principal" "test" {
  display_name = azuread_service_principal.test.display_name
}
`, ServicePrincipalResource{}.basic(data))
}

func (ServicePrincipalDataSource) byObjectId(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_service_principal" "test" {
  object_id = azuread_service_principal.test.object_id
}
`, ServicePrincipalResource{}.basic(data))
}
