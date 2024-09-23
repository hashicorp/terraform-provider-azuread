package me

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoveAllDevicesFromManagementOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type RemoveAllDevicesFromManagementOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultRemoveAllDevicesFromManagementOperationOptions() RemoveAllDevicesFromManagementOperationOptions {
	return RemoveAllDevicesFromManagementOperationOptions{}
}

func (o RemoveAllDevicesFromManagementOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RemoveAllDevicesFromManagementOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o RemoveAllDevicesFromManagementOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// RemoveAllDevicesFromManagement - Invoke action removeAllDevicesFromManagement. Retire all devices from management for
// this user
func (c MeClient) RemoveAllDevicesFromManagement(ctx context.Context, options RemoveAllDevicesFromManagementOperationOptions) (result RemoveAllDevicesFromManagementOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/me/removeAllDevicesFromManagement",
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
