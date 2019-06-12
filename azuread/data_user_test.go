package azuread

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccDataSourceAzureADUser_byUserPrincipalName(t *testing.T) {
	dataSourceName := "data.azuread_user.test"
	id := acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)
	password := id + "p@$$wR2"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureADUserDataSource_byUserPrincipalName(id, password),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataSourceName, "user_principal_name"),
					resource.TestCheckResourceAttrSet(dataSourceName, "account_enabled"),
					resource.TestCheckResourceAttrSet(dataSourceName, "display_name"),
					resource.TestCheckResourceAttrSet(dataSourceName, "mail_nickname"),
				),
			},
		},
	})
}

func TestAccDataSourceAzureADUser_byObjectId(t *testing.T) {
	dataSourceName := "data.azuread_user.test"
	id := acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)
	password := id + "p@$$wR2"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureADUserDataSource_byObjectId(id, password),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataSourceName, "user_principal_name"),
					resource.TestCheckResourceAttrSet(dataSourceName, "account_enabled"),
					resource.TestCheckResourceAttrSet(dataSourceName, "display_name"),
					resource.TestCheckResourceAttrSet(dataSourceName, "mail_nickname"),
				),
			},
		},
	})
}

func testAccAzureADUserDataSource_byUserPrincipalName(id, password string) string {
	return fmt.Sprintf(`
%s

data "azuread_user" "test" {
	user_principal_name = "${azuread_user.test.user_principal_name}"
}
`, testAccADUser_basic(id, password))
}

func testAccAzureADUserDataSource_byObjectId(id, password string) string {
	return fmt.Sprintf(`
%s

data "azuread_user" "test" {
	object_id = "${azuread_user.test.object_id}"
}
`, testAccADUser_basic(id, password))
}
