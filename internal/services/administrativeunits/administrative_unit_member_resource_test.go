package administrativeunits_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/administrativeunits/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
)

type AdministrativeUnitMemberResource struct{}

func TestAccAdministrativeUnitMember_group(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_administrative_unit_member", "test")
	r := AdministrativeUnitMemberResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.group(data),
			Check: resource.ComposeTestCheckFunc(
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

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.oneUser(data),
			Check: resource.ComposeTestCheckFunc(
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

	dataA.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.oneUser(dataA),
			Check: resource.ComposeTestCheckFunc(
				check.That(dataA.ResourceName).ExistsInAzure(r),
				check.That(dataA.ResourceName).Key("administrative_unit_object_id").IsUuid(),
				check.That(dataA.ResourceName).Key("member_object_id").IsUuid(),
			),
		},
		dataA.ImportStep(),
		{
			Config: r.twoUsers(dataA),
			Check: resource.ComposeTestCheckFunc(
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
			Check: resource.ComposeTestCheckFunc(
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

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.group(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.RequiresImportErrorStep(r.requiresImport(data)),
	})
}

func (r AdministrativeUnitMemberResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.AdministrativeUnits.AdministrativeUnitsClient
	client.BaseClient.DisableRetries = true

	id, err := parse.AdministrativeUnitMemberID(state.ID)
	if err != nil {
		return nil, fmt.Errorf("parsing Administrative Unit Member ID: %v", err)
	}

	if _, status, err := client.GetMember(ctx, id.AdministrativeUnitId, id.MemberId); err != nil {
		if status == http.StatusNotFound {
			return utils.Bool(false), nil
		}
		return nil, fmt.Errorf("failed to retrieve administrative unit member %q (administrative unit ID: %q): %+v", id.MemberId, id.AdministrativeUnitId, err)
	}

	return utils.Bool(true), nil
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
