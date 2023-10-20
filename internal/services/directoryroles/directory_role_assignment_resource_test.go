// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package directoryroles_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
)

type DirectoryRoleAssignmentResource struct{}

func TestAccDirectoryRoleAssignment_servicePrincipal(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_directory_role_assignment", "test")
	r := DirectoryRoleAssignmentResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.servicePrincipal(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("role_id").IsUuid(),
				check.That(data.ResourceName).Key("principal_object_id").IsUuid(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccDirectoryRoleAssignment_servicePrincipalWithCustomRole(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_directory_role_assignment", "test")
	r := DirectoryRoleAssignmentResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.servicePrincipalCustomRole(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("role_id").IsUuid(),
				check.That(data.ResourceName).Key("principal_object_id").IsUuid(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccDirectoryRoleAssignment_servicePrincipalScopedApplication(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_directory_role_assignment", "test")
	r := DirectoryRoleAssignmentResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.servicePrincipalScopedApplication(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("role_id").IsUuid(),
				check.That(data.ResourceName).Key("principal_object_id").IsUuid(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccDirectoryRoleAssignment_user(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_directory_role_assignment", "testA")
	r := DirectoryRoleAssignmentResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.oneUser(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("role_id").IsUuid(),
				check.That(data.ResourceName).Key("principal_object_id").IsUuid(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccDirectoryRoleAssignment_userWithCustomRole(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_directory_role_assignment", "testA")
	r := DirectoryRoleAssignmentResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.oneUserCustomRole(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("role_id").IsUuid(),
				check.That(data.ResourceName).Key("principal_object_id").IsUuid(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccDirectoryRoleAssignment_multipleUser(t *testing.T) {
	dataA := acceptance.BuildTestData(t, "azuread_directory_role_assignment", "testA")
	dataB := acceptance.BuildTestData(t, "azuread_directory_role_assignment", "testB")
	r := DirectoryRoleAssignmentResource{}

	dataA.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.oneUser(dataA),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(dataA.ResourceName).ExistsInAzure(r),
				check.That(dataA.ResourceName).Key("role_id").IsUuid(),
				check.That(dataA.ResourceName).Key("principal_object_id").IsUuid(),
			),
		},
		dataA.ImportStep(),
		{
			Config: r.twoUsers(dataA),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(dataA.ResourceName).ExistsInAzure(r),
				check.That(dataA.ResourceName).Key("role_id").IsUuid(),
				check.That(dataA.ResourceName).Key("principal_object_id").IsUuid(),
				check.That(dataB.ResourceName).ExistsInAzure(r),
				check.That(dataB.ResourceName).Key("role_id").IsUuid(),
				check.That(dataB.ResourceName).Key("principal_object_id").IsUuid(),
			),
		},
		dataA.ImportStep(),
		dataB.ImportStep(),
		{
			Config: r.oneUser(dataA),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(dataA.ResourceName).ExistsInAzure(r),
				check.That(dataA.ResourceName).Key("role_id").IsUuid(),
				check.That(dataA.ResourceName).Key("principal_object_id").IsUuid(),
			),
		},
		dataA.ImportStep(),
	})
}

func (r DirectoryRoleAssignmentResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.DirectoryRoles.RoleAssignmentsClient
	client.BaseClient.DisableRetries = true
	defer func() { client.BaseClient.DisableRetries = false }()

	if _, status, err := client.Get(ctx, state.ID, odata.Query{}); err != nil {
		if status == http.StatusNotFound {
			return pointer.To(false), nil
		}
		return nil, fmt.Errorf("failed to retrieve directory role assignment %q: %+v", state.ID, err)
	}

	return pointer.To(true), nil
}

func (DirectoryRoleAssignmentResource) templateThreeUsers(data acceptance.TestData) string {
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

func (r DirectoryRoleAssignmentResource) servicePrincipal(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_application" "test" {
  display_name = "acctestServicePrincipal-%[2]d"
}

resource "azuread_service_principal" "test" {
  application_id = azuread_application.test.application_id
}

resource "azuread_directory_role_assignment" "test" {
  role_id             = azuread_directory_role.test.template_id
  principal_object_id = azuread_service_principal.test.object_id
}
`, DirectoryRoleResource{}.byTemplateId(data), data.RandomInteger)
}

func (r DirectoryRoleAssignmentResource) servicePrincipalScopedApplication(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_directory_role" "test" {
  display_name = "Cloud application administrator"
}

resource "azuread_application" "admin_test" {
  display_name = "acctestApplicationAdministrator-%[1]d"
}

resource "azuread_application" "test" {
  display_name = "acctestServicePrincipal-%[1]d"
}

resource "azuread_service_principal" "test" {
  application_id = azuread_application.test.application_id
}

resource "azuread_directory_role_assignment" "test" {
  role_id             = azuread_directory_role.test.template_id
  principal_object_id = azuread_service_principal.test.object_id
  directory_scope_id  = format("/%%s", azuread_application.admin_test.object_id)
}
`, data.RandomInteger)
}

func (r DirectoryRoleAssignmentResource) servicePrincipalCustomRole(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_application" "test" {
  display_name = "acctestServicePrincipal-%[2]d"
}

resource "azuread_service_principal" "test" {
  application_id = azuread_application.test.application_id
}

resource "azuread_directory_role_assignment" "test" {
  role_id             = azuread_custom_directory_role.test.object_id
  principal_object_id = azuread_service_principal.test.object_id
}
`, CustomDirectoryRoleResource{}.basic(data), data.RandomInteger)
}

func (r DirectoryRoleAssignmentResource) oneUser(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s
%[2]s

resource "azuread_directory_role_assignment" "testA" {
  role_id             = azuread_directory_role.test.template_id
  principal_object_id = azuread_user.testA.object_id
}
`, DirectoryRoleResource{}.byTemplateId(data), r.templateThreeUsers(data))
}

func (r DirectoryRoleAssignmentResource) oneUserCustomRole(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s
%[2]s

resource "azuread_directory_role_assignment" "testA" {
  role_id             = azuread_custom_directory_role.test.object_id
  principal_object_id = azuread_user.testA.object_id
}
`, CustomDirectoryRoleResource{}.basic(data), r.templateThreeUsers(data))
}

func (r DirectoryRoleAssignmentResource) twoUsers(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s
%[2]s

resource "azuread_directory_role_assignment" "testA" {
  role_id             = azuread_directory_role.test.template_id
  principal_object_id = azuread_user.testA.object_id
}

resource "azuread_directory_role_assignment" "testB" {
  role_id             = azuread_directory_role.test.template_id
  principal_object_id = azuread_user.testB.object_id
}
`, DirectoryRoleResource{}.byTemplateId(data), r.templateThreeUsers(data))
}
