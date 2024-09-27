package me

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WipeManagedAppRegistrationsByDeviceTagOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type WipeManagedAppRegistrationsByDeviceTagOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultWipeManagedAppRegistrationsByDeviceTagOperationOptions() WipeManagedAppRegistrationsByDeviceTagOperationOptions {
	return WipeManagedAppRegistrationsByDeviceTagOperationOptions{}
}

func (o WipeManagedAppRegistrationsByDeviceTagOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o WipeManagedAppRegistrationsByDeviceTagOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o WipeManagedAppRegistrationsByDeviceTagOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// WipeManagedAppRegistrationsByDeviceTag - Invoke action wipeManagedAppRegistrationsByDeviceTag. Issues a wipe
// operation on an app registration with specified device tag.
func (c MeClient) WipeManagedAppRegistrationsByDeviceTag(ctx context.Context, input WipeManagedAppRegistrationsByDeviceTagRequest, options WipeManagedAppRegistrationsByDeviceTagOperationOptions) (result WipeManagedAppRegistrationsByDeviceTagOperationResponse, err error) {
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
		Path:          "/me/wipeManagedAppRegistrationsByDeviceTag",
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
