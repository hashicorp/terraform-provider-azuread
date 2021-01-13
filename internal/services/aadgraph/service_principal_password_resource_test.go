package aadgraph_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/terraform-providers/terraform-provider-azuread/internal/acceptance"
	"github.com/terraform-providers/terraform-provider-azuread/internal/acceptance/check"
	"github.com/terraform-providers/terraform-provider-azuread/internal/clients"
	"github.com/terraform-providers/terraform-provider-azuread/internal/services/aadgraph/graph"
	"github.com/terraform-providers/terraform-provider-azuread/internal/utils"
)

type ServicePrincipalPasswordResource struct{}

func TestAccServicePrincipalPassword_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_service_principal_password", "test")
	endDate := time.Now().AddDate(0, 5, 27).UTC().Format(time.RFC3339)
	r := ServicePrincipalPasswordResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data, endDate),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("key_id").Exists(),
			),
		},
		data.ImportStep("value"),
	})
}

func TestAccServicePrincipalPassword_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_service_principal_password", "test")
	startDate := time.Now().AddDate(0, 0, 7).UTC().Format(time.RFC3339)
	endDate := time.Now().AddDate(0, 5, 27).UTC().Format(time.RFC3339)
	r := ServicePrincipalPasswordResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.complete(data, startDate, endDate),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("key_id").HasValue(data.RandomID),
			),
		},
		data.ImportStep("value"),
	})
}

func TestAccServicePrincipalPassword_relativeEndDate(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_service_principal_password", "test")
	r := ServicePrincipalPasswordResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.relativeEndDate(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("key_id").Exists(),
			),
		},
		data.ImportStep("end_date_relative", "value"),
	})
}

func TestAccServicePrincipalPassword_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_service_principal_password", "test")
	endDate := time.Now().AddDate(0, 5, 27).UTC().Format(time.RFC3339)
	r := ServicePrincipalPasswordResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data, endDate),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.RequiresImportErrorStep(r.requiresImport(data, endDate)),
	})
}

func (r ServicePrincipalPasswordResource) Exists(ctx context.Context, clients *clients.AadClient, state *terraform.InstanceState) (*bool, error) {
	id, err := graph.ParsePasswordId(state.ID)
	if err != nil {
		return nil, fmt.Errorf("Service Principal Password ID: %v", err)
	}

	resp, err := clients.AadGraph.ServicePrincipalsClient.Get(ctx, id.ObjectId)
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			return nil, fmt.Errorf("Service Principal with object ID %q does not exist", id.ObjectId)
		}
		return nil, fmt.Errorf("failed to retrieve Service Principal with object ID %q: %+v", id.ObjectId, err)
	}

	credentials, err := clients.AadGraph.ServicePrincipalsClient.ListPasswordCredentials(ctx, id.ObjectId)
	if err != nil {
		return nil, fmt.Errorf("listing Password Credentials for Service Principal %q: %+v", id.ObjectId, err)
	}

	cred := graph.PasswordCredentialResultFindByKeyId(credentials, id.KeyId)
	if cred != nil {
		return utils.Bool(true), nil
	}

	return nil, fmt.Errorf("Password Credential %q was not found for Service Principal %q", id.KeyId, id.ObjectId)
}

func (ServicePrincipalPasswordResource) template(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name = "acctestServicePrincipal-%[1]d"
}

resource "azuread_service_principal" "test" {
  application_id = azuread_application.test.application_id
}
`, data.RandomInteger)
}

func (r ServicePrincipalPasswordResource) basic(data acceptance.TestData, endDate string) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_service_principal_password" "test" {
  service_principal_id = azuread_service_principal.test.id
  value                = "%[2]s"
  end_date             = "%[3]s"
}
`, r.template(data), data.RandomPassword, endDate)
}

func (r ServicePrincipalPasswordResource) complete(data acceptance.TestData, startDate, endDate string) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_service_principal_password" "test" {
  service_principal_id = azuread_service_principal.test.id
  description          = "terraform"
  key_id               = "%[2]s"
  start_date           = "%[3]s"
  end_date             = "%[4]s"
  value                = "%[5]s"
}
`, r.template(data), data.RandomID, startDate, endDate, data.RandomPassword)
}

func (r ServicePrincipalPasswordResource) relativeEndDate(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_service_principal_password" "test" {
  service_principal_id = azuread_service_principal.test.id
  value                = "%[2]s"
  end_date_relative    = "8760h"
}
`, r.template(data), data.RandomPassword)
}

func (r ServicePrincipalPasswordResource) requiresImport(data acceptance.TestData, endDate string) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_service_principal_password" "import" {
  key_id               = azuread_service_principal_password.test.key_id
  service_principal_id = azuread_service_principal_password.test.service_principal_id
  value                = azuread_service_principal_password.test.value
  end_date             = azuread_service_principal_password.test.end_date
}
`, r.basic(data, endDate))
}
