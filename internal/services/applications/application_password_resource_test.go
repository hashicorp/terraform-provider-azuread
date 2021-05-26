package applications_test

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/aadgraph"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/applications/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
)

type ApplicationPasswordResource struct{}

func TestAccApplicationPassword_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_password", "test")
	r := ApplicationPasswordResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("end_date").Exists(),
				check.That(data.ResourceName).Key("key_id").Exists(),
				check.That(data.ResourceName).Key("start_date").Exists(),
				check.That(data.ResourceName).Key("value").Exists(),
			),
		},
	})
}

func TestAccApplicationPassword_updateDeprecated(t *testing.T) {
	// TODO: remove this test in v2.0
	if v := os.Getenv("AAD_USE_MICROSOFT_GRAPH"); v != "" {
		t.Skipf("Test skipped when using MS Graph")
	}

	data := acceptance.BuildTestData(t, "azuread_application_password", "test")
	startDate := time.Now().AddDate(0, 0, 7).UTC().Format(time.RFC3339)
	endDate := time.Now().AddDate(0, 5, 27).UTC().Format(time.RFC3339)
	r := ApplicationPasswordResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.completeAadGraph(data, startDate, endDate),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("end_date").Exists(),
				check.That(data.ResourceName).Key("key_id").Exists(),
				check.That(data.ResourceName).Key("start_date").Exists(),
				check.That(data.ResourceName).Key("value").Exists(),
			),
		},
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("end_date").Exists(),
				check.That(data.ResourceName).Key("key_id").Exists(),
				check.That(data.ResourceName).Key("start_date").Exists(),
				check.That(data.ResourceName).Key("value").Exists(),
			),
		},
	})
}

func TestAccApplicationPassword_basicAadGraph(t *testing.T) {
	// TODO: remove this test in v2.0
	if v := os.Getenv("AAD_USE_MICROSOFT_GRAPH"); v != "" {
		t.Skipf("Test skipped when using MS Graph")
	}

	data := acceptance.BuildTestData(t, "azuread_application_password", "test")
	endDate := time.Now().AddDate(0, 5, 27).UTC().Format(time.RFC3339)
	r := ApplicationPasswordResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basicAadGraph(data, endDate),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("key_id").Exists(),
				check.That(data.ResourceName).Key("end_date").HasValue(endDate),
				check.That(data.ResourceName).Key("value").Exists(),
			),
		},
	})
}

func TestAccApplicationPassword_completeAadGraph(t *testing.T) {
	// TODO: remove this test in v2.0
	if v := os.Getenv("AAD_USE_MICROSOFT_GRAPH"); v != "" {
		t.Skipf("Test skipped when using MS Graph")
	}

	data := acceptance.BuildTestData(t, "azuread_application_password", "test")
	startDate := time.Now().AddDate(0, 0, 7).UTC().Format(time.RFC3339)
	endDate := time.Now().AddDate(0, 5, 27).UTC().Format(time.RFC3339)
	r := ApplicationPasswordResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.completeAadGraph(data, startDate, endDate),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("key_id").HasValue(data.RandomID),
				check.That(data.ResourceName).Key("start_date").HasValue(startDate),
				check.That(data.ResourceName).Key("end_date").HasValue(endDate),
				check.That(data.ResourceName).Key("value").Exists(),
			),
		},
	})
}

func TestAccApplicationPassword_relativeEndDateAadGraph(t *testing.T) {
	// TODO: remove this test in v2.0
	if v := os.Getenv("AAD_USE_MICROSOFT_GRAPH"); v != "" {
		t.Skipf("Test skipped when using MS Graph")
	}

	data := acceptance.BuildTestData(t, "azuread_application_password", "test")
	r := ApplicationPasswordResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.relativeEndDateAadGraph(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("key_id").Exists(),
				check.That(data.ResourceName).Key("end_date").Exists(),
				check.That(data.ResourceName).Key("end_date_relative").HasValue("8760h"),
				check.That(data.ResourceName).Key("value").Exists(),
			),
		},
	})
}

func (r ApplicationPasswordResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	id, err := parse.PasswordID(state.ID)
	if err != nil {
		return nil, fmt.Errorf("parsing Application Password ID: %v", err)
	}

	if clients.EnableMsGraphBeta {
		app, status, err := clients.Applications.MsClient.Get(ctx, id.ObjectId)
		if err != nil {
			if status == http.StatusNotFound {
				return nil, fmt.Errorf("Application with object ID %q does not exist", id.ObjectId)
			}
			return nil, fmt.Errorf("failed to retrieve Application with object ID %q: %+v", id.ObjectId, err)
		}

		if app.PasswordCredentials != nil {
			for _, cred := range *app.PasswordCredentials {
				if cred.KeyId != nil && *cred.KeyId == id.KeyId {
					return utils.Bool(true), nil
				}
			}
		}
	} else {
		resp, err := clients.Applications.AadClient.Get(ctx, id.ObjectId)
		if err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
				return nil, fmt.Errorf("Application with object ID %q does not exist", id.ObjectId)
			}
			return nil, fmt.Errorf("failed to retrieve Application with object ID %q: %+v", id.ObjectId, err)
		}

		credentials, err := clients.Applications.AadClient.ListPasswordCredentials(ctx, id.ObjectId)
		if err != nil {
			return nil, fmt.Errorf("listing Password Credentials for Application %q: %+v", id.ObjectId, err)
		}

		cred := aadgraph.PasswordCredentialResultFindByKeyId(credentials, id.KeyId)
		if cred != nil {
			return utils.Bool(true), nil
		}
	}

	return nil, fmt.Errorf("Password Credential %q was not found for Application %q", id.KeyId, id.ObjectId)
}

func (ApplicationPasswordResource) template(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name = "acctestAppPassword-%[1]d"
}
`, data.RandomInteger)
}

func (r ApplicationPasswordResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_application_password" "test" {
  application_object_id = azuread_application.test.object_id
}
`, r.template(data))
}

func (r ApplicationPasswordResource) basicAadGraph(data acceptance.TestData, endDate string) string {
	// TODO: remove this config in v2.0
	return fmt.Sprintf(`
%[1]s

resource "azuread_application_password" "test" {
  application_object_id = azuread_application.test.id
  value                 = "%[2]s"
  end_date              = "%[3]s"
}
`, r.template(data), data.RandomPassword, endDate)
}

func (r ApplicationPasswordResource) completeAadGraph(data acceptance.TestData, startDate, endDate string) string {
	// TODO: remove this config in v2.0
	return fmt.Sprintf(`
%[1]s

resource "azuread_application_password" "test" {
  application_object_id = azuread_application.test.id
  description           = "terraform"
  key_id                = "%[2]s"
  start_date            = "%[3]s"
  end_date              = "%[4]s"
  value                 = "%[5]s"
}
`, r.template(data), data.RandomID, startDate, endDate, data.RandomPassword)
}

func (r ApplicationPasswordResource) relativeEndDateAadGraph(data acceptance.TestData) string {
	// TODO: remove this config in v2.0
	return fmt.Sprintf(`
%[1]s

resource "azuread_application_password" "test" {
  application_object_id = azuread_application.test.id
  value                 = "%[2]s"
  end_date_relative     = "8760h"
}
`, r.template(data), data.RandomPassword)
}
