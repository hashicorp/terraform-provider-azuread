package privilegedaccessgroupassignmentschedulerequest

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

type CancelPrivilegedAccessGroupAssignmentScheduleRequestOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CancelPrivilegedAccessGroupAssignmentScheduleRequestOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCancelPrivilegedAccessGroupAssignmentScheduleRequestOperationOptions() CancelPrivilegedAccessGroupAssignmentScheduleRequestOperationOptions {
	return CancelPrivilegedAccessGroupAssignmentScheduleRequestOperationOptions{}
}

func (o CancelPrivilegedAccessGroupAssignmentScheduleRequestOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CancelPrivilegedAccessGroupAssignmentScheduleRequestOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CancelPrivilegedAccessGroupAssignmentScheduleRequestOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CancelPrivilegedAccessGroupAssignmentScheduleRequest - Invoke action cancel. Cancel an access assignment request to a
// group whose membership and ownership are governed by PIM.
func (c PrivilegedAccessGroupAssignmentScheduleRequestClient) CancelPrivilegedAccessGroupAssignmentScheduleRequest(ctx context.Context, id stable.IdentityGovernancePrivilegedAccessGroupAssignmentScheduleRequestId, options CancelPrivilegedAccessGroupAssignmentScheduleRequestOperationOptions) (result CancelPrivilegedAccessGroupAssignmentScheduleRequestOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/cancel", id.ID()),
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
