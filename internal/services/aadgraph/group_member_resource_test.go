package aadgraph_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-azuread/internal/acceptance"
	"github.com/terraform-providers/terraform-provider-azuread/internal/clients"
	"github.com/terraform-providers/terraform-provider-azuread/internal/services/aadgraph/graph"
	"github.com/terraform-providers/terraform-provider-azuread/internal/utils"
)

func TestAccGroupMember_group(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_group_member", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckGroupMemberDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGroupMember_group(data.RandomInteger),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(data.ResourceName, "group_object_id"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "member_object_id"),
				),
			},
			data.ImportStep(),
		},
	})
}

func TestAccGroupMember_servicePrincipal(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_group_member", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckGroupMemberDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGroupMember_servicePrincipal(data.RandomInteger),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(data.ResourceName, "group_object_id"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "member_object_id"),
				),
			},
			data.ImportStep(),
		},
	})
}

func TestAccGroupMember_user(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_group_member", "testA")
	pw := "utils@$$wR2" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckGroupMemberDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGroupMember_oneUser(data.RandomInteger, pw),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(data.ResourceName, "group_object_id"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "member_object_id"),
				),
			},
			data.ImportStep(),
		},
	})
}

func TestAccGroupMember_multipleUser(t *testing.T) {
	dataA := acceptance.BuildTestData(t, "azuread_group_member", "testA")
	dataB := acceptance.BuildTestData(t, "azuread_group_member", "testA")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckGroupMemberDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGroupMember_oneUser(dataA.RandomInteger, dataA.RandomPassword),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataA.ResourceName, "group_object_id"),
					resource.TestCheckResourceAttrSet(dataA.ResourceName, "member_object_id"),
				),
			},
			dataA.ImportStep(),
			{
				Config: testAccGroupMember_twoUsers(dataA.RandomInteger, dataA.RandomPassword),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataA.ResourceName, "group_object_id"),
					resource.TestCheckResourceAttrSet(dataA.ResourceName, "member_object_id"),
					resource.TestCheckResourceAttrSet(dataB.ResourceName, "group_object_id"),
					resource.TestCheckResourceAttrSet(dataB.ResourceName, "member_object_id"),
				),
			},
			// we rerun the config so the group resource updates with the number of members
			{
				Config: testAccGroupMember_twoUsers(dataA.RandomInteger, dataA.RandomPassword),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("azuread_group.test", "members.#", "2"),
				),
			},
			dataA.ImportStep(),
			{
				Config: testAccGroupMember_oneUser(dataA.RandomInteger, dataA.RandomPassword),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataA.ResourceName, "group_object_id"),
					resource.TestCheckResourceAttrSet(dataA.ResourceName, "member_object_id"),
				),
			},
			// we rerun the config so the group resource updates with the number of members
			{
				Config: testAccGroupMember_oneUser(dataA.RandomInteger, dataA.RandomPassword),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("azuread_group.test", "members.#", "1"),
				),
			},
		},
	})
}

func testCheckGroupMemberDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "azuread_group_member" {
			continue
		}

		client := acceptance.AzureADProvider.Meta().(*clients.AadClient).AadGraph.GroupsClient
		ctx := acceptance.AzureADProvider.Meta().(*clients.AadClient).StopContext

		groupID := rs.Primary.Attributes["group_object_id"]
		memberID := rs.Primary.Attributes["member_object_id"]

		// see if group exists
		if resp, err := client.Get(ctx, groupID); err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
				continue
			}

			return fmt.Errorf("retrieving Group with ID %q: %+v", groupID, err)
		}

		members, err := graph.GroupAllMembers(ctx, client, groupID)
		if err != nil {
			return fmt.Errorf("retrieving Group members (groupObjectId: %q): %+v", groupID, err)
		}

		var memberObjectID string
		for _, objectID := range members {
			if objectID == memberID {
				memberObjectID = objectID
			}
		}

		if memberObjectID != "" {
			return fmt.Errorf("Group member still exists:\n%#v", memberObjectID)
		}
	}

	return nil
}

func testAccGroupMember_group(id int) string {
	return fmt.Sprintf(`

resource "azuread_group" "test" {
  name = "acctestGroup-%[1]d"
}

resource "azuread_group" "member" {
  name = "acctestGroup-%[1]d-Member"
}

resource "azuread_group_member" "test" {
  group_object_id  = azuread_group.test.object_id
  member_object_id = azuread_group.member.object_id
}

`, id)
}

func testAccGroupMember_servicePrincipal(id int) string {
	return fmt.Sprintf(`

resource "azuread_application" "test" {
  name = "acctestApp-%[1]d"
}

resource "azuread_service_principal" "test" {
  application_id = azuread_application.test.application_id
}

resource "azuread_group" "test" {
  name = "acctestGroup-%[1]d"
}

resource "azuread_group_member" "test" {
  group_object_id  = azuread_group.test.object_id
  member_object_id = azuread_service_principal.test.object_id
}

`, id)
}

func testAccGroupMember_oneUser(id int, password string) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_group" "test" {
  name = "acctestGroup-%[2]d"
}

resource "azuread_group_member" "testA" {
  group_object_id  = azuread_group.test.object_id
  member_object_id = azuread_user.testA.object_id
}

`, testAccUser_threeUsersABC(id, password), id)
}

func testAccGroupMember_twoUsers(id int, password string) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_group" "test" {
  name = "acctestGroup-%[2]d"
}

resource "azuread_group_member" "testA" {
  group_object_id  = azuread_group.test.object_id
  member_object_id = azuread_user.testA.object_id
}

resource "azuread_group_member" "testB" {
  group_object_id  = azuread_group.test.object_id
  member_object_id = azuread_user.testB.object_id
}

`, testAccUser_threeUsersABC(id, password), id)
}
