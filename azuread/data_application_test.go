package azuread

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccAzureADApplicationDataSource_byObjectId(t *testing.T) {
	dataSourceName := "data.azuread_application.test"
	id := uuid.New().String()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADApplication_basic(id),
			},
			{
				Config: testAccAzureADApplicationDataSource_objectId(id),
				Check: resource.ComposeTestCheckFunc(
					testCheckADApplicationExists(dataSourceName),
					resource.TestCheckResourceAttr(dataSourceName, "name", fmt.Sprintf("acctest%s", id)),
					resource.TestCheckResourceAttr(dataSourceName, "homepage", fmt.Sprintf("https://acctest%s", id)),
					resource.TestCheckResourceAttr(dataSourceName, "identifier_uris.#", "0"),
					resource.TestCheckResourceAttr(dataSourceName, "reply_urls.#", "0"),
					resource.TestCheckResourceAttr(dataSourceName, "required_resource_access.#", "0"),
					resource.TestCheckResourceAttr(dataSourceName, "type", "webapp/api"),
					resource.TestCheckResourceAttr(dataSourceName, "oauth2_allow_implicit_flow", "false"),
					resource.TestCheckResourceAttr(dataSourceName, "oauth2_permissions.#", "1"),
					resource.TestCheckResourceAttr(dataSourceName, "oauth2_permissions.0.admin_consent_description", fmt.Sprintf("Access %s", fmt.Sprintf("acctest%s", id))),
					resource.TestCheckResourceAttrSet(dataSourceName, "application_id"),
				),
			},
		},
	})
}

func TestAccAzureADApplicationDataSource_byObjectIdComplete(t *testing.T) {
	dataSourceName := "data.azuread_application.test"
	id := uuid.New().String()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADApplication_complete(id),
			},
			{
				Config: testAccAzureADApplicationDataSource_objectIdComplete(id),
				Check: resource.ComposeTestCheckFunc(
					testCheckADApplicationExists(dataSourceName),
					resource.TestCheckResourceAttr(dataSourceName, "name", fmt.Sprintf("acctest%s", id)),
					resource.TestCheckResourceAttr(dataSourceName, "homepage", fmt.Sprintf("https://homepage-%s", id)),
					resource.TestCheckResourceAttr(dataSourceName, "identifier_uris.#", "1"),
					resource.TestCheckResourceAttr(dataSourceName, "reply_urls.#", "1"),
					resource.TestCheckResourceAttr(dataSourceName, "oauth2_allow_implicit_flow", "true"),
					resource.TestCheckResourceAttr(dataSourceName, "required_resource_access.#", "2"),
					resource.TestCheckResourceAttr(dataSourceName, "group_membership_claims", "All"),
					resource.TestCheckResourceAttrSet(dataSourceName, "application_id"),
				),
			},
		},
	})
}

func TestAccAzureADApplicationDataSource_byName(t *testing.T) {
	dataSourceName := "data.azuread_application.test"
	id := uuid.New().String()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADApplication_basic(id),
			},
			{
				Config: testAccAzureADApplicationDataSource_name(id),
				Check: resource.ComposeTestCheckFunc(
					testCheckADApplicationExists(dataSourceName),
					resource.TestCheckResourceAttr(dataSourceName, "name", fmt.Sprintf("acctest%s", id)),
					resource.TestCheckResourceAttr(dataSourceName, "homepage", fmt.Sprintf("https://acctest%s", id)),
					resource.TestCheckResourceAttr(dataSourceName, "identifier_uris.#", "0"),
					resource.TestCheckResourceAttr(dataSourceName, "reply_urls.#", "0"),
					resource.TestCheckResourceAttr(dataSourceName, "required_resource_access.#", "0"),
					resource.TestCheckResourceAttr(dataSourceName, "oauth2_allow_implicit_flow", "false"),
					resource.TestCheckResourceAttrSet(dataSourceName, "application_id"),
				),
			},
		},
	})
}

func testAccAzureADApplicationDataSource_objectId(id string) string {
	template := testAccADApplication_basic(id)
	return fmt.Sprintf(`
%s

data "azuread_application" "test" {
  object_id = "${azuread_application.test.object_id}"
}
`, template)
}

func testAccAzureADApplicationDataSource_objectIdComplete(id string) string {
	template := testAccADApplication_complete(id)
	return fmt.Sprintf(`
%s

data "azuread_application" "test" {
  object_id = "${azuread_application.test.object_id}"
}
`, template)
}

func testAccAzureADApplicationDataSource_name(id string) string {
	template := testAccADApplication_basic(id)
	return fmt.Sprintf(`
%s

data "azuread_application" "test" {
  name = "${azuread_application.test.name}"
}
`, template)
}
