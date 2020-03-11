package azuread

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/ar"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/graph"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/tf"
)

func TestAccAzureADGroupMember_user(t *testing.T) {
	rn := "azuread_group_member.testA"
	id := tf.AccRandTimeInt()
	pw := "p@$$wR2" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureADGroupMemberDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureADGroupMember_oneUser(id, pw),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(rn, "group_object_id"),
					resource.TestCheckResourceAttrSet(rn, "member_object_id"),
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

func TestAccAzureADGroupMember_multipleUser(t *testing.T) {
	rna := "azuread_group_member.testA"
	rnb := "azuread_group_member.testB"
	id := tf.AccRandTimeInt()
	pw := "p@$$wR2" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureADGroupMemberDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureADGroupMember_oneUser(id, pw),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(rna, "group_object_id"),
					resource.TestCheckResourceAttrSet(rna, "member_object_id"),
				),
			},
			{
				ResourceName:      rna,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccAzureADGroupMember_twoUsers(id, pw),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(rna, "group_object_id"),
					resource.TestCheckResourceAttrSet(rna, "member_object_id"),
					resource.TestCheckResourceAttrSet(rnb, "group_object_id"),
					resource.TestCheckResourceAttrSet(rnb, "member_object_id"),
				),
			},
			// we rerun the config so the group resource updates with the number of members
			{
				Config: testAccAzureADGroupMember_twoUsers(id, pw),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("azuread_group.test", "members.#", "2"),
				),
			},
			{
				ResourceName:      rna,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccAzureADGroupMember_oneUser(id, pw),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(rna, "group_object_id"),
					resource.TestCheckResourceAttrSet(rna, "member_object_id"),
				),
			},
			{
				Config: testAccAzureADGroupMember_oneUser(id, pw),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("azuread_group.test", "members.#", "1"),
				),
			},
		},
	})
}

func TestAccAzureADGroupMember_group(t *testing.T) {
	rn := "azuread_group_member.test"
	id := tf.AccRandTimeInt()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureADGroupMemberDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureADGroupMember_group(id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(rn, "group_object_id"),
					resource.TestCheckResourceAttrSet(rn, "member_object_id"),
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

func TestAccAzureADGroupMember_servicePrincipal(t *testing.T) {
	rn := "azuread_group_member.test"
	id := tf.AccRandTimeInt()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureADGroupMemberDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureADGroupMember_servicePrincipal(id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(rn, "group_object_id"),
					resource.TestCheckResourceAttrSet(rn, "member_object_id"),
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

func testCheckAzureADGroupMemberDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "azuread_group_member" {
			continue
		}

		client := testAccProvider.Meta().(*ArmClient).groupsClient
		ctx := testAccProvider.Meta().(*ArmClient).StopContext

		groupID := rs.Primary.Attributes["group_object_id"]
		memberID := rs.Primary.Attributes["member_object_id"]

		// see if group exists
		if resp, err := client.Get(ctx, groupID); err != nil {
			if ar.ResponseWasNotFound(resp.Response) {
				continue
			}

			return fmt.Errorf("Error retrieving Azure AD Group with ID %q: %+v", groupID, err)
		}

		members, err := graph.GroupAllMembers(client, ctx, groupID)
		if err != nil {
			return fmt.Errorf("Error retrieving Azure AD Group members (groupObjectId: %q): %+v", groupID, err)
		}

		var memberObjectID string
		for _, objectID := range members {
			if objectID == memberID {
				memberObjectID = objectID
			}
		}

		if memberObjectID != "" {
			return fmt.Errorf("Azure AD group member still exists:\n%#v", memberObjectID)
		}
	}

	return nil
}

func testAccAzureADGroupMember_oneUser(id int, password string) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_group" "test" {
  name = "acctestGroup-%[2]d"
}

resource "azuread_group_member" "testA" {
  group_object_id  = "${azuread_group.test.object_id}"
  member_object_id = "${azuread_user.testA.object_id}"
}

`, testAccADUser_threeUsersABC(id, password), id)
}

func testAccAzureADGroupMember_twoUsers(id int, password string) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_group" "test" {
  name = "acctestGroup-%[2]d"
}

resource "azuread_group_member" "testA" {
  group_object_id  = "${azuread_group.test.object_id}"
  member_object_id = "${azuread_user.testA.object_id}"
}

resource "azuread_group_member" "testB" {
  group_object_id  = "${azuread_group.test.object_id}"
  member_object_id = "${azuread_user.testB.object_id}"
}

`, testAccADUser_threeUsersABC(id, password), id)
}

func testAccAzureADGroupMember_group(id int) string {
	return fmt.Sprintf(`

resource "azuread_group" "test" {
  name = "acctestGroup-%[1]d"
}

resource "azuread_group" "member" {
  name = "acctestGroup-%[1]d-Member"
}

resource "azuread_group_member" "test" {
  group_object_id  = "${azuread_group.test.object_id}"
  member_object_id = "${azuread_group.member.object_id}"
}

`, id)
}

func testAccAzureADGroupMember_servicePrincipal(id int) string {
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

resource "azuread_group_member" "test" {
  group_object_id  = "${azuread_group.test.object_id}"
  member_object_id = "${azuread_service_principal.test.object_id}"
}

`, id)
}
