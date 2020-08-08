package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/ar"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/graph"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/internal/acceptance"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/internal/clients"
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

func testCheckApplicationKeyExists(name string) resource.TestCheckFunc { //nolint unparam
	return func(s *terraform.State) error {
		client := acceptance.AzureADProvider.Meta().(*clients.AadClient).ApplicationsClient
		ctx := acceptance.AzureADProvider.Meta().(*clients.AadClient).StopContext

		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %q", name)
		}

		id, err := graph.ParseCredentialId(rs.Primary.ID)
		if err != nil {
			return fmt.Errorf("parsing Application Key Credential ID: %v", err)
		}
		resp, err := client.Get(ctx, id.ObjectId)
		if err != nil {
			if ar.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("Bad: Azure AD Application %q does not exist", id.ObjectId)
			}
			return fmt.Errorf("Bad: Get on Azure AD applicationsClient: %+v", err)
		}

		credentials, err := client.ListKeyCredentials(ctx, id.ObjectId)
		if err != nil {
			return fmt.Errorf("Error Listing Key Credentials for Application %q: %+v", id.ObjectId, err)
		}

		cred := graph.KeyCredentialResultFindByKeyId(credentials, id.KeyId)
		if cred != nil {
			return nil
		}

		return fmt.Errorf("Key Credential %q was not found in Application %q", id.KeyId, id.ObjectId)
	}
}

func testCheckApplicationKeyCheckDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		client := acceptance.AzureADProvider.Meta().(*clients.AadClient).ApplicationsClient
		ctx := acceptance.AzureADProvider.Meta().(*clients.AadClient).StopContext

		if rs.Type != "azuread_application_certificate" {
			continue
		}

		id, err := graph.ParseCredentialId(rs.Primary.ID)
		if err != nil {
			return fmt.Errorf("error parsing Application Credential ID: %v", err)
		}

		resp, err := client.Get(ctx, id.ObjectId)
		if err != nil {
			if ar.ResponseWasNotFound(resp.Response) {
				return nil
			}

			return err
		}

		return fmt.Errorf("Azure AD Application Key Credential still exists:\n%#v", resp)
	}

	return nil
}

func TestAccAzureApplicationCertificate_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_certificate", "test")
	keyType := "AsymmetricX509Cert"
	endDate := time.Now().AddDate(0, 6, 0).UTC().Format(time.RFC3339)
	value := testCertificateApplication

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckApplicationKeyCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADObjectCertificateApplication_basic(data.RandomInteger, keyType, endDate, value),
				Check: resource.ComposeTestCheckFunc(
					testCheckApplicationKeyExists(data.ResourceName),
				),
			},
			data.ImportStep("end_date_relative", "value"),
		},
	})
}

func TestAccAzureApplicationCertificate_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_certificate", "test")
	keyId := uuid.New().String()
	keyType := "AsymmetricX509Cert"
	startDate := time.Now().AddDate(0, 0, 7).UTC().Format(time.RFC3339)
	endDate := time.Now().AddDate(0, 6, 0).UTC().Format(time.RFC3339)
	value := testCertificateApplication

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckApplicationKeyCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccApplicationCertificate_complete(data.RandomInteger, keyId, keyType, startDate, endDate, value),
				Check: resource.ComposeTestCheckFunc(
					testCheckApplicationKeyExists(data.ResourceName),
				),
			},
			data.ImportStep("end_date_relative", "value"),
		},
	})
}

func TestAccAzureApplicationCertificate_relativeEndDate(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_certificate", "test")
	keyType := "AsymmetricX509Cert"
	value := testCertificateApplication

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckApplicationKeyCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccApplicationCertificate_relativeEndDate(data.RandomInteger, keyType, value),
				Check: resource.ComposeTestCheckFunc(
					// can't assert on Value since it's not returned
					testCheckApplicationKeyExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "end_date"),
				),
			},
			data.ImportStep("end_date_relative", "value"),
		},
	})
}

func TestAccAzureApplicationCertificate_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_certificate", "test")
	keyType := "AsymmetricX509Cert"
	endDate := time.Now().AddDate(0, 6, 0).UTC().Format(time.RFC3339)
	value := testCertificateApplication

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckApplicationKeyCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADObjectCertificateApplication_basic(data.RandomInteger, keyType, endDate, value),
				Check: resource.ComposeTestCheckFunc(
					testCheckApplicationKeyExists(data.ResourceName),
				),
			},
			data.RequiresImportErrorStep(testAccApplicationCertificate_requiresImport(data.RandomInteger, keyType, endDate, value)),
		},
	})
}

func testAccApplicationCertificate_template(ri int) string {
	return fmt.Sprintf(`
resource "azuread_application" "tests" {
  name = "acctestApp-%d"
}
`, ri)
}

func testAccADObjectCertificateApplication_basic(ri int, keyType, endDate, value string) string {
	return fmt.Sprintf(`
%s

resource "azuread_application_certificate" "test" {
  application_object_id = azuread_application.tests.id
  type                  = "%s"
  end_date              = "%s"
  value                 = <<EOT
%s
EOT
}
`, testAccApplicationCertificate_template(ri), keyType, endDate, value)
}

func testAccApplicationCertificate_complete(ri int, keyId, keyType, startDate, endDate, value string) string {
	return fmt.Sprintf(`
%s

resource "azuread_application_certificate" "test" {
  application_object_id = azuread_application.tests.id
  key_id                = "%s"
  type                  = "%s"
  start_date            = "%s"
  end_date              = "%s"
  value                 = <<EOT
%s
EOT
}
`, testAccApplicationCertificate_template(ri), keyId, keyType, startDate, endDate, value)
}

func testAccApplicationCertificate_relativeEndDate(ri int, keyType, value string) string {
	return fmt.Sprintf(`
%s

resource "azuread_application_certificate" "test" {
  application_object_id = azuread_application.tests.id
  end_date_relative     = "4320h"
  type                  = "%s"
  value                 = <<EOT
%s
EOT
}
`, testAccApplicationCertificate_template(ri), keyType, value)
}

func testAccApplicationCertificate_requiresImport(ri int, keyType, endDate, value string) string {
	template := testAccADObjectCertificateApplication_basic(ri, keyType, endDate, value)
	return fmt.Sprintf(`
%s

resource "azuread_application_certificate" "import" {
  application_object_id = azuread_application_certificate.test.application_object_id
  key_id                = azuread_application_certificate.test.key_id
  type                  = azuread_application_certificate.test.type
  end_date              = azuread_application_certificate.test.end_date
  value                 = azuread_application_certificate.test.value
}
`, template)
}
