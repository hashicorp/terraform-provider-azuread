package azuread

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/ar"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/graph"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/tf"
)

func TestAccAzureADGroupOwner_user(t *testing.T) {
	rna := "azuread_group_owner.testA"
	rnb := "azuread_group_owner.testB"
	id := tf.AccRandTimeInt()
	pw := "p@$$wR2" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureADGroupOwnerDestroy(rna),
		Steps: []resource.TestStep{
			{
				Config: testAccAzureADGroupOwner_oneUser(id, pw),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(rna, "group_object_id"),
					resource.TestCheckResourceAttrSet(rna, "owner_object_id"),
				),
			},
			{
				Config: testAccAzureADGroupOwner_twoUsers(id, pw),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(rna, "group_object_id"),
					resource.TestCheckResourceAttrSet(rna, "owner_object_id"),
					resource.TestCheckResourceAttrSet(rnb, "group_object_id"),
					resource.TestCheckResourceAttrSet(rnb, "owner_object_id"),
				),
			},
			// we rerun the config so the group resource updates with the number of members
			{
				Config: testAccAzureADGroupOwner_twoUsers(id, pw),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("azuread_group.test", "owners.#", "2"),
				),
			},
			{
				ResourceName:      rnb,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccAzureADGroupOwner_oneUser(id, pw),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(rna, "group_object_id"),
					resource.TestCheckResourceAttrSet(rna, "owner_object_id"),
				),
			},
			{
				Config: testAccAzureADGroupOwner_twoUsers(id, pw),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("azuread_group.test", "owners.#", "1"),
				),
			},
			{
				ResourceName:      rna,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAzureADGroupOwner_ServicePrincipal(t *testing.T) {
	rn := "azuread_group_owner.test"
	id := tf.AccRandTimeInt()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureADGroupOwnerDestroy(rn),
		Steps: []resource.TestStep{
			{
				Config: testAccAzureADGroupOwner_ServicePrincipal(id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(rn, "group_object_id"),
					resource.TestCheckResourceAttrSet(rn, "owner_object_id"),
				),
			},
			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testCheckAzureADGroupOwnerDestroy(ignoreResource string) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "azuread_group_owner" {
				continue
			}

			if rs.Primary.ID == ignoreResource {
				continue // there always has to be one owner remaining so lets ignore it
			}

			client := testAccProvider.Meta().(*ArmClient).groupsClient
			ctx := testAccProvider.Meta().(*ArmClient).StopContext

			groupID := rs.Primary.Attributes["group_object_id"]
			ownerID := rs.Primary.Attributes["owner_object_id"]

			// see if group exists
			if resp, err := client.Get(ctx, groupID); err != nil {
				if ar.ResponseWasNotFound(resp.Response) {
					continue
				}

				return fmt.Errorf("Error retrieving Azure AD Group with ID %q: %+v", groupID, err)
			}

			owners, err := graph.GroupAllOwners(client, ctx, groupID)
			if err != nil {
				return fmt.Errorf("Error retrieving Azure AD Group owners (groupObjectId: %q): %+v", groupID, err)
			}

			var ownerObjectID string
			for _, objectID := range owners {
				if objectID == ownerID {
					ownerObjectID = objectID
				}
			}

			if ownerObjectID != "" {
				return fmt.Errorf("Azure AD group owner still exists:\n%#v", ownerObjectID)
			}
		}

		return nil
	}
}

func testAccAzureADGroupOwner_oneUser(id int, password string) string {
	return fmt.Sprintf(`
%[1]s
	
resource "azuread_group" "test" {
	name = "acctestGroup-%[2]d"
}

resource "azuread_group_owner" "testA" {
	group_object_id 	= "${azuread_group.test.object_id}"
	owner_object_id 	= "${azuread_user.testA.object_id}"
}

`, testAccADUser_threeUsersABC(id, password), id)
}

func testAccAzureADGroupOwner_twoUsers(id int, password string) string {
	return fmt.Sprintf(`
%[1]s
	
resource "azuread_group" "test" {
	name = "acctestGroup-%[2]d"
}

resource "azuread_group_owner" "testA" {
	group_object_id 	= "${azuread_group.test.object_id}"
	owner_object_id 	= "${azuread_user.testA.object_id}"
}

resource "azuread_group_owner" "testB" {
	group_object_id 	= "${azuread_group.test.object_id}"
	owner_object_id 	= "${azuread_user.testB.object_id}"
}

`, testAccADUser_threeUsersABC(id, password), id)
}

func testAccAzureADGroupOwner_ServicePrincipal(id int) string {
	return fmt.Sprintf(`

resource "azuread_application" "test" {
	name = "acctestApp-%[1]d"
}

resource "azuread_service_principal" "test" {
	application_id = "${azuread_application.test.application_id}"
}

resource "azuread_group" "test" {
	name = "acctestGroup-%[1]d"
}

resource "azuread_group_owner" "test" {
	group_object_id  = "${azuread_group.test.object_id}"
	owner_object_id = "${azuread_service_principal.test.object_id}"
}

`, id)
}
