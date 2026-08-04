// Copyright IBM Corp. 2014, 2025
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
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
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

func TestAccServicePrincipalTokenSigningCertificate_multiple(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_service_principal_token_signing_certificate", "blue")
	r := servicePrincipalTokenSigningCertificateResource{}

	blueAddress := "azuread_service_principal_token_signing_certificate.blue"
	greenAddress := "azuread_service_principal_token_signing_certificate.green"
	blueEndDate := time.Now().AddDate(0, 2, 0).UTC().Format(time.RFC3339)
	greenEndDate := time.Now().AddDate(0, 4, 0).UTC().Format(time.RFC3339)
	var greenKeyId string

	data.ResourceTest(t, r, []acceptance.TestStep{
		{
			Config: r.multiple(data, blueEndDate, greenEndDate),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(blueAddress).ExistsInAzure(r),
				check.That(greenAddress).ExistsInAzure(r),
				resource.TestCheckResourceAttrWith(greenAddress, "key_id", func(value string) error {
					greenKeyId = value
					return nil
				}),
			),
		},
		{
			// Removing "blue" from config forces its deletion. Only "blue" should be removed;
			// "green" must survive untouched with an unchanged key_id.
			Config: r.multipleGreenOnly(data, greenEndDate),
			Check: acceptance.ComposeTestCheckFunc(
				check.That(greenAddress).ExistsInAzure(r),
				resource.TestCheckResourceAttrWith(greenAddress, "key_id", func(value string) error {
					if value != greenKeyId {
						return fmt.Errorf("expected %q key_id to remain %q, got %q - deleting %q appears to have affected it", greenAddress, greenKeyId, value, blueAddress)
					}
					return nil
				}),
			),
		},
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

func (r servicePrincipalTokenSigningCertificateResource) multiple(data acceptance.TestData, blueEndDate, greenEndDate string) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_service_principal_token_signing_certificate" "blue" {
  service_principal_id = azuread_service_principal.test.id
  display_name         = "CN=acctestTokenSigningCertBlue-%[2]d"
  end_date             = "%[3]s"
}

resource "azuread_service_principal_token_signing_certificate" "green" {
  service_principal_id = azuread_service_principal.test.id
  display_name         = "CN=acctestTokenSigningCertGreen-%[2]d"
  end_date             = "%[4]s"
}
`, r.template(data), data.RandomInteger, blueEndDate, greenEndDate)
}

func (r servicePrincipalTokenSigningCertificateResource) multipleGreenOnly(data acceptance.TestData, greenEndDate string) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_service_principal_token_signing_certificate" "green" {
  service_principal_id = azuread_service_principal.test.id
  display_name         = "CN=acctestTokenSigningCertGreen-%[2]d"
  end_date             = "%[3]s"
}
`, r.template(data), data.RandomInteger, greenEndDate)
}
