package aadgraph_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/terraform-providers/terraform-provider-azuread/internal/acceptance"
)

func TestAccApplicationDataSource_byObjectId(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_application", "test")
	pw := "utils@$$wR2" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.ProviderFactories,
		CheckDestroy:      testCheckApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccApplicationDataSource_objectId(data.RandomInteger, pw),
				Check: resource.ComposeTestCheckFunc(
					testCheckApplicationExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "application_id"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "object_id"),
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
				),
			},
		},
	})
}

func TestAccApplicationDataSource_byApplicationId(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_application", "test")
	pw := "utils@$$wR2" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.ProviderFactories,
		CheckDestroy:      testCheckApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccApplicationDataSource_applicationId(data.RandomInteger, pw),
				Check: resource.ComposeTestCheckFunc(
					testCheckApplicationExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "application_id"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "object_id"),
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
				),
			},
		},
	})
}

func TestAccApplicationDataSource_byName(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_application", "test")
	pw := "utils@$$wR2" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.ProviderFactories,
		CheckDestroy:      testCheckApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccApplicationDataSource_name(data.RandomInteger, pw),
				Check: resource.ComposeTestCheckFunc(
					testCheckApplicationExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "application_id"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "object_id"),
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
				),
			},
		},
	})
}

func testAccApplicationDataSource_objectId(ri int, pw string) string {
	template := testAccApplication_complete(ri, pw)
	return fmt.Sprintf(`
%s

data "azuread_application" "test" {
  object_id = azuread_application.test.object_id
}
`, template)
}

func testAccApplicationDataSource_applicationId(ri int, pw string) string {
	template := testAccApplication_complete(ri, pw)
	return fmt.Sprintf(`
%s

data "azuread_application" "test" {
  application_id = azuread_application.test.application_id
}
`, template)
}

func testAccApplicationDataSource_name(ri int, pw string) string {
	template := testAccApplication_complete(ri, pw)
	return fmt.Sprintf(`
%s

data "azuread_application" "test" {
  name = azuread_application.test.name
}
`, template)
}
