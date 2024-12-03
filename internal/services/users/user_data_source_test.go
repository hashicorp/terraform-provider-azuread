// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package users_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
)

type UserDataSource struct{}

func TestAccUserDataSource_byUserPrincipalName(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_user", "test")
	r := UserDataSource{}

	data.DataSourceTest(t, []acceptance.TestStep{{
		Config: r.byUserPrincipalName(data),
		Check:  r.testCheckFunc(data),
	}})
}

func TestAccUserDataSource_byUserPrincipalNameNonexistent(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_user", "test")

	data.DataSourceTest(t, []acceptance.TestStep{{
		Config:      UserDataSource{}.byUserPrincipalNameNonexistent(data),
		ExpectError: regexp.MustCompile("User with UPN \"[^\"]+\" was not found"),
	}})
}

func TestAccUserDataSource_byObjectId(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_user", "test")
	r := UserDataSource{}

	data.DataSourceTest(t, []acceptance.TestStep{{
		Config: r.byObjectId(data),
		Check:  r.testCheckFunc(data),
	}})
}

func TestAccUserDataSource_byObjectIdNonexistent(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_user", "test")

	data.DataSourceTest(t, []acceptance.TestStep{{
		Config:      UserDataSource{}.byObjectIdNonexistent(),
		ExpectError: regexp.MustCompile("User not found with object ID:"),
	}})
}

func TestAccUserDataSource_byMailNickname(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_user", "test")
	r := UserDataSource{}

	data.DataSourceTest(t, []acceptance.TestStep{{
		Config: r.byMailNickname(data),
		Check:  r.testCheckFunc(data),
	}})
}

func TestAccUserDataSource_byMailNicknameNonexistent(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_user", "test")

	data.DataSourceTest(t, []acceptance.TestStep{{
		Config:      UserDataSource{}.byMailNicknameNonexistent(data),
		ExpectError: regexp.MustCompile("User not found with email alias:"),
	}})
}

func TestAccUserDataSource_byMail(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_user", "test")
	r := UserDataSource{}

	data.DataSourceTest(t, []acceptance.TestStep{{
		Config: r.byMail(data),
		Check:  r.testCheckFunc(data),
	}})
}

func TestAccUserDataSource_byMailNonexistent(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_user", "test")

	data.DataSourceTest(t, []acceptance.TestStep{{
		Config:      UserDataSource{}.byMailNonexistent(data),
		ExpectError: regexp.MustCompile("User not found with mail:"),
	}})
}

func TestAccUserDataSource_byEmployeeId(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_user", "test")
	r := UserDataSource{}

	data.DataSourceTest(t, []acceptance.TestStep{{
		Config: r.byEmployeeId(data),
		Check:  r.testCheckFunc(data),
	}})
}

func TestAccUserDataSource_byEmployeeIdNonexistent(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_user", "test")

	data.DataSourceTest(t, []acceptance.TestStep{{
		Config:      UserDataSource{}.byEmployeeIdNonexistent(data),
		ExpectError: regexp.MustCompile("User not found with employee ID:"),
	}})
}

func (UserDataSource) testCheckFunc(data acceptance.TestData) acceptance.TestCheckFunc {
	return acceptance.ComposeTestCheckFunc(
		check.That(data.ResourceName).Key("account_enabled").Exists(),
		check.That(data.ResourceName).Key("city").HasValue(fmt.Sprintf("acctestUser-%d-City", data.RandomInteger)),
		check.That(data.ResourceName).Key("company_name").HasValue(fmt.Sprintf("acctestUser-%d-Company", data.RandomInteger)),
		check.That(data.ResourceName).Key("country").HasValue(fmt.Sprintf("acctestUser-%d-Country", data.RandomInteger)),
		check.That(data.ResourceName).Key("department").HasValue(fmt.Sprintf("acctestUser-%d-Dept", data.RandomInteger)),
		check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctestUser-%d-DisplayName", data.RandomInteger)),
		check.That(data.ResourceName).Key("employee_hire_date").HasValue("2018-01-01T01:02:03Z"),
		check.That(data.ResourceName).Key("given_name").HasValue(fmt.Sprintf("acctestUser-%d-GivenName", data.RandomInteger)),
		check.That(data.ResourceName).Key("job_title").HasValue(fmt.Sprintf("acctestUser-%d-Job", data.RandomInteger)),
		// check.That(data.ResourceName).Key("mail").Exists(), // TODO only set for O365 domains
		check.That(data.ResourceName).Key("mail_nickname").HasValue(fmt.Sprintf("acctestUser-%d-MailNickname", data.RandomInteger)),
		check.That(data.ResourceName).Key("manager_id").IsUuid(),
		check.That(data.ResourceName).Key("mobile_phone").HasValue("(555) 555-5555"),
		check.That(data.ResourceName).Key("object_id").IsUuid(),
		check.That(data.ResourceName).Key("office_location").HasValue(fmt.Sprintf("acctestUser-%d-OfficeLocation", data.RandomInteger)),
		check.That(data.ResourceName).Key("onpremises_immutable_id").Exists(),
		check.That(data.ResourceName).Key("postal_code").HasValue("111111"),
		check.That(data.ResourceName).Key("state").HasValue(fmt.Sprintf("acctestUser-%d-State", data.RandomInteger)),
		check.That(data.ResourceName).Key("street_address").HasValue(fmt.Sprintf("acctestUser-%d-Street", data.RandomInteger)),
		check.That(data.ResourceName).Key("surname").HasValue(fmt.Sprintf("acctestUser-%d-Surname", data.RandomInteger)),
		check.That(data.ResourceName).Key("usage_location").HasValue("NO"),
		check.That(data.ResourceName).Key("user_principal_name").Exists(),
		check.That(data.ResourceName).Key("user_type").HasValue("Member"),
	)
}

func (UserDataSource) byUserPrincipalName(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_user" "test" {
  user_principal_name = azuread_user.test.user_principal_name
}
`, UserResource{}.complete(data))
}

func (UserDataSource) byUserPrincipalNameNonexistent(data acceptance.TestData) string {
	return fmt.Sprintf(`
data "azuread_domains" "test" {
  only_initial = true
}

data "azuread_user" "test" {
  user_principal_name = "not-a-real-user-%[1]d${data.azuread_domains.test.domains.0.domain_name}"
}
`, data.RandomInteger)
}

func (UserDataSource) byObjectId(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_user" "test" {
  object_id = azuread_user.test.object_id
}
`, UserResource{}.complete(data))
}

func (UserDataSource) byObjectIdNonexistent() string {
	return `
data "azuread_user" "test" {
  object_id = "00000000-0000-0000-0000-000000000000"
}
`
}

func (UserDataSource) byMailNickname(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_user" "test" {
  mail_nickname = azuread_user.test.mail_nickname
}
`, UserResource{}.complete(data))
}

func (UserDataSource) byMailNicknameNonexistent(data acceptance.TestData) string {
	return fmt.Sprintf(`
data "azuread_domains" "test" {
  only_initial = true
}

data "azuread_user" "test" {
  mail_nickname = "not-a-real-user-%[1]d${data.azuread_domains.test.domains.0.domain_name}"
}
`, data.RandomInteger)
}

func (UserDataSource) byMail(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_user" "test" {
  mail = azuread_user.test.mail
}
`, UserResource{}.complete(data))
}

func (UserDataSource) byMailNonexistent(data acceptance.TestData) string {
	return fmt.Sprintf(`
data "azuread_domains" "test" {
  only_initial = true
}

data "azuread_user" "test" {
  mail = "not-a-real-user-%[1]d${data.azuread_domains.test.domains.0.domain_name}"
}
`, data.RandomInteger)
}

func (UserDataSource) byEmployeeId(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_user" "test" {
  employee_id = azuread_user.test.employee_id
}
`, UserResource{}.complete(data))
}

func (UserDataSource) byEmployeeIdNonexistent(data acceptance.TestData) string {
	return `
data "azuread_domains" "test" {
  only_initial = true
}

data "azuread_user" "test" {
  employee_id = "not-a-real-employeeid"
}
`
}
