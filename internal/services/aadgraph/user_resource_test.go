package aadgraph_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-azuread/internal/acceptance"
	"github.com/terraform-providers/terraform-provider-azuread/internal/clients"
	"github.com/terraform-providers/terraform-provider-azuread/internal/utils"
)

func TestAccUser_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_user", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccUser_basic(data.RandomInteger, data.RandomPassword),
				Check: resource.ComposeTestCheckFunc(
					testCheckUserExists(data.ResourceName),
				),
			},
			data.ImportStep("force_password_change", "password"),
		},
	})
}

func TestAccUser_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_user", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccUser_complete(data.RandomInteger, data.RandomPassword),
				Check: resource.ComposeTestCheckFunc(
					testCheckUserExists(data.ResourceName),
				),
			},
			data.ImportStep("force_password_change", "password"),
		},
	})
}

func TestAccUser_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_user", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccUser_basic(data.RandomInteger, data.RandomPassword),
				Check: resource.ComposeTestCheckFunc(
					testCheckUserExists(data.ResourceName),
				),
			},
			data.ImportStep("force_password_change", "password"),
			{
				Config: testAccUser_complete(data.RandomInteger, data.RandomPassword),
				Check: resource.ComposeTestCheckFunc(
					testCheckUserExists(data.ResourceName),
				),
			},
			data.ImportStep("force_password_change", "password"),
			{
				Config: testAccUser_basic(data.RandomInteger, data.RandomPassword),
				Check: resource.ComposeTestCheckFunc(
					testCheckUserExists(data.ResourceName),
				),
			},
			data.ImportStep("force_password_change", "password"),
		},
	})
}

func TestAccUser_threeUsersABC(t *testing.T) {
	dataA := acceptance.BuildTestData(t, "azuread_user", "testA")
	dataB := acceptance.BuildTestData(t, "azuread_user", "testB")
	dataC := acceptance.BuildTestData(t, "azuread_user", "testC")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccUser_threeUsersABC(dataA.RandomInteger, dataA.RandomPassword),
				Check: resource.ComposeTestCheckFunc(
					testCheckUserExists(dataA.ResourceName),
					testCheckUserExists(dataB.ResourceName),
					testCheckUserExists(dataC.ResourceName),
				),
			},
			dataA.ImportStep("force_password_change", "password"),
			dataB.ImportStep("force_password_change", "password"),
			dataC.ImportStep("force_password_change", "password"),
		},
	})
}

func testCheckUserExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %q", name)
		}

		client := acceptance.AzureADProvider.Meta().(*clients.AadClient).AadGraph.UsersClient
		ctx := acceptance.AzureADProvider.Meta().(*clients.AadClient).StopContext
		resp, err := client.Get(ctx, rs.Primary.ID)

		if err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
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

		client := acceptance.AzureADProvider.Meta().(*clients.AadClient).AadGraph.UsersClient
		ctx := acceptance.AzureADProvider.Meta().(*clients.AadClient).StopContext
		resp, err := client.Get(ctx, rs.Primary.ID)

		if err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
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
  force_password_change = true

  display_name    = "acctestUser-%[1]d-Updated"
  given_name      = "acctestUser-%[1]d-GivenName"
  surname         = "acctestUser-%[1]d-Surname"
  mail_nickname   = "acctestUser-%[1]d-Updated"
  account_enabled = false
  password        = "%[2]s"
  usage_location  = "NO"
  immutable_id    = "%[1]d"

  job_title      = "acctestUser-%[1]d-Job"
  department     = "acctestUser-%[1]d-Dept"
  company_name   = "acctestUser-%[1]d-Company"
  street_address = "acctestUser-%[1]d-Street"
  state          = "acctestUser-%[1]d-State"
  city           = "acctestUser-%[1]d-City"
  country        = "acctestUser-%[1]d-Country"
  postal_code    = "111111"
  mobile         = "(555) 555-5555"

  physical_delivery_office_name = "acctestUser-%[1]d-PDON"
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
