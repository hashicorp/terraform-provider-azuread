package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"golang.org/x/oauth2"
)

const (
	msiDefaultApiVersion = "2018-02-01"
	msiDefaultEndpoint   = "http://169.254.169.254/metadata/identity/oauth2/token"
	msiDefaultTimeout    = 10 * time.Second
)

// MsiAuthorizer is an Authorizer which supports managed service identity.
type MsiAuthorizer struct {
	ctx  context.Context
	conf *MsiConfig
}

// Token returns an access token acquired from the metadata endpoint.
func (a *MsiAuthorizer) Token() (*oauth2.Token, error) {
	query := url.Values{
		"api-version": []string{a.conf.MsiApiVersion},
		"resource":    []string{a.conf.Resource},
	}
	url := fmt.Sprintf("%s?%s", a.conf.MsiEndpoint, query.Encode())

	body, err := azureMetadata(a.ctx, url)
	if err != nil {
		return nil, fmt.Errorf("MsiAuthorizer: failed to request token from metadata endpoint: %v", err)
	}

	// TODO: surface the client ID for use by callers
	var tokenRes struct {
		AccessToken  string      `json:"access_token"`
		ClientID     string      `json:"client_id"`
		Resource     string      `json:"resource"`
		TokenType    string      `json:"token_type"`
		ExpiresIn    interface{} `json:"expires_in"`     // relative seconds from now
		ExpiresOn    interface{} `json:"expires_on"`     // timestamp
		ExtExpiresIn interface{} `json:"ext_expires_in"` // relative seconds from now
	}
	if err := json.Unmarshal(body, &tokenRes); err != nil {
		return nil, fmt.Errorf("MsiAuthorizer: failed to unmarshal token: %v", err)
	}

	token := &oauth2.Token{
		AccessToken: tokenRes.AccessToken,
		TokenType:   tokenRes.TokenType,
	}

	var secs time.Duration
	if exp, ok := tokenRes.ExpiresIn.(string); ok && exp != "" {
		if v, err := strconv.Atoi(exp); err == nil {
			secs = time.Duration(v)
		}
	} else if exp, ok := tokenRes.ExpiresIn.(int64); ok {
		secs = time.Duration(exp)
	} else if exp, ok := tokenRes.ExpiresIn.(float64); ok {
		secs = time.Duration(exp)
	}
	if secs > 0 {
		token.Expiry = time.Now().Add(secs * time.Second)
	}

	return token, nil
}

// MsiConfig configures an MsiAuthorizer.
type MsiConfig struct {
	MsiApiVersion string
	MsiEndpoint   string
	Resource      string
}

// NewMsiConfig returns a new MsiConfig with a configured metadata endpoint and resource.
func NewMsiConfig(ctx context.Context, resource string, msiEndpoint string) (*MsiConfig, error) {
	endpoint := msiDefaultEndpoint
	if msiEndpoint != "" {
		endpoint = msiEndpoint
	}

	// validate the metadata endpoint
	e, err := url.Parse(endpoint)
	if err != nil {
		return nil, fmt.Errorf("NewMsiConfig: invalid MSI endpoint configured: %q", endpoint)
	}

	// determine the generic metadata URL and check if we can reach it
	e.Path = "/metadata"
	e.RawQuery = url.Values{
		"api-version": []string{msiDefaultApiVersion},
		"format":      []string{"text"},
	}.Encode()

	_, err = azureMetadata(ctx, e.String())
	if err != nil {
		return nil, fmt.Errorf("NewMsiConfig: could not validate MSI endpoint: %v", err)
	}

	return &MsiConfig{
		Resource:      resource,
		MsiApiVersion: msiDefaultApiVersion,
		MsiEndpoint:   endpoint,
	}, nil
}

// TokenSource provides a source for obtaining access tokens using MsiAuthorizer.
func (c *MsiConfig) TokenSource(ctx context.Context) Authorizer {
	return NewCachedAuthorizer(&MsiAuthorizer{ctx: ctx, conf: c})
}

func azureMetadata(ctx context.Context, url string) (body []byte, err error) {
	var req *http.Request
	req, err = http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return
	}
	req.Header = http.Header{
		"Metadata": []string{"true"},
	}
	client := &http.Client{
		Timeout: msiDefaultTimeout,
	}
	var resp *http.Response
	resp, err = client.Do(req)
	if err != nil {
		return
	}
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if c := resp.StatusCode; c < 200 || c > 299 {
		err = fmt.Errorf("received HTTP status %d", resp.StatusCode)
		return
	}
	return
}
