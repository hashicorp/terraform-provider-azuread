package azuread

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccDataSourceAzureADGroup_byName(t *testing.T) {
	dataSourceName := "data.azuread_group.test"
	id := uuid.New().String()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureADGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureADGroup(id),
			},
			{
				Config: testAccDataSourceAzureADGroup_name(id),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureADGroupExists(dataSourceName),
					resource.TestCheckResourceAttr(dataSourceName, "name", fmt.Sprintf("acctest%s", id)),
				),
			},
		},
	})
}

func TestAccDataSourceAzureADGroup_byObjectId(t *testing.T) {
	dataSourceName := "data.azuread_group.test"
	id := uuid.New().String()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureADGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureADGroup(id),
			},
			{
				Config: testAccDataSourceAzureADGroup_objectId(id),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureADGroupExists(dataSourceName),
					resource.TestCheckResourceAttr(dataSourceName, "name", fmt.Sprintf("acctest%s", id)),
				),
			},
		},
	})
}

func testAccDataSourceAzureADGroup_name(id string) string {
	return fmt.Sprintf(`
%s

data "azuread_group" "test" {
  name = "${azuread_group.test.name}"
}
`, testAccAzureADGroup(id))
}

func testAccDataSourceAzureADGroup_objectId(id string) string {
	return fmt.Sprintf(`
%s

data "azuread_group" "test" {
  object_id = "${azuread_group.test.object_id}"
}
`, testAccAzureADGroup(id))
}
