package azuread

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/acctest"
	"testing"

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

func testAccAzureADUserDataSource_byUserPrincipalName(id, password string) string {
	template := testAccADUser_basic(id, password)
	return fmt.Sprintf(`

%s

data "azuread_user" "test" {
	user_principal_name = "${azuread_user.test.user_principal_name}"
}
`, template)
}
