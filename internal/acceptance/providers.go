package acceptance

import (
	"sync"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-providers/terraform-provider-azuread/internal/provider"
)

var once sync.Once

func EnsureProvidersAreInitialised() {
	once.Do(func() {
		AzureADProvider = provider.AzureADProvider()
		ProviderFactories = map[string]func() (*schema.Provider, error){
			"azuread": func() (*schema.Provider, error) { return AzureADProvider, nil },
		}
		SupportedProviders = map[string]*schema.Provider{ // TODO deprecated
			"azuread": AzureADProvider,
		}
	})
}
