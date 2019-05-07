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

//Test with application
func TestAccAzureADObjectPasswordApplication_basic(t *testing.T) {
	resourceName := "azuread_object_password.test"
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
		CheckDestroy: testCheckADApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADObjectPasswordApplication_basic(applicationId, value),
				Check: resource.ComposeTestCheckFunc(
					// can't assert on Value since it's not returned
					testCheckADObjectPasswordApplicationExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "start_date"),
					resource.TestCheckResourceAttrSet(resourceName, "key_id"),
					resource.TestCheckResourceAttr(resourceName, "end_date", "2020-01-01T01:02:03Z"),
				),
			},
		},
	})
}

func TestAccAzureADObjectPasswordApplication_requiresImport(t *testing.T) {
	if !requireResourcesToBeImported {
		t.Skip("Skipping since resources aren't required to be imported")
		return
	}

	resourceName := "azuread_object_password.test"
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
		CheckDestroy: testCheckADApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADObjectPasswordApplication_basic(applicationId, value),
				Check: resource.ComposeTestCheckFunc(
					testCheckADObjectPasswordApplicationExists(resourceName),
				),
			},
			{
				Config:      testAccADObjectPasswordApplication_requiresImport(applicationId, value),
				ExpectError: testRequiresImportError("azuread_object_password"),
			},
		},
	})
}

func TestAccAzureADObjectPasswordApplication_customKeyId(t *testing.T) {
	resourceName := "azuread_object_password.test"
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
		CheckDestroy: testCheckADApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADApplicationPassword_customKeyId(applicationId, keyId, value),
				Check: resource.ComposeTestCheckFunc(
					// can't assert on Value since it's not returned
					testCheckADObjectPasswordApplicationExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "start_date"),
					resource.TestCheckResourceAttr(resourceName, "key_id", keyId),
					resource.TestCheckResourceAttr(resourceName, "end_date", "2020-01-01T01:02:03Z"),
				),
			},
		},
	})
}

func TestAccAzureADObjectPasswordApplication_relativeEndDate(t *testing.T) {
	resourceName := "azuread_application.test"
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
		CheckDestroy: testCheckADApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADApplicationPassword_relativeEndDate(applicationId, value),
				Check: resource.ComposeTestCheckFunc(
					// can't assert on Value since it's not returned
					testCheckADObjectPasswordApplicationExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "start_date"),
					resource.TestCheckResourceAttrSet(resourceName, "key_id"),
					resource.TestCheckResourceAttrSet(resourceName, "end_date"),
				),
			},
		},
	})
}

//Test with service principal
func TestAccAzureADObjectPasswordServicePrincipal_basic(t *testing.T) {
	resourceName := "azuread_object_password.test"
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
		CheckDestroy: testCheckADApplicationDestroy,
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

func TestAccAzureADObjectPasswordServicePrincipal_requiresImport(t *testing.T) {
	if !requireResourcesToBeImported {
		t.Skip("Skipping since resources aren't required to be imported")
		return
	}

	resourceName := "azuread_object_password.test"
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
				ExpectError: testRequiresImportError("azuread_object_password"),
			},
		},
	})
}

func TestAccAzureADObjectPasswordServicePrincipal_customKeyId(t *testing.T) {
	resourceName := "azuread_object_password.test"
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
					testCheckADObjectPasswordServicePrincipalExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "start_date"),
					resource.TestCheckResourceAttr(resourceName, "key_id", keyId),
					resource.TestCheckResourceAttr(resourceName, "end_date", "2020-01-01T01:02:03Z"),
				),
			},
		},
	})
}

func TestAccAzureADObjectPasswordServicePrincipal_relativeEndDate(t *testing.T) {
	resourceName := "azuread_application.test"
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

func testCheckADObjectPasswordApplicationExists(name string) resource.TestCheckFunc { //nolint unparam
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %q", name)
		}

		client := testAccProvider.Meta().(*ArmClient).applicationsClient
		ctx := testAccProvider.Meta().(*ArmClient).StopContext

		id := strings.Split(rs.Primary.ID, "/")
		applicationId := id[0]
		keyId := id[1]
		resp, err := client.Get(ctx, applicationId)

		if err != nil {
			if ar.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("Bad: Azure AD Application %q does not exist", applicationId)
			}
			return fmt.Errorf("Bad: Get on Azure AD applicationsClient: %+v", err)
		}

		credentials, err := client.ListPasswordCredentials(ctx, applicationId)
		if err != nil {
			return fmt.Errorf("Error Listing Password Credentials for Application %q: %+v", applicationId, err)
		}

		for _, credential := range *credentials.Value {
			if credential.KeyID == nil {
				continue
			}

			if *credential.KeyID == keyId {
				return nil
			}
		}

		return fmt.Errorf("Password Credential %q was not found in Application %q", keyId, applicationId)
	}
}

// Application templates
func testAccADApplicationPassword_template(applicationId string) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name = "acctestspa%s"
}
`, applicationId)
}

func testAccADObjectPasswordApplication_basic(applicationId, value string) string {
	return fmt.Sprintf(`
%s

resource "azuread_object_password" "test" {
  object_id            = "${azuread_application.test.id}"
  value                = "%s"
  end_date             = "2020-01-01T01:02:03Z"
}
`, testAccADApplicationPassword_template(applicationId), value)
}

func testAccADObjectPasswordApplication_requiresImport(applicationId, value string) string {
	template := testAccADObjectPasswordApplication_basic(applicationId, value)
	return fmt.Sprintf(`
%s

resource "azuread_object_password" "import" {
  key_id               = "${azuread_object_password.test.key_id}"
  object_id            = "${azuread_object_password.test.application_id}"
  value                = "${azuread_object_password.test.value}"
  end_date             = "${azuread_object_password.test.end_date}"
}
`, template)
}

func testAccADApplicationPassword_customKeyId(applicationId, keyId, value string) string {
	return fmt.Sprintf(`
%s

resource "azuread_object_password" "test" {
  object_id            = "${azuread_application.test.id}"
  key_id               = "%s"
  value                = "%s"
  end_date             = "2020-01-01T01:02:03Z"
}
`, testAccADApplicationPassword_template(applicationId), keyId, value)
}

func testAccADApplicationPassword_relativeEndDate(applicationId, value string) string {
	return fmt.Sprintf(`
%s

resource "azuread_object_password" "test" {
  object_id            = "${azuread_application.test.id}"
  value                = "%s"
  end_date_relative    = "8760h"
}
`, testAccADApplicationPassword_template(applicationId), value)
}

func testCheckADObjectPasswordServicePrincipalExists(name string) resource.TestCheckFunc { //nolint unparam
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %q", name)
		}

		client := testAccProvider.Meta().(*ArmClient).servicePrincipalsClient
		ctx := testAccProvider.Meta().(*ArmClient).StopContext

		id := strings.Split(rs.Primary.ID, "/")
		servicePrincipalId := id[0]
		keyId := id[1]
		resp, err := client.Get(ctx, servicePrincipalId)

		if err != nil {
			if ar.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("Bad: Azure AD Service Principal %q does not exist", servicePrincipalId)
			}
			return fmt.Errorf("Bad: Get on Azure AD servicePrincipalsClient: %+v", err)
		}

		credentials, err := client.ListPasswordCredentials(ctx, servicePrincipalId)
		if err != nil {
			return fmt.Errorf("Error Listing Password Credentials for Service Principal %q: %+v", servicePrincipalId, err)
		}

		for _, credential := range *credentials.Value {
			if credential.KeyID == nil {
				continue
			}

			if *credential.KeyID == keyId {
				return nil
			}
		}

		return fmt.Errorf("Password Credential %q was not found in Service Principal %q", keyId, servicePrincipalId)
	}
}

// Service Principal templates
func testAccADObjectPasswordServicePrincipal_template(applicationId string) string {
	return fmt.Sprintf(`
%s

resource "azuread_service_principal" "test" {
  application_id = "${azuread_application.test.application_id}"
}
`, testAccADApplicationPassword_template(applicationId))
}

func testAccADObjectPasswordServicePrincipal_basic(applicationId, value string) string {
	return fmt.Sprintf(`
%s

resource "azuread_object_password" "test" {
  object_id            = "${azuread_service_principal.test.id}"
  value                = "%s"
  end_date             = "2020-01-01T01:02:03Z"
}
`, testAccADObjectPasswordServicePrincipal_template(applicationId), value)
}

func testAccADObjectPasswordServicePrincipal_requiresImport(applicationId, value string) string {
	template := testAccADServicePrincipalPassword_basic(applicationId, value)
	return fmt.Sprintf(`
%s

resource "azuread_object_password" "import" {
  key_id               = "${azuread_object_password.test.key_id}"
  object_id            = "${azuread_object_password.test.service_principal_id}"
  value                = "${azuread_object_password.test.value}"
  end_date             = "${azuread_object_password.test.end_date}"
}
`, template)
}

func testAccADObjectPasswordServicePrincipal_customKeyId(applicationId, keyId, value string) string {
	return fmt.Sprintf(`
%s

resource "azuread_object_password" "test" {
  object_id            = "${azuread_service_principal.test.id}"
  key_id               = "%s"
  value                = "%s"
  end_date             = "2020-01-01T01:02:03Z"
}
`, testAccADObjectPasswordServicePrincipal_template(applicationId), keyId, value)
}

func testAccADObjectPasswordServicePrincipal_relativeEndDate(applicationId, value string) string {
	return fmt.Sprintf(`
%s

resource "azuread_object_password" "test" {
  object_id            = "${azuread_service_principal.test.id}"
  value                = "%s"
  end_date_relative    = "8760h"
}
`, testAccADObjectPasswordServicePrincipal_template(applicationId), value)
}
