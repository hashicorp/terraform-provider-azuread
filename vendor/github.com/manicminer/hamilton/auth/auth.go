package auth

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"golang.org/x/crypto/pkcs12"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"golang.org/x/oauth2/endpoints"
	"golang.org/x/oauth2/microsoft"
	"io/ioutil"
	"strings"

	microsoft2 "github.com/manicminer/hamilton/auth/microsoft"
)

type Config struct {
	//Environment    string
	TenantID string
	ClientID string

	// Azure CLI Tokens Auth
	// TODO: NOT YET SUPPORTED
	EnableAzureCliToken bool

	// Managed Service Identity Auth
	// TODO: NOT YET SUPPORTED
	EnableMsiAuth bool
	MsiEndpoint   string

	// Service Principal (Client Cert) Auth
	EnableClientCertAuth bool
	ClientCertPath       string
	ClientCertPassword   string

	// Service Principal (Client Secret) Auth
	EnableClientSecretAuth bool
	ClientSecret           string
	ClientSecretDocsLink   string
}

type Authorizer interface {
	Token() (*oauth2.Token, error)
}

func (c *Config) NewAuthorizer(ctx context.Context) (Authorizer, error) {
	if c.EnableClientCertAuth && strings.TrimSpace(c.ClientID) != "" && strings.TrimSpace(c.ClientCertPath) != "" {
		a, err := NewClientCertificateAuthorizer(ctx, c.TenantID, c.ClientID, c.ClientCertPath, c.ClientCertPassword)
		if err != nil {
			return nil, fmt.Errorf("could not configure ClientCertificate Authorizer: %s", err)
		}
		if a != nil {
			return a, nil
		}
	}

	if c.EnableClientSecretAuth && strings.TrimSpace(c.ClientID) != "" && strings.TrimSpace(c.ClientSecret) != "" {
		a, err := NewClientSecretAuthorizer(ctx, c.TenantID, c.ClientID, c.ClientSecret)
		if err != nil {
			return nil, fmt.Errorf("could not configure ClientCertificate Authorizer: %s", err)
		}
		if a != nil {
			return a, nil
		}
	}

	return nil, fmt.Errorf("no Authorizer could be configured, please check your configuration")
}

func NewClientSecretAuthorizer(ctx context.Context, tenantId, clientId, clientSecret string) (Authorizer, error) {
	conf := clientcredentials.Config{
		AuthStyle:    oauth2.AuthStyleInParams,
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Scopes:       []string{"https://graph.microsoft.com/.default"},
		TokenURL:     endpoints.AzureAD(tenantId).TokenURL,
	}
	return conf.TokenSource(ctx), nil
}

func NewClientCertificateAuthorizer(ctx context.Context, tenantId, clientId, pfxPath, pfxPass string) (Authorizer, error) {
	pfx, err := ioutil.ReadFile(pfxPath)
	if err != nil {
		return nil, fmt.Errorf("could not read pkcs12 store at %q: %s", pfxPath, err)
	}

	key, cert, err := pkcs12.Decode(pfx, pfxPass)
	if err != nil {
		return nil, fmt.Errorf("could not decode pkcs12 credential store: %s", err)
	}

	priv, ok := key.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("unsupported non-rsa key was found in pkcs12 store %q", pfxPath)
	}

	conf := microsoft2.Config{
		ClientID:    clientId,
		PrivateKey:  x509.MarshalPKCS1PrivateKey(priv),
		Certificate: cert.Raw,
		Scopes:      []string{"https://graph.microsoft.com/.default"},
		TokenURL:    microsoft.AzureADEndpoint(tenantId).TokenURL,
	}
	return conf.TokenSource(ctx), nil
}
