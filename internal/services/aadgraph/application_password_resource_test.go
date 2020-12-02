package aadgraph_test

import (
	"context"
	"fmt"
	"github.com/terraform-providers/terraform-provider-azuread/internal/acceptance/check"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/terraform-providers/terraform-provider-azuread/internal/acceptance"
	"github.com/terraform-providers/terraform-provider-azuread/internal/clients"
	"github.com/terraform-providers/terraform-provider-azuread/internal/services/aadgraph/graph"
	"github.com/terraform-providers/terraform-provider-azuread/internal/utils"
)

type ApplicationPasswordResource struct{}

func TestAccApplicationPassword_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_password", "test")
	data.AdditionalData["end_date"] = time.Now().AddDate(0, 5, 27).UTC().Format(time.RFC3339)
	r := ApplicationPasswordResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("key_id").Exists(),
			),
		},
		data.ImportStep("value"),
	})
}

func TestAccApplicationPassword_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_password", "test")
	data.AdditionalData["end_date"] = time.Now().AddDate(0, 5, 27).UTC().Format(time.RFC3339)
	r := ApplicationPasswordResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.complete(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("key_id").HasValue(data.RandomID),
			),
		},
		data.ImportStep("value"),
	})
}

func TestAccApplicationPassword_relativeEndDate(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_password", "test")
	r := ApplicationPasswordResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.relativeEndDate(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("key_id").Exists(),
				check.That(data.ResourceName).Key("end_date").Exists(),
			),
		},
		data.ImportStep("end_date_relative", "value"),
	})
}

func TestAccApplicationPassword_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_password", "test")
	data.AdditionalData["end_date"] = time.Now().AddDate(0, 5, 27).UTC().Format(time.RFC3339)
	r := ApplicationPasswordResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("key_id").Exists(),
			),
		},
		data.RequiresImportErrorStep(r.requiresImport(data)),
	})
}

func (a ApplicationPasswordResource) Exists(ctx context.Context, clients *clients.AadClient, state *terraform.InstanceState) (*bool, error) {
	id, err := graph.ParsePasswordId(state.ID)
	if err != nil {
		return nil, fmt.Errorf("parsing Application Password ID: %v", err)
	}

	resp, err := clients.AadGraph.ApplicationsClient.Get(ctx, id.ObjectId)
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			return nil, fmt.Errorf("Application with object ID %q does not exist", id.ObjectId)
		}
		return nil, fmt.Errorf("failed to retrieve Application with object ID %q: %+v", id.ObjectId, err)
	}

	credentials, err := clients.AadGraph.ApplicationsClient.ListPasswordCredentials(ctx, id.ObjectId)
	if err != nil {
		return nil, fmt.Errorf("listing Password Credentials for Application %q: %+v", id.ObjectId, err)
	}

	cred := graph.PasswordCredentialResultFindByKeyId(credentials, id.KeyId)
	if cred != nil {
		return utils.Bool(true), nil
	}

	return nil, fmt.Errorf("Password Credential %q was not found for Application %q", id.KeyId, id.ObjectId)
}

func (ApplicationPasswordResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_application_password" "test" {
  application_object_id = azuread_application.test.id
  value                 = "%[2]s"
  end_date              = "%[3]s"
}
`, ApplicationResource{}.basic(data), data.RandomPassword, data.AdditionalData["end_date"])
}

func (ApplicationPasswordResource) complete(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_application_password" "test" {
  application_object_id = azuread_application.test.id
  description           = "terraform"
  key_id                = "%[2]s"
  value                 = "%[3]s"
  end_date              = "%[4]s"
}
`, ApplicationResource{}.basic(data), data.RandomID, data.RandomPassword, data.AdditionalData["end_date"])
}

func (ApplicationPasswordResource) relativeEndDate(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_application_password" "test" {
  application_object_id = azuread_application.test.id
  value                 = "%[2]s"
  end_date_relative     = "8760h"
}
`, ApplicationResource{}.basic(data), data.RandomPassword)
}

func (ApplicationPasswordResource) requiresImport(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_application_password" "import" {
  application_object_id = azuread_application_password.test.application_object_id
  key_id                = azuread_application_password.test.key_id
  value                 = azuread_application_password.test.value
  end_date              = azuread_application_password.test.end_date
}
`, ApplicationPasswordResource{}.basic(data))
}
