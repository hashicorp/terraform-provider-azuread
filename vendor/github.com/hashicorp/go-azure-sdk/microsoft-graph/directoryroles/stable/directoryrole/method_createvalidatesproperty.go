package directoryrole

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateValidatesPropertyOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateValidatesPropertyOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateValidatesPropertyOperationOptions() CreateValidatesPropertyOperationOptions {
	return CreateValidatesPropertyOperationOptions{}
}

func (o CreateValidatesPropertyOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateValidatesPropertyOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateValidatesPropertyOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateValidatesProperty - Invoke action validateProperties. Validate that a Microsoft 365 group's display name or
// mail nickname complies with naming policies. Clients can use this API to determine whether a display name or mail
// nickname is valid before trying to create a Microsoft 365 group. To validate the properties of an existing group, use
// the group: validateProperties function. The following policy validations are performed for the display name and mail
// nickname properties: 1. Validate the prefix and suffix naming policy 2. Validate the custom banned words policy 3.
// Validate that the mail nickname is unique This API only returns the first validation failure that is encountered. If
// the properties fail multiple validations, only the first validation failure is returned. However, you can validate
// both the mail nickname and the display name and receive a collection of validation errors if you're only validating
// the prefix and suffix naming policy. To learn more about configuring naming policies, see Configure naming policy.
func (c DirectoryRoleClient) CreateValidatesProperty(ctx context.Context, input CreateValidatesPropertyRequest, options CreateValidatesPropertyOperationOptions) (result CreateValidatesPropertyOperationResponse, err error) {
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
		Path:          "/directoryRoles/validateProperties",
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
