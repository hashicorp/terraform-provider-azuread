// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package identitygovernance_test

import (
	"context"
	"fmt"
	"regexp"
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

func TestAccAccessPackageResourcePackageAssociation_completeWithGroup(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_access_package_resource_package_association", "test")
	r := AccessPackageResourcePackageAssociationResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config:  r.completeWithGroup(data),
			Destroy: false,
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("access_type").HasValue("Member"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccAccessPackageResourcePackageAssociation_completeWithGroupOwner(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_access_package_resource_package_association", "test")
	r := AccessPackageResourcePackageAssociationResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config:  r.completeWithGroupOwner(data),
			Destroy: false,
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("access_type").HasValue("Owner"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccAccessPackageResourcePackageAssociation_completeWithApplication(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_access_package_resource_package_association", "test")
	r := AccessPackageResourcePackageAssociationResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config:  r.completeWithApplication(data),
			Destroy: false,
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("access_type").IsUuid(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccAccessPackageResourcePackageAssociation_invalidAccessType(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_access_package_resource_package_association", "test")
	r := AccessPackageResourcePackageAssociationResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config:      r.invalidAccessType(data),
			ExpectError: regexp.MustCompile(`expected access_type to be one of \[Member Owner\]`),
		},
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

func (AccessPackageResourcePackageAssociationResource) completeWithGroup(data acceptance.TestData) string {
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
  access_type                     = "Member"
}
`, data.RandomInteger)
}

func (AccessPackageResourcePackageAssociationResource) completeWithGroupOwner(data acceptance.TestData) string {
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
  access_type                     = "Owner"
}
`, data.RandomInteger)
}

func (AccessPackageResourcePackageAssociationResource) completeWithApplication(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

data "azuread_application_published_app_ids" "well_known" {}

resource "azuread_service_principal" "msgraph" {
  client_id    = data.azuread_application_published_app_ids.well_known.result.MicrosoftGraph
  use_existing = true
}

resource "azuread_application" "test" {
  display_name = "acctest-packageAssociationResource-%[1]d"

  required_resource_access {
    resource_app_id = data.azuread_application_published_app_ids.well_known.result.MicrosoftGraph

    resource_access {
      id   = azuread_service_principal.msgraph.app_role_ids["User.Read.All"]
      type = "Role"
    }
  }
}

resource "azuread_service_principal" "test" {
  client_id = azuread_application.test.client_id
}

resource "azuread_app_role_assignment" "test" {
  app_role_id         = azuread_service_principal.msgraph.app_role_ids["User.Read.All"]
  principal_object_id = azuread_service_principal.test.object_id
  resource_object_id  = azuread_service_principal.msgraph.object_id
}

resource "azuread_access_package_catalog" "test_catalog" {
  display_name = "test-catalog-%[1]d"
  description  = "Test catalog %[1]d"
}

resource "azuread_access_package_resource_catalog_association" "test" {
  catalog_id             = azuread_access_package_catalog.test_catalog.id
  resource_origin_id     = azuread_service_principal.test.object_id
  resource_origin_system = "AadApplication"
}

resource "azuread_access_package" "test" {
  display_name = "test-package-%[1]d"
  description  = "Test Package %[1]d"
  catalog_id   = azuread_access_package_catalog.test_catalog.id
}

resource "azuread_access_package_resource_package_association" "test" {
  access_package_id               = azuread_access_package.test.id
  catalog_resource_association_id = azuread_access_package_resource_catalog_association.test.id
  access_type                     = azuread_service_principal.msgraph.app_role_ids["User.Read.All"]
}
`, data.RandomInteger)
}

func (AccessPackageResourcePackageAssociationResource) invalidAccessType(data acceptance.TestData) string {
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
  access_type                     = "InvalidValue"
}
`, data.RandomInteger)
}
