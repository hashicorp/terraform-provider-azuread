package aadgraph_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-azuread/azuread/internal/acceptance"
)

func TestAccUsersDataSource_byUserPrincipalNames(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_users", "test")
	password := "p@$$wR2" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { acceptance.PreCheck(t) },
		Providers: acceptance.SupportedProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccUsersDataSource_byUserPrincipalNames(data.RandomInteger, password),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(data.ResourceName, "user_principal_names.#", "2"),
					resource.TestCheckResourceAttr(data.ResourceName, "object_ids.#", "2"),
					resource.TestCheckResourceAttr(data.ResourceName, "users.#", "2"),
				),
			},
		},
	})
}

func TestAccUsersDataSource_byUserPrincipalNamesIgnoreMissing(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_users", "test")
	password := "p@$$wR2" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { acceptance.PreCheck(t) },
		Providers: acceptance.SupportedProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccUsersDataSource_byUserPrincipalNamesIgnoreMissing(data.RandomInteger, password),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(data.ResourceName, "user_principal_names.#", "3"),
					resource.TestCheckResourceAttr(data.ResourceName, "object_ids.#", "3"),
					resource.TestCheckResourceAttr(data.ResourceName, "users.#", "3"),
				),
			},
		},
	})
}

func TestAccUsersDataSource_byObjectIds(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_users", "test")
	password := "p@$$wR2" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { acceptance.PreCheck(t) },
		Providers: acceptance.SupportedProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccUsersDataSource_byObjectIds(data.RandomInteger, password),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(data.ResourceName, "user_principal_names.#", "2"),
					resource.TestCheckResourceAttr(data.ResourceName, "object_ids.#", "2"),
					resource.TestCheckResourceAttr(data.ResourceName, "users.#", "2"),
				),
			},
		},
	})
}

func TestAccUsersDataSource_byObjectIdsIgnoreMissing(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_users", "test")
	password := "p@$$wR2" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { acceptance.PreCheck(t) },
		Providers: acceptance.SupportedProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccUsersDataSource_byObjectIdsIgnoreMissing(data.RandomInteger, password),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(data.ResourceName, "user_principal_names.#", "2"),
					resource.TestCheckResourceAttr(data.ResourceName, "object_ids.#", "2"),
					resource.TestCheckResourceAttr(data.ResourceName, "users.#", "2"),
				),
			},
		},
	})
}

func TestAccUsersDataSource_byMailNicknames(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_users", "test")
	password := "p@$$wR2" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { acceptance.PreCheck(t) },
		Providers: acceptance.SupportedProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccUsersDataSource_byMailNicknames(data.RandomInteger, password),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(data.ResourceName, "user_principal_names.#", "2"),
					resource.TestCheckResourceAttr(data.ResourceName, "object_ids.#", "2"),
					resource.TestCheckResourceAttr(data.ResourceName, "mail_nicknames.#", "2"),
					resource.TestCheckResourceAttr(data.ResourceName, "users.#", "2"),
				),
			},
		},
	})
}

func TestAccUsersDataSource_byMailNicknamesIgnoreMissing(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_users", "test")
	password := "p@$$wR2" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { acceptance.PreCheck(t) },
		Providers: acceptance.SupportedProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccUsersDataSource_byMailNicknamesIgnoreMissing(data.RandomInteger, password),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(data.ResourceName, "user_principal_names.#", "2"),
					resource.TestCheckResourceAttr(data.ResourceName, "object_ids.#", "2"),
					resource.TestCheckResourceAttr(data.ResourceName, "mail_nicknames.#", "2"),
					resource.TestCheckResourceAttr(data.ResourceName, "users.#", "2"),
				),
			},
		},
	})
}

func TestAccUsersDataSource_noNames(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_users", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { acceptance.PreCheck(t) },
		Providers: acceptance.SupportedProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccUsersDataSource_noNames(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(data.ResourceName, "user_principal_names.#", "0"),
					resource.TestCheckResourceAttr(data.ResourceName, "object_ids.#", "0"),
					resource.TestCheckResourceAttr(data.ResourceName, "mail_nicknames.#", "0"),
				),
			},
		},
	})
}

func testAccUsersDataSource_byUserPrincipalNames(id int, password string) string {
	return fmt.Sprintf(`
%s

data "azuread_users" "test" {
  user_principal_names = [azuread_user.testA.user_principal_name, azuread_user.testB.user_principal_name]
}
`, testAccUser_threeUsersABC(id, password))
}

func testAccUsersDataSource_byUserPrincipalNamesIgnoreMissing(id int, password string) string {
	return fmt.Sprintf(`
%s

data "azuread_users" "test" {
  ignore_missing = true

  user_principal_names = [
    azuread_user.testA.user_principal_name,
    azuread_user.testB.user_principal_name,
    azuread_user.testC.user_principal_name,
    "not-a-real-user-%d${data.azuread_domains.tenant_domain.domains.0.domain_name}",
  ]
}
`, testAccUser_threeUsersABC(id, password), id)
}

func testAccUsersDataSource_byObjectIds(id int, password string) string {
	return fmt.Sprintf(`
%s

data "azuread_users" "test" {
  object_ids = [azuread_user.testA.object_id, azuread_user.testB.object_id]
}
`, testAccUser_threeUsersABC(id, password))
}

func testAccUsersDataSource_byObjectIdsIgnoreMissing(id int, password string) string {
	return fmt.Sprintf(`
%s

data "azuread_users" "test" {
  ignore_missing = true

  object_ids = [
    azuread_user.testA.object_id,
    azuread_user.testB.object_id,
    "00000000-0000-0000-0000-000000000000"
  ]
}
`, testAccUser_threeUsersABC(id, password))
}

func testAccUsersDataSource_byMailNicknames(id int, password string) string {
	return fmt.Sprintf(`
%s

data "azuread_users" "test" {
  mail_nicknames = [azuread_user.testA.mail_nickname, azuread_user.testB.mail_nickname]
}
`, testAccUser_threeUsersABC(id, password))
}

func testAccUsersDataSource_byMailNicknamesIgnoreMissing(id int, password string) string {
	return fmt.Sprintf(`
%s

data "azuread_users" "test" {
  ignore_missing = true

  mail_nicknames = [
    azuread_user.testA.mail_nickname,
    azuread_user.testB.mail_nickname,
    "not-a-real-user-%d${data.azuread_domains.tenant_domain.domains.0.domain_name}"
  ]
}
`, testAccUser_threeUsersABC(id, password), id)
}

func testAccUsersDataSource_noNames() string {
	return `
data "azuread_users" "test" {
  user_principal_names = []
}
`
}
