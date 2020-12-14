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

type AzureCliAuthorizer struct {
	ctx      context.Context
	conf     *AzureCliConfig
	TenantID string
}

func (a AzureCliAuthorizer) Token() (*oauth2.Token, error) {
	// We don't need to handle token caching and refreshing since az-cli does that for us
	var token struct {
		AccessToken string `json:"accessToken"`
		ExpiresOn   string `json:"expiresOn"`
		Tenant      string `json:"tenant"`
		TokenType   string `json:"tokenType"`
	}
	cmd := []string{"account", "get-access-token", "--resource-type=ms-graph", "--tenant", a.TenantID, "-o=json"}
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

type AzureCliConfig struct {
	TenantID string
}

func (c *AzureCliConfig) TokenSource(ctx context.Context) Authorizer {
	var tenantId string
	if validTenantId, err := regexp.MatchString("^[a-zA-Z0-9._-]+$", c.TenantID); err == nil && validTenantId {
		tenantId = c.TenantID
	} else {
		var account struct {
			ID       string `json:"id"`
			TenantID string `json:"tenantId"`
		}
		cmd := []string{"account", "show", "-o=json"}
		err := jsonUnmarshalAzCmd(&account, cmd...)
		if err == nil {
			tenantId = account.TenantID
		}
	}

	return &AzureCliAuthorizer{
		ctx:      ctx,
		conf:     c,
		TenantID: tenantId,
	}
}

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
		err := fmt.Errorf("waiting for the Azure CLI: %+v", err)
		if stdErrStr := stderr.String(); stdErrStr != "" {
			err = fmt.Errorf("%s: %s", err, strings.TrimSpace(stdErrStr))
		}
		return err
	}

	if err := json.Unmarshal([]byte(stdout.String()), &i); err != nil {
		return fmt.Errorf("unmarshaling the result of Azure CLI: %v", err)
	}

	return nil
}
