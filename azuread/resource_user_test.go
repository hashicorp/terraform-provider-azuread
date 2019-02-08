package azuread

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/ar"
)

func TestAccAzureADUser_basic(t *testing.T) {
	resourceName := "azuread_user.test"
	id := strconv.Itoa(rand.Intn(9999))
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

func TestAccAzureADUser_update(t *testing.T) {
	resourceName := "azuread_user.test"
	id := strconv.Itoa(rand.Intn(9999))
	password := id + "p@$$wR2"
	updatedId := strconv.Itoa(rand.Intn(9999))
	updatedPassword := updatedId + "p@$$wR2"

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
					resource.TestCheckResourceAttr(resourceName, "display_name", fmt.Sprintf("acctest%s", id)),
					resource.TestCheckResourceAttr(resourceName, "mail_nickname", fmt.Sprintf("acctest%s", id)),
					resource.TestCheckResourceAttr(resourceName, "account_enabled", "true"),
				),
			},
			{
				Config: testAccADUser_basic(updatedId, updatedPassword),
				Check: resource.ComposeTestCheckFunc(
					testCheckADUserExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "user_principal_name"),
					resource.TestCheckResourceAttr(resourceName, "display_name", fmt.Sprintf("acctest%s", updatedId)),
					resource.TestCheckResourceAttr(resourceName, "mail_nickname", fmt.Sprintf("acctest%s", updatedId)),
					resource.TestCheckResourceAttr(resourceName, "account_enabled", "true"),
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
	user_principal_name = "acctest%s@${data.azuread_domains.tenant_domain.domains.0.domain_name}"
	display_name = "acctest%s"
	mail_nickname = "acctest%s"
	account_enabled = true
	password = "%s"
	force_password_change = true
}
`, id, id, id, password)
}
