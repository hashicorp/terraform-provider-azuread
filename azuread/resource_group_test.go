package azuread

import (
	"fmt"
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
				Check:  assertResourceWithMemberCount(id, "0"),
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

	config := testAccAzureADGroupWithThreeMembers(id)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureADGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check:  assertResourceWithMemberCount(id, "3"),
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
				Check:  assertResourceWithMemberCount(id, "0"),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAzureADGroup_diverse(t *testing.T) {
	resourceName := "azuread_group.test"
	id, err := uuid.GenerateUUID()
	if err != nil {
		t.Fatal(err)
	}
	config := testAccAzureADGroupWithDiverseMembers(id)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureADGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check:  assertResourceWithMemberCount(id, "3"),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAzureADGroup_progression(t *testing.T) {
	resourceName := "azuread_group.test"
	id, err := uuid.GenerateUUID()
	if err != nil {
		t.Fatal(err)
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureADGroupDestroy,
		Steps: []resource.TestStep{
			// Empty group with 0 members
			{
				Config: testAccAzureADGroup(id),
				Check:  assertResourceWithMemberCount(id, "0"),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Group with 1 member
			{
				Config: testAccAzureADGroupWithOneMember(id),
				Check:  assertResourceWithMemberCount(id, "1"),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Group with multiple members
			{
				Config: testAccAzureADGroupWithThreeMembers(id),
				Check:  assertResourceWithMemberCount(id, "3"),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Group with a different member
			{
				Config: testAccAzureADGroupWithServicePrincipal(id),
				Check:  assertResourceWithMemberCount(id, "1"),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Empty group with 0 members
			{
				Config: testAccAzureADGroup(id),
				Check:  assertResourceWithMemberCount(id, "0"),
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

func assertResourceWithMemberCount(id string, memberCount string) resource.TestCheckFunc {
	resourceName := "azuread_group.test"

	return resource.ComposeTestCheckFunc(
		testCheckAzureADGroupExists(resourceName),
		resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("acctest%s", id)),
		resource.TestCheckResourceAttrSet(resourceName, "object_id"),
		resource.TestCheckResourceAttr(resourceName, "members.#", memberCount),
	)
}

func testAccAzureADGroup(id string) string {
	return fmt.Sprintf(`
resource "azuread_group" "test" {
  name = "acctest%s"
  members = []
}
`, id)
}

func testAccAzureADGroupWithDiverseMembers(id string) string {
	return fmt.Sprintf(`
data "azuread_domains" "tenant_domain" {
	only_initial = true
}

resource "azuread_application" "test_app_%[1]s" {
  name = "app%[1]s"
}

resource "azuread_service_principal" "test_sp_%[1]s" {
  application_id = azuread_application.test_app_%[1]s.application_id
}

resource "azuread_group" "test_g_%[1]s" {
  name   = "acctest%[1]s"
}

resource "azuread_user" "acctest_user_%[1]s" {
	user_principal_name   = "acctest.%[1]s@${data.azuread_domains.tenant_domain.domains.0.domain_name}"
	display_name          = "acctest-%[1]s"
	password              = "%[1]s"
}

resource "azuread_group" "test" {
  name   = "acctest%[1]s"
  members = [ azuread_user.acctest_user_%[1]s.object_id, azuread_group.test_g_%[1]s.object_id, azuread_service_principal.test_sp_%[1]s.object_id ]
}
`, id)
}

func testAccAzureADGroupWithOneMember(id string) string {
	return fmt.Sprintf(`
data "azuread_domains" "tenant_domain" {
	only_initial = true
}

resource "azuread_user" "acctest_user_%[1]s" {
	user_principal_name   = "acctest.%[1]s@${data.azuread_domains.tenant_domain.domains.0.domain_name}"
	display_name          = "acctest-%[1]s"
	password              = "%[1]s"
}

resource "azuread_group" "test" {
  name   = "acctest%[1]s"
  members = [ azuread_user.acctest_user_%[1]s.object_id ]
}
`, id)
}

func testAccAzureADGroupWithThreeMembers(id string) string {
	return fmt.Sprintf(`
data "azuread_domains" "tenant_domain" {
	only_initial = true
}

resource "azuread_user" "acctest_user_%[2]s" {
	user_principal_name   = "acctest.%[2]s@${data.azuread_domains.tenant_domain.domains.0.domain_name}"
	display_name          = "acctest-%[2]s"
	password              = "%[2]s"
}

resource "azuread_user" "acctest_user_%[3]s" {
	user_principal_name   = "acctest.%[3]s@${data.azuread_domains.tenant_domain.domains.0.domain_name}"
	display_name          = "acctest-%[3]s"
	password              = "%[3]s"
}

resource "azuread_user" "acctest_user_%[4]s" {
	user_principal_name   = "acctest.%[4]s@${data.azuread_domains.tenant_domain.domains.0.domain_name}"
	display_name          = "acctest-%[4]s"
	password              = "%[4]s"
}

resource "azuread_group" "test" {
  name   = "acctest%[1]s"
  members = [ azuread_user.acctest_user_%[2]s.object_id, azuread_user.acctest_user_%[3]s.object_id, azuread_user.acctest_user_%[4]s.object_id ]
}
`, id, id+"a", id+"b", id+"c")
}

func testAccAzureADGroupWithServicePrincipal(id string) string {
	return fmt.Sprintf(`
resource "azuread_application" "test_app_%[1]s" {
  name = "app%[1]s"
}

resource "azuread_service_principal" "test_sp_%[1]s" {
  application_id = azuread_application.test_app_%[1]s.application_id
}

resource "azuread_group" "test" {
  name   = "acctest%[1]s"
  members = [ azuread_service_principal.test_sp_%[1]s.object_id ]
}
`, id)
}
