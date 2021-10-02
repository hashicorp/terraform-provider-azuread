package identitygovernance_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/manicminer/hamilton/odata"

	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
)

type AccessPackageCatalogResource struct{}

func TestAccAccessPackageCatalog_simple(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_access_package_catalog", "test")
	r := AccessPackageCatalogResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.APC(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("id").Exists(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccAccessPackageCatalog_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_access_package_catalog", "test")
	r := AccessPackageCatalogResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.APC(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.APCUpdate(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.APC(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func (r AccessPackageCatalogResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	_, status, err := clients.IdentityGovernance.AccessPackageCatalogClient.Get(ctx, state.ID, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return nil, fmt.Errorf("Named Location with object ID %q does not exist", state.ID)
		}
		return nil, fmt.Errorf("failed to retrieve Named Location with object ID %q: %+v", state.ID, err)
	}

	return utils.Bool(true), nil
}

func (AccessPackageCatalogResource) APC(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_access_package_catalog" "test" {
  display_name = "acctestAPC-%[1]d"
  catalog_status = "Unpublished"
  description = "My test Catalog"
  is_externally_visible = true
}
`, data.RandomInteger)
}

func (AccessPackageCatalogResource) APCUpdate(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_access_package_catalog" "test" {
  display_name = "acctestAPC-%[1]d-updated"
  catalog_status = "Published"
  description = "My test Catalog Updated"
  is_externally_visible = false
}
`, data.RandomInteger)
}