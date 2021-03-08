package serviceprincipals_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance"
	"github.com/hashicorp/terraform-provider-azuread/internal/acceptance/check"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/aadgraph"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/serviceprincipals/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
)

const servicePrincipalCertificatePem string = `-----BEGIN CERTIFICATE-----
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

const servicePrincipalCertificateBase64 string = `MIIDLDCCAhSgAwIBAgIQLSZ4E7hXTw+nb8YavHIoLjANBgkqhkiG9w0BAQsFADAYMRYwFAYDVQQDEw10ZXN0cmcta3YtYXBwMB4XDTIxMDExOTE4MjczMVoXDTIyMDExOTE4MzczMVowGDEWMBQGA1UEAxMNdGVzdHJnLWt2LWFwcDCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAL+HfZqmm57ngIYuzQBWnH2Yw/7u4h5xJ+F4/7U5nGABcUJQOMH+bXpZUz6LpwaXfQ70l8zmEPQv2qIfDs8TzcH0DOi2mOgM8eQoaUkOUeu4AXBBcRcVgTURH5HkSbEYMsxyaiinrvn5+KoQJcgVj8dZdcN+YxZr+ZgTaHGxjirTJEt6aGt+zr2gsZi8m8qGAQuIJbhPvBUk36VmriEIQR3ReigjT0yRCwBezsXL7EZ+WEdZB6p2UFGkXLq7coSkEA9UHLB0pMtLn74RbN6S395VnW4Vk3fgSfinysfIdro5UChC9R6OA9pWSgR0dxRw5AO0JMU8YHZnsajedpGREUUCAwEAAaNyMHAwDgYDVR0PAQH/BAQDAgG+MAkGA1UdEwQCMAAwEwYDVR0lBAwwCgYIKwYBBQUHAwEwHwYDVR0jBBgwFoAUFQUOvXHBPa0ZpNNwARFn0+k4dD4wHQYDVR0OBBYEFBUFDr1xwT2tGaTTcAERZ9PpOHQ+MA0GCSqGSIb3DQEBCwUAA4IBAQCf1LwnUVoBoYluey2kTn5rdI4As0pMg9nfec8xAWh3BjbYTjElcce+IP73TLzTPLe1lR2PlY/QcHuvfx8Orkm4JLBHrIUEcDh+G12qjMKU1GYtSUFj/QYwAPvesjryO0ow2XP+JgTj4yxVyXTcYfwT4t7gBlV50B0BXY/us3Mu/laczlN+xIonIPzIX5ZZQZBwDNmc1EPUjSN7KZ9AzvrMB6EcXYrP7IM7xEJzhCb/zgdIgGYGp0sMxOb8EZnMnIYmForwLbvryAZ3iN1RPCmxjXPRex1IfWxUY1PhCLjU3LUch6aHHx3YYp/GMg8j5DlfyD4WtIxqIUpJP5uE/e8Q`

const servicePrincipalCertificateHex string = `3082032C30820214A00302010202102D267813B8574F0FA76FC61ABC72282E300D06092A864886F70D01010B05003018311630140603550403130D7465737472672D6B762D617070301E170D3231303131393138323733315A170D3232303131393138333733315A3018311630140603550403130D7465737472672D6B762D61707030820122300D06092A864886F70D01010105000382010F003082010A0282010100BF877D9AA69B9EE780862ECD00569C7D98C3FEEEE21E7127E178FFB5399C600171425038C1FE6D7A59533E8BA706977D0EF497CCE610F42FDAA21F0ECF13CDC1F40CE8B698E80CF1E42869490E51EBB80170417117158135111F91E449B11832CC726A28A7AEF9F9F8AA1025C8158FC75975C37E63166BF998136871B18E2AD3244B7A686B7ECEBDA0B198BC9BCA86010B8825B84FBC1524DFA566AE2108411DD17A28234F4C910B005ECEC5CBEC467E58475907AA765051A45CBABB7284A4100F541CB074A4CB4B9FBE116CDE92DFDE559D6E159377E049F8A7CAC7C876BA39502842F51E8E03DA564A0474771470E403B424C53C607667B1A8DE76919111450203010001A3723070300E0603551D0F0101FF0404030201BE30090603551D130402300030130603551D25040C300A06082B06010505070301301F0603551D2304183016801415050EBD71C13DAD19A4D370011167D3E938743E301D0603551D0E0416041415050EBD71C13DAD19A4D370011167D3E938743E300D06092A864886F70D01010B050003820101009FD4BC27515A01A1896E7B2DA44E7E6B748E00B34A4C83D9DF79CF310168770636D84E312571C7BE20FEF74CBCD33CB7B5951D8F958FD0707BAF7F1F0EAE49B824B047AC850470387E1B5DAA8CC294D4662D494163FD063000FBDEB23AF23B4A30D973FE2604E3E32C55C974DC61FC13E2DEE0065579D01D015D8FEEB3732EFE569CCE537EC48A2720FCC85F96594190700CD99CD443D48D237B299F40CEFACC07A11C5D8ACFEC833BC442738426FFCE0748806606A74B0CC4E6FC1199CC9C8626168AF02DBBEBC8067788DD513C29B18D73D17B1D487D6C546353E108B8D4DCB51C87A6871F1DD8629FC6320F23E4395FC83E16B48C6A214A493F9B84FDEF10`

type ServicePrincipalCertificateResource struct{}

func TestAccServicePrincipalCertificate_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_service_principal_certificate", "test")
	endDate := time.Now().AddDate(0, 3, 27).UTC().Format(time.RFC3339)
	r := ServicePrincipalCertificateResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data, endDate),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("key_id").Exists(),
			),
		},
		data.ImportStep("encoding", "value"),
	})
}

func TestAccServicePrincipalCertificate_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_service_principal_certificate", "test")
	startDate := time.Now().AddDate(0, 0, 7).UTC().Format(time.RFC3339)
	endDate := time.Now().AddDate(0, 3, 27).UTC().Format(time.RFC3339)
	r := ServicePrincipalCertificateResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.complete(data, startDate, endDate),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("key_id").Exists(),
			),
		},
		data.ImportStep("encoding", "end_date_relative", "value"),
	})
}

func TestAccServicePrincipalCertificate_base64Cert(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_service_principal_certificate", "test")
	endDate := time.Now().AddDate(0, 3, 27).UTC().Format(time.RFC3339)
	r := ServicePrincipalCertificateResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.base64Cert(data, endDate),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("key_id").Exists(),
			),
		},
		data.ImportStep("encoding", "end_date_relative", "value"),
	})
}

func TestAccServicePrincipalCertificate_hexCert(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_service_principal_certificate", "test")
	endDate := time.Now().AddDate(0, 3, 27).UTC().Format(time.RFC3339)
	r := ServicePrincipalCertificateResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.hexCert(data, endDate),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("key_id").Exists(),
			),
		},
		data.ImportStep("encoding", "end_date_relative", "value"),
	})
}

func TestAccServicePrincipalCertificate_relativeEndDate(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_service_principal_certificate", "test")
	r := ServicePrincipalCertificateResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.relativeEndDate(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("key_id").Exists(),
				check.That(data.ResourceName).Key("end_date").Exists(),
			),
		},
		data.ImportStep("encoding", "end_date_relative", "value"),
	})
}

func TestAccServicePrincipalCertificate_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_service_principal_certificate", "test")
	endDate := time.Now().AddDate(0, 3, 27).UTC().Format(time.RFC3339)
	r := ServicePrincipalCertificateResource{}

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

func (r ServicePrincipalCertificateResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	id, err := parse.CertificateID(state.ID)
	if err != nil {
		return nil, fmt.Errorf("parsing Service Principal Certificate ID: %v", err)
	}

	resp, err := clients.ServicePrincipals.AadClient.Get(ctx, id.ObjectId)
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			return nil, fmt.Errorf("Service Principal with object ID %q does not exist", id.ObjectId)
		}
		return nil, fmt.Errorf("Service Principal with object ID %q does not exist", id.ObjectId)
	}

	credentials, err := clients.ServicePrincipals.AadClient.ListKeyCredentials(ctx, id.ObjectId)
	if err != nil {
		return nil, fmt.Errorf("listing Key Credentials for Service Principal %q: %+v", id.ObjectId, err)
	}

	cred := aadgraph.KeyCredentialResultFindByKeyId(credentials, id.KeyId)
	if cred != nil {
		return utils.Bool(true), nil
	}

	return nil, fmt.Errorf("Key Credential %q was not found for Service Principal %q", id.KeyId, id.ObjectId)
}

func (ServicePrincipalCertificateResource) template(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name = "acctestServicePrincipal-%[1]d"
}

resource "azuread_service_principal" "test" {
  application_id = azuread_application.test.application_id
}
`, data.RandomInteger)
}

func (r ServicePrincipalCertificateResource) basic(data acceptance.TestData, endDate string) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_service_principal_certificate" "test" {
  service_principal_id = azuread_service_principal.test.id
  type                 = "AsymmetricX509Cert"
  end_date             = "%[2]s"
  value                = <<EOT
%[3]s
EOT
}
`, r.template(data), endDate, servicePrincipalCertificatePem)
}

func (r ServicePrincipalCertificateResource) complete(data acceptance.TestData, startDate, endDate string) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_service_principal_certificate" "test" {
  service_principal_id = azuread_service_principal.test.id
  key_id               = "%[2]s"
  type                 = "AsymmetricX509Cert"
  start_date           = "%[3]s"
  end_date             = "%[4]s"
  encoding             = "pem"
  value                = <<EOT
%[5]s
EOT
}
`, r.template(data), data.RandomID, startDate, endDate, servicePrincipalCertificatePem)
}

func (r ServicePrincipalCertificateResource) base64Cert(data acceptance.TestData, endDate string) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_service_principal_certificate" "test" {
  service_principal_id = azuread_service_principal.test.id
  type                 = "AsymmetricX509Cert"
  end_date             = "%[2]s"
  encoding             = "base64"
  value                = <<EOT
%[3]s
EOT
}
`, r.template(data), endDate, servicePrincipalCertificateBase64)
}

func (r ServicePrincipalCertificateResource) hexCert(data acceptance.TestData, endDate string) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_service_principal_certificate" "test" {
  service_principal_id = azuread_service_principal.test.id
  type                 = "AsymmetricX509Cert"
  end_date             = "%[2]s"
  encoding             = "hex"
  value                = <<EOT
%[3]s
EOT
}
`, r.template(data), endDate, servicePrincipalCertificateHex)
}

func (r ServicePrincipalCertificateResource) relativeEndDate(data acceptance.TestData) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_service_principal_certificate" "test" {
  service_principal_id = azuread_service_principal.test.id
  end_date_relative    = "2280h"
  type                 = "AsymmetricX509Cert"
  value                = <<EOT
%[2]s
EOT
}
`, r.template(data), servicePrincipalCertificatePem)
}

func (r ServicePrincipalCertificateResource) requiresImport(data acceptance.TestData, endDate string) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_service_principal_certificate" "import" {
  service_principal_id = azuread_service_principal_certificate.test.service_principal_id
  key_id               = azuread_service_principal_certificate.test.key_id
  type                 = azuread_service_principal_certificate.test.type
  end_date             = azuread_service_principal_certificate.test.end_date
  value                = azuread_service_principal_certificate.test.value
}
`, r.basic(data, endDate))
}
