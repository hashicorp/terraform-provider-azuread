// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serviceprincipals_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/serviceprincipals/parse"
)

type servicePrincipalTokenSigningCertificateResource struct{}

func TestAccServicePrincipalTokenSigningCertificate_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_service_principal_token_signing_certificate", "test")
	r := servicePrincipalTokenSigningCertificateResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.basic(data),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("display_name").Exists(),
				check.That(data.ResourceName).Key("end_date").Exists(),
				check.That(data.ResourceName).Key("key_id").Exists(),
				check.That(data.ResourceName).Key("thumbprint").Exists(),
				check.That(data.ResourceName).Key("value").Exists(),
			),
		},
		data.ImportStep(),
	})
}

func TestAccServicePrincipalTokenSigningCertificate_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_service_principal_token_signing_certificate", "test")
	endDate := time.Now().AddDate(0, 3, 27).UTC().Format(time.RFC3339)
	r := servicePrincipalTokenSigningCertificateResource{}

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.complete(data, endDate),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("display_name").Exists(),
				check.That(data.ResourceName).Key("end_date").Exists(),
				check.That(data.ResourceName).Key("key_id").Exists(),
				check.That(data.ResourceName).Key("thumbprint").Exists(),
				check.That(data.ResourceName).Key("value").Exists(),
			),
		},
		data.ImportStep(),
	})
}

func (r servicePrincipalTokenSigningCertificateResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	client := clients.ServicePrincipals.ServicePrincipalsClient
	client.BaseClient.DisableRetries = true
	defer func() { client.BaseClient.DisableRetries = false }()

	id, err := parse.SigningCertificateID(state.ID)
	if err != nil {
		return nil, fmt.Errorf("parsing Service Principal Token Signing Certificate ID: %v", err)
	}

	servicePrincipal, status, err := client.Get(ctx, id.ObjectId, odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return nil, fmt.Errorf("Service Principal with object ID %q does not exist", id.ObjectId)
		}
		return nil, fmt.Errorf("failed to retrieve Service Principal with object ID %q: %+v", id.ObjectId, err)
	}

	if servicePrincipal.KeyCredentials != nil {
		for _, cred := range *servicePrincipal.KeyCredentials {
			if cred.KeyId != nil && *cred.KeyId == id.KeyId {
				return pointer.To(true), nil
			}
		}
	}

	return nil, fmt.Errorf("Token Signing Key Credential %q was not found for Service Principal %q", id.KeyId, id.ObjectId)
}

func (servicePrincipalTokenSigningCertificateResource) template(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  display_name = "acctestServicePrincipal-%[1]d"
}

resource "azuread_service_principal" "test" {
  application_id = azuread_application.test.application_id
}
`, data.RandomInteger)
}

func (r servicePrincipalTokenSigningCertificateResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_service_principal_token_signing_certificate" "test" {
  service_principal_id = azuread_service_principal.test.id
}
`, r.template(data))
}

func (r servicePrincipalTokenSigningCertificateResource) complete(data acceptance.TestData, endDate string) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_service_principal_token_signing_certificate" "test" {
  service_principal_id = azuread_service_principal.test.id
  display_name         = "CN=acctestTokenSigningCert-%[2]s"
  end_date             = "%[3]s"
}
`, r.template(data), data.RandomID, endDate)
}
