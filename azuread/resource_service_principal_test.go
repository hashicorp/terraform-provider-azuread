package azuread

import (
	"fmt"
	"testing"

	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/ar"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccAzureADServicePrincipal_basic(t *testing.T) {
	resourceName := "azuread_service_principal.test"
	id := uuid.New().String()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADServicePrincipalDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADServicePrincipal_basic(id),
				Check: resource.ComposeTestCheckFunc(
					testCheckADServicePrincipalExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "display_name"),
					resource.TestCheckResourceAttrSet(resourceName, "application_id"),
					resource.TestCheckResourceAttrSet(resourceName, "object_id"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAzureADServicePrincipal_complete(t *testing.T) {
	resourceName := "azuread_service_principal.test"
	id := uuid.New().String()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADServicePrincipalDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADServicePrincipal_complete(id),
				Check: resource.ComposeTestCheckFunc(
					testCheckADServicePrincipalExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "3"),
					resource.TestCheckResourceAttrSet(resourceName, "object_id"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testCheckADServicePrincipalExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %q", name)
		}

		client := testAccProvider.Meta().(*ArmClient).servicePrincipalsClient
		ctx := testAccProvider.Meta().(*ArmClient).StopContext
		resp, err := client.Get(ctx, rs.Primary.ID)

		if err != nil {
			if ar.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("Bad: Azure AD Service Principal %q does not exist", rs.Primary.ID)
			}
			return fmt.Errorf("Bad: Get on Azure AD servicePrincipalsClient: %+v", err)
		}

		return nil
	}
}

func testCheckADServicePrincipalDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "azuread_service_principal" {
			continue
		}

		client := testAccProvider.Meta().(*ArmClient).servicePrincipalsClient
		ctx := testAccProvider.Meta().(*ArmClient).StopContext
		resp, err := client.Get(ctx, rs.Primary.ID)

		if err != nil {
			if ar.ResponseWasNotFound(resp.Response) {
				return nil
			}

			return err
		}

		return fmt.Errorf("Azure AD Service Principal still exists:\n%#v", resp)
	}

	return nil
}

func testAccADServicePrincipal_basic(id string) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name = "acctestspa%s"
}

resource "azuread_service_principal" "test" {
  application_id = "${azuread_application.test.application_id}"
}
`, id)
}

func testAccADServicePrincipal_complete(id string) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name = "acctestspa%s"
}

resource "azuread_service_principal" "test" {
  application_id = "${azuread_application.test.application_id}"

  tags = ["test", "multiple", "CapitalS"]
}
`, id)
}
