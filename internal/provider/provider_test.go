package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/go-azure-helpers/authentication"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestProvider(t *testing.T) {
	if err := AzureADProvider().(*schema.Provider).InternalValidate(); err != nil {
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

	provider := AzureADProvider().(*schema.Provider)

	// Support only Azure CLI authentication
	provider.ConfigureFunc = func(d *schema.ResourceData) (interface{}, error) {
		builder := &authentication.Builder{
			TenantID:              d.Get("tenant_id").(string),
			MetadataHost:          d.Get("metadata_host").(string),
			Environment:           d.Get("environment").(string),
			SupportsAzureCliToken: true,
			TenantOnly:            true,
		}

		return buildClient(provider, builder, "")
	}

	err := provider.Configure(terraform.NewResourceConfigRaw(nil))
	if err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestAccProvider_servicePrincipalAuth(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		return
	}

	provider := AzureADProvider().(*schema.Provider)

	// Support only Service Principal authentication (certificate or secret)
	provider.ConfigureFunc = func(d *schema.ResourceData) (interface{}, error) {
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

		return buildClient(provider, builder, "")
	}

	err := provider.Configure(terraform.NewResourceConfigRaw(nil))
	if err != nil {
		t.Fatalf("err: %s", err)
	}
}
