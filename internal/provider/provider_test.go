package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"os"
	"testing"

	"github.com/hashicorp/go-azure-helpers/authentication"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestProvider(t *testing.T) {
	if err := AzureADProvider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ = AzureADProvider()
}

func TestAccProvider_cliAuth(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		return
	}

	provider := AzureADProvider()
	ctx := context.Background()

	// Support only Azure CLI authentication
	provider.ConfigureContextFunc = func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		builder := &authentication.Builder{
			TenantID:              d.Get("tenant_id").(string),
			MetadataHost:          d.Get("metadata_host").(string),
			Environment:           d.Get("environment").(string),
			SupportsAzureCliToken: true,
			TenantOnly:            true,
		}

		return buildClient(ctx, provider, builder, "")
	}

	d := provider.Configure(ctx, terraform.NewResourceConfigRaw(nil))
	if d != nil && d.HasError() {
		t.Fatalf("err: %+v", d)
	}
}

func TestAccProvider_servicePrincipalAuth(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		return
	}

	provider := AzureADProvider()
	ctx := context.Background()

	// Support only Service Principal authentication (certificate or secret)
	provider.ConfigureContextFunc = func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		builder := &authentication.Builder{
			ClientID:                 d.Get("client_id").(string),
			ClientSecret:             d.Get("client_secret").(string),
			TenantID:                 d.Get("tenant_id").(string),
			MetadataHost:             d.Get("metadata_host").(string),
			Environment:              d.Get("environment").(string),
			ClientCertPassword:       d.Get("client_certificate_password").(string),
			ClientCertPath:           d.Get("client_certificate_path").(string),
			SupportsClientCertAuth:   true,
			SupportsClientSecretAuth: true,
			TenantOnly:               true,
		}

		return buildClient(ctx, provider, builder, "")
	}

	d := provider.Configure(ctx, terraform.NewResourceConfigRaw(nil))
	if d != nil && d.HasError() {
		t.Fatalf("err: %+v", d)
	}
}
