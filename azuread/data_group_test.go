package azuread

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/tf"
)

func TestAccDataSourceAzureADGroup_byName(t *testing.T) {
	dsn := "data.azuread_group.test"
	id := tf.AccRandTimeInt()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureADGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureADGroup_basic(id),
			},
			{
				Config: testAccDataSourceAzureADGroup_name(id),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureADGroupExists(dsn),
					resource.TestCheckResourceAttr(dsn, "name", fmt.Sprintf("acctestGroup-%d", id)),
				),
			},
		},
	})
}

func TestAccDataSourceAzureADGroup_byObjectId(t *testing.T) {
	dsn := "data.azuread_group.test"
	id := tf.AccRandTimeInt()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureADGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureADGroup_basic(id),
			},
			{
				Config: testAccDataSourceAzureADGroup_objectId(id),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureADGroupExists(dsn),
					resource.TestCheckResourceAttr(dsn, "name", fmt.Sprintf("acctestGroup-%d", id)),
				),
			},
		},
	})
}

func testAccDataSourceAzureADGroup_name(id int) string {
	return fmt.Sprintf(`
%s

data "azuread_group" "test" {
  name = "${azuread_group.test.name}"
}
`, testAccAzureADGroup_basic(id))
}

func testAccDataSourceAzureADGroup_objectId(id int) string {
	return fmt.Sprintf(`
%s

data "azuread_group" "test" {
  object_id = "${azuread_group.test.object_id}"
}
`, testAccAzureADGroup_basic(id))
}
