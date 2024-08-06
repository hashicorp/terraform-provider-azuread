// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applications_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/applications/parse"
)

type ApplicationPasswordResource struct{}

func TestAccApplicationPassword_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_password", "test")
	r := ApplicationPasswordResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("end_date").Exists(),
				check.That(data.ResourceName).Key("key_id").Exists(),
				check.That(data.ResourceName).Key("start_date").Exists(),
				check.That(data.ResourceName).Key("value").Exists(),
			),
		},
	})
}

func TestAccApplicationPassword_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_password", "test")
	startDate := time.Now().AddDate(0, 0, 7).UTC().Format(time.RFC3339)
	endDate := time.Now().AddDate(0, 5, 27).UTC().Format(time.RFC3339)
	r := ApplicationPasswordResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.complete(data, startDate, endDate),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("end_date").Exists(),
				check.That(data.ResourceName).Key("key_id").Exists(),
				check.That(data.ResourceName).Key("start_date").Exists(),
				check.That(data.ResourceName).Key("value").Exists(),
			),
		},
	})
}

func TestAccApplicationPassword_relativeEndDate(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_password", "test")
	r := ApplicationPasswordResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.relativeEndDate(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("end_date").Exists(),
				check.That(data.ResourceName).Key("end_date_relative").HasValue("8760h"),
				check.That(data.ResourceName).Key("key_id").Exists(),
				check.That(data.ResourceName).Key("start_date").Exists(),
				check.That(data.ResourceName).Key("value").Exists(),
			),
		},
	})
}

func TestAccApplicationPassword_deprecatedId(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_password", "test")
	r := ApplicationPasswordResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.deprecatedId(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_object_id").Exists(),
				check.That(data.ResourceName).Key("end_date").Exists(),
				check.That(data.ResourceName).Key("key_id").Exists(),
				check.That(data.ResourceName).Key("start_date").Exists(),
				check.That(data.ResourceName).Key("value").Exists(),
			),
		},
	})
}

func TestAccApplicationPassword_deprecatedId2(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_password", "test")
	r := ApplicationPasswordResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.deprecatedId2(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_object_id").Exists(),
				check.That(data.ResourceName).Key("end_date").Exists(),
				check.That(data.ResourceName).Key("key_id").Exists(),
				check.That(data.ResourceName).Key("start_date").Exists(),
				check.That(data.ResourceName).Key("value").Exists(),
			),
		},
	})
}

func TestAccApplicationPassword_with_ApplicationInlinePassword(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_password", "test")
	application := "azuread_application.test"

	r := ApplicationPasswordResource{}
	aR := ApplicationResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.passwordsCombined(data, true),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_object_id").Exists(),
				check.That(data.ResourceName).Key("end_date").Exists(),
				check.That(data.ResourceName).Key("key_id").Exists(),
				check.That(data.ResourceName).Key("start_date").Exists(),
				check.That(data.ResourceName).Key("value").Exists(),
				/* azuread_application */
				check.That(application).ExistsInAzure(aR),
				check.That(application).Key("password.#").HasValue("1"),
				check.That(application).Key("password.0.key_id").Exists(),
				check.That(application).Key("password.0.value").Exists(),
				check.That(application).Key("password.0.start_date").Exists(),
				check.That(application).Key("password.0.end_date").Exists(),
				check.That(application).Key("password.0.display_name").HasValue(fmt.Sprintf("acctest-appPassword-%s", data.RandomString)),
			),
		},
		{
			Config: r.passwordsCombined(data, false),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(application).ExistsInAzure(aR),
			),
		},
		{
			RefreshState: true,
			Check: acceptance.ComposeTestCheckFunc(
				check.That(application).ExistsInAzure(aR),
				check.That(application).Key("password.#").HasValue("0"),
			),
		},
	})
}

func (r ApplicationPasswordResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.Applications.ApplicationsClientBeta
	client.BaseClient.DisableRetries = true
	defer func() { client.BaseClient.DisableRetries = false }()

	id, err := parse.PasswordID(state.ID)
	if err != nil {
		return nil, fmt.Errorf("parsing Application Password ID: %v", err)
	}

	app, status, err := client.Get(ctx, id.ObjectId, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return nil, fmt.Errorf("Application with object ID %q does not exist", id.ObjectId)
		}
		return nil, fmt.Errorf("failed to retrieve Application with object ID %q: %+v", id.ObjectId, err)
	}

	if app.PasswordCredentials != nil {
		for _, cred := range *app.PasswordCredentials {
			if cred.KeyId != nil && *cred.KeyId == id.KeyId {
				return pointer.To(true), nil
			}
		}
	}

	return nil, fmt.Errorf("Password Credential %q was not found for Application %q", id.KeyId, id.ObjectId)
}

func (ApplicationPasswordResource) template(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  display_name = "acctestAppPassword-%[1]d"
}
`, data.RandomInteger)
}

func (r ApplicationPasswordResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_application_password" "test" {
  application_id = azuread_application.test.id
}
`, r.template(data))
}

func (r ApplicationPasswordResource) complete(data acceptance.TestData, startDate, endDate string) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_application_password" "test" {
  application_id = azuread_application.test.id
  display_name   = "terraform-%[2]s"
  start_date     = "%[3]s"
  end_date       = "%[4]s"
}
`, r.template(data), data.RandomString, startDate, endDate)
}

func (r ApplicationPasswordResource) relativeEndDate(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_application_password" "test" {
  application_id    = azuread_application.test.id
  display_name      = "terraform-%[2]s"
  end_date_relative = "8760h"
}
`, r.template(data), data.RandomString)
}

func (r ApplicationPasswordResource) deprecatedId(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_application_password" "test" {
  application_object_id = azuread_application.test.object_id
}
`, r.template(data))
}

func (r ApplicationPasswordResource) deprecatedId2(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_application_password" "test" {
  application_object_id = azuread_application.test.id
}
`, r.template(data))
}

func (r ApplicationPasswordResource) passwordsCombined(data acceptance.TestData, renderPassword bool) string {
	return fmt.Sprintf(`
provider "azuread" {}

data "azuread_client_config" "current" {}

resource "azuread_application" "test" {
  display_name = "acctest-appPassword-%[1]d"
  owners       = [data.azuread_client_config.current.object_id]

  %[3]s
}

resource "azuread_application_password" "test" {
  application_object_id = azuread_application.test.id
  display_name          = "acctest-application-password-%[2]s"
}


`, data.RandomInteger, data.RandomString, r.applicationPassword(data.RandomString, renderPassword))
}

func (r ApplicationPasswordResource) applicationPassword(randomString string, renderPassword bool) string {
	if renderPassword {
		return fmt.Sprintf(`
  password {
    display_name = "acctest-appPassword-%[1]s"
  }
`, randomString)
	}

	return ""

}
