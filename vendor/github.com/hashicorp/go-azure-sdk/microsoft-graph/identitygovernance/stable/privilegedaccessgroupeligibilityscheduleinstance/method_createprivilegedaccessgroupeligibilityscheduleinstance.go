package privilegedaccessgroupeligibilityscheduleinstance

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreatePrivilegedAccessGroupEligibilityScheduleInstanceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.PrivilegedAccessGroupEligibilityScheduleInstance
}

type CreatePrivilegedAccessGroupEligibilityScheduleInstanceOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreatePrivilegedAccessGroupEligibilityScheduleInstanceOperationOptions() CreatePrivilegedAccessGroupEligibilityScheduleInstanceOperationOptions {
	return CreatePrivilegedAccessGroupEligibilityScheduleInstanceOperationOptions{}
}

func (o CreatePrivilegedAccessGroupEligibilityScheduleInstanceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreatePrivilegedAccessGroupEligibilityScheduleInstanceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreatePrivilegedAccessGroupEligibilityScheduleInstanceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreatePrivilegedAccessGroupEligibilityScheduleInstance - Create new navigation property to
// eligibilityScheduleInstances for identityGovernance
func (c PrivilegedAccessGroupEligibilityScheduleInstanceClient) CreatePrivilegedAccessGroupEligibilityScheduleInstance(ctx context.Context, input stable.PrivilegedAccessGroupEligibilityScheduleInstance, options CreatePrivilegedAccessGroupEligibilityScheduleInstanceOperationOptions) (result CreatePrivilegedAccessGroupEligibilityScheduleInstanceOperationResponse, err error) {
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
		Path:          "/identityGovernance/privilegedAccess/group/eligibilityScheduleInstances",
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

	var model stable.PrivilegedAccessGroupEligibilityScheduleInstance
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
