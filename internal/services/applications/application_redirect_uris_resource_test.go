// Copyright IBM Corp. 2014, 2025
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

func TestAccApplicationRedirectUris_all(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_redirect_uris", "test_public")
	data2 := acceptance.BuildTestData(t, "azuread_application_redirect_uris", "test_spa")
	data3 := acceptance.BuildTestData(t, "azuread_application_redirect_uris", "test_web")
	r := ApplicationRedirectUrisResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.all(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
				check.That(data2.ResourceName).ExistsInAzure(r),
				check.That(data2.ResourceName).Key("application_id").Exists(),
				check.That(data3.ResourceName).ExistsInAzure(r),
				check.That(data3.ResourceName).Key("application_id").Exists(),
			),
		},
		data.ImportStep(),
		data2.ImportStep(),
		data3.ImportStep(),
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
	client := clients.Applications.ApplicationClient

	id, err := parse.ParseRedirectUrisID(state.ID)
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
		return nil, fmt.Errorf("retrieving %s: app was nil", id)
	}

	switch id.UriType {
	case applications.RedirectUriTypePublicClient:
		if app.PublicClient != nil && app.PublicClient.RedirectUris != nil && len(*app.PublicClient.RedirectUris) > 0 {
			return pointer.To(true), nil
		}
	case applications.RedirectUriTypeSPA:
		if app.Spa != nil && app.Spa.RedirectUris != nil && len(*app.Spa.RedirectUris) > 0 {
			return pointer.To(true), nil
		}
	case applications.RedirectUriTypeWeb:
		if app.Web != nil && app.Web.RedirectUris != nil && len(*app.Web.RedirectUris) > 0 {
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
    "https://mobile.hashitown.example.com-%[1]d.com/",
    "https://beta.hashitown.example.com-%[1]d.com/",
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
    "https://app.hashitown.example.com-%[1]d.com/",
    "https://classic.hashitown.example.com-%[1]d.com/",
    "urn:ietf:wg:oauth:2.0:oob",
  ]
}
`, data.RandomInteger)
}

func (ApplicationRedirectUrisResource) all(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application_registration" "test" {
  display_name = "acctest-RedirectUris-%[1]d"
}

resource "azuread_application_redirect_uris" "test_public" {
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

resource "azuread_application_redirect_uris" "test_spa" {
  application_id = azuread_application_registration.test.id
  type           = "SPA"

  redirect_uris = [
    "https://mobile.hashitown.example.com-%[1]d.com/",
    "https://beta.hashitown.example.com-%[1]d.com/",
  ]
}

resource "azuread_application_redirect_uris" "test_web" {
  application_id = azuread_application_registration.test.id
  type           = "Web"

  redirect_uris = [
    "https://app.hashitown.example.com-%[1]d.com/",
    "https://classic.hashitown.example.com-%[1]d.com/",
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
