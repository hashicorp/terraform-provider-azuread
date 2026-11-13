// Copyright IBM Corp. 2019, 2026
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

// Several associations in one access package, so refreshes read the package from the cache. In step 1 the second
// association is created after the first has cached the package, so its Read fails if Create does not evict it.
// Step 2 deletes one association and creates another. Each apply starts with an empty cache.
func TestAccAccessPackageResourcePackageAssociation_multiple(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_access_package_resource_package_association", "test")
	r := AccessPackageResourcePackageAssociationResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.multiple(data, "second", 1),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That("azuread_access_package_resource_package_association.second").ExistsInAzure(r),
			),
		},
		{
			Config: r.multiple(data, "third", 2),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That("azuread_access_package_resource_package_association.third").ExistsInAzure(r),
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
  display_name     = "acctest-access-package-resource-catalog-association-%[1]d"
  security_enabled = true
}

resource "azuread_access_package_catalog" "test_catalog" {
  display_name = "acctest-catalog-%[1]d"
  description  = "Test catalog %[1]d"
}

resource "azuread_access_package_resource_catalog_association" "test" {
  catalog_id             = azuread_access_package_catalog.test_catalog.id
  resource_origin_id     = azuread_group.test_group.object_id
  resource_origin_system = "AadGroup"
}

resource "azuread_access_package" "test" {
  display_name = "acctest-package-%[1]d"
  description  = "Test Package %[1]d"
  catalog_id   = azuread_access_package_catalog.test_catalog.id
}

resource "azuread_access_package_resource_package_association" "test" {
  access_package_id               = azuread_access_package.test.id
  catalog_resource_association_id = azuread_access_package_resource_catalog_association.test.id
}
`, data.RandomInteger)
}

// multiple has the association "test" on group 0 and the association named by other on group index, in one access package.
// It uses count, not for_each: the test framework cannot read for_each keys from state. The association named by other
// depends on "test", so Terraform creates it after the Read of "test" has cached the access package.
func (AccessPackageResourcePackageAssociationResource) multiple(data acceptance.TestData, other string, index int) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_group" "test" {
  count            = 3
  display_name     = "acctest-access-package-resource-package-association-${count.index}-%[1]d"
  security_enabled = true
}

resource "azuread_access_package_catalog" "test" {
  display_name = "acctest-catalog-%[1]d"
  description  = "Test catalog %[1]d"
}

resource "azuread_access_package_resource_catalog_association" "test" {
  count                  = 3
  catalog_id             = azuread_access_package_catalog.test.id
  resource_origin_id     = azuread_group.test[count.index].object_id
  resource_origin_system = "AadGroup"
}

resource "azuread_access_package" "test" {
  display_name = "acctest-package-%[1]d"
  description  = "Test Package %[1]d"
  catalog_id   = azuread_access_package_catalog.test.id
}

resource "azuread_access_package_resource_package_association" "test" {
  access_package_id               = azuread_access_package.test.id
  catalog_resource_association_id = azuread_access_package_resource_catalog_association.test[0].id
}

resource "azuread_access_package_resource_package_association" "%[2]s" {
  access_package_id               = azuread_access_package.test.id
  catalog_resource_association_id = azuread_access_package_resource_catalog_association.test[%[3]d].id

  depends_on = [azuread_access_package_resource_package_association.test]
}
`, data.RandomInteger, other, index)
}
