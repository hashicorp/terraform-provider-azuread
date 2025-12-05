// Copyright IBM Corp. 2019, 2025
// SPDX-License-Identifier: MPL-2.0

package administrativeunits_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/stable/administrativeunitscopedrolemember"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
)

type AdministrativeUnitRoleMemberResource struct{}

func TestAccAdministrativeUnitRoleMember_user(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_administrative_unit_role_member", "test")
	r := AdministrativeUnitRoleMemberResource{}

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

func TestAccAdministrativeUnitRoleMember_multipleUser(t *testing.T) {
	dataA := acceptance.BuildTestData(t, "azuread_administrative_unit_role_member", "testA")
	dataB := acceptance.BuildTestData(t, "azuread_administrative_unit_role_member", "testB")
	dataC := acceptance.BuildTestData(t, "azuread_administrative_unit_role_member", "testC")
	r := AdministrativeUnitRoleMemberResource{}

	dataA.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.threeUsers(dataA),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(dataA.ResourceName).ExistsInAzure(r),
				check.That(dataA.ResourceName).Key("role_object_id").IsUuid(),
				check.That(dataA.ResourceName).Key("member_object_id").IsUuid(),
				check.That(dataB.ResourceName).ExistsInAzure(r),
				check.That(dataB.ResourceName).Key("role_object_id").IsUuid(),
				check.That(dataB.ResourceName).Key("member_object_id").IsUuid(),
				check.That(dataC.ResourceName).ExistsInAzure(r),
				check.That(dataC.ResourceName).Key("role_object_id").IsUuid(),
				check.That(dataC.ResourceName).Key("member_object_id").IsUuid(),
			),
		},
		dataA.ImportStep(),
		dataB.ImportStep(),
		dataC.ImportStep(),
	})
}

func TestAccAdministrativeUnitRoleMember_group(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_administrative_unit_role_member", "test")
	r := AdministrativeUnitRoleMemberResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.group(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("role_object_id").IsUuid(),
				check.That(data.ResourceName).Key("member_object_id").IsUuid(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccAdministrativeUnitRoleMember_servicePrincipal(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_administrative_unit_role_member", "test")
	r := AdministrativeUnitRoleMemberResource{}

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

func (r AdministrativeUnitRoleMemberResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.AdministrativeUnits.AdministrativeUnitScopedRoleMemberClient

	id, err := stable.ParseDirectoryAdministrativeUnitIdScopedRoleMemberID(state.ID)
	if err != nil {
		return nil, fmt.Errorf("parsing Directory Role Member ID: %v", err)
	}

	if resp, err := client.GetAdministrativeUnitScopedRoleMember(ctx, stable.NewDirectoryAdministrativeUnitIdScopedRoleMemberID(id.AdministrativeUnitId, id.ScopedRoleMembershipId), administrativeunitscopedrolemember.DefaultGetAdministrativeUnitScopedRoleMemberOperationOptions()); err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return pointer.To(false), nil
		}
		return nil, fmt.Errorf("failed to retrieve administrative unit role membership %q (AU ID: %q): %+v", id.ScopedRoleMembershipId, id.AdministrativeUnitId, err)
	}

	return pointer.To(true), nil
}

func (AdministrativeUnitRoleMemberResource) templateThreeUsers(data acceptance.TestData) string {
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

func (AdministrativeUnitRoleMemberResource) templateGroup(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_group" "test" {
  display_name       = "acctestGroup-%[1]d"
  assignable_to_role = true
  security_enabled   = true
}
`, data.RandomInteger)
}

func (AdministrativeUnitRoleMemberResource) templateServicePrincipal(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  display_name = "acctestServicePrincipal-%[1]d"
}

resource "azuread_service_principal" "test" {
  client_id = azuread_application.test.client_id
}
`, data.RandomInteger)
}

func (AdministrativeUnitRoleMemberResource) roleByTemplateId(_ acceptance.TestData) string {
	return `
resource "azuread_directory_role" "test" {
  template_id = "644ef478-e28f-4e28-b9dc-3fdde9aa0b1f" // Printer administrator
}
`
}

func (r AdministrativeUnitRoleMemberResource) oneUser(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s
%[2]s
%[3]s

resource "azuread_administrative_unit_role_member" "test" {
  role_object_id                = azuread_directory_role.test.object_id
  member_object_id              = azuread_user.testA.object_id
  administrative_unit_object_id = azuread_administrative_unit.test.object_id
}
`, AdministrativeUnitRoleMemberResource{}.roleByTemplateId(data), r.templateThreeUsers(data), AdministrativeUnitResource{}.basic(data))
}

func (r AdministrativeUnitRoleMemberResource) threeUsers(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s
%[2]s
%[3]s

resource "azuread_administrative_unit_role_member" "testA" {
  role_object_id                = azuread_directory_role.test.object_id
  member_object_id              = azuread_user.testA.object_id
  administrative_unit_object_id = azuread_administrative_unit.test.object_id
}

resource "azuread_administrative_unit_role_member" "testB" {
  role_object_id                = azuread_directory_role.test.object_id
  member_object_id              = azuread_user.testB.object_id
  administrative_unit_object_id = azuread_administrative_unit.test.object_id
}

resource "azuread_administrative_unit_role_member" "testC" {
  role_object_id                = azuread_directory_role.test.object_id
  member_object_id              = azuread_user.testC.object_id
  administrative_unit_object_id = azuread_administrative_unit.test.object_id
}
`, AdministrativeUnitRoleMemberResource{}.roleByTemplateId(data), r.templateThreeUsers(data), AdministrativeUnitResource{}.basic(data))
}

func (r AdministrativeUnitRoleMemberResource) group(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s
%[2]s
%[3]s

resource "azuread_administrative_unit_role_member" "test" {
  role_object_id                = azuread_directory_role.test.object_id
  member_object_id              = azuread_group.test.object_id
  administrative_unit_object_id = azuread_administrative_unit.test.object_id
}
`, AdministrativeUnitRoleMemberResource{}.roleByTemplateId(data), r.templateGroup(data), AdministrativeUnitResource{}.basic(data))
}

func (r AdministrativeUnitRoleMemberResource) servicePrincipal(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s
%[2]s
%[3]s

resource "azuread_administrative_unit_role_member" "test" {
  role_object_id                = azuread_directory_role.test.object_id
  member_object_id              = azuread_service_principal.test.object_id
  administrative_unit_object_id = azuread_administrative_unit.test.object_id
}
`, AdministrativeUnitRoleMemberResource{}.roleByTemplateId(data), r.templateServicePrincipal(data), AdministrativeUnitResource{}.basic(data))
}
