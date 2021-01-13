package microsoft

import (
	"context"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/jws" //nolint:staticcheck

)

type assertionSource struct {
	ctx  context.Context
	conf *Config
}

func (a assertionSource) Token() (*oauth2.Token, error) {
	crt := a.conf.Certificate
	if der, _ := pem.Decode(a.conf.Certificate); der != nil {
		crt = der.Bytes
	}
	cert, err := x509.ParseCertificate(crt)
	if err != nil {
		return nil, fmt.Errorf("oauth2: cannot parse certificate: %v", err)
	}
	s := sha1.Sum(cert.Raw)
	fp := base64.URLEncoding.EncodeToString(s[:])
	h := jws.Header{
		Algorithm: "RS256",
		Typ:       "JWT",
		KeyID:     fp,
	}

	claimSet := &jws.ClaimSet{
		Iss: a.conf.ClientID,
		Sub: a.conf.ClientID,
		Aud: a.conf.TokenURL,
	}
	if t := a.conf.Expires; t > 0 {
		claimSet.Exp = time.Now().Add(t).Unix()
	}
	if aud := a.conf.Audience; aud != "" {
		claimSet.Aud = aud
	}

	pk, err := parseKey(a.conf.PrivateKey)
	if err != nil {
		return nil, err
	}

	payload, err := jws.Encode(&h, claimSet, pk)
	if err != nil {
		return nil, err
	}

	hc := oauth2.NewClient(a.ctx, nil)
	v := url.Values{
		"client_assertion":      {payload},
		"client_assertion_type": {"urn:ietf:params:oauth:client-assertion-type:jwt-bearer"},
		"client_id":             {a.conf.ClientID},
		"grant_type":            {"client_credentials"},
	}
	if a.conf.Resource != "" {
		v["resource"] = []string{a.conf.Resource}
	} else {
		v["scope"] = []string{strings.Join(a.conf.Scopes, " ")}
	}
	resp, err := hc.PostForm(a.conf.TokenURL, v)
	if err != nil {
		return nil, fmt.Errorf("oauth2: cannot fetch token: %v", err)
	}

	return token(resp)
}

// parseKey returns an rsa.PrivateKey containing the provided binary key data.
// If the provided key is PEM encoded, it is decoded first.
func parseKey(key []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(key)
	if block != nil {
		key = block.Bytes
	}
	parsedKey, err := x509.ParsePKCS8PrivateKey(key)
	if err != nil {
		parsedKey, err = x509.ParsePKCS1PrivateKey(key)
		if err != nil {
			return nil, fmt.Errorf("private key should be a PEM or plain PKCS1 or PKCS8; parse error: %v", err)
		}
	}
	parsed, ok := parsedKey.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("private key is invalid")
	}
	return parsed, nil
}
