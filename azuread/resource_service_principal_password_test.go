package azuread

import (
	"fmt"
	"strings"
	"testing"

	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/ar"

	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccAzureADServicePrincipalPassword_basic(t *testing.T) {
	resourceName := "azuread_service_principal_password.test"
	applicationId, err := uuid.GenerateUUID()
	if err != nil {
		t.Fatal(err)
	}
	value, err := uuid.GenerateUUID()
	if err != nil {
		t.Fatal(err)
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADServicePrincipalDestroy,
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
	applicationId, err := uuid.GenerateUUID()
	if err != nil {
		t.Fatal(err)
	}
	value, err := uuid.GenerateUUID()
	if err != nil {
		t.Fatal(err)
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADServicePrincipalDestroy,
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
	applicationId, err := uuid.GenerateUUID()
	if err != nil {
		t.Fatal(err)
	}
	keyId, err := uuid.GenerateUUID()
	if err != nil {
		t.Fatal(err)
	}
	value, err := uuid.GenerateUUID()
	if err != nil {
		t.Fatal(err)
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADServicePrincipalDestroy,
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
	applicationId, err := uuid.GenerateUUID()
	if err != nil {
		t.Fatal(err)
	}
	value, err := uuid.GenerateUUID()
	if err != nil {
		t.Fatal(err)
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADServicePrincipalDestroy,
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

func testCheckADServicePrincipalPasswordExists(name string) resource.TestCheckFunc { //nolint unparam
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %q", name)
		}

		client := testAccProvider.Meta().(*ArmClient).servicePrincipalsClient
		ctx := testAccProvider.Meta().(*ArmClient).StopContext

		id := strings.Split(rs.Primary.ID, "/")
		objectId := id[0]
		keyId := id[1]
		resp, err := client.Get(ctx, objectId)

		if err != nil {
			if ar.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("Bad: Azure AD Service Principal %q does not exist", objectId)
			}
			return fmt.Errorf("Bad: Get on Azure AD servicePrincipalsClient: %+v", err)
		}

		credentials, err := client.ListPasswordCredentials(ctx, objectId)
		if err != nil {
			return fmt.Errorf("Error Listing Password Credentials for Service Principal %q: %+v", objectId, err)
		}

		for _, credential := range *credentials.Value {
			if credential.KeyID == nil {
				continue
			}

			if *credential.KeyID == keyId {
				return nil
			}
		}

		return fmt.Errorf("Password Credential %q was not found in Service Principal %q", keyId, objectId)
	}
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
