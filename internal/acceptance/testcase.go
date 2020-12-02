package acceptance

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/terraform-providers/terraform-provider-azuread/internal/acceptance/helpers"
	"github.com/terraform-providers/terraform-provider-azuread/internal/acceptance/types"
	"github.com/terraform-providers/terraform-provider-azuread/internal/clients"
)

// lintignore:AT001
func (td TestData) DataSourceTest(t *testing.T, steps []resource.TestStep) {
	testCase := resource.TestCase{
		PreCheck: func() { PreCheck(t) },
		Steps:    steps,
	}

	td.runAcceptanceTest(t, testCase)
}

func (td TestData) ResourceTest(t *testing.T, testResource types.TestResource, steps []resource.TestStep) {
	testCase := resource.TestCase{
		PreCheck: func() { PreCheck(t) },
		CheckDestroy: func(s *terraform.State) error {
			client := buildClient()
			return helpers.CheckDestroyedFunc(client, testResource, td.ResourceType, td.resourceLabel)(s)
		},
		Steps: steps,
	}

	td.runAcceptanceTest(t, testCase)
}

func buildClient() *clients.AadClient {
	// if enableBinaryTesting {
	//   TODO: build up a client on demand
	//   NOTE: this'll want caching/a singleton, and likely RP registration etc disabled, since otherwise this'll become
	//   		 extremely expensive - and this doesn't need access to the provider feature toggles
	// }

	return AzureADProvider.Meta().(*clients.AadClient)
}

func (td TestData) runAcceptanceTest(t *testing.T, testCase resource.TestCase) {
	testCase.ProviderFactories = map[string]func() (*schema.Provider, error){
		"azuread": func() (*schema.Provider, error) { return AzureADProvider, nil },
	}

	resource.ParallelTest(t, testCase)
}
