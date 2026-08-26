// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package identitygovernance_test

import (
	"context"
	"fmt"
	"os"
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
				check.That(data.ResourceName).Key("access_type").HasValue("Member"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccAccessPackageResourcePackageAssociation_completeWithAccessTypeOwner(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_access_package_resource_package_association", "test")
	r := AccessPackageResourcePackageAssociationResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config:  r.completeWithAccessTypeOwner(data),
			Destroy: false,
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("access_type").HasValue("Owner"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccAccessPackageResourcePackageAssociation_completeWithAccessTypeApplication(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_access_package_resource_package_association", "test")
	r := AccessPackageResourcePackageAssociationResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config:  r.completeWithAccessTypeApplication(data),
			Destroy: false,
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("access_type").IsUuid(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccAccessPackageResourcePackageAssociation_completeWithAccessTypeApplicationRoleName(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_access_package_resource_package_association", "test")
	r := AccessPackageResourcePackageAssociationResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config:  r.completeWithAccessTypeApplicationRoleName(data),
			Destroy: false,
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("access_type").HasValue("TestRole"),
			),
		},
		// An imported association has no configuration to compare against, so the role is read
		// back under its originId rather than the display name the configuration named it by.
		data.ImportStep("access_type"),
	})
}

// A SharePoint site can't be provisioned by the acceptance tests, so this is skipped unless
// AZUREAD_TEST_SHAREPOINT_SITE_URL and AZUREAD_TEST_SHAREPOINT_ROLE_ORIGIN_ID name an existing
// site and a role on it.
func TestAccAccessPackageResourcePackageAssociation_completeWithSharePointSite(t *testing.T) {
	siteURL := os.Getenv("AZUREAD_TEST_SHAREPOINT_SITE_URL")
	roleOriginID := os.Getenv("AZUREAD_TEST_SHAREPOINT_ROLE_ORIGIN_ID")
	if siteURL == "" || roleOriginID == "" {
		t.Skipf("Acceptance test skipped unless env 'AZUREAD_TEST_SHAREPOINT_SITE_URL' and 'AZUREAD_TEST_SHAREPOINT_ROLE_ORIGIN_ID' set")
	}

	data := acceptance.BuildTestData(t, "azuread_access_package_resource_package_association", "test")
	r := AccessPackageResourcePackageAssociationResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config:  r.completeWithSharePointSite(data, siteURL, roleOriginID),
			Destroy: false,
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("access_type").HasValue(roleOriginID),
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
  access_type                     = "Member"
}
`, data.RandomInteger)
}

func (AccessPackageResourcePackageAssociationResource) completeWithAccessTypeOwner(data acceptance.TestData) string {
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
  access_type                     = "Owner"
}
`, data.RandomInteger)
}

func (AccessPackageResourcePackageAssociationResource) completeWithAccessTypeApplication(data acceptance.TestData) string {
	// The application defines its own app role and that role's id is used as the
	// access_type. This keeps the fixture self-contained: it doesn't read or update
	// the well-known Microsoft Graph service principal (which needs privileges to
	// patch a first-party SP), so it runs with Application.ReadWrite.All alone. It
	// also mirrors reality: an AadApplication resource's roles are the application's
	// own app roles. The role id is a fixed UUID reused as the access_type.
	const appRoleId = "9f3c8e21-6b4d-4f2a-8d7c-1e5a9b0c3d4e"

	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application" "test" {
  display_name = "acctest-packageAssociationResource-%[1]d"

  app_role {
    allowed_member_types = ["User"]
    description          = "Test role for access package resource association"
    display_name         = "TestRole"
    enabled              = true
    id                   = "%[2]s"
    value                = "Test.Role"
  }
}

resource "azuread_service_principal" "test" {
  client_id = azuread_application.test.client_id
}

# Test-only: entitlement management can't find a service principal that was just
# created until it has replicated. Real usage onboards a pre-existing application,
# so this delay is an artifact of the test creating its own SP inline, not a
# limitation of the resource.
resource "time_sleep" "wait_for_sp" {
  depends_on      = [azuread_service_principal.test]
  create_duration = "60s"
}

resource "azuread_access_package_catalog" "test_catalog" {
  display_name = "acctest-catalog-%[1]d"
  description  = "Test catalog %[1]d"
}

resource "azuread_access_package_resource_catalog_association" "test" {
  catalog_id             = azuread_access_package_catalog.test_catalog.id
  resource_origin_id     = azuread_service_principal.test.object_id
  resource_origin_system = "AadApplication"

  depends_on = [time_sleep.wait_for_sp]
}

resource "azuread_access_package" "test" {
  display_name = "acctest-package-%[1]d"
  description  = "Test Package %[1]d"
  catalog_id   = azuread_access_package_catalog.test_catalog.id
}

resource "azuread_access_package_resource_package_association" "test" {
  access_package_id               = azuread_access_package.test.id
  catalog_resource_association_id = azuread_access_package_resource_catalog_association.test.id
  access_type                     = "%[2]s"
}
`, data.RandomInteger, appRoleId)
}

func (AccessPackageResourcePackageAssociationResource) completeWithAccessTypeApplicationRoleName(data acceptance.TestData) string {
	// The same app role as completeWithAccessTypeApplication, attached by the display name it is
	// shown under instead of its ID.
	const appRoleId = "9f3c8e21-6b4d-4f2a-8d7c-1e5a9b0c3d4e"

	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application" "test" {
  display_name = "acctest-packageAssociationResource-%[1]d"

  app_role {
    allowed_member_types = ["User"]
    description          = "Test role for access package resource association"
    display_name         = "TestRole"
    enabled              = true
    id                   = "%[2]s"
    value                = "Test.Role"
  }
}

resource "azuread_service_principal" "test" {
  client_id = azuread_application.test.client_id
}

# Entitlement management can't see a service principal until it has replicated.
resource "time_sleep" "wait_for_sp" {
  depends_on      = [azuread_service_principal.test]
  create_duration = "60s"
}

resource "azuread_access_package_catalog" "test_catalog" {
  display_name = "acctest-catalog-%[1]d"
  description  = "Test catalog %[1]d"
}

resource "azuread_access_package_resource_catalog_association" "test" {
  catalog_id             = azuread_access_package_catalog.test_catalog.id
  resource_origin_id     = azuread_service_principal.test.object_id
  resource_origin_system = "AadApplication"

  depends_on = [time_sleep.wait_for_sp]
}

resource "azuread_access_package" "test" {
  display_name = "acctest-package-%[1]d"
  description  = "Test Package %[1]d"
  catalog_id   = azuread_access_package_catalog.test_catalog.id
}

resource "azuread_access_package_resource_package_association" "test" {
  access_package_id               = azuread_access_package.test.id
  catalog_resource_association_id = azuread_access_package_resource_catalog_association.test.id
  access_type                     = "TestRole"
}
`, data.RandomInteger, appRoleId)
}

func (AccessPackageResourcePackageAssociationResource) completeWithSharePointSite(data acceptance.TestData, siteURL, roleOriginID string) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_access_package_catalog" "test_catalog" {
  display_name = "acctest-catalog-%[1]d"
  description  = "Test catalog %[1]d"
}

resource "azuread_access_package_resource_catalog_association" "test" {
  catalog_id             = azuread_access_package_catalog.test_catalog.id
  resource_origin_id     = %[2]q
  resource_origin_system = "SharePointOnline"
}

resource "azuread_access_package" "test" {
  display_name = "acctest-package-%[1]d"
  description  = "Test Package %[1]d"
  catalog_id   = azuread_access_package_catalog.test_catalog.id
}

resource "azuread_access_package_resource_package_association" "test" {
  access_package_id               = azuread_access_package.test.id
  catalog_resource_association_id = azuread_access_package_resource_catalog_association.test.id
  access_type                     = %[3]q
}
`, data.RandomInteger, siteURL, roleOriginID)
}
