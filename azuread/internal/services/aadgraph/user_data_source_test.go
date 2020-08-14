package aadgraph_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/tf"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/internal/acceptance"
)

func TestAccAzureADUserDataSource_byUserPrincipalName(t *testing.T) {
	dsn := "data.azuread_user.tests"
	id := tf.AccRandTimeInt()
	password := "p@$$wR2" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { acceptance.PreCheck(t) },
		Providers: acceptance.SupportedProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureADUserDataSource_byUserPrincipalName(id, password),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dsn, "user_principal_name"),
					resource.TestCheckResourceAttrSet(dsn, "account_enabled"),
					resource.TestCheckResourceAttrSet(dsn, "display_name"),
					resource.TestCheckResourceAttrSet(dsn, "mail_nickname"),
				),
			},
		},
	})
}

func TestAccAzureADUserDataSource_byUserPrincipalNameNonexistent(t *testing.T) {
	ri := tf.AccRandTimeInt()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { acceptance.PreCheck(t) },
		Providers: acceptance.SupportedProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccAzureADUserDataSource_byUserPrincipalNameNonexistent(ri),
				ExpectError: regexp.MustCompile("Azure AD User not found with UPN:"),
			},
		},
	})
}

func TestAccAzureADUserDataSource_byObjectId(t *testing.T) {
	dsn := "data.azuread_user.tests"
	id := tf.AccRandTimeInt()
	password := "p@$$wR2" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { acceptance.PreCheck(t) },
		Providers: acceptance.SupportedProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureADUserDataSource_byObjectId(id, password),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dsn, "user_principal_name"),
					resource.TestCheckResourceAttrSet(dsn, "account_enabled"),
					resource.TestCheckResourceAttrSet(dsn, "display_name"),
					resource.TestCheckResourceAttrSet(dsn, "mail_nickname"),
				),
			},
		},
	})
}

func TestAccAzureADUserDataSource_byObjectIdNonexistent(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { acceptance.PreCheck(t) },
		Providers: acceptance.SupportedProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccAzureADUserDataSource_byObjectIdNonexistent(),
				ExpectError: regexp.MustCompile("Azure AD User not found with object ID:"),
			},
		},
	})
}

func TestAccAzureADUserDataSource_byMailNickname(t *testing.T) {
	dsn := "data.azuread_user.tests"
	id := tf.AccRandTimeInt()
	password := "p@$$wR2" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { acceptance.PreCheck(t) },
		Providers: acceptance.SupportedProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureADUserDataSource_byMailNickname(id, password),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dsn, "user_principal_name"),
					resource.TestCheckResourceAttrSet(dsn, "account_enabled"),
					resource.TestCheckResourceAttrSet(dsn, "display_name"),
					resource.TestCheckResourceAttrSet(dsn, "mail_nickname"),
				),
			},
		},
	})
}

func TestAccAzureADUserDataSource_byMailNicknameNonexistent(t *testing.T) {
	ri := tf.AccRandTimeInt()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { acceptance.PreCheck(t) },
		Providers: acceptance.SupportedProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccAzureADUserDataSource_byMailNicknameNonexistent(ri),
				ExpectError: regexp.MustCompile("Azure AD User not found with email alias:"),
			},
		},
	})
}

func testAccAzureADUserDataSource_byUserPrincipalName(id int, password string) string {
	return fmt.Sprintf(`
%s

data "azuread_user" "tests" {
  user_principal_name = azuread_user.tests.user_principal_name
}
`, testAccUser_basic(id, password))
}

func testAccAzureADUserDataSource_byUserPrincipalNameNonexistent(ri int) string {
	return fmt.Sprintf(`
data "azuread_domains" "tenant_domain" {
  only_initial = true
}

data "azuread_user" "tests" {
  user_principal_name = "not-a-real-user-%d${data.azuread_domains.tenant_domain.domains.0.domain_name}"
}
`, ri)
}

func testAccAzureADUserDataSource_byObjectId(id int, password string) string {
	return fmt.Sprintf(`
%s

data "azuread_user" "tests" {
  object_id = azuread_user.tests.object_id
}
`, testAccUser_basic(id, password))
}

func testAccAzureADUserDataSource_byObjectIdNonexistent() string {
	return `
data "azuread_user" "tests" {
  object_id = "00000000-0000-0000-0000-000000000000"
}
`
}

func testAccAzureADUserDataSource_byMailNickname(id int, password string) string {
	return fmt.Sprintf(`
%s

data "azuread_user" "tests" {
  mail_nickname = azuread_user.tests.mail_nickname
}
`, testAccUser_basic(id, password))
}

func testAccAzureADUserDataSource_byMailNicknameNonexistent(ri int) string {
	return fmt.Sprintf(`
data "azuread_domains" "tenant_domain" {
  only_initial = true
}

data "azuread_user" "tests" {
  mail_nickname = "not-a-real-user-%d${data.azuread_domains.tenant_domain.domains.0.domain_name}"
}
`, ri)
}
