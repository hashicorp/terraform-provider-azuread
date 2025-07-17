// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/go-azure-sdk/sdk/auth"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-azuread/internal/clients"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
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
	t.Skip("skipping test for now, as it requires Azure CLI authentication to be set up in the environment which if not compatible with MFA Requirement")

	if os.Getenv("TF_ACC") == "" {
		t.Skip("TF_ACC not set")
	}

	provider := AzureADProvider()
	ctx := context.Background()

	// Support only Azure CLI authentication
	provider.ConfigureContextFunc = func(ctx context.Context, d *pluginsdk.ResourceData) (interface{}, pluginsdk.Diagnostics) {
		envName := d.Get("environment").(string)
		env, err := environments.FromName(envName)
		if err != nil {
			t.Fatalf("configuring environment %q: %v", envName, err)
		}

		tenantId, err := getTenantId(d)
		if err != nil {
			return nil, pluginsdk.DiagFromErr(err)
		}

		authConfig := &auth.Credentials{
			Environment:                       *env,
			TenantID:                          *tenantId,
			EnableAuthenticatingUsingAzureCLI: true,
		}

		return buildClient(ctx, provider, authConfig, "")
	}

	d := provider.Configure(ctx, terraform.NewResourceConfigRaw(nil))
	if d != nil && d.HasError() {
		t.Fatalf("err: %+v", d)
	}

	if errs := testCheckProvider(provider); len(errs) > 0 {
		for _, err := range errs {
			t.Error(err)
		}
	}
}

func TestAccProvider_clientCertificateAuth(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("TF_ACC not set")
	}

	provider := AzureADProvider()
	ctx := context.Background()

	// Support only client certificate authentication
	provider.ConfigureContextFunc = func(ctx context.Context, d *pluginsdk.ResourceData) (interface{}, pluginsdk.Diagnostics) {
		envName := d.Get("environment").(string)
		env, err := environments.FromName(envName)
		if err != nil {
			t.Fatalf("configuring environment %q: %v", envName, err)
		}

		clientId, err := getClientId(d)
		if err != nil {
			return nil, pluginsdk.DiagFromErr(err)
		}

		tenantId, err := getTenantId(d)
		if err != nil {
			return nil, pluginsdk.DiagFromErr(err)
		}

		authConfig := &auth.Credentials{
			Environment:               *env,
			TenantID:                  *tenantId,
			ClientID:                  *clientId,
			ClientCertificatePath:     d.Get("client_certificate_path").(string),
			ClientCertificatePassword: d.Get("client_certificate_password").(string),
			EnableAuthenticatingUsingClientCertificate: true,
		}

		return buildClient(ctx, provider, authConfig, "")
	}

	d := provider.Configure(ctx, terraform.NewResourceConfigRaw(nil))
	if d != nil && d.HasError() {
		t.Fatalf("err: %+v", d)
	}

	if errs := testCheckProvider(provider); len(errs) > 0 {
		for _, err := range errs {
			t.Error(err)
		}
	}
}

func TestAccProvider_clientCertificateInlineAuth(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("TF_ACC not set")
	}

	provider := AzureADProvider()
	ctx := context.Background()

	// Support only client certificate authentication
	provider.ConfigureContextFunc = func(ctx context.Context, d *pluginsdk.ResourceData) (interface{}, pluginsdk.Diagnostics) {
		var certData []byte
		if encodedCert := d.Get("client_certificate").(string); encodedCert != "" {
			var err error
			certData, err = decodeCertificate(encodedCert)
			if err != nil {
				return nil, pluginsdk.DiagFromErr(err)
			}
		}

		envName := d.Get("environment").(string)
		env, err := environments.FromName(envName)
		if err != nil {
			t.Fatalf("configuring environment %q: %v", envName, err)
		}

		clientId, err := getClientId(d)
		if err != nil {
			return nil, pluginsdk.DiagFromErr(err)
		}

		tenantId, err := getTenantId(d)
		if err != nil {
			return nil, pluginsdk.DiagFromErr(err)
		}

		authConfig := &auth.Credentials{
			Environment:               *env,
			TenantID:                  *tenantId,
			ClientID:                  *clientId,
			ClientCertificateData:     certData,
			ClientCertificatePassword: d.Get("client_certificate_password").(string),
			EnableAuthenticatingUsingClientCertificate: true,
		}

		return buildClient(ctx, provider, authConfig, "")
	}

	d := provider.Configure(ctx, terraform.NewResourceConfigRaw(nil))
	if d != nil && d.HasError() {
		t.Fatalf("err: %+v", d)
	}

	if errs := testCheckProvider(provider); len(errs) > 0 {
		for _, err := range errs {
			t.Error(err)
		}
	}
}

func TestAccProvider_clientSecretAuth(t *testing.T) {
	t.Run("fromEnvironment", testAccProvider_clientSecretAuthFromEnvironment)
	t.Run("fromFiles", testAccProvider_clientSecretAuthFromFiles)
}

func testAccProvider_clientSecretAuthFromEnvironment(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("TF_ACC not set")
	}
	if os.Getenv("ARM_CLIENT_ID") == "" {
		t.Skip("ARM_CLIENT_ID not set")
	}
	if os.Getenv("ARM_CLIENT_SECRET") == "" {
		t.Skip("ARM_CLIENT_SECRET not set")
	}

	// Ensure we are running using the expected env-vars
	// t.SetEnv does automatic cleanup / resets the values after the test
	t.Setenv("ARM_CLIENT_ID_FILE_PATH", "")
	t.Setenv("ARM_CLIENT_SECRET_FILE_PATH", "")

	provider := AzureADProvider()
	ctx := context.Background()

	// Support only client secret authentication
	provider.ConfigureContextFunc = func(ctx context.Context, d *pluginsdk.ResourceData) (interface{}, pluginsdk.Diagnostics) {
		envName := d.Get("environment").(string)
		env, err := environments.FromName(envName)
		if err != nil {
			t.Fatalf("configuring environment %q: %v", envName, err)
		}

		clientId, err := getClientId(d)
		if err != nil {
			return nil, pluginsdk.DiagFromErr(err)
		}

		clientSecret, err := getClientSecret(d)
		if err != nil {
			return nil, pluginsdk.DiagFromErr(err)
		}

		tenantId, err := getTenantId(d)
		if err != nil {
			return nil, pluginsdk.DiagFromErr(err)
		}

		authConfig := &auth.Credentials{
			Environment:                           *env,
			TenantID:                              *tenantId,
			ClientID:                              *clientId,
			ClientSecret:                          *clientSecret,
			EnableAuthenticatingUsingClientSecret: true,
		}

		return buildClient(ctx, provider, authConfig, "")
	}

	d := provider.Configure(ctx, terraform.NewResourceConfigRaw(nil))
	if d != nil && d.HasError() {
		t.Fatalf("err: %+v", d)
	}

	if errs := testCheckProvider(provider); len(errs) > 0 {
		for _, err := range errs {
			t.Error(err)
		}
	}
}

func testAccProvider_clientSecretAuthFromFiles(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("TF_ACC not set")
	}
	if os.Getenv("ARM_CLIENT_ID_FILE_PATH") == "" {
		t.Skip("ARM_CLIENT_ID_FILE_PATH not set")
	}
	if os.Getenv("ARM_CLIENT_SECRET_FILE_PATH") == "" {
		t.Skip("ARM_CLIENT_SECRET_FILE_PATH not set")
	}

	// Ensure we are running using the expected env-vars
	// t.SetEnv does automatic cleanup / resets the values after the test
	t.Setenv("ARM_CLIENT_ID", "")
	t.Setenv("ARM_CLIENT_SECRET", "")

	provider := AzureADProvider()
	ctx := context.Background()

	// Support only client secret authentication
	provider.ConfigureContextFunc = func(ctx context.Context, d *pluginsdk.ResourceData) (interface{}, pluginsdk.Diagnostics) {
		envName := d.Get("environment").(string)
		env, err := environments.FromName(envName)
		if err != nil {
			t.Fatalf("configuring environment %q: %v", envName, err)
		}

		clientId, err := getClientId(d)
		if err != nil {
			return nil, pluginsdk.DiagFromErr(err)
		}

		clientSecret, err := getClientSecret(d)
		if err != nil {
			return nil, pluginsdk.DiagFromErr(err)
		}

		tenantId, err := getTenantId(d)
		if err != nil {
			return nil, pluginsdk.DiagFromErr(err)
		}

		authConfig := &auth.Credentials{
			Environment:                           *env,
			TenantID:                              *tenantId,
			ClientID:                              *clientId,
			ClientSecret:                          *clientSecret,
			EnableAuthenticatingUsingClientSecret: true,
		}

		return buildClient(ctx, provider, authConfig, "")
	}

	d := provider.Configure(ctx, terraform.NewResourceConfigRaw(nil))
	if d != nil && d.HasError() {
		t.Fatalf("err: %+v", d)
	}

	if errs := testCheckProvider(provider); len(errs) > 0 {
		for _, err := range errs {
			t.Error(err)
		}
	}
}

func TestAccProvider_genericOidcAuth(t *testing.T) {
	t.Run("fromEnvironment", testAccProvider_genericOidcAuthFromEnvironment)
	t.Run("fromFiles", testAccProvider_genericOidcAuthFromFiles)
}

func testAccProvider_genericOidcAuthFromEnvironment(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("TF_ACC not set")
	}
	if os.Getenv("ARM_OIDC_TOKEN_FILE_PATH") == "" {
		t.Skip("ARM_OIDC_TOKEN_FILE_PATH not set")
	}

	// Ensure we are running using the expected env-vars
	// t.SetEnv does automatic cleanup / resets the values after the test
	t.Setenv("ARM_OIDC_TOKEN", "")

	provider := AzureADProvider()
	ctx := context.Background()

	// Support only oidc authentication
	provider.ConfigureContextFunc = func(ctx context.Context, d *pluginsdk.ResourceData) (interface{}, pluginsdk.Diagnostics) {
		envName := d.Get("environment").(string)
		env, err := environments.FromName(envName)
		if err != nil {
			t.Fatalf("configuring environment %q: %v", envName, err)
		}

		idToken, err := getOidcToken(d)
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

		authConfig := &auth.Credentials{
			Environment:                   *env,
			TenantID:                      *tenantId,
			ClientID:                      *clientId,
			OIDCAssertionToken:            *idToken,
			EnableAuthenticationUsingOIDC: true,
		}

		return buildClient(ctx, provider, authConfig, "")
	}

	d := provider.Configure(ctx, terraform.NewResourceConfigRaw(nil))
	if d != nil && d.HasError() {
		t.Fatalf("err: %+v", d)
	}

	if errs := testCheckProvider(provider); len(errs) > 0 {
		for _, err := range errs {
			t.Error(err)
		}
	}
}

func testAccProvider_genericOidcAuthFromFiles(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("TF_ACC not set")
	}
	if os.Getenv("ARM_OIDC_TOKEN") == "" {
		t.Skip("ARM_OIDC_TOKEN not set")
	}

	// Ensure we are running using the expected env-vars
	// t.SetEnv does automatic cleanup / resets the values after the test
	t.Setenv("ARM_OIDC_TOKEN_FILE_PATH", "")

	provider := AzureADProvider()
	ctx := context.Background()

	// Support only oidc authentication
	provider.ConfigureContextFunc = func(ctx context.Context, d *pluginsdk.ResourceData) (interface{}, pluginsdk.Diagnostics) {
		envName := d.Get("environment").(string)
		env, err := environments.FromName(envName)
		if err != nil {
			t.Fatalf("configuring environment %q: %v", envName, err)
		}

		idToken, err := getOidcToken(d)
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

		authConfig := &auth.Credentials{
			Environment:                   *env,
			TenantID:                      *tenantId,
			ClientID:                      *clientId,
			OIDCAssertionToken:            *idToken,
			EnableAuthenticationUsingOIDC: true,
		}

		return buildClient(ctx, provider, authConfig, "")
	}

	d := provider.Configure(ctx, terraform.NewResourceConfigRaw(nil))
	if d != nil && d.HasError() {
		t.Fatalf("err: %+v", d)
	}

	if errs := testCheckProvider(provider); len(errs) > 0 {
		for _, err := range errs {
			t.Error(err)
		}
	}
}

func TestAccProvider_githubOidcAuth(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("TF_ACC not set")
	}
	if os.Getenv("ACTIONS_ID_TOKEN_REQUEST_URL") == "" {
		t.Skip("ACTIONS_ID_TOKEN_REQUEST_URL not set")
	}
	if os.Getenv("ACTIONS_ID_TOKEN_REQUEST_TOKEN") == "" {
		t.Skip("ACTIONS_ID_TOKEN_REQUEST_TOKEN not set")
	}

	provider := AzureADProvider()
	ctx := context.Background()

	// Support only oidc authentication
	provider.ConfigureContextFunc = func(ctx context.Context, d *pluginsdk.ResourceData) (interface{}, pluginsdk.Diagnostics) {
		envName := d.Get("environment").(string)
		env, err := environments.FromName(envName)
		if err != nil {
			t.Fatalf("configuring environment %q: %v", envName, err)
		}

		clientId, err := getClientId(d)
		if err != nil {
			return nil, pluginsdk.DiagFromErr(err)
		}

		tenantId, err := getTenantId(d)
		if err != nil {
			return nil, pluginsdk.DiagFromErr(err)
		}

		authConfig := &auth.Credentials{
			Environment:                         *env,
			TenantID:                            *tenantId,
			ClientID:                            *clientId,
			OIDCTokenRequestToken:               d.Get("oidc_request_token").(string),
			OIDCTokenRequestURL:                 d.Get("oidc_request_url").(string),
			EnableAuthenticationUsingGitHubOIDC: true,
		}

		return buildClient(ctx, provider, authConfig, "")
	}

	d := provider.Configure(ctx, terraform.NewResourceConfigRaw(nil))
	if d != nil && d.HasError() {
		t.Fatalf("err: %+v", d)
	}

	if errs := testCheckProvider(provider); len(errs) > 0 {
		for _, err := range errs {
			t.Error(err)
		}
	}
}

func TestAccProvider_adoOidcAuth(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		t.Skip("TF_ACC not set")
	}
	if os.Getenv("SYSTEM_ACCESSTOKEN") == "" {
		t.Skip("SYSTEM_ACCESSTOKEN not set")
	}
	if os.Getenv("SYSTEM_OIDCREQUESTURI") == "" {
		t.Skip("SYSTEM_OIDCREQUESTURI not set")
	}
	if os.Getenv("ARM_ADO_PIPELINE_SERVICE_CONNECTION_ID") == "" {
		t.Skip("ARM_ADO_PIPELINE_SERVICE_CONNECTION_ID")
	}

	provider := AzureADProvider()
	ctx := context.Background()

	// Support only oidc authentication
	provider.ConfigureContextFunc = func(ctx context.Context, d *pluginsdk.ResourceData) (interface{}, pluginsdk.Diagnostics) {
		envName := d.Get("environment").(string)
		env, err := environments.FromName(envName)
		if err != nil {
			t.Fatalf("configuring environment %q: %v", envName, err)
		}

		clientId, err := getClientId(d)
		if err != nil {
			return nil, pluginsdk.DiagFromErr(err)
		}

		tenantId, err := getTenantId(d)
		if err != nil {
			return nil, pluginsdk.DiagFromErr(err)
		}

		authConfig := &auth.Credentials{
			Environment:                              *env,
			TenantID:                                 *tenantId,
			ClientID:                                 *clientId,
			OIDCTokenRequestToken:                    d.Get("oidc_request_token").(string),
			OIDCTokenRequestURL:                      d.Get("oidc_request_url").(string),
			ADOPipelineServiceConnectionID:           d.Get("ado_pipeline_service_connection_id").(string),
			EnableAuthenticationUsingADOPipelineOIDC: true,
		}

		return buildClient(ctx, provider, authConfig, "")
	}

	d := provider.Configure(ctx, terraform.NewResourceConfigRaw(nil))
	if d != nil && d.HasError() {
		t.Fatalf("err: %+v", d)
	}

	if errs := testCheckProvider(provider); len(errs) > 0 {
		for _, err := range errs {
			t.Error(err)
		}
	}
}

func testCheckProvider(provider *schema.Provider) (errs []error) {
	client := provider.Meta().(*clients.Client)

	if endpoint, ok := client.Environment.MicrosoftGraph.Endpoint(); !ok || *endpoint == "" {
		errs = append(errs, fmt.Errorf("MsGraphEndpoint was empty in client.Environment"))
	}

	if client.ClientID == "" {
		errs = append(errs, fmt.Errorf("client.ClientID was empty"))
	}

	if client.TenantID == "" {
		errs = append(errs, fmt.Errorf("client.TenantID was empty"))
	}

	if client.ObjectID == "" {
		errs = append(errs, fmt.Errorf("client.ObjectID was empty"))
	}

	if client.Claims.TenantId == "" {
		errs = append(errs, fmt.Errorf("TenantId was not populated in client.Claims"))
	}

	if client.Claims.ObjectId == "" {
		errs = append(errs, fmt.Errorf("ObjectId was not populated in client.Claims"))
	}

	return //nolint:nakedret
}
