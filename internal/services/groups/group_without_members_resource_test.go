// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package groups_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	groupBeta "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/group"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
)

type GroupWithoutMembersResource struct{}

func TestAccGroupWithoutMembers_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_group_without_members", "test")
	r := GroupWithoutMembersResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctestGroup-%d", data.RandomInteger)),
			),
		},
		data.ImportStep(),
	})
}

func TestAccGroupWithoutMembers_basicUnified(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_group_without_members", "test")
	r := GroupWithoutMembersResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basicUnified(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctestGroup-%d", data.RandomInteger)),
			),
		},
		data.ImportStep(),
	})
}

func TestAccGroupWithoutMembers_completeUnified(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_group_without_members", "test")
	r := GroupWithoutMembersResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.completeUnified(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccGroupWithoutMembers_updateUnified(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_group_without_members", "test")
	r := GroupWithoutMembersResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basicUnified(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.unified(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.completeUnified(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.unified(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccGroupWithoutMembers_assignableToRole(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_group_without_members", "test")
	r := GroupWithoutMembersResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.assignableToRole(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccGroupWithoutMembers_behaviors(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_group_without_members", "test")
	r := GroupWithoutMembersResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.behaviors(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccGroupWithoutMembers_dynamicMembership(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_group_without_members", "test")
	r := GroupWithoutMembersResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.dynamicMembership(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccGroupWithoutMembers_callerOwner(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_group_without_members", "test")
	r := GroupWithoutMembersResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.withCallerAsOwner(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccGroupWithoutMembers_owners(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_group_without_members", "test")
	r := GroupWithoutMembersResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("owners.#").HasValue("1"),
			),
		},
		data.ImportStep(),
		{
			Config: r.withOneOwner(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("owners.#").HasValue("1"),
			),
		},
		data.ImportStep(),
		{
			Config: r.withThreeOwners(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("owners.#").HasValue("3"),
			),
		},
		data.ImportStep(),
		{
			Config: r.withOneOwner(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("owners.#").HasValue("1"),
			),
		},
		data.ImportStep(),
		{
			Config: r.withServicePrincipalOwner(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("owners.#").HasValue("1"),
			),
		},
		data.ImportStep(),
		{
			Config: r.withDiverseOwners(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("owners.#").HasValue("2"),
			),
		},
		data.ImportStep(),
		{
			Config: r.removeOwners(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("owners.#").HasValue("0"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccGroupWithoutMembers_preventDuplicateNamesPass(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_group_without_members", "test")
	r := GroupWithoutMembersResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.preventDuplicateNamesPass(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctestGroup-%d", data.RandomInteger)),
			),
		},
		data.ImportStep("prevent_duplicate_names"),
	})
}

func TestAccGroupWithoutMembers_preventDuplicateNamesForceNew(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_group_without_members", "test")
	r := GroupWithoutMembersResource{}

	data.ResourceTestIgnoreRecreate(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		{
			Config: r.preventDuplicateNamesForceNew(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctestGroup-%d", data.RandomInteger)),
			),
		},
		data.ImportStep("prevent_duplicate_names"),
	})
}

func TestAccGroupWithoutMembers_provisioning(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_group_without_members", "test")
	r := GroupWithoutMembersResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.provisioning(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccGroupWithoutMembers_unifiedExtraSettings(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_group_without_members", "test")
	r := GroupWithoutMembersResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.unifiedWithExtraSettings(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.unifiedAsUser(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.unifiedWithExtraSettings(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccGroupWithoutMembers_preventDuplicateNamesFail(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_group", "test")
	r := GroupWithoutMembersResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		data.RequiresImportErrorStep(r.preventDuplicateNamesFail(data)),
	})
}

func TestAccGroupWithoutMembers_visibility(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_group_without_members", "test")
	r := GroupWithoutMembersResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.visibility(data, "Private"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.visibility(data, "Public"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccGroupWithoutMembers_administrativeUnit(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_group_without_members", "test")
	r := GroupWithoutMembersResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.administrativeUnits(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("administrative_unit_ids.#").HasValue("2"),
			),
		},
		data.ImportStep("administrative_unit_ids"),
		{
			Config: r.administrativeUnitsWithoutAssociation(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("administrative_unit_ids.#").HasValue("0"),
			),
		},
		data.ImportStep("administrative_unit_ids"),
		{
			Config: r.administrativeUnits(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("administrative_unit_ids.#").HasValue("2"),
			),
		},
		data.ImportStep("administrative_unit_ids"),
	})
}

func TestAccGroupWithoutMembers_writeback(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_group_without_members", "test")
	r := GroupWithoutMembersResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.withWriteback(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("onpremises_group_type").HasValue("UniversalSecurityGroup"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccGroupWithoutMembers_writebackUpdate(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_group_without_members", "test")
	r := GroupWithoutMembersResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.withWriteback(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("onpremises_group_type").HasValue("UniversalSecurityGroup"),
			),
		},
		data.ImportStep(),
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccGroupWithoutMembers_writebackUnified(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_group_without_members", "test")
	r := GroupWithoutMembersResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.unifiedWithWriteback(data, "UniversalDistributionGroup"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("onpremises_group_type").HasValue("UniversalDistributionGroup"),
			),
		},
		data.ImportStep(),
		{
			Config: r.unifiedWithWriteback(data, "UniversalMailEnabledSecurityGroup"),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("onpremises_group_type").HasValue("UniversalMailEnabledSecurityGroup"),
			),
		},
		data.ImportStep(),
	})
}

func (r GroupWithoutMembersResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.Groups.GroupClientBeta

	id, err := beta.ParseGroupID(state.ID)
	if err != nil {
		return nil, err
	}

	resp, err := client.GetGroup(ctx, *id, groupBeta.DefaultGetGroupOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return pointer.To(false), nil
		}
		return nil, fmt.Errorf("failed to retrieve %s: %v", id, err)
	}
	return pointer.To(true), nil
}

func (GroupWithoutMembersResource) templateDiverseDirectoryObjects(data acceptance.TestData) string {
	return fmt.Sprintf(`
data "azuread_domains" "test" {
  only_initial = true
}

resource "azuread_application" "test" {
  display_name = "acctestGroup-%[1]d"
}

resource "azuread_service_principal" "test" {
  client_id = azuread_application.test.client_id
}

resource "azuread_user" "test" {
  user_principal_name = "acctestGroup.%[1]d@${data.azuread_domains.test.domains.0.domain_name}"
  display_name        = "acctestGroup-%[1]d"
  password            = "%[2]s"
}

resource "azuread_group_without_members" "member" {
  display_name     = "acctestGroup-%[1]d-Member"
  security_enabled = true
}
`, data.RandomInteger, data.RandomPassword)
}

func (GroupWithoutMembersResource) templateThreeUsers(data acceptance.TestData) string {
	return fmt.Sprintf(`
data "azuread_domains" "test" {
  only_initial = true
}

resource "azuread_user" "testA" {
  user_principal_name = "acctestGroup.%[1]d.A@${data.azuread_domains.test.domains.0.domain_name}"
  display_name        = "acctestGroup-%[1]d-A"
  password            = "%[2]s"
}

resource "azuread_user" "testB" {
  user_principal_name = "acctestGroup.%[1]d.B@${data.azuread_domains.test.domains.0.domain_name}"
  display_name        = "acctestGroup-%[1]d-B"
  mail_nickname       = "acctestGroup-%[1]d-B"
  password            = "%[2]s"
}

resource "azuread_user" "testC" {
  user_principal_name = "acctestGroup.%[1]d.C@${data.azuread_domains.test.domains.0.domain_name}"
  display_name        = "acctestGroup-%[1]d-C"
  password            = "%[2]s"
}
`, data.RandomInteger, data.RandomPassword)
}

func (GroupWithoutMembersResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_group_without_members" "test" {
  display_name     = "acctestGroup-%[1]d"
  security_enabled = true
}
`, data.RandomInteger)
}

func (GroupWithoutMembersResource) removeOwners(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_group_without_members" "test" {
  display_name     = "acctestGroup-%[1]d"
  security_enabled = true
  owners           = []
}
`, data.RandomInteger)
}

func (GroupWithoutMembersResource) basicUnified(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_group_without_members" "test" {
  display_name     = "acctestGroup-%[1]d"
  types            = ["Unified"]
  mail_enabled     = true
  mail_nickname    = "acctestGroup-%[1]d"
  security_enabled = false
}
`, data.RandomInteger)
}

func (GroupWithoutMembersResource) unified(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_group_without_members" "test" {
  display_name     = "acctestGroup-%[1]d"
  description      = "Please delete me as this is a.test.AD group!"
  types            = ["Unified"]
  mail_enabled     = true
  mail_nickname    = "acctest.Group-%[1]d"
  security_enabled = true
  theme            = "Pink"
}
`, data.RandomInteger)
}

func (r GroupWithoutMembersResource) unifiedAsUser(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {
  client_id           = ""
  client_id_file_path = ""
  use_cli             = true
}

%[1]s
`, r.unified(data))
}

func (GroupWithoutMembersResource) unifiedWithExtraSettings(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {
  client_id           = ""
  client_id_file_path = ""
  use_cli             = true
}

resource "azuread_group_without_members" "test" {
  display_name     = "acctestGroup-%[1]d"
  description      = "Please delete me as this is a.test.AD group!"
  types            = ["Unified"]
  mail_enabled     = true
  mail_nickname    = "acctest.Group-%[1]d"
  security_enabled = true
  theme            = "Pink"

  auto_subscribe_new_members = true
  external_senders_allowed   = true
  hide_from_address_lists    = true
  hide_from_outlook_clients  = true
}
`, data.RandomInteger)
}

func (GroupWithoutMembersResource) unifiedWithWriteback(data acceptance.TestData, onPremisesGroupType string) string {
	return fmt.Sprintf(`
resource "azuread_group_without_members" "test" {
  display_name     = "acctestGroup-%[1]d"
  description      = "Please delete me as this is a.test.AD group!"
  types            = ["Unified"]
  mail_enabled     = true
  mail_nickname    = "acctest.Group-%[1]d"
  security_enabled = true
  theme            = "Pink"

  writeback_enabled     = true
  onpremises_group_type = %[2]q
}
`, data.RandomInteger, onPremisesGroupType)
}

func (GroupWithoutMembersResource) completeUnified(data acceptance.TestData) string {
	return fmt.Sprintf(`
data "azuread_domains" "test" {
  only_initial = true
}

data "azuread_client_config" "test" {}

resource "azuread_user" "test" {
  user_principal_name = "acctestGroup.%[1]d@${data.azuread_domains.test.domains.0.domain_name}"
  display_name        = "acctestGroup-%[1]d"
  password            = "%[2]s"
}

resource "azuread_group_without_members" "test" {
  description      = "Please delete me as this is a.test.AD group!"
  display_name     = "acctestGroup-complete-%[1]d"
  types            = ["Unified"]
  mail_enabled     = true
  mail_nickname    = "acctest.Group-%[1]d"
  security_enabled = true
  owners           = [azuread_user.test.object_id]
  theme            = "Purple"
}
`, data.RandomInteger, data.RandomPassword)
}

func (GroupWithoutMembersResource) assignableToRole(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_group_without_members" "test" {
  assignable_to_role = true
  display_name       = "acctestGroup-assignableToRole-%[1]d"
  security_enabled   = true
}
`, data.RandomInteger)
}

func (GroupWithoutMembersResource) behaviors(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_group_without_members" "test" {
  display_name  = "acctestGroup-behaviors-%[1]d"
  mail_enabled  = true
  mail_nickname = "acctestGroup-behaviors-%[1]d"
  types         = ["Unified"]

  behaviors = [
    "AllowOnlyMembersToPost",
    "HideGroupInOutlook",
    "SubscribeNewGroupMembers",
    "WelcomeEmailDisabled"
  ]
}
`, data.RandomInteger)
}

func (GroupWithoutMembersResource) dynamicMembership(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_group_without_members" "test" {
  display_name     = "acctestGroup-%[1]d"
  description      = "Please delete me as this is a.test.AD group!"
  types            = ["DynamicMembership", "Unified"]
  mail_enabled     = true
  mail_nickname    = "acctest.Group-%[1]d"
  security_enabled = true

  dynamic_membership {
    enabled = true
    rule    = "user.department -eq \"Sales\""
  }
}
`, data.RandomInteger)
}

func (GroupWithoutMembersResource) provisioning(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_group_without_members" "test" {
  display_name  = "acctestGroup-behaviors-%[1]d"
  mail_enabled  = true
  mail_nickname = "acctestGroup-behaviors-%[1]d"
  types         = ["Unified"]

  provisioning_options = ["Team"]
}
`, data.RandomInteger)
}

func (GroupWithoutMembersResource) visibility(data acceptance.TestData, visibility string) string {
	return fmt.Sprintf(`
resource "azuread_group_without_members" "test" {
  display_name  = "acctestGroup-visibility-%[1]d"
  mail_enabled  = true
  mail_nickname = "acctestGroup-visibility-%[1]d"
  types         = ["Unified"]
  visibility    = "%[2]s"
}
`, data.RandomInteger, visibility)
}

func (r GroupWithoutMembersResource) withOneOwner(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_group_without_members" "test" {
  display_name     = "acctestGroup-%[2]d"
  security_enabled = true
  owners           = [azuread_user.testA.object_id]
}
`, r.templateThreeUsers(data), data.RandomInteger)
}

func (GroupWithoutMembersResource) withCallerAsOwner(data acceptance.TestData) string {
	return fmt.Sprintf(`
data "azuread_client_config" "test" {}

resource "azuread_group_without_members" "test" {
  display_name     = "acctestGroup-%[1]d"
  security_enabled = true
  owners           = [data.azuread_client_config.test.object_id]
}
`, data.RandomInteger)
}

func (GroupWithoutMembersResource) withServicePrincipalOwner(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  display_name = "acctestGroup-%[1]d"
}

resource "azuread_service_principal" "test" {
  client_id = azuread_application.test.client_id
}

resource "azuread_group_without_members" "test" {
  display_name     = "acctestGroup-%[1]d"
  security_enabled = true
  owners           = [azuread_service_principal.test.object_id]
}
`, data.RandomInteger)
}

func (r GroupWithoutMembersResource) withDiverseOwners(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_group_without_members" "test" {
  display_name     = "acctestGroup-%[2]d"
  security_enabled = true
  owners = [
    azuread_service_principal.test.object_id,
    azuread_user.test.object_id,
  ]
}
`, r.templateDiverseDirectoryObjects(data), data.RandomInteger)
}

func (r GroupWithoutMembersResource) withThreeOwners(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_group_without_members" "test" {
  display_name     = "acctestGroup-%[2]d"
  security_enabled = true
  owners = [
    azuread_user.testA.object_id,
    azuread_user.testB.object_id,
    azuread_user.testC.object_id
  ]
}
`, r.templateThreeUsers(data), data.RandomInteger)
}

func (GroupWithoutMembersResource) preventDuplicateNamesPass(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_group_without_members" "test" {
  display_name            = "acctestGroup-%[1]d"
  security_enabled        = true
  prevent_duplicate_names = true
}
`, data.RandomInteger)
}

func (r GroupWithoutMembersResource) preventDuplicateNamesFail(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_group_without_members" "duplicate" {
  display_name            = azuread_group_without_members.test.display_name
  security_enabled        = true
  prevent_duplicate_names = true
}
`, r.basic(data))
}

func (GroupWithoutMembersResource) preventDuplicateNamesForceNew(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_group_without_members" "test" {
  display_name            = "acctestGroup-%[1]d"
  security_enabled        = true
  prevent_duplicate_names = true

  assignable_to_role = true
}
`, data.RandomInteger)
}

func (r GroupWithoutMembersResource) administrativeUnits(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_administrative_unit" "test" {
  display_name = "acctestGroup-administrative-unit-%[1]d"
}

resource "azuread_administrative_unit" "test2" {
  display_name = "acctestGroup-administrative-unit-%[1]d"
}

resource "azuread_group_without_members" "test" {
  display_name            = "acctestGroup-%[1]d"
  security_enabled        = true
  administrative_unit_ids = [azuread_administrative_unit.test.object_id, azuread_administrative_unit.test2.object_id]
}
`, data.RandomInteger, data.RandomInteger)
}

func (r GroupWithoutMembersResource) administrativeUnitsWithoutAssociation(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_administrative_unit" "test" {
  display_name = "acctestGroup-administrative-unit-%[1]d"
}

resource "azuread_administrative_unit" "test2" {
  display_name = "acctestGroup-administrative-unit-%[1]d"
}

resource "azuread_group_without_members" "test" {
  display_name     = "acctestGroup-%[1]d"
  security_enabled = true
}
`, data.RandomInteger, data.RandomInteger)
}

func (GroupWithoutMembersResource) withWriteback(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_group_without_members" "test" {
  display_name     = "acctestGroup-%[1]d"
  security_enabled = true

  writeback_enabled     = true
  onpremises_group_type = "UniversalSecurityGroup"
}
`, data.RandomInteger)
}
