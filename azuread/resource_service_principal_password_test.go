package azuread

import (
	"fmt"
	"testing"

	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/ar"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/graph"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func testCheckADServicePrincipalPasswordExists(name string) resource.TestCheckFunc { //nolint unparam
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*ArmClient).servicePrincipalsClient
		ctx := testAccProvider.Meta().(*ArmClient).StopContext

		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %q", name)
		}

		id, err := graph.ParsePasswordCredentialId(rs.Primary.ID)
		if err != nil {
			return fmt.Errorf("error Service Principal Password Credential ID: %v", err)
		}

		resp, err := client.Get(ctx, id.ObjectId)
		if err != nil {
			if ar.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("Bad: Azure AD Service Principal %q does not exist", id.ObjectId)
			}
			return fmt.Errorf("Bad: Get on Azure AD servicePrincipalsClient: %+v", err)
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

func testCheckADServicePrincipalPasswordCheckDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		client := testAccProvider.Meta().(*ArmClient).applicationsClient
		ctx := testAccProvider.Meta().(*ArmClient).StopContext

		if rs.Type != "azuread_service_principal_password" {
			continue
		}

		id, err := graph.ParsePasswordCredentialId(rs.Primary.ID)
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

func TestAccAzureADServicePrincipalPassword_basic(t *testing.T) {
	resourceName := "azuread_service_principal_password.test"
	applicationId := uuid.New().String()
	value := uuid.New().String()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADServicePrincipalPasswordCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADServicePrincipalPassword_basic(applicationId, value),
				Check: resource.ComposeTestCheckFunc(
					// can't assert on Value since it's not returned
					testCheckADServicePrincipalPasswordExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "start_date"),
					resource.TestCheckResourceAttrSet(resourceName, "key_id"),
					resource.TestCheckResourceAttr(resourceName, "end_date", "2020-01-01T01:02:03Z"),
				),
			},
		},
	})
}

func TestAccAzureADServicePrincipalPassword_requiresImport(t *testing.T) {
	if !requireResourcesToBeImported {
		t.Skip("Skipping since resources aren't required to be imported")
		return
	}

	resourceName := "azuread_service_principal_password.test"
	applicationId := uuid.New().String()
	value := uuid.New().String()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADServicePrincipalPasswordCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADServicePrincipalPassword_basic(applicationId, value),
				Check: resource.ComposeTestCheckFunc(
					testCheckADServicePrincipalPasswordExists(resourceName),
				),
			},
			{
				Config:      testAccADServicePrincipalPassword_requiresImport(applicationId, value),
				ExpectError: testRequiresImportError("azuread_service_principal_password"),
			},
		},
	})
}

func TestAccAzureADServicePrincipalPassword_customKeyId(t *testing.T) {
	resourceName := "azuread_service_principal_password.test"
	applicationId := uuid.New().String()
	keyId := uuid.New().String()
	value := uuid.New().String()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADServicePrincipalPasswordCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADServicePrincipalPassword_customKeyId(applicationId, keyId, value),
				Check: resource.ComposeTestCheckFunc(
					// can't assert on Value since it's not returned
					testCheckADServicePrincipalPasswordExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "start_date"),
					resource.TestCheckResourceAttr(resourceName, "key_id", keyId),
					resource.TestCheckResourceAttr(resourceName, "end_date", "2020-01-01T01:02:03Z"),
				),
			},
		},
	})
}

func TestAccAzureADServicePrincipalPassword_relativeEndDate(t *testing.T) {
	resourceName := "azuread_service_principal_password.test"
	applicationId := uuid.New().String()
	value := uuid.New().String()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADServicePrincipalPasswordCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADServicePrincipalPassword_relativeEndDate(applicationId, value),
				Check: resource.ComposeTestCheckFunc(
					// can't assert on Value since it's not returned
					testCheckADServicePrincipalPasswordExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "start_date"),
					resource.TestCheckResourceAttrSet(resourceName, "key_id"),
					resource.TestCheckResourceAttrSet(resourceName, "end_date"),
				),
			},
		},
	})
}

func testAccADServicePrincipalPassword_template(applicationId string) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name = "acctestspa%s"
}

resource "azuread_service_principal" "test" {
  application_id = "${azuread_application.test.application_id}"
}
`, applicationId)
}

func testAccADServicePrincipalPassword_basic(applicationId, value string) string {
	return fmt.Sprintf(`
%s

resource "azuread_service_principal_password" "test" {
  service_principal_id = "${azuread_service_principal.test.id}"
  value                = "%s"
  end_date             = "2020-01-01T01:02:03Z"
}
`, testAccADServicePrincipalPassword_template(applicationId), value)
}

func testAccADServicePrincipalPassword_requiresImport(applicationId, value string) string {
	template := testAccADServicePrincipalPassword_basic(applicationId, value)
	return fmt.Sprintf(`
%s

resource "azuread_service_principal_password" "import" {
  key_id               = "${azuread_service_principal_password.test.key_id}"
  service_principal_id = "${azuread_service_principal_password.test.service_principal_id}"
  value                = "${azuread_service_principal_password.test.value}"
  end_date             = "${azuread_service_principal_password.test.end_date}"
}
`, template)
}

func testAccADServicePrincipalPassword_customKeyId(applicationId, keyId, value string) string {
	return fmt.Sprintf(`
%s

resource "azuread_service_principal_password" "test" {
  service_principal_id = "${azuread_service_principal.test.id}"
  key_id               = "%s"
  value                = "%s"
  end_date             = "2020-01-01T01:02:03Z"
}
`, testAccADServicePrincipalPassword_template(applicationId), keyId, value)
}

func testAccADServicePrincipalPassword_relativeEndDate(applicationId, value string) string {
	return fmt.Sprintf(`
%s

resource "azuread_service_principal_password" "test" {
  service_principal_id = "${azuread_service_principal.test.id}"
  value                = "%s"
  end_date_relative    = "8760h"
}
`, testAccADServicePrincipalPassword_template(applicationId), value)
}
