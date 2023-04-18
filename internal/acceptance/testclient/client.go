package testclient

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/hashicorp/go-azure-sdk/sdk/auth"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
)

var (
	_client    *clients.Client
	clientLock = &sync.Mutex{}
)

func Build(tenantId string) (*clients.Client, error) {
	clientLock.Lock()
	defer clientLock.Unlock()

	if _client == nil {
		var (
			ctx          = context.Background()
			metadataHost = os.Getenv("ARM_METADATA_HOSTNAME")

			env *environments.Environment
			err error
		)

		envName, exists := os.LookupEnv("ARM_ENVIRONMENT")
		if !exists {
			envName = "public"
		}

		if metadataHost != "" {
			if env, err = environments.FromEndpoint(ctx, fmt.Sprintf("https://%s", metadataHost), envName); err != nil {
				return nil, fmt.Errorf("building test client: %+v", err)
			}
		} else if env, err = environments.FromName(envName); err != nil {
			return nil, fmt.Errorf("building test client: %+v", err)
		}

		if tenantId == "" {
			tenantId = os.Getenv("ARM_TENANT_ID")
		}

		authConfig := auth.Credentials{
			Environment: *env,
			ClientID:    os.Getenv("ARM_CLIENT_ID"),
			TenantID:    tenantId,

			ClientCertificatePath:     os.Getenv("ARM_CLIENT_CERTIFICATE_PATH"),
			ClientCertificatePassword: os.Getenv("ARM_CLIENT_CERTIFICATE_PASSWORD"),
			ClientSecret:              os.Getenv("ARM_CLIENT_SECRET"),

			EnableAuthenticatingUsingClientCertificate: true,
			EnableAuthenticatingUsingClientSecret:      true,
			EnableAuthenticatingUsingAzureCLI:          false,
			EnableAuthenticatingUsingManagedIdentity:   false,
			EnableAuthenticationUsingOIDC:              false,
			EnableAuthenticationUsingGitHubOIDC:        false,
		}

		builder := clients.ClientBuilder{
			AuthConfig:       &authConfig,
			TerraformVersion: os.Getenv("TERRAFORM_CORE_VERSION"),
		}

		client, err := builder.Build(ctx)
		if err != nil {
			return nil, fmt.Errorf("building test client: %+v", err)
		}

		_client = client
	}

	return _client, nil
}
