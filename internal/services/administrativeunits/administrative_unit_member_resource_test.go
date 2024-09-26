// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package administrativeunits_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/stable/administrativeunitmember"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
)

type AdministrativeUnitMemberResource struct{}

func TestAccAdministrativeUnitMember_group(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_administrative_unit_member", "test")
	r := AdministrativeUnitMemberResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.group(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("administrative_unit_object_id").IsUuid(),
				check.That(data.ResourceName).Key("member_object_id").IsUuid(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccAdministrativeUnitMember_user(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_administrative_unit_member", "testA")
	r := AdministrativeUnitMemberResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.oneUser(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("administrative_unit_object_id").IsUuid(),
				check.That(data.ResourceName).Key("member_object_id").IsUuid(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccAdministrativeUnitMember_multipleUsers(t *testing.T) {
	dataA := acceptance.BuildTestData(t, "azuread_administrative_unit_member", "testA")
	dataB := acceptance.BuildTestData(t, "azuread_administrative_unit_member", "testB")
	r := AdministrativeUnitMemberResource{}

	dataA.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.oneUser(dataA),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(dataA.ResourceName).ExistsInAzure(r),
				check.That(dataA.ResourceName).Key("administrative_unit_object_id").IsUuid(),
				check.That(dataA.ResourceName).Key("member_object_id").IsUuid(),
			),
		},
		dataA.ImportStep(),
		{
			Config: r.twoUsers(dataA),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(dataA.ResourceName).ExistsInAzure(r),
				check.That(dataA.ResourceName).Key("administrative_unit_object_id").IsUuid(),
				check.That(dataA.ResourceName).Key("member_object_id").IsUuid(),
				check.That(dataB.ResourceName).ExistsInAzure(r),
				check.That(dataB.ResourceName).Key("administrative_unit_object_id").IsUuid(),
				check.That(dataB.ResourceName).Key("member_object_id").IsUuid(),
			),
		},
		dataA.ImportStep(),
		dataB.ImportStep(),
		{
			Config: r.oneUser(dataA),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(dataA.ResourceName).ExistsInAzure(r),
				check.That(dataA.ResourceName).Key("administrative_unit_object_id").IsUuid(),
				check.That(dataA.ResourceName).Key("member_object_id").IsUuid(),
			),
		},
		dataA.ImportStep(),
	})
}

func TestAccAdministrativeUnitMember_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_administrative_unit_member", "test")
	r := AdministrativeUnitMemberResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.group(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.RequiresImportErrorStep(r.requiresImport(data)),
	})
}

func (r AdministrativeUnitMemberResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.AdministrativeUnits.AdministrativeUnitMemberClient

	id, err := stable.ParseDirectoryAdministrativeUnitIdMemberID(state.ID)
	if err != nil {
		return nil, fmt.Errorf("parsing Administrative Unit Member ID: %v", err)
	}

	options := administrativeunitmember.ListAdministrativeUnitMembersOperationOptions{
		Filter: pointer.To(fmt.Sprintf("id eq '%s'", id.DirectoryObjectId)),
	}
	resp, err := client.ListAdministrativeUnitMembers(ctx, stable.NewDirectoryAdministrativeUnitID(id.AdministrativeUnitId), options)
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return pointer.To(false), nil
		}
		return nil, fmt.Errorf("failed to retrieve administrative unit member %q (administrative unit ID: %q): %+v", id.DirectoryObjectId, id.AdministrativeUnitId, err)
	}

	if resp.Model != nil {
		for _, member := range *resp.Model {
			if pointer.From(member.DirectoryObject().Id) == id.DirectoryObjectId {
				return pointer.To(true), nil
			}
		}
	}

	return pointer.To(false), nil
}

func (AdministrativeUnitMemberResource) templateThreeUsers(data acceptance.TestData) string {
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

func (r AdministrativeUnitMemberResource) group(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_group" "member" {
  display_name     = "acctest-AdministrativeUnitMember-%[2]d"
  security_enabled = true
  lifecycle {
    ignore_changes = [administrative_unit_ids]
  }
}

resource "azuread_administrative_unit_member" "test" {
  administrative_unit_object_id = azuread_administrative_unit.test.object_id
  member_object_id              = azuread_group.member.object_id
}
`, AdministrativeUnitResource{}.basic(data), data.RandomInteger)
}

func (r AdministrativeUnitMemberResource) oneUser(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s
%[2]s

resource "azuread_administrative_unit_member" "testA" {
  administrative_unit_object_id = azuread_administrative_unit.test.object_id
  member_object_id              = azuread_user.testA.object_id
}
`, AdministrativeUnitResource{}.basic(data), r.templateThreeUsers(data))
}

func (r AdministrativeUnitMemberResource) twoUsers(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s
%[2]s

resource "azuread_administrative_unit_member" "testA" {
  administrative_unit_object_id = azuread_administrative_unit.test.object_id
  member_object_id              = azuread_user.testA.object_id
}

resource "azuread_administrative_unit_member" "testB" {
  administrative_unit_object_id = azuread_administrative_unit.test.object_id
  member_object_id              = azuread_user.testB.object_id
}
`, AdministrativeUnitResource{}.basic(data), r.templateThreeUsers(data))
}

func (r AdministrativeUnitMemberResource) requiresImport(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_administrative_unit_member" "import" {
  administrative_unit_object_id = azuread_administrative_unit_member.test.administrative_unit_object_id
  member_object_id              = azuread_administrative_unit_member.test.member_object_id
}
`, r.group(data))
}
