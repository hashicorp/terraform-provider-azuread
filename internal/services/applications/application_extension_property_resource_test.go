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
	"github.com/hashicorp/terraform-provider-azuread/internal/services/applications/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
)

type ApplicationExtensionPropertyResource struct{}

func TestAccApplicationExtensionProperty_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_extension_property", "test")
	r := ApplicationExtensionPropertyResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.complete(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_id").Exists(),
				check.That(data.ResourceName).Key("data_type").Exists(),
				//check.That(data.ResourceName).Key("target_objects").Exists(),
			),
		},
		data.ImportStep(),
	})
}

func (r ApplicationExtensionPropertyResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.Applications.ApplicationsClient
	client.BaseClient.DisableRetries = true
	extId, err := parse.ApplicationExtensionPropertyID(state.ID)
	if err != nil {
		return nil, err
	}

	appExts, status, err := client.ListExtensions(ctx, extId.ApplicationId, odata.Query{
		Filter: "id eq '" + extId.ExtensionPropertyId + "'",
	})
	if err != nil {
		if status == http.StatusNotFound {
			return nil, fmt.Errorf("Extension property with ID %q does not exist in Application ID %q", extId.ExtensionPropertyId, extId.ApplicationId)
		}
		return nil, fmt.Errorf("failed to retrieve Extension property ID %q Application ID %q: %+v", extId.ExtensionPropertyId, extId.ApplicationId, err)
	}
	appExt := (*appExts)[0]
	return utils.Bool(appExt.Id != nil && *appExt.Id == extId.ExtensionPropertyId), nil
}

func (ApplicationExtensionPropertyResource) complete(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

resource "azuread_application" "testapp" {
	display_name = "acctest-APP-%[1]d"
}

resource "azuread_application_extension_property" "test" {
  application_id  = azuread_application.testapp.id
  name            = "acctest_APP_complete_%[1]d"
  data_type       = "String"
  target_objects = [ "User" ]
}
`, data.RandomInteger)
}
