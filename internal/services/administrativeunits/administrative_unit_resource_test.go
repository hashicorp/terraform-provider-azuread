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
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/stable/administrativeunit"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
)

type AdministrativeUnitResource struct{}

func TestAccAdministrativeUnit_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_administrative_unit", "test")
	r := AdministrativeUnitResource{}

	data.ResourceTestIgnoreDangling(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("object_id").IsUuid(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccAdministrativeUnit_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_administrative_unit", "test")
	r := AdministrativeUnitResource{}

	data.ResourceTestIgnoreDangling(t, r, []acceptance.TestStep{
		{
			Config: r.complete(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("object_id").IsUuid(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccAdministrativeUnit_withMembers(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_administrative_unit", "test")
	r := AdministrativeUnitResource{}

	data.ResourceTestIgnoreDangling(t, r, []acceptance.TestStep{
		{
			Config: r.withMembers(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("object_id").IsUuid(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccGroup_preventDuplicateNamesPass(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_administrative_unit", "test")
	r := AdministrativeUnitResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.preventDuplicateNamesPass(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctestAdministrativeUnit-%d", data.RandomInteger)),
			),
		},
		data.ImportStep("prevent_duplicate_names"),
	})
}

func TestAccGroup_preventDuplicateNamesFail(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_administrative_unit", "test")
	r := AdministrativeUnitResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		data.RequiresImportErrorStep(r.preventDuplicateNamesFail(data)),
	})
}

func (r AdministrativeUnitResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.AdministrativeUnits.AdministrativeUnitClient

	id, err := stable.ParseDirectoryAdministrativeUnitID(state.ID)
	if err != nil {
		return nil, err
	}

	resp, err := client.GetAdministrativeUnit(ctx, *id, administrativeunit.DefaultGetAdministrativeUnitOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return nil, fmt.Errorf("%s does not exist", id)
		}
		return nil, fmt.Errorf("failed to retrieve %s: %v", id, err)
	}
	return pointer.To(resp.Model != nil), nil
}

func (AdministrativeUnitResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_administrative_unit" "test" {
  display_name = "acctestAdministrativeUnit-%[1]d"
}
`, data.RandomInteger)
}

func (AdministrativeUnitResource) complete(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_administrative_unit" "test" {
  display_name = "acctestAdministrativeUnit-%[1]d"
  description  = "testing administrative units %[2]s"

  hidden_membership_enabled = true
}
`, data.RandomInteger, data.RandomString)
}

func (AdministrativeUnitResource) withMembers(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

data "azuread_domains" "test" {
  only_initial = true
}

resource "azuread_group" "member" {
  display_name     = "acctest-AdministrativeUnitMember-%[1]d"
  security_enabled = true
  lifecycle {
    ignore_changes = [administrative_unit_ids]
  }
}

resource "azuread_user" "memberA" {
  user_principal_name = "acctestAdministrativeUnitMember.%[1]d.A@${data.azuread_domains.test.domains.0.domain_name}"
  display_name        = "acctestAdministrativeUnitMember-%[1]d-A"
  password            = "%[2]s"
}

resource "azuread_user" "memberB" {
  user_principal_name = "acctestAdministrativeUnitMember.%[1]d.B@${data.azuread_domains.test.domains.0.domain_name}"
  display_name        = "acctestAdministrativeUnitMember-%[1]d-B"
  password            = "%[2]s"
}

resource "azuread_user" "memberC" {
  user_principal_name = "acctestAdministrativeUnitMember.%[1]d.C@${data.azuread_domains.test.domains.0.domain_name}"
  display_name        = "acctestAdministrativeUnitMember-%[1]d-C"
  password            = "%[2]s"
}

resource "azuread_administrative_unit" "test" {
  display_name = "acctestAdministrativeUnit-%[1]d"
  members = [
    azuread_group.member.object_id,
    azuread_user.memberA.object_id,
    azuread_user.memberB.object_id,
    azuread_user.memberC.object_id,
  ]
}
`, data.RandomInteger, data.RandomPassword)
}

func (AdministrativeUnitResource) preventDuplicateNamesPass(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_administrative_unit" "test" {
  display_name            = "acctestAdministrativeUnit-%[1]d"
  prevent_duplicate_names = true
}
`, data.RandomInteger)
}

func (r AdministrativeUnitResource) preventDuplicateNamesFail(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_administrative_unit" "duplicate" {
  display_name            = azuread_administrative_unit.test.display_name
  prevent_duplicate_names = true
}
`, r.basic(data))
}
