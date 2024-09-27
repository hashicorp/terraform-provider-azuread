// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serviceprincipals_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/stable/serviceprincipal"
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
	client := clients.ServicePrincipals.ServicePrincipalClient

	id, err := parse.SigningCertificateID(state.ID)
	if err != nil {
		return nil, fmt.Errorf("parsing Service Principal Token Signing Certificate ID: %v", err)
	}

	servicePrincipalId := stable.NewServicePrincipalID(id.ObjectId)

	resp, err := client.GetServicePrincipal(ctx, servicePrincipalId, serviceprincipal.DefaultGetServicePrincipalOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return nil, fmt.Errorf("%s does not exist", servicePrincipalId)
		}
		return nil, fmt.Errorf("failed to retrieve %s: %v", servicePrincipalId, err)
	}

	if resp.Model != nil && resp.Model.KeyCredentials != nil {
		for _, cred := range *resp.Model.KeyCredentials {
			if cred.KeyId.GetOrZero() == id.KeyId {
				return pointer.To(true), nil
			}
		}
	}

	return pointer.To(false), nil
}

func (servicePrincipalTokenSigningCertificateResource) template(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  display_name = "acctestServicePrincipal-%[1]d"
}

resource "azuread_service_principal" "test" {
  client_id = azuread_application.test.client_id
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
