package applications_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/manicminer/hamilton/odata"

	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
)

type ApplicationExtensionPropertyResource struct{}

func TestAccApplicationExtensionProperty_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_extension_property", "test")
	r := ApplicationExtensionPropertyResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
				check.That(data.ResourceName).Key("object_id").Exists(),
				check.That(data.ResourceName).Key("display_name").HasValue(fmt.Sprintf("acctest-APP-%d", data.RandomInteger)),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApplicationExtensionProperty_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_extension_property", "test")
	r := ApplicationResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.complete(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
				check.That(data.ResourceName).Key("object_id").Exists(),
			),
		},
		data.ImportStep(),
	})
}

func (r ApplicationExtensionPropertyResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.Applications.ApplicationsClient
	client.BaseClient.DisableRetries = true
	app, status, err := client.Get(ctx, state.ID, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return nil, fmt.Errorf("Application with object ID %q does not exist", state.ID)
		}
		return nil, fmt.Errorf("failed to retrieve Application with object ID %q: %+v", state.ID, err)
	}
	return utils.Bool(app.ID != nil && *app.ID == state.ID), nil
}

func (ApplicationExtensionPropertyResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application_extension_property" "test" {
  name = "my-ext-prop-%[1]d"
  data_type = "string"
}
`, data.RandomInteger)
}

func (ApplicationExtensionPropertyResource) complete(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

data "azuread_application" "b2capptest" {
  name = "b2c-extensions-app. Do not modify. Used by AADB2C for storing user data."
}

resource "azuread_applmication_extension_property" "test" {
  application_id  = data.azuread_application.b2capptest.id
  name            = "acctest-APP-complete-%[1]d"
  data_type       = "string"
  app_display_name= "AccTest APP Complete %[1]d"
  target_objects = [ "Users" ]
  is_synced_from_on_premises = false
}
`, data.RandomInteger)
}
