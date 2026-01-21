// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package synchronization_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/stable/synchronizationsecret"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
)

type SynchronizationSecretResource struct{}

func TestAccSynchronizationSecret_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_synchronization_secret", "test")
	r := SynchronizationSecretResource{}

	data.ResourceTestIgnoreDangling(t, r, []acceptance.TestStep{
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
	client := clients.Synchronization.SynchronizationSecretClient

	id, err := stable.ParseServicePrincipalID(state.ID)
	if err != nil {
		return nil, err
	}

	resp, err := client.ListSynchronizationSecrets(ctx, *id, synchronizationsecret.DefaultListSynchronizationSecretsOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return pointer.To(false), nil
		}
		return nil, fmt.Errorf("retrieving synchronization secrets for %s", id)
	}

	if resp.Model == nil {
		return pointer.To(false), nil
	}

	return pointer.To(true), nil
}

func (r SynchronizationSecretResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azuread" {}

data "azuread_client_config" "test" {}

data "azuread_application_template" "test" {
  display_name = "Azure Databricks SCIM Provisioning Connector"
}

resource "azuread_application_from_template" "test" {
  display_name = "acctestSynchronizationJob-%[1]d"
  template_id  = data.azuread_application_template.test.template_id
}

data "azuread_service_principal" "test" {
  object_id = azuread_application_from_template.test.service_principal_object_id
}

resource "azuread_synchronization_secret" "test" {
  service_principal_id = data.azuread_service_principal.test.id

  credential {
    key   = "BaseAddress"
    value = "https://test-address.azuredatabricks.net"
  }
  credential {
    key   = "SecretToken"
    value = "password-value"
  }
}
`, data.RandomInteger)
}
