// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applications_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/federatedidentitycredential"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/applications/parse"
)

type ApplicationFederatedIdentityCredentialResource struct{}

func TestAccApplicationFederatedIdentityCredential_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_federated_identity_credential", "test")
	r := ApplicationFederatedIdentityCredentialResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("credential_id").Exists(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApplicationFederatedIdentityCredential_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_federated_identity_credential", "test")
	r := ApplicationFederatedIdentityCredentialResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.complete(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("credential_id").Exists(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApplicationFederatedIdentityCredential_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_federated_identity_credential", "test")
	r := ApplicationFederatedIdentityCredentialResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("credential_id").Exists(),
			),
		},
		data.ImportStep(),
		{
			Config: r.complete(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("credential_id").Exists(),
			),
		},
		data.ImportStep(),
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("credential_id").Exists(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApplicationFederatedIdentityCredential_deprecatedId(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_federated_identity_credential", "test")
	r := ApplicationFederatedIdentityCredentialResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.deprecatedId(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_object_id").Exists(),
				check.That(data.ResourceName).Key("credential_id").Exists(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccApplicationFederatedIdentityCredential_deprecatedId2(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_federated_identity_credential", "test")
	r := ApplicationFederatedIdentityCredentialResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.deprecatedId2(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_object_id").Exists(),
				check.That(data.ResourceName).Key("credential_id").Exists(),
			),
		},
		data.ImportStep("application_object_id"),
	})
}

func (r ApplicationFederatedIdentityCredentialResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.Applications.ApplicationFederatedIdentityCredential

	id, err := parse.FederatedIdentityCredentialID(state.ID)
	if err != nil {
		return nil, fmt.Errorf("parsing Application Federated Identity Credential ID: %v", err)
	}

	credentialId := stable.NewApplicationIdFederatedIdentityCredentialID(id.ObjectId, id.KeyId)

	resp, err := client.GetFederatedIdentityCredential(ctx, credentialId, federatedidentitycredential.DefaultGetFederatedIdentityCredentialOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return pointer.To(false), nil
		}
		return nil, fmt.Errorf("failed to retrieve %s: %+v", credentialId, err)
	}

	return pointer.To(resp.Model != nil), nil
}

func (ApplicationFederatedIdentityCredentialResource) template(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  display_name = "acctestFederatedIdentityCredential-%[1]d"
}
`, data.RandomInteger)
}

func (r ApplicationFederatedIdentityCredentialResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_application_federated_identity_credential" "test" {
  application_id = azuread_application.test.id
  display_name   = "hashitown-%[2]s"
  audiences      = ["api://HashiTownLikesAzureAD"]
  issuer         = "https://tokens.hashitown.net"
  subject        = "%[3]s"
}
`, r.template(data), data.RandomString, data.RandomID)
}

func (r ApplicationFederatedIdentityCredentialResource) complete(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_application_federated_identity_credential" "test" {
  application_id = azuread_application.test.id
  display_name   = "hashitown-%[2]s"
  description    = "Funtime tokens for HashiTown"
  audiences      = ["api://HashiTownLikesAzureAD"]
  issuer         = "https://vending.hashitown.net"
  subject        = "%[3]s"
}
`, r.template(data), data.RandomString, data.UUID())
}

func (r ApplicationFederatedIdentityCredentialResource) deprecatedId(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_application_federated_identity_credential" "test" {
  application_object_id = azuread_application.test.object_id
  display_name          = "hashitown-%[2]s"
  audiences             = ["api://HashiTownLikesAzureAD"]
  issuer                = "https://tokens.hashitown.net"
  subject               = "%[3]s"
}
`, r.template(data), data.RandomString, data.RandomID)
}

func (r ApplicationFederatedIdentityCredentialResource) deprecatedId2(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_application_federated_identity_credential" "test" {
  application_object_id = azuread_application.test.id
  display_name          = "hashitown-%[2]s"
  audiences             = ["api://HashiTownLikesAzureAD"]
  issuer                = "https://tokens.hashitown.net"
  subject               = "%[3]s"
}
`, r.template(data), data.RandomString, data.RandomID)
}
