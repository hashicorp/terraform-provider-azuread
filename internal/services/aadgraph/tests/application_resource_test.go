package aadgraph_test

import (
	"context"
	"fmt"
	"github.com/terraform-providers/terraform-provider-azuread/internal/acceptance/check"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/terraform-providers/terraform-provider-azuread/internal/acceptance"
	"github.com/terraform-providers/terraform-provider-azuread/internal/clients"
	"github.com/terraform-providers/terraform-provider-azuread/internal/utils"
)

type ApplicationResource struct {}
var r = ApplicationResource{}

func TestAccApplication_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azuread_application", "test")

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
				check.That(data.ResourceName).Key("application_type").HasValue("web"),
			),
		},
		data.ImportStep(),
	})
}

//func TestAccApplication_basic(t *testing.T) {
//	data := acceptance.BuildTestData(t, "azuread_application", "test")
//
//	resource.ParallelTest(t, resource.TestCase{
//		PreCheck:          func() { acceptance.PreCheck(t) },
//		ProviderFactories: acceptance.ProviderFactories,
//		CheckDestroy:      testCheckApplicationDestroy,
//		Steps: []resource.TestStep{
//			{
//				Config: testAccApplication_basic(data.RandomInteger),
//				Check: resource.ComposeTestCheckFunc(
//					testCheckApplicationExists(data.ResourceName),
//					resource.TestCheckResourceAttrSet(data.ResourceName, "application_id"),
//					resource.TestCheckResourceAttrSet(data.ResourceName, "object_id"),
//				),
//			},
//			data.ImportStep(),
//		},
//	})
//}
//
//func TestAccApplication_complete(t *testing.T) {
//	data := acceptance.BuildTestData(t, "azuread_application", "test")
//	pw := "utils@$$wR2" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)
//
//	resource.ParallelTest(t, resource.TestCase{
//		PreCheck:          func() { acceptance.PreCheck(t) },
//		ProviderFactories: acceptance.ProviderFactories,
//		CheckDestroy:      testCheckApplicationDestroy,
//		Steps: []resource.TestStep{
//			{
//				Config: testAccApplication_complete(data.RandomInteger, pw),
//				Check: resource.ComposeTestCheckFunc(
//					testCheckApplicationExists(data.ResourceName),
//					resource.TestCheckResourceAttrSet(data.ResourceName, "application_id"),
//					resource.TestCheckResourceAttrSet(data.ResourceName, "object_id"),
//				),
//			},
//			data.ImportStep(),
//		},
//	})
//}
//
//func TestAccApplication_update(t *testing.T) {
//	data := acceptance.BuildTestData(t, "azuread_application", "test")
//	updatedri := tf.AccRandTimeInt()
//	pw := "utils@$$wR2" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)
//
//	resource.ParallelTest(t, resource.TestCase{
//		PreCheck:          func() { acceptance.PreCheck(t) },
//		ProviderFactories: acceptance.ProviderFactories,
//		CheckDestroy:      testCheckApplicationDestroy,
//		Steps: []resource.TestStep{
//			{
//				Config: testAccApplication_basic(data.RandomInteger),
//				Check: resource.ComposeTestCheckFunc(
//					testCheckApplicationExists(data.ResourceName),
//					resource.TestCheckResourceAttrSet(data.ResourceName, "application_id"),
//					resource.TestCheckResourceAttrSet(data.ResourceName, "object_id"),
//				),
//			},
//			data.ImportStep(),
//			{
//				Config: testAccApplication_complete(updatedri, pw),
//				Check: resource.ComposeTestCheckFunc(
//					testCheckApplicationExists(data.ResourceName),
//					resource.TestCheckResourceAttrSet(data.ResourceName, "application_id"),
//					resource.TestCheckResourceAttrSet(data.ResourceName, "object_id"),
//				),
//			},
//			data.ImportStep(),
//			{
//				Config: testAccApplication_basicEmpty(data.RandomInteger),
//				Check: resource.ComposeTestCheckFunc(
//					testCheckApplicationExists(data.ResourceName),
//					resource.TestCheckResourceAttrSet(data.ResourceName, "application_id"),
//					resource.TestCheckResourceAttrSet(data.ResourceName, "object_id"),
//				),
//			},
//			data.ImportStep(),
//		},
//	})
//}
//
//func TestAccApplication_publicClient(t *testing.T) {
//	data := acceptance.BuildTestData(t, "azuread_application", "test")
//
//	resource.ParallelTest(t, resource.TestCase{
//		PreCheck:          func() { acceptance.PreCheck(t) },
//		ProviderFactories: acceptance.ProviderFactories,
//		CheckDestroy:      testCheckApplicationDestroy,
//		Steps: []resource.TestStep{
//			{
//				Config: testAccApplication_publicClient(data.RandomInteger),
//				Check: resource.ComposeTestCheckFunc(
//					testCheckApplicationExists(data.ResourceName),
//				),
//			},
//			data.ImportStep(),
//		},
//	})
//}
//
//func TestAccApplication_appRoles(t *testing.T) {
//	data := acceptance.BuildTestData(t, "azuread_application", "test")
//
//	resource.ParallelTest(t, resource.TestCase{
//		PreCheck:          func() { acceptance.PreCheck(t) },
//		ProviderFactories: acceptance.ProviderFactories,
//		CheckDestroy:      testCheckApplicationDestroy,
//		Steps: []resource.TestStep{
//			{
//				Config: testAccApplication_appRoles(data.RandomInteger),
//				Check: resource.ComposeTestCheckFunc(
//					testCheckApplicationExists(data.ResourceName),
//					resource.TestCheckResourceAttr(data.ResourceName, "app_role.#", "1"),
//				),
//			},
//			data.ImportStep(),
//		},
//	})
//}
//
//func TestAccApplication_appRolesNoValue(t *testing.T) {
//	data := acceptance.BuildTestData(t, "azuread_application", "test")
//
//	resource.ParallelTest(t, resource.TestCase{
//		PreCheck:          func() { acceptance.PreCheck(t) },
//		ProviderFactories: acceptance.ProviderFactories,
//		CheckDestroy:      testCheckApplicationDestroy,
//		Steps: []resource.TestStep{
//			{
//				Config: testAccApplication_appRolesNoValue(data.RandomInteger),
//				Check: resource.ComposeTestCheckFunc(
//					testCheckApplicationExists(data.ResourceName),
//					resource.TestCheckResourceAttr(data.ResourceName, "app_role.#", "1"),
//				),
//			},
//			data.ImportStep(),
//		},
//	})
//}
//
//func TestAccApplication_appRolesUpdate(t *testing.T) {
//	data := acceptance.BuildTestData(t, "azuread_application", "test")
//
//	resource.ParallelTest(t, resource.TestCase{
//		PreCheck:          func() { acceptance.PreCheck(t) },
//		ProviderFactories: acceptance.ProviderFactories,
//		CheckDestroy:      testCheckApplicationDestroy,
//		Steps: []resource.TestStep{
//			{
//				Config: testAccApplication_basic(data.RandomInteger),
//				Check: resource.ComposeTestCheckFunc(
//					testCheckApplicationExists(data.ResourceName),
//					resource.TestCheckResourceAttr(data.ResourceName, "app_role.#", "0"),
//				),
//			},
//			data.ImportStep(),
//			{
//				Config: testAccApplication_appRoles(data.RandomInteger),
//				Check: resource.ComposeTestCheckFunc(
//					testCheckApplicationExists(data.ResourceName),
//					resource.TestCheckResourceAttr(data.ResourceName, "app_role.#", "1"),
//				),
//			},
//			data.ImportStep(),
//			{
//				Config: testAccApplication_appRolesUpdate(data.RandomInteger),
//				Check: resource.ComposeTestCheckFunc(
//					testCheckApplicationExists(data.ResourceName),
//					resource.TestCheckResourceAttr(data.ResourceName, "app_role.#", "2"),
//				),
//			},
//			data.ImportStep(),
//			{
//				Config: testAccApplication_appRolesEmpty(data.RandomInteger),
//				Check: resource.ComposeTestCheckFunc(
//					testCheckApplicationExists(data.ResourceName),
//					resource.TestCheckResourceAttr(data.ResourceName, "app_role.#", "0"),
//				),
//			},
//			data.ImportStep(),
//		},
//	})
//}
//
//func TestAccApplication_groupMembershipClaimsUpdate(t *testing.T) {
//	data := acceptance.BuildTestData(t, "azuread_application", "test")
//
//	resource.ParallelTest(t, resource.TestCase{
//		PreCheck:          func() { acceptance.PreCheck(t) },
//		ProviderFactories: acceptance.ProviderFactories,
//		CheckDestroy:      testCheckApplicationDestroy,
//		Steps: []resource.TestStep{
//			{
//				Config: testAccApplication_basic(data.RandomInteger),
//				Check: resource.ComposeTestCheckFunc(
//					testCheckApplicationExists(data.ResourceName),
//				),
//			},
//			data.ImportStep(),
//			{
//				Config: testAccApplication_withGroupMembershipClaimsDirectoryRole(data.RandomInteger),
//				Check: resource.ComposeTestCheckFunc(
//					testCheckApplicationExists(data.ResourceName),
//					resource.TestCheckResourceAttr(data.ResourceName, "group_membership_claims", "DirectoryRole"),
//				),
//			},
//			data.ImportStep(),
//			{
//				Config: testAccApplication_withGroupMembershipClaimsSecurityGroup(data.RandomInteger),
//				Check: resource.ComposeTestCheckFunc(
//					testCheckApplicationExists(data.ResourceName),
//					resource.TestCheckResourceAttr(data.ResourceName, "group_membership_claims", "SecurityGroup"),
//				),
//			},
//			data.ImportStep(),
//			{
//				Config: testAccApplication_withGroupMembershipClaimsApplicationGroup(data.RandomInteger),
//				Check: resource.ComposeTestCheckFunc(
//					testCheckApplicationExists(data.ResourceName),
//					resource.TestCheckResourceAttr(data.ResourceName, "group_membership_claims", "ApplicationGroup"),
//				),
//			},
//			data.ImportStep(),
//		},
//	})
//}
//
//func TestAccApplication_native(t *testing.T) {
//	data := acceptance.BuildTestData(t, "azuread_application", "test")
//
//	resource.ParallelTest(t, resource.TestCase{
//		PreCheck:          func() { acceptance.PreCheck(t) },
//		ProviderFactories: acceptance.ProviderFactories,
//		CheckDestroy:      testCheckApplicationDestroy,
//		Steps: []resource.TestStep{
//			{
//				Config: testAccApplication_native(data.RandomInteger),
//				Check: resource.ComposeTestCheckFunc(
//					testCheckApplicationExists(data.ResourceName),
//					resource.TestCheckResourceAttr(data.ResourceName, "name", fmt.Sprintf("acctest-APP-%[1]d", data.RandomInteger)),
//					resource.TestCheckResourceAttr(data.ResourceName, "homepage", ""),
//					resource.TestCheckResourceAttr(data.ResourceName, "type", "native"),
//					resource.TestCheckResourceAttr(data.ResourceName, "identifier_uris.#", "0"),
//				),
//			},
//			data.ImportStep(),
//		},
//	})
//}
//
//func TestAccApplication_nativeReplyUrls(t *testing.T) {
//	data := acceptance.BuildTestData(t, "azuread_application", "test")
//
//	resource.ParallelTest(t, resource.TestCase{
//		PreCheck:          func() { acceptance.PreCheck(t) },
//		ProviderFactories: acceptance.ProviderFactories,
//		CheckDestroy:      testCheckApplicationDestroy,
//		Steps: []resource.TestStep{
//			{
//				Config: testAccApplication_nativeReplyUrls(data.RandomInteger),
//				Check: resource.ComposeTestCheckFunc(
//					testCheckApplicationExists(data.ResourceName),
//					resource.TestCheckResourceAttr(data.ResourceName, "name", fmt.Sprintf("acctest-APP-%[1]d", data.RandomInteger)),
//					resource.TestCheckResourceAttr(data.ResourceName, "type", "native"),
//					resource.TestCheckResourceAttr(data.ResourceName, "reply_urls.#", "1"),
//				),
//			},
//			data.ImportStep(),
//		},
//	})
//}
//
//func TestAccApplication_nativeUpdate(t *testing.T) {
//	data := acceptance.BuildTestData(t, "azuread_application", "test")
//	pw := "utils@$$wR2" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)
//
//	resource.ParallelTest(t, resource.TestCase{
//		PreCheck:          func() { acceptance.PreCheck(t) },
//		ProviderFactories: acceptance.ProviderFactories,
//		CheckDestroy:      testCheckApplicationDestroy,
//		Steps: []resource.TestStep{
//			{
//				Config: testAccApplication_basic(data.RandomInteger),
//				Check: resource.ComposeTestCheckFunc(
//					testCheckApplicationExists(data.ResourceName),
//					resource.TestCheckResourceAttr(data.ResourceName, "type", "webapp/api"),
//					resource.TestCheckResourceAttr(data.ResourceName, "identifier_uris.#", "0"),
//				),
//			},
//			data.ImportStep(),
//			{
//				Config: testAccApplication_native(data.RandomInteger),
//				Check: resource.ComposeTestCheckFunc(
//					testCheckApplicationExists(data.ResourceName),
//					resource.TestCheckResourceAttr(data.ResourceName, "type", "native"),
//					resource.TestCheckResourceAttr(data.ResourceName, "identifier_uris.#", "0"),
//				),
//			},
//			data.ImportStep(),
//			{
//				Config: testAccApplication_complete(data.RandomInteger, pw),
//				Check: resource.ComposeTestCheckFunc(
//					testCheckApplicationExists(data.ResourceName),
//					resource.TestCheckResourceAttr(data.ResourceName, "type", "webapp/api"),
//					resource.TestCheckResourceAttr(data.ResourceName, "identifier_uris.#", "1"),
//				),
//			},
//			data.ImportStep(),
//		},
//	})
//}
//
//func TestAccApplication_nativeAppDoesNotAllowIdentifierUris(t *testing.T) {
//	data := acceptance.BuildTestData(t, "azuread_application", "test")
//
//	resource.ParallelTest(t, resource.TestCase{
//		PreCheck:          func() { acceptance.PreCheck(t) },
//		ProviderFactories: acceptance.ProviderFactories,
//		CheckDestroy:      testCheckApplicationDestroy,
//		Steps: []resource.TestStep{
//			{
//				Config:      testAccApplication_nativeAppDoesNotAllowIdentifierUris(data.RandomInteger),
//				ExpectError: regexp.MustCompile("not required for a native application"),
//			},
//		},
//	})
//}
//
//func TestAccApplication_oauth2PermissionsUpdate(t *testing.T) {
//	data := acceptance.BuildTestData(t, "azuread_application", "test")
//
//	resource.ParallelTest(t, resource.TestCase{
//		PreCheck:          func() { acceptance.PreCheck(t) },
//		ProviderFactories: acceptance.ProviderFactories,
//		CheckDestroy:      testCheckApplicationDestroy,
//		Steps: []resource.TestStep{
//			{
//				Config: testAccApplication_basic(data.RandomInteger),
//				Check: resource.ComposeTestCheckFunc(
//					testCheckApplicationExists(data.ResourceName),
//					resource.TestCheckResourceAttr(data.ResourceName, "oauth2_permissions.#", "1"),
//				),
//			},
//			data.ImportStep(),
//			{
//				Config: testAccApplication_oauth2Permissions(data.RandomInteger),
//				Check: resource.ComposeTestCheckFunc(
//					testCheckApplicationExists(data.ResourceName),
//					resource.TestCheckResourceAttr(data.ResourceName, "oauth2_permissions.#", "2"),
//				),
//			},
//			data.ImportStep(),
//			{
//				Config: testAccApplication_oauth2PermissionsEmpty(data.RandomInteger),
//				Check: resource.ComposeTestCheckFunc(
//					testCheckApplicationExists(data.ResourceName),
//					resource.TestCheckResourceAttr(data.ResourceName, "oauth2_permissions.#", "0"),
//				),
//			},
//			data.ImportStep(),
//		},
//	})
//}
//
//func TestAccApplication_preventDuplicateNamesOk(t *testing.T) {
//	data := acceptance.BuildTestData(t, "azuread_application", "test")
//
//	resource.ParallelTest(t, resource.TestCase{
//		PreCheck:          func() { acceptance.PreCheck(t) },
//		ProviderFactories: acceptance.ProviderFactories,
//		CheckDestroy:      testCheckApplicationDestroy,
//		Steps: []resource.TestStep{
//			{
//				Config: testAccApplication_preventDuplicateNamesOk(data.RandomInteger),
//				Check: resource.ComposeTestCheckFunc(
//					testCheckApplicationExists(data.ResourceName),
//					resource.TestCheckResourceAttr(data.ResourceName, "name", fmt.Sprintf("acctest-APP-%[1]d", data.RandomInteger)),
//				),
//			},
//			data.ImportStep("prevent_duplicate_names"),
//		},
//	})
//}
//
//func TestAccApplication_preventDuplicateNamesFail(t *testing.T) {
//	data := acceptance.BuildTestData(t, "azuread_application", "test")
//
//	resource.ParallelTest(t, resource.TestCase{
//		PreCheck:          func() { acceptance.PreCheck(t) },
//		ProviderFactories: acceptance.ProviderFactories,
//		CheckDestroy:      testCheckApplicationDestroy,
//		Steps: []resource.TestStep{
//			{
//				Config:      testAccApplication_preventDuplicateNamesFail(data.RandomInteger),
//				ExpectError: regexp.MustCompile("existing Application .+ was found"),
//			},
//		},
//	})
//}
//
//func TestAccApplication_duplicateAppRolesOauth2PermissionsValues(t *testing.T) {
//	data := acceptance.BuildTestData(t, "azuread_application", "test")
//
//	resource.ParallelTest(t, resource.TestCase{
//		PreCheck:          func() { acceptance.PreCheck(t) },
//		ProviderFactories: acceptance.ProviderFactories,
//		CheckDestroy:      testCheckApplicationDestroy,
//		Steps: []resource.TestStep{
//			{
//				Config:      testAccApplication_duplicateAppRolesOauth2PermissionsValues(data.RandomInteger),
//				ExpectError: regexp.MustCompile("validation failed: duplicate app_role / oauth2_permissions value found:"),
//			},
//		},
//	})
//}
//
//func TestAccApplication_ownersUpdate(t *testing.T) {
//	data := acceptance.BuildTestData(t, "azuread_application", "test")
//	pw := "utils@$$wR2" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)
//
//	resource.ParallelTest(t, resource.TestCase{
//		PreCheck:          func() { acceptance.PreCheck(t) },
//		ProviderFactories: acceptance.ProviderFactories,
//		CheckDestroy:      testCheckApplicationDestroy,
//		Steps: []resource.TestStep{
//			{
//				Config: testAccApplication_basic(data.RandomInteger),
//				Check: resource.ComposeTestCheckFunc(
//					testCheckApplicationExists(data.ResourceName),
//					resource.TestCheckResourceAttr(data.ResourceName, "owners.#", "1"),
//				),
//			},
//			data.ImportStep(),
//			{
//				Config: testAccApplication_removeOwners(data.RandomInteger, pw),
//				Check: resource.ComposeTestCheckFunc(
//					testCheckApplicationExists(data.ResourceName),
//					resource.TestCheckResourceAttr(data.ResourceName, "owners.#", "0"),
//				),
//			},
//			data.ImportStep(),
//			{
//				Config: testAccApplication_singleOwner(data.RandomInteger, pw),
//				Check: resource.ComposeTestCheckFunc(
//					testCheckApplicationExists(data.ResourceName),
//					resource.TestCheckResourceAttr(data.ResourceName, "owners.#", "1"),
//				),
//			},
//			data.ImportStep(),
//			{
//				Config: testAccApplication_threeOwners(data.RandomInteger, pw),
//				Check: resource.ComposeTestCheckFunc(
//					testCheckApplicationExists(data.ResourceName),
//					resource.TestCheckResourceAttr(data.ResourceName, "owners.#", "3"),
//				),
//			},
//			data.ImportStep(),
//			{
//				Config: testAccApplication_removeOwners(data.RandomInteger, pw),
//				Check: resource.ComposeTestCheckFunc(
//					testCheckApplicationExists(data.ResourceName),
//					resource.TestCheckResourceAttr(data.ResourceName, "owners.#", "0"),
//				),
//			},
//			data.ImportStep(),
//		},
//	})
//}

func (a ApplicationResource) Exists(ctx context.Context, clients *clients.AadClient, state *terraform.InstanceState) (*bool, error) {
	resp, err := clients.AadGraph.ApplicationsClient.Get(ctx, state.ID)

	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			return nil, fmt.Errorf("Application with object ID %q does not exist", state.ID)
		}
		return nil, fmt.Errorf("failed to retrieve Application with object ID %q: %+v", state.ID, err)
	}

	return utils.Bool(resp.ObjectID != nil && *resp.ObjectID == state.ID), nil
}

//func testCheckApplicationDestroy(s *terraform.State) error {
//	for _, rs := range s.RootModule().Resources {
//		if rs.Type != "azuread_application" {
//			continue
//		}
//
//		client := acceptance.AzureADProvider.Meta().(*clients.AadClient).AadGraph.ApplicationsClient
//		ctx := acceptance.AzureADProvider.Meta().(*clients.AadClient).StopContext
//		resp, err := client.Get(ctx, rs.Primary.ID)
//
//		if err != nil {
//			if utils.ResponseWasNotFound(resp.Response) {
//				return nil
//			}
//
//			return err
//		}
//
//		return fmt.Errorf("Application still exists:\n%#v", resp)
//	}
//
//	return nil
//}

func (ApplicationResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "azuread_application" "test" {
  name = "acctest-APP-%[1]d"
}
`, data.RandomInteger)
}

//func testAccApplication_basicEmpty(ri int) string {
//	return fmt.Sprintf(`
//resource "azuread_application" "test" {
//  name                    = "acctest-APP-%[1]d"
//  identifier_uris         = []
//  oauth2_permissions      = []
//  reply_urls              = []
//  group_membership_claims = "None"
//}
//`, ri)
//}
//
//func testAccApplication_publicClient(ri int) string {
//	return fmt.Sprintf(`
//resource "azuread_application" "test" {
//  name          = "acctest-APP-%[1]d"
//  type          = "native"
//  public_client = true
//}
//`, ri)
//}
//
//func testAccApplication_withGroupMembershipClaimsDirectoryRole(ri int) string {
//	return fmt.Sprintf(`
//resource "azuread_application" "test" {
//  name                    = "acctest-APP-%[1]d"
//  group_membership_claims = "DirectoryRole"
//}
//`, ri)
//}
//
//func testAccApplication_withGroupMembershipClaimsSecurityGroup(ri int) string {
//	return fmt.Sprintf(`
//resource "azuread_application" "test" {
//  name                    = "acctest-APP-%[1]d"
//  group_membership_claims = "SecurityGroup"
//}
//`, ri)
//}
//
//func testAccApplication_withGroupMembershipClaimsApplicationGroup(ri int) string {
//	return fmt.Sprintf(`
//resource "azuread_application" "test" {
//  name                    = "acctest-APP-%[1]d"
//  group_membership_claims = "ApplicationGroup"
//}
//`, ri)
//}
//
//func testAccApplication_complete(ri int, pw string) string {
//	return fmt.Sprintf(`
//%[1]s
//
//resource "azuread_application" "test" {
//  name                       = "acctest-APP-%[2]d"
//  homepage                   = "https://homepage-%[2]d"
//  identifier_uris            = ["api://hashicorptestapp-%[2]d"]
//  reply_urls                 = ["https://unittest.hashicorptest.com"]
//  logout_url                 = "https://log.me.out"
//  available_to_other_tenants = true
//  group_membership_claims    = "All"
//  oauth2_allow_implicit_flow = true
//  type                       = "webapp/api"
//
//  required_resource_access {
//    resource_app_id = "00000003-0000-0000-c000-000000000000"
//
//    resource_access {
//      id   = "7ab1d382-f21e-4acd-a863-ba3e13f7da61"
//      type = "Role"
//    }
//
//    resource_access {
//      id   = "e1fe6dd8-ba31-4d61-89e7-88639da4683d"
//      type = "Scope"
//    }
//
//    resource_access {
//      id   = "06da0dbc-49e2-44d2-8312-53f166ab848a"
//      type = "Scope"
//    }
//  }
//
//  required_resource_access {
//    resource_app_id = "00000002-0000-0000-c000-000000000000"
//
//    resource_access {
//      id   = "311a71cc-e848-46a1-bdf8-97ff7156d8e6"
//      type = "Scope"
//    }
//  }
//
//  oauth2_permissions {
//    admin_consent_description  = "Administer the application"
//    admin_consent_display_name = "Administer"
//    is_enabled                 = true
//    type                       = "Admin"
//    value                      = "administer"
//  }
//
//  oauth2_permissions {
//    admin_consent_description  = "Allow the application to access acctest-APP-%[2]d on behalf of the signed-in user."
//    admin_consent_display_name = "Access acctest-APP-%[2]d"
//    is_enabled                 = true
//    type                       = "User"
//    user_consent_description   = "Allow the application to access acctest-APP-%[2]d on your behalf."
//    user_consent_display_name  = "Access acctest-APP-%[2]d"
//    value                      = "user_impersonation"
//  }
//
//  optional_claims {
//    access_token {
//      name = "myclaim"
//    }
//
//    access_token {
//      name = "otherclaim"
//    }
//
//    id_token {
//      name                  = "userclaim"
//      source                = "user"
//      essential             = true
//      additional_properties = ["emit_as_roles"]
//    }
//  }
//
//  owners = [azuread_user.test.object_id]
//}
//`, testAccUser_basic(ri, pw), ri)
//}
//
//func testAccApplication_appRoles(ri int) string {
//	return fmt.Sprintf(`
//resource "azuread_application" "test" {
//  name = "acctest-APP-%[1]d"
//
//  app_role {
//    allowed_member_types = [
//      "User",
//      "Application",
//    ]
//
//    description  = "Admins can manage roles and perform all task actions"
//    display_name = "Admin"
//    is_enabled   = true
//    value        = "Admin"
//  }
//}
//`, ri)
//}
//
//func testAccApplication_appRolesNoValue(ri int) string {
//	return fmt.Sprintf(`
//resource "azuread_application" "test" {
//  name = "acctest-APP-%[1]d"
//
//  app_role {
//    allowed_member_types = ["User"]
//    description          = "Admins can manage roles and perform all task actions"
//    display_name         = "Admin"
//    is_enabled           = true
//  }
//}
//`, ri)
//}
//
//func testAccApplication_appRolesUpdate(ri int) string {
//	return fmt.Sprintf(`
//resource "azuread_application" "test" {
//  name = "acctestApp-%d"
//
//  app_role {
//    allowed_member_types = ["User"]
//    description          = "Admins can manage roles and perform all task actions"
//    display_name         = "Admin"
//    is_enabled           = true
//    value                = ""
//  }
//
//  app_role {
//    allowed_member_types = ["User"]
//    description          = "ReadOnly roles have limited query access"
//    display_name         = "ReadOnly"
//    is_enabled           = true
//    value                = "User"
//  }
//}
//`, ri)
//}
//
//func testAccApplication_appRolesEmpty(ri int) string {
//	return fmt.Sprintf(`
//resource "azuread_application" "test" {
//  name = "acctestApp-%d"
//
//  app_role = []
//}
//`, ri)
//}
//
//func testAccApplication_oauth2Permissions(ri int) string {
//	return fmt.Sprintf(`
//resource "azuread_application" "test" {
//  name = "acctest-APP-%[1]d"
//
//  oauth2_permissions {
//    admin_consent_description  = "Allow the application to access acctest-APP-%[1]d on behalf of the signed-in user."
//    admin_consent_display_name = "Access acctest-APP-%[1]d"
//    is_enabled                 = true
//    type                       = "User"
//    user_consent_description   = "Allow the application to access acctest-APP-%[1]d on your behalf."
//    user_consent_display_name  = "Access acctest-APP-%[1]d"
//    value                      = "user_impersonation"
//  }
//
//  oauth2_permissions {
//    admin_consent_description  = "Administer the application"
//    admin_consent_display_name = "Administer"
//    is_enabled                 = true
//    type                       = "Admin"
//    value                      = "administer"
//  }
//}
//`, ri)
//}
//
//func testAccApplication_oauth2PermissionsEmpty(ri int) string {
//	return fmt.Sprintf(`
//resource "azuread_application" "test" {
//  name               = "acctest-APP-%[1]d"
//  oauth2_permissions = []
//}
//`, ri)
//}
//
//func testAccApplication_native(ri int) string {
//	return fmt.Sprintf(`
//resource "azuread_application" "test" {
//  name = "acctest-APP-%[1]d"
//  type = "native"
//}
//`, ri)
//}
//
//func testAccApplication_nativeReplyUrls(ri int) string {
//	return fmt.Sprintf(`
//resource "azuread_application" "test" {
//  name       = "acctest-APP-%[1]d"
//  type       = "native"
//  reply_urls = ["urn:ietf:wg:oauth:2.0:oob"]
//}
//`, ri)
//}
//
//func testAccApplication_nativeAppDoesNotAllowIdentifierUris(ri int) string {
//	return fmt.Sprintf(`
//resource "azuread_application" "test" {
//  name            = "acctest-APP-%[1]d"
//  identifier_uris = ["http://%[1]d.hashicorptest.com"]
//  type            = "native"
//}
//`, ri)
//}
//
//func testAccApplication_preventDuplicateNamesOk(ri int) string {
//	return fmt.Sprintf(`
//resource "azuread_application" "test" {
//  name                    = "acctest-APP-%[1]d"
//  prevent_duplicate_names = true
//}
//`, ri)
//}
//
//func testAccApplication_preventDuplicateNamesFail(ri int) string {
//	return fmt.Sprintf(`
//%s
//
//resource "azuread_application" "duplicate" {
//  name                    = azuread_application.test.name
//  prevent_duplicate_names = true
//}
//`, testAccApplication_basic(ri))
//}
//
//func testAccApplication_duplicateAppRolesOauth2PermissionsValues(ri int) string {
//	return fmt.Sprintf(`
//resource "azuread_application" "test" {
//  name = "acctest-APP-%[1]d"
//
//  app_role {
//    allowed_member_types = ["User"]
//    description          = "Admins can manage roles and perform all task actions"
//    display_name         = "Admin"
//    is_enabled           = true
//    value                = "administer"
//  }
//
//  oauth2_permissions {
//    admin_consent_description  = "Administer the application"
//    admin_consent_display_name = "Administer"
//    is_enabled                 = true
//    type                       = "Admin"
//    value                      = "administer"
//  }
//}
//`, ri)
//}
//
//func testAccApplication_singleOwner(ri int, pw string) string {
//	return fmt.Sprintf(`
//%[1]s
//
//resource "azuread_application" "test" {
//  name = "acctest-APP-%[2]d"
//  owners = [
//    azuread_user.testA.object_id,
//  ]
//}
//`, testAccUser_threeUsersABC(ri, pw), ri)
//}
//
//func testAccApplication_threeOwners(ri int, pw string) string {
//	return fmt.Sprintf(`
//%[1]s
//
//resource "azuread_application" "test" {
//  name = "acctest-APP-%[2]d"
//  owners = [
//    azuread_user.testA.object_id,
//    azuread_user.testB.object_id,
//    azuread_user.testC.object_id,
//  ]
//}
//`, testAccUser_threeUsersABC(ri, pw), ri)
//}
//
//func testAccApplication_removeOwners(ri int, pw string) string {
//	return fmt.Sprintf(`
//%[1]s
//
//resource "azuread_application" "test" {
//  name   = "acctest-APP-%[2]d"
//  owners = []
//}
//`, testAccUser_threeUsersABC(ri, pw), ri)
//}
