package azuread

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccDataSourceAzureADDomains_basic(t *testing.T) {
	dataSourceName := "data.azuread_domains.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: `data "azuread_domains" "test" {}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataSourceName, "domains.0.domain_name"),
					resource.TestCheckResourceAttrSet(dataSourceName, "domains.0.is_default"),
					resource.TestCheckResourceAttrSet(dataSourceName, "domains.0.is_initial"),
					resource.TestCheckResourceAttrSet(dataSourceName, "domains.0.is_verified"),
				),
			},
		},
	})
}

func TestAccDataSourceAzureADDomains_tenantDomainOnly(t *testing.T) {
	dataSourceName := "data.azuread_domains.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: `data "azuread_domains" "test" {
					tenant_domain_only = true
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataSourceName, "domains.0.domain_name"),
					resource.TestCheckResourceAttrSet(dataSourceName, "domains.0.is_default"),
					resource.TestCheckResourceAttrSet(dataSourceName, "domains.0.is_initial"),
					resource.TestCheckResourceAttrSet(dataSourceName, "domains.0.is_verified"),
				),
			},
		},
	})
}
