package me

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateExportPersonalDataOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateExportPersonalDataOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateExportPersonalDataOperationOptions() CreateExportPersonalDataOperationOptions {
	return CreateExportPersonalDataOperationOptions{}
}

func (o CreateExportPersonalDataOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateExportPersonalDataOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateExportPersonalDataOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateExportPersonalData - Invoke action exportPersonalData. Submit a data policy operation request from a company
// administrator or an application to export an organizational user's data. This data includes the user's data stored in
// OneDrive and their activity reports. For more information about exporting data while complying with regulations, see
// Data Subject Requests and the GDPR and CCPA.
func (c MeClient) CreateExportPersonalData(ctx context.Context, input CreateExportPersonalDataRequest, options CreateExportPersonalDataOperationOptions) (result CreateExportPersonalDataOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/me/exportPersonalData",
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
