package base

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/manicminer/hamilton/auth"
	"github.com/manicminer/hamilton/base/odata"
	"github.com/manicminer/hamilton/environments"
)

type ApiVersion string

const (
	Version10   ApiVersion = "v1.0"
	VersionBeta ApiVersion = "beta"
)

const (
	defaultInitialBackoff = 1 * time.Second
	defaultBackoffCap     = 64 * time.Second
	requestAttempts       = 10
)

// ValidStatusFunc is a function that tests whether an HTTP response is considered valid for the particular request.
type ValidStatusFunc func(response *http.Response, o *odata.OData) bool

// HttpRequestInput is any type that can validate the response to an HTTP request.
type HttpRequestInput interface {
	GetValidStatusCodes() []int
	GetValidStatusFunc() ValidStatusFunc
}

// Uri represents a Microsoft Graph endpoint.
type Uri struct {
	Entity      string
	Params      url.Values
	HasTenantId bool
}

// GraphClient is any suitable HTTP client.
type GraphClient = *http.Client

// Client is a base client to be used by clients for specific entities.
// It can send GET, POST, PUT, PATCH and DELETE requests to Microsoft Graph and is API version and tenant aware.
type Client struct {
	// Endpoint is the base endpoint for Microsoft Graph, usually "https://graph.microsoft.com".
	Endpoint environments.ApiEndpoint

	// ApiVersion is the Microsoft Graph API version to use.
	ApiVersion ApiVersion

	// TenantId is the tenant ID to use in requests.
	TenantId string

	// UserAgent is the HTTP user agent string to send in requests.
	UserAgent string

	// Authorizer is anything that can provide an access token with which to authorize requests.
	Authorizer auth.Authorizer

	httpClient GraphClient
}

// NewClient returns a new Client configured with the specified API version and tenant ID.
func NewClient(apiVersion ApiVersion, tenantId string) Client {
	return Client{
		Endpoint:   environments.MsGraphGlobal.Endpoint,
		ApiVersion: apiVersion,
		TenantId:   tenantId,
		httpClient: http.DefaultClient,
	}
}

// buildUri is used by the package to build a complete URI string for API requests.
func (c Client) buildUri(uri Uri) (string, error) {
	newUrl, err := url.Parse(string(c.Endpoint))
	if err != nil {
		return "", err
	}
	newUrl.Path = "/" + string(c.ApiVersion)
	if uri.HasTenantId {
		newUrl.Path = fmt.Sprintf("%s/%s", newUrl.Path, c.TenantId)
	}
	newUrl.Path = fmt.Sprintf("%s/%s", newUrl.Path, strings.TrimLeft(uri.Entity, "/"))
	if uri.Params != nil {
		newUrl.RawQuery = uri.Params.Encode()
	}
	return newUrl.String(), nil
}

// performRequest is used by the package to send an HTTP request to the API.
func (c Client) performRequest(req *http.Request, input HttpRequestInput) (*http.Response, int, *odata.OData, error) {
	var status int

	if c.Authorizer != nil {
		token, err := c.Authorizer.Token()
		if err != nil {
			return nil, status, nil, err
		}
		token.SetAuthHeader(req)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	if c.UserAgent != "" {
		req.Header.Add("User-Agent", c.UserAgent)
	}

	var resp *http.Response
	var o odata.OData
	var err error

	var backoffPower func(int64, int64) int64
	backoffPower = func(base, exp int64) int64 {
		if exp <= 1 {
			return base
		}
		return base * backoffPower(base, exp-1)
	}

	var attempts, backoff, multiplier int64
	for attempts = 0; attempts < requestAttempts; attempts++ {
		// sleep after the previous failed attempt
		if attempts > 0 {
			time.Sleep(time.Duration(backoff))
		}

		// default exponential backoff
		multiplier++
		backoff = int64(defaultInitialBackoff) * backoffPower(2, multiplier)
		if cap := int64(defaultBackoffCap); backoff > cap {
			backoff = cap
		}

		resp, err = c.httpClient.Do(req)
		if err != nil {
			return nil, status, nil, err
		}

		d, err := odata.FromResponse(resp)
		if err != nil {
			return nil, status, &o, err
		}
		if d != nil {
			o = *d
		}

		status = resp.StatusCode
		if !containsStatusCode(input.GetValidStatusCodes(), status) {
			f := input.GetValidStatusFunc()
			if f != nil && f(resp, &o) {
				return resp, status, &o, nil
			}

			// rate limiting
			if containsStatusCode([]int{424, 429, 503}, status) {
				if retryAfter := resp.Header.Get("Retry-After"); retryAfter != "" {
					if r, err := strconv.ParseFloat(retryAfter, 64); err == nil && r > 0 {
						// Retry-After header detected, use that instead of default backoff
						backoff = int64(r * float64(time.Second))
						multiplier = 0
					}
				}
				continue
			}

			var errText string
			switch {
			case o.Error != nil && o.Error.String() != "":
				errText = fmt.Sprintf("OData error: %s", o.Error)
			default:
				defer resp.Body.Close()
				respBody, _ := ioutil.ReadAll(resp.Body)
				errText = fmt.Sprintf("response: %s", respBody)
			}
			return nil, status, &o, fmt.Errorf("unexpected status %d with %s", resp.StatusCode, errText)
		}

		break
	}

	return resp, status, &o, nil
}

// containsStatusCode determines whether the returned status code is in the []int of expected status codes.
func containsStatusCode(expected []int, actual int) bool {
	for _, v := range expected {
		if actual == v {
			return true
		}
	}

	return false
}
