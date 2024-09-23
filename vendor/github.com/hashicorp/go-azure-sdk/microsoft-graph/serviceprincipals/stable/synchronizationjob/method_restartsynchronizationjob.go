package synchronizationjob

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestartSynchronizationJobOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type RestartSynchronizationJobOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultRestartSynchronizationJobOperationOptions() RestartSynchronizationJobOperationOptions {
	return RestartSynchronizationJobOperationOptions{}
}

func (o RestartSynchronizationJobOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RestartSynchronizationJobOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o RestartSynchronizationJobOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// RestartSynchronizationJob - Invoke action restart. Restart a stopped synchronization job, forcing it to reprocess all
// the objects in the directory. Optionally clears existing the synchronization state and previous errors.
func (c SynchronizationJobClient) RestartSynchronizationJob(ctx context.Context, id stable.ServicePrincipalIdSynchronizationJobId, input RestartSynchronizationJobRequest, options RestartSynchronizationJobOperationOptions) (result RestartSynchronizationJobOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/restart", id.ID()),
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
