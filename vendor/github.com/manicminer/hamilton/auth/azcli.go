package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"os/exec"
	"regexp"
	"strings"
	"time"
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
	cmd := []string{"account", "get-access-token", fmt.Sprintf("--resource-type=%s", resourceType), "--tenant", a.conf.TenantID, "-o=json"}
	err := jsonUnmarshalAzCmd(&token, cmd...)
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
	// check az-cli version

	// check tenant id
	validTenantId, err := regexp.MatchString("^[a-zA-Z0-9._-]+$", tenantId)
	if err != nil {
		return nil, fmt.Errorf("could not parse tenant ID %q: %s", tenantId, err)
	}

	if !validTenantId {
		var account struct {
			ID       string `json:"id"`
			TenantID string `json:"tenantId"`
		}
		cmd := []string{"account", "show", "-o=json"}
		err := jsonUnmarshalAzCmd(&account, cmd...)
		if err != nil {
			return nil, fmt.Errorf("obtaining tenant ID: %s", err)
		}
		tenantId = account.TenantID
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

// jsonUnmarshalAzCmd executes an Azure CLI command and unmarshals the JSON output.
func jsonUnmarshalAzCmd(i interface{}, arg ...string) error {
	var stderr bytes.Buffer
	var stdout bytes.Buffer

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
