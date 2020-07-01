package azuread

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

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
				Config: testAccDataSourceAzureADGroup_name(id),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureADGroupExists(dsn),
					resource.TestCheckResourceAttr(dsn, "name", fmt.Sprintf("acctestGroup-%d", id)),
				),
			},
		},
	})
}

func TestAccDataSourceAzureADGroup_byCaseInsensitiveName(t *testing.T) {
	dsn := "data.azuread_group.test"
	id := tf.AccRandTimeInt()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureADGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceAzureADGroup_caseInsensitiveName(id),
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
				Config: testAccDataSourceAzureADGroup_objectId(id),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureADGroupExists(dsn),
					resource.TestCheckResourceAttr(dsn, "name", fmt.Sprintf("acctestGroup-%d", id)),
				),
			},
		},
	})
}

func TestAccDataSourceAzureADGroup_members(t *testing.T) {
	dsn := "data.azuread_group.test"
	id := tf.AccRandTimeInt()
	pw := "p@$$wR2" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureADGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceAzureADGroup_members(id, pw),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureADGroupExists(dsn),
					resource.TestCheckResourceAttr(dsn, "name", fmt.Sprintf("acctestGroup-%d", id)),
					resource.TestCheckResourceAttr(dsn, "members.#", "3"),
				),
			},
		},
	})
}

func TestAccDataSourceAzureADGroup_owners(t *testing.T) {
	dsn := "data.azuread_group.test"
	id := tf.AccRandTimeInt()
	pw := "p@$$wR2" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureADGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceAzureADGroup_owners(id, pw),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureADGroupExists(dsn),
					resource.TestCheckResourceAttr(dsn, "name", fmt.Sprintf("acctestGroup-%d", id)),
					resource.TestCheckResourceAttr(dsn, "owners.#", "3"),
				),
			},
		},
	})
}

func testAccDataSourceAzureADGroup_name(id int) string {
	return fmt.Sprintf(`
%s

data "azuread_group" "test" {
  name = azuread_group.test.name
}
`, testAccAzureADGroup_basic(id))
}

func testAccDataSourceAzureADGroup_caseInsensitiveName(id int) string {
	return fmt.Sprintf(`
%s

data "azuread_group" "test" {
  name = upper(azuread_group.test.name)
}
`, testAccAzureADGroup_basic(id))
}

func testAccDataSourceAzureADGroup_objectId(id int) string {
	return fmt.Sprintf(`
%s

data "azuread_group" "test" {
  object_id = azuread_group.test.object_id
}
`, testAccAzureADGroup_basic(id))
}

func testAccDataSourceAzureADGroup_members(id int, password string) string {
	return fmt.Sprintf(`
%s

data "azuread_group" "test" {
  object_id = azuread_group.test.object_id
}
`, testAccAzureADGroupWithThreeMembers(id, password))
}

func testAccDataSourceAzureADGroup_owners(id int, password string) string {
	return fmt.Sprintf(`
%s

data "azuread_group" "test" {
  object_id = azuread_group.test.object_id
}
`, testAccAzureADGroupWithThreeOwners(id, password))
}
