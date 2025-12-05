// Copyright IBM Corp. 2019, 2025
// SPDX-License-Identifier: MPL-2.0

package identitygovernance_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/entitlementmanagementaccesspackagecatalog"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
)

type AccessPackageCatalogResource struct{}

func TestAccAccessPackageCatalog_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_access_package_catalog", "test")
	r := AccessPackageCatalogResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccAccessPackageCatalog_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_access_package_catalog", "test")
	r := AccessPackageCatalogResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.complete(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccAccessPackageCatalog_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_access_package_catalog", "test")
	r := AccessPackageCatalogResource{}

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

func (AccessPackageCatalogResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.IdentityGovernance.AccessPackageCatalogClient
	id := beta.NewIdentityGovernanceEntitlementManagementAccessPackageCatalogID(state.ID)

	resp, err := client.GetEntitlementManagementAccessPackageCatalog(ctx, id, entitlementmanagementaccesspackagecatalog.DefaultGetEntitlementManagementAccessPackageCatalogOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return pointer.To(false), nil
		}
		return nil, fmt.Errorf("failed to retrieve %s: %v", id, err)
	}

	return pointer.To(true), nil
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
