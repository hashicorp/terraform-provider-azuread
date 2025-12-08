// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"flag"

	"github.com/glueckkanja/terraform-provider-azuread/internal/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := &plugin.ServeOpts{
		Debug:        false,
		ProviderAddr: "registry.terraform.io/hashicorp/azuread",
		ProviderFunc: provider.AzureADProvider,
	}

	plugin.Serve(opts)
}
