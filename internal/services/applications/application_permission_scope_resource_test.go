// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applications_test

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/applications/parse"
)

type ApplicationPermissionScopeResource struct{}

func TestAccApplicationPermissionScope_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_permission_scope", "test")
	r := ApplicationPermissionScopeResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
				check.That(data.ResourceName).Key("scope_id").Exists(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApplicationPermissionScope_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_permission_scope", "test")
	r := ApplicationPermissionScopeResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.complete(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
				check.That(data.ResourceName).Key("scope_id").Exists(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApplicationPermissionScope_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_permission_scope", "test")
	r := ApplicationPermissionScopeResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
				check.That(data.ResourceName).Key("scope_id").Exists(),
			),
		},
		data.ImportStep(),
		{
			Config: r.complete(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
				check.That(data.ResourceName).Key("scope_id").Exists(),
			),
		},
		data.ImportStep(),
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
				check.That(data.ResourceName).Key("scope_id").Exists(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApplicationPermissionScope_multiple(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_permission_scope", "test")
	data2 := acceptance.BuildTestData(t, "azuread_application_permission_scope", "test2")
	r := ApplicationPermissionScopeResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.multiple(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
				check.That(data.ResourceName).Key("scope_id").Exists(),
				check.That(data2.ResourceName).ExistsInAzure(r),
				check.That(data2.ResourceName).Key("application_id").Exists(),
				check.That(data2.ResourceName).Key("scope_id").Exists(),
			),
		},
		data.ImportStep(),
		data2.ImportStep(),
	})
}

func TestAccApplicationPermissionScope_multipleUpdate(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_permission_scope", "test")
	data2 := acceptance.BuildTestData(t, "azuread_application_permission_scope", "test2")
	r := ApplicationPermissionScopeResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.multiple(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
				check.That(data.ResourceName).Key("scope_id").Exists(),
				check.That(data2.ResourceName).ExistsInAzure(r),
				check.That(data2.ResourceName).Key("application_id").Exists(),
				check.That(data2.ResourceName).Key("scope_id").Exists(),
			),
		},
		data.ImportStep(),
		data2.ImportStep(),
		{
			Config: r.multipleUpdate(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
				check.That(data.ResourceName).Key("scope_id").Exists(),
				check.That(data2.ResourceName).ExistsInAzure(r),
				check.That(data2.ResourceName).Key("application_id").Exists(),
				check.That(data2.ResourceName).Key("scope_id").Exists(),
			),
		},
		data.ImportStep(),
		data2.ImportStep(),
		{
			Config: r.multiple(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
				check.That(data.ResourceName).Key("scope_id").Exists(),
				check.That(data2.ResourceName).ExistsInAzure(r),
				check.That(data2.ResourceName).Key("application_id").Exists(),
				check.That(data2.ResourceName).Key("scope_id").Exists(),
			),
		},
		data.ImportStep(),
		data2.ImportStep(),
	})
}

func TestAccApplicationPermissionScope_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_permission_scope", "test")
	r := ApplicationPermissionScopeResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
				check.That(data.ResourceName).Key("scope_id").Exists(),
			),
		},
		data.RequiresImportErrorStep(r.requiresImport(data)),
	})
}

func (r ApplicationPermissionScopeResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.Applications.ApplicationsClient
	client.BaseClient.DisableRetries = true
	defer func() { client.BaseClient.DisableRetries = false }()

	id, err := parse.ParsePermissionScopeID(state.ID)
	if err != nil {
		return nil, err
	}

	result, status, err := client.Get(ctx, id.ApplicationId, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return pointer.To(false), nil
		}
		return nil, fmt.Errorf("retrieving %s: %+v", id, err)
	}
	if result == nil {
		return nil, fmt.Errorf("retrieving %s: result was nil", id)
	}

	if result.Api == nil || result.Api.OAuth2PermissionScopes == nil {
		return pointer.To(false), nil
	}

	for _, scope := range *result.Api.OAuth2PermissionScopes {
		if strings.EqualFold(*scope.ID, id.ScopeID) {
			return pointer.To(true), nil
		}
	}

	return pointer.To(false), nil
}

func (ApplicationPermissionScopeResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application_registration" "test" {
  display_name = "acctest-AppRegistration-%[1]d"
}

resource "azuread_application_permission_scope" "test" {
  application_id = azuread_application_registration.test.id
  scope_id       = "%[2]s"
  value          = "administer"

  admin_consent_description  = "Administer the application"
  admin_consent_display_name = "Administer"
}
`, data.RandomInteger, data.RandomID)
}

func (ApplicationPermissionScopeResource) complete(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application_registration" "test" {
  display_name = "acctest-AppRegistration-%[1]d"
}

resource "azuread_application_permission_scope" "test" {
  application_id = azuread_application_registration.test.id
  scope_id       = "%[2]s"
  type           = "Admin"
  value          = "administer"

  admin_consent_description  = "Administer the application"
  admin_consent_display_name = "Administer"
  user_consent_description   = "Allow the application to access acctest-APP-%[1]d on your behalf."
  user_consent_display_name  = "Access acctest-APP-%[1]d"
}
`, data.RandomInteger, data.RandomID)
}

func (ApplicationPermissionScopeResource) multiple(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application_registration" "test" {
  display_name = "acctest-AppRegistration-%[1]d"
}

resource "random_uuid" "test" {}
resource "random_uuid" "test2" {}

resource "azuread_application_permission_scope" "test" {
  application_id = azuread_application_registration.test.id
  scope_id       = random_uuid.test.id
  value          = "administer"

  admin_consent_description  = "Administer the application"
  admin_consent_display_name = "Administer"
}

resource "azuread_application_permission_scope" "test2" {
  application_id = azuread_application_registration.test.id
  scope_id       = random_uuid.test2.id
  value          = "operate"

  admin_consent_description  = "Operate the application"
  admin_consent_display_name = "Operate"
}
`, data.RandomInteger)
}

func (ApplicationPermissionScopeResource) multipleUpdate(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application_registration" "test" {
  display_name = "acctest-AppRegistration-%[1]d"
}

resource "random_uuid" "test" {}
resource "random_uuid" "test2" {}

resource "azuread_application_permission_scope" "test" {
  application_id = azuread_application_registration.test.id
  scope_id       = random_uuid.test.id
  value          = "administrate"

  admin_consent_description  = "Administrate the application"
  admin_consent_display_name = "Administrate"
}

resource "azuread_application_permission_scope" "test2" {
  application_id = azuread_application_registration.test.id
  scope_id       = random_uuid.test2.id
  value          = "view"

  admin_consent_description  = "View the application"
  admin_consent_display_name = "View"
}
`, data.RandomInteger)
}

func (ApplicationPermissionScopeResource) requiresImport(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application_registration" "test" {
  display_name = "acctest-AppRegistration-%[1]d"
}

resource "azuread_application_permission_scope" "test" {
  application_id = azuread_application_registration.test.id
  scope_id       = "%[2]s"
  value          = "administer"

  admin_consent_description  = "Administer the application"
  admin_consent_display_name = "Administer"
}

resource "azuread_application_permission_scope" "import" {
  application_id = azuread_application_permission_scope.test.application_id
  scope_id       = azuread_application_permission_scope.test.scope_id
  value          = azuread_application_permission_scope.test.value

  admin_consent_description  = azuread_application_permission_scope.test.admin_consent_description
  admin_consent_display_name = azuread_application_permission_scope.test.admin_consent_display_name
}
`, data.RandomInteger, data.RandomID)
}
