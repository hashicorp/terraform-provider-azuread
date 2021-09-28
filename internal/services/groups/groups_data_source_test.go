package groups_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/manicminer/hamilton/odata"

	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
)

type GroupsDataSource struct{}

func TestAccGroupsDataSource_byDisplayNames(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_groups", "test")
	r := GroupsDataSource{}

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: r.byDisplayNames(data),
			Check: resource.ComposeAggregateTestCheckFunc(
				check.That(data.ResourceName).Key("display_names.#").HasValue("2"),
				check.That(data.ResourceName).Key("object_ids.#").HasValue("2"),
			),
		},
	})
}

func TestAccGroupsDataSource_byObjectIds(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_groups", "test")
	r := GroupsDataSource{}

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: r.byObjectIds(data),
			Check: resource.ComposeAggregateTestCheckFunc(
				check.That(data.ResourceName).Key("display_names.#").HasValue("2"),
				check.That(data.ResourceName).Key("object_ids.#").HasValue("2"),
			),
		},
	})
}

func TestAccGroupsDataSource_noNames(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_groups", "test")

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: GroupsDataSource{}.noNames(),
			Check: resource.ComposeAggregateTestCheckFunc(
				check.That(data.ResourceName).Key("display_names.#").HasValue("0"),
				check.That(data.ResourceName).Key("object_ids.#").HasValue("0"),
			),
		},
	})
}

func TestAccGroupsDataSource_returnAll(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_groups", "test")

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: GroupsDataSource{}.returnAll(),
			Check: resource.ComposeAggregateTestCheckFunc(
				check.That(data.ResourceName).Key("display_names.#").Exists(),
				check.That(data.ResourceName).Key("object_ids.#").Exists(),
			),
		},
	})
}

func TestAccGroupsDataSource_returnAllMailEnabled(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_groups", "test")

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: GroupsDataSource{}.returnAllMailEnabled(data),
			Check: resource.ComposeAggregateTestCheckFunc(
				check.That(data.ResourceName).Key("display_names.#").Exists(),
				check.That(data.ResourceName).Key("object_ids.#").Exists(),
				check.That(data.ResourceName).Key("object_ids").ValidatesWith(testCheckHasOnlyMailEnabledGroups()),
			),
		},
	})
}

func TestAccGroupsDataSource_returnAllSecurityEnabled(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_groups", "test")

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: GroupsDataSource{}.returnAllSecurityEnabled(data),
			Check: resource.ComposeAggregateTestCheckFunc(
				check.That(data.ResourceName).Key("display_names.#").Exists(),
				check.That(data.ResourceName).Key("object_ids.#").Exists(),
				check.That(data.ResourceName).Key("object_ids").ValidatesWith(testCheckHasOnlySecurityEnabledGroups()),
			),
		},
	})
}

func TestAccGroupsDataSource_returnAllMailNotSecurityEnabled(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_groups", "test")

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: GroupsDataSource{}.returnAllMailNotSecurityEnabled(data),
			Check: resource.ComposeAggregateTestCheckFunc(
				check.That(data.ResourceName).Key("display_names.#").Exists(),
				check.That(data.ResourceName).Key("object_ids.#").Exists(),
				check.That(data.ResourceName).Key("object_ids").ValidatesWith(testCheckHasOnlyMailEnabledGroupsNotSecurityEnabledGroups()),
			),
		},
	})
}

func TestAccGroupsDataSource_returnAllSecurityNotMailEnabled(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_groups", "test")

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: GroupsDataSource{}.returnAllSecurityNotMailEnabled(data),
			Check: resource.ComposeAggregateTestCheckFunc(
				check.That(data.ResourceName).Key("display_names.#").Exists(),
				check.That(data.ResourceName).Key("object_ids.#").Exists(),
				check.That(data.ResourceName).Key("object_ids").ValidatesWith(testCheckHasOnlySecurityEnabledGroupsNotMailEnabledGroups()),
			),
		},
	})
}

func testCheckHasOnlyMailEnabledGroups() check.KeyValidationFunc {
	return testCheckGroupsDataSource(true, false, false, false)
}

func testCheckHasOnlySecurityEnabledGroups() check.KeyValidationFunc {
	return testCheckGroupsDataSource(false, true, false, false)
}

func testCheckHasOnlyMailEnabledGroupsNotSecurityEnabledGroups() check.KeyValidationFunc {
	return testCheckGroupsDataSource(true, false, false, true)
}

func testCheckHasOnlySecurityEnabledGroupsNotMailEnabledGroups() check.KeyValidationFunc {
	return testCheckGroupsDataSource(false, true, true, false)
}

func testCheckGroupsDataSource(hasMailGroupsOnly, hasSecurityGroupsOnly, hasNoMailGroups, hasNoSecurityGroups bool) check.KeyValidationFunc {
	return func(ctx context.Context, clients *clients.Client, values []interface{}) error {
		client := clients.Groups.GroupsClient

		for _, v := range values {
			oid := v.(string)
			group, _, err := client.Get(ctx, oid, odata.Query{})
			if err != nil {
				return fmt.Errorf("retrieving group with object ID %q: %+oid", oid, err)
			}
			if group == nil {
				return fmt.Errorf("retrieving group with object ID %q: group was nil", oid)
			}
			if group.ID == nil {
				return fmt.Errorf("retrieving group with object ID %q: ID was nil", oid)
			}
			if group.DisplayName == nil {
				return fmt.Errorf("retrieving group with object ID %q: DisplayName was nil", oid)
			}
			if hasMailGroupsOnly && group.MailEnabled != nil && !*group.MailEnabled {
				return fmt.Errorf("expected only mail-enabled groups, encountered group %q (object ID: %q) which is not mail-enabled", *group.DisplayName, *group.ID)
			}
			if hasSecurityGroupsOnly && group.SecurityEnabled != nil && !*group.SecurityEnabled {
				return fmt.Errorf("expected only security-enabled groups, encountered group %q (object ID: %q) which is not security-enabled", *group.DisplayName, *group.ID)
			}
			if hasNoMailGroups && group.MailEnabled != nil && *group.MailEnabled {
				return fmt.Errorf("expected no mail-enabled groups, encountered group %q (object ID: %q) which is mail-enabled", *group.DisplayName, *group.ID)
			}
			if hasNoSecurityGroups && group.SecurityEnabled != nil && *group.SecurityEnabled {
				return fmt.Errorf("expected no security-enabled groups, encountered group %q (object ID: %q) which is security-enabled", *group.DisplayName, *group.ID)
			}
		}

		return nil
	}
}

func (GroupsDataSource) template(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_group" "testA" {
  display_name     = "acctestGroupA-%[1]d"
  security_enabled = true
}

resource "azuread_group" "testB" {
  display_name  = "acctestGroupB-%[1]d"
  mail_enabled  = true
  mail_nickname = "acctestGroupB-%[1]d"
  types         = ["Unified"]
}
`, data.RandomInteger)
}

func (r GroupsDataSource) byDisplayNames(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_groups" "test" {
  display_names = [azuread_group.testA.display_name, azuread_group.testB.display_name]
}
`, r.template(data))
}

func (r GroupsDataSource) byObjectIds(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_groups" "test" {
  object_ids = [azuread_group.testA.object_id, azuread_group.testB.object_id]
}
`, r.template(data))
}

func (GroupsDataSource) noNames() string {
	return `
data "azuread_groups" "test" {
  display_names = []
}
`
}

func (GroupsDataSource) returnAll() string {
	return `
data "azuread_groups" "test" {
  return_all = true
}
`
}

func (GroupsDataSource) returnAllMailEnabled(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_groups" "test" {
  mail_enabled = true
  return_all   = true
}
`, GroupResource{}.multipleGroupsTemplate(data))
}

func (GroupsDataSource) returnAllSecurityEnabled(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_groups" "test" {
  return_all       = true
  security_enabled = true
}
`, GroupResource{}.multipleGroupsTemplate(data))
}

func (GroupsDataSource) returnAllMailNotSecurityEnabled(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_groups" "test" {
  mail_enabled     = true
  return_all       = true
  security_enabled = false
}
`, GroupResource{}.multipleGroupsTemplate(data))
}

func (GroupsDataSource) returnAllSecurityNotMailEnabled(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_groups" "test" {
  mail_enabled     = false
  return_all       = true
  security_enabled = true
}
`, GroupResource{}.multipleGroupsTemplate(data))
}
