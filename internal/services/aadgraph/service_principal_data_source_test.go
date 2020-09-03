package aadgraph_test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-azuread/internal/acceptance"
)

func TestAccServicePrincipalDataSource_byApplicationId(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_service_principal", "test")
	id := uuid.New().String()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckServicePrincipalDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccServicePrincipalDataSource_byApplicationId(id),
				Check: resource.ComposeTestCheckFunc(
					testCheckServicePrincipalExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "application_id"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "object_id"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "display_name"),
					resource.TestCheckResourceAttr(data.ResourceName, "app_roles.#", "0"),
					resource.TestCheckResourceAttr(data.ResourceName, "oauth2_permissions.#", "1"),
					resource.TestCheckResourceAttr(data.ResourceName, "oauth2_permissions.0.admin_consent_description", fmt.Sprintf("Allow the application to access %s on behalf of the signed-in user.", fmt.Sprintf("acctestApp-%s", id))),
				),
			},
		},
	})
}

func TestAccServicePrincipalDataSource_byDisplayName(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_service_principal", "test")
	id := uuid.New().String()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckServicePrincipalDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccServicePrincipalDataSource_byDisplayName(id),
				Check: resource.ComposeTestCheckFunc(
					testCheckServicePrincipalExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "application_id"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "object_id"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "display_name"),
				),
			},
		},
	})
}

func TestAccServicePrincipalDataSource_byObjectId(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_service_principal", "test")
	id := uuid.New().String()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckServicePrincipalDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccServicePrincipalDataSource_byObjectId(id),
				Check: resource.ComposeTestCheckFunc(
					testCheckServicePrincipalExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "application_id"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "object_id"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "display_name"),
				),
			},
		},
	})
}

func testAccServicePrincipalDataSource_byApplicationId(id string) string {
	template := testAccServicePrincipal_basic(id)
	return fmt.Sprintf(`
%s

data "azuread_service_principal" "test" {
  application_id = azuread_service_principal.test.application_id
}
`, template)
}

func testAccServicePrincipalDataSource_byDisplayName(id string) string {
	template := testAccServicePrincipal_basic(id)
	return fmt.Sprintf(`
%s

data "azuread_service_principal" "test" {
  display_name = azuread_service_principal.test.display_name
}
`, template)
}

func testAccServicePrincipalDataSource_byObjectId(id string) string {
	template := testAccServicePrincipal_basic(id)
	return fmt.Sprintf(`
%s

data "azuread_service_principal" "test" {
  object_id = azuread_service_principal.test.object_id
}
`, template)
}
