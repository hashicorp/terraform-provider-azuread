package main

import (
	"github.com/terraform-providers/terraform-provider-azuread/azuread"

	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: azuread.Provider})
}
