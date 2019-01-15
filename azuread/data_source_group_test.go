package azuread

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/ar"
)

func TestAccDataSourceAzureADGroup_byDisplayName(t *testing.T) {
	dataSourceName := "data.azuread_group.test"
	id := uuid.New().String()
	config := testAccDataSourceAzureRMAzureADGroup_byDisplayName(id)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		// CheckDestroy:
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMActiveDirectoryServiceGroupExists(dataSourceName),
					resource.TestCheckResourceAttrSet(dataSourceName, "mail"),
					resource.TestCheckResourceAttrSet(dataSourceName, "object_id"),
					resource.TestCheckResourceAttrSet(dataSourceName, "display_name"),
					resource.TestCheckResourceAttrSet(dataSourceName, "id"),
				),
			},
		},
	})
}

func TestAccDataSourceAzureRMAzureADGroup_byObjectId(t *testing.T) {
	dataSourceName := "data.azuread_group.test"
	id := uuid.New().String()
	config := testAccDataSourceAzureRMAzureADGroup_byObjectId(id)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		// CheckDestroy:
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMActiveDirectoryServiceGroupExists(dataSourceName),
					resource.TestCheckResourceAttrSet(dataSourceName, "mail"),
					resource.TestCheckResourceAttrSet(dataSourceName, "object_id"),
					resource.TestCheckResourceAttrSet(dataSourceName, "display_name"),
					resource.TestCheckResourceAttrSet(dataSourceName, "id"),
				),
			},
		},
	})
}

func testAccDataSourceAzureRMAzureADGroup_byDisplayName(id string) string {
	return fmt.Sprintf(`

data "azuread_group" "test" {
  display_name = "my-cool-group"
}
`)
}

func testAccDataSourceAzureRMAzureADGroup_byObjectId(id string) string {
	return fmt.Sprintf(`

data "azuread_group" "test" {
  object_id = "30a6a0d5-fcb4-xxxx-xxxx-98bfb06a67a1"
}
`)
}

func testCheckAzureRMActiveDirectoryServiceGroupExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %q", name)
		}

		client := testAccProvider.Meta().(*ArmClient).groupsClient
		ctx := testAccProvider.Meta().(*ArmClient).StopContext
		resp, err := client.Get(ctx, rs.Primary.ID)

		if err != nil {
			if ar.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("Bad: Azure AD Group %q does not exist", rs.Primary.ID)
			}
			return fmt.Errorf("Bad: Get on Azure AD groupsClient: %+v", err)
		}

		return nil
	}
}
