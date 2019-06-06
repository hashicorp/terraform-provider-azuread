package azuread

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/ar"
)

func TestAccAzureADUser_basic(t *testing.T) {
	resourceName := "azuread_user.test"
	id := acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)
	password := id + "p@$$wR2"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADUser_basic(id, password),
				Check: resource.ComposeTestCheckFunc(
					testCheckADUserExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "user_principal_name"),
					resource.TestCheckResourceAttrSet(resourceName, "object_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", fmt.Sprintf("acctest%s", id)),
					resource.TestCheckResourceAttr(resourceName, "mail_nickname", fmt.Sprintf("acctest%s", id)),
					resource.TestCheckResourceAttr(resourceName, "account_enabled", "true"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"force_password_change",
					"password", // not returned from API, sensitive
				},
			},
		},
	})
}

func TestAccAzureADUser_complete(t *testing.T) {
	resourceName := "azuread_user.test"
	id := acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)
	password := id + "p@$$wR2"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADUser_complete(id, password),
				Check: resource.ComposeTestCheckFunc(
					testCheckADUserExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "user_principal_name"),
					resource.TestCheckResourceAttrSet(resourceName, "object_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", fmt.Sprintf("acctestupdate%s", id)),
					resource.TestCheckResourceAttr(resourceName, "mail_nickname", fmt.Sprintf("acctestupdate%s", id)),
					resource.TestCheckResourceAttr(resourceName, "account_enabled", "false"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"force_password_change",
					"password", // not returned from API, sensitive
				},
			},
		},
	})
}

func TestAccAzureADUser_update(t *testing.T) {
	resourceName := "azuread_user.test"
	id := acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)
	password := id + "p@$$wRd"
	updatedPassword := id + "p@$$wRd2"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADUser_basic(id, password),
				Check: resource.ComposeTestCheckFunc(
					testCheckADUserExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "user_principal_name"),
					resource.TestCheckResourceAttrSet(resourceName, "object_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", fmt.Sprintf("acctest%s", id)),
					resource.TestCheckResourceAttr(resourceName, "mail_nickname", fmt.Sprintf("acctest%s", id)),
					resource.TestCheckResourceAttr(resourceName, "account_enabled", "true"),
				),
			},
			{
				Config: testAccADUser_complete(id, updatedPassword),
				Check: resource.ComposeTestCheckFunc(
					testCheckADUserExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "user_principal_name"),
					resource.TestCheckResourceAttrSet(resourceName, "object_id"),
					resource.TestCheckResourceAttr(resourceName, "display_name", fmt.Sprintf("acctestupdate%s", id)),
					resource.TestCheckResourceAttr(resourceName, "mail_nickname", fmt.Sprintf("acctestupdate%s", id)),
					resource.TestCheckResourceAttr(resourceName, "account_enabled", "false"),
				),
			},
			{
				Config: testAccADUser_multiple(id, password),
				Check: resource.ComposeTestCheckFunc(
					testCheckADUserExists("azuread_user.testA"),
					testCheckADUserExists("azuread_user.testB"),
					resource.TestCheckResourceAttrSet("azuread_user.testA", "user_principal_name"),
					resource.TestCheckResourceAttr("azuread_user.testA", "display_name", fmt.Sprintf("acctestA%s", id)),
					resource.TestCheckResourceAttr("azuread_user.testA", "mail_nickname", fmt.Sprintf("acctestA%s", id)),
					resource.TestCheckResourceAttrSet("azuread_user.testB", "user_principal_name"),
					resource.TestCheckResourceAttr("azuread_user.testB", "display_name", fmt.Sprintf("acctest_display%s", id)),
					resource.TestCheckResourceAttr("azuread_user.testB", "mail_nickname", fmt.Sprintf("acctest_mail%s", id)),
				),
			},
		},
	})
}

func testCheckADUserExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %q", name)
		}

		client := testAccProvider.Meta().(*ArmClient).usersClient
		ctx := testAccProvider.Meta().(*ArmClient).StopContext
		resp, err := client.Get(ctx, rs.Primary.ID)

		if err != nil {
			if ar.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("Bad: Azure AD User %q does not exist", rs.Primary.ID)
			}
			return fmt.Errorf("Bad: Get on Azure AD usersClient: %+v", err)
		}

		return nil
	}
}

func testCheckADUserDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "azuread_user" {
			continue
		}

		client := testAccProvider.Meta().(*ArmClient).usersClient
		ctx := testAccProvider.Meta().(*ArmClient).StopContext
		resp, err := client.Get(ctx, rs.Primary.ID)

		if err != nil {
			if ar.ResponseWasNotFound(resp.Response) {
				return nil
			}

			return err
		}

		return fmt.Errorf("Azure AD User still exists:\n%#v", resp)
	}

	return nil
}

func testAccADUser_basic(id string, password string) string {
	return fmt.Sprintf(`

data "azuread_domains" "tenant_domain" {
	only_initial = true
}

resource "azuread_user" "test" {
	user_principal_name   = "acctest%[1]s@${data.azuread_domains.tenant_domain.domains.0.domain_name}"
	display_name          = "acctest%[1]s"
	password              = "%[2]s"
}
`, id, password)
}

func testAccADUser_complete(id string, password string) string {
	return fmt.Sprintf(`

data "azuread_domains" "tenant_domain" {
	only_initial = true
}

resource "azuread_user" "test" {
	user_principal_name   = "acctest%[1]s@${data.azuread_domains.tenant_domain.domains.0.domain_name}"
	display_name          = "acctestupdate%[1]s"
	mail_nickname         = "acctestupdate%[1]s"
	account_enabled       = false
	password              = "%[2]s"
	force_password_change = true
}
`, id, password)
}

func testAccADUser_multiple(id string, password string) string {
	return fmt.Sprintf(`

data "azuread_domains" "tenant_domain" {
	only_initial = true
}

resource "azuread_user" "testA" {
	user_principal_name   = "acctestA%[1]s@${data.azuread_domains.tenant_domain.domains.0.domain_name}"
	display_name          = "acctestA%[1]s"
	password              = "%[2]s"
}

resource "azuread_user" "testB" {
	user_principal_name   = "acctestB%[1]s@${data.azuread_domains.tenant_domain.domains.0.domain_name}"
	display_name          = "acctest_display%[1]s"
	mail_nickname         = "acctest_mail%[1]s"
	password              = "%[2]s"
}
`, id, password)
}
