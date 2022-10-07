package identitygovernance_test

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
)

type AccessPackageResourceCatalogAssociationResource struct{}

func TestAccAccessPackageResourceCatalogAssociation_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_access_package_resource_catalog_association", "test")
	r := AccessPackageResourceCatalogAssociationResource{}

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

func (AccessPackageResourceCatalogAssociationResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.IdentityGovernance.AccessPackageResourceClient
	client.BaseClient.DisableRetries = true

	var exists = false
	defer func() {
		if err := recover(); err != nil {
			log.Printf("[DEBUG] This needs to be fixed in the upstream libaray: %v", err)
			exists = false
		}
		exists = false
	}()

	ids := strings.Split(state.ID, idDelimitor)
	catalogId := ids[0]
	resourceOriginId := ids[1]
	catalogResource, status, err := client.Get(ctx, catalogId, resourceOriginId)
	if err != nil {
		if status == http.StatusNotFound {
			return nil, fmt.Errorf("Access package catalog association with object ID %q does not exist", state.ID)
		}
		return nil, fmt.Errorf("failed to retrieve access package catalog association with object ID %q: %+v", state.ID, err)
	}
	exists = catalogResource.ID != nil && *catalogResource.OriginId == state.Attributes["resource_origin_id"]

	return &exists, nil
}

func (AccessPackageResourceCatalogAssociationResource) complete(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_group" "test_group" {
	display_name     = "test-access-package-resource-catalog-association-%[1]d"
	security_enabled = true
}

resource "azuread_access_package_catalog" "test_catalog" {
	display_name = "test-catalog-%[1]d"	
  	description  = "Test catalog %[1]d"
}

resource "azuread_access_package_resource_catalog_association" "test" {
  catalog_id             = azuread_access_package_catalog.test_catalog.id
  resource_origin_id     = azuread_group.test_group.object_id
  resource_origin_system = "AadGroup"
}
`, data.RandomInteger)
}
