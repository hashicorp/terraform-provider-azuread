// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package users_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
)

type UsersDataSource struct{}

func TestAccUsersDataSource_byUserPrincipalNames(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_users", "test")

	data.DataSourceTest(t, []acceptance.TestStep{{
		Config: UsersDataSource{}.byUserPrincipalNames(data),
		Check: acceptance.ComposeTestCheckFunc(
			check.That(data.ResourceName).Key("user_principal_names.#").HasValue("2"),
			check.That(data.ResourceName).Key("object_ids.#").HasValue("2"),
			check.That(data.ResourceName).Key("mail_nicknames.#").HasValue("2"),
			check.That(data.ResourceName).Key("employee_ids.#").HasValue("2"),
			check.That(data.ResourceName).Key("users.#").HasValue("2"),
		),
	}})
}

func TestAccUsersDataSource_byUserPrincipalNamesIgnoreMissing(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_users", "test")

	data.DataSourceTest(t, []acceptance.TestStep{{
		Config: UsersDataSource{}.byUserPrincipalNamesIgnoreMissing(data),
		Check: acceptance.ComposeTestCheckFunc(
			check.That(data.ResourceName).Key("user_principal_names.#").HasValue("3"),
			check.That(data.ResourceName).Key("object_ids.#").HasValue("3"),
			check.That(data.ResourceName).Key("mail_nicknames.#").HasValue("3"),
			check.That(data.ResourceName).Key("employee_ids.#").HasValue("2"),
			check.That(data.ResourceName).Key("users.#").HasValue("3"),
		),
	}})
}

func TestAccUsersDataSource_byObjectIds(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_users", "test")

	data.DataSourceTest(t, []acceptance.TestStep{{
		Config: UsersDataSource{}.byObjectIds(data),
		Check: acceptance.ComposeTestCheckFunc(
			check.That(data.ResourceName).Key("user_principal_names.#").HasValue("2"),
			check.That(data.ResourceName).Key("object_ids.#").HasValue("2"),
			check.That(data.ResourceName).Key("mail_nicknames.#").HasValue("2"),
			check.That(data.ResourceName).Key("employee_ids.#").HasValue("2"),
			check.That(data.ResourceName).Key("users.#").HasValue("2"),
		),
	}})
}

func TestAccUsersDataSource_byObjectIdsIgnoreMissing(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_users", "test")

	data.DataSourceTest(t, []acceptance.TestStep{{
		Config: UsersDataSource{}.byObjectIdsIgnoreMissing(data),
		Check: acceptance.ComposeTestCheckFunc(
			check.That(data.ResourceName).Key("user_principal_names.#").HasValue("2"),
			check.That(data.ResourceName).Key("object_ids.#").HasValue("2"),
			check.That(data.ResourceName).Key("mail_nicknames.#").HasValue("2"),
			check.That(data.ResourceName).Key("employee_ids.#").HasValue("2"),
			check.That(data.ResourceName).Key("users.#").HasValue("2"),
		),
	}})
}

func TestAccUsersDataSource_byMailNicknames(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_users", "test")

	data.DataSourceTest(t, []acceptance.TestStep{{
		Config: UsersDataSource{}.byMailNicknames(data),
		Check: acceptance.ComposeTestCheckFunc(
			check.That(data.ResourceName).Key("user_principal_names.#").HasValue("2"),
			check.That(data.ResourceName).Key("object_ids.#").HasValue("2"),
			check.That(data.ResourceName).Key("mail_nicknames.#").HasValue("2"),
			check.That(data.ResourceName).Key("employee_ids.#").HasValue("2"),
			check.That(data.ResourceName).Key("users.#").HasValue("2"),
		),
	}})
}

func TestAccUsersDataSource_byMailNicknamesIgnoreMissing(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_users", "test")

	data.DataSourceTest(t, []acceptance.TestStep{{
		Config: UsersDataSource{}.byMailNicknamesIgnoreMissing(data),
		Check: acceptance.ComposeTestCheckFunc(
			check.That(data.ResourceName).Key("user_principal_names.#").HasValue("2"),
			check.That(data.ResourceName).Key("object_ids.#").HasValue("2"),
			check.That(data.ResourceName).Key("mail_nicknames.#").HasValue("2"),
			check.That(data.ResourceName).Key("employee_ids.#").HasValue("2"),
			check.That(data.ResourceName).Key("users.#").HasValue("2"),
		),
	}})
}

func TestAccUsersDataSource_byMails(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_users", "test")

	data.DataSourceTest(t, []acceptance.TestStep{{
		Config: UsersDataSource{}.byMails(data),
		Check: acceptance.ComposeTestCheckFunc(
			check.That(data.ResourceName).Key("user_principal_names.#").HasValue("2"),
			check.That(data.ResourceName).Key("object_ids.#").HasValue("2"),
			check.That(data.ResourceName).Key("mail_nicknames.#").HasValue("2"),
			check.That(data.ResourceName).Key("mails.#").HasValue("2"),
			check.That(data.ResourceName).Key("employee_ids.#").HasValue("2"),
			check.That(data.ResourceName).Key("users.#").HasValue("2"),
		),
	}})
}

func TestAccUsersDataSource_byMailsIgnoreMissing(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_users", "test")

	data.DataSourceTest(t, []acceptance.TestStep{{
		Config: UsersDataSource{}.byMailsIgnoreMissing(data),
		Check: acceptance.ComposeTestCheckFunc(
			check.That(data.ResourceName).Key("user_principal_names.#").HasValue("2"),
			check.That(data.ResourceName).Key("object_ids.#").HasValue("2"),
			check.That(data.ResourceName).Key("mail_nicknames.#").HasValue("2"),
			check.That(data.ResourceName).Key("mails.#").HasValue("2"),
			check.That(data.ResourceName).Key("employee_ids.#").HasValue("2"),
			check.That(data.ResourceName).Key("users.#").HasValue("2"),
		),
	}})
}

func TestAccUsersDataSource_byEmployeeIds(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_users", "test")

	data.DataSourceTest(t, []acceptance.TestStep{{
		Config: UsersDataSource{}.byEmployeeIds(data),
		Check: acceptance.ComposeTestCheckFunc(
			check.That(data.ResourceName).Key("user_principal_names.#").HasValue("2"),
			check.That(data.ResourceName).Key("object_ids.#").HasValue("2"),
			check.That(data.ResourceName).Key("mail_nicknames.#").HasValue("2"),
			check.That(data.ResourceName).Key("employee_ids.#").HasValue("2"),
			check.That(data.ResourceName).Key("users.#").HasValue("2"),
		),
	}})
}

func TestAccUsersDataSource_byEmployeeIdsIgnoreMissing(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_users", "test")

	data.DataSourceTest(t, []acceptance.TestStep{{
		Config: UsersDataSource{}.byEmployeeIdsIgnoreMissing(data),
		Check: acceptance.ComposeTestCheckFunc(
			check.That(data.ResourceName).Key("user_principal_names.#").HasValue("2"),
			check.That(data.ResourceName).Key("object_ids.#").HasValue("2"),
			check.That(data.ResourceName).Key("mail_nicknames.#").HasValue("2"),
			check.That(data.ResourceName).Key("employee_ids.#").HasValue("2"),
			check.That(data.ResourceName).Key("users.#").HasValue("2"),
		),
	}})
}

func TestAccUsersDataSource_noNames(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_users", "test")

	data.DataSourceTest(t, []acceptance.TestStep{{
		Config: UsersDataSource{}.noNames(),
		Check: acceptance.ComposeTestCheckFunc(
			check.That(data.ResourceName).Key("user_principal_names.#").HasValue("0"),
			check.That(data.ResourceName).Key("object_ids.#").HasValue("0"),
			check.That(data.ResourceName).Key("mail_nicknames.#").HasValue("0"),
			check.That(data.ResourceName).Key("employee_ids.#").HasValue("0"),
			check.That(data.ResourceName).Key("users.#").HasValue("0"),
		),
	}})
}

func TestAccUsersDataSource_returnAll(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_users", "test")

	data.DataSourceTest(t, []acceptance.TestStep{{
		Config: UsersDataSource{}.returnAll(),
		Check: acceptance.ComposeTestCheckFunc(
			check.That(data.ResourceName).Key("user_principal_names.#").Exists(),
			check.That(data.ResourceName).Key("object_ids.#").Exists(),
			check.That(data.ResourceName).Key("mail_nicknames.#").Exists(),
			check.That(data.ResourceName).Key("employee_ids.#").Exists(),
			check.That(data.ResourceName).Key("users.#").Exists(),
		),
	}})
}

func (UsersDataSource) byUserPrincipalNames(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_users" "test" {
  user_principal_names = [azuread_user.testA.user_principal_name, azuread_user.testB.user_principal_name]
}
`, UserResource{}.threeUsersABC(data))
}

func (UsersDataSource) byUserPrincipalNamesIgnoreMissing(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_users" "test" {
  ignore_missing = true

  user_principal_names = [
    azuread_user.testA.user_principal_name,
    "not-a-real-user-%[2]d${data.azuread_domains.test.domains.0.domain_name}",
    azuread_user.testB.user_principal_name,
    azuread_user.testC.user_principal_name,
  ]
}
`, UserResource{}.threeUsersABC(data), data.RandomInteger)
}

func (UsersDataSource) byObjectIds(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_users" "test" {
  object_ids = [azuread_user.testA.object_id, azuread_user.testB.object_id]
}
`, UserResource{}.threeUsersABC(data))
}

func (UsersDataSource) byObjectIdsIgnoreMissing(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_users" "test" {
  ignore_missing = true

  object_ids = [
    azuread_user.testA.object_id,
    "00000000-0000-0000-0000-000000000000",
    azuread_user.testB.object_id,
  ]
}
`, UserResource{}.threeUsersABC(data))
}

func (UsersDataSource) byMailNicknames(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_users" "test" {
  mail_nicknames = [azuread_user.testA.mail_nickname, azuread_user.testB.mail_nickname]
}
`, UserResource{}.threeUsersABC(data))
}

func (UsersDataSource) byMailNicknamesIgnoreMissing(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_users" "test" {
  ignore_missing = true

  mail_nicknames = [
    azuread_user.testA.mail_nickname,
    "not-a-real-user-%[2]d${data.azuread_domains.test.domains.0.domain_name}",
    azuread_user.testB.mail_nickname,
  ]
}
`, UserResource{}.threeUsersABC(data), data.RandomInteger)
}

func (UsersDataSource) byMails(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_users" "test" {
  mails = [azuread_user.testA.mail, azuread_user.testB.mail]
}
`, UserResource{}.threeUsersABC(data))
}

func (UsersDataSource) byMailsIgnoreMissing(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_users" "test" {
  ignore_missing = true

  mails = [
    azuread_user.testA.mail,
    "not-a-real-user-%[2]d${data.azuread_domains.test.domains.0.domain_name}",
    azuread_user.testB.mail,
  ]
}
`, UserResource{}.threeUsersABC(data), data.RandomInteger)
}

func (UsersDataSource) byEmployeeIds(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_users" "test" {
  employee_ids = [azuread_user.testA.employee_id, azuread_user.testB.employee_id]
}
`, UserResource{}.threeUsersABC(data))
}

func (UsersDataSource) byEmployeeIdsIgnoreMissing(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_users" "test" {
  ignore_missing = true

  employee_ids = [
    azuread_user.testA.employee_id,
    "not-a-real-employee-id",
    azuread_user.testB.employee_id,
  ]
}
`, UserResource{}.threeUsersABC(data))
}

func (UsersDataSource) noNames() string {
	return `
data "azuread_users" "test" {
  user_principal_names = []
}
`
}

func (UsersDataSource) returnAll() string {
	return `
data "azuread_users" "test" {
  return_all = true
}
`
}
