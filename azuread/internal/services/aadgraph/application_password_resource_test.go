package aadgraph_test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/ar"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/graph"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/internal/acceptance"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/internal/clients"
)

func testCheckApplicationPasswordExists(name string) resource.TestCheckFunc { //nolint unparam
	return func(s *terraform.State) error {
		client := acceptance.AzureADProvider.Meta().(*clients.AadClient).AadGraph.ApplicationsClient
		ctx := acceptance.AzureADProvider.Meta().(*clients.AadClient).StopContext

		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %q", name)
		}

		id, err := graph.ParseCredentialId(rs.Primary.ID)
		if err != nil {
			return fmt.Errorf("error parsing Application Password Credential ID: %v", err)
		}
		resp, err := client.Get(ctx, id.ObjectId)
		if err != nil {
			if ar.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("Bad: Azure AD Application  %q does not exist", id.ObjectId)
			}
			return fmt.Errorf("Bad: Get on Azure AD applicationsClient: %+v", err)
		}

		credentials, err := client.ListPasswordCredentials(ctx, id.ObjectId)
		if err != nil {
			return fmt.Errorf("Error Listing Password Credentials for Application %q: %+v", id.ObjectId, err)
		}

		cred := graph.PasswordCredentialResultFindByKeyId(credentials, id.KeyId)
		if cred != nil {
			return nil
		}

		return fmt.Errorf("Password Credential %q was not found in Application %q", id.KeyId, id.ObjectId)
	}
}

func testCheckApplicationPasswordCheckDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		client := acceptance.AzureADProvider.Meta().(*clients.AadClient).AadGraph.ApplicationsClient
		ctx := acceptance.AzureADProvider.Meta().(*clients.AadClient).StopContext

		if rs.Type != "azuread_application_password" {
			continue
		}

		id, err := graph.ParseCredentialId(rs.Primary.ID)
		if err != nil {
			return fmt.Errorf("error parsing Application Password Credential ID: %v", err)
		}

		resp, err := client.Get(ctx, id.ObjectId)
		if err != nil {
			if ar.ResponseWasNotFound(resp.Response) {
				return nil
			}

			return err
		}

		return fmt.Errorf("Application Password Credential still exists:\n%#v", resp)
	}

	return nil
}

func TestAccApplicationPassword_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_password", "test")
	applicationId := uuid.New().String()
	value := uuid.New().String()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckApplicationPasswordCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADObjectPasswordApplication_basic(applicationId, value),
				Check: resource.ComposeTestCheckFunc(
					testCheckApplicationPasswordExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "start_date"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "key_id"),
					resource.TestCheckResourceAttr(data.ResourceName, "end_date", "2099-01-01T01:02:03Z"),
				),
			},
			data.ImportStep("value"),
		},
	})
}

func TestAccApplicationPassword_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_password", "test")
	applicationId := uuid.New().String()
	value := uuid.New().String()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckApplicationPasswordCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADObjectPasswordApplication_basic(applicationId, value),
				Check: resource.ComposeTestCheckFunc(
					testCheckApplicationPasswordExists(data.ResourceName),
				),
			},
			data.RequiresImportErrorStep(testAccApplicationPassword_requiresImport(applicationId, value)),
		},
	})
}

func TestAccApplicationPassword_customKeyId(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_password", "test")
	applicationId := uuid.New().String()
	keyId := uuid.New().String()
	value := uuid.New().String()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckApplicationPasswordCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccApplicationPassword_customKeyId(applicationId, keyId, value),
				Check: resource.ComposeTestCheckFunc(
					testCheckApplicationPasswordExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "start_date"),
					resource.TestCheckResourceAttr(data.ResourceName, "key_id", keyId),
					resource.TestCheckResourceAttr(data.ResourceName, "end_date", "2099-01-01T01:02:03Z"),
				),
			},
			data.ImportStep("value"),
		},
	})
}

func TestAccApplicationPassword_description(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_password", "test")
	applicationId := uuid.New().String()
	value := uuid.New().String()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckApplicationPasswordCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccApplicationPassword_description(applicationId, value),
				Check: resource.ComposeTestCheckFunc(
					testCheckApplicationPasswordExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "start_date"),
					resource.TestCheckResourceAttr(data.ResourceName, "description", "terraform"),
					resource.TestCheckResourceAttr(data.ResourceName, "end_date", "2099-01-01T01:02:03Z"),
				),
			},
			data.ImportStep("value"),
		},
	})
}

func TestAccApplicationPassword_relativeEndDate(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_password", "test")
	applicationId := uuid.New().String()
	value := uuid.New().String()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckApplicationPasswordCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccApplicationPassword_relativeEndDate(applicationId, value),
				Check: resource.ComposeTestCheckFunc(
					testCheckApplicationPasswordExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "start_date"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "key_id"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "end_date"),
				),
			},
			data.ImportStep("value", "end_date_relative"),
		},
	})
}

func testAccApplicationPassword_template(applicationId string) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name = "acctestApp-%s"
}
`, applicationId)
}

func testAccADObjectPasswordApplication_basic(applicationId, value string) string {
	return fmt.Sprintf(`
%s

resource "azuread_application_password" "test" {
  application_object_id = azuread_application.test.id
  value                 = "%s"
  end_date              = "2099-01-01T01:02:03Z"
}
`, testAccApplicationPassword_template(applicationId), value)
}

func testAccApplicationPassword_requiresImport(applicationId, value string) string {
	template := testAccADObjectPasswordApplication_basic(applicationId, value)
	return fmt.Sprintf(`
%s

resource "azuread_application_password" "import" {
  application_object_id = azuread_application_password.test.application_object_id
  key_id                = azuread_application_password.test.key_id
  value                 = azuread_application_password.test.value
  end_date              = azuread_application_password.test.end_date
}
`, template)
}

func testAccApplicationPassword_customKeyId(applicationId, keyId, value string) string {
	return fmt.Sprintf(`
%s

resource "azuread_application_password" "test" {
  application_object_id = azuread_application.test.id
  key_id                = "%s"
  value                 = "%s"
  end_date              = "2099-01-01T01:02:03Z"
}
`, testAccApplicationPassword_template(applicationId), keyId, value)
}

func testAccApplicationPassword_description(applicationId, value string) string {
	return fmt.Sprintf(`
%s

resource "azuread_application_password" "test" {
  application_object_id = azuread_application.test.id
  description           = "terraform"
  value                 = "%s"
  end_date              = "2099-01-01T01:02:03Z"
}
`, testAccApplicationPassword_template(applicationId), value)
}

func testAccApplicationPassword_relativeEndDate(applicationId, value string) string {
	return fmt.Sprintf(`
%s

resource "azuread_application_password" "test" {
  application_object_id = azuread_application.test.id
  value                 = "%s"
  end_date_relative     = "8760h"
}
`, testAccApplicationPassword_template(applicationId), value)
}
