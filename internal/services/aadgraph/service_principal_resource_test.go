package aadgraph_test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-azuread/internal/acceptance"
	"github.com/terraform-providers/terraform-provider-azuread/internal/clients"
	"github.com/terraform-providers/terraform-provider-azuread/internal/utils"
)

func TestAccServicePrincipal_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_service_principal", "test")
	id := uuid.New().String()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckServicePrincipalDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccServicePrincipal_basic(id),
				Check: resource.ComposeTestCheckFunc(
					testCheckServicePrincipalExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "application_id"),
					resource.TestCheckResourceAttr(data.ResourceName, "oauth2_permissions.#", "1"),
					resource.TestCheckResourceAttr(data.ResourceName, "app_role_assignment_required", "false"),
					resource.TestCheckResourceAttr(data.ResourceName, "oauth2_permissions.0.admin_consent_description", fmt.Sprintf("Allow the application to access %s on behalf of the signed-in user.", fmt.Sprintf("acctestApp-%s", id))),
					resource.TestCheckResourceAttrSet(data.ResourceName, "object_id"),
				),
			},
			data.ImportStep(),
		},
	})
}

func TestAccServicePrincipal_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_service_principal", "test")
	id := uuid.New().String()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckServicePrincipalDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccServicePrincipal_complete(id),
				Check: resource.ComposeTestCheckFunc(
					testCheckServicePrincipalExists(data.ResourceName),
					resource.TestCheckResourceAttr(data.ResourceName, "app_role_assignment_required", "true"),
					resource.TestCheckResourceAttr(data.ResourceName, "tags.#", "3"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "object_id"),
				),
			},
			data.ImportStep(),
		},
	})
}

func TestAccServicePrincipal_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_service_principal", "test")
	id := uuid.New().String()
	updatedId := uuid.New().String()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckServicePrincipalDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccServicePrincipal_basic(id),
				Check: resource.ComposeTestCheckFunc(
					testCheckServicePrincipalExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "application_id"),
					resource.TestCheckResourceAttr(data.ResourceName, "app_role_assignment_required", "false"),
				),
			},
			data.ImportStep(),
			{
				Config: testAccServicePrincipal_complete(updatedId),
				Check: resource.ComposeTestCheckFunc(
					testCheckServicePrincipalExists(data.ResourceName),
					resource.TestCheckResourceAttr(data.ResourceName, "tags.#", "3"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "object_id"),
					resource.TestCheckResourceAttr(data.ResourceName, "app_role_assignment_required", "true"),
				),
			},
			data.ImportStep(),
			{
				Config: testAccServicePrincipal_basic(id),
				Check: resource.ComposeTestCheckFunc(
					testCheckServicePrincipalExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "display_name"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "application_id"),
					resource.TestCheckResourceAttr(data.ResourceName, "app_role_assignment_required", "false"),
				),
			},
			data.ImportStep(),
		},
	})
}

func testCheckServicePrincipalExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %q", name)
		}

		client := acceptance.AzureADProvider.Meta().(*clients.AadClient).AadGraph.ServicePrincipalsClient
		ctx := acceptance.AzureADProvider.Meta().(*clients.AadClient).StopContext
		resp, err := client.Get(ctx, rs.Primary.ID)

		if err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("Bad: Azure AD Service Principal %q does not exist", rs.Primary.ID)
			}
			return fmt.Errorf("Bad: Get on Azure AD ServicePrincipalsClient: %+v", err)
		}

		return nil
	}
}

func testCheckServicePrincipalDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "azuread_service_principal" {
			continue
		}

		client := acceptance.AzureADProvider.Meta().(*clients.AadClient).AadGraph.ServicePrincipalsClient
		ctx := acceptance.AzureADProvider.Meta().(*clients.AadClient).StopContext
		resp, err := client.Get(ctx, rs.Primary.ID)

		if err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
				return nil
			}

			return err
		}

		return fmt.Errorf("Azure AD Service Principal still exists:\n%#v", resp)
	}

	return nil
}

func testAccServicePrincipal_basic(id string) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name = "acctestApp-%s"
}

resource "azuread_service_principal" "test" {
  application_id = azuread_application.test.application_id
}
`, id)
}

func testAccServicePrincipal_complete(id string) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name = "acctesttApp-%s"
}

resource "azuread_service_principal" "test" {
  application_id               = azuread_application.test.application_id
  app_role_assignment_required = true

  tags = ["test", "multiple", "CapitalS"]
}
`, id)
}
