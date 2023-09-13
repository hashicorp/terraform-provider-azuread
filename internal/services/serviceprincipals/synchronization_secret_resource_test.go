// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serviceprincipals_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/serviceprincipals/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
)

type SynchronizationSecretResource struct{}

func TestAccSynchronizationSecret_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_synchronization_secret", "test")
	r := SynchronizationSecretResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("credential.#").HasValue("2"),
				check.That(data.ResourceName).Key("credential.0.key").HasValue("BaseAddress"),
				check.That(data.ResourceName).Key("credential.0.value").HasValue("https://test-address.azuredatabricks.net"),
				check.That(data.ResourceName).Key("credential.1.key").HasValue("SecretToken"),
				check.That(data.ResourceName).Key("credential.1.value").HasValue("password-value"),
			),
		},
	})
}

func (r SynchronizationSecretResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.ServicePrincipals.SynchronizationJobClient
	client.BaseClient.DisableRetries = true
	defer func() { client.BaseClient.DisableRetries = false }()

	id, err := parse.SynchronizationSecretID(state.ID)
	if err != nil {
		return nil, fmt.Errorf("Parsing synchronization secret from service principal %v", err)
	}

	_, status, err := client.GetSecrets(ctx, id.ServicePrincipalId)
	if err != nil {
		if status == http.StatusNotFound {
			return nil, fmt.Errorf("Synchronization secrets for service principal %q was not found ", id.ServicePrincipalId)
		}
		return nil, fmt.Errorf("Retrieving synchronization secrets for service principal %q", id.ServicePrincipalId)
	}
	return utils.Bool(true), nil
}

func (SynchronizationSecretResource) template(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

data "azuread_client_config" "test" {}

data "azuread_application_template" "test" {
  display_name = "Azure Databricks SCIM Provisioning Connector"
}

resource "azuread_application" "test" {
  display_name = "acctestSynchronizationJob-%[1]d"
  owners       = [data.azuread_client_config.test.object_id]
  template_id  = data.azuread_application_template.test.template_id
}

resource "azuread_service_principal" "test" {
  application_id = azuread_application.test.application_id
  owners         = [data.azuread_client_config.test.object_id]
  use_existing   = true
}
`, data.RandomInteger)
}

func (r SynchronizationSecretResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_synchronization_secret" "test" {
  service_principal_id = azuread_service_principal.test.id

  credential {
    key   = "BaseAddress"
    value = "https://test-address.azuredatabricks.net"
  }
  credential {
    key   = "SecretToken"
    value = "password-value"
  }
}
`, r.template(data))
}
