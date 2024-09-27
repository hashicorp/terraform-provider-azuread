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

type RetryServiceProvisioningOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type RetryServiceProvisioningOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultRetryServiceProvisioningOperationOptions() RetryServiceProvisioningOperationOptions {
	return RetryServiceProvisioningOperationOptions{}
}

func (o RetryServiceProvisioningOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RetryServiceProvisioningOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o RetryServiceProvisioningOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// RetryServiceProvisioning - Invoke action retryServiceProvisioning. Retry the user service provisioning.
func (c UserClient) RetryServiceProvisioning(ctx context.Context, id beta.UserId, options RetryServiceProvisioningOperationOptions) (result RetryServiceProvisioningOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/retryServiceProvisioning", id.ID()),
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
