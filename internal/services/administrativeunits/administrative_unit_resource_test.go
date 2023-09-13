// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package administrativeunits_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
)

type AdministrativeUnitResource struct{}

func TestAccAdministrativeUnit_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_administrative_unit", "test")
	r := AdministrativeUnitResource{}

	data.ResourceTestIgnoreDangling(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
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

	data.ResourceTestIgnoreDangling(t, r, []resource.TestStep{
		{
			Config: r.complete(data),
			Check: resource.ComposeTestCheckFunc(
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

	data.ResourceTestIgnoreDangling(t, r, []resource.TestStep{
		{
			Config: r.withMembers(data),
			Check: resource.ComposeTestCheckFunc(
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

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.preventDuplicateNamesPass(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctestAdministrativeUnit-%d", data.RandomInteger)),
			),
		},
		data.ImportStep("prevent_duplicate_names"),
	})
}

func TestAccGroup_preventDuplicateNamesFail(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_administrative_unit", "test")
	r := AdministrativeUnitResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		data.RequiresImportErrorStep(r.preventDuplicateNamesFail(data)),
	})
}

func (r AdministrativeUnitResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.AdministrativeUnits.AdministrativeUnitsClient
	client.BaseClient.DisableRetries = true
	defer func() { client.BaseClient.DisableRetries = false }()

	role, status, err := client.Get(ctx, state.ID, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return nil, fmt.Errorf("Administratove Unit with object ID %q does not exist", state.ID)
		}
		return nil, fmt.Errorf("failed to retrieve Administratove Unit with object ID %q: %+v", state.ID, err)
	}
	return utils.Bool(role.ID != nil && *role.ID == state.ID), nil
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
