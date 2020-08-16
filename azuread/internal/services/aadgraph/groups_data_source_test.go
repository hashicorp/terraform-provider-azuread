package aadgraph_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-azuread/azuread/internal/acceptance"
)

func TestAccGroupsDataSource_byUserPrincipalNames(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_groups", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { acceptance.PreCheck(t) },
		Providers: acceptance.SupportedProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccGroupsDataSource_byDisplayNames(data.RandomInteger),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(data.ResourceName, "names.#", "2"),
					resource.TestCheckResourceAttr(data.ResourceName, "object_ids.#", "2"),
				),
			},
		},
	})
}

func TestAccGroupsDataSource_byObjectIds(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_groups", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { acceptance.PreCheck(t) },
		Providers: acceptance.SupportedProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccGroupsDataSource_byObjectIds(data.RandomInteger),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(data.ResourceName, "names.#", "2"),
					resource.TestCheckResourceAttr(data.ResourceName, "object_ids.#", "2"),
				),
			},
		},
	})
}

func TestAccGroupsDataSource_noNames(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_groups", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { acceptance.PreCheck(t) },
		Providers: acceptance.SupportedProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccGroupsDataSource_noNames(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(data.ResourceName, "names.#", "0"),
					resource.TestCheckResourceAttr(data.ResourceName, "object_ids.#", "0"),
				),
			},
		},
	})
}

func testAccGroup_multiple(id int) string {
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

func testAccGroupsDataSource_byDisplayNames(id int) string {
	return fmt.Sprintf(`
%s

data "azuread_groups" "test" {
  names = [azuread_group.testA.name, azuread_group.testB.name]
}
`, testAccGroup_multiple(id))
}

func testAccGroupsDataSource_byObjectIds(id int) string {
	return fmt.Sprintf(`
%s

data "azuread_groups" "test" {
  object_ids = [azuread_group.testA.object_id, azuread_group.testB.object_id]
}
`, testAccGroup_multiple(id))
}
func testAccGroupsDataSource_noNames() string {
	return `
data "azuread_groups" "test" {
  names = []
}
`
}
