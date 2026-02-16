package administrativeunitmember

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoveAdministrativeUnitMemberRefOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type RemoveAdministrativeUnitMemberRefOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultRemoveAdministrativeUnitMemberRefOperationOptions() RemoveAdministrativeUnitMemberRefOperationOptions {
	return RemoveAdministrativeUnitMemberRefOperationOptions{}
}

func (o RemoveAdministrativeUnitMemberRefOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o RemoveAdministrativeUnitMemberRefOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o RemoveAdministrativeUnitMemberRefOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// RemoveAdministrativeUnitMemberRef - Remove a member. Use this API to remove a member (user, group, or device) from an
// administrative unit.
func (c AdministrativeUnitMemberClient) RemoveAdministrativeUnitMemberRef(ctx context.Context, id stable.DirectoryAdministrativeUnitIdMemberId, options RemoveAdministrativeUnitMemberRefOperationOptions) (result RemoveAdministrativeUnitMemberRefOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/$ref", id.ID()),
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
