// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package directoryroles_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/directoryroles/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
)

type DirectoryRoleMemberResource struct{}

func TestAccDirectoryRoleMember_servicePrincipal(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_directory_role_member", "test")
	r := DirectoryRoleMemberResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.servicePrincipal(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("role_object_id").IsUuid(),
				check.That(data.ResourceName).Key("member_object_id").IsUuid(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccDirectoryRoleMember_user(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_directory_role_member", "testA")
	r := DirectoryRoleMemberResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.oneUser(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("role_object_id").IsUuid(),
				check.That(data.ResourceName).Key("member_object_id").IsUuid(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccDirectoryRoleMember_multipleUser(t *testing.T) {
	dataA := acceptance.BuildTestData(t, "azuread_directory_role_member", "testA")
	dataB := acceptance.BuildTestData(t, "azuread_directory_role_member", "testB")
	r := DirectoryRoleMemberResource{}

	dataA.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.oneUser(dataA),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(dataA.ResourceName).ExistsInAzure(r),
				check.That(dataA.ResourceName).Key("role_object_id").IsUuid(),
				check.That(dataA.ResourceName).Key("member_object_id").IsUuid(),
			),
		},
		dataA.ImportStep(),
		{
			Config: r.twoUsers(dataA),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(dataA.ResourceName).ExistsInAzure(r),
				check.That(dataA.ResourceName).Key("role_object_id").IsUuid(),
				check.That(dataA.ResourceName).Key("member_object_id").IsUuid(),
				check.That(dataB.ResourceName).ExistsInAzure(r),
				check.That(dataB.ResourceName).Key("role_object_id").IsUuid(),
				check.That(dataB.ResourceName).Key("member_object_id").IsUuid(),
			),
		},
		dataA.ImportStep(),
		dataB.ImportStep(),
		{
			Config: r.oneUser(dataA),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(dataA.ResourceName).ExistsInAzure(r),
				check.That(dataA.ResourceName).Key("role_object_id").IsUuid(),
				check.That(dataA.ResourceName).Key("member_object_id").IsUuid(),
			),
		},
		dataA.ImportStep(),
	})
}

func TestAccDirectoryRoleMember_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_directory_role_member", "test")
	r := DirectoryRoleMemberResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.servicePrincipal(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.RequiresImportErrorStep(r.requiresImport(data)),
	})
}

func (r DirectoryRoleMemberResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.DirectoryRoles.DirectoryRolesClient
	client.BaseClient.DisableRetries = true
	defer func() { client.BaseClient.DisableRetries = false }()

	id, err := parse.DirectoryRoleMemberID(state.ID)
	if err != nil {
		return nil, fmt.Errorf("parsing Directory Role Member ID: %v", err)
	}

	if _, status, err := client.GetMember(ctx, id.DirectoryRoleId, id.MemberId); err != nil {
		if status == http.StatusNotFound {
			return utils.Bool(false), nil
		}
		return nil, fmt.Errorf("failed to retrieve directory role member %q (role ID: %q): %+v", id.MemberId, id.DirectoryRoleId, err)
	}

	return utils.Bool(true), nil
}

func (DirectoryRoleMemberResource) templateThreeUsers(data acceptance.TestData) string {
	return fmt.Sprintf(`
data "azuread_domains" "test" {
  only_initial = true
}

resource "azuread_user" "testA" {
  user_principal_name = "acctestUser.%[1]d.A@${data.azuread_domains.test.domains.0.domain_name}"
  display_name        = "acctestUser-%[1]d-A"
  password            = "%[2]s"
}

resource "azuread_user" "testB" {
  user_principal_name = "acctestUser.%[1]d.B@${data.azuread_domains.test.domains.0.domain_name}"
  display_name        = "acctestUser-%[1]d-B"
  mail_nickname       = "acctestUser-%[1]d-B"
  password            = "%[2]s"
}

resource "azuread_user" "testC" {
  user_principal_name = "acctestUser.%[1]d.C@${data.azuread_domains.test.domains.0.domain_name}"
  display_name        = "acctestUser-%[1]d-C"
  password            = "%[2]s"
}
`, data.RandomInteger, data.RandomPassword)
}

func (r DirectoryRoleMemberResource) servicePrincipal(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_application" "test" {
  display_name = "acctestServicePrincipal-%[2]d"
}

resource "azuread_service_principal" "test" {
  application_id = azuread_application.test.application_id
}

resource "azuread_directory_role_member" "test" {
  role_object_id   = azuread_directory_role.test.object_id
  member_object_id = azuread_service_principal.test.object_id
}
`, DirectoryRoleResource{}.byTemplateId(data), data.RandomInteger)
}

func (r DirectoryRoleMemberResource) oneUser(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s
%[2]s

resource "azuread_directory_role_member" "testA" {
  role_object_id   = azuread_directory_role.test.object_id
  member_object_id = azuread_user.testA.object_id
}
`, DirectoryRoleResource{}.byTemplateId(data), r.templateThreeUsers(data))
}

func (r DirectoryRoleMemberResource) twoUsers(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s
%[2]s

resource "azuread_directory_role_member" "testA" {
  role_object_id   = azuread_directory_role.test.object_id
  member_object_id = azuread_user.testA.object_id
}

resource "azuread_directory_role_member" "testB" {
  role_object_id   = azuread_directory_role.test.object_id
  member_object_id = azuread_user.testB.object_id
}
`, DirectoryRoleResource{}.byTemplateId(data), r.templateThreeUsers(data))
}

func (r DirectoryRoleMemberResource) requiresImport(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_directory_role_member" "import" {
  role_object_id   = azuread_directory_role_member.test.role_object_id
  member_object_id = azuread_directory_role_member.test.member_object_id
}
`, r.servicePrincipal(data))
}
