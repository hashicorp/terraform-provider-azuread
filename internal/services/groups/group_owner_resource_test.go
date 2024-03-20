package groups_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/groups/parse"
)

type GroupOwnerResource struct{}

func TestAccGroupOwner_group(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_group_owner", "test")
	r := GroupOwnerResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.group(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("group_object_id").IsUuid(),
				check.That(data.ResourceName).Key("owner_object_id").IsUuid(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccGroupOwner_servicePrincipal(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_group_owner", "test")
	r := GroupOwnerResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.servicePrincipal(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("group_object_id").IsUuid(),
				check.That(data.ResourceName).Key("owner_object_id").IsUuid(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccGroupOwner_user(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_group_owner", "testA")
	r := GroupOwnerResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.oneUser(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("group_object_id").IsUuid(),
				check.That(data.ResourceName).Key("owner_object_id").IsUuid(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccGroupOwner_multipleUser(t *testing.T) {
	dataA := acceptance.BuildTestData(t, "azuread_group_owner", "testA")
	dataB := acceptance.BuildTestData(t, "azuread_group_owner", "testB")
	r := GroupOwnerResource{}

	dataA.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.oneUser(dataA),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(dataA.ResourceName).ExistsInAzure(r),
				check.That(dataA.ResourceName).Key("group_object_id").IsUuid(),
				check.That(dataA.ResourceName).Key("owner_object_id").IsUuid(),
			),
		},
		dataA.ImportStep(),
		{
			Config: r.twoUsers(dataA),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(dataA.ResourceName).ExistsInAzure(r),
				check.That(dataA.ResourceName).Key("group_object_id").IsUuid(),
				check.That(dataA.ResourceName).Key("owner_object_id").IsUuid(),
				check.That(dataB.ResourceName).ExistsInAzure(r),
				check.That(dataB.ResourceName).Key("group_object_id").IsUuid(),
				check.That(dataB.ResourceName).Key("owner_object_id").IsUuid(),
			),
		},
		// we rerun the config so the group resource updates with the number of owners
		{
			Config: r.twoUsers(dataA),
			Check: acceptance.ComposeTestCheckFunc(
				check.That("azuread_group.test").Key("owners.#").HasValue("2"),
			),
		},
		dataA.ImportStep(),
		{
			Config: r.oneUser(dataA),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(dataA.ResourceName).ExistsInAzure(r),
				check.That(dataA.ResourceName).Key("group_object_id").IsUuid(),
				check.That(dataA.ResourceName).Key("owner_object_id").IsUuid(),
			),
		},
		// we rerun the config so the group resource updates with the number of owners
		{
			Config: r.oneUser(dataA),
			Check: acceptance.ComposeTestCheckFunc(
				check.That("azuread_group.test").Key("owners.#").HasValue("1"),
			),
		},
	})
}

func TestAccGroupMember_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_group_owner", "test")
	r := GroupOwnerResource{}

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

func (r GroupOwnerResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.Groups.GroupsClient
	client.BaseClient.DisableRetries = true
	defer func() { client.BaseClient.DisableRetries = false }()

	id, err := parse.GroupMemberID(state.ID)
	if err != nil {
		return nil, fmt.Errorf("parsing Group Member ID: %v", err)
	}

	owners, _, err := client.ListMembers(ctx, id.GroupId)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve Group owners (groupId: %q): %+v", id.GroupId, err)
	}

	if owners != nil {
		for _, objectId := range *owners {
			if strings.EqualFold(objectId, id.MemberId) {
				return pointer.To(true), nil
			}
		}
	}

	return nil, fmt.Errorf("Member %q was not found in Group %q", id.MemberId, id.GroupId)
}

func (GroupOwnerResource) template(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_group" "test" {
  display_name     = "acctestGroup-%[1]d"
  security_enabled = true
}
`, data.RandomInteger)
}

func (GroupOwnerResource) templateThreeUsers(data acceptance.TestData) string {
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

func (r GroupOwnerResource) group(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_group" "owner" {
  display_name     = "acctestGroup-%[2]d-Member"
  security_enabled = true
}

resource "azuread_group_owner" "test" {
  group_object_id  = azuread_group.test.object_id
  owner_object_id = azuread_group.owner.object_id
}
`, r.template(data), data.RandomInteger)
}

func (r GroupOwnerResource) servicePrincipal(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_application" "test" {
  display_name = "acctestServicePrincipal-%[2]d"
}

resource "azuread_service_principal" "test" {
  application_id = azuread_application.test.application_id
}

resource "azuread_group_owner" "test" {
  group_object_id  = azuread_group.test.object_id
  owner_object_id = azuread_service_principal.test.object_id
}
`, r.template(data), data.RandomInteger)
}

func (r GroupOwnerResource) oneUser(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s
%[2]s

resource "azuread_group_owner" "testA" {
  group_object_id  = azuread_group.test.object_id
  owner_object_id = azuread_user.testA.object_id
}
`, r.template(data), r.templateThreeUsers(data))
}

func (r GroupOwnerResource) twoUsers(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s
%[2]s

resource "azuread_group_owner" "testA" {
  group_object_id  = azuread_group.test.object_id
  owner_object_id = azuread_user.testA.object_id
}

resource "azuread_group_owner" "testB" {
  group_object_id  = azuread_group.test.object_id
  owner_object_id = azuread_user.testB.object_id
}
`, r.template(data), r.templateThreeUsers(data))
}

func (r GroupOwnerResource) requiresImport(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_group_owner" "import" {
  group_object_id  = azuread_group_owner.test.group_object_id
  owner_object_id = azuread_group_owner.test.owner_object_id
}
`, r.group(data))
}
