// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package acceptance

import (
	"os"
	"sync"

	"github.com/glueckkanja/terraform-provider-azuread/internal/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var AzureADProvider *schema.Provider
var once sync.Once

func init() {
	if os.Getenv("TF_ACC") == "" {
		return
	}
	EnsureProvidersAreInitialised()
}

func EnsureProvidersAreInitialised() {
	once.Do(func() {
		AzureADProvider = provider.AzureADProvider()
	})
}
