// Copyright IBM Corp. 2019, 2025
// SPDX-License-Identifier: MPL-2.0

package applications_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/application"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/applications/parse"
)

type ApplicationKnownClientsResource struct{}

func TestAccApplicationKnownClients_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_known_clients", "test")
	r := ApplicationKnownClientsResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApplicationKnownClients_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_known_clients", "test")
	r := ApplicationKnownClientsResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
			),
		},
		data.ImportStep(),
		{
			Config: r.update(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
			),
		},
		data.ImportStep(),
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApplicationKnownClients_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_known_clients", "test")
	r := ApplicationKnownClientsResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
			),
		},
		data.RequiresImportErrorStep(r.requiresImport(data)),
	})
}

func (r ApplicationKnownClientsResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.Applications.ApplicationClient

	id, err := parse.ParseKnownClientsID(state.ID)
	if err != nil {
		return nil, err
	}

	applicationId := stable.NewApplicationID(id.ApplicationId)

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

	if app.Api != nil && app.Api.KnownClientApplications != nil && len(*app.Api.KnownClientApplications) > 0 {
		return pointer.To(true), nil
	}

	return pointer.To(false), nil
}

func (ApplicationKnownClientsResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application_registration" "test" {
  display_name = "acctest-KnownClients-%[1]d"
}

resource "azuread_application_registration" "client" {
  display_name = "acctest-Client-%[1]d"
}

resource "azuread_application_registration" "client2" {
  display_name = "acctest-Client2-%[1]d"
}

resource "azuread_application_known_clients" "test" {
  application_id = azuread_application_registration.test.id
  known_client_ids = [
    azuread_application_registration.client.client_id,
    azuread_application_registration.client2.client_id,
  ]
}
`, data.RandomInteger, data.RandomPassword)
}

func (ApplicationKnownClientsResource) update(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application_registration" "test" {
  display_name = "acctest-KnownClients-%[1]d"
}

resource "azuread_application_registration" "client" {
  display_name = "acctest-Client-%[1]d"
}

resource "azuread_application_registration" "client2" {
  display_name = "acctest-Client2-%[1]d"
}

resource "azuread_application_known_clients" "test" {
  application_id = azuread_application_registration.test.id
  known_client_ids = [
    azuread_application_registration.client2.client_id,
  ]
}
`, data.RandomInteger, data.RandomPassword)
}

func (ApplicationKnownClientsResource) requiresImport(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application_registration" "test" {
  display_name = "acctest-KnownClients-%[1]d"
}

resource "azuread_application_registration" "client" {
  display_name = "acctest-Client-%[1]d"
}

resource "azuread_application_known_clients" "test" {
  application_id   = azuread_application_registration.test.id
  known_client_ids = [azuread_application_registration.client.client_id]
}

resource "azuread_application_known_clients" "import" {
  application_id   = azuread_application_known_clients.test.application_id
  known_client_ids = azuread_application_known_clients.test.known_client_ids
}
`, data.RandomInteger, data.RandomPassword)
}
