package azuread

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/ar"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/tf"
)

func TestAccAzureADDomainService_basic(t *testing.T) {
	resourceName := "azuread_domain_service.test"
	ri := tf.AccRandTimeInt()
	location := testLocation()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureADDomainServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureADDomainService_basic(ri, location),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "filtered_sync", "true"),
					resource.TestCheckResourceAttr(resourceName, "domain_controller_ip_address.#", "2"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"ldaps.0.pfx_certificate",
					"ldaps.0.pfx_certificate_password",
				},
			},
		},
	})
}

func TestAccAzureADDomainService_complete(t *testing.T) {
	resourceName := "azuread_domain_service.test"
	ri := tf.AccRandTimeInt()
	location := testLocation()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureADDomainServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureADDomainService_complete(ri, location),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "filtered_sync", "false"),
					resource.TestCheckResourceAttr(resourceName, "domain_controller_ip_address.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "ldaps.0.ldaps", "true"),
					resource.TestCheckResourceAttr(resourceName, "ldaps.0.external_access", "true"),
					resource.TestCheckResourceAttrSet(resourceName, "ldaps.0.external_access_ip_address"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"ldaps.0.pfx_certificate",
					"ldaps.0.pfx_certificate_password",
				},
			},
		},
	})
}

func TestAccAzureADDomainService_update(t *testing.T) {
	resourceName := "azuread_domain_service.test"
	ri := tf.AccRandTimeInt()
	location := testLocation()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testCheckAzureADDomainServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureADDomainService_basic(ri, location),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "filtered_sync", "true"),
					resource.TestCheckResourceAttr(resourceName, "domain_controller_ip_address.#", "2"),
				),
			},
			{
				Config: testAccAzureADDomainService_complete(ri, location),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "filtered_sync", "false"),
					resource.TestCheckResourceAttr(resourceName, "domain_controller_ip_address.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "ldaps.0.ldaps", "true"),
					resource.TestCheckResourceAttr(resourceName, "ldaps.0.external_access", "true"),
					resource.TestCheckResourceAttrSet(resourceName, "ldaps.0.external_access_ip_address"),
				),
			},
			{
				Config: testAccAzureADDomainService_basic(ri, location),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "filtered_sync", "true"),
					resource.TestCheckResourceAttr(resourceName, "domain_controller_ip_address.#", "2"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"ldaps.0.pfx_certificate",
					"ldaps.0.pfx_certificate_password",
				},
			},
		},
	})
}

func testCheckAzureADDomainServiceDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*ArmClient).domainServicesClient
	ctx := testAccProvider.Meta().(*ArmClient).StopContext

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "azuread_domain_service" {
			continue
		}

		name := rs.Primary.Attributes["name"]
		resourceGroup := rs.Primary.Attributes["resource_group_name"]

		resp, err := client.Get(ctx, resourceGroup, name)
		if err != nil {
			if ar.ResponseWasNotFound(resp.Response) {
				return nil
			}

			return err
		}

		if resp.StatusCode != http.StatusNotFound {
			return fmt.Errorf("Domain Service still exists:\n%#v", resp)
		}
	}

	return nil
}

func testAccAzureADDomainService_basic(rInt int, location string) string {
	return fmt.Sprintf(`
%s

resource "azuread_domain_service" "test" {
  name                  = "test.onmicrosoft.com"
  location              = "${azurerm_resource_group.test.location}"
  resource_group_name   = "${azurerm_resource_group.test.name}"
  subnet_id             = "${azurerm_subnet.test.id}"
}
`, testAccAzureADDomainService_template(rInt, location))
}

func testAccAzureADDomainService_complete(rInt int, location string) string {
	return fmt.Sprintf(`
%s

resource "azuread_domain_service" "test" {
  name                  = "test.onmicrosoft.com"
  location              = "${azurerm_resource_group.test.location}"
  resource_group_name   = "${azurerm_resource_group.test.name}"
  subnet_id             = "${azurerm_subnet.test.id}"
  filtered_sync         = false

  security {
    ntlm_v1 = true
    tls_v1 = true
    sync_ntlm_passwords = true
  }

  ldaps {
	external_access = true
    ldaps = true
    pfx_certificate = "${filebase64("testdata/domain_service_test.pfx")}"
    pfx_certificate_password = "test"
  }
}
`, testAccAzureADDomainService_template(rInt, location))
}

func testAccAzureADDomainService_template(rInt int, location string) string {
	return fmt.Sprintf(`
resource "azurerm_resource_group" "test" {
  name     = "acctestRG-aadds-%d"
  location = "%s"
}

resource "azurerm_network_security_group" "test" {
  name                = "acctestNSG-%d"
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"

  security_rule {
    name                        = "AllowSyncWithAzureAD"
    priority                    = 101
    direction                   = "Inbound"
    access                      = "Allow"
    protocol                    = "Tcp"
    source_port_range           = "*"
    destination_port_range      = "443"
    source_address_prefix       = "AzureActiveDirectoryDomainServices"
    destination_address_prefix  = "*"
  }

  security_rule {
    name                        = "AllowRD"
    priority                    = 201
    direction                   = "Inbound"
    access                      = "Allow"
    protocol                    = "Tcp"
    source_port_range           = "*"
    destination_port_range      = "3389"
    source_address_prefix       = "CorpNetSaw"
    destination_address_prefix  = "*"
  }

  security_rule {
    name                        = "AllowPSRemoting"
    priority                    = 301
    direction                   = "Inbound"
    access                      = "Allow"
    protocol                    = "Tcp"
    source_port_range           = "*"
    destination_port_range      = "5986"
    source_address_prefix       = "AzureActiveDirectoryDomainServices"
    destination_address_prefix  = "*"
  }

  security_rule {
    name                        = "AllowLDAPS"
    priority                    = 401
    direction                   = "Inbound"
    access                      = "Allow"
    protocol                    = "Tcp"
    source_port_range           = "*"
    destination_port_range      = "636"
    source_address_prefix       = "*"
    destination_address_prefix  = "*"
  }
}

resource "azurerm_virtual_network" "test" {
  name                = "acctestVnet-%d"
  address_space       = ["10.0.1.0/24"]
  location            = "${azurerm_resource_group.test.location}"
  resource_group_name = "${azurerm_resource_group.test.name}"

  lifecycle {
    ignore_changes = [ dns_servers ]
  }
}

resource "azurerm_subnet" "test" {
  name                 = "acctestSubnet-%d"
  virtual_network_name = "${azurerm_virtual_network.test.name}"
  resource_group_name = "${azurerm_resource_group.test.name}"
  address_prefix       = "10.0.1.0/24"
  network_security_group_id = "${azurerm_network_security_group.test.id}"
}
`, rInt, location, rInt, rInt, rInt)
}
