package azuread

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/tf"
)

func TestAccAzureADGroupsDataSource_byUserPrincipalNames(t *testing.T) {
	dsn := "data.azuread_groups.test"
	id := tf.AccRandTimeInt()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureADGroupsDataSource_byDisplayNames(id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dsn, "names.#", "2"),
					resource.TestCheckResourceAttr(dsn, "object_ids.#", "2"),
				),
			},
		},
	})
}

func TestAccAzureADGroupsDataSource_byObjectIds(t *testing.T) {
	dsn := "data.azuread_groups.test"
	id := tf.AccRandTimeInt()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureADGroupsDataSource_byObjectIds(id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dsn, "names.#", "2"),
					resource.TestCheckResourceAttr(dsn, "object_ids.#", "2"),
				),
			},
		},
	})
}

func TestAccAzureADGroupsDataSource_noNames(t *testing.T) {
	dsn := "data.azuread_groups.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureADGroupsDataSource_noNames(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dsn, "names.#", "0"),
					resource.TestCheckResourceAttr(dsn, "object_ids.#", "0"),
				),
			},
		},
	})
}

func testAccAzureADGroup_multiple(id int) string {
	return fmt.Sprintf(`
resource "azuread_group" "testA" {
  name    = "acctestGroup-%[1]d"
  members = []
}

resource "azuread_group" "testB" {
  name    = "acctestGroup-%[1]d"
  members = []
}
`, id)
}

func testAccAzureADGroupsDataSource_byDisplayNames(id int) string {
	return fmt.Sprintf(`
%s

data "azuread_groups" "test" {
  names = [azuread_group.testA.name, azuread_group.testB.name]
}
`, testAccAzureADGroup_multiple(id))
}

func testAccAzureADGroupsDataSource_byObjectIds(id int) string {
	return fmt.Sprintf(`
%s

data "azuread_groups" "test" {
  object_ids = [azuread_group.testA.object_id, azuread_group.testB.object_id]
}
`, testAccAzureADGroup_multiple(id))
}
func testAccAzureADGroupsDataSource_noNames() string {
	return `
data "azuread_groups" "test" {
  names = []
}
`
}
