package entitlementmanagementassignmentpolicy

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SetEntitlementManagementAssignmentPolicyOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SetEntitlementManagementAssignmentPolicyOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultSetEntitlementManagementAssignmentPolicyOperationOptions() SetEntitlementManagementAssignmentPolicyOperationOptions {
	return SetEntitlementManagementAssignmentPolicyOperationOptions{}
}

func (o SetEntitlementManagementAssignmentPolicyOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SetEntitlementManagementAssignmentPolicyOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SetEntitlementManagementAssignmentPolicyOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SetEntitlementManagementAssignmentPolicy - Update accessPackageAssignmentPolicy. Update an existing
// accessPackageAssignmentPolicy object to change one or more of its properties, such as the display name or
// description.
func (c EntitlementManagementAssignmentPolicyClient) SetEntitlementManagementAssignmentPolicy(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementAssignmentPolicyId, input stable.AccessPackageAssignmentPolicy, options SetEntitlementManagementAssignmentPolicyOperationOptions) (result SetEntitlementManagementAssignmentPolicyOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPut,
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
