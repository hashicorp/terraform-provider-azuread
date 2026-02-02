package privilegedaccessgroupassignmentscheduleinstance

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreatePrivilegedAccessGroupAssignmentScheduleInstanceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.PrivilegedAccessGroupAssignmentScheduleInstance
}

type CreatePrivilegedAccessGroupAssignmentScheduleInstanceOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreatePrivilegedAccessGroupAssignmentScheduleInstanceOperationOptions() CreatePrivilegedAccessGroupAssignmentScheduleInstanceOperationOptions {
	return CreatePrivilegedAccessGroupAssignmentScheduleInstanceOperationOptions{}
}

func (o CreatePrivilegedAccessGroupAssignmentScheduleInstanceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreatePrivilegedAccessGroupAssignmentScheduleInstanceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreatePrivilegedAccessGroupAssignmentScheduleInstanceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreatePrivilegedAccessGroupAssignmentScheduleInstance - Create new navigation property to assignmentScheduleInstances
// for identityGovernance
func (c PrivilegedAccessGroupAssignmentScheduleInstanceClient) CreatePrivilegedAccessGroupAssignmentScheduleInstance(ctx context.Context, input stable.PrivilegedAccessGroupAssignmentScheduleInstance, options CreatePrivilegedAccessGroupAssignmentScheduleInstanceOperationOptions) (result CreatePrivilegedAccessGroupAssignmentScheduleInstanceOperationResponse, err error) {
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
		Path:          "/identityGovernance/privilegedAccess/group/assignmentScheduleInstances",
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

	var model stable.PrivilegedAccessGroupAssignmentScheduleInstance
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
