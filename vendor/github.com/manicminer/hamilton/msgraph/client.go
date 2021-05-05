package msgraph

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/manicminer/hamilton/auth"
	"github.com/manicminer/hamilton/environments"
	"github.com/manicminer/hamilton/odata"
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

	httpClient *http.Client
}

// NewClient returns a new Client configured with the specified API version and tenant ID.
func NewClient(apiVersion ApiVersion, tenantId string) Client {
	return Client{
		Endpoint:   environments.MsGraphGlobal.Endpoint,
		ApiVersion: apiVersion,
		TenantId:   tenantId,
		UserAgent:  "Hamilton (Go-http-client/1.1)",
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
	var o *odata.OData
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

		o, err = odata.FromResponse(resp)
		if err != nil {
			return nil, status, o, err
		}

		status = resp.StatusCode
		if !containsStatusCode(input.GetValidStatusCodes(), status) {
			f := input.GetValidStatusFunc()
			if f != nil && f(resp, o) {
				return resp, status, o, nil
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
				respBody, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					return nil, status, o, fmt.Errorf("unexpected status %d, could not read response body", resp.StatusCode)
				}
				errText = fmt.Sprintf("response: %s", respBody)
			}
			return nil, status, o, fmt.Errorf("unexpected status %d with %s", resp.StatusCode, errText)
		}

		break
	}

	return resp, status, o, nil
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

// DeleteHttpRequestInput configures a DELETE request.
type DeleteHttpRequestInput struct {
	ValidStatusCodes []int
	ValidStatusFunc  ValidStatusFunc
	Uri              Uri
}

// GetValidStatusCodes returns a []int of status codes considered valid for a DELETE request.
func (i DeleteHttpRequestInput) GetValidStatusCodes() []int {
	return i.ValidStatusCodes
}

// GetValidStatusFunc returns a function used to evaluate whether the response to a DELETE request is considered valid.
func (i DeleteHttpRequestInput) GetValidStatusFunc() ValidStatusFunc {
	return i.ValidStatusFunc
}

// Delete performs a DELETE request.
func (c Client) Delete(ctx context.Context, input DeleteHttpRequestInput) (*http.Response, int, *odata.OData, error) {
	var status int
	url, err := c.buildUri(input.Uri)
	if err != nil {
		return nil, status, nil, fmt.Errorf("unable to make request: %v", err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, http.NoBody)
	if err != nil {
		return nil, status, nil, err
	}
	resp, status, o, err := c.performRequest(req, input)
	if err != nil {
		return nil, status, o, err
	}
	return resp, status, o, nil
}

// GetHttpRequestInput configures a GET request.
type GetHttpRequestInput struct {
	ValidStatusCodes []int
	ValidStatusFunc  ValidStatusFunc
	Uri              Uri
	rawUri           string
}

// GetValidStatusCodes returns a []int of status codes considered valid for a GET request.
func (i GetHttpRequestInput) GetValidStatusCodes() []int {
	return i.ValidStatusCodes
}

// GetValidStatusFunc returns a function used to evaluate whether the response to a GET request is considered valid.
func (i GetHttpRequestInput) GetValidStatusFunc() ValidStatusFunc {
	return i.ValidStatusFunc
}

// Get performs a GET request.
func (c Client) Get(ctx context.Context, input GetHttpRequestInput) (*http.Response, int, *odata.OData, error) {
	var status int

	// Check for a raw uri, else build one from the Uri field
	url := input.rawUri
	if url == "" {
		var err error
		url, err = c.buildUri(input.Uri)
		if err != nil {
			return nil, status, nil, fmt.Errorf("unable to make request: %v", err)
		}
	}

	// Build a new request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, status, nil, err
	}

	// Perform the request
	resp, status, o, err := c.performRequest(req, input)
	if err != nil {
		return nil, status, o, err
	}

	// Check for json content before handling pagination
	contentType := strings.ToLower(resp.Header.Get("Content-Type"))
	if strings.HasPrefix(contentType, "application/json") {
		// Read the response body and close it
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, status, o, fmt.Errorf("could not parse response body")
		}
		resp.Body.Close()

		// Unmarshall firstOdata
		var firstOdata odata.OData
		if err := json.Unmarshal(respBody, &firstOdata); err != nil {
			return nil, status, o, err
		}

		if firstOdata.NextLink == nil || firstOdata.Value == nil {
			// No more pages, reassign response body and return
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(respBody))
			return resp, status, o, nil
		}

		// Get the next page, recursively
		nextInput := input
		nextInput.rawUri = *firstOdata.NextLink
		nextResp, status, o, err := c.Get(ctx, nextInput)
		if err != nil {
			return resp, status, o, err
		}

		// Read the next page response body and close it
		nextRespBody, err := ioutil.ReadAll(nextResp.Body)
		if err != nil {
			return nil, status, o, fmt.Errorf("could not parse response body")
		}
		nextResp.Body.Close()

		// Unmarshall firstOdata from the next page
		var nextOdata odata.OData
		if err := json.Unmarshal(nextRespBody, &nextOdata); err != nil {
			return resp, status, o, err
		}

		if nextOdata.Value != nil {
			// Next page has results, append to current page
			value := append(*firstOdata.Value, *nextOdata.Value...)
			nextOdata.Value = &value
		}

		// Marshal the entire result, along with fields from the final page
		newJson, err := json.Marshal(nextOdata)
		if err != nil {
			return resp, status, o, err
		}

		// Reassign the response body
		resp.Body = ioutil.NopCloser(bytes.NewBuffer(newJson))
	}

	return resp, status, o, nil
}

// PatchHttpRequestInput configures a PATCH request.
type PatchHttpRequestInput struct {
	Body             []byte
	ValidStatusCodes []int
	ValidStatusFunc  ValidStatusFunc
	Uri              Uri
}

// GetValidStatusCodes returns a []int of status codes considered valid for a PATCH request.
func (i PatchHttpRequestInput) GetValidStatusCodes() []int {
	return i.ValidStatusCodes
}

// GetValidStatusFunc returns a function used to evaluate whether the response to a PATCH request is considered valid.
func (i PatchHttpRequestInput) GetValidStatusFunc() ValidStatusFunc {
	return i.ValidStatusFunc
}

// Patch performs a PATCH request.
func (c Client) Patch(ctx context.Context, input PatchHttpRequestInput) (*http.Response, int, *odata.OData, error) {
	var status int
	url, err := c.buildUri(input.Uri)
	if err != nil {
		return nil, status, nil, fmt.Errorf("unable to make request: %v", err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, url, bytes.NewBuffer(input.Body))
	if err != nil {
		return nil, status, nil, err
	}
	resp, status, o, err := c.performRequest(req, input)
	if err != nil {
		return nil, status, o, err
	}
	return resp, status, o, nil
}

// PostHttpRequestInput configures a POST request.
type PostHttpRequestInput struct {
	Body             []byte
	ValidStatusCodes []int
	ValidStatusFunc  ValidStatusFunc
	Uri              Uri
}

// GetValidStatusCodes returns a []int of status codes considered valid for a POST request.
func (i PostHttpRequestInput) GetValidStatusCodes() []int {
	return i.ValidStatusCodes
}

// GetValidStatusFunc returns a function used to evaluate whether the response to a POST request is considered valid.
func (i PostHttpRequestInput) GetValidStatusFunc() ValidStatusFunc {
	return i.ValidStatusFunc
}

// Post performs a POST request.
func (c Client) Post(ctx context.Context, input PostHttpRequestInput) (*http.Response, int, *odata.OData, error) {
	var status int
	url, err := c.buildUri(input.Uri)
	if err != nil {
		return nil, status, nil, fmt.Errorf("unable to make request: %v", err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(input.Body))
	if err != nil {
		return nil, status, nil, err
	}
	resp, status, o, err := c.performRequest(req, input)
	if err != nil {
		return nil, status, o, err
	}
	return resp, status, o, nil
}

// PutHttpRequestInput configures a PUT request.
type PutHttpRequestInput struct {
	Body             []byte
	ValidStatusCodes []int
	ValidStatusFunc  ValidStatusFunc
	Uri              Uri
}

// GetValidStatusCodes returns a []int of status codes considered valid for a PUT request.
func (i PutHttpRequestInput) GetValidStatusCodes() []int {
	return i.ValidStatusCodes
}

// GetValidStatusFunc returns a function used to evaluate whether the response to a PUT request is considered valid.
func (i PutHttpRequestInput) GetValidStatusFunc() ValidStatusFunc {
	return i.ValidStatusFunc
}

// Put performs a PUT request.
func (c Client) Put(ctx context.Context, input PutHttpRequestInput) (*http.Response, int, *odata.OData, error) {
	var status int
	url, err := c.buildUri(input.Uri)
	if err != nil {
		return nil, status, nil, fmt.Errorf("unable to make request: %v", err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, bytes.NewBuffer(input.Body))
	if err != nil {
		return nil, status, nil, err
	}
	resp, status, o, err := c.performRequest(req, input)
	if err != nil {
		return nil, status, o, err
	}
	return resp, status, o, nil
}
