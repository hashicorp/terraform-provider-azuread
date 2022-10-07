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

type AccessPackageResource struct{}

func TestAccAccessPackage_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_access_package", "test")
	r := AccessPackageResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("catalog_id"),
	})
}

func TestAccAccessPackage_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_access_package", "test")
	r := AccessPackageResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.complete(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("catalog_id"),
	})
}

func TestAccAccessPackage_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_access_package", "test")
	r := AccessPackageResource{}

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

func (AccessPackageResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.IdentityGovernance.AccessPackageClient
	client.BaseClient.DisableRetries = true

	accessPackage, status, err := client.Get(ctx, state.ID, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return nil, fmt.Errorf("Access package with object ID %q does not exist", state.ID)
		}
		return nil, fmt.Errorf("failed to retrieve access package with object ID %q: %+v", state.ID, err)
	}
	return utils.Bool(accessPackage.ID != nil && *accessPackage.ID == state.ID), nil
}

func (AccessPackageResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_access_package_catalog" "test_catalog" {
	display_name = "test-catalog-%[1]d"	
  	description  = "Test catalog %[1]d"
}

resource "azuread_access_package" "test" {
  display_name = "access-package-%[1]d"
  description  = "Access Package %[1]d"
  catalog_id   = azuread_access_package_catalog.test_catalog.id
}
`, data.RandomInteger)
}

func (AccessPackageResource) complete(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_access_package_catalog" "test_catalog" {
	display_name = "test-catalog-%[1]d"	
  	description  = "Test catalog %[1]d"
}

resource "azuread_access_package" "test" {
  display_name = "access-package-%[1]d"
  description  = "Access Package %[1]d"
  is_hidden    = true
  catalog_id   = azuread_access_package_catalog.test_catalog.id
}
`, data.RandomInteger)
}
