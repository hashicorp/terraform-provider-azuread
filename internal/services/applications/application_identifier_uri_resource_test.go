// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applications_test

import (
	"context"
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/glueckkanja/terraform-provider-azuread/internal/acceptance"
	"github.com/glueckkanja/terraform-provider-azuread/internal/acceptance/check"
	"github.com/glueckkanja/terraform-provider-azuread/internal/clients"
	"github.com/glueckkanja/terraform-provider-azuread/internal/services/applications/parse"
	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/application"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

type ApplicationIdentifierUriResource struct{}

func TestAccApplicationIdentifierUri_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_identifier_uri", "test")
	r := ApplicationIdentifierUriResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
				check.That(data.ResourceName).Key("identifier_uri").Exists(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApplicationIdentifierUri_multiple(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_identifier_uri", "test")
	data2 := acceptance.BuildTestData(t, "azuread_application_identifier_uri", "test2")
	r := ApplicationIdentifierUriResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.multiple(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
				check.That(data.ResourceName).Key("identifier_uri").Exists(),
				check.That(data2.ResourceName).ExistsInAzure(r),
				check.That(data2.ResourceName).Key("application_id").Exists(),
				check.That(data2.ResourceName).Key("identifier_uri").Exists(),
			),
		},
		data.ImportStep(),
		data2.ImportStep(),
	})
}

func TestAccApplicationIdentifierUri_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_identifier_uri", "test")
	r := ApplicationIdentifierUriResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
				check.That(data.ResourceName).Key("identifier_uri").Exists(),
			),
		},
		data.RequiresImportErrorStep(r.requiresImport(data)),
	})
}

func (r ApplicationIdentifierUriResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.Applications.ApplicationClient

	id, err := parse.ParseIdentifierUriID(state.ID)
	if err != nil {
		return nil, err
	}

	applicationId := stable.NewApplicationID(id.ApplicationId)

	uriFromIdSegment, err := base64.StdEncoding.DecodeString(id.IdentifierUri)
	if err != nil {
		return nil, fmt.Errorf("failed to decode identifierUri from resource ID: %+v", err)
	}

	resp, err := client.GetApplication(ctx, applicationId, application.DefaultGetApplicationOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return pointer.To(false), nil
		}
		return nil, fmt.Errorf("retrieving %s: %+v", id, err)
	}

	app := resp.Model
	if app == nil {
		return nil, fmt.Errorf("retrieving %s: model was nil", id)
	}

	if app.IdentifierUris == nil {
		return pointer.To(false), nil
	}

	for _, existingUri := range *app.IdentifierUris {
		if existingUri == string(uriFromIdSegment) {
			return pointer.To(true), nil
		}
	}

	return pointer.To(false), nil
}

func (ApplicationIdentifierUriResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application_registration" "test" {
  display_name = "acctest-AppRegistration-%[1]d"
}

resource "azuread_application_identifier_uri" "test" {
  application_id = azuread_application_registration.test.id
  identifier_uri = "api://hashicorptestapp-%[1]d"
}
`, data.RandomInteger)
}

func (ApplicationIdentifierUriResource) multiple(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application_registration" "test" {
  display_name = "acctest-AppRegistration-%[1]d"
}

resource "azuread_application_identifier_uri" "test" {
  application_id = azuread_application_registration.test.id
  identifier_uri = "api://hashicorptestapp-%[1]d"
}

resource "azuread_application_identifier_uri" "test2" {
  application_id = azuread_application_registration.test.id
  identifier_uri = "api://${azuread_application_registration.test.client_id}"
}
`, data.RandomInteger)
}

func (ApplicationIdentifierUriResource) requiresImport(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application_registration" "test" {
  display_name = "acctest-AppRegistration-%[1]d"
}

resource "azuread_application_identifier_uri" "test" {
  application_id = azuread_application_registration.test.id
  identifier_uri = "api://hashicorptestapp-%[1]d"
}

resource "azuread_application_identifier_uri" "import" {
  application_id = azuread_application_identifier_uri.test.application_id
  identifier_uri = azuread_application_identifier_uri.test.identifier_uri
}
`, data.RandomInteger, data.RandomID)
}
