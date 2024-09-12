// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identitygovernance_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementroleassignment"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
)

type AccessPackageCatalogRoleAssignmentResource struct{}

func TestAccAccessPackageCatalogRoleAssignmentResource_group(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_access_package_catalog_role_assignment", "test")
	r := AccessPackageCatalogRoleAssignmentResource{}

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: r.group(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("catalog_id").IsUuid(),
				check.That(data.ResourceName).Key("principal_object_id").IsUuid(),
				check.That(data.ResourceName).Key("role_id").Exists(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccAccessPackageCatalogRoleAssignmentResource_servicePrincipal(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_access_package_catalog_role_assignment", "test")
	r := AccessPackageCatalogRoleAssignmentResource{}

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: r.servicePrincipal(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("catalog_id").IsUuid(),
				check.That(data.ResourceName).Key("principal_object_id").IsUuid(),
				check.That(data.ResourceName).Key("role_id").Exists(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccAccessPackageCatalogRoleAssignmentResource_user(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_access_package_catalog_role_assignment", "test")
	r := AccessPackageCatalogRoleAssignmentResource{}

	data.DataSourceTest(t, []acceptance.TestStep{
		{
			Config: r.user(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("catalog_id").IsUuid(),
				check.That(data.ResourceName).Key("principal_object_id").IsUuid(),
				check.That(data.ResourceName).Key("role_id").Exists(),
			),
		},
		data.ImportStep(),
	})
}

func (r AccessPackageCatalogRoleAssignmentResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.IdentityGovernance.RoleAssignmentClient
	id := beta.NewRoleManagementEntitlementManagementRoleAssignmentID(state.ID)

	if resp, err := client.GetEntitlementManagementRoleAssignment(ctx, id, entitlementmanagementroleassignment.DefaultGetEntitlementManagementRoleAssignmentOperationOptions()); err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return pointer.To(false), nil
		}
		return nil, fmt.Errorf("failed to retrieve %s: %+v", id, err)
	}

	return pointer.To(true), nil
}

func (AccessPackageCatalogRoleAssignmentResource) group(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_access_package_catalog_role" "test" {
  display_name = "Catalog owner"
}

resource "azuread_group" "test" {
  display_name     = "acctestGroup-%[2]d"
  security_enabled = true
}

resource "azuread_access_package_catalog_role_assignment" "test" {
  role_id             = data.azuread_access_package_catalog_role.test.object_id
  catalog_id          = azuread_access_package_catalog.test.id
  principal_object_id = azuread_group.test.object_id
}
`, AccessPackageCatalogResource{}.basic(data), data.RandomInteger)
}

func (AccessPackageCatalogRoleAssignmentResource) servicePrincipal(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_access_package_catalog_role" "test" {
  display_name = "Catalog owner"
}

data "azuread_client_config" "test" {}

resource "azuread_access_package_catalog_role_assignment" "test" {
  role_id             = data.azuread_access_package_catalog_role.test.object_id
  catalog_id          = azuread_access_package_catalog.test.id
  principal_object_id = data.azuread_client_config.test.object_id
}
`, AccessPackageCatalogResource{}.basic(data), data.RandomInteger, data.RandomPassword)
}

func (AccessPackageCatalogRoleAssignmentResource) user(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_access_package_catalog_role" "test" {
  display_name = "Catalog owner"
}

data "azuread_domains" "test" {
  only_initial = true
}

resource "azuread_user" "test" {
  user_principal_name = "acctestUser'%[2]d@${data.azuread_domains.test.domains.0.domain_name}"
  display_name        = "acctestUser-%[2]d"
  password            = "%[3]s"
}

resource "azuread_access_package_catalog_role_assignment" "test" {
  role_id             = data.azuread_access_package_catalog_role.test.object_id
  catalog_id          = azuread_access_package_catalog.test.id
  principal_object_id = azuread_user.test.object_id
}
`, AccessPackageCatalogResource{}.basic(data), data.RandomInteger, data.RandomPassword)
}
