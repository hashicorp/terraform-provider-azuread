package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/go-version"
	"golang.org/x/oauth2"
)

const (
	azureCliMinimumVersion   = "2.0.81"
	azureCliNextMajorVersion = "3.0.0"
)

// AzureCliAuthorizer is an Authorizer which supports the Azure CLI.
type AzureCliAuthorizer struct {
	// TenantID is optional and forces selection of the specified tenant. Must be a valid UUID.
	TenantID string

	ctx  context.Context
	conf *AzureCliConfig
}

// Token returns an access token using the Azure CLI as an authentication mechanism.
func (a AzureCliAuthorizer) Token() (*oauth2.Token, error) {
	// We don't need to handle token caching and refreshing since az-cli does that for us
	var token struct {
		AccessToken string `json:"accessToken"`
		ExpiresOn   string `json:"expiresOn"`
		Tenant      string `json:"tenant"`
		TokenType   string `json:"tokenType"`
	}
	var resourceType string
	switch a.conf.Api {
	case MsGraph:
		resourceType = "ms-graph"
	case AadGraph:
		resourceType = "aad-graph"
	}
	err := jsonUnmarshalAzCmd(&token, "account", "get-access-token", fmt.Sprintf("--resource-type=%s", resourceType), "--tenant", a.conf.TenantID)
	if err != nil {
		return nil, err
	}

	return &oauth2.Token{
		AccessToken: token.AccessToken,
		TokenType:   token.TokenType,
		Expiry:      time.Time{},
	}, nil
}

// AzureCliConfig configures an AzureCliAuthorizer.
type AzureCliConfig struct {
	Api      Api
	TenantID string
}

// NewAzureCliConfig validates the supplied tenant ID and returns a new AzureCliConfig.
func NewAzureCliConfig(api Api, tenantId string) (*AzureCliConfig, error) {
	var err error

	// check az-cli version
	if err = checkAzVersion(); err != nil {
		return nil, err
	}

	// check tenant id
	tenantId, err = checkTenantId(tenantId)
	if err != nil {
		return nil, err
	}
	if tenantId == "" {
		return nil, errors.New("invalid tenantId or unable to determine tenantId")
	}

	return &AzureCliConfig{Api: api, TenantID: tenantId}, nil
}

// TokenSource provides a source for obtaining access tokens using AzureCliAuthorizer.
func (c *AzureCliConfig) TokenSource(ctx context.Context) Authorizer {
	return &AzureCliAuthorizer{
		TenantID: c.TenantID,
		ctx:      ctx,
		conf:     c,
	}
}

// checkAzVersion tries to determine the version of Azure CLI in the path and checks for a compatible version
func checkAzVersion() error {
	var cliVersion *struct {
		AzureCli          *string      `json:"azure-cli,omitempty"`
		AzureCliCore      *string      `json:"azure-cli-core,omitempty"`
		AzureCliTelemetry *string      `json:"azure-cli-telemetry,omitempty"`
		Extensions        *interface{} `json:"extensions,omitempty"`
	}
	err := jsonUnmarshalAzCmd(&cliVersion, "version")
	if err != nil {
		return fmt.Errorf("could not parse Azure CLI version: %v", err)
	}

	if cliVersion.AzureCli == nil {
		return fmt.Errorf("could not detect Azure CLI version. Please ensure you have installed Azure CLI version %s or newer", azureCliMinimumVersion)
	}

	actual, err := version.NewVersion(*cliVersion.AzureCli)
	if err != nil {
		return fmt.Errorf("could not parse detected Azure CLI version %q: %+v", *cliVersion.AzureCli, err)
	}

	supported, err := version.NewVersion(azureCliMinimumVersion)
	if err != nil {
		return fmt.Errorf("could not parse supported Azure CLI version: %+v", err)
	}

	nextMajor, err := version.NewVersion(azureCliNextMajorVersion)
	if err != nil {
		return fmt.Errorf("could not parse next major Azure CLI version: %+v", err)
	}

	if nextMajor.LessThanOrEqual(actual) {
		return fmt.Errorf("unsupported Azure CLI version %q detected, please install a version newer than %s but older than %s", actual, supported, nextMajor)
	}

	if actual.LessThan(supported) {
		return fmt.Errorf("unsupported Azure CLI version %q detected, please install version %s or newer and ensure the `az` command is in your path", actual, supported)
	}

	return nil
}

// checkTenantId validates the supplied tenant ID, and tries to determine the default tenant if a valid one is not supplied.
func checkTenantId(tenantId string) (string, error) {
	validTenantId, err := regexp.MatchString("^[a-zA-Z0-9._-]+$", tenantId)
	if err != nil {
		return "", fmt.Errorf("could not parse tenant ID %q: %s", tenantId, err)
	}

	if !validTenantId {
		var account struct {
			ID       string `json:"id"`
			TenantID string `json:"tenantId"`
		}
		err := jsonUnmarshalAzCmd(&account, "account", "show")
		if err != nil {
			return "", fmt.Errorf("obtaining tenant ID: %s", err)
		}
		tenantId = account.TenantID
	}

	return tenantId, nil
}

// jsonUnmarshalAzCmd executes an Azure CLI command and unmarshals the JSON output.
func jsonUnmarshalAzCmd(i interface{}, arg ...string) error {
	var stderr bytes.Buffer
	var stdout bytes.Buffer

	arg = append(arg, "-o=json")
	cmd := exec.Command("az", arg...)
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	if err := cmd.Start(); err != nil {
		err := fmt.Errorf("launching Azure CLI: %+v", err)
		if stdErrStr := stderr.String(); stdErrStr != "" {
			err = fmt.Errorf("%s: %s", err, strings.TrimSpace(stdErrStr))
		}
		return err
	}

	if err := cmd.Wait(); err != nil {
		err := fmt.Errorf("running Azure CLI: %+v", err)
		if stdErrStr := stderr.String(); stdErrStr != "" {
			err = fmt.Errorf("%s: %s", err, strings.TrimSpace(stdErrStr))
		}
		return err
	}

	if err := json.Unmarshal(stdout.Bytes(), &i); err != nil {
		return fmt.Errorf("unmarshaling the output of Azure CLI: %v", err)
	}

	return nil
}
