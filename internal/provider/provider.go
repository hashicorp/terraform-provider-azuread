package provider

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/go-azure-helpers/authentication"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/manicminer/hamilton/auth"
	"github.com/manicminer/hamilton/environments"

	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
)

// Microsoft’s Terraform Partner ID is this specific GUID
const terraformPartnerId = "222c6c49-1b0a-5959-a213-6608f9eb8820"

type ServiceRegistration interface {
	// Name is the name of this Service
	Name() string

	// WebsiteCategories returns a list of categories which can be used for the sidebar
	WebsiteCategories() []string

	// SupportedDataSources returns the supported Data Sources supported by this Service
	SupportedDataSources() map[string]*schema.Resource

	// SupportedResources returns the supported Resources supported by this Service
	SupportedResources() map[string]*schema.Resource
}

// AzureADProvider returns a schema.Provider.
func AzureADProvider() *schema.Provider {
	// avoids this showing up in test output
	var debugLog = func(f string, v ...interface{}) {
		if os.Getenv("TF_LOG") == "" {
			return
		}

		if os.Getenv("TF_ACC") != "" {
			return
		}

		log.Printf(f, v...)
	}

	dataSources := make(map[string]*schema.Resource)
	resources := make(map[string]*schema.Resource)
	for _, service := range SupportedServices() {
		debugLog("[DEBUG] Registering Resources for %q..", service.Name())
		for k, v := range service.SupportedResources() {
			if existing := resources[k]; existing != nil {
				panic(fmt.Sprintf("An existing Resource exists for %q", k))
			}

			resources[k] = v
		}

		debugLog("[DEBUG] Registering Data Sources for %q..", service.Name())
		for k, v := range service.SupportedDataSources() {
			if existing := dataSources[k]; existing != nil {
				panic(fmt.Sprintf("An existing Data Source exists for %q", k))
			}

			dataSources[k] = v
		}
	}

	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"client_id": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ARM_CLIENT_ID", ""),
				Description: "The Client ID which should be used for service principal authentication.",
			},

			"tenant_id": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ARM_TENANT_ID", ""),
				Description: "The Tenant ID which should be used. Works with all authentication methods except Managed Identity.",
			},

			"metadata_host": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ARM_METADATA_HOSTNAME", ""),
				Description: "[DEPRECATED] The Hostname which should be used for the Azure Metadata Service.",
			},

			"environment": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ARM_ENVIRONMENT", "global"),
				Description: "The cloud environment which should be used. Possible values are `global` (formerly `public`), `usgovernment`, `dod`, `germany`, and `china`. Defaults to `global`.",
			},

			// Client Certificate specific fields
			"client_certificate_password": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ARM_CLIENT_CERTIFICATE_PASSWORD", ""),
			},

			"client_certificate_path": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ARM_CLIENT_CERTIFICATE_PATH", ""),
				Description: "The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service Principal using a Client Certificate.",
			},

			// Client Secret specific fields
			"client_secret": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ARM_CLIENT_SECRET", ""),
				Description: "The password to decrypt the Client Certificate. For use when authenticating as a Service Principal using a Client Certificate",
			},

			// CLI authentication specific fields
			"use_cli": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ARM_USE_CLI", true),
				Description: "Allow Azure CLI to be used for Authentication.",
			},

			// Managed Identity specific fields
			"use_msi": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ARM_USE_MSI", false),
				Description: "Allow Managed Identity to be used for Authentication.",
			},

			"msi_endpoint": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ARM_MSI_ENDPOINT", ""),
				Description: "The path to a custom endpoint for Managed Identity - in most circumstances this should be detected automatically. ",
			},

			// Managed Tracking GUID for User-agent
			"partner_id": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.Any(validation.IsUUID, validation.StringIsEmpty),
				DefaultFunc:  schema.EnvDefaultFunc("ARM_PARTNER_ID", ""),
				Description:  "A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution.",
			},

			"disable_terraform_partner_id": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ARM_DISABLE_TERRAFORM_PARTNER_ID", false),
				Description: "Disable the Terraform Partner ID which is used if a custom `partner_id` isn't specified.",
			},

			// MS Graph beta
			// TODO: remove in v2.0
			"use_microsoft_graph": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("AAD_USE_MICROSOFT_GRAPH", false),
				Description: "Beta: Use the Microsoft Graph API, instead of the legacy Azure Active Directory Graph API, where supported.",
			},
		},

		ResourcesMap:   resources,
		DataSourcesMap: dataSources,
	}

	p.ConfigureContextFunc = providerConfigure(p)

	return p
}

func providerConfigure(p *schema.Provider) schema.ConfigureContextFunc {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		environment, aadEnvironment := environment(d.Get("environment").(string))

		// Microsoft Graph beta opt-in
		enableMsGraph := d.Get("use_microsoft_graph").(bool)

		var authConfig *auth.Config
		if enableMsGraph {
			authConfig = &auth.Config{
				Environment:            environment,
				TenantID:               d.Get("tenant_id").(string),
				ClientID:               d.Get("client_id").(string),
				ClientCertPassword:     d.Get("client_certificate_password").(string),
				ClientCertPath:         d.Get("client_certificate_path").(string),
				ClientSecret:           d.Get("client_secret").(string),
				EnableClientCertAuth:   true,
				EnableClientSecretAuth: true,
				EnableAzureCliToken:    d.Get("use_cli").(bool),
				EnableMsiAuth:          d.Get("use_msi").(bool),
				MsiEndpoint:            d.Get("msi_endpoint").(string),
			}
		}

		aadBuilder := &authentication.Builder{
			ClientID:           d.Get("client_id").(string),
			ClientSecret:       d.Get("client_secret").(string),
			TenantID:           d.Get("tenant_id").(string),
			MetadataHost:       d.Get("metadata_host").(string),
			Environment:        aadEnvironment,
			MsiEndpoint:        d.Get("msi_endpoint").(string),
			ClientCertPassword: d.Get("client_certificate_password").(string),
			ClientCertPath:     d.Get("client_certificate_path").(string),

			// Feature Toggles
			SupportsClientCertAuth:         true,
			SupportsClientSecretAuth:       true,
			SupportsManagedServiceIdentity: d.Get("use_msi").(bool),
			SupportsAzureCliToken:          d.Get("use_cli").(bool),
			TenantOnly:                     true,
		}

		// only one pid can be interpreted currently
		// hence, send partner ID if present, otherwise send Terraform GUID
		// unless users have opted out
		partnerId := d.Get("partner_id").(string)
		if partnerId == "" && !d.Get("disable_terraform_partner_id").(bool) {
			partnerId = terraformPartnerId
		}

		return buildClient(ctx, p, authConfig, aadBuilder, partnerId, enableMsGraph)
	}
}

// TODO: v2.0 pull out authentication.Builder and derived configuration
func buildClient(ctx context.Context, p *schema.Provider, authConfig *auth.Config, b *authentication.Builder, partnerId string, enableMsGraph bool) (*clients.Client, diag.Diagnostics) {
	aadConfig, err := b.Build()
	if err != nil {
		return nil, tf.ErrorDiagF(err, "Building AzureAD Client")
	}

	clientBuilder := clients.ClientBuilder{
		AuthConfig:       authConfig,
		AadAuthConfig:    aadConfig,
		EnableMsGraph:    enableMsGraph,
		PartnerID:        partnerId,
		TerraformVersion: p.TerraformVersion,
	}

	stopCtx, ok := schema.StopContext(ctx) //nolint:SA1019
	if !ok {
		stopCtx = ctx
	}

	client, err := clientBuilder.Build(stopCtx)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return client, nil
}

func environment(name string) (env environments.Environment, aadEnv string) {
	switch name {
	case "global", "public":
		env = environments.Global
		aadEnv = "public"
	case "usgovernment", "usgovernmentl4":
		env = environments.USGovernmentL4
		aadEnv = "usgovernment"
	case "dod", "usgovernmentl5":
		env = environments.USGovernmentL5
		aadEnv = "usgovernment"
	case "german", "germany":
		env = environments.Germany
		aadEnv = "german"
	case "china":
		env = environments.China
		aadEnv = "china"
	}
	return
}
