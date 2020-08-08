package acceptance

import (
	"sync"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-azuread/azuread/internal/provider"
)

var once sync.Once

func EnsureProvidersAreInitialised() {
	once.Do(func() {
		AzureADProvider = provider.AzureADProvider().(*schema.Provider)
		SupportedProviders = map[string]terraform.ResourceProvider{
			"azuread": AzureADProvider,
		}
	})
}
