package user

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WipeManagedAppRegistrationsByAzureAdDeviceIdOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type WipeManagedAppRegistrationsByAzureAdDeviceIdOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultWipeManagedAppRegistrationsByAzureAdDeviceIdOperationOptions() WipeManagedAppRegistrationsByAzureAdDeviceIdOperationOptions {
	return WipeManagedAppRegistrationsByAzureAdDeviceIdOperationOptions{}
}

func (o WipeManagedAppRegistrationsByAzureAdDeviceIdOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o WipeManagedAppRegistrationsByAzureAdDeviceIdOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o WipeManagedAppRegistrationsByAzureAdDeviceIdOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// WipeManagedAppRegistrationsByAzureAdDeviceId - Invoke action wipeManagedAppRegistrationsByAzureAdDeviceId. Issues a
// wipe operation on an app registration with specified aad device Id.
func (c UserClient) WipeManagedAppRegistrationsByAzureAdDeviceId(ctx context.Context, id beta.UserId, input WipeManagedAppRegistrationsByAzureAdDeviceIdRequest, options WipeManagedAppRegistrationsByAzureAdDeviceIdOperationOptions) (result WipeManagedAppRegistrationsByAzureAdDeviceIdOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/wipeManagedAppRegistrationsByAzureAdDeviceId", id.ID()),
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
