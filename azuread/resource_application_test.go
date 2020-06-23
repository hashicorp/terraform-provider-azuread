package azuread

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/ar"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/tf"
)

func TestAccAzureADApplication_basic(t *testing.T) {
	resourceName := "azuread_application.test"
	ri := tf.AccRandTimeInt()
	appName := fmt.Sprintf("acctest-APP-%d", ri)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADApplication_basic(ri),
				Check: resource.ComposeTestCheckFunc(
					testCheckADApplicationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", appName),
					resource.TestCheckResourceAttr(resourceName, "homepage", fmt.Sprintf("https://%s", appName)),
					resource.TestCheckResourceAttr(resourceName, "oauth2_allow_implicit_flow", "false"),
					resource.TestCheckResourceAttr(resourceName, "type", "webapp/api"),
					resource.TestCheckResourceAttr(resourceName, "oauth2_permissions.#", "1"),
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
	ri := tf.AccRandTimeInt()
	pw := "p@$$wR2" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADApplication_complete(ri, pw),
				Check: resource.ComposeTestCheckFunc(
					testCheckADApplicationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("acctest-APP-%[1]d", ri)),
					resource.TestCheckResourceAttr(resourceName, "homepage", fmt.Sprintf("https://homepage-%d", ri)),
					resource.TestCheckResourceAttr(resourceName, "oauth2_allow_implicit_flow", "true"),
					resource.TestCheckResourceAttr(resourceName, "identifier_uris.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "identifier_uris.0", fmt.Sprintf("http://%d.hashicorptest.com/00000000-0000-0000-0000-00000000", ri)),
					resource.TestCheckResourceAttr(resourceName, "reply_urls.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "group_membership_claims", "All"),
					resource.TestCheckResourceAttr(resourceName, "optional_claims.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "optional_claims.0.access_token.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "optional_claims.0.id_token.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "required_resource_access.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "oauth2_permissions.#", "2"),
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
	ri := tf.AccRandTimeInt()
	updatedri := tf.AccRandTimeInt()
	pw := "p@$$wR2" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADApplication_basic(ri),
				Check: resource.ComposeTestCheckFunc(
					testCheckADApplicationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("acctest-APP-%[1]d", ri)),
					resource.TestCheckResourceAttr(resourceName, "homepage", fmt.Sprintf("https://acctest-APP-%d", ri)),
					resource.TestCheckResourceAttr(resourceName, "identifier_uris.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "reply_urls.#", "0"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccADApplication_complete(updatedri, pw),
				Check: resource.ComposeTestCheckFunc(
					testCheckADApplicationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("acctest-APP-%[1]d", updatedri)),
					resource.TestCheckResourceAttr(resourceName, "homepage", fmt.Sprintf("https://homepage-%d", updatedri)),
					resource.TestCheckResourceAttr(resourceName, "identifier_uris.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "identifier_uris.0", fmt.Sprintf("http://%d.hashicorptest.com/00000000-0000-0000-0000-00000000", updatedri)),
					resource.TestCheckResourceAttr(resourceName, "reply_urls.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "reply_urls.3714513888", "http://unittest.hashicorptest.com"),
					resource.TestCheckResourceAttr(resourceName, "optional_claims.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "optional_claims.0.access_token.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "optional_claims.0.id_token.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "required_resource_access.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "oauth2_permissions.#", "2"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccADApplication_basicEmpty(ri),
				Check: resource.ComposeTestCheckFunc(
					testCheckADApplicationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("acctest-APP-%[1]d", ri)),
					resource.TestCheckResourceAttr(resourceName, "identifier_uris.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "reply_urls.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "optional_claims.#", "0"),
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
	ri := tf.AccRandTimeInt()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADApplication_http_homepage(ri),
				Check: resource.ComposeTestCheckFunc(
					testCheckADApplicationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("acctest-APP-%[1]d", ri)),
					resource.TestCheckResourceAttr(resourceName, "homepage", fmt.Sprintf("http://homepage-%d", ri)),
					resource.TestCheckResourceAttr(resourceName, "oauth2_allow_implicit_flow", "false"),
					resource.TestCheckResourceAttr(resourceName, "type", "webapp/api"),
					resource.TestCheckResourceAttr(resourceName, "oauth2_permissions.#", "1"),
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
	ri := tf.AccRandTimeInt()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADApplication_publicClient(ri),
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

func TestAccAzureADApplication_availableToOtherTenants(t *testing.T) {
	resourceName := "azuread_application.test"
	ri := tf.AccRandTimeInt()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADApplication_availableToOtherTenants(ri),
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
	ri := tf.AccRandTimeInt()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADApplication_appRoles(ri),
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

func TestAccAzureADApplication_appRolesNoValue(t *testing.T) {
	rn := "azuread_application.test"
	ri := tf.AccRandTimeInt()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADApplication_appRolesNoValue(ri),
				Check: resource.ComposeTestCheckFunc(
					testCheckADApplicationExists(rn),
					resource.TestCheckResourceAttr(rn, "app_role.#", "1"),
				),
			},
			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAzureADApplication_appRolesUpdate(t *testing.T) {
	rn := "azuread_application.test"
	ri := tf.AccRandTimeInt()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADApplication_appRoles(ri),
				Check: resource.ComposeTestCheckFunc(
					testCheckADApplicationExists(rn),
					resource.TestCheckResourceAttr(rn, "app_role.#", "1"),
				),
			},
			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccADApplication_appRolesUpdate(ri),
				Check: resource.ComposeTestCheckFunc(
					testCheckADApplicationExists(rn),
					resource.TestCheckResourceAttr(rn, "app_role.#", "2"),
				),
			},
			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAzureADApplication_appRolesDelete(t *testing.T) {
	rn := "azuread_application.test"
	ri := tf.AccRandTimeInt()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADApplication_appRolesUpdate(ri),
				Check: resource.ComposeTestCheckFunc(
					testCheckADApplicationExists(rn),
					resource.TestCheckResourceAttr(rn, "app_role.#", "2"),
				),
			},
			{
				Config: testAccADApplication_appRoles(ri),
				Check: resource.ComposeTestCheckFunc(
					testCheckADApplicationExists(rn),
					resource.TestCheckResourceAttr(rn, "app_role.#", "1"),
				),
			},
		},
	})
}

func TestAccAzureADApplication_groupMembershipClaimsUpdate(t *testing.T) {
	rn := "azuread_application.test"
	ri := tf.AccRandTimeInt()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADApplication_basic(ri),
				Check: resource.ComposeTestCheckFunc(
					testCheckADApplicationExists(rn),
				),
			},
			{
				Config: testAccADApplication_withGroupMembershipClaimsDirectoryRole(ri),
				Check: resource.ComposeTestCheckFunc(
					testCheckADApplicationExists(rn),
					resource.TestCheckResourceAttr(rn, "group_membership_claims", "DirectoryRole"),
				),
			},
			{
				Config: testAccADApplication_withGroupMembershipClaimsSecurityGroup(ri),
				Check: resource.ComposeTestCheckFunc(
					testCheckADApplicationExists(rn),
					resource.TestCheckResourceAttr(rn, "group_membership_claims", "SecurityGroup"),
				),
			},
			{
				Config: testAccADApplication_withGroupMembershipClaimsApplicationGroup(ri),
				Check: resource.ComposeTestCheckFunc(
					testCheckADApplicationExists(rn),
					resource.TestCheckResourceAttr(rn, "group_membership_claims", "ApplicationGroup"),
				),
			},
			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAzureADApplication_native(t *testing.T) {
	rn := "azuread_application.test"
	ri := tf.AccRandTimeInt()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADApplication_native(ri),
				Check: resource.ComposeTestCheckFunc(
					testCheckADApplicationExists(rn),
					resource.TestCheckResourceAttr(rn, "name", fmt.Sprintf("acctest-APP-%[1]d", ri)),
					resource.TestCheckResourceAttr(rn, "homepage", ""),
					resource.TestCheckResourceAttr(rn, "type", "native"),
					resource.TestCheckResourceAttr(rn, "identifier_uris.#", "0"),
					resource.TestCheckResourceAttrSet(rn, "application_id"),
				),
			},
			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAzureADApplication_nativeReplyUrls(t *testing.T) {
	rn := "azuread_application.test"
	ri := tf.AccRandTimeInt()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADApplication_nativeReplyUrls(ri),
				Check: resource.ComposeTestCheckFunc(
					testCheckADApplicationExists(rn),
					resource.TestCheckResourceAttr(rn, "name", fmt.Sprintf("acctest-APP-%[1]d", ri)),
					resource.TestCheckResourceAttr(rn, "type", "native"),
					resource.TestCheckResourceAttr(rn, "reply_urls.#", "1"),
					resource.TestCheckResourceAttr(rn, "reply_urls.3637476042", "urn:ietf:wg:oauth:2.0:oob"),
					resource.TestCheckResourceAttrSet(rn, "application_id"),
				),
			},
			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAzureADApplication_nativeUpdate(t *testing.T) {
	rn := "azuread_application.test"
	ri := tf.AccRandTimeInt()
	pw := "p@$$wR2" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADApplication_basic(ri),
				Check: resource.ComposeTestCheckFunc(
					testCheckADApplicationExists(rn),
					resource.TestCheckResourceAttr(rn, "name", fmt.Sprintf("acctest-APP-%[1]d", ri)),
					resource.TestCheckResourceAttr(rn, "homepage", fmt.Sprintf("https://acctest-APP-%d", ri)),
					resource.TestCheckResourceAttr(rn, "type", "webapp/api"),
					resource.TestCheckResourceAttr(rn, "identifier_uris.#", "0"),
					resource.TestCheckResourceAttrSet(rn, "application_id"),
				),
			},
			{
				Config: testAccADApplication_native(ri),
				Check: resource.ComposeTestCheckFunc(
					testCheckADApplicationExists(rn),
					resource.TestCheckResourceAttr(rn, "name", fmt.Sprintf("acctest-APP-%[1]d", ri)),
					resource.TestCheckResourceAttr(rn, "homepage", fmt.Sprintf("https://acctest-APP-%d", ri)),
					resource.TestCheckResourceAttr(rn, "type", "native"),
					resource.TestCheckResourceAttr(rn, "identifier_uris.#", "0"),
					resource.TestCheckResourceAttrSet(rn, "application_id"),
				),
			},
			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccADApplication_complete(ri, pw),
				Check: resource.ComposeTestCheckFunc(
					testCheckADApplicationExists(rn),
					resource.TestCheckResourceAttr(rn, "name", fmt.Sprintf("acctest-APP-%[1]d", ri)),
					resource.TestCheckResourceAttr(rn, "homepage", fmt.Sprintf("https://homepage-%d", ri)),
					resource.TestCheckResourceAttr(rn, "identifier_uris.#", "1"),
					resource.TestCheckResourceAttr(rn, "identifier_uris.0", fmt.Sprintf("http://%d.hashicorptest.com/00000000-0000-0000-0000-00000000", ri)),
					resource.TestCheckResourceAttr(rn, "reply_urls.#", "1"),
					resource.TestCheckResourceAttr(rn, "required_resource_access.#", "2"),
					resource.TestCheckResourceAttrSet(rn, "application_id"),
				),
			},
			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAzureADApplication_native_app_does_not_allow_identifier_uris(t *testing.T) {
	ri := tf.AccRandTimeInt()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccADApplication_native_app_does_not_allow_identifier_uris(ri),
				ExpectError: regexp.MustCompile("identifier_uris is not required for a native application"),
			},
		},
	})
}

func TestAccAzureADApplication_oauth2PermissionsUpdate(t *testing.T) {
	resourceName := "azuread_application.test"
	ri := tf.AccRandTimeInt()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccADApplication_basic(ri),
				Check: resource.ComposeTestCheckFunc(
					testCheckADApplicationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("acctest-APP-%[1]d", ri)),
					//resource.TestCheckResourceAttr(resourceName, "identifier_uris.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "oauth2_permissions.#", "1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccADApplication_oauth2Permissions(ri),
				Check: resource.ComposeTestCheckFunc(
					testCheckADApplicationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("acctest-APP-%[1]d", ri)),
					resource.TestCheckResourceAttr(resourceName, "oauth2_permissions.#", "2"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccADApplication_oauth2PermissionsEmpty(ri),
				Check: resource.ComposeTestCheckFunc(
					testCheckADApplicationExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("acctest-APP-%[1]d", ri)),
					resource.TestCheckResourceAttr(resourceName, "oauth2_permissions.#", "0"),
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

func TestAccAzureADApplication_preventDuplicateNames(t *testing.T) {
	ri := tf.AccRandTimeInt()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckADApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccADApplication_duplicateName(ri),
				ExpectError: regexp.MustCompile("existing Application .+ was found"),
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

func testAccADApplication_basic(ri int) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name = "acctest-APP-%[1]d"
}
`, ri)
}

func testAccADApplication_basicEmpty(ri int) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name                    = "acctest-APP-%[1]d"
  identifier_uris         = []
  reply_urls              = []
  group_membership_claims = "None"
}
`, ri)
}

func testAccADApplication_http_homepage(ri int) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name     = "acctest-APP-%[1]d"
  homepage = "http://homepage-%[1]d"
}
`, ri)
}

func testAccADApplication_publicClient(ri int) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name          = "acctest-APP-%[1]d"
  type          = "native"
  public_client = true
}
`, ri)
}

func testAccADApplication_availableToOtherTenants(ri int) string {
	return fmt.Sprintf(`
data "azuread_domains" "tenant_domain" {
  only_initial = true
}

resource "azuread_application" "test" {
  name                       = "acctest-APP-%[1]d"
  identifier_uris            = ["https://%[1]d.${data.azuread_domains.tenant_domain.domains.0.domain_name}"]
  available_to_other_tenants = true
}
`, ri)
}

func testAccADApplication_withGroupMembershipClaimsDirectoryRole(ri int) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name                    = "acctest-APP-%[1]d"
  group_membership_claims = "DirectoryRole"
}
`, ri)
}

func testAccADApplication_withGroupMembershipClaimsSecurityGroup(ri int) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name                    = "acctest-APP-%[1]d"
  group_membership_claims = "SecurityGroup"
}
`, ri)
}

func testAccADApplication_withGroupMembershipClaimsApplicationGroup(ri int) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name                    = "acctest-APP-%[1]d"
  group_membership_claims = "ApplicationGroup"
}
`, ri)
}

func testAccADApplication_complete(ri int, pw string) string {
	return fmt.Sprintf(`
%[1]s

resource "azuread_application" "test" {
  name                       = "acctest-APP-%[2]d"
  homepage                   = "https://homepage-%[2]d"
  identifier_uris            = ["http://%[2]d.hashicorptest.com/00000000-0000-0000-0000-00000000"]
  reply_urls                 = ["http://unittest.hashicorptest.com"]
  logout_url                 = "http://log.me.out"
  oauth2_allow_implicit_flow = true

  group_membership_claims = "All"

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

  oauth2_permissions {
    admin_consent_description  = "Administer the application"
    admin_consent_display_name = "Administer"
    is_enabled                 = true
    type                       = "Admin"
    value                      = "administer"
  }

  oauth2_permissions {
    admin_consent_description  = "Allow the application to access acctest-APP-%[2]d on behalf of the signed-in user."
    admin_consent_display_name = "Access acctest-APP-%[2]d"
    is_enabled                 = true
    type                       = "User"
    user_consent_description   = "Allow the application to access acctest-APP-%[2]d on your behalf."
    user_consent_display_name  = "Access acctest-APP-%[2]d"
    value                      = "user_impersonation"
  }

  optional_claims {
    access_token {
      name = "myclaim"
    }

    access_token {
      name = "otherclaim"
    }

    id_token {
      name                  = "userclaim"
      source                = "user"
      essential             = true
      additional_properties = ["emit_as_roles"]
    }
  }

  owners = [azuread_user.test.object_id]
}
`, testAccADUser_basic(ri, pw), ri)
}

func testAccADApplication_appRoles(ri int) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name = "acctest-APP-%[1]d"

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
`, ri)
}

func testAccADApplication_appRolesNoValue(ri int) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name = "acctest-APP-%[1]d"

  app_role {
    allowed_member_types = ["User"]
    description          = "Admins can manage roles and perform all task actions"
    display_name         = "Admin"
    is_enabled           = true
  }
}
`, ri)
}

func testAccADApplication_appRolesUpdate(ri int) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name = "acctestApp-%d"

  app_role {
    allowed_member_types = ["User"]
    description          = "Admins can manage roles and perform all task actions"
    display_name         = "Admin"
    is_enabled           = true
    value                = ""
  }

  app_role {
    allowed_member_types = ["User"]
    description          = "ReadOnly roles have limited query access"
    display_name         = "ReadOnly"
    is_enabled           = true
    value                = "User"
  }
}
`, ri)
}

func testAccADApplication_oauth2Permissions(ri int) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name = "acctest-APP-%[1]d"

  oauth2_permissions {
    admin_consent_description  = "Administer the application"
    admin_consent_display_name = "Administer"
    is_enabled                 = true
    type                       = "Admin"
    value                      = "administer"
  }

  oauth2_permissions {
    admin_consent_description  = "Allow the application to access acctest-APP-%[1]d on behalf of the signed-in user."
    admin_consent_display_name = "Access acctest-APP-%[1]d"
    is_enabled                 = true
    type                       = "User"
    user_consent_description   = "Allow the application to access acctest-APP-%[1]d on your behalf."
    user_consent_display_name  = "Access acctest-APP-%[1]d"
    value                      = "user_impersonation"
  }
}
`, ri)
}

func testAccADApplication_oauth2PermissionsEmpty(ri int) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name               = "acctest-APP-%[1]d"
  oauth2_permissions = []
}
`, ri)
}

func testAccADApplication_native(ri int) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name = "acctest-APP-%[1]d"
  type = "native"
}
`, ri)
}

func testAccADApplication_nativeReplyUrls(ri int) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name       = "acctest-APP-%[1]d"
  type       = "native"
  reply_urls = ["urn:ietf:wg:oauth:2.0:oob"]
}
`, ri)
}

func testAccADApplication_native_app_does_not_allow_identifier_uris(ri int) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name            = "acctest-APP-%[1]d"
  identifier_uris = ["http://%[1]d.hashicorptest.com"]
  type            = "native"
}
`, ri)
}

func testAccADApplication_duplicateName(ri int) string {
	return fmt.Sprintf(`
%s

resource "azuread_application" "duplicate" {
  name                    = azuread_application.test.name
  prevent_duplicate_names = true
}
`, testAccADApplication_basic(ri))
}
