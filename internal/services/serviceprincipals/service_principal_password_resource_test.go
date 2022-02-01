package serviceprincipals_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/manicminer/hamilton/odata"

	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/serviceprincipals/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
)

type ServicePrincipalPasswordResource struct{}

func TestAccServicePrincipalPassword_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_service_principal_password", "test")
	r := ServicePrincipalPasswordResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("key_id").Exists(),
				check.That(data.ResourceName).Key("start_date").Exists(),
				check.That(data.ResourceName).Key("end_date").Exists(),
				check.That(data.ResourceName).Key("value").Exists(),
			),
		},
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
				check.That(data.ResourceName).Key("key_id").Exists(),
				check.That(data.ResourceName).Key("start_date").Exists(),
				check.That(data.ResourceName).Key("end_date").Exists(),
				check.That(data.ResourceName).Key("value").Exists(),
			),
		},
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
				check.That(data.ResourceName).Key("end_date").Exists(),
				check.That(data.ResourceName).Key("end_date_relative").HasValue("8760h"),
				check.That(data.ResourceName).Key("key_id").Exists(),
				check.That(data.ResourceName).Key("start_date").Exists(),
				check.That(data.ResourceName).Key("value").Exists(),
			),
		},
	})
}

func (r ServicePrincipalPasswordResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.ServicePrincipals.ServicePrincipalsClient
	client.BaseClient.DisableRetries = true

	id, err := parse.PasswordID(state.ID)
	if err != nil {
		return nil, fmt.Errorf("parsing Service Principal Password ID: %v", err)
	}

	servicePrincipal, status, err := client.Get(ctx, id.ObjectId, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return nil, fmt.Errorf("Service Principal with object ID %q does not exist", id.ObjectId)
		}
		return nil, fmt.Errorf("failed to retrieve Service Principal with object ID %q: %+v", id.ObjectId, err)
	}

	if servicePrincipal.PasswordCredentials != nil {
		for _, cred := range *servicePrincipal.PasswordCredentials {
			if cred.KeyId != nil && *cred.KeyId == id.KeyId {
				return utils.Bool(true), nil
			}
		}
	}

	return nil, fmt.Errorf("Password Credential %q was not found for Service Principal %q", id.KeyId, id.ObjectId)
}

func (r ServicePrincipalPasswordResource) template(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  display_name = "acctestServicePrincipal-%[1]d"
}

resource "azuread_service_principal" "test" {
  application_id = azuread_application.test.application_id
}
`, data.RandomInteger)
}

func (r ServicePrincipalPasswordResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_service_principal_password" "test" {
  service_principal_id = azuread_service_principal.test.object_id
}
`, r.template(data))
}

func (r ServicePrincipalPasswordResource) complete(data acceptance.TestData, startDate, endDate string) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_service_principal_password" "test" {
  service_principal_id = azuread_service_principal.test.object_id
  display_name         = "terraform-%[2]s"
  start_date           = "%[3]s"
  end_date             = "%[4]s"
}
`, r.template(data), data.RandomString, startDate, endDate)
}

func (r ServicePrincipalPasswordResource) relativeEndDate(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_service_principal_password" "test" {
  service_principal_id = azuread_service_principal.test.object_id
  display_name         = "terraform-%[2]s"
  end_date_relative    = "8760h"
}
`, r.template(data), data.RandomString)
}
