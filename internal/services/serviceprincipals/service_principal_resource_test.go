package serviceprincipals_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
)

type ServicePrincipalResource struct{}

func TestAccServicePrincipal_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_service_principal", "test")
	r := ServicePrincipalResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccServicePrincipal_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_service_principal", "test")
	r := ServicePrincipalResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.complete(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccServicePrincipal_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_service_principal", "test")
	r := ServicePrincipalResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.complete(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func (r ServicePrincipalResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	var id *string

	if clients.EnableMsGraphBeta {
		app, status, err := clients.ServicePrincipals.MsClient.Get(ctx, state.ID)
		if err != nil {
			if status == http.StatusNotFound {
				return nil, fmt.Errorf("Service Principal with object ID %q does not exist", state.ID)
			}
			return nil, fmt.Errorf("failed to retrieve Service Principal with object ID %q: %+v", state.ID, err)
		}
		id = app.ID
	} else {
		resp, err := clients.ServicePrincipals.AadClient.Get(ctx, state.ID)

		if err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
				return nil, fmt.Errorf("Service Principal with object ID %q does not exist", state.ID)
			}
			return nil, fmt.Errorf("failed to retrieve Service Principal with object ID %q: %+v", state.ID, err)
		}
		id = resp.ObjectID
	}

	return utils.Bool(id != nil && *id == state.ID), nil
}

func (ServicePrincipalResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name = "acctestServicePrincipal-%[1]d"
}

resource "azuread_service_principal" "test" {
  application_id = azuread_application.test.application_id
}
`, data.RandomInteger)
}

func (ServicePrincipalResource) complete(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name = "acctestServicePrincipal-%[1]d"
}

resource "azuread_service_principal" "test" {
  application_id               = azuread_application.test.application_id
  app_role_assignment_required = true

  tags = ["test", "multiple", "CapitalS"]
}
`, data.RandomInteger)
}
