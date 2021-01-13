package aadgraph_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/terraform-providers/terraform-provider-azuread/internal/acceptance"
	"github.com/terraform-providers/terraform-provider-azuread/internal/acceptance/check"
)

type ApplicationDataSource struct{}

func TestAccApplicationDataSource_byObjectId(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_application", "test")
	r := ApplicationDataSource{}

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: r.objectId(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("application_id").IsUuid(),
				check.That(data.ResourceName).Key("object_id").IsUuid(),
				check.That(data.ResourceName).Key("name").HasValue(fmt.Sprintf("acctest-APP-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("homepage").HasValue(fmt.Sprintf("https://homepage-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("identifier_uris.#").HasValue("1"),
				check.That(data.ResourceName).Key("reply_urls.#").HasValue("1"),
				check.That(data.ResourceName).Key("oauth2_allow_implicit_flow").HasValue("true"),
				check.That(data.ResourceName).Key("optional_claims.#").HasValue("1"),
				check.That(data.ResourceName).Key("optional_claims.0.access_token.#").HasValue("2"),
				check.That(data.ResourceName).Key("optional_claims.0.id_token.#").HasValue("1"),
				check.That(data.ResourceName).Key("required_resource_access.#").HasValue("2"),
				check.That(data.ResourceName).Key("group_membership_claims").HasValue("All"),
			),
		},
	})
}

func TestAccApplicationDataSource_byApplicationId(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_application", "test")
	r := ApplicationDataSource{}

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: r.applicationId(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("application_id").IsUuid(),
				check.That(data.ResourceName).Key("object_id").IsUuid(),
				check.That(data.ResourceName).Key("name").HasValue(fmt.Sprintf("acctest-APP-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("homepage").HasValue(fmt.Sprintf("https://homepage-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("identifier_uris.#").HasValue("1"),
				check.That(data.ResourceName).Key("reply_urls.#").HasValue("1"),
				check.That(data.ResourceName).Key("oauth2_allow_implicit_flow").HasValue("true"),
				check.That(data.ResourceName).Key("optional_claims.#").HasValue("1"),
				check.That(data.ResourceName).Key("optional_claims.0.access_token.#").HasValue("2"),
				check.That(data.ResourceName).Key("optional_claims.0.id_token.#").HasValue("1"),
				check.That(data.ResourceName).Key("required_resource_access.#").HasValue("2"),
				check.That(data.ResourceName).Key("group_membership_claims").HasValue("All"),
			),
		},
	})
}

func TestAccApplicationDataSource_byName(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_application", "test")
	r := ApplicationDataSource{}

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: r.name(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("application_id").IsUuid(),
				check.That(data.ResourceName).Key("object_id").IsUuid(),
				check.That(data.ResourceName).Key("name").HasValue(fmt.Sprintf("acctest-APP-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("homepage").HasValue(fmt.Sprintf("https://homepage-%d", data.RandomInteger)),
				check.That(data.ResourceName).Key("identifier_uris.#").HasValue("1"),
				check.That(data.ResourceName).Key("reply_urls.#").HasValue("1"),
				check.That(data.ResourceName).Key("oauth2_allow_implicit_flow").HasValue("true"),
				check.That(data.ResourceName).Key("optional_claims.#").HasValue("1"),
				check.That(data.ResourceName).Key("optional_claims.0.access_token.#").HasValue("2"),
				check.That(data.ResourceName).Key("optional_claims.0.id_token.#").HasValue("1"),
				check.That(data.ResourceName).Key("required_resource_access.#").HasValue("2"),
				check.That(data.ResourceName).Key("group_membership_claims").HasValue("All"),
			),
		},
	})
}

func (ApplicationDataSource) objectId(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_application" "test" {
  object_id = azuread_application.test.object_id
}
`, ApplicationResource{}.complete(data))
}

func (ApplicationDataSource) applicationId(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_application" "test" {
  application_id = azuread_application.test.application_id
}
`, ApplicationResource{}.complete(data))
}

func (ApplicationDataSource) name(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_application" "test" {
  name = azuread_application.test.name
}
`, ApplicationResource{}.complete(data))
}
