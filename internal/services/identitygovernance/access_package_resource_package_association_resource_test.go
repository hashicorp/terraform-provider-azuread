// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package identitygovernance_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/identitygovernance"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/identitygovernance/parse"
)

type AccessPackageResourcePackageAssociationResource struct{}

func TestAccAccessPackageResourcePackageAssociation_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_access_package_resource_package_association", "test")
	r := AccessPackageResourcePackageAssociationResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config:  r.complete(data),
			Destroy: false,
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccAccessPackageResourcePackageAssociation_eligibleMember(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_access_package_resource_package_association", "test")
	r := AccessPackageResourcePackageAssociationResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config:  r.eligibleMember(data),
			Destroy: false,
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("access_type").HasValue("Eligible Member"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccAccessPackageResourcePackageAssociation_eligibleOwner(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_access_package_resource_package_association", "test")
	r := AccessPackageResourcePackageAssociationResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config:  r.eligibleOwner(data),
			Destroy: false,
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("access_type").HasValue("Eligible Owner"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccAccessPackageResourcePackageAssociation_applicationAppRole(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_access_package_resource_package_association", "test")
	r := AccessPackageResourcePackageAssociationResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config:  r.applicationAppRole(data),
			Destroy: false,
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("resource_role_origin_id").Exists(),
				check.That(data.ResourceName).Key("resource_role_display_name").Exists(),
			),
		},
		data.ImportStep(),
	})
}

func (AccessPackageResourcePackageAssociationResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.IdentityGovernance.AccessPackageClient

	resourceId, err := parse.AccessPackageResourcePackageAssociationID(state.ID)
	if err != nil {
		return nil, err
	}

	id := beta.NewIdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageResourceRoleScopeID(resourceId.AccessPackageId, resourceId.ResourceRoleScopeId)

	roleScope, err := identitygovernance.GetAccessPackageResourcesRoleScope(ctx, client, id)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve %s: %v", id, err)
	}

	return pointer.To(roleScope != nil), nil
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

func (AccessPackageResourcePackageAssociationResource) eligibleMember(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_group" "test_group" {
  display_name     = "test-access-package-eligible-member-%[1]d"
  security_enabled = true
}

resource "azuread_group_role_management_policy" "test_member" {
  group_id = azuread_group.test_group.object_id
  role_id  = "member"

  eligible_assignment_rules {
    expiration_required = false
  }
}

resource "azuread_access_package_catalog" "test_catalog" {
  display_name = "test-catalog-eligible-member-%[1]d"
  description  = "Test catalog eligible member %[1]d"
}

resource "azuread_access_package_resource_catalog_association" "test" {
  catalog_id             = azuread_access_package_catalog.test_catalog.id
  resource_origin_id     = azuread_group.test_group.object_id
  resource_origin_system = "AadGroup"
}

resource "azuread_access_package" "test" {
  display_name = "test-package-eligible-member-%[1]d"
  description  = "Test Package Eligible Member %[1]d"
  catalog_id   = azuread_access_package_catalog.test_catalog.id
}

resource "azuread_access_package_resource_package_association" "test" {
  access_package_id               = azuread_access_package.test.id
  catalog_resource_association_id = azuread_access_package_resource_catalog_association.test.id
  access_type                     = "Eligible Member"

  depends_on = [azuread_group_role_management_policy.test_member]
}
`, data.RandomInteger)
}

func (AccessPackageResourcePackageAssociationResource) eligibleOwner(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_group" "test_group" {
  display_name     = "test-access-package-eligible-owner-%[1]d"
  security_enabled = true
}

resource "azuread_group_role_management_policy" "test_owner" {
  group_id = azuread_group.test_group.object_id
  role_id  = "owner"

  eligible_assignment_rules {
    expiration_required = false
  }
}

resource "azuread_access_package_catalog" "test_catalog" {
  display_name = "test-catalog-eligible-owner-%[1]d"
  description  = "Test catalog eligible owner %[1]d"
}

resource "azuread_access_package_resource_catalog_association" "test" {
  catalog_id             = azuread_access_package_catalog.test_catalog.id
  resource_origin_id     = azuread_group.test_group.object_id
  resource_origin_system = "AadGroup"
}

resource "azuread_access_package" "test" {
  display_name = "test-package-eligible-owner-%[1]d"
  description  = "Test Package Eligible Owner %[1]d"
  catalog_id   = azuread_access_package_catalog.test_catalog.id
}

resource "azuread_access_package_resource_package_association" "test" {
  access_package_id               = azuread_access_package.test.id
  catalog_resource_association_id = azuread_access_package_resource_catalog_association.test.id
  access_type                     = "Eligible Owner"

  depends_on = [azuread_group_role_management_policy.test_owner]
}
`, data.RandomInteger)
}

func (AccessPackageResourcePackageAssociationResource) applicationAppRole(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application" "test" {
  display_name = "test-access-package-app-%[1]d"

  app_role {
    allowed_member_types = ["User"]
    description          = "Test AppRole"
    display_name         = "TestRole"
    id                   = "c3e219dc-0a44-4341-8660-868d389c9558"
    value                = "TestRole"
  }
}

resource "azuread_service_principal" "test" {
  client_id = azuread_application.test.client_id
}

resource "azuread_access_package_catalog" "test_catalog" {
  display_name = "test-catalog-app-role-%[1]d"
  description  = "Test catalog app role %[1]d"
}

resource "azuread_access_package_resource_catalog_association" "test" {
  catalog_id             = azuread_access_package_catalog.test_catalog.id
  resource_origin_id     = azuread_service_principal.test.object_id
  resource_origin_system = "AadApplication"
}

resource "azuread_access_package" "test" {
  display_name = "test-package-app-role-%[1]d"
  description  = "Test Package App Role %[1]d"
  catalog_id   = azuread_access_package_catalog.test_catalog.id
}

resource "azuread_access_package_resource_package_association" "test" {
  access_package_id               = azuread_access_package.test.id
  catalog_resource_association_id = azuread_access_package_resource_catalog_association.test.id
  resource_role_origin_id         = tolist(azuread_application.test.app_role)[0].id
}
`, data.RandomInteger)
}
