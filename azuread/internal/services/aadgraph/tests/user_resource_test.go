package tests

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/ar"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/tf"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/internal/acceptance"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/internal/clients"
)

func TestAccUser_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_user", "test")
	pw := "p@$$wRd" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccUser_basic(data.RandomInteger, pw),
				Check: resource.ComposeTestCheckFunc(
					testCheckUserExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "user_principal_name"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "object_id"),
					resource.TestCheckResourceAttr(data.ResourceName, "display_name", fmt.Sprintf("acctestUser-%d", data.RandomInteger)),
					resource.TestCheckResourceAttr(data.ResourceName, "mail_nickname", fmt.Sprintf("acctestUser.%d", data.RandomInteger)),
					resource.TestCheckResourceAttr(data.ResourceName, "account_enabled", "true"),
				),
			},
			data.ImportStep("force_password_change", "password"),
		},
	})
}

func TestAccUser_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_user", "test")
	pw := "p@$$wRd" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccUser_complete(data.RandomInteger, pw),
				Check: resource.ComposeTestCheckFunc(
					testCheckUserExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "user_principal_name"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "object_id"),
					resource.TestCheckResourceAttr(data.ResourceName, "display_name", fmt.Sprintf("acctestUser-%d-Updated", data.RandomInteger)),
					resource.TestCheckResourceAttr(data.ResourceName, "mail_nickname", fmt.Sprintf("acctestUser-%d-Updated", data.RandomInteger)),
					resource.TestCheckResourceAttr(data.ResourceName, "account_enabled", "false"),
					resource.TestCheckResourceAttr(data.ResourceName, "immutable_id", strconv.Itoa(data.RandomInteger)),
				),
			},
			data.ImportStep("force_password_change", "password"),
		},
	})
}

func TestAccUser_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_user", "test")
	pw1 := "p@$$wRd" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)
	pw2 := "p@$$wRd2" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccUser_basic(data.RandomInteger, pw1),
				Check: resource.ComposeTestCheckFunc(
					testCheckUserExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "user_principal_name"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "object_id"),
					resource.TestCheckResourceAttr(data.ResourceName, "display_name", fmt.Sprintf("acctestUser-%d", data.RandomInteger)),
					resource.TestCheckResourceAttr(data.ResourceName, "mail_nickname", fmt.Sprintf("acctestUser.%d", data.RandomInteger)),
					resource.TestCheckResourceAttr(data.ResourceName, "account_enabled", "true"),
				),
			},
			data.ImportStep("force_password_change", "password"),
			{
				Config: testAccUser_complete(data.RandomInteger, pw2),
				Check: resource.ComposeTestCheckFunc(
					testCheckUserExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "user_principal_name"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "object_id"),
					resource.TestCheckResourceAttr(data.ResourceName, "display_name", fmt.Sprintf("acctestUser-%d-Updated", data.RandomInteger)),
					resource.TestCheckResourceAttr(data.ResourceName, "mail_nickname", fmt.Sprintf("acctestUser-%d-Updated", data.RandomInteger)),
					resource.TestCheckResourceAttr(data.ResourceName, "account_enabled", "false"),
					resource.TestCheckResourceAttr(data.ResourceName, "immutable_id", strconv.Itoa(data.RandomInteger)),
				),
			},
			data.ImportStep("force_password_change", "password"),
		},
	})
}

func TestAccUser_threeUsersABC(t *testing.T) {
	ri := tf.AccRandTimeInt()
	pw := "p@$$wRd" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccUser_threeUsersABC(ri, pw),
				Check: resource.ComposeTestCheckFunc(
					testCheckUserExists("azuread_user.testA"),
					testCheckUserExists("azuread_user.testB"),
					resource.TestCheckResourceAttrSet("azuread_user.testA", "user_principal_name"),
					resource.TestCheckResourceAttr("azuread_user.testA", "display_name", fmt.Sprintf("acctestUser-%d-A", ri)),
					resource.TestCheckResourceAttr("azuread_user.testA", "mail_nickname", fmt.Sprintf("acctestUser.%d.A", ri)),
					resource.TestCheckResourceAttrSet("azuread_user.testB", "user_principal_name"),
					resource.TestCheckResourceAttr("azuread_user.testB", "display_name", fmt.Sprintf("acctestUser-%d-B", ri)),
					resource.TestCheckResourceAttr("azuread_user.testB", "mail_nickname", fmt.Sprintf("acctestUser-%d-B", ri)),
				),
			},
			{
				ResourceName:            "azuread_user.testA",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"force_password_change", "password"},
			},
			{
				ResourceName:            "azuread_user.testB",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"force_password_change", "password"},
			},
			{
				ResourceName:            "azuread_user.testC",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"force_password_change", "password"},
			},
		},
	})
}

func testCheckUserExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %q", name)
		}

		client := acceptance.AzureADProvider.Meta().(*clients.AadClient).UsersClient
		ctx := acceptance.AzureADProvider.Meta().(*clients.AadClient).StopContext
		resp, err := client.Get(ctx, rs.Primary.ID)

		if err != nil {
			if ar.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("Bad: User %q does not exist", rs.Primary.ID)
			}
			return fmt.Errorf("Bad: Get on UsersClient: %+v", err)
		}

		return nil
	}
}

func testCheckUserDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "azuread_user" {
			continue
		}

		client := acceptance.AzureADProvider.Meta().(*clients.AadClient).UsersClient
		ctx := acceptance.AzureADProvider.Meta().(*clients.AadClient).StopContext
		resp, err := client.Get(ctx, rs.Primary.ID)

		if err != nil {
			if ar.ResponseWasNotFound(resp.Response) {
				return nil
			}

			return err
		}

		return fmt.Errorf("User still exists:\n%#v", resp)
	}

	return nil
}

func testAccUser_basic(id int, password string) string {
	return fmt.Sprintf(`
data "azuread_domains" "tenant_domain" {
  only_initial = true
}

resource "azuread_user" "test" {
  user_principal_name = "acctestUser.%[1]d@${data.azuread_domains.tenant_domain.domains.0.domain_name}"
  display_name        = "acctestUser-%[1]d"
  password            = "%[2]s"
}
`, id, password)
}

func testAccUser_complete(id int, password string) string {
	return fmt.Sprintf(`
data "azuread_domains" "tenant_domain" {
  only_initial = true
}

resource "azuread_user" "test" {
  user_principal_name   = "acctestUser.%[1]d@${data.azuread_domains.tenant_domain.domains.0.domain_name}"
  display_name          = "acctestUser-%[1]d-Updated"
  mail_nickname         = "acctestUser-%[1]d-Updated"
  account_enabled       = false
  password              = "%[2]s"
  force_password_change = true
  usage_location        = "NO"
  immutable_id          = "%[1]d"
}
`, id, password)
}

func testAccUser_threeUsersABC(id int, password string) string {
	return fmt.Sprintf(`
data "azuread_domains" "tenant_domain" {
  only_initial = true
}

resource "azuread_user" "testA" {
  user_principal_name = "acctestUser.%[1]d.A@${data.azuread_domains.tenant_domain.domains.0.domain_name}"
  display_name        = "acctestUser-%[1]d-A"
  password            = "%[2]s"
}

resource "azuread_user" "testB" {
  user_principal_name = "acctestUser.%[1]d.B@${data.azuread_domains.tenant_domain.domains.0.domain_name}"
  display_name        = "acctestUser-%[1]d-B"
  mail_nickname       = "acctestUser-%[1]d-B"
  password            = "%[2]s"
}

resource "azuread_user" "testC" {
  user_principal_name = "acctestUser.%[1]d.C@${data.azuread_domains.tenant_domain.domains.0.domain_name}"
  display_name        = "acctestUser-%[1]d-C"
  password            = "%[2]s"
}
`, id, password)
}
