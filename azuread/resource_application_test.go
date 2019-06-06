package azuread

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/ar"
)

func TestAccAzureADApplication_basic(t *testing.T) {
	resourceName := "azuread_application.test"
	id := uuid.New().String()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADApplication_basic(id),
				Check: resource.ComposeTestCheckFunc(
					testCheckADApplicationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("acctest%s", id)),
					resource.TestCheckResourceAttr(resourceName, "homepage", fmt.Sprintf("https://acctest%s", id)),
					resource.TestCheckResourceAttr(resourceName, "type", "webapp/api"),
					resource.TestCheckResourceAttr(resourceName, "oauth2_permissions.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "oauth2_permissions.0.admin_consent_description", fmt.Sprintf("Access %s", fmt.Sprintf("acctest%s", id))),
					resource.TestCheckResourceAttrSet(resourceName, "application_id"),
					resource.TestCheckResourceAttrSet(resourceName, "object_id"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAzureADApplication_complete(t *testing.T) {
	resourceName := "azuread_application.test"
	id := uuid.New().String()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADApplication_complete(id),
				Check: resource.ComposeTestCheckFunc(
					testCheckADApplicationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("acctest%s", id)),
					resource.TestCheckResourceAttr(resourceName, "homepage", fmt.Sprintf("https://homepage-%s", id)),
					resource.TestCheckResourceAttr(resourceName, "identifier_uris.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "identifier_uris.0", fmt.Sprintf("http://%s.hashicorptest.com/00000000-0000-0000-0000-00000000", id)),
					resource.TestCheckResourceAttr(resourceName, "reply_urls.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "group_membership_claims", "All"),
					resource.TestCheckResourceAttr(resourceName, "required_resource_access.#", "2"),
					resource.TestCheckResourceAttrSet(resourceName, "application_id"),
					resource.TestCheckResourceAttrSet(resourceName, "object_id"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAzureADApplication_update(t *testing.T) {
	resourceName := "azuread_application.test"
	id := uuid.New().String()
	updatedId := uuid.New().String()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADApplication_basic(id),
				Check: resource.ComposeTestCheckFunc(
					testCheckADApplicationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("acctest%s", id)),
					resource.TestCheckResourceAttr(resourceName, "homepage", fmt.Sprintf("https://acctest%s", id)),
					resource.TestCheckResourceAttr(resourceName, "identifier_uris.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "reply_urls.#", "0"),
				),
			},
			{
				Config: testAccADApplication_complete(updatedId),
				Check: resource.ComposeTestCheckFunc(
					testCheckADApplicationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("acctest%s", updatedId)),
					resource.TestCheckResourceAttr(resourceName, "homepage", fmt.Sprintf("https://homepage-%s", updatedId)),
					resource.TestCheckResourceAttr(resourceName, "identifier_uris.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "identifier_uris.0", fmt.Sprintf("http://%s.hashicorptest.com/00000000-0000-0000-0000-00000000", updatedId)),
					resource.TestCheckResourceAttr(resourceName, "reply_urls.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "reply_urls.3714513888", "http://unittest.hashicorptest.com"),
					resource.TestCheckResourceAttr(resourceName, "required_resource_access.#", "2"),
				),
			},
		},
	})
}

func TestAccAzureADApplication_availableToOtherTenants(t *testing.T) {
	resourceName := "azuread_application.test"
	id := uuid.New().String()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADApplication_availableToOtherTenants(id),
				Check: resource.ComposeTestCheckFunc(
					testCheckADApplicationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "available_to_other_tenants", "true"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAzureADApplication_groupMembershipClaimsUpdate(t *testing.T) {
	resourceName := "azuread_application.test"
	id := uuid.New().String()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADApplication_basic(id),
				Check: resource.ComposeTestCheckFunc(
					testCheckADApplicationExists(resourceName),
				),
			},
			{
				Config: testAccADApplication_withGroupMembershipClaimsAll(id),
				Check: resource.ComposeTestCheckFunc(
					testCheckADApplicationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "group_membership_claims", "All"),
				),
			},
			{
				Config: testAccADApplication_withGroupMembershipClaimsSecurityGroup(id),
				Check: resource.ComposeTestCheckFunc(
					testCheckADApplicationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "group_membership_claims", "SecurityGroup"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAzureADApplication_native(t *testing.T) {
	resourceName := "azuread_application.test"
	id := uuid.New().String()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADApplication_native(id),
				Check: resource.ComposeTestCheckFunc(
					testCheckADApplicationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("acctest%s", id)),
					resource.TestCheckResourceAttr(resourceName, "homepage", ""),
					resource.TestCheckResourceAttr(resourceName, "type", "native"),
					resource.TestCheckResourceAttr(resourceName, "identifier_uris.#", "0"),
					resource.TestCheckResourceAttrSet(resourceName, "application_id"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAzureADApplication_nativeReplyUrls(t *testing.T) {
	resourceName := "azuread_application.test"
	id := uuid.New().String()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADApplication_nativeReplyUrls(id),
				Check: resource.ComposeTestCheckFunc(
					testCheckADApplicationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("acctest%s", id)),
					resource.TestCheckResourceAttr(resourceName, "type", "native"),
					resource.TestCheckResourceAttr(resourceName, "reply_urls.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "reply_urls.3637476042", "urn:ietf:wg:oauth:2.0:oob"),
					resource.TestCheckResourceAttrSet(resourceName, "application_id"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAzureADApplication_nativeUpdate(t *testing.T) {
	resourceName := "azuread_application.test"
	id := uuid.New().String()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADApplication_basic(id),
				Check: resource.ComposeTestCheckFunc(
					testCheckADApplicationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("acctest%s", id)),
					resource.TestCheckResourceAttr(resourceName, "homepage", fmt.Sprintf("https://acctest%s", id)),
					resource.TestCheckResourceAttr(resourceName, "type", "webapp/api"),
					resource.TestCheckResourceAttr(resourceName, "identifier_uris.#", "0"),
					resource.TestCheckResourceAttrSet(resourceName, "application_id"),
				),
			},
			{
				Config: testAccADApplication_native(id),
				Check: resource.ComposeTestCheckFunc(
					testCheckADApplicationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("acctest%s", id)),
					resource.TestCheckResourceAttr(resourceName, "homepage", fmt.Sprintf("https://acctest%s", id)),
					resource.TestCheckResourceAttr(resourceName, "type", "native"),
					resource.TestCheckResourceAttr(resourceName, "identifier_uris.#", "0"),
					resource.TestCheckResourceAttrSet(resourceName, "application_id"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccADApplication_complete(id),
				Check: resource.ComposeTestCheckFunc(
					testCheckADApplicationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("acctest%s", id)),
					resource.TestCheckResourceAttr(resourceName, "homepage", fmt.Sprintf("https://homepage-%s", id)),
					resource.TestCheckResourceAttr(resourceName, "identifier_uris.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "identifier_uris.0", fmt.Sprintf("http://%s.hashicorptest.com/00000000-0000-0000-0000-00000000", id)),
					resource.TestCheckResourceAttr(resourceName, "reply_urls.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "required_resource_access.#", "2"),
					resource.TestCheckResourceAttrSet(resourceName, "application_id"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAzureADApplication_native_app_does_not_allow_identifier_uris(t *testing.T) {
	id := uuid.New().String()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccADApplication_native_app_does_not_allow_identifier_uris(id),
				ExpectError: regexp.MustCompile("identifier_uris is not required for a native application"),
			},
		},
	})
}

func testCheckADApplicationExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %q", name)
		}

		client := testAccProvider.Meta().(*ArmClient).applicationsClient
		ctx := testAccProvider.Meta().(*ArmClient).StopContext
		resp, err := client.Get(ctx, rs.Primary.ID)

		if err != nil {
			if ar.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("Bad: Azure AD Application %q does not exist", rs.Primary.ID)
			}
			return fmt.Errorf("Bad: Get on Azure AD applicationsClient: %+v", err)
		}

		return nil
	}
}

func testCheckADApplicationDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "azuread_application" {
			continue
		}

		client := testAccProvider.Meta().(*ArmClient).applicationsClient
		ctx := testAccProvider.Meta().(*ArmClient).StopContext
		resp, err := client.Get(ctx, rs.Primary.ID)

		if err != nil {
			if ar.ResponseWasNotFound(resp.Response) {
				return nil
			}

			return err
		}

		return fmt.Errorf("Azure AD Application still exists:\n%#v", resp)
	}

	return nil
}

func testAccADApplication_basic(id string) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name = "acctest%s"
}
`, id)
}

func testAccADApplication_availableToOtherTenants(id string) string {
	return fmt.Sprintf(`

data "azuread_domains" "tenant_domain" {
	only_initial = true
}

resource "azuread_application" "test" {
  name                       = "acctest%s"
  identifier_uris            = ["https://%s.${data.azuread_domains.tenant_domain.domains.0.domain_name}"]
  available_to_other_tenants = true
}
`, id, id)
}

func testAccADApplication_withGroupMembershipClaimsAll(id string) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name                       = "acctest%s"
  group_membership_claims    = "All"
}
`, id)
}

func testAccADApplication_withGroupMembershipClaimsSecurityGroup(id string) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name                       = "acctest%s"
  group_membership_claims    = "SecurityGroup"
}
`, id)
}

func testAccADApplication_complete(id string) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name                       = "acctest%s"
  homepage                   = "https://homepage-%s"
  identifier_uris            = ["http://%s.hashicorptest.com/00000000-0000-0000-0000-00000000"]
  reply_urls                 = ["http://unittest.hashicorptest.com"]
  oauth2_allow_implicit_flow = true
  group_membership_claims    = "All"

  required_resource_access {
    resource_app_id = "00000003-0000-0000-c000-000000000000"

    resource_access {
      id = "7ab1d382-f21e-4acd-a863-ba3e13f7da61"
      type = "Role"
    }

    resource_access {
      id = "e1fe6dd8-ba31-4d61-89e7-88639da4683d"
      type = "Scope"
    }

    resource_access {
      id = "06da0dbc-49e2-44d2-8312-53f166ab848a"
      type = "Scope"
    }
  }

  required_resource_access {
    resource_app_id = "00000002-0000-0000-c000-000000000000"

    resource_access {
      id = "311a71cc-e848-46a1-bdf8-97ff7156d8e6"
      type = "Scope"
    }
  }
}
`, id, id, id)
}

func testAccADApplication_native(id string) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name = "acctest%s"
  type = "native"
}
`, id)
}

func testAccADApplication_nativeReplyUrls(id string) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name       = "acctest%s"
  type       = "native"
  reply_urls = ["urn:ietf:wg:oauth:2.0:oob"]
}
`, id)
}

func testAccADApplication_native_app_does_not_allow_identifier_uris(id string) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name            = "acctest%s"
  identifier_uris = ["http://%s.hashicorptest.com"]
  type            = "native"
}
`, id, id)
}
