package azuread

import (
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccClientConfigDataSource_basic(t *testing.T) {
	dsn := "data.azuread_client_config.current"
	clientId := os.Getenv("ARM_CLIENT_ID")
	tenantId := os.Getenv("ARM_TENANT_ID")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckArmClientConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dsn, "client_id", clientId),
					resource.TestCheckResourceAttr(dsn, "tenant_id", tenantId),
					testAzureRMClientConfigGUIDAttr(dsn, "object_id"),
				),
			},
		},
	})
}

func TestAccClientConfigDataSource_basicDeprecated(t *testing.T) { // TODO: remove in v1.0
	dsn := "data.azuread_client_config.current"
	clientId := os.Getenv("ARM_CLIENT_ID")
	tenantId := os.Getenv("ARM_TENANT_ID")
	subscriptionId := os.Getenv("ARM_SUBSCRIPTION_ID")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckArmClientConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dsn, "client_id", clientId),
					resource.TestCheckResourceAttr(dsn, "tenant_id", tenantId),
					resource.TestCheckResourceAttr(dsn, "subscription_id", subscriptionId),
					testAzureRMClientConfigGUIDAttr(dsn, "object_id"),
				),
			},
		},
	})
}

func testAzureRMClientConfigGUIDAttr(name, key string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		r, err := regexp.Compile("^[A-Fa-f0-9]{8}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{12}$")
		if err != nil {
			return err
		}

		return resource.TestMatchResourceAttr(name, key, r)(s)
	}
}

const testAccCheckArmClientConfig_basic = `
data "azuread_client_config" "current" { }
`
