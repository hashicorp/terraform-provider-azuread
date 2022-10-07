package identitygovernance_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/manicminer/hamilton/msgraph"
)

type AccessPackageCatalogDataSource struct{}

func TestAccAccessPackageCatalogDataSource_byId(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_access_package_catalog", "test")
	r := AccessPackageCatalogDataSource{}

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: r.byId(data),
			Check:  r.testCheckFunc(data),
		},
	})
}

func TestAccAccessPackageCatalogDataSource_byDisplayName(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_access_package_catalog", "test")
	r := AccessPackageCatalogDataSource{}

	data.DataSourceTest(t, []resource.TestStep{
		{
			Config: r.byDisplayName(data),
			Check:  r.testCheckFunc(data),
		},
	})
}

func (AccessPackageCatalogDataSource) testCheckFunc(data acceptance.TestData) resource.TestCheckFunc {
	return resource.ComposeTestCheckFunc(
		check.That(data.ResourceName).Key("description").HasValue(fmt.Sprintf("Test access package catalog %[1]d", data.RandomInteger)),
		check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("test-access-package-catalog-%[1]d", data.RandomInteger)),
		check.That(data.ResourceName).Key("state").HasValue(msgraph.AccessPackageCatalogStateUnpublished),
		check.That(data.ResourceName).Key("is_externally_visible").HasValue("false"),
	)
}

func (AccessPackageCatalogDataSource) byId(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_access_package_catalog" "test" {
  object_id = azuread_access_package_catalog.test.id
}
`, AccessPackageCatalogResource{}.complete(data))
}

func (AccessPackageCatalogDataSource) byDisplayName(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_access_package_catalog" "test" {
  display_name = azuread_access_package_catalog.test.display_name
}
`, AccessPackageCatalogResource{}.complete(data))
}
