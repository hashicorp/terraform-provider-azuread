package oauth2permissiongrant

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetOAuth2PermissionGrantOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.OAuth2PermissionGrant
}

type GetOAuth2PermissionGrantOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetOAuth2PermissionGrantOperationOptions() GetOAuth2PermissionGrantOperationOptions {
	return GetOAuth2PermissionGrantOperationOptions{}
}

func (o GetOAuth2PermissionGrantOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetOAuth2PermissionGrantOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetOAuth2PermissionGrantOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetOAuth2PermissionGrant - Get oAuth2PermissionGrant (a delegated permission grant). Retrieve the properties of a
// single delegated permission grant represented by an oAuth2PermissionGrant object. An oAuth2PermissionGrant represents
// delegated permissions which have been granted for a client application to access an API on behalf of a signed-in
// user.
func (c OAuth2PermissionGrantClient) GetOAuth2PermissionGrant(ctx context.Context, id stable.OAuth2PermissionGrantId, options GetOAuth2PermissionGrantOperationOptions) (result GetOAuth2PermissionGrantOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var model stable.OAuth2PermissionGrant
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
