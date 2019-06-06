package azuread

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccAzureADServicePrincipalDataSource_byApplicationId(t *testing.T) {
	dataSourceName := "data.azuread_service_principal.test"
	id := uuid.New().String()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADServicePrincipalDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureADServicePrincipalDataSource_byApplicationId(id),
				Check: resource.ComposeTestCheckFunc(
					testCheckADServicePrincipalExists(dataSourceName),
					resource.TestCheckResourceAttrSet(dataSourceName, "application_id"),
					resource.TestCheckResourceAttrSet(dataSourceName, "object_id"),
					resource.TestCheckResourceAttrSet(dataSourceName, "display_name"),
				),
			},
		},
	})
}

func TestAccAzureADServicePrincipalDataSource_byDisplayName(t *testing.T) {
	dataSourceName := "data.azuread_service_principal.test"
	id := uuid.New().String()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADServicePrincipalDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureADServicePrincipalDataSource_byDisplayName(id),
				Check: resource.ComposeTestCheckFunc(
					testCheckADServicePrincipalExists(dataSourceName),
					resource.TestCheckResourceAttrSet(dataSourceName, "application_id"),
					resource.TestCheckResourceAttrSet(dataSourceName, "object_id"),
					resource.TestCheckResourceAttrSet(dataSourceName, "display_name"),
				),
			},
		},
	})
}

func TestAccAzureADServicePrincipalDataSource_byObjectId(t *testing.T) {
	dataSourceName := "data.azuread_service_principal.test"
	id := uuid.New().String()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADServicePrincipalDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureADServicePrincipalDataSource_byObjectId(id),
				Check: resource.ComposeTestCheckFunc(
					testCheckADServicePrincipalExists(dataSourceName),
					resource.TestCheckResourceAttrSet(dataSourceName, "application_id"),
					resource.TestCheckResourceAttrSet(dataSourceName, "object_id"),
					resource.TestCheckResourceAttrSet(dataSourceName, "display_name"),
				),
			},
		},
	})
}

func testAccAzureADServicePrincipalDataSource_byApplicationId(id string) string {
	template := testAccADServicePrincipal_basic(id)
	return fmt.Sprintf(`
%s

data "azuread_service_principal" "test" {
  application_id = "${azuread_service_principal.test.application_id}"
}
`, template)
}

func testAccAzureADServicePrincipalDataSource_byDisplayName(id string) string {
	template := testAccADServicePrincipal_basic(id)
	return fmt.Sprintf(`
%s

data "azuread_service_principal" "test" {
  display_name = "${azuread_service_principal.test.display_name}"
}
`, template)
}

func testAccAzureADServicePrincipalDataSource_byObjectId(id string) string {
	template := testAccADServicePrincipal_basic(id)
	return fmt.Sprintf(`
%s

data "azuread_service_principal" "test" {
  object_id = "${azuread_service_principal.test.object_id}"
}
`, template)
}
