package groups_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/terraform-providers/terraform-provider-azuread/internal/acceptance"
	"github.com/terraform-providers/terraform-provider-azuread/internal/acceptance/check"
)

type GroupDataSource struct{}

func TestAccGroupDataSource_byName(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_group", "test")

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: GroupDataSource{}.name(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("name").HasValue(fmt.Sprintf("acctestGroup-%d", data.RandomInteger)),
			),
		},
	})
}
func TestAccGroupDataSource_byNameWithSecurity(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_group", "test")

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: GroupDataSource{}.nameSecurity(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("name").HasValue(fmt.Sprintf("acctestGroup-%d", data.RandomInteger)),
			),
		},
	})
}

func TestAccGroupDataSource_byNameDeprecated(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_group", "test")

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: GroupDataSource{}.nameDeprecated(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("name").HasValue(fmt.Sprintf("acctestGroup-%d", data.RandomInteger)),
			),
		},
	})
}

func TestAccGroupDataSource_byCaseInsensitiveName(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_group", "test")

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: GroupDataSource{}.caseInsensitiveName(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("name").HasValue(fmt.Sprintf("acctestGroup-%d", data.RandomInteger)),
			),
		},
	})
}

func TestAccGroupDataSource_byObjectId(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_group", "test")

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: GroupDataSource{}.objectId(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("name").HasValue(fmt.Sprintf("acctestGroup-%d", data.RandomInteger)),
			),
		},
	})
}

func TestAccGroupDataSource_byObjectIdWithSecurity(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_group", "test")

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: GroupDataSource{}.objectIdSecurity(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("name").HasValue(fmt.Sprintf("acctestGroup-%d", data.RandomInteger)),
			),
		},
	})
}

func TestAccGroupDataSource_members(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_group", "test")

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: GroupDataSource{}.members(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("name").HasValue(fmt.Sprintf("acctestGroup-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("members.#").HasValue("3"),
			),
		},
	})
}

func TestAccGroupDataSource_owners(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_group", "test")

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: GroupDataSource{}.owners(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("name").HasValue(fmt.Sprintf("acctestGroup-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("owners.#").HasValue("3"),
			),
		},
	})
}

func (GroupDataSource) name(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_group" "test" {
  display_name = azuread_group.test.name
}
`, GroupResource{}.basic(data))
}

func (GroupDataSource) nameDeprecated(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_group" "test" {
  name = azuread_group.test.name
}
`, GroupResource{}.basic(data))
}

func (GroupDataSource) nameSecurity(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_group" "test" {
  display_name     = azuread_group.test.name
  mail_enabled     = false
  security_enabled = true
}
`, GroupResource{}.basic(data))
}

func (GroupDataSource) caseInsensitiveName(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s
data "azuread_group" "test" {
  display_name = upper(azuread_group.test.name)
}
`, GroupResource{}.basic(data))
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

func (GroupDataSource) owners(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_group" "test" {
  object_id = azuread_group.test.object_id
}
`, GroupResource{}.withThreeOwners(data))
}
