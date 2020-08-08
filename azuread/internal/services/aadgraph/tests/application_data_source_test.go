package tests

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-azuread/azuread/internal/acceptance"
)

func TestAccAzureADApplicationDataSource_byObjectId(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_application", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccApplication_basic(data.RandomInteger),
			},
			{
				Config: testAccAzureADApplicationDataSource_objectId(data.RandomInteger),
				Check: resource.ComposeTestCheckFunc(
					testCheckApplicationExists(data.ResourceName),
					resource.TestCheckResourceAttr(data.ResourceName, "name", fmt.Sprintf("acctest-APP-%d", data.RandomInteger)),
					resource.TestCheckResourceAttr(data.ResourceName, "homepage", fmt.Sprintf("https://acctest-APP-%d", data.RandomInteger)),
					resource.TestCheckResourceAttr(data.ResourceName, "identifier_uris.#", "0"),
					resource.TestCheckResourceAttr(data.ResourceName, "reply_urls.#", "0"),
					resource.TestCheckResourceAttr(data.ResourceName, "optional_claims.#", "0"),
					resource.TestCheckResourceAttr(data.ResourceName, "required_resource_access.#", "0"),
					resource.TestCheckResourceAttr(data.ResourceName, "type", "webapp/api"),
					resource.TestCheckResourceAttr(data.ResourceName, "app_roles.#", "0"),
					resource.TestCheckResourceAttr(data.ResourceName, "oauth2_allow_implicit_flow", "false"),
					resource.TestCheckResourceAttr(data.ResourceName, "oauth2_permissions.#", "1"),
					resource.TestCheckResourceAttr(data.ResourceName, "oauth2_permissions.0.admin_consent_description", fmt.Sprintf("Allow the application to access %s on behalf of the signed-in user.", fmt.Sprintf("acctest-APP-%d", data.RandomInteger))),
					resource.TestCheckResourceAttrSet(data.ResourceName, "application_id"),
				),
			},
		},
	})
}

func TestAccAzureADApplicationDataSource_byObjectIdComplete(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_application", "test")
	pw := "p@$$wR2" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccApplication_complete(data.RandomInteger, pw),
			},
			{
				Config: testAccAzureADApplicationDataSource_objectIdComplete(data.RandomInteger, pw),
				Check: resource.ComposeTestCheckFunc(
					testCheckApplicationExists(data.ResourceName),
					resource.TestCheckResourceAttr(data.ResourceName, "name", fmt.Sprintf("acctest-APP-%d", data.RandomInteger)),
					resource.TestCheckResourceAttr(data.ResourceName, "homepage", fmt.Sprintf("https://homepage-%d", data.RandomInteger)),
					resource.TestCheckResourceAttr(data.ResourceName, "identifier_uris.#", "1"),
					resource.TestCheckResourceAttr(data.ResourceName, "reply_urls.#", "1"),
					resource.TestCheckResourceAttr(data.ResourceName, "oauth2_allow_implicit_flow", "true"),
					resource.TestCheckResourceAttr(data.ResourceName, "optional_claims.#", "1"),
					resource.TestCheckResourceAttr(data.ResourceName, "optional_claims.0.access_token.#", "2"),
					resource.TestCheckResourceAttr(data.ResourceName, "optional_claims.0.id_token.#", "1"),
					resource.TestCheckResourceAttr(data.ResourceName, "required_resource_access.#", "2"),
					resource.TestCheckResourceAttr(data.ResourceName, "group_membership_claims", "All"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "application_id"),
				),
			},
		},
	})
}

func TestAccAzureADApplicationDataSource_byApplicationId(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_application", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccApplication_basic(data.RandomInteger),
			},
			{
				Config: testAccAzureADApplicationDataSource_applicationId(data.RandomInteger),
				Check: resource.ComposeTestCheckFunc(
					testCheckApplicationExists(data.ResourceName),
					resource.TestCheckResourceAttr(data.ResourceName, "name", fmt.Sprintf("acctest-APP-%d", data.RandomInteger)),
					resource.TestCheckResourceAttr(data.ResourceName, "homepage", fmt.Sprintf("https://acctest-APP-%d", data.RandomInteger)),
					resource.TestCheckResourceAttr(data.ResourceName, "identifier_uris.#", "0"),
					resource.TestCheckResourceAttr(data.ResourceName, "reply_urls.#", "0"),
					resource.TestCheckResourceAttr(data.ResourceName, "optional_claims.#", "0"),
					resource.TestCheckResourceAttr(data.ResourceName, "required_resource_access.#", "0"),
					resource.TestCheckResourceAttr(data.ResourceName, "oauth2_allow_implicit_flow", "false"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "application_id"),
				),
			},
		},
	})
}

func TestAccAzureADApplicationDataSource_byName(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_application", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccApplication_basic(data.RandomInteger),
			},
			{
				Config: testAccAzureADApplicationDataSource_name(data.RandomInteger),
				Check: resource.ComposeTestCheckFunc(
					testCheckApplicationExists(data.ResourceName),
					resource.TestCheckResourceAttr(data.ResourceName, "name", fmt.Sprintf("acctest-APP-%d", data.RandomInteger)),
					resource.TestCheckResourceAttr(data.ResourceName, "homepage", fmt.Sprintf("https://acctest-APP-%d", data.RandomInteger)),
					resource.TestCheckResourceAttr(data.ResourceName, "identifier_uris.#", "0"),
					resource.TestCheckResourceAttr(data.ResourceName, "reply_urls.#", "0"),
					resource.TestCheckResourceAttr(data.ResourceName, "optional_claims.#", "0"),
					resource.TestCheckResourceAttr(data.ResourceName, "required_resource_access.#", "0"),
					resource.TestCheckResourceAttr(data.ResourceName, "oauth2_allow_implicit_flow", "false"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "application_id"),
				),
			},
		},
	})
}

func testAccAzureADApplicationDataSource_objectId(ri int) string {
	template := testAccApplication_basic(ri)
	return fmt.Sprintf(`
%s

data "azuread_application" "test" {
  object_id = azuread_application.test.object_id
}
`, template)
}

func testAccAzureADApplicationDataSource_objectIdComplete(ri int, pw string) string {
	template := testAccApplication_complete(ri, pw)
	return fmt.Sprintf(`
%s

data "azuread_application" "test" {
  object_id = azuread_application.test.object_id
}
`, template)
}

func testAccAzureADApplicationDataSource_applicationId(ri int) string {
	template := testAccApplication_basic(ri)
	return fmt.Sprintf(`
%s

data "azuread_application" "test" {
  application_id = azuread_application.test.application_id
}
`, template)
}

func testAccAzureADApplicationDataSource_name(ri int) string {
	template := testAccApplication_basic(ri)
	return fmt.Sprintf(`
%s

data "azuread_application" "test" {
  name = azuread_application.test.name
}
`, template)
}
