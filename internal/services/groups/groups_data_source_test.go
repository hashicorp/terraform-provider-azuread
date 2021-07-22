package groups_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
)

type GroupsDataSource struct{}

func TestAccGroupsDataSource_byDisplayNames(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_groups", "test")
	r := GroupsDataSource{}

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: r.byDisplayNames(data),
			Check: resource.ComposeTestCheckFunc(
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
			Check: resource.ComposeTestCheckFunc(
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
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("display_names.#").HasValue("0"),
				check.That(data.ResourceName).Key("object_ids.#").HasValue("0"),
			),
		},
	})
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
