package aadgraph_test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-azuread/azuread/internal/acceptance"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/internal/ar"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/internal/clients"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/internal/services/aadgraph/graph"
)

func testCheckServicePrincipalPasswordExists(name string) resource.TestCheckFunc { //nolint unparam
	return func(s *terraform.State) error {
		client := acceptance.AzureADProvider.Meta().(*clients.AadClient).AadGraph.ServicePrincipalsClient
		ctx := acceptance.AzureADProvider.Meta().(*clients.AadClient).StopContext

		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %q", name)
		}

		id, err := graph.ParseCredentialId(rs.Primary.ID)
		if err != nil {
			return fmt.Errorf("error Service Principal Password Credential ID: %v", err)
		}

		resp, err := client.Get(ctx, id.ObjectId)
		if err != nil {
			if ar.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("Bad: Azure AD Service Principal %q does not exist", id.ObjectId)
			}
			return fmt.Errorf("Bad: Get on Azure AD ServicePrincipalsClient: %+v", err)
		}

		credentials, err := client.ListPasswordCredentials(ctx, id.ObjectId)
		if err != nil {
			return fmt.Errorf("Error Listing Password Credentials for Service Principal %q: %+v", id.ObjectId, err)
		}

		cred := graph.PasswordCredentialResultFindByKeyId(credentials, id.KeyId)
		if cred != nil {
			return nil
		}

		return fmt.Errorf("Password Credential %q was not found in Service Principal %q", id.KeyId, id.ObjectId)
	}
}

func testCheckServicePrincipalPasswordCheckDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		client := acceptance.AzureADProvider.Meta().(*clients.AadClient).AadGraph.ApplicationsClient
		ctx := acceptance.AzureADProvider.Meta().(*clients.AadClient).StopContext

		if rs.Type != "azuread_service_principal_password" {
			continue
		}

		id, err := graph.ParseCredentialId(rs.Primary.ID)
		if err != nil {
			return fmt.Errorf("error parsing Service Principal Password Credential ID: %v", err)
		}

		resp, err := client.Get(ctx, id.ObjectId)
		if err != nil {
			if ar.ResponseWasNotFound(resp.Response) {
				return nil
			}

			return err
		}

		return fmt.Errorf("Azure AD Service Principal Password Credential still exists:\n%#v", resp)
	}

	return nil
}

func TestAccServicePrincipalPassword_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_service_principal_password", "test")
	applicationId := uuid.New().String()
	value := uuid.New().String()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckServicePrincipalPasswordCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccServicePrincipalPassword_basic(applicationId, value),
				Check: resource.ComposeTestCheckFunc(
					testCheckServicePrincipalPasswordExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "start_date"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "key_id"),
					resource.TestCheckResourceAttr(data.ResourceName, "end_date", "2099-01-01T01:02:03Z"),
				),
			},
			data.ImportStep("value"),
		},
	})
}

func TestAccServicePrincipalPassword_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_service_principal_password", "test")
	applicationId := uuid.New().String()
	value := uuid.New().String()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckServicePrincipalPasswordCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccServicePrincipalPassword_basic(applicationId, value),
				Check: resource.ComposeTestCheckFunc(
					testCheckServicePrincipalPasswordExists(data.ResourceName),
				),
			},
			data.RequiresImportErrorStep(testAccServicePrincipalPassword_requiresImport(applicationId, value)),
		},
	})
}

func TestAccServicePrincipalPassword_customKeyId(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_service_principal_password", "test")
	applicationId := uuid.New().String()
	keyId := uuid.New().String()
	value := uuid.New().String()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckServicePrincipalPasswordCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccServicePrincipalPassword_customKeyId(applicationId, keyId, value),
				Check: resource.ComposeTestCheckFunc(
					testCheckServicePrincipalPasswordExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "start_date"),
					resource.TestCheckResourceAttr(data.ResourceName, "key_id", keyId),
					resource.TestCheckResourceAttr(data.ResourceName, "end_date", "2099-01-01T01:02:03Z"),
				),
			},
			data.ImportStep("value"),
		},
	})
}

func TestAccServicePrincipalPassword_description(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_service_principal_password", "test")
	applicationId := uuid.New().String()
	value := uuid.New().String()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckServicePrincipalPasswordCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccServicePrincipalPassword_description(applicationId, value),
				Check: resource.ComposeTestCheckFunc(
					testCheckServicePrincipalPasswordExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "start_date"),
					resource.TestCheckResourceAttr(data.ResourceName, "description", "terraform"),
					resource.TestCheckResourceAttr(data.ResourceName, "end_date", "2099-01-01T01:02:03Z"),
				),
			},
			data.ImportStep("value"),
		},
	})
}

func TestAccServicePrincipalPassword_relativeEndDate(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_service_principal_password", "test")
	applicationId := uuid.New().String()
	value := uuid.New().String()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckServicePrincipalPasswordCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccServicePrincipalPassword_relativeEndDate(applicationId, value),
				Check: resource.ComposeTestCheckFunc(
					testCheckServicePrincipalPasswordExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "start_date"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "key_id"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "end_date"),
				),
			},
			data.ImportStep("end_date_relative", "value"),
		},
	})
}

func testAccServicePrincipalPassword_template(applicationId string) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name = "acctestApp-%s"
}

resource "azuread_service_principal" "test" {
  application_id = azuread_application.test.application_id
}
`, applicationId)
}

func testAccServicePrincipalPassword_basic(applicationId, value string) string {
	return fmt.Sprintf(`
%s

resource "azuread_service_principal_password" "test" {
  service_principal_id = azuread_service_principal.test.id
  value                = "%s"
  end_date             = "2099-01-01T01:02:03Z"
}
`, testAccServicePrincipalPassword_template(applicationId), value)
}

func testAccServicePrincipalPassword_requiresImport(applicationId, value string) string {
	template := testAccServicePrincipalPassword_basic(applicationId, value)
	return fmt.Sprintf(`
%s

resource "azuread_service_principal_password" "import" {
  key_id               = azuread_service_principal_password.test.key_id
  service_principal_id = azuread_service_principal_password.test.service_principal_id
  value                = azuread_service_principal_password.test.value
  end_date             = azuread_service_principal_password.test.end_date
}
`, template)
}

func testAccServicePrincipalPassword_customKeyId(applicationId, keyId, value string) string {
	return fmt.Sprintf(`
%s

resource "azuread_service_principal_password" "test" {
  service_principal_id = azuread_service_principal.test.id
  key_id               = "%s"
  value                = "%s"
  end_date             = "2099-01-01T01:02:03Z"
}
`, testAccServicePrincipalPassword_template(applicationId), keyId, value)
}

func testAccServicePrincipalPassword_description(applicationId, value string) string {
	return fmt.Sprintf(`
%s

resource "azuread_service_principal_password" "test" {
  service_principal_id = azuread_service_principal.test.id
  description          = "terraform"
  value                = "%s"
  end_date             = "2099-01-01T01:02:03Z"
}
`, testAccServicePrincipalPassword_template(applicationId), value)
}

func testAccServicePrincipalPassword_relativeEndDate(applicationId, value string) string {
	return fmt.Sprintf(`
%s

resource "azuread_service_principal_password" "test" {
  service_principal_id = azuread_service_principal.test.id
  value                = "%s"
  end_date_relative    = "8760h"
}
`, testAccServicePrincipalPassword_template(applicationId), value)
}
