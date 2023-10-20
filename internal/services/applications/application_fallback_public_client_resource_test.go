// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applications_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/applications/parse"
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

func TestAccApplicationFallbackPublicClient_update(t *testing.T) {
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
		{
			Config: r.enabled(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
				check.That(data.ResourceName).Key("enabled").HasValue("true"),
			),
		},
		data.ImportStep(),
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
	client := clients.Applications.ApplicationsClient
	client.BaseClient.DisableRetries = true
	defer func() { client.BaseClient.DisableRetries = false }()

	id, err := parse.ParseFallbackPublicClientID(state.ID)
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

	if result.IsFallbackPublicClient == nil {
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
