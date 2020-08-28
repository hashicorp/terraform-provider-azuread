package aadgraph_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-azuread/internal/acceptance"
	"github.com/terraform-providers/terraform-provider-azuread/internal/clients"
	"github.com/terraform-providers/terraform-provider-azuread/internal/services/aadgraph/graph"
	"github.com/terraform-providers/terraform-provider-azuread/internal/utils"
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
			return fmt.Errorf("parsing Application Password Credential ID: %v", err)
		}
		resp, err := client.Get(ctx, id.ObjectId)
		if err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("Bad: Application  %q does not exist", id.ObjectId)
			}
			return fmt.Errorf("Bad: Get on applicationsClient: %+v", err)
		}

		credentials, err := client.ListPasswordCredentials(ctx, id.ObjectId)
		if err != nil {
			return fmt.Errorf("listing Password Credentials for Application %q: %+v", id.ObjectId, err)
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
			return fmt.Errorf("parsing Application Password Credential ID: %v", err)
		}

		resp, err := client.Get(ctx, id.ObjectId)
		if err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckApplicationPasswordCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccApplicationPassword_basic(data.RandomInteger, data.RandomPassword),
				Check: resource.ComposeTestCheckFunc(
					testCheckApplicationPasswordExists(data.ResourceName),
				),
			},
			data.ImportStep("value"),
		},
	})
}

func TestAccApplicationPassword_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_password", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckApplicationPasswordCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccApplicationPassword_complete(data.RandomInteger, data.RandomID, data.RandomPassword),
				Check: resource.ComposeTestCheckFunc(
					testCheckApplicationPasswordExists(data.ResourceName),
				),
			},
			data.ImportStep("value"),
		},
	})
}

func TestAccApplicationPassword_relativeEndDate(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_password", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckApplicationPasswordCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccApplicationPassword_relativeEndDate(data.RandomInteger, data.RandomPassword),
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

func TestAccApplicationPassword_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application_password", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckApplicationPasswordCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccApplicationPassword_basic(data.RandomInteger, data.RandomPassword),
				Check: resource.ComposeTestCheckFunc(
					testCheckApplicationPasswordExists(data.ResourceName),
				),
			},
			data.RequiresImportErrorStep(testAccApplicationPassword_requiresImport(data.RandomInteger, data.RandomPassword)),
		},
	})
}

func testAccApplicationPassword_template(ri int) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name = "acctestApp-%d"
}
`, ri)
}

func testAccApplicationPassword_basic(ri int, value string) string {
	return fmt.Sprintf(`
%s

resource "azuread_application_password" "test" {
  application_object_id = azuread_application.test.id
  value                 = "%s"
  end_date              = "2099-01-01T01:02:03Z"
}
`, testAccApplicationPassword_template(ri), value)
}

func testAccApplicationPassword_complete(ri int, keyId, value string) string {
	return fmt.Sprintf(`
%s

resource "azuread_application_password" "test" {
  application_object_id = azuread_application.test.id
  description           = "terraform"
  key_id                = "%s"
  value                 = "%s"
  end_date              = "2099-01-01T01:02:03Z"
}
`, testAccApplicationPassword_template(ri), keyId, value)
}

func testAccApplicationPassword_relativeEndDate(ri int, value string) string {
	return fmt.Sprintf(`
%s

resource "azuread_application_password" "test" {
  application_object_id = azuread_application.test.id
  value                 = "%s"
  end_date_relative     = "8760h"
}
`, testAccApplicationPassword_template(ri), value)
}

func testAccApplicationPassword_requiresImport(ri int, value string) string {
	template := testAccApplicationPassword_basic(ri, value)
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
