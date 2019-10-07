package azuread

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

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
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("acctestApp-%s", id)),
					resource.TestCheckResourceAttr(resourceName, "homepage", fmt.Sprintf("https://acctestApp-%s", id)),
					resource.TestCheckResourceAttr(resourceName, "oauth2_allow_implicit_flow", "false"),
					resource.TestCheckResourceAttr(resourceName, "type", "webapp/api"),
					resource.TestCheckResourceAttr(resourceName, "oauth2_permissions.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "oauth2_permissions.0.admin_consent_description", fmt.Sprintf("Allow the application to access %s on behalf of the signed-in user.", fmt.Sprintf("acctestApp-%s", id))),
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

func TestAccAzureADApplication_http_homepage(t *testing.T) {
	resourceName := "azuread_application.test"
	id := uuid.New().String()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADApplication_http_homepage(id),
				Check: resource.ComposeTestCheckFunc(
					testCheckADApplicationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("acctestApp-%s", id)),
					resource.TestCheckResourceAttr(resourceName, "homepage", fmt.Sprintf("http://homepage-%s", id)),
					resource.TestCheckResourceAttr(resourceName, "oauth2_allow_implicit_flow", "false"),
					resource.TestCheckResourceAttr(resourceName, "type", "webapp/api"),
					resource.TestCheckResourceAttr(resourceName, "oauth2_permissions.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "oauth2_permissions.0.admin_consent_description", fmt.Sprintf("Allow the application to access %s on behalf of the signed-in user.", fmt.Sprintf("acctestApp-%s", id))),
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
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("acctestApp-%s", id)),
					resource.TestCheckResourceAttr(resourceName, "homepage", fmt.Sprintf("https://homepage-%s", id)),
					resource.TestCheckResourceAttr(resourceName, "oauth2_allow_implicit_flow", "true"),
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

func TestAccAzureADApplication_publicClient(t *testing.T) {
	resourceName := "azuread_application.test"
	id := uuid.New().String()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADApplication_publicClient(id),
				Check: resource.ComposeTestCheckFunc(
					testCheckADApplicationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "public_client", "true"),
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
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("acctestApp-%s", id)),
					resource.TestCheckResourceAttr(resourceName, "homepage", fmt.Sprintf("https://acctestApp-%s", id)),
					resource.TestCheckResourceAttr(resourceName, "identifier_uris.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "reply_urls.#", "0"),
				),
			},
			{
				Config: testAccADApplication_complete(updatedId),
				Check: resource.ComposeTestCheckFunc(
					testCheckADApplicationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("acctestApp-%s", updatedId)),
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

func TestAccAzureADApplication_appRoles(t *testing.T) {
	resourceName := "azuread_application.test"
	id := uuid.New().String()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADApplication_appRoles(id),
				Check: resource.ComposeTestCheckFunc(
					testCheckADApplicationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "app_role.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "app_role.3282540397.allowed_member_types.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "app_role.3282540397.allowed_member_types.2550101162", "Application"),
					resource.TestCheckResourceAttr(resourceName, "app_role.3282540397.allowed_member_types.2906997583", "User"),
					resource.TestCheckResourceAttr(resourceName, "app_role.3282540397.description", "Admins can manage roles and perform all task actions"),
					resource.TestCheckResourceAttr(resourceName, "app_role.3282540397.display_name", "Admin"),
					resource.TestCheckResourceAttr(resourceName, "app_role.3282540397.is_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "app_role.3282540397.value", "Admin"),
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

func TestAccAzureADApplication_appRolesUpdate(t *testing.T) {
	resourceName := "azuread_application.test"
	id := uuid.New().String()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADApplication_appRoles(id),
				Check: resource.ComposeTestCheckFunc(
					testCheckADApplicationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "app_role.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "app_role.3282540397.allowed_member_types.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "app_role.3282540397.allowed_member_types.2550101162", "Application"),
					resource.TestCheckResourceAttr(resourceName, "app_role.3282540397.allowed_member_types.2906997583", "User"),
					resource.TestCheckResourceAttr(resourceName, "app_role.3282540397.description", "Admins can manage roles and perform all task actions"),
					resource.TestCheckResourceAttr(resourceName, "app_role.3282540397.display_name", "Admin"),
					resource.TestCheckResourceAttr(resourceName, "app_role.3282540397.is_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "app_role.3282540397.value", "Admin"),
				),
			},
			{
				Config: testAccADApplication_appRolesUpdate(id),
				Check: resource.ComposeTestCheckFunc(
					testCheckADApplicationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "app_role.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "app_role.1786747921.allowed_member_types.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "app_role.1786747921.allowed_member_types.2906997583", "User"),
					resource.TestCheckResourceAttr(resourceName, "app_role.1786747921.description", "ReadOnly roles have limited query access"),
					resource.TestCheckResourceAttr(resourceName, "app_role.1786747921.display_name", "ReadOnly"),
					resource.TestCheckResourceAttr(resourceName, "app_role.1786747921.is_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "app_role.1786747921.value", "User"),
					resource.TestCheckResourceAttr(resourceName, "app_role.2608972077.allowed_member_types.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "app_role.2608972077.allowed_member_types.2906997583", "User"),
					resource.TestCheckResourceAttr(resourceName, "app_role.2608972077.description", "Admins can manage roles and perform all task actions"),
					resource.TestCheckResourceAttr(resourceName, "app_role.2608972077.display_name", "Admin"),
					resource.TestCheckResourceAttr(resourceName, "app_role.2608972077.is_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "app_role.2608972077.value", "Admin"),
				),
			},
		},
	})
}

func TestAccAzureADApplication_appRolesDelete(t *testing.T) {
	resourceName := "azuread_application.test"
	id := uuid.New().String()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADApplication_appRolesUpdate(id),
				Check: resource.ComposeTestCheckFunc(
					testCheckADApplicationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "app_role.#", "2"),
				),
			},
			{
				Config: testAccADApplication_appRoles(id),
				Check: resource.ComposeTestCheckFunc(
					testCheckADApplicationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "app_role.#", "1"),
				),
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
				Config: testAccADApplication_withGroupMembershipClaimsDirectoryRole(id),
				Check: resource.ComposeTestCheckFunc(
					testCheckADApplicationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "group_membership_claims", "DirectoryRole"),
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
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("acctestApp-%s", id)),
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
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("acctestApp-%s", id)),
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
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("acctestApp-%s", id)),
					resource.TestCheckResourceAttr(resourceName, "homepage", fmt.Sprintf("https://acctestApp-%s", id)),
					resource.TestCheckResourceAttr(resourceName, "type", "webapp/api"),
					resource.TestCheckResourceAttr(resourceName, "identifier_uris.#", "0"),
					resource.TestCheckResourceAttrSet(resourceName, "application_id"),
				),
			},
			{
				Config: testAccADApplication_native(id),
				Check: resource.ComposeTestCheckFunc(
					testCheckADApplicationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("acctestApp-%s", id)),
					resource.TestCheckResourceAttr(resourceName, "homepage", fmt.Sprintf("https://acctestApp-%s", id)),
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
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("acctestApp-%s", id)),
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
  name = "acctestApp-%s"
}
`, id)
}

func testAccADApplication_http_homepage(id string) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name      = "acctestApp-%s"
  homepage  = "http://homepage-%s"
}
`, id, id)
}

func testAccADApplication_publicClient(id string) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name          = "acctestApp-%s"
  type          = "native"
  public_client = true
}
`, id)
}

func testAccADApplication_availableToOtherTenants(id string) string {
	return fmt.Sprintf(`
data "azuread_domains" "tenant_domain" {
  only_initial = true
}

resource "azuread_application" "test" {
  name                       = "acctestApp-%s"
  identifier_uris            = ["https://%s.${data.azuread_domains.tenant_domain.domains.0.domain_name}"]
  available_to_other_tenants = true
}
`, id, id)
}

func testAccADApplication_withGroupMembershipClaimsDirectoryRole(id string) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name                    = "acctestApp-%s"
  group_membership_claims = "DirectoryRole"
}
`, id)
}

func testAccADApplication_withGroupMembershipClaimsSecurityGroup(id string) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name                    = "acctestApp-%s"
  group_membership_claims = "SecurityGroup"
}
`, id)
}

func testAccADApplication_complete(id string) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name                       = "acctestApp-%s"
  homepage                   = "https://homepage-%s"
  identifier_uris            = ["http://%s.hashicorptest.com/00000000-0000-0000-0000-00000000"]
  reply_urls                 = ["http://unittest.hashicorptest.com"]
  oauth2_allow_implicit_flow = true
  
  group_membership_claims    = "All"

  required_resource_access {
    resource_app_id = "00000003-0000-0000-c000-000000000000"

    resource_access {
      id   = "7ab1d382-f21e-4acd-a863-ba3e13f7da61"
      type = "Role"
    }

    resource_access {
      id   = "e1fe6dd8-ba31-4d61-89e7-88639da4683d"
      type = "Scope"
    }

    resource_access {
      id   = "06da0dbc-49e2-44d2-8312-53f166ab848a"
      type = "Scope"
    }
  }

  required_resource_access {
    resource_app_id = "00000002-0000-0000-c000-000000000000"

    resource_access {
      id   = "311a71cc-e848-46a1-bdf8-97ff7156d8e6"
      type = "Scope"
    }
  }
}
`, id, id, id)
}

func testAccADApplication_appRoles(id string) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name = "acctestApp-%s"

  app_role {
    allowed_member_types = [
      "User",
      "Application",
    ]

    description  = "Admins can manage roles and perform all task actions"
    display_name = "Admin"
    is_enabled   = true
    value        = "Admin"
  }
}
`, id)
}

func testAccADApplication_appRolesUpdate(id string) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name = "acctestApp-%s"

  app_role {
    allowed_member_types = ["User"]
    description          = "Admins can manage roles and perform all task actions"
    display_name         = "Admin"
    is_enabled           = true
    value                = "Admin"
  }

  app_role {
    allowed_member_types = ["User"]
    description          = "ReadOnly roles have limited query access"
    display_name         = "ReadOnly"
    is_enabled           = true
    value                = "User"
  }
}
`, id)
}

func testAccADApplication_native(id string) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name = "acctestApp-%s"
  type = "native"
}
`, id)
}

func testAccADApplication_nativeReplyUrls(id string) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name       = "acctestApp-%s"
  type       = "native"
  reply_urls = ["urn:ietf:wg:oauth:2.0:oob"]
}
`, id)
}

func testAccADApplication_native_app_does_not_allow_identifier_uris(id string) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name            = "acctestApp-%s"
  identifier_uris = ["http://%s.hashicorptest.com"]
  type            = "native"
}
`, id, id)
}
