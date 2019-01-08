package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/terraform-providers/terraform-provider-azuread/azuread"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: azuread.Provider})
}
