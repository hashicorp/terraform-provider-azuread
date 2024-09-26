// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package users_test

import (
	"context"
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/user"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
)

type UserResource struct{}

func TestAccUser_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_user", "test")
	r := UserResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("force_password_change", "password"),
	})
}

func TestAccUser_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_user", "test")
	r := UserResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.complete(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("force_password_change", "password"),
	})
}

func TestAccUser_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_user", "test")
	r := UserResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("force_password_change", "password"),
		{
			Config: r.complete(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("force_password_change", "password"),
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("force_password_change", "password"),
	})
}

func TestAccUser_threeUsersABC(t *testing.T) {
	dataA := acceptance.BuildTestData(t, "azuread_user", "testA")
	dataB := acceptance.BuildTestData(t, "azuread_user", "testB")
	dataC := acceptance.BuildTestData(t, "azuread_user", "testC")
	r := UserResource{}

	dataA.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.threeUsersABC(dataA),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(dataA.ResourceName).ExistsInAzure(r),
				check.That(dataB.ResourceName).ExistsInAzure(r),
				check.That(dataC.ResourceName).ExistsInAzure(r),
			),
		},
		dataA.ImportStep("force_password_change", "password"),
		dataB.ImportStep("force_password_change", "password"),
		dataC.ImportStep("force_password_change", "password"),
	})
}

func TestAccUser_withRandomProvider(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_user", "test")
	r := UserResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.withRandomProvider(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("force_password_change", "password"),
	})
}

func TestAccUser_passwordOmitted(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_user", "test")
	r := UserResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config:      r.passwordOmitted(data),
			ExpectError: regexp.MustCompile("`password` is required when creating a new user"),
		},
	})
}

func TestAccUser_passwordInvalid(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_user", "test")
	r := UserResource{}
	firstPassword := data.RandomPassword
	secondPassword := "B"

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.setPassword(data, firstPassword),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		{
			Config:      r.setPassword(data, secondPassword),
			ExpectError: regexp.MustCompile("specified password does not comply"),
		},
		{
			RefreshState:       true,
			ExpectNonEmptyPlan: true,
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("password").HasValue(firstPassword),
			),
		},
	})
}

func (r UserResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.Users.UserClient

	id, err := stable.ParseUserID(state.ID)
	if err != nil {
		return nil, err
	}

	resp, err := client.GetUser(ctx, *id, user.DefaultGetUserOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return pointer.To(false), nil
		}
		return nil, fmt.Errorf("failed to retrieve %s: %+v", id, err)
	}
	return pointer.To(true), nil
}

func (UserResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

data "azuread_domains" "test" {
  only_initial = true
}

resource "azuread_user" "test" {
  user_principal_name = "acctestUser'%[1]d@${data.azuread_domains.test.domains.0.domain_name}"
  display_name        = "acctestUser-%[1]d"
  password            = "%[2]s"
}
`, data.RandomInteger, data.RandomPassword)
}

func (UserResource) complete(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

data "azuread_domains" "test" {
  only_initial = true
}

resource "azuread_user" "manager" {
  user_principal_name = "acctestManager.%[1]d@${data.azuread_domains.test.domains.0.domain_name}"
  display_name        = "acctestManager-%[1]d"
  password            = "%[2]s"
}

resource "azuread_user" "test" {
  user_principal_name = "acctestUser'%[1]d-complete@${data.azuread_domains.test.domains.0.domain_name}"
  mail                = "acctestUser.%[1]d@hashicorp.biz"
  mail_nickname       = "acctestUser-%[1]d-MailNickname"
  other_mails         = ["acctestUser.%[1]d@hashicorp.net", "acctestUser.%[1]d@hashicorp.org"]

  account_enabled         = false
  manager_id              = azuread_user.manager.object_id
  onpremises_immutable_id = "%[1]d"
  usage_location          = "NO"

  password                    = "%[2]s"
  force_password_change       = true
  disable_strong_password     = true
  disable_password_expiration = true

  age_group                  = "NotAdult"
  business_phones            = ["12345678901"]
  company_name               = "acctestUser-%[1]d-Company"
  consent_provided_for_minor = "NotRequired"
  cost_center                = "acctestUser-%[1]d-CostCenter"
  department                 = "acctestUser-%[1]d-Dept"
  display_name               = "acctestUser-%[1]d-DisplayName"
  division                   = "acctestUser-%[1]d-Division"
  employee_id                = "%[3]s%[3]s"
  employee_type              = "Contractor"
  fax_number                 = "(555) 555-5555"
  given_name                 = "acctestUser-%[1]d-GivenName"
  job_title                  = "acctestUser-%[1]d-Job"
  mobile_phone               = "(555) 555-5555"
  office_location            = "acctestUser-%[1]d-OfficeLocation"
  preferred_language         = "es-CO"
  show_in_address_list       = false
  surname                    = "acctestUser-%[1]d-Surname"

  street_address = "acctestUser-%[1]d-Street"
  state          = "acctestUser-%[1]d-State"
  city           = "acctestUser-%[1]d-City"
  country        = "acctestUser-%[1]d-Country"
  postal_code    = "111111"
}
`, data.RandomInteger, data.RandomPassword, data.RandomString)
}

func (UserResource) threeUsersABC(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

data "azuread_domains" "test" {
  only_initial = true
}

resource "azuread_user" "testA" {
  user_principal_name = "acctestUser'%[1]d.A@${data.azuread_domains.test.domains.0.domain_name}"
  display_name        = "acctestUser-%[1]d-A"
  employee_id         = "A%[3]s%[3]s"
  mail                = "acctestUser-%[1]d-A@${data.azuread_domains.test.domains.0.domain_name}"
  password            = "%[2]s"
}

resource "azuread_user" "testB" {
  user_principal_name = "acctestUser.%[1]d.B@${data.azuread_domains.test.domains.0.domain_name}"
  display_name        = "acctestUser-%[1]d-B"
  mail_nickname       = "acctestUser-%[1]d-B"
  mail                = "acctestUser-%[1]d-B@${data.azuread_domains.test.domains.0.domain_name}"
  employee_id         = "B%[3]s%[3]s"
  password            = "%[2]s"
}

resource "azuread_user" "testC" {
  user_principal_name = "acctestUser.%[1]d.C@${data.azuread_domains.test.domains.0.domain_name}"
  display_name        = "acctestUser-%[1]d-C"
  password            = "%[2]s"
}
`, data.RandomInteger, data.RandomPassword, data.RandomString)
}

func (UserResource) withRandomProvider(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}
provider "random" {}

data "azuread_domains" "test" {
  only_initial = true
}

resource "random_password" "test" {
  length = 32
}

resource "azuread_user" "test" {
  user_principal_name = "acctestUser.%[1]d@${data.azuread_domains.test.domains.0.domain_name}"
  display_name        = "acctestUser-%[1]d"
  password            = random_password.test.result
}
`, data.RandomInteger)
}

func (UserResource) passwordOmitted(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

data "azuread_domains" "test" {
  only_initial = true
}

resource "azuread_user" "test" {
  user_principal_name = "acctestUser.%[1]d@${data.azuread_domains.test.domains.0.domain_name}"
  display_name        = "acctestUser-%[1]d"
}
`, data.RandomInteger)
}

func (UserResource) setPassword(data acceptance.TestData, password string) string {
	return fmt.Sprintf(`
provider "azuread" {}

data "azuread_domains" "test" {
  only_initial = true
}

resource "azuread_user" "test" {
  user_principal_name = "acctestUser'%[1]d@${data.azuread_domains.test.domains.0.domain_name}"
  display_name        = "acctestUser-%[1]d"
  password            = "%[2]s"
}
`, data.RandomInteger, password)
}
