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

func testCheckServicePrincipalPasswordExists(name string) resource.TestCheckFunc { //nolint unparam
	return func(s *terraform.State) error {
		client := acceptance.AzureADProvider.Meta().(*clients.AadClient).AadGraph.ServicePrincipalsClient
		ctx := acceptance.AzureADProvider.Meta().(*clients.AadClient).StopContext

		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %q", name)
		}

		id, err := graph.ParsePasswordId(rs.Primary.ID)
		if err != nil {
			return fmt.Errorf("Service Principal Password Credential ID: %v", err)
		}

		resp, err := client.Get(ctx, id.ObjectId)
		if err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("Bad: Service Principal %q does not exist", id.ObjectId)
			}
			return fmt.Errorf("Bad: Get on ServicePrincipalsClient: %+v", err)
		}

		credentials, err := client.ListPasswordCredentials(ctx, id.ObjectId)
		if err != nil {
			return fmt.Errorf("listing Password Credentials for Service Principal %q: %+v", id.ObjectId, err)
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

		id, err := graph.ParsePasswordId(rs.Primary.ID)
		if err != nil {
			return fmt.Errorf("parsing Service Principal Password Credential ID: %v", err)
		}

		resp, err := client.Get(ctx, id.ObjectId)
		if err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
				return nil
			}

			return err
		}

		return fmt.Errorf("Service Principal Password Credential still exists:\n%#v", resp)
	}

	return nil
}

func TestAccServicePrincipalPassword_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_service_principal_password", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckServicePrincipalPasswordCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccServicePrincipalPassword_basic(data.RandomInteger, data.RandomPassword),
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

func TestAccServicePrincipalPassword_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_service_principal_password", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckServicePrincipalPasswordCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccServicePrincipalPassword_complete(data.RandomInteger, data.RandomID, data.RandomPassword),
				Check: resource.ComposeTestCheckFunc(
					testCheckServicePrincipalPasswordExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "start_date"),
					resource.TestCheckResourceAttr(data.ResourceName, "key_id", data.RandomID),
					resource.TestCheckResourceAttr(data.ResourceName, "end_date", "2099-01-01T01:02:03Z"),
				),
			},
			data.ImportStep("value"),
		},
	})
}

func TestAccServicePrincipalPassword_relativeEndDate(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_service_principal_password", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckServicePrincipalPasswordCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccServicePrincipalPassword_relativeEndDate(data.RandomInteger, data.RandomPassword),
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

func TestAccServicePrincipalPassword_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_service_principal_password", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckServicePrincipalPasswordCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccServicePrincipalPassword_basic(data.RandomInteger, data.RandomPassword),
				Check: resource.ComposeTestCheckFunc(
					testCheckServicePrincipalPasswordExists(data.ResourceName),
				),
			},
			data.RequiresImportErrorStep(testAccServicePrincipalPassword_requiresImport(data.RandomInteger, data.RandomPassword)),
		},
	})
}

func testAccServicePrincipalPassword_template(ri int) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name = "acctestApp-%d"
}

resource "azuread_service_principal" "test" {
  application_id = azuread_application.test.application_id
}
`, ri)
}

func testAccServicePrincipalPassword_basic(ri int, value string) string {
	return fmt.Sprintf(`
%s

resource "azuread_service_principal_password" "test" {
  service_principal_id = azuread_service_principal.test.id
  value                = "%s"
  end_date             = "2099-01-01T01:02:03Z"
}
`, testAccServicePrincipalPassword_template(ri), value)
}

func testAccServicePrincipalPassword_complete(ri int, keyId, value string) string {
	return fmt.Sprintf(`
%s

resource "azuread_service_principal_password" "test" {
  service_principal_id = azuread_service_principal.test.id
  description          = "terraform"
  key_id               = "%s"
  value                = "%s"
  end_date             = "2099-01-01T01:02:03Z"
}
`, testAccServicePrincipalPassword_template(ri), keyId, value)
}

func testAccServicePrincipalPassword_relativeEndDate(ri int, value string) string {
	return fmt.Sprintf(`
%s

resource "azuread_service_principal_password" "test" {
  service_principal_id = azuread_service_principal.test.id
  value                = "%s"
  end_date_relative    = "8760h"
}
`, testAccServicePrincipalPassword_template(ri), value)
}

func testAccServicePrincipalPassword_requiresImport(ri int, value string) string {
	return fmt.Sprintf(`
%s

resource "azuread_service_principal_password" "import" {
  key_id               = azuread_service_principal_password.test.key_id
  service_principal_id = azuread_service_principal_password.test.service_principal_id
  value                = azuread_service_principal_password.test.value
  end_date             = azuread_service_principal_password.test.end_date
}
`, testAccServicePrincipalPassword_basic(ri, value))
}
