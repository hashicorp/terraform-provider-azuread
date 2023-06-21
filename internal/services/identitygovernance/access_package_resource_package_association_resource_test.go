// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

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
	"github.com/hashicorp/terraform-provider-azuread/internal/services/identitygovernance/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
)

type AccessPackageResourcePackageAssociationResource struct{}

func TestAccAccessPackageResourcePackageAssociation_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_access_package_resource_package_association", "test")
	r := AccessPackageResourcePackageAssociationResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config:  r.complete(data),
			Destroy: false,
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func (AccessPackageResourcePackageAssociationResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.IdentityGovernance.AccessPackageResourceRoleScopeClient
	client.BaseClient.DisableRetries = true
	defer func() { client.BaseClient.DisableRetries = false }()

	id, err := parse.AccessPackageResourcePackageAssociationID(state.ID)
	if err != nil {
		return nil, err
	}

	_, status, err := client.Get(ctx, id.AccessPackageId, id.ResourcePackageAssociationId)
	if err != nil {
		if status == http.StatusNotFound {
			return utils.Bool(false), nil
		}

		return nil, fmt.Errorf("failed to retrieve access package resource association with ID %q: %+v", id.ID(), err)
	}

	return utils.Bool(true), nil
}

func (AccessPackageResourcePackageAssociationResource) complete(data acceptance.TestData) string {
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

resource "azuread_access_package" "test" {
  display_name = "test-package-%[1]d"
  description  = "Test Package %[1]d"
  catalog_id   = azuread_access_package_catalog.test_catalog.id
}

resource "azuread_access_package_resource_package_association" "test" {
  access_package_id               = azuread_access_package.test.id
  catalog_resource_association_id = azuread_access_package_resource_catalog_association.test.id
}
`, data.RandomInteger)
}
