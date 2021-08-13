package common

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/meta"
	"github.com/manicminer/hamilton/auth"
	"github.com/manicminer/hamilton/environments"
	"github.com/manicminer/hamilton/msgraph"

	"github.com/hashicorp/terraform-provider-azuread/version"
)

type ClientOptions struct {
	Environment environments.Environment
	TenantID    string

	PartnerID        string
	TerraformVersion string

	Authorizer auth.Authorizer
}

func (o ClientOptions) ConfigureClient(c *msgraph.Client) {
	c.Authorizer = o.Authorizer
	c.Endpoint = o.Environment.MsGraph.Endpoint
	c.UserAgent = o.userAgent(c.UserAgent)

	if c.RequestMiddlewares == nil {
		c.RequestMiddlewares = &[]msgraph.RequestMiddleware{}
	}
	if c.ResponseMiddlewares == nil {
		c.ResponseMiddlewares = &[]msgraph.ResponseMiddleware{}
	}
	*c.RequestMiddlewares = append(*c.RequestMiddlewares, o.requestLogger)
	*c.ResponseMiddlewares = append(*c.ResponseMiddlewares, o.responseLogger)

	c.RetryableClient.RetryMax = 20
}

func (o ClientOptions) requestLogger(req *http.Request) (*http.Request, error) {
	if req == nil {
		return nil, nil
	}

	// Don't log the Authorization header
	authHeaderName := "Authorization"
	authHeaderValue := req.Header.Get(authHeaderName)
	if authHeaderValue != "" {
		req.Header.Del(authHeaderName)
	}

	if dump, err := httputil.DumpRequestOut(req, true); err == nil {
		log.Printf("[DEBUG] Begin AzureAD Request: ==========================================\n%s\n========================================= End AzureAD Request\n", dump)
	} else {
		// fallback to basic message
		log.Printf("[DEBUG] AzureAD Request: %s %s\n", req.Method, req.URL)
	}

	if authHeaderValue != "" {
		req.Header.Add(authHeaderName, authHeaderValue)
	}
	return req, nil
}

func (o ClientOptions) responseLogger(req *http.Request, resp *http.Response) (*http.Response, error) {
	if resp != nil {
		if dump, err2 := httputil.DumpResponse(resp, true); err2 == nil {
			log.Printf("[DEBUG] Begin AzureAD Response for %s %s: ==========================================\n%s\n========================================== End AzureAD Response\n", req.Method, req.URL, dump)
		} else {
			log.Printf("[DEBUG] AzureAD Response: %s for %s %s\n", resp.Status, req.Method, req.URL)
		}
	} else {
		log.Printf("[DEBUG] AzureAD Request for %s %s completed with no response", req.Method, req.URL)
	}
	return resp, nil
}

func (o ClientOptions) userAgent(sdkUserAgent string) (userAgent string) {
	tfUserAgent := fmt.Sprintf("HashiCorp Terraform/%s (+https://www.terraform.io) Terraform Plugin SDK/%s", o.TerraformVersion, meta.SDKVersionString())
	providerUserAgent := fmt.Sprintf("%s terraform-provider-azuread/%s", tfUserAgent, version.ProviderVersion)
	userAgent = strings.TrimSpace(fmt.Sprintf("%s %s", providerUserAgent, sdkUserAgent))

	// append the CloudShell version to the user agent if it exists
	if azureAgent := os.Getenv("AZURE_HTTP_USER_AGENT"); azureAgent != "" {
		userAgent = fmt.Sprintf("%s %s", userAgent, azureAgent)
	}

	if o.PartnerID != "" {
		userAgent = fmt.Sprintf("%s pid-%s", userAgent, o.PartnerID)
	}

	return
}
