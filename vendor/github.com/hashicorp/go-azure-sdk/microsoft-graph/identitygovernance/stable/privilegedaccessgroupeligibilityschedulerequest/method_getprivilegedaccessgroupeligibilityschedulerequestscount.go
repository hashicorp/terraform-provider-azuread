package privilegedaccessgroupeligibilityschedulerequest

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetPrivilegedAccessGroupEligibilityScheduleRequestsCountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetPrivilegedAccessGroupEligibilityScheduleRequestsCountOperationOptions struct {
	Filter    *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Search    *string
}

func DefaultGetPrivilegedAccessGroupEligibilityScheduleRequestsCountOperationOptions() GetPrivilegedAccessGroupEligibilityScheduleRequestsCountOperationOptions {
	return GetPrivilegedAccessGroupEligibilityScheduleRequestsCountOperationOptions{}
}

func (o GetPrivilegedAccessGroupEligibilityScheduleRequestsCountOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetPrivilegedAccessGroupEligibilityScheduleRequestsCountOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	return &out
}

func (o GetPrivilegedAccessGroupEligibilityScheduleRequestsCountOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetPrivilegedAccessGroupEligibilityScheduleRequestsCount - Get the number of the resource
func (c PrivilegedAccessGroupEligibilityScheduleRequestClient) GetPrivilegedAccessGroupEligibilityScheduleRequestsCount(ctx context.Context, options GetPrivilegedAccessGroupEligibilityScheduleRequestsCountOperationOptions) (result GetPrivilegedAccessGroupEligibilityScheduleRequestsCountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "text/plain",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          "/identityGovernance/privilegedAccess/group/eligibilityScheduleRequests/$count",
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

	var model []byte
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
