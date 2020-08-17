package aadgraph_test

import (
	"fmt"
	"github.com/terraform-providers/terraform-provider-azuread/internal/utils"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-azuread/internal/acceptance"
	"github.com/terraform-providers/terraform-provider-azuread/internal/clients"
	"github.com/terraform-providers/terraform-provider-azuread/internal/tf"
)

func TestAccGroup_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_group", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGroup_basic(data.RandomInteger),
				Check:  testCheckGroupBasic(data.RandomInteger, "0", "0"),
			},
			data.ImportStep(),
		},
	})
}

func TestAccGroup_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_group", "test")
	pw := "utils@$$wR2" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGroup_complete(data.RandomInteger, pw),
				Check:  testCheckGroupBasic(data.RandomInteger, "1", "1"),
			},
			data.ImportStep(),
		},
	})
}

func TestAccGroup_owners(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_group", "test")
	pw := "utils@$$wR2" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGroupWithThreeOwners(data.RandomInteger, pw),
				Check:  testCheckGroupBasic(data.RandomInteger, "0", "3"),
			},
			data.ImportStep(),
		},
	})
}

func TestAccGroup_members(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_group", "test")
	pw := "utils@$$wR2" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGroupWithThreeMembers(data.RandomInteger, pw),
				Check:  testCheckGroupBasic(data.RandomInteger, "3", "0"),
			},
			data.ImportStep(),
		},
	})
}

func TestAccGroup_membersAndOwners(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_group", "test")
	pw := "utils@$$wR2" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGroupWithOwnersAndMembers(data.RandomInteger, pw),
				Check:  testCheckGroupBasic(data.RandomInteger, "2", "1"),
			},
			data.ImportStep(),
		},
	})
}

func TestAccGroup_membersDiverse(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_group", "test")
	pw := "utils@$$wR2" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGroupWithDiverseMembers(data.RandomInteger, pw),
				Check:  testCheckGroupBasic(data.RandomInteger, "3", "0"),
			},
			data.ImportStep(),
		},
	})
}

func TestAccGroup_ownersDiverse(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_group", "test")
	pw := "utils@$$wR2" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGroupWithDiverseOwners(data.RandomInteger, pw),
				Check:  testCheckGroupBasic(data.RandomInteger, "0", "2"),
			},
			data.ImportStep(),
		},
	})
}

func TestAccGroup_membersUpdate(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_group", "test")
	pw := "utils@$$wR2" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckGroupDestroy,
		Steps: []resource.TestStep{
			// Empty group with 0 members
			{
				Config: testAccGroup_basic(data.RandomInteger),
				Check:  testCheckGroupBasic(data.RandomInteger, "0", "0"),
			},
			data.ImportStep(),
			// Group with 1 member
			{
				Config: testAccGroupWithOneMember(data.RandomInteger, pw),
				Check:  testCheckGroupBasic(data.RandomInteger, "1", "0"),
			},
			data.ImportStep(),
			// Group with multiple members
			{
				Config: testAccGroupWithThreeMembers(data.RandomInteger, pw),
				Check:  testCheckGroupBasic(data.RandomInteger, "3", "0"),
			},
			data.ImportStep(),
			// Group with a different member
			{
				Config: testAccGroupWithServicePrincipalMember(data.RandomInteger),
				Check:  testCheckGroupBasic(data.RandomInteger, "1", "0"),
			},
			data.ImportStep(),
			// Empty group with 0 members
			{
				Config: testAccGroup_basic(data.RandomInteger),
				Check:  testCheckGroupBasic(data.RandomInteger, "0", "0"),
			},
			data.ImportStep(),
		},
	})
}

func TestAccGroup_ownersUpdate(t *testing.T) {
	rn := "azuread_group.test"
	id := tf.AccRandTimeInt()
	pw := "utils@$$wR2" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckGroupDestroy,
		Steps: []resource.TestStep{
			// Empty group with 0 owners
			{
				Config: testAccGroup_basic(id),
				Check:  testCheckGroupBasic(id, "0", "0"),
			},
			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Group with multiple owners
			{
				Config: testAccGroupWithThreeOwners(id, pw),
				Check:  testCheckGroupBasic(id, "0", "3"),
			},
			// Group with 1 owners
			{
				Config: testAccGroupWithOneOwners(id, pw),
				Check:  testCheckGroupBasic(id, "0", "1"),
			},
			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Group with a different owners
			{
				Config: testAccGroupWithServicePrincipalOwner(id),
				Check:  testCheckGroupBasic(id, "0", "1"),
			},
			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Empty group with 0 owners is not possible
		},
	})
}

func TestAccGroup_preventDuplicateNames(t *testing.T) {
	ri := tf.AccRandTimeInt()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccGroup_duplicateName(ri),
				ExpectError: regexp.MustCompile("existing Azure Active Directory Group .+ was found"),
			},
		},
	})
}

func testCheckGroupExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %q", name)
		}

		client := acceptance.AzureADProvider.Meta().(*clients.AadClient).AadGraph.GroupsClient
		ctx := acceptance.AzureADProvider.Meta().(*clients.AadClient).StopContext
		resp, err := client.Get(ctx, rs.Primary.ID)

		if err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("Bad: Azure AD Group %q does not exist", rs.Primary.ID)
			}
			return fmt.Errorf("Bad: Get on Azure AD GroupsClient: %+v", err)
		}

		return nil
	}
}

func testCheckGroupDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "azuread_group" {
			continue
		}

		client := acceptance.AzureADProvider.Meta().(*clients.AadClient).AadGraph.GroupsClient
		ctx := acceptance.AzureADProvider.Meta().(*clients.AadClient).StopContext
		resp, err := client.Get(ctx, rs.Primary.ID)

		if err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
				return nil
			}

			return err
		}

		return fmt.Errorf("Azure AD group still exists:\n%#v", resp)
	}

	return nil
}

func testCheckGroupBasic(id int, memberCount, ownerCount string) resource.TestCheckFunc {
	resourceName := "azuread_group.test"

	return resource.ComposeTestCheckFunc(
		testCheckGroupExists(resourceName),
		resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("acctestGroup-%d", id)),
		resource.TestCheckResourceAttrSet(resourceName, "object_id"),
		resource.TestCheckResourceAttr(resourceName, "members.#", memberCount),
		resource.TestCheckResourceAttr(resourceName, "owners.#", ownerCount),
	)
}

func testAccGroup_basic(id int) string {
	return fmt.Sprintf(`
resource "azuread_group" "test" {
  name    = "acctestGroup-%d"
  members = []
}
`, id)
}

func testAccGroup_complete(id int, password string) string {
	return fmt.Sprintf(`
%s

resource "azuread_group" "test" {
  name        = "acctestGroup-%d"
  description = "Please delete me as this is a.test.AD group!"
  members     = [azuread_user.test.object_id]
  owners      = [azuread_user.test.object_id]
}
`, testAccUser_basic(id, password), id)
}

func testAccDiverseDirectoryObjects(id int, password string) string {
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
  name = "acctestGroup-%[1]d-Member"
}

resource "azuread_user" "test" {
  user_principal_name = "acctestUser.%[1]d@${data.azuread_domains.tenant_domain.domains.0.domain_name}"
  display_name        = "acctestUser-%[1]d"
  password            = "%[2]s"
}
`, id, password)
}

func testAccGroupWithDiverseMembers(id int, password string) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_group" "test" {
  name    = "acctestGroup-%[2]d"
  members = [azuread_user.test.object_id, azuread_group.member.object_id, azuread_service_principal.test.object_id]
}
`, testAccDiverseDirectoryObjects(id, password), id)
}

func testAccGroupWithDiverseOwners(id int, password string) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_group" "test" {
  name   = "acctestGroup-%[2]d"
  owners = [azuread_user.test.object_id, azuread_service_principal.test.object_id]
}
`, testAccDiverseDirectoryObjects(id, password), id)
}

func testAccGroupWithOneMember(id int, password string) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_group" "test" {
  name    = "acctestGroup-%[2]d"
  members = [azuread_user.test.object_id]
}
`, testAccUser_basic(id, password), id)
}

func testAccGroupWithOneOwners(id int, password string) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_group" "test" {
  name   = "acctestGroup-%[2]d"
  owners = [azuread_user.test.object_id]
}
`, testAccUser_basic(id, password), id)
}

func testAccGroupWithThreeMembers(id int, password string) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_group" "test" {
  name    = "acctestGroup-%[2]d"
  members = [azuread_user.testA.object_id, azuread_user.testB.object_id, azuread_user.testC.object_id]
}
`, testAccUser_threeUsersABC(id, password), id)
}

func testAccGroupWithThreeOwners(id int, password string) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_group" "test" {
  name   = "acctestGroup-%[2]d"
  owners = [azuread_user.testA.object_id, azuread_user.testB.object_id, azuread_user.testC.object_id]
}
`, testAccUser_threeUsersABC(id, password), id)
}

func testAccGroupWithOwnersAndMembers(id int, password string) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_group" "test" {
  name    = "acctestGroup-%[2]d"
  owners  = [azuread_user.testA.object_id]
  members = [azuread_user.testB.object_id, azuread_user.testC.object_id]
}
`, testAccUser_threeUsersABC(id, password), id)
}

func testAccGroupWithServicePrincipalMember(id int) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name = "acctestApp-%[1]d"
}

resource "azuread_service_principal" "test" {
  application_id = azuread_application.test.application_id
}

resource "azuread_group" "test" {
  name    = "acctestGroup-%[1]d"
  members = [azuread_service_principal.test.object_id]
}
`, id)
}

func testAccGroupWithServicePrincipalOwner(id int) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name = "acctestApp-%[1]d"
}

resource "azuread_service_principal" "test" {
  application_id = azuread_application.test.application_id
}

resource "azuread_group" "test" {
  name   = "acctestGroup-%[1]d"
  owners = [azuread_service_principal.test.object_id]
}
`, id)
}

func testAccGroup_duplicateName(id int) string {
	return fmt.Sprintf(`
%s

resource "azuread_group" "duplicate" {
  name                    = azuread_group.test.name
  prevent_duplicate_names = true
}
`, testAccGroup_basic(id))
}
