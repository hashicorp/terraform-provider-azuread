package aadgraph_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/terraform-providers/terraform-provider-azuread/internal/acceptance"
)

func TestAccDomainsDataSource_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_domains", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDomainsDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(data.ResourceName, "domains.0.domain_name"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "domains.0.authentication_type"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "domains.0.is_default"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "domains.0.is_initial"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "domains.0.is_verified"),
				),
			},
		},
	})
}

func TestAccDomainsDataSource_onlyDefault(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_domains", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDomainsDataSource_onlyDefault,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(data.ResourceName, "domains.0.domain_name"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "domains.0.is_default"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "domains.0.is_verified"),
					resource.TestCheckResourceAttr(data.ResourceName, "domains.0.is_default", "true"),
				),
			},
		},
	})
}

func TestAccDomainsDataSource_onlyInitial(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azuread_domains", "test")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDomainsDataSource_onlyInitial,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(data.ResourceName, "domains.0.domain_name"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "domains.0.is_default"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "domains.0.is_verified"),
					resource.TestCheckResourceAttr(data.ResourceName, "domains.0.is_initial", "true"),
				),
			},
		},
	})
}

const (
	testAccDomainsDataSource_basic = `data "azuread_domains" "test" {}`

	testAccDomainsDataSource_onlyDefault = `
data "azuread_domains" "test" {
  only_default = true
}
`

	testAccDomainsDataSource_onlyInitial = `
data "azuread_domains" "test" {
  only_initial = true
}
`
)
