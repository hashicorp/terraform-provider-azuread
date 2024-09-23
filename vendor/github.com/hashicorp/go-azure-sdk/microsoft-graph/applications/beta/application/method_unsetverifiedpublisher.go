package application

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnsetVerifiedPublisherOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UnsetVerifiedPublisherOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUnsetVerifiedPublisherOperationOptions() UnsetVerifiedPublisherOperationOptions {
	return UnsetVerifiedPublisherOperationOptions{}
}

func (o UnsetVerifiedPublisherOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UnsetVerifiedPublisherOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UnsetVerifiedPublisherOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UnsetVerifiedPublisher - Invoke action unsetVerifiedPublisher. Unset the verifiedPublisher previously set on an
// application, removing all verified publisher properties. For more information, see Publisher verification.
func (c ApplicationClient) UnsetVerifiedPublisher(ctx context.Context, id beta.ApplicationId, options UnsetVerifiedPublisherOperationOptions) (result UnsetVerifiedPublisherOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/unsetVerifiedPublisher", id.ID()),
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

	return
}
