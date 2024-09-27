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

type SetVerifiedPublisherOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SetVerifiedPublisherOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultSetVerifiedPublisherOperationOptions() SetVerifiedPublisherOperationOptions {
	return SetVerifiedPublisherOperationOptions{}
}

func (o SetVerifiedPublisherOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SetVerifiedPublisherOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SetVerifiedPublisherOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SetVerifiedPublisher - Invoke action setVerifiedPublisher. Set the the verifiedPublisher on an application. For more
// information, including prerequisites to setting a verified publisher, see Publisher verification.
func (c ApplicationClient) SetVerifiedPublisher(ctx context.Context, id beta.ApplicationId, input SetVerifiedPublisherRequest, options SetVerifiedPublisherOperationOptions) (result SetVerifiedPublisherOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/setVerifiedPublisher", id.ID()),
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
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
