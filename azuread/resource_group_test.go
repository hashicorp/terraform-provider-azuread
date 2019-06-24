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
				Check:  assertResourceWithMemberCount(id, resourceName, "0"),
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
		memberUuid, _ := uuid.GenerateUUID()
		members = append(members, memberUuid)
	}

	config := testAccAzureADGroupWithMembers(id, members)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureADGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check:  assertResourceWithMemberCount(id, resourceName, "5"),
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
				Check:  assertResourceWithMemberCount(id, resourceName, "0"),
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
				Check:  assertResourceWithMemberCount(id, resourceName, "3"),
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
				Config: groupWithMembers(id, ""),
				Check:  assertResourceWithMemberCount(id, resourceName, "0"),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Group with 1 member
			{
				Config: user(id) + groupWithMembers(id, fmt.Sprintf("azuread_user.acctest_user_%[1]s.id", id)),
				Check:  assertResourceWithMemberCount(id, resourceName, "1"),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Group with multiple members
			{
				Config: user(id+"a") + user(id+"b") + user(id+"c") + groupWithMembers(id, fmt.Sprintf("azuread_user.acctest_user_%[1]s.id, azuread_user.acctest_user_%[2]s.id, azuread_user.acctest_user_%[3]s.id", id+"a", id+"b", id+"c")),
				Check:  assertResourceWithMemberCount(id, resourceName, "3"),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Group with a different member
			{
				Config: servicePrincipal(id) + groupWithMembers(id, fmt.Sprintf("azuread_service_principal.test_sp_%[1]s.id", id)),
				Check:  assertResourceWithMemberCount(id, resourceName, "1"),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Empty group with 0 members
			{
				Config: groupWithMembers(id, ""),
				Check:  assertResourceWithMemberCount(id, resourceName, "0"),
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

func assertResourceWithMemberCount(id string, resourceName string, memberCount string) resource.TestCheckFunc {
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
}
`, id)
}

func testAccAzureADGroupWithDiverseMembers(id string) string {
	var sb strings.Builder

	sb.WriteString(servicePrincipal(id))
	sb.WriteString(group(id))
	sb.WriteString(user(id))
	sb.WriteString(groupWithMembers(id, fmt.Sprintf("azuread_user.acctest_user_%[1]s.id, azuread_group.test_g_%[1]s.id, azuread_service_principal.test_sp_%[1]s.id", id)))

	return sb.String()
}

func testAccAzureADGroupWithMembers(id string, members []string) string {
	var sb strings.Builder

	sb.WriteString(`
data "azuread_domains" "tenant_domain" {
	only_initial = true
}`)

	for _, member := range members {
		fmt.Fprintf(&sb, `
resource "azuread_user" "acctest_user_%[1]s" {
	user_principal_name   = "acctest.%[1]s@${data.azuread_domains.tenant_domain.domains.0.domain_name}"
	display_name          = "acctest-%[1]s"
	password              = "%[1]s"
}
`, member)
	}

	fmt.Fprintf(&sb, `
resource "azuread_group" "test" {
  name   = "acctest%s"
  members = [ %s ]
}
`, id, strings.Join(formatMembersAsUser(members), ", "))

	return sb.String()
}

func formatMembersAsUser(members []string) []string {
	vsm := make([]string, len(members))
	for i, v := range members {
		vsm[i] = "azuread_user.acctest_user_" + v + ".id"
	}
	return vsm
}

func servicePrincipal(id string) string {
	return fmt.Sprintf(`
resource "azuread_application" "test_app_%[1]s" {
  name = "app%[1]s"
}

resource "azuread_service_principal" "test_sp_%[1]s" {
  application_id = azuread_application.test_app_%[1]s.application_id
}

`, id)
}

func group(id string) string {
	return fmt.Sprintf(`
resource "azuread_group" "test_g_%[1]s" {
  name   = "acctest%[1]s"
}
`, id)
}

func groupWithMembers(id string, hclMemberString string) string {
	return fmt.Sprintf(`
data "azuread_domains" "tenant_domain" {
	only_initial = true
}

resource "azuread_group" "test" {
  name   = "acctest%[1]s"
  members = [ %[2]s ]
}
`, id, hclMemberString)
}

func user(id string) string {
	return fmt.Sprintf(`
resource "azuread_user" "acctest_user_%[1]s" {
	user_principal_name   = "acctest.%[1]s@${data.azuread_domains.tenant_domain.domains.0.domain_name}"
	display_name          = "acctest-%[1]s"
	password              = "%[1]s"
}
`, id)
}
