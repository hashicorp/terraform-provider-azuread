// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identitygovernance_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
)

type AccessPackageResource struct{}

func TestAccAccessPackage_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_access_package", "test")
	r := AccessPackageResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("catalog_id"),
	})
}

func TestAccAccessPackage_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_access_package", "test")
	r := AccessPackageResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.complete(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep("catalog_id"),
	})
}

func TestAccAccessPackage_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_access_package", "test")
	r := AccessPackageResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.complete(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func (AccessPackageResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.IdentityGovernance.AccessPackageClient
	client.BaseClient.DisableRetries = true
	defer func() { client.BaseClient.DisableRetries = false }()

	_, status, err := client.Get(ctx, state.ID, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return utils.Bool(false), nil
		}
		return nil, fmt.Errorf("failed to retrieve access package with ID %q: %+v", state.ID, err)
	}
	return utils.Bool(true), nil
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
  hidden       = true
  catalog_id   = azuread_access_package_catalog.test_catalog.id
}
`, data.RandomInteger)
}
