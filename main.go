package main

import (
	"flag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	if debugMode {
		plugin.Serve(&plugin.ServeOpts{
			ProviderFunc: Provider,
			ProviderAddr: "registry.terraform.io/hashicorp/azuread",
			Debug:        true,
		})
	} else {
		plugin.Serve(&plugin.ServeOpts{
			ProviderFunc: Provider,
		})
	}
}
