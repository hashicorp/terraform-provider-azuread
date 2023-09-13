// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identitygovernance_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/identitygovernance/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
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

func TestAccAccessPackageResourceCatalogAssociation_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_access_package_resource_catalog_association", "test")
	r := AccessPackageResourceCatalogAssociationResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.complete(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.RequiresImportErrorStep(r.requiresImport(data)),
	})
}

func (r AccessPackageResourceCatalogAssociationResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.IdentityGovernance.AccessPackageResourceClient
	client.BaseClient.DisableRetries = true
	defer func() { client.BaseClient.DisableRetries = false }()

	id, err := parse.AccessPackageResourceCatalogAssociationID(state.ID)
	if err != nil {
		return nil, err
	}

	_, status, err := client.Get(ctx, id.CatalogId, id.OriginId)
	if err != nil {
		if status == http.StatusNotFound {
			return utils.Bool(false), nil
		}

		return nil, fmt.Errorf("failed to retrieve access package catalog association with ID %q: %+v", id.ID(), err)
	}

	return utils.Bool(true), nil
}

func (r AccessPackageResourceCatalogAssociationResource) complete(data acceptance.TestData) string {
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

func (r AccessPackageResourceCatalogAssociationResource) requiresImport(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_access_package_resource_catalog_association" "import" {
  catalog_id             = azuread_access_package_resource_catalog_association.test.catalog_id
  resource_origin_id     = azuread_access_package_resource_catalog_association.test.resource_origin_id
  resource_origin_system = azuread_access_package_resource_catalog_association.test.resource_origin_system
}
`, r.complete(data), data.RandomInteger)
}
