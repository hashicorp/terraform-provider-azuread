package azuread

import (
	"fmt"
	"testing"

	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccDataSourceAzureADGroup_byName(t *testing.T) {
	dataSourceName := "data.azuread_group.test"
	id, err := uuid.GenerateUUID()
	if err != nil {
		t.Fatal(err)
	}
	config := testAccDataSourceAzureADGroup_name(id)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureADGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureADGroup(id),
			},
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureADGroupExists(dataSourceName),
					resource.TestCheckResourceAttr(dataSourceName, "name", fmt.Sprintf("acctest%s", id)),
				),
			},
		},
	})
}

func testAccDataSourceAzureADGroup_name(id string) string {
	template := testAccAzureADGroup(id)
	return fmt.Sprintf(`
%s

data "azuread_group" "test" {
  name = "${azuread_group.test.name}"
}
`, template)
}
