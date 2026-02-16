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
func (c UserClient) CreateExportPersonalData(ctx context.Context, id beta.UserId, input CreateExportPersonalDataRequest, options CreateExportPersonalDataOperationOptions) (result CreateExportPersonalDataOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/exportPersonalData", id.ID()),
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
