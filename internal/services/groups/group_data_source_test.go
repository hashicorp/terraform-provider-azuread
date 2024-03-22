// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package groups_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
)

type GroupDataSource struct{}

func TestAccGroupDataSource_byDisplayName(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_group", "test")

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: GroupDataSource{}.displayName(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctestGroup-%d", data.RandomInteger)),
			),
		},
	})
}
func TestAccGroupDataSource_byDisplayNameWithSecurity(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_group", "test")

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: GroupDataSource{}.displayNameSecurity(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctestGroup-%d", data.RandomInteger)),
			),
		},
	})
}

func TestAccGroupDataSource_byDisplayNameWithSecurityNotMail(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_group", "test")

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: GroupDataSource{}.displayNameSecurityNotMail(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctestGroup-%d", data.RandomInteger)),
			),
		},
	})
}

func TestAccGroupDataSource_byCaseInsensitiveDisplayName(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_group", "test")

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: GroupDataSource{}.caseInsensitiveDisplayName(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctestGroup-%d", data.RandomInteger)),
			),
		},
	})
}

func TestAccGroupDataSource_byMailNickname(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_group", "test")

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: GroupDataSource{}.mailNickname(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctestGroup-%d", data.RandomInteger)),
			),
		},
	})
}

func TestAccGroupDataSource_byMailNicknameWithSecurity(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_group", "test")

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: GroupDataSource{}.mailNicknameSecurity(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctestGroup-%d", data.RandomInteger)),
			),
		},
	})
}

func TestAccGroupDataSource_byMailNicknameWithSecurityNotMail(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_group", "test")

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: GroupDataSource{}.mailNicknameSecurityNotMail(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctestGroup-%d", data.RandomInteger)),
			),
		},
	})
}

func TestAccGroupDataSource_byObjectId(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_group", "test")

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: GroupDataSource{}.objectId(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctestGroup-%d", data.RandomInteger)),
			),
		},
	})
}

func TestAccGroupDataSource_byObjectIdWithSecurity(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_group", "test")

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: GroupDataSource{}.objectIdSecurity(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctestGroup-%d", data.RandomInteger)),
			),
		},
	})
}

func TestAccGroupDataSource_dynamicMembership(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_group", "test")

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: GroupDataSource{}.dynamicMembership(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctestGroup-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("dynamic_membership.#").HasValue("1"),
				check.That(data.ResourceName).Key("dynamic_membership.0.enabled").HasValue("true"),
				check.That(data.ResourceName).Key("dynamic_membership.0.rule").HasValue("user.department -eq \"Sales\""),
			),
		},
	})
}

func TestAccGroupDataSource_members(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_group", "test")

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: GroupDataSource{}.members(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctestGroup-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("members.#").HasValue("3"),
			),
		},
	})
}

func TestAccGroupDataSource_transitiveMembers(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_group", "test")

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: GroupDataSource{}.transitiveMembers(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctestGroup-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("members.#").HasValue("4"),
			),
		},
	})
}

func TestAccGroupDataSource_owners(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_group", "test")

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: GroupDataSource{}.owners(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctestGroup-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("owners.#").HasValue("3"),
			),
		},
	})
}

func TestAccGroupDataSource_unifiedExtraSettings(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_group", "test")

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: GroupDataSource{}.unifiedWithExtraSettings(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctestGroup-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("auto_subscribe_new_members").HasValue("true"),
				check.That(data.ResourceName).Key("external_senders_allowed").HasValue("true"),
				check.That(data.ResourceName).Key("hide_from_address_lists").HasValue("true"),
				check.That(data.ResourceName).Key("hide_from_outlook_clients").HasValue("true"),
			),
		},
	})
}

func TestAccGroupDataSource_writeback(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_group", "test")

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: GroupDataSource{}.writeback(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctestGroup-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("writeback_enabled").HasValue("true"),
				check.That(data.ResourceName).Key("onpremises_group_type").HasValue("UniversalSecurityGroup"),
			),
		},
	})
}

func (GroupDataSource) displayName(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_group" "test" {
  display_name = azuread_group.test.display_name
}
`, GroupResource{}.basic(data))
}

func (GroupDataSource) displayNameSecurity(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_group" "test" {
  display_name     = azuread_group.test.display_name
  security_enabled = true
}
`, GroupResource{}.basic(data))
}

func (GroupDataSource) displayNameSecurityNotMail(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s
%[2]s

data "azuread_group" "test" {
  display_name     = azuread_group.test.display_name
  mail_enabled     = false
  security_enabled = true
}
`, GroupResource{}.basic(data), GroupResource{}.basicUnified(data))
}

func (GroupDataSource) caseInsensitiveDisplayName(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s
data "azuread_group" "test" {
  display_name = upper(azuread_group.test.display_name)
}
`, GroupResource{}.basic(data))
}

func (GroupDataSource) mailNickname(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_group" "test" {
  mail_nickname = azuread_group.test.mail_nickname
}
`, GroupResource{}.basic(data))
}

func (GroupDataSource) mailNicknameSecurity(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_group" "test" {
  mail_nickname    = azuread_group.test.mail_nickname
  security_enabled = true
}
`, GroupResource{}.basic(data))
}

func (GroupDataSource) mailNicknameSecurityNotMail(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s
%[2]s

data "azuread_group" "test" {
  mail_nickname    = azuread_group.test.mail_nickname
  mail_enabled     = false
  security_enabled = true
}
`, GroupResource{}.basic(data), GroupResource{}.basicUnified(data))
}

func (GroupDataSource) objectId(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_group" "test" {
  object_id = azuread_group.test.object_id
}
`, GroupResource{}.basic(data))
}

func (GroupDataSource) objectIdSecurity(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_group" "test" {
  object_id        = azuread_group.test.object_id
  mail_enabled     = false
  security_enabled = true
}
`, GroupResource{}.basic(data))
}

func (GroupDataSource) members(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_group" "test" {
  object_id = azuread_group.test.object_id
}
`, GroupResource{}.withThreeMembers(data))
}

func (GroupDataSource) transitiveMembers(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_group" "test" {
  object_id                  = azuread_group.test.object_id
  include_transitive_members = true
}
`, GroupResource{}.withTransitiveMembers(data))
}

func (GroupDataSource) dynamicMembership(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_group" "test" {
  object_id = azuread_group.test.object_id
}
`, GroupResource{}.dynamicMembership(data))
}

func (GroupDataSource) owners(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_group" "test" {
  object_id = azuread_group.test.object_id
}
`, GroupResource{}.withThreeOwners(data))
}

func (GroupDataSource) unifiedWithExtraSettings(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_group" "test" {
  object_id = azuread_group.test.object_id
}
`, GroupResource{}.unifiedWithExtraSettings(data))
}

func (GroupDataSource) writeback(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_group" "test" {
  display_name = azuread_group.test.display_name
}
`, GroupResource{}.withWriteback(data))
}
