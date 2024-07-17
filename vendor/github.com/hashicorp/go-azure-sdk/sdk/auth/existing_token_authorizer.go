package auth

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/environments"
	"golang.org/x/oauth2"
)

type ExistingTokenAuthorizerOptions struct {
	//Api describes the Azure API being used
	Api environments.Api

	//Token is the pre-existing token to use.
	Token string
}

var _ Authorizer = &ExistingTokenAuthorizer{}

type ExistingTokenAuthorizer struct {
	conf *existingTokenConfig
}

func NewExistingTokenAuthorizer(ctx context.Context, options ExistingTokenAuthorizerOptions) (Authorizer, error) {
	resource, err := environments.Resource(options.Api)

	if err != nil {
		return nil, fmt.Errorf("determining resource for api %q: %+v", options.Api.Name(), err)
	}

	conf, err := newExistingTokenConfig(*resource, options.Token)

	if err != nil {
		return nil, err
	}

	return conf.TokenSource(ctx)
}

type existingTokenConfig struct {
	Resource string
	Token    string
}

func newExistingTokenConfig(resource string, token string) (*existingTokenConfig, error) {
	if resource == "" {
		return nil, fmt.Errorf("resource is required")
	}

	if token == "" {
		return nil, fmt.Errorf("token is required")
	}

	return &existingTokenConfig{
		Resource: resource,
		Token:    token,
	}, nil
}

func (c *existingTokenConfig) TokenSource(_ context.Context) (Authorizer, error) {
	return NewCachedAuthorizer(&ExistingTokenAuthorizer{
		conf: c,
	})
}

func (a *ExistingTokenAuthorizer) AuxiliaryTokens(_ context.Context, _ *http.Request) ([]*oauth2.Token, error) {
	// auxiliary tokens are not supported with existing authentication, so just return an empty slice
	return []*oauth2.Token{}, nil
}

func (a *ExistingTokenAuthorizer) Token(ctx context.Context, _ *http.Request) (*oauth2.Token, error) {
	if a.conf == nil {
		return nil, fmt.Errorf("could not request token: conf is nil")
	}

	claims, err := getJwtClaims(a.conf.Token)

	if err != nil {
		return nil, fmt.Errorf("could not get claims from token: %v", err)
	}

	if claims["exp"].(float64)-float64(time.Now().UTC().Unix()) < 60 {
		return nil, fmt.Errorf("token has expired")
	}

	return &oauth2.Token{
		AccessToken: a.conf.Token,
		TokenType:   "Bearer",
	}, nil
}

func getJwtClaims(token string) (map[string]interface{}, error) {
	tokenPayload := strings.Split(strings.Split(token, ".")[1], ".")[0]

	decodedPayload, err := base64.RawStdEncoding.DecodeString(tokenPayload)

	if err != nil {
		return nil, err
	}

	var tokenPayloadMap map[string]interface{}

	err = json.Unmarshal(decodedPayload, &tokenPayloadMap)

	if err != nil {
		return nil, err
	}

	return tokenPayloadMap, nil
}
