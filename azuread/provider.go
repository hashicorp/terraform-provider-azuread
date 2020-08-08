package azuread

import (
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-azuread/azuread/internal/provider"
)

func Provider() terraform.ResourceProvider {
	return provider.AzureADProvider()
}
