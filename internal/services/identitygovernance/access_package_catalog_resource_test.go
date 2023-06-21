// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identitygovernance_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
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
	defer func() { client.BaseClient.DisableRetries = false }()

	_, status, err := client.Get(ctx, state.ID, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return utils.Bool(false), nil
		}

		return nil, fmt.Errorf("failed to retrieve access package catalog with ID %q: %+v", state.ID, err)
	}
	return utils.Bool(true), nil
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
  display_name       = "test-access-package-catalog-%[1]d"
  description        = "Test access package catalog %[1]d"
  externally_visible = false
  published          = false
}
`, data.RandomInteger)
}
