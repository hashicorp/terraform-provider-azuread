package transitivemember

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetTransitiveMemberGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.Group
}

type GetTransitiveMemberGroupOperationOptions struct {
	ConsistencyLevel *odata.ConsistencyLevel
	Expand           *odata.Expand
	Metadata         *odata.Metadata
	RetryFunc        client.RequestRetryFunc
	Select           *[]string
}

func DefaultGetTransitiveMemberGroupOperationOptions() GetTransitiveMemberGroupOperationOptions {
	return GetTransitiveMemberGroupOperationOptions{}
}

func (o GetTransitiveMemberGroupOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetTransitiveMemberGroupOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.ConsistencyLevel != nil {
		out.ConsistencyLevel = *o.ConsistencyLevel
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetTransitiveMemberGroupOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetTransitiveMemberGroup - List group transitive members. Get a list of a group's members. A group can have different
// object types as members. For more information about supported member types for different groups, see Group
// membership. This operation is transitive and returns a flat list of all nested members. An attempt to filter by an
// OData cast that represents an unsupported member type returns a 400 Bad Request error with the
// Request_UnsupportedQuery code.
func (c TransitiveMemberClient) GetTransitiveMemberGroup(ctx context.Context, id beta.GroupIdTransitiveMemberId, options GetTransitiveMemberGroupOperationOptions) (result GetTransitiveMemberGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/group", id.ID()),
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

	var model beta.Group
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
