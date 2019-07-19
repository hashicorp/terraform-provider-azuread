package azuread

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/tf"

	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/ar"
)

func TestAccAzureADGroup_basic(t *testing.T) {
	rn := "azuread_group.test"
	id := tf.AccRandTimeInt()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureADGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureADGroup_basic(id),
				Check:  testCheckAzureAdGroupBasic(id, "0", "0"),
			},
			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAzureADGroup_complete(t *testing.T) {
	rn := "azuread_group.test"
	id := tf.AccRandTimeInt()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureADGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureADGroup_basic(id),
				Check:  testCheckAzureAdGroupBasic(id, "0", "0"),
			},
			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAzureADGroup_owners(t *testing.T) {
	rn := "azuread_group.test"
	id := tf.AccRandTimeInt()
	pw := "p@$$wR2" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureADGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureADGroupWithThreeOwners(id, pw),
				Check:  testCheckAzureAdGroupBasic(id, "0", "3"),
			},
			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAzureADGroup_members(t *testing.T) {
	rn := "azuread_group.test"
	id := tf.AccRandTimeInt()
	pw := "p@$$wR2" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureADGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureADGroupWithThreeMembers(id, pw),
				Check:  testCheckAzureAdGroupBasic(id, "3", "0"),
			},
			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAzureADGroup_membersAndOwners(t *testing.T) {
	rn := "azuread_group.test"
	id := tf.AccRandTimeInt()
	pw := "p@$$wR2" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureADGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureADGroupWithOwnersAndMembers(id, pw),
				Check:  testCheckAzureAdGroupBasic(id, "2", "1"),
			},
			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAzureADGroup_membersDiverse(t *testing.T) {
	rn := "azuread_group.test"
	id := tf.AccRandTimeInt()
	pw := "p@$$wR2" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureADGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureADGroupWithDiverseMembers(id, pw),
				Check:  testCheckAzureAdGroupBasic(id, "3", "0"),
			},
			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAzureADGroup_membersUpdate(t *testing.T) {
	rn := "azuread_group.test"
	id := tf.AccRandTimeInt()
	pw := "p@$$wR2" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureADGroupDestroy,
		Steps: []resource.TestStep{
			// Empty group with 0 members
			{
				Config: testAccAzureADGroup_basic(id),
				Check:  testCheckAzureAdGroupBasic(id, "0", "0"),
			},
			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Group with 1 member
			{
				Config: testAccAzureADGroupWithOneMember(id, pw),
				Check:  testCheckAzureAdGroupBasic(id, "1", "0"),
			},
			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Group with multiple members
			{
				Config: testAccAzureADGroupWithThreeMembers(id, pw),
				Check:  testCheckAzureAdGroupBasic(id, "3", "0"),
			},
			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Group with a different member
			{
				Config: testAccAzureADGroupWithServicePrincipal(id),
				Check:  testCheckAzureAdGroupBasic(id, "1", "0"),
			},
			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Empty group with 0 members
			{
				Config: testAccAzureADGroup_basic(id),
				Check:  testCheckAzureAdGroupBasic(id, "0", "0"),
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

func testCheckAzureAdGroupBasic(id int, memberCount, ownerCount string) resource.TestCheckFunc {
	resourceName := "azuread_group.test"

	return resource.ComposeTestCheckFunc(
		testCheckAzureADGroupExists(resourceName),
		resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("acctestGroup-%d", id)),
		resource.TestCheckResourceAttrSet(resourceName, "object_id"),
		resource.TestCheckResourceAttr(resourceName, "members.#", memberCount),
		resource.TestCheckResourceAttr(resourceName, "owners.#", ownerCount),
	)
}

func testAccAzureADGroup_basic(id int) string {
	return fmt.Sprintf(`
resource "azuread_group" "test" {
  name    = "acctestGroup-%d"
  members = []
}
`, id)
}

func testAccAzureADGroupWithDiverseMembers(id int, password string) string {
	return fmt.Sprintf(`
data "azuread_domains" "tenant_domain" {
	only_initial = true
}

resource "azuread_application" "test" {
  name = "acctestApp-%[1]d"
}

resource "azuread_service_principal" "test" {
  application_id = azuread_application.test.application_id
}

resource "azuread_group" "member" {
  name   = "acctestGroup-%[1]d-Member"
}

resource "azuread_user" "test" {
	user_principal_name   = "acctestUser.%[1]d@${data.azuread_domains.tenant_domain.domains.0.domain_name}"
	display_name          = "acctestUser-%[1]d"
	password              = "%[2]s"
}

resource "azuread_group" "test" {
  name   = "acctestGroup-%[1]d"
  members = [ azuread_user.test.object_id, azuread_group.member.object_id, azuread_service_principal.test.object_id ]
}
`, id, password)
}

func testAccAzureADGroupWithOneMember(id int, password string) string {
	return fmt.Sprintf(`
data "azuread_domains" "tenant_domain" {
	only_initial = true
}

resource "azuread_user" "test" {
	user_principal_name   = "acctestUser.%[1]d@${data.azuread_domains.tenant_domain.domains.0.domain_name}"
	display_name          = "acctestUser-%[1]d"
	password              = "%[2]s"
}

resource "azuread_group" "test" {
  name   = "acctestGroup-%[1]d"
  members = [ azuread_user.test.object_id ]
}
`, id, password)
}

func testAccAzureADGroupWithThreeMembers(id int, password string) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_group" "test" {
  name    = "acctestGroup-%[2]d"
  members = [ azuread_user.testA.object_id, azuread_user.testB.object_id, azuread_user.testC.object_id ]
}
`, testAccADUser_threeUsersABC(id, password), id)
}

func testAccAzureADGroupWithThreeOwners(id int, password string) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_group" "test" {
  name   = "acctestGroup-%[2]d"
  owners = [ azuread_user.testA.object_id, azuread_user.testB.object_id, azuread_user.testC.object_id ]
}
`, testAccADUser_threeUsersABC(id, password), id)
}

func testAccAzureADGroupWithOwnersAndMembers(id int, password string) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_group" "test" {
  name    = "acctestGroup-%[2]d"
  owners  = [ azuread_user.testA.object_id ]
  members = [ azuread_user.testB.object_id, azuread_user.testC.object_id ]
}
`, testAccADUser_threeUsersABC(id, password), id)
}

func testAccAzureADGroupWithServicePrincipal(id int) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name = "acctestApp-%[1]d"
}

resource "azuread_service_principal" "test" {
  application_id = azuread_application.test.application_id
}

resource "azuread_group" "test" {
  name    = "acctestGroup-%[1]d"
  members = [ azuread_service_principal.test.object_id ]
}
`, id)
}
