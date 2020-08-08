package tests

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-azuread/azuread/internal/acceptance"
)

func TestAccDataSourceGroup_byName(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_group", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGroup_name(data.RandomInteger),
				Check: resource.ComposeTestCheckFunc(
					testCheckGroupExists(data.ResourceName),
					resource.TestCheckResourceAttr(data.ResourceName, "name", fmt.Sprintf("acctestGroup-%d", data.RandomInteger)),
				),
			},
		},
	})
}

func TestAccDataSourceGroup_byObjectId(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_group", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGroup_objectId(data.RandomInteger),
				Check: resource.ComposeTestCheckFunc(
					testCheckGroupExists(data.ResourceName),
					resource.TestCheckResourceAttr(data.ResourceName, "name", fmt.Sprintf("acctestGroup-%d", data.RandomInteger)),
				),
			},
		},
	})
}

func TestAccDataSourceGroup_members(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_group", "test")
	pw := "p@$$wR2" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGroup_members(data.RandomInteger, pw),
				Check: resource.ComposeTestCheckFunc(
					testCheckGroupExists(data.ResourceName),
					resource.TestCheckResourceAttr(data.ResourceName, "name", fmt.Sprintf("acctestGroup-%d", data.RandomInteger)),
					resource.TestCheckResourceAttr(data.ResourceName, "members.#", "3"),
				),
			},
		},
	})
}

func TestAccDataSourceGroup_owners(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_group", "test")
	pw := "p@$$wR2" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGroup_owners(data.RandomInteger, pw),
				Check: resource.ComposeTestCheckFunc(
					testCheckGroupExists(data.ResourceName),
					resource.TestCheckResourceAttr(data.ResourceName, "name", fmt.Sprintf("acctestGroup-%d", data.RandomInteger)),
					resource.TestCheckResourceAttr(data.ResourceName, "owners.#", "3"),
				),
			},
		},
	})
}

func testAccDataSourceGroup_name(id int) string {
	return fmt.Sprintf(`
%s

data "azuread_group" "test" {
  name = azuread_group.test.name
}
`, testAccGroup_basic(id))
}

func testAccDataSourceGroup_objectId(id int) string {
	return fmt.Sprintf(`
%s

data "azuread_group" "test" {
  object_id = azuread_group.test.object_id
}
`, testAccGroup_basic(id))
}

func testAccDataSourceGroup_members(id int, password string) string {
	return fmt.Sprintf(`
%s

data "azuread_group" "test" {
  object_id = azuread_group.test.object_id
}
`, testAccGroupWithThreeMembers(id, password))
}

func testAccDataSourceGroup_owners(id int, password string) string {
	return fmt.Sprintf(`
%s

data "azuread_group" "test" {
  object_id = azuread_group.test.object_id
}
`, testAccGroupWithThreeOwners(id, password))
}
