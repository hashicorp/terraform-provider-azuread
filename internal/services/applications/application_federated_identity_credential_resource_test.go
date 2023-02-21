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

type ApplicationFederatedIdentityCredentialResource struct{}

func TestAccApplicationFederatedIdentityCredential_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_federated_identity_credential", "test")
	r := ApplicationFederatedIdentityCredentialResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
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

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.complete(data),
			Check: resource.ComposeTestCheckFunc(
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

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("credential_id").Exists(),
			),
		},
		data.ImportStep(),
		{
			Config: r.complete(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("credential_id").Exists(),
			),
		},
		data.ImportStep(),
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("credential_id").Exists(),
			),
		},
		data.ImportStep(),
	})
}

func (r ApplicationFederatedIdentityCredentialResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.Applications.ApplicationsClient
	client.BaseClient.DisableRetries = true

	id, err := parse.FederatedIdentityCredentialID(state.ID)
	if err != nil {
		return nil, fmt.Errorf("parsing Application Federated Identity Credential ID: %v", err)
	}

	credential, status, err := client.GetFederatedIdentityCredential(ctx, id.ObjectId, id.KeyId, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return nil, fmt.Errorf("Federated Identity Credential %q for Application with object ID %q does not exist", id.KeyId, id.ObjectId)
		}
		return nil, fmt.Errorf("failed to retrieve Federated Identity Credential %q for Application with object ID %q: %+v", id.KeyId, id.ObjectId, err)
	}

	return utils.Bool(credential != nil), nil
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
  application_object_id = azuread_application.test.object_id
  display_name          = "hashitown-%[2]s"
  audiences             = ["api://HashiTownLikesAzureAD"]
  issuer                = "https://tokens.hashitown.net"
  subject               = "%[3]s"
}
`, r.template(data), data.RandomString, data.RandomID)
}

func (r ApplicationFederatedIdentityCredentialResource) complete(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_application_federated_identity_credential" "test" {
  application_object_id = azuread_application.test.object_id
  display_name          = "hashitown-%[2]s"
  description           = "Funtime tokens for HashiTown"
  audiences             = ["api://HashiTownLikesAzureAD"]
  issuer                = "https://vending.hashitown.net"
  subject               = "%[3]s"
}
`, r.template(data), data.RandomString, data.UUID())
}
