// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package common

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/auth"
	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/meta"
	"github.com/hashicorp/terraform-provider-azuread/version"
)

type contextKey string

type ClientOptions struct {
	Environment environments.Environment
	TenantID    string

	PartnerID        string
	TerraformVersion string

	Authorizer auth.Authorizer
}

func (o ClientOptions) Configure(c *msgraph.Client) {
	c.SetAuthorizer(o.Authorizer)
	c.SetUserAgent(o.userAgent(c.UserAgent))
	c.AppendRequestMiddleware(o.requestLogger)
	c.AppendResponseMiddleware(o.responseLogger)
}

func (o ClientOptions) requestLogger(req *http.Request) (*http.Request, error) {
	if req == nil {
		return nil, nil
	}

	requestId, err := uuid.GenerateUUID()
	if err != nil {
		return nil, err
	}

	ctx := req.Context()
	newReq := req.WithContext(context.WithValue(ctx, contextKey("requestId"), requestId))

	// Don't log the Authorization header
	authHeaderName := "Authorization"
	authHeaderValue := newReq.Header.Get(authHeaderName)
	if authHeaderValue != "" {
		newReq.Header.Del(authHeaderName)
	}

	if dump, err := httputil.DumpRequestOut(newReq, true); err == nil {
		log.Printf(`[DEBUG] ============================ Begin AzureAD Request ============================
Request ID: %s

%s
============================= End AzureAD Request =============================
`, requestId, dump)
	} else {
		// fallback to basic message
		log.Printf("[DEBUG] AzureAD Request %s: %s %s\n", requestId, newReq.Method, newReq.URL)
	}

	if authHeaderValue != "" {
		newReq.Header.Add(authHeaderName, authHeaderValue)
	}
	return newReq, nil
}

func (o ClientOptions) responseLogger(req *http.Request, resp *http.Response) (*http.Response, error) {
	requestId := "UNKNOWN"

	if req != nil {
		if v := req.Context().Value(contextKey("requestId")); v != nil {
			requestId = v.(string)
		}
	}

	if resp != nil {
		if dump, err2 := httputil.DumpResponse(resp, true); err2 == nil {
			log.Printf(`[DEBUG] ============================ Begin AzureAD Response ===========================
%s %s
Request ID: %s

%s
============================= End AzureAD Response ============================
`, req.Method, req.URL, requestId, dump)
		} else {
			log.Printf("[DEBUG] AzureAD Response: %s for %s (%s %s)\n", resp.Status, requestId, req.Method, req.URL)
		}
	} else {
		log.Printf("[DEBUG] AzureAD Request for %s (%s %s) completed with no response", requestId, req.Method, req.URL)
	}

	return resp, nil
}

func (o ClientOptions) userAgent(sdkUserAgent string) (userAgent string) {
	tfUserAgent := fmt.Sprintf("HashiCorp Terraform/%s (+https://www.terraform.io) Terraform Plugin SDK/%s", o.TerraformVersion, meta.SDKVersionString()) //nolint:staticcheck
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
