package identitygovernance_test

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
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/manicminer/hamilton/odata"
)

type AccessPackageCatalogResource struct{}

func TestAccAccessPackageCatalog_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_access_package_catalog", "test")
	r := AccessPackageCatalogResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccAccessPackageCatalog_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_access_package_catalog", "test")
	r := AccessPackageCatalogResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.complete(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
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
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.complete(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func (AccessPackageCatalogResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.IdentityGovernance.AccessPackageCatalogClient
	client.BaseClient.DisableRetries = true

	accessPackageCatalog, status, err := client.Get(ctx, state.ID, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return nil, fmt.Errorf("Access package catalog with object ID %q does not exist", state.ID)
		}
		return nil, fmt.Errorf("failed to retrieve access package catalog with object ID %q: %+v", state.ID, err)
	}
	return utils.Bool(accessPackageCatalog.ID != nil && *accessPackageCatalog.ID == state.ID), nil
}

func (AccessPackageCatalogResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_access_package_catalog" "test" {
  display_name = "test-access-package-catalog-%[1]d"
  description  = "Test access package catalog %[1]d"
}
`, data.RandomInteger)
}

func (AccessPackageCatalogResource) complete(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_access_package_catalog" "test" {
  display_name 			= "test-access-package-catalog-%[1]d"
  description  			= "Test access package catalog %[1]d"
  state		   			= "unpublished"
  is_externally_visible = false
}
`, data.RandomInteger)
}
