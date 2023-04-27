package serviceprincipals_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
)

type ServicePrincipalDelegatedPermissionGrantResource struct{}

func TestAccServicePrincipalDelegatedPermissionGrant_allUsers(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_service_principal_delegated_permission_grant", "test")
	r := ServicePrincipalDelegatedPermissionGrantResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.allUsers(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccServicePrincipalDelegatedPermissionGrant_singleUser(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_service_principal_delegated_permission_grant", "test")
	r := ServicePrincipalDelegatedPermissionGrantResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.singleUser(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func (r ServicePrincipalDelegatedPermissionGrantResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.ServicePrincipals.DelegatedPermissionGrantsClient
	client.BaseClient.DisableRetries = true
	defer func() { client.BaseClient.DisableRetries = false }()

	if _, status, err := client.Get(ctx, state.ID, odata.Query{}); err != nil {
		if status == http.StatusNotFound {
			return nil, fmt.Errorf("Delegated Permission Grant with ID %q does not exist", state.ID)
		}
		return nil, fmt.Errorf("failed to retrieve Delegated Permission Grant with ID %q: %+v", state.ID, err)
	}

	return utils.Bool(true), nil
}

func (r ServicePrincipalDelegatedPermissionGrantResource) template(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

data "azuread_application_published_app_ids" "well_known" {}

resource "azuread_service_principal" "msgraph" {
  application_id = data.azuread_application_published_app_ids.well_known.result.MicrosoftGraph
  use_existing   = true
}

resource "azuread_application" "test" {
  display_name = "acctest-APP-%[1]d"

  required_resource_access {
    resource_app_id = data.azuread_application_published_app_ids.well_known.result.MicrosoftGraph

    resource_access {
      id   = azuread_service_principal.msgraph.oauth2_permission_scope_ids["openid"]
      type = "Scope"
    }

    resource_access {
      id   = azuread_service_principal.msgraph.oauth2_permission_scope_ids["User.Read"]
      type = "Scope"
    }
  }
}

resource "azuread_service_principal" "test" {
  application_id = azuread_application.test.application_id
}
`, data.RandomInteger)
}

func (r ServicePrincipalDelegatedPermissionGrantResource) allUsers(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_service_principal_delegated_permission_grant" "test" {
  service_principal_object_id          = azuread_service_principal.test.object_id
  resource_service_principal_object_id = azuread_service_principal.msgraph.object_id
  claim_values                         = ["openid", "User.Read.All"]
}
`, r.template(data))
}

func (r ServicePrincipalDelegatedPermissionGrantResource) singleUser(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

data "azuread_domains" "test" {
  only_initial = true
}

resource "azuread_user" "test" {
  user_principal_name = "acctestUser'%[2]d@${data.azuread_domains.test.domains.0.domain_name}"
  display_name        = "acctestUser-%[2]d"
  password            = "%[3]s"
}

resource "azuread_service_principal_delegated_permission_grant" "test" {
  service_principal_object_id          = azuread_service_principal.test.object_id
  resource_service_principal_object_id = azuread_service_principal.msgraph.object_id
  claim_values                         = ["openid", "User.Read.All"]
}
`, r.template(data), data.RandomInteger, data.RandomPassword)
}
