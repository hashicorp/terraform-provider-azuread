// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applications_test

import (
	"context"
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

type ApplicationFallbackPublicClientResource struct{}

func TestAccApplicationFallbackPublicClient_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_fallback_public_client", "test")
	r := ApplicationFallbackPublicClientResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.enabled(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
				check.That(data.ResourceName).Key("enabled").HasValue("true"),
			),
		},
		data.ImportStep(),
	})
}
func TestAccApplicationFallbackPublicClient_disabled(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_fallback_public_client", "test")
	r := ApplicationFallbackPublicClientResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.disabled(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
				check.That(data.ResourceName).Key("enabled").HasValue("false"),
			),
		},
		data.ImportStep(),
	})
}

func (r ApplicationFallbackPublicClientResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.Applications.ApplicationClient

	id, err := parse.ParseFallbackPublicClientID(state.ID)
	if err != nil {
		return nil, err
	}

	applicationId := stable.NewApplicationID(id.ApplicationId)

	resp, err := client.GetApplication(ctx, applicationId, application.DefaultGetApplicationOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return pointer.To(false), nil
		}
		return nil, fmt.Errorf("retrieving %s: %+v", applicationId, err)
	}

	app := resp.Model
	if app == nil {
		return nil, fmt.Errorf("retrieving %s: model was nil", applicationId)
	}

	if app.IsFallbackPublicClient == nil {
		return pointer.To(false), nil
	}

	return pointer.To(true), nil
}

func (ApplicationFallbackPublicClientResource) enabled(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application_registration" "test" {
  display_name = "acctest-FallbackPublicClient-%[1]d"
}

resource "azuread_application_fallback_public_client" "test" {
  application_id = azuread_application_registration.test.id
  enabled        = true
}
`, data.RandomInteger, data.RandomID)
}

func (ApplicationFallbackPublicClientResource) disabled(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application_registration" "test" {
  display_name = "acctest-FallbackPublicClient-%[1]d"
}

resource "azuread_application_fallback_public_client" "test" {
  application_id = azuread_application_registration.test.id
  enabled        = false
}
`, data.RandomInteger)
}
