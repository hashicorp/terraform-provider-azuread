package microsoft

import (
	"context"
	"time"

	"golang.org/x/oauth2"
)

type AuthType int

const (
	AuthTypeAssertion AuthType = iota
	AuthTypeSecret
)

// Config is the configuration for using client credentials flow.
//
// For more information see:
// https://docs.microsoft.com/en-us/azure/active-directory/develop/v2-oauth2-client-creds-grant-flow#get-a-token
// https://docs.microsoft.com/en-us/azure/active-directory/develop/active-directory-certificate-credentials
type Config struct {
	// ClientID is the application's ID.
	ClientID string

	// ClientSecret is the application's secret.
	ClientSecret string

	// PrivateKey contains the contents of an RSA private key or the
	// contents of a PEM file that contains a private key. The provided
	// private key is used to sign JWT assertions.
	// PEM containers with a passphrase are not supported.
	// Use the following command to convert a PKCS 12 file into a PEM.
	//
	//    $ openssl pkcs12 -in key.p12 -out key.pem -nodes
	//
	PrivateKey []byte

	// Certificate contains the (optionally PEM encoded) X509 certificate registered
	// for the application with which you are authenticating.
	Certificate []byte

	// Resource specifies an API resource for which to request access (used for v1 tokens)
	Resource string

	// Scopes specifies a list of requested permission scopes (used for v2 tokens)
	Scopes []string

	// TokenURL is the token endpoint. Typically you can use the AzureADEndpoint
	// function to obtain this value, but it may change for non-public clouds.
	TokenURL string

	// Expires optionally specifies how long the token is valid for.
	Expires time.Duration

	// Audience optionally specifies the intended audience of the
	// request.  If empty, the value of TokenURL is used as the
	// intended audience.
	Audience string
}

// TokenSource returns a JWT TokenSource using the configuration
// in c and the HTTP client from the provided context.
func (c *Config) TokenSource(ctx context.Context, authType AuthType) (source oauth2.TokenSource) {
	switch authType {
	case AuthTypeAssertion:
		source = oauth2.ReuseTokenSource(nil, assertionSource{ctx, c})
	case AuthTypeSecret:
		source = oauth2.ReuseTokenSource(nil, secretSource{ctx, c})
	}
	return
}
