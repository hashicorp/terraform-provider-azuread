package user

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

type WipeManagedAppRegistrationByDeviceTagOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type WipeManagedAppRegistrationByDeviceTagOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultWipeManagedAppRegistrationByDeviceTagOperationOptions() WipeManagedAppRegistrationByDeviceTagOperationOptions {
	return WipeManagedAppRegistrationByDeviceTagOperationOptions{}
}

func (o WipeManagedAppRegistrationByDeviceTagOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o WipeManagedAppRegistrationByDeviceTagOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o WipeManagedAppRegistrationByDeviceTagOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// WipeManagedAppRegistrationByDeviceTag - Invoke action wipeManagedAppRegistrationByDeviceTag. Issues a wipe operation
// on an app registration with specified device tag.
func (c UserClient) WipeManagedAppRegistrationByDeviceTag(ctx context.Context, id beta.UserId, input WipeManagedAppRegistrationByDeviceTagRequest, options WipeManagedAppRegistrationByDeviceTagOperationOptions) (result WipeManagedAppRegistrationByDeviceTagOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/wipeManagedAppRegistrationByDeviceTag", id.ID()),
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
