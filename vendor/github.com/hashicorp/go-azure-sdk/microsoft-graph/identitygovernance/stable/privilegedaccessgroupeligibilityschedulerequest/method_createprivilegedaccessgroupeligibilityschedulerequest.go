package privilegedaccessgroupeligibilityschedulerequest

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreatePrivilegedAccessGroupEligibilityScheduleRequestOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.PrivilegedAccessGroupEligibilityScheduleRequest
}

type CreatePrivilegedAccessGroupEligibilityScheduleRequestOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreatePrivilegedAccessGroupEligibilityScheduleRequestOperationOptions() CreatePrivilegedAccessGroupEligibilityScheduleRequestOperationOptions {
	return CreatePrivilegedAccessGroupEligibilityScheduleRequestOperationOptions{}
}

func (o CreatePrivilegedAccessGroupEligibilityScheduleRequestOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreatePrivilegedAccessGroupEligibilityScheduleRequestOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreatePrivilegedAccessGroupEligibilityScheduleRequestOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreatePrivilegedAccessGroupEligibilityScheduleRequest - Create eligibilityScheduleRequest. Create a new
// privilegedAccessGroupEligibilityScheduleRequest object.
func (c PrivilegedAccessGroupEligibilityScheduleRequestClient) CreatePrivilegedAccessGroupEligibilityScheduleRequest(ctx context.Context, input stable.PrivilegedAccessGroupEligibilityScheduleRequest, options CreatePrivilegedAccessGroupEligibilityScheduleRequestOperationOptions) (result CreatePrivilegedAccessGroupEligibilityScheduleRequestOperationResponse, err error) {
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
		Path:          "/identityGovernance/privilegedAccess/group/eligibilityScheduleRequests",
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

	var model stable.PrivilegedAccessGroupEligibilityScheduleRequest
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
