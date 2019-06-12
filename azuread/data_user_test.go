package azuread

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"

	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/tf"
)

func TestAccAzureADUserDataSource_byUserPrincipalName(t *testing.T) {
	dsn := "data.azuread_user.test"
	id := tf.AccRandTimeInt()
	password := "p@$$wR2" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureADUserDataSource_byUserPrincipalName(id, password),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dsn, "user_principal_name"),
					resource.TestCheckResourceAttrSet(dsn, "account_enabled"),
					resource.TestCheckResourceAttrSet(dsn, "display_name"),
					resource.TestCheckResourceAttrSet(dsn, "mail_nickname"),
				),
			},
		},
	})
}

func TestAccAzureADUserDataSource_byObjectId(t *testing.T) {
	dsn := "data.azuread_user.test"
	id := tf.AccRandTimeInt()
	password := "p@$$wR2" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureADUserDataSource_byObjectId(id, password),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dsn, "user_principal_name"),
					resource.TestCheckResourceAttrSet(dsn, "account_enabled"),
					resource.TestCheckResourceAttrSet(dsn, "display_name"),
					resource.TestCheckResourceAttrSet(dsn, "mail_nickname"),
				),
			},
		},
	})
}

func testAccAzureADUserDataSource_byUserPrincipalName(id int, password string) string {
	return fmt.Sprintf(`
%s

data "azuread_user" "test" {
	user_principal_name = "${azuread_user.test.user_principal_name}"
}
`, testAccADUser_basic(id, password))
}

func testAccAzureADUserDataSource_byObjectId(id int, password string) string {
	return fmt.Sprintf(`
%s

data "azuread_user" "test" {
	object_id = "${azuread_user.test.object_id}"
}
`, testAccADUser_basic(id, password))
}
