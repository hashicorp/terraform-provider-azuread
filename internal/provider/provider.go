// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/auth"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/sdk"
)

// Terraform's Microsoft Partner ID is this specific GUID
const terraformPartnerId = "222c6c49-1b0a-5959-a213-6608f9eb8820"

type ServiceRegistration interface {
	// Name is the name of this Service
	Name() string

	// WebsiteCategories returns a list of categories which can be used for the sidebar
	WebsiteCategories() []string

	// SupportedDataSources returns the supported Data Sources supported by this Service
	SupportedDataSources() map[string]*pluginsdk.Resource

	// SupportedResources returns the supported Resources supported by this Service
	SupportedResources() map[string]*pluginsdk.Resource
}

// AzureADProvider returns a schema.Provider.
func AzureADProvider() *schema.Provider {
	dataSources := make(map[string]*pluginsdk.Resource)
	resources := make(map[string]*pluginsdk.Resource)

	// first handle the typed services
	for _, service := range SupportedTypedServices() {
		logEntry("[DEBUG] Registering Data Sources for %q..", service.Name())
		for _, ds := range service.DataSources() {
			key := ds.ResourceType()
			if existing := dataSources[key]; existing != nil {
				panic(fmt.Sprintf("An existing Data Source exists for %q", key))
			}

			wrapper := sdk.NewDataSourceWrapper(ds)
			dataSource, err := wrapper.DataSource()
			if err != nil {
				panic(fmt.Errorf("creating Wrapper for Data Source %q: %+v", key, err))
			}

			dataSources[key] = dataSource
		}

		logEntry("[DEBUG] Registering Resources for %q..", service.Name())
		for _, r := range service.Resources() {
			key := r.ResourceType()
			if existing := resources[key]; existing != nil {
				panic(fmt.Sprintf("An existing Resource exists for %q", key))
			}

			wrapper := sdk.NewResourceWrapper(r)
			resource, err := wrapper.Resource()
			if err != nil {
				panic(fmt.Errorf("creating Wrapper for Resource %q: %+v", key, err))
			}
			resources[key] = resource
		}
	}

	// then handle the untyped services
	for _, service := range SupportedUntypedServices() {
		logEntry("[DEBUG] Registering Data Sources for %q..", service.Name())
		for k, v := range service.SupportedDataSources() {
			if existing := dataSources[k]; existing != nil {
				panic(fmt.Sprintf("An existing Data Source exists for %q", k))
			}

			dataSources[k] = v
		}

		logEntry("[DEBUG] Registering Resources for %q..", service.Name())
		for k, v := range service.SupportedResources() {
			if existing := resources[k]; existing != nil {
				panic(fmt.Sprintf("An existing Resource exists for %q", k))
			}

			resources[k] = v
		}
	}

	p := &schema.Provider{
		Schema: map[string]*pluginsdk.Schema{
			"client_id": {
				Type:        pluginsdk.TypeString,
				Optional:    true,
				DefaultFunc: pluginsdk.EnvDefaultFunc("ARM_CLIENT_ID", ""),
				Description: "The Client ID which should be used for service principal authentication",
			},

			"client_id_file_path": {
				Type:        pluginsdk.TypeString,
				Optional:    true,
				DefaultFunc: pluginsdk.EnvDefaultFunc("ARM_CLIENT_ID_FILE_PATH", ""),
				Description: "The path to a file containing the Client ID which should be used for service principal authentication",
			},

			"tenant_id": {
				Type:        pluginsdk.TypeString,
				Optional:    true,
				DefaultFunc: pluginsdk.EnvDefaultFunc("ARM_TENANT_ID", ""),
				Description: "The Tenant ID which should be used. Works with all authentication methods except Managed Identity",
			},

			"environment": {
				Type:        pluginsdk.TypeString,
				Required:    true,
				DefaultFunc: pluginsdk.EnvDefaultFunc("ARM_ENVIRONMENT", "global"),
				Description: "The cloud environment which should be used. Possible values are: `global` (also `public`), `usgovernmentl4` (also `usgovernment`), `usgovernmentl5` (also `dod`), and `china`. Defaults to `global`. Not used and should not be specified when `metadata_host` is specified.",
			},

			"metadata_host": {
				Type:        pluginsdk.TypeString,
				Required:    true,
				DefaultFunc: pluginsdk.EnvDefaultFunc("ARM_METADATA_HOSTNAME", ""),
				Description: "The Hostname which should be used for the Azure Metadata Service.",
			},

			// Client Certificate specific fields
			"client_certificate": {
				Type:        pluginsdk.TypeString,
				Optional:    true,
				DefaultFunc: pluginsdk.EnvDefaultFunc("ARM_CLIENT_CERTIFICATE", ""),
				Description: "Base64 encoded PKCS#12 certificate bundle to use when authenticating as a Service Principal using a Client Certificate",
			},

			"client_certificate_password": {
				Type:        pluginsdk.TypeString,
				Optional:    true,
				DefaultFunc: pluginsdk.EnvDefaultFunc("ARM_CLIENT_CERTIFICATE_PASSWORD", ""),
				Description: "The password to decrypt the Client Certificate. For use when authenticating as a Service Principal using a Client Certificate",
			},

			"client_certificate_path": {
				Type:        pluginsdk.TypeString,
				Optional:    true,
				DefaultFunc: pluginsdk.EnvDefaultFunc("ARM_CLIENT_CERTIFICATE_PATH", ""),
				Description: "The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service Principal using a Client Certificate",
			},

			// Client Secret specific fields
			"client_secret": {
				Type:        pluginsdk.TypeString,
				Optional:    true,
				DefaultFunc: pluginsdk.EnvDefaultFunc("ARM_CLIENT_SECRET", ""),
				Description: "The application password to use when authenticating as a Service Principal using a Client Secret",
			},

			"client_secret_file_path": {
				Type:        pluginsdk.TypeString,
				Optional:    true,
				DefaultFunc: pluginsdk.EnvDefaultFunc("ARM_CLIENT_SECRET_FILE_PATH", ""),
				Description: "The path to a file containing the application password to use when authenticating as a Service Principal using a Client Secret",
			},

			// OIDC specific fields
			"use_oidc": {
				Type:        pluginsdk.TypeBool,
				Optional:    true,
				DefaultFunc: pluginsdk.EnvDefaultFunc("ARM_USE_OIDC", false),
				Description: "Allow OpenID Connect to be used for authentication",
			},

			"oidc_token": {
				Type:        pluginsdk.TypeString,
				Optional:    true,
				DefaultFunc: pluginsdk.EnvDefaultFunc("ARM_OIDC_TOKEN", ""),
				Description: "The ID token for use when authenticating as a Service Principal using OpenID Connect.",
			},

			"oidc_token_file_path": {
				Type:        pluginsdk.TypeString,
				Optional:    true,
				DefaultFunc: pluginsdk.EnvDefaultFunc("ARM_OIDC_TOKEN_FILE_PATH", ""),
				Description: "The path to a file containing an ID token for use when authenticating as a Service Principal using OpenID Connect.",
			},

			"oidc_request_token": {
				Type:        pluginsdk.TypeString,
				Optional:    true,
				DefaultFunc: pluginsdk.MultiEnvDefaultFunc([]string{"ARM_OIDC_REQUEST_TOKEN", "ACTIONS_ID_TOKEN_REQUEST_TOKEN"}, ""),
				Description: "The bearer token for the request to the OIDC provider. For use when authenticating as a Service Principal using OpenID Connect.",
			},

			"oidc_request_url": {
				Type:        pluginsdk.TypeString,
				Optional:    true,
				DefaultFunc: pluginsdk.MultiEnvDefaultFunc([]string{"ARM_OIDC_REQUEST_URL", "ACTIONS_ID_TOKEN_REQUEST_URL"}, ""),
				Description: "The URL for the OIDC provider from which to request an ID token. For use when authenticating as a Service Principal using OpenID Connect.",
			},

			// Azure AKS Workload Identity fields
			"use_aks_workload_identity": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ARM_USE_AKS_WORKLOAD_IDENTITY", false),
				Description: "Allow Azure AKS Workload Identity to be used for Authentication.",
			},

			// CLI authentication specific fields
			"use_cli": {
				Type:        pluginsdk.TypeBool,
				Optional:    true,
				DefaultFunc: pluginsdk.EnvDefaultFunc("ARM_USE_CLI", true),
				Description: "Allow Azure CLI to be used for Authentication",
			},

			// Managed Identity specific fields
			"use_msi": {
				Type:        pluginsdk.TypeBool,
				Optional:    true,
				DefaultFunc: pluginsdk.EnvDefaultFunc("ARM_USE_MSI", false),
				Description: "Allow Managed Identity to be used for Authentication",
			},

			"msi_endpoint": {
				Type:        pluginsdk.TypeString,
				Optional:    true,
				DefaultFunc: pluginsdk.EnvDefaultFunc("ARM_MSI_ENDPOINT", ""),
				Description: "The path to a custom endpoint for Managed Identity - in most circumstances this should be detected automatically",
			},

			// Managed Tracking GUID for User-agent
			"partner_id": {
				Type:         pluginsdk.TypeString,
				Optional:     true,
				ValidateFunc: validation.Any(validation.IsUUID, validation.StringIsEmpty),
				DefaultFunc:  pluginsdk.EnvDefaultFunc("ARM_PARTNER_ID", ""),
				Description:  "A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution",
			},

			"disable_terraform_partner_id": {
				Type:        pluginsdk.TypeBool,
				Optional:    true,
				DefaultFunc: pluginsdk.EnvDefaultFunc("ARM_DISABLE_TERRAFORM_PARTNER_ID", false),
				Description: "Disable the Terraform Partner ID, which is used if a custom `partner_id` isn't specified",
			},
		},

		ResourcesMap:   resources,
		DataSourcesMap: dataSources,
	}

	p.ConfigureContextFunc = providerConfigure(p)

	return p
}

func providerConfigure(p *schema.Provider) schema.ConfigureContextFunc {
	return func(ctx context.Context, d *pluginsdk.ResourceData) (interface{}, pluginsdk.Diagnostics) {
		var certData []byte
		if encodedCert := d.Get("client_certificate").(string); encodedCert != "" {
			var err error
			certData, err = decodeCertificate(encodedCert)
			if err != nil {
				return nil, pluginsdk.DiagFromErr(err)
			}
		}

		idToken, err := getOidcToken(d)
		if err != nil {
			return nil, pluginsdk.DiagFromErr(err)
		}

		clientSecret, err := getClientSecret(d)
		if err != nil {
			return nil, pluginsdk.DiagFromErr(err)
		}

		clientId, err := getClientId(d)
		if err != nil {
			return nil, pluginsdk.DiagFromErr(err)
		}

		tenantId, err := getTenantId(d)
		if err != nil {
			return nil, pluginsdk.DiagFromErr(err)
		}

		var (
			env *environments.Environment

			envName      = d.Get("environment").(string)
			metadataHost = d.Get("metadata_host").(string)
		)

		if metadataHost != "" {
			logEntry("[DEBUG] Configuring cloud environment from Metadata Service at %q", metadataHost)
			if env, err = environments.FromEndpoint(ctx, fmt.Sprintf("https://%s", metadataHost)); err != nil {
				return nil, pluginsdk.DiagFromErr(err)
			}
		} else {
			logEntry("[DEBUG] Configuring built-in cloud environment by name: %q", envName)
			if env, err = environments.FromName(envName); err != nil {
				return nil, pluginsdk.DiagFromErr(err)
			}
		}

		if env.MicrosoftGraph == nil {
			return nil, pluginsdk.DiagErrorf("Microsoft Graph was not configured for the specified environment")
		} else if endpoint, ok := env.MicrosoftGraph.Endpoint(); !ok || *endpoint == "" {
			return nil, pluginsdk.DiagErrorf("Microsoft Graph endpoint could not be determined for the specified environment")
		}

		var (
			enableAzureCli        = d.Get("use_cli").(bool)
			enableManagedIdentity = d.Get("use_msi").(bool)
			enableOidc            = d.Get("use_oidc").(bool) || d.Get("use_aks_workload_identity").(bool)
		)

		authConfig := &auth.Credentials{
			Environment: *env,
			ClientID:    *clientId,
			TenantID:    *tenantId,

			ClientCertificateData:     certData,
			ClientCertificatePassword: d.Get("client_certificate_password").(string),
			ClientCertificatePath:     d.Get("client_certificate_path").(string),
			ClientSecret:              *clientSecret,

			OIDCAssertionToken:          *idToken,
			GitHubOIDCTokenRequestURL:   d.Get("oidc_request_url").(string),
			GitHubOIDCTokenRequestToken: d.Get("oidc_request_token").(string),

			CustomManagedIdentityEndpoint: d.Get("msi_endpoint").(string),

			EnableAuthenticatingUsingAzureCLI:          enableAzureCli,
			EnableAuthenticatingUsingClientCertificate: true,
			EnableAuthenticatingUsingClientSecret:      true,
			EnableAuthenticatingUsingManagedIdentity:   enableManagedIdentity,
			EnableAuthenticationUsingGitHubOIDC:        enableOidc,
			EnableAuthenticationUsingOIDC:              enableOidc,
		}

		// only one pid can be interpreted currently
		// hence, send partner ID if present, otherwise send Terraform GUID
		// unless users have opted out
		partnerId := d.Get("partner_id").(string)
		if partnerId == "" && !d.Get("disable_terraform_partner_id").(bool) {
			partnerId = terraformPartnerId
		}

		return buildClient(ctx, p, authConfig, partnerId)
	}
}

func buildClient(ctx context.Context, p *schema.Provider, authConfig *auth.Credentials, partnerId string) (*clients.Client, pluginsdk.Diagnostics) {
	clientBuilder := clients.ClientBuilder{
		AuthConfig:       authConfig,
		PartnerID:        partnerId,
		TerraformVersion: p.TerraformVersion,
	}

	stopCtx, ok := schema.StopContext(ctx) //nolint:staticcheck
	if !ok {
		stopCtx = ctx
	}

	client, err := clientBuilder.Build(stopCtx)
	if err != nil {
		return nil, pluginsdk.DiagFromErr(err)
	}

	return client, nil
}
