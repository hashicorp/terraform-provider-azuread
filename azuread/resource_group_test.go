package azuread

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/ar"
)

func TestAccAzureADGroup_basic(t *testing.T) {
	resourceName := "azuread_group.test"
	id, err := uuid.GenerateUUID()
	if err != nil {
		t.Fatal(err)
	}
	config := testAccAzureADGroup(id)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureADGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureADGroupExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("acctest%s", id)),
					resource.TestCheckResourceAttrSet(resourceName, "object_id"),
					resource.TestCheckNoResourceAttr(resourceName, "members"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAzureADGroup_members(t *testing.T) {
	resourceName := "azuread_group.test"
	id, err := uuid.GenerateUUID()
	if err != nil {
		t.Fatal(err)
	}

	members := make([]string, 0)

	for i := 0; i < 5; i++ {
		memberUuid, err := uuid.GenerateUUID()
		if err != nil {
			t.Fatal(err)
		}
		members = append(members, "\""+memberUuid+"\"")
	}

	config := testAccAzureADGroupWithMembers(id, members)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureADGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureADGroupExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("acctest%s", id)),
					resource.TestCheckResourceAttrSet(resourceName, "object_id"),
					resource.TestCheckResourceAttr(resourceName, "members", fmt.Sprintf("[ %s ]", strings.Join(members, ", "))),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAzureADGroup_complete(t *testing.T) {
	resourceName := "azuread_group.test"
	id, err := uuid.GenerateUUID()
	if err != nil {
		t.Fatal(err)
	}
	config := testAccAzureADGroup(id)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureADGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureADGroupExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("acctest%s", id)),
					resource.TestCheckResourceAttrSet(resourceName, "object_id"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testCheckAzureADGroupExists(name string) resource.TestCheckFunc {
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

func testCheckAzureADGroupDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "azuread_group" {
			continue
		}

		client := testAccProvider.Meta().(*ArmClient).groupsClient
		ctx := testAccProvider.Meta().(*ArmClient).StopContext
		resp, err := client.Get(ctx, rs.Primary.ID)

		if err != nil {
			if ar.ResponseWasNotFound(resp.Response) {
				return nil
			}

			return err
		}

		return fmt.Errorf("Azure AD group still exists:\n%#v", resp)
	}

	return nil
}

func testAccAzureADGroup(id string) string {
	return fmt.Sprintf(`
resource "azuread_group" "test" {
  name = "acctest%s"
}
`, id)
}

func testAccAzureADGroupWithMembers(id string, members []string) string {
	return fmt.Sprintf(`
resource "azuread_group" "test" {
  name = "acctest%s"
  members = [ %s ]
}
`, id, strings.Join(members, ", "))
}
