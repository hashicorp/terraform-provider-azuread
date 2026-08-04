package directoryroleeligibilityscheduleinstance

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateDirectoryRoleEligibilityScheduleInstanceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateDirectoryRoleEligibilityScheduleInstanceOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateDirectoryRoleEligibilityScheduleInstanceOperationOptions() UpdateDirectoryRoleEligibilityScheduleInstanceOperationOptions {
	return UpdateDirectoryRoleEligibilityScheduleInstanceOperationOptions{}
}

func (o UpdateDirectoryRoleEligibilityScheduleInstanceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateDirectoryRoleEligibilityScheduleInstanceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateDirectoryRoleEligibilityScheduleInstanceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateDirectoryRoleEligibilityScheduleInstance - Update the navigation property roleEligibilityScheduleInstances in
// roleManagement
func (c DirectoryRoleEligibilityScheduleInstanceClient) UpdateDirectoryRoleEligibilityScheduleInstance(ctx context.Context, id stable.RoleManagementDirectoryRoleEligibilityScheduleInstanceId, input stable.UnifiedRoleEligibilityScheduleInstance, options UpdateDirectoryRoleEligibilityScheduleInstanceOperationOptions) (result UpdateDirectoryRoleEligibilityScheduleInstanceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
