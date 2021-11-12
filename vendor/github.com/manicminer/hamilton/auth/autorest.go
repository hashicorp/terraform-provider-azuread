package auth

import (
	"fmt"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
	"golang.org/x/oauth2"
)

// AutorestAuthorizerWrapper is an Authorizer which sources tokens from an autorest.BearerAuthorizer
type AutorestAuthorizerWrapper struct {
	bearerAuthorizer *autorest.BearerAuthorizer
}

// Token returns an access token using an autorest.BearerAuthorizer struct
func (a AutorestAuthorizerWrapper) Token() (*oauth2.Token, error) {
	tokenProvider := a.bearerAuthorizer.TokenProvider()
	if refresher, ok := tokenProvider.(adal.Refresher); ok {
		if err := refresher.EnsureFresh(); err != nil {
			return nil, err
		}
	}

	var adalToken adal.Token
	if spToken, ok := tokenProvider.(*adal.ServicePrincipalToken); ok {
		adalToken = spToken.Token()
	}

	if adalToken.AccessToken == "" {
		return nil, fmt.Errorf("could not obtain access token via supplied autorest.BearerAuthorizer")
	}

	return &oauth2.Token{
		AccessToken:  adalToken.AccessToken,
		TokenType:    adalToken.Type,
		RefreshToken: adalToken.RefreshToken,
		Expiry:       adalToken.Expires(),
	}, nil
}
