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

type ApplicationFromTemplateResource struct{}

func TestAccApplicationFromTemplate_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_from_template", "test")
	r := ApplicationFromTemplateResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("template_id").Exists(),
				check.That(data.ResourceName).Key("application_id").Exists(),
				check.That(data.ResourceName).Key("application_object_id").Exists(),
				check.That(data.ResourceName).Key("service_principal_id").Exists(),
				check.That(data.ResourceName).Key("service_principal_object_id").Exists(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApplicationFromTemplate_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_from_template", "test")
	r := ApplicationFromTemplateResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("template_id").Exists(),
				check.That(data.ResourceName).Key("application_id").Exists(),
				check.That(data.ResourceName).Key("application_object_id").Exists(),
				check.That(data.ResourceName).Key("service_principal_id").Exists(),
				check.That(data.ResourceName).Key("service_principal_object_id").Exists(),
			),
		},
		data.ImportStep(),
		{
			Config: r.update(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("template_id").Exists(),
				check.That(data.ResourceName).Key("application_id").Exists(),
				check.That(data.ResourceName).Key("application_object_id").Exists(),
				check.That(data.ResourceName).Key("service_principal_id").Exists(),
				check.That(data.ResourceName).Key("service_principal_object_id").Exists(),
			),
		},
		data.ImportStep(),
	})
}

func (r ApplicationFromTemplateResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.Applications.ApplicationClient

	id, err := parse.ParseFromTemplateID(state.ID)
	if err != nil {
		return nil, err
	}

	applicationId := stable.NewApplicationID(id.ApplicationId)

	// Check the application exists
	resp, err := client.GetApplication(ctx, applicationId, application.DefaultGetApplicationOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return pointer.To(false), nil
		}
		return nil, fmt.Errorf("retrieving %s: %+v", id, err)
	}

	return pointer.To(true), nil
}

func (ApplicationFromTemplateResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application_from_template" "test" {
  display_name = "acctest-FromTemplate-%[1]d"
  template_id  = "%[2]s"
}

data "azuread_application" "test" {
  object_id = azuread_application_from_template.test.application_object_id
}

data "azuread_service_principal" "test" {
  object_id = azuread_application_from_template.test.service_principal_object_id
}
`, data.RandomInteger, testApplicationTemplateId)
}

func (ApplicationFromTemplateResource) update(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application_from_template" "test" {
  display_name = "acctest-FromTemplateUpdated-%[1]d"
  template_id  = "%[2]s"
}

data "azuread_application" "test" {
  object_id = azuread_application_from_template.test.application_object_id
}

data "azuread_service_principal" "test" {
  object_id = azuread_application_from_template.test.service_principal_object_id
}
`, data.RandomInteger, testApplicationTemplateId)
}
