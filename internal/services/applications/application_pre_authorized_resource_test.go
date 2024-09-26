// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applications_test

import (
	"context"
	"fmt"
	"strings"
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

type ApplicationPreAuthorizedResource struct{}

func TestAccApplicationPreAuthorized_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_pre_authorized", "test")
	r := ApplicationPreAuthorizedResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("authorized_client_id").Exists(),
				check.That(data.ResourceName).Key("permission_ids.#").HasValue("2"),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApplicationPreAuthorized_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_pre_authorized", "test")
	r := ApplicationPreAuthorizedResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.RequiresImportErrorStep(r.requiresImport(data)),
	})
}

func TestAccApplicationPreAuthorized_multipleCreateDestroy(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_pre_authorized", "authorize_1")
	data2 := acceptance.BuildTestData(t, "azuread_application", "authorizer")
	r := ApplicationPreAuthorizedResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.multiple(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.multipleDestroy(data2),
		},
		{
			// This step should catch any failed destroys from the previous step by throwing an ImportAsExists error
			Config: r.multiple(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func (ApplicationPreAuthorizedResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.Applications.ApplicationClient

	id, err := parse.ApplicationPreAuthorizedID(state.ID)
	if err != nil {
		return nil, fmt.Errorf("parsing Pre-Authorized Application ID: %v", err)
	}

	applicationId := stable.NewApplicationID(id.ObjectId)

	resp, err := client.GetApplication(ctx, applicationId, application.DefaultGetApplicationOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return nil, fmt.Errorf("Application with object ID %q does not exist", id.ObjectId)
		}
		return nil, fmt.Errorf("failed to retrieve %s: %+v", applicationId, err)
	}

	app := resp.Model
	if app != nil && app.Api != nil && app.Api.PreAuthorizedApplications != nil {
		for _, a := range *app.Api.PreAuthorizedApplications {
			if strings.EqualFold(a.AppId.GetOrZero(), id.AppId) {
				return pointer.To(true), nil
			}
		}
	}

	return nil, fmt.Errorf("Pre-Authorized Application %q was not found for Application %q", id.AppId, id.ObjectId)
}

func (ApplicationPreAuthorizedResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_application" "authorized" {
  display_name = "acctestApp-authorized-%[1]d"
}

resource "azuread_application" "authorizer" {
  display_name = "acctestApp-authorizer-%[1]d"

  api {
    oauth2_permission_scope {
      admin_consent_description  = "Administer the application"
      admin_consent_display_name = "Administer"
      enabled                    = true
      id                         = "%[2]s"
      type                       = "Admin"
      value                      = "administer"
    }

    oauth2_permission_scope {
      admin_consent_description  = "Access the application"
      admin_consent_display_name = "Access"
      enabled                    = true
      id                         = "%[3]s"
      type                       = "User"
      user_consent_description   = "Access the application"
      user_consent_display_name  = "Access"
      value                      = "user_impersonation"
    }
  }
}

resource "azuread_application_pre_authorized" "test" {
  application_id       = azuread_application.authorizer.id
  authorized_client_id = azuread_application.authorized.client_id
  permission_ids       = ["%[2]s", "%[3]s"]
}
`, data.RandomInteger, data.UUID(), data.UUID())
}

func (r ApplicationPreAuthorizedResource) requiresImport(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_application_pre_authorized" "import" {
  application_id       = azuread_application_pre_authorized.test.application_id
  authorized_client_id = azuread_application_pre_authorized.test.authorized_client_id
  permission_ids       = ["%[2]s"]
}
`, r.basic(data), data.UUID())
}

func (ApplicationPreAuthorizedResource) multiple(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_application" "authorized_1" {
  display_name = "acctestApp-authorized-1-%[1]d"
}

resource "azuread_application" "authorized_2" {
  display_name = "acctestApp-authorized-2-%[1]d"
}

resource "azuread_application" "authorized_3" {
  display_name = "acctestApp-authorized-3-%[1]d"
}

resource "azuread_application" "authorizer" {
  display_name = "acctestApp-authorizer-%[1]d"

  api {
    oauth2_permission_scope {
      admin_consent_description  = "Administer the application"
      admin_consent_display_name = "Administer"
      enabled                    = true
      id                         = "11111111-1111-1111-1111-111111111111"
      type                       = "Admin"
      value                      = "administer"
    }
  }
}

resource "azuread_application_pre_authorized" "authorize_1" {
  application_id       = azuread_application.authorizer.id
  authorized_client_id = azuread_application.authorized_1.client_id
  permission_ids       = ["11111111-1111-1111-1111-111111111111"]
}

resource "azuread_application_pre_authorized" "authorize_2" {
  application_id       = azuread_application.authorizer.id
  authorized_client_id = azuread_application.authorized_2.client_id
  permission_ids       = ["11111111-1111-1111-1111-111111111111"]
}

resource "azuread_application_pre_authorized" "authorize_3" {
  application_id       = azuread_application.authorizer.id
  authorized_client_id = azuread_application.authorized_3.client_id
  permission_ids       = ["11111111-1111-1111-1111-111111111111"]
}
`, data.RandomInteger)
}

func (ApplicationPreAuthorizedResource) multipleDestroy(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_application" "authorized_1" {
  display_name = "acctestApp-authorized-1-%[1]d"
}

resource "azuread_application" "authorized_2" {
  display_name = "acctestApp-authorized-2-%[1]d"
}

resource "azuread_application" "authorized_3" {
  display_name = "acctestApp-authorized-3-%[1]d"
}

resource "azuread_application" "authorizer" {
  display_name = "acctestApp-authorizer-%[1]d"

  api {
    oauth2_permission_scope {
      admin_consent_description  = "Administer the application"
      admin_consent_display_name = "Administer"
      enabled                    = true
      id                         = "11111111-1111-1111-1111-111111111111"
      type                       = "Admin"
      value                      = "administer"
    }
  }
}
`, data.RandomInteger)
}
