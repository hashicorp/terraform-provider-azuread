package base

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/manicminer/hamilton/auth"
	"github.com/manicminer/hamilton/base/odata"
	"github.com/manicminer/hamilton/environments"
)

type ApiVersion string

const (
	Version10   ApiVersion = "v1.0"
	VersionBeta ApiVersion = "beta"
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

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, status, nil, err
	}

	var o odata.OData

	// Check for json content before looking for odata metadata
	contentType := strings.ToLower(resp.Header.Get("Content-Type"))
	if strings.HasPrefix(contentType, "application/json") {
		// Read the response body and close it
		respBody, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return nil, status, nil, fmt.Errorf("could not read response body: %s", err)
		}

		// Unmarshall odata
		if err := json.Unmarshal(respBody, &o); err != nil {
			return nil, status, nil, err
		}

		// Reassign the response body
		resp.Body = ioutil.NopCloser(bytes.NewBuffer(respBody))
	}

	status = resp.StatusCode
	if !containsStatusCode(input.GetValidStatusCodes(), status) {
		f := input.GetValidStatusFunc()
		if f != nil && f(resp, &o) {
			return resp, status, &o, nil
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
