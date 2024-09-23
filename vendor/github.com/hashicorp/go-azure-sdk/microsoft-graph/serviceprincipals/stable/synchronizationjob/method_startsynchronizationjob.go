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

type StartSynchronizationJobOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type StartSynchronizationJobOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultStartSynchronizationJobOperationOptions() StartSynchronizationJobOperationOptions {
	return StartSynchronizationJobOperationOptions{}
}

func (o StartSynchronizationJobOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o StartSynchronizationJobOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o StartSynchronizationJobOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// StartSynchronizationJob - Invoke action start. Start an existing synchronization job. If the job is in a paused
// state, it continues processing changes from the point where it was paused. If the job is in quarantine, the
// quarantine status is cleared. Don't create scripts to call the start job continuously while it's running because that
// can cause the service to stop running. Use the start job only when the job is currently paused or in quarantine.
func (c SynchronizationJobClient) StartSynchronizationJob(ctx context.Context, id stable.ServicePrincipalIdSynchronizationJobId, options StartSynchronizationJobOperationOptions) (result StartSynchronizationJobOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/start", id.ID()),
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
