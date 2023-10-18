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
	"github.com/hashicorp/terraform-provider-azuread/internal/services/applications"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/applications/parse"
)

type ApplicationRedirectUrisResource struct{}

func TestAccApplicationRedirectUris_publicClient(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_redirect_uris", "test")
	r := ApplicationRedirectUrisResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.publicClient(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApplicationRedirectUris_spa(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_redirect_uris", "test")
	r := ApplicationRedirectUrisResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.spa(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApplicationRedirectUris_web(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_redirect_uris", "test")
	r := ApplicationRedirectUrisResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.web(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApplicationRedirectUris_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_redirect_uris", "test")
	r := ApplicationRedirectUrisResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.web(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
			),
		},
		data.RequiresImportErrorStep(r.requiresImport(data)),
	})
}

func (r ApplicationRedirectUrisResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.Applications.ApplicationsClient
	client.BaseClient.DisableRetries = true
	defer func() { client.BaseClient.DisableRetries = false }()

	id, err := parse.ParseRedirectUrisID(state.ID)
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

	switch id.UriType {
	case applications.RedirectUriTypePublicClient:
		if result.PublicClient != nil && result.PublicClient.RedirectUris != nil && len(*result.PublicClient.RedirectUris) > 0 {
			return pointer.To(true), nil
		}
	case applications.RedirectUriTypeSPA:
		if result.Spa != nil && result.Spa.RedirectUris != nil && len(*result.Spa.RedirectUris) > 0 {
			return pointer.To(true), nil
		}
	case applications.RedirectUriTypeWeb:
		if result.Web != nil && result.Web.RedirectUris != nil && len(*result.Web.RedirectUris) > 0 {
			return pointer.To(true), nil
		}
	}

	return pointer.To(false), nil
}

func (ApplicationRedirectUrisResource) publicClient(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application_registration" "test" {
  display_name = "acctest-RedirectUris-%[1]d"
}

resource "azuread_application_redirect_uris" "test" {
  application_id = azuread_application_registration.test.id
  type           = "PublicClient"

  redirect_uris = [
    "myapp://auth",
    "sample.mobile.app.bundie.id://auth",
    "https://login.microsoftonline.com/common/oauth2/nativeclient",
    "https://login.live.com/oauth20_desktop.srf",
    "ms-appx-web://Microsoft.AAD.BrokerPlugin/00000000-1111-1111-1111-222222222222",
    "urn:ietf:wg:oauth:2.0:foo",
  ]
}
`, data.RandomInteger)
}

func (ApplicationRedirectUrisResource) spa(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application_registration" "test" {
  display_name = "acctest-RedirectUris-%[1]d"
}

resource "azuread_application_redirect_uris" "test" {
  application_id = azuread_application_registration.test.id
  type           = "SPA"

  redirect_uris = [
    "https://mobile.hashitown-%[1]d.com/",
    "https://beta.hashitown-%[1]d.com/",
  ]
}
`, data.RandomInteger)
}

func (ApplicationRedirectUrisResource) web(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application_registration" "test" {
  display_name = "acctest-RedirectUris-%[1]d"
}

resource "azuread_application_redirect_uris" "test" {
  application_id = azuread_application_registration.test.id
  type           = "Web"

  redirect_uris = [
    "https://app.hashitown-%[1]d.com/",
    "https://classic.hashitown-%[1]d.com/",
    "urn:ietf:wg:oauth:2.0:oob",
  ]
}
`, data.RandomInteger)
}

func (r ApplicationRedirectUrisResource) requiresImport(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_application_redirect_uris" "import" {
  application_id = azuread_application_redirect_uris.test.application_id
  type           = azuread_application_redirect_uris.test.type
  redirect_uris  = azuread_application_redirect_uris.test.redirect_uris
}
`, r.web(data))
}
