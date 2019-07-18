package azuread

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/ar"
)

func TestAccAzureADGroupMember_User(t *testing.T) {
	resourceName := "azuread_group_member.test"
	id := acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)
	password := id + "p@$$wR2"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureADGroupMemberDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureADGroupMember_User(id, password),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "group_object_id"),
					resource.TestCheckResourceAttrSet(resourceName, "member_object_id"),
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

func TestAccAzureADGroupMember_Group(t *testing.T) {
	resourceName := "azuread_group_member.test"
	id := acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureADGroupMemberDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureADGroupMember_Group(id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "group_object_id"),
					resource.TestCheckResourceAttrSet(resourceName, "member_object_id"),
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

func TestAccAzureADGroupMember_ServicePrincipal(t *testing.T) {
	resourceName := "azuread_group_member.test"
	id := acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureADGroupMemberDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureADGroupMember_ServicePrincipal(id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "group_object_id"),
					resource.TestCheckResourceAttrSet(resourceName, "member_object_id"),
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

func testCheckAzureADGroupMemberDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "azuread_group_member" {
			continue
		}

		client := testAccProvider.Meta().(*ArmClient).groupsClient
		ctx := testAccProvider.Meta().(*ArmClient).StopContext

		groupID := rs.Primary.Attributes["group_object_id"]
		memberID := rs.Primary.Attributes["member_object_id"]

		members, err := client.GetGroupMembersComplete(ctx, groupID)
		if err != nil {
			if ar.ResponseWasNotFound(members.Response().Response) {
				return nil
			}

			return err
		}

		var memberObjectID string
		for members.NotDone() {
			// possible members are users, groups or service principals
			// we try to 'cast' each result as the corresponding type and diff
			// if we found the object we're looking for
			user, _ := members.Value().AsUser()
			if user != nil {
				if *user.ObjectID == memberID {
					memberObjectID = *user.ObjectID
					// we successfully found the directory object we're looking for, we can stop looping
					// through the results
					break
				}
			}

			group, _ := members.Value().AsADGroup()
			if group != nil {
				if *group.ObjectID == memberID {
					memberObjectID = *group.ObjectID
					// we successfully found the directory object we're looking for, we can stop looping
					// through the results
					break
				}
			}

			servicePrincipal, _ := members.Value().AsServicePrincipal()
			if servicePrincipal != nil {
				if *servicePrincipal.ObjectID == memberID {
					memberObjectID = *servicePrincipal.ObjectID
					// we successfully found the directory object we're looking for, we can stop looping
					// through the results
					break
				}
			}

			err = members.NextWithContext(ctx)
			if err != nil {
				return fmt.Errorf("Error listing Azure AD Group Members: %s", err)
			}
		}

		if memberObjectID != "" {
			return fmt.Errorf("Azure AD group member still exists:\n%#v", memberObjectID)
		}
	}

	return nil
}

func testAccAzureADGroupMember_User(id string, password string) string {
	return fmt.Sprintf(`

data "azuread_domains" "tenant_domain" {
	only_initial = true
}

resource "azuread_user" "test" {
	user_principal_name   = "acctestUser.%[1]s.A@${data.azuread_domains.tenant_domain.domains.0.domain_name}"
	display_name          = "acctestUser-%[1]s-A"
	password              = "%[2]s"
}
	
resource "azuread_group" "test" {
	name = "acctestGroup-%[1]s"
}

resource "azuread_group_member" "test" {
	group_object_id 	= "${azuread_group.test.object_id}"
	member_object_id 	= "${azuread_user.test.object_id}"
}

`, id, password)
}

func testAccAzureADGroupMember_Group(id string) string {
	return fmt.Sprintf(`
	
resource "azuread_group" "test" {
	name = "acctestGroup%[1]s"
}

resource "azuread_group" "member" {
	name = "acctestGroup-%[1]s-Member"
}

resource "azuread_group_member" "test" {
	group_object_id 	= "${azuread_group.test.object_id}"
	member_object_id 	= "${azuread_group.member.object_id}"
}

`, id)
}

func testAccAzureADGroupMember_ServicePrincipal(id string) string {
	return fmt.Sprintf(`

resource "azuread_application" "test" {
	name = "acctestApp%[1]s"
}

resource "azuread_service_principal" "test" {
	application_id = "${azuread_application.test.application_id}"
}

resource "azuread_group" "test" {
	name = "acctestGroup-%[1]s"
}

resource "azuread_group_member" "test" {
	group_object_id  = "${azuread_group.test.object_id}"
	member_object_id = "${azuread_service_principal.test.object_id}"
}

`, id)
}
