package approleassignedto

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

type DeleteAppRoleAssignedToOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteAppRoleAssignedToOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDeleteAppRoleAssignedToOperationOptions() DeleteAppRoleAssignedToOperationOptions {
	return DeleteAppRoleAssignedToOperationOptions{}
}

func (o DeleteAppRoleAssignedToOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteAppRoleAssignedToOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteAppRoleAssignedToOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteAppRoleAssignedTo - Delete appRoleAssignedTo. Deletes an appRoleAssignment that a user, group, or client
// service principal has been granted for a resource service principal.
func (c AppRoleAssignedToClient) DeleteAppRoleAssignedTo(ctx context.Context, id stable.ServicePrincipalIdAppRoleAssignedToId, options DeleteAppRoleAssignedToOperationOptions) (result DeleteAppRoleAssignedToOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          id.ID(),
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
