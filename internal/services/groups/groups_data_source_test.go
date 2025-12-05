// Copyright IBM Corp. 2019, 2025
// SPDX-License-Identifier: MPL-2.0

package groups_test

import (
	"context"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	groupBeta "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/group"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
)

type GroupsDataSource struct{}

func TestAccGroupsDataSource_byDisplayNames(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_groups", "test")
	r := GroupsDataSource{}

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: r.template(data),
		},
		{
			Config: r.byDisplayNames(data),
			Check: acceptance.ComposeAggregateTestCheckFunc(
				check.That(data.ResourceName).Key("display_names.#").HasValue("2"),
				check.That(data.ResourceName).Key("object_ids.#").HasValue("2"),
			),
		},
	})
}

func TestAccGroupsDataSource_byDisplayNamesIgnoreMissing(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_groups", "test")
	r := GroupsDataSource{}

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: r.template(data),
		},
		{
			Config: r.byDisplayNamesIgnoreMissing(data),
			Check: acceptance.ComposeAggregateTestCheckFunc(
				check.That(data.ResourceName).Key("display_names.#").HasValue("2"),
				check.That(data.ResourceName).Key("object_ids.#").HasValue("2"),
			),
		},
	})
}

func TestAccGroupsDataSource_byDisplayNamePrefix(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_groups", "test")
	r := GroupsDataSource{}
	moreThanZero := regexp.MustCompile("^[1-9][0-9]*$")

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: r.template(data),
		},
		{
			Config: r.byDisplayNamePrefix(data),
			Check: acceptance.ComposeAggregateTestCheckFunc(
				check.That(data.ResourceName).Key("display_names.#").MatchesRegex(moreThanZero),
				check.That(data.ResourceName).Key("object_ids.#").MatchesRegex(moreThanZero),
			),
		},
	})
}

func TestAccGroupsDataSource_byObjectIds(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_groups", "test")
	r := GroupsDataSource{}

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: r.template(data),
		},
		{
			Config: r.byObjectIds(data),
			Check: acceptance.ComposeAggregateTestCheckFunc(
				check.That(data.ResourceName).Key("display_names.#").HasValue("2"),
				check.That(data.ResourceName).Key("object_ids.#").HasValue("2"),
			),
		},
	})
}

func TestAccGroupsDataSource_noNames(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_groups", "test")

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: GroupsDataSource{}.noNames(),
			Check: acceptance.ComposeAggregateTestCheckFunc(
				check.That(data.ResourceName).Key("display_names.#").HasValue("0"),
				check.That(data.ResourceName).Key("object_ids.#").HasValue("0"),
			),
		},
	})
}

func TestAccGroupsDataSource_returnAll(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_groups", "test")
	r := GroupsDataSource{}

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: r.template(data),
		},
		{
			Config: r.returnAll(),
			Check: acceptance.ComposeAggregateTestCheckFunc(
				check.That(data.ResourceName).Key("display_names.#").Exists(),
				check.That(data.ResourceName).Key("object_ids.#").Exists(),
			),
		},
	})
}

func TestAccGroupsDataSource_returnAllMailEnabled(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_groups", "test")
	r := GroupsDataSource{}

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: r.template(data),
		},
		{
			Config: r.returnAllMailEnabled(data),
			Check: acceptance.ComposeAggregateTestCheckFunc(
				check.That(data.ResourceName).Key("display_names.#").Exists(),
				check.That(data.ResourceName).Key("object_ids.#").Exists(),
				check.That(data.ResourceName).Key("object_ids").ValidatesWith(testCheckHasOnlyMailEnabledGroups()),
			),
		},
	})
}

func TestAccGroupsDataSource_returnAllSecurityEnabled(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_groups", "test")
	r := GroupsDataSource{}

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: r.template(data),
		},
		{
			Config: r.returnAllSecurityEnabled(data),
			Check: acceptance.ComposeAggregateTestCheckFunc(
				check.That(data.ResourceName).Key("display_names.#").Exists(),
				check.That(data.ResourceName).Key("object_ids.#").Exists(),
				check.That(data.ResourceName).Key("object_ids").ValidatesWith(testCheckHasOnlySecurityEnabledGroups()),
			),
		},
	})
}

func TestAccGroupsDataSource_returnAllMailNotSecurityEnabled(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_groups", "test")
	r := GroupsDataSource{}

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: r.template(data),
		},
		{
			Config: r.returnAllMailNotSecurityEnabled(data),
			Check: acceptance.ComposeAggregateTestCheckFunc(
				check.That(data.ResourceName).Key("display_names.#").Exists(),
				check.That(data.ResourceName).Key("object_ids.#").Exists(),
				check.That(data.ResourceName).Key("object_ids").ValidatesWith(testCheckHasOnlyMailEnabledGroupsNotSecurityEnabledGroups()),
			),
		},
	})
}

func TestAccGroupsDataSource_returnAllSecurityNotMailEnabled(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_groups", "test")
	r := GroupsDataSource{}

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: r.template(data),
		},
		{
			Config: r.returnAllSecurityNotMailEnabled(data),
			Check: acceptance.ComposeAggregateTestCheckFunc(
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
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, 5*time.Minute)
		defer cancel()
		client := clients.Groups.GroupClientBeta

		for _, v := range values {
			id := beta.NewGroupID(v.(string))
			resp, err := client.GetGroup(ctx, id, groupBeta.DefaultGetGroupOperationOptions())
			if err != nil {
				return fmt.Errorf("retrieving %s: %v", id, err)
			}
			group := resp.Model
			if group == nil {
				return fmt.Errorf("retrieving %s: group was nil", id)
			}
			if group.Id == nil {
				return fmt.Errorf("retrieving %s: ID was nil", id)
			}
			if group.DisplayName == nil {
				return fmt.Errorf("retrieving %s: DisplayName was nil", id)
			}
			if hasMailGroupsOnly && !group.MailEnabled.GetOrZero() {
				return fmt.Errorf("expected only mail-enabled groups, encountered %s which is not mail-enabled", id)
			}
			if hasSecurityGroupsOnly && !group.SecurityEnabled.GetOrZero() {
				return fmt.Errorf("expected only security-enabled groups, encountered %s which is not security-enabled", id)
			}
			if hasNoMailGroups && group.MailEnabled.GetOrZero() {
				return fmt.Errorf("expected no mail-enabled groups, encountered %s which is mail-enabled", id)
			}
			if hasNoSecurityGroups && group.SecurityEnabled.GetOrZero() {
				return fmt.Errorf("expected no security-enabled groups, encountered %s which is security-enabled", id)
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

resource "azuread_group" "testC" {
  display_name     = "acctestGroupC-%[1]d"
  mail_enabled     = true
  mail_nickname    = "acctestGroupC%[1]d"
  types            = ["Unified"]
  security_enabled = true
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

func (r GroupsDataSource) byDisplayNamesIgnoreMissing(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_groups" "test" {
  ignore_missing = true

  display_names = [
    azuread_group.testA.display_name,
    "not-a-real-group-%[2]d",
    azuread_group.testB.display_name,
  ]
}
`, r.template(data), data.RandomInteger)
}

func (r GroupsDataSource) byDisplayNamePrefix(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_groups" "test" {
  display_name_prefix = "acctestGroup"
  depends_on          = [azuread_group.testA, azuread_group.testB]
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

func (r GroupsDataSource) returnAllMailEnabled(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_groups" "test" {
  mail_enabled = true
  return_all   = true
}
`, r.template(data))
}

func (r GroupsDataSource) returnAllSecurityEnabled(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_groups" "test" {
  return_all       = true
  security_enabled = true
}
`, r.template(data))
}

func (r GroupsDataSource) returnAllMailNotSecurityEnabled(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_groups" "test" {
  mail_enabled     = true
  return_all       = true
  security_enabled = false
}
`, r.template(data))
}

func (r GroupsDataSource) returnAllSecurityNotMailEnabled(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_groups" "test" {
  mail_enabled     = false
  return_all       = true
  security_enabled = true
}
`, r.template(data))
}
