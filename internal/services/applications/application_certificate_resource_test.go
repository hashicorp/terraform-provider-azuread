package applications_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/terraform-providers/terraform-provider-azuread/internal/acceptance"
	"github.com/terraform-providers/terraform-provider-azuread/internal/acceptance/check"
	"github.com/terraform-providers/terraform-provider-azuread/internal/clients"
	"github.com/terraform-providers/terraform-provider-azuread/internal/helpers/aadgraph"
	"github.com/terraform-providers/terraform-provider-azuread/internal/services/applications/parse"
	"github.com/terraform-providers/terraform-provider-azuread/internal/utils"
)

const testCertificateApplication string = `-----BEGIN CERTIFICATE-----
MIIDGjCCAgICCQDAQlCA1jw1BjANBgkqhkiG9w0BAQsFADBPMQswCQYDVQQGEwJV
UzELMAkGA1UECAwCQ0ExFzAVBgNVBAoMDkhhc2hpQ29ycCwgSW5jMRowGAYDVQQD
DBFoYXNoaWNvcnB0ZXN0LmNvbTAeFw0yMDA1MzEyMDI2MTFaFw0yMTA1MzEyMDI2
MTFaME8xCzAJBgNVBAYTAlVTMQswCQYDVQQIDAJDQTEXMBUGA1UECgwOSGFzaGlD
b3JwLCBJbmMxGjAYBgNVBAMMEWhhc2hpY29ycHRlc3QuY29tMIIBIjANBgkqhkiG
9w0BAQEFAAOCAQ8AMIIBCgKCAQEAmVsolhGq5SKLB4OJrQqlBmqOMTZLNKeAved5
f0pNmr9Sdb7VapA8vVALENJq0XWDsY5blrsDevYtPVh3ZXKfZpmJwbtcy5ZJ+B6B
HVXZHO2Ep3RzZ/bFOxXXvmtUGmOpJxJyHlXao4mz2LuNHUa6GJqDjVaZ3ZL2EjFa
Q2nRIorIYoERHxEcogpQqF0csL70hi8Ho/dFpfsKhooYfzWRgj4KncMbjWYb82L6
qPhnonETOGHgpqXP7FhGLkm0pfWL1oz21hZRWRVh6+RjBQ4gm3iRrzffyjvvdVL8
molsNguVYF5Km7D/oRTILImPgXLlKnB/XRpORvhDK4zw8ewiewIDAQABMA0GCSqG
SIb3DQEBCwUAA4IBAQB7XGxsoU1DF7ECvuR5T9cPrW60DxjPP7uEp2UQeEZgkCDD
WpAqmhXrtn2fRpKVBCFzTVZZKW3/f8W3ivGR9hnVTmPtvjG9n9wLq997k8h2GfD8
z4YaA8vEwqmrMYX8azeTGL/JQHFIleA6YPokdybgP6aC9hIWdSVXdV4G5kgN3/GV
lWOie5Wf9IZotaDzExG7P38mGzQOLtZVCnIxGyy6/Q1E5n5PUxc/9i/iY6xlC2Ih
HraQzsK7BNxC5NSwwirT95JH+Xd8rvWu+bCveJz3mnZ3sgolCoxL6Hv1uD2UOZb5
rCHdW31vp5PYNJaSkYL0j259Ogb8crkIzDr3Z8YF
-----END CERTIFICATE-----`

type ApplicationCertificateResource struct{}

func TestAccApplicationCertificate_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_certificate", "test")
	endDate := time.Now().AddDate(0, 3, 27).UTC().Format(time.RFC3339)
	r := ApplicationCertificateResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data, endDate),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("key_id").Exists(),
			),
		},
		data.ImportStep("end_date_relative", "value"),
	})
}

func TestAccApplicationCertificate_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_certificate", "test")
	startDate := time.Now().AddDate(0, 0, 7).UTC().Format(time.RFC3339)
	endDate := time.Now().AddDate(0, 3, 27).UTC().Format(time.RFC3339)
	r := ApplicationCertificateResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.complete(data, startDate, endDate),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("key_id").Exists(),
			),
		},
		data.ImportStep("end_date_relative", "value"),
	})
}

func TestAccApplicationCertificate_relativeEndDate(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_certificate", "test")
	r := ApplicationCertificateResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.relativeEndDate(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("key_id").Exists(),
				check.That(data.ResourceName).Key("end_date").Exists(),
			),
		},
		data.ImportStep("end_date_relative", "value"),
	})
}

func TestAccApplicationCertificate_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_certificate", "test")
	endDate := time.Now().AddDate(0, 3, 27).UTC().Format(time.RFC3339)
	r := ApplicationCertificateResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data, endDate),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("key_id").Exists(),
			),
		},
		data.RequiresImportErrorStep(r.requiresImport(data, endDate)),
	})
}

func (ApplicationCertificateResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	id, err := parse.CertificateID(state.ID)
	if err != nil {
		return nil, fmt.Errorf("parsing Application Certificate ID: %v", err)
	}

	resp, err := clients.Applications.AadClient.Get(ctx, id.ObjectId)
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			return nil, fmt.Errorf("Application with object ID %q does not exist", id.ObjectId)
		}
		return nil, fmt.Errorf("failed to retrieve Application with object ID %q: %+v", id.ObjectId, err)
	}

	credentials, err := clients.Applications.AadClient.ListKeyCredentials(ctx, id.ObjectId)
	if err != nil {
		return nil, fmt.Errorf("listing Key Credentials for Application %q: %+v", id.ObjectId, err)
	}

	cred := aadgraph.KeyCredentialResultFindByKeyId(credentials, id.KeyId)
	if cred != nil {
		return utils.Bool(true), nil
	}

	return nil, fmt.Errorf("Key Credential %q was not found for Application %q", id.KeyId, id.ObjectId)
}

func (ApplicationCertificateResource) template(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name = "acctestApp-%[1]d"
}
`, data.RandomInteger)
}

func (r ApplicationCertificateResource) basic(data acceptance.TestData, endDate string) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_application_certificate" "test" {
  application_object_id = azuread_application.test.id
  type                  = "AsymmetricX509Cert"
  end_date              = "%[2]s"
  value                 = <<EOT
%[3]s
EOT
}
`, r.template(data), endDate, testCertificateApplication)
}

func (r ApplicationCertificateResource) complete(data acceptance.TestData, startDate, endDate string) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_application_certificate" "test" {
  application_object_id = azuread_application.test.id
  key_id                = "%[2]s"
  type                  = "AsymmetricX509Cert"
  start_date            = "%[3]s"
  end_date              = "%[4]s"
  value                 = <<EOT
%[5]s
EOT
}
`, r.template(data), data.RandomID, startDate, endDate, testCertificateApplication)
}

func (r ApplicationCertificateResource) relativeEndDate(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_application_certificate" "test" {
  application_object_id = azuread_application.test.id
  end_date_relative     = "2280h"
  type                  = "AsymmetricX509Cert"
  value                 = <<EOT
%[2]s
EOT
}
`, r.template(data), testCertificateApplication)
}

func (r ApplicationCertificateResource) requiresImport(data acceptance.TestData, endDate string) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_application_certificate" "import" {
  application_object_id = azuread_application_certificate.test.application_object_id
  key_id                = azuread_application_certificate.test.key_id
  type                  = azuread_application_certificate.test.type
  end_date              = azuread_application_certificate.test.end_date
  value                 = azuread_application_certificate.test.value
}
`, r.basic(data, endDate))
}
