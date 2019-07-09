package azuread

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/ar"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/tf"
)

func TestAccAzureADUser_basic(t *testing.T) {
	rn := "azuread_user.test"
	id := tf.AccRandTimeInt()
	pw := "p@$$wRd" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADUser_basic(id, pw),
				Check: resource.ComposeTestCheckFunc(
					testCheckADUserExists(rn),
					resource.TestCheckResourceAttrSet(rn, "user_principal_name"),
					resource.TestCheckResourceAttrSet(rn, "object_id"),
					resource.TestCheckResourceAttr(rn, "display_name", fmt.Sprintf("acctest%d", id)),
					resource.TestCheckResourceAttr(rn, "mail_nickname", fmt.Sprintf("acctest%d", id)),
					resource.TestCheckResourceAttr(rn, "account_enabled", "true"),
				),
			},
			{
				ResourceName:      rn,
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
	rn := "azuread_user.test"
	id := tf.AccRandTimeInt()
	pw := "p@$$wRd" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADUser_complete(id, pw),
				Check: resource.ComposeTestCheckFunc(
					testCheckADUserExists(rn),
					resource.TestCheckResourceAttrSet(rn, "user_principal_name"),
					resource.TestCheckResourceAttrSet(rn, "object_id"),
					resource.TestCheckResourceAttr(rn, "display_name", fmt.Sprintf("acctestupdate%d", id)),
					resource.TestCheckResourceAttr(rn, "mail_nickname", fmt.Sprintf("acctestupdate%d", id)),
					resource.TestCheckResourceAttr(rn, "account_enabled", "false"),
				),
			},
			{
				ResourceName:      rn,
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
	rn := "azuread_user.test"
	id := tf.AccRandTimeInt()
	pw1 := "p@$$wRd" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)
	pw2 := "p@$$wRd2" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADUser_basic(id, pw1),
				Check: resource.ComposeTestCheckFunc(
					testCheckADUserExists(rn),
					resource.TestCheckResourceAttrSet(rn, "user_principal_name"),
					resource.TestCheckResourceAttrSet(rn, "object_id"),
					resource.TestCheckResourceAttr(rn, "display_name", fmt.Sprintf("acctest%d", id)),
					resource.TestCheckResourceAttr(rn, "mail_nickname", fmt.Sprintf("acctest%d", id)),
					resource.TestCheckResourceAttr(rn, "account_enabled", "true"),
				),
			},
			{
				Config: testAccADUser_complete(id, pw2),
				Check: resource.ComposeTestCheckFunc(
					testCheckADUserExists(rn),
					resource.TestCheckResourceAttrSet(rn, "user_principal_name"),
					resource.TestCheckResourceAttrSet(rn, "object_id"),
					resource.TestCheckResourceAttr(rn, "display_name", fmt.Sprintf("acctestupdate%d", id)),
					resource.TestCheckResourceAttr(rn, "mail_nickname", fmt.Sprintf("acctestupdate%d", id)),
					resource.TestCheckResourceAttr(rn, "account_enabled", "false"),
				),
			},
			{
				Config: testAccADUser_multiple(id, pw1),
				Check: resource.ComposeTestCheckFunc(
					testCheckADUserExists("azuread_user.testA"),
					testCheckADUserExists("azuread_user.testB"),
					resource.TestCheckResourceAttrSet("azuread_user.testA", "user_principal_name"),
					resource.TestCheckResourceAttr("azuread_user.testA", "display_name", fmt.Sprintf("acctestA%d", id)),
					resource.TestCheckResourceAttr("azuread_user.testA", "mail_nickname", fmt.Sprintf("acctestA%d", id)),
					resource.TestCheckResourceAttrSet("azuread_user.testB", "user_principal_name"),
					resource.TestCheckResourceAttr("azuread_user.testB", "display_name", fmt.Sprintf("acctest_display%d", id)),
					resource.TestCheckResourceAttr("azuread_user.testB", "mail_nickname", fmt.Sprintf("acctest_mail%d", id)),
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

func testAccADUser_basic(id int, password string) string {
	return fmt.Sprintf(`
data "azuread_domains" "tenant_domain" {
  only_initial = true
}

resource "azuread_user" "test" {
  user_principal_name = "acctest%[1]d@${data.azuread_domains.tenant_domain.domains.0.domain_name}"
  display_name        = "acctest%[1]d"
  password            = "%[2]s"
}
`, id, password)
}

func testAccADUser_complete(id int, password string) string {
	return fmt.Sprintf(`
data "azuread_domains" "tenant_domain" {
  only_initial = true
}

resource "azuread_user" "test" {
  user_principal_name   = "acctest%[1]d@${data.azuread_domains.tenant_domain.domains.0.domain_name}"
  display_name          = "acctestupdate%[1]d"
  mail_nickname         = "acctestupdate%[1]d"
  account_enabled       = false
  password              = "%[2]s"
  force_password_change = true
}
`, id, password)
}

func testAccADUser_multiple(id int, password string) string {
	return fmt.Sprintf(`
data "azuread_domains" "tenant_domain" {
  only_initial = true
}

resource "azuread_user" "testA" {
  user_principal_name = "acctestA%[1]d@${data.azuread_domains.tenant_domain.domains.0.domain_name}"
  display_name        = "acctestA%[1]d"
  password            = "%[2]s"
}

resource "azuread_user" "testB" {
  user_principal_name = "acctestB%[1]d@${data.azuread_domains.tenant_domain.domains.0.domain_name}"
  display_name        = "acctest_display%[1]d"
  mail_nickname       = "acctest_mail%[1]d"
  password            = "%[2]s"
}
`, id, password)
}
