// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applications_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/federatedidentitycredential"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
)

type ApplicationFlexibleFederatedIdentityCredentialResource struct{}

func TestAccApplicationFlexibleFederatedIdentityCredential_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_flexible_federated_identity_credential", "test")
	r := ApplicationFlexibleFederatedIdentityCredentialResource{}

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

func TestAccApplicationFlexibleFederatedIdentityCredential_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_flexible_federated_identity_credential", "test")
	r := ApplicationFlexibleFederatedIdentityCredentialResource{}

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

func TestAccApplicationFlexibleFederatedIdentityCredential_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_flexible_federated_identity_credential", "test")
	r := ApplicationFlexibleFederatedIdentityCredentialResource{}

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

func (r ApplicationFlexibleFederatedIdentityCredentialResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.Applications.ApplicationFlexibleFederatedIdentityCredential

	id, err := beta.ParseApplicationIdFederatedIdentityCredentialID(state.ID)
	if err != nil {
		return nil, fmt.Errorf("parsing Application Federated Identity Credential ID: %v", err)
	}

	credentialId := beta.NewApplicationIdFederatedIdentityCredentialID(id.ApplicationId, id.FederatedIdentityCredentialId)

	resp, err := client.GetFederatedIdentityCredential(ctx, credentialId, federatedidentitycredential.DefaultGetFederatedIdentityCredentialOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return pointer.To(false), nil
		}
		return nil, fmt.Errorf("failed to retrieve %s: %+v", credentialId, err)
	}

	return pointer.To(resp.Model != nil), nil
}

func (ApplicationFlexibleFederatedIdentityCredentialResource) template(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  display_name = "acctestFederatedIdentityCredential-%[1]d"
}
`, data.RandomInteger)
}

func (r ApplicationFlexibleFederatedIdentityCredentialResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_application_flexible_federated_identity_credential" "test" {
  application_id             = azuread_application.test.id
  display_name   			 = "hashitown.example.com-%[2]s"
  claims_matching_expression = "claims['sub'] matches 'repo:contoso/contoso-repo:ref:refs/heads/*' and claims['job_workflow_ref'] matches 'contoso/contoso-prod/.github/workflows/*.yml@refs/heads/main'"
  audience      			 = "api://HashiTownLikesAzureAD"
  issuer         			 = "https://token.actions.githubusercontent.com"
}
`, r.template(data), data.RandomString)
}

func (r ApplicationFlexibleFederatedIdentityCredentialResource) complete(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_application_flexible_federated_identity_credential" "test" {
  application_id             = azuread_application.test.id
  display_name               = "hashitown.example.com-%[2]s"
  claims_matching_expression = "claims['sub'] matches 'repo:contoso/contoso-repo:ref:refs/heads/*' and claims['job_workflow_ref'] matches 'contoso/contoso-prod/.github/workflows/update.yml@refs/heads/main'"
  description                = "Funtime tokens for HashiTown"
  audience                   = "api://HashiTownLikesAzureAD"
  issuer                     = "https://token.actions.githubusercontent.com"
}
`, r.template(data), data.RandomString, data.UUID())
}
