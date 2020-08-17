package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func TestProvider(t *testing.T) {
	if err := AzureADProvider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ = AzureADProvider()
}
