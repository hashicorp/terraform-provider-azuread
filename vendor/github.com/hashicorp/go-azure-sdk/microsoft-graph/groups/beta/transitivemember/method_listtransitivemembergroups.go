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

type ListTransitiveMemberGroupsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Group
}

type ListTransitiveMemberGroupsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Group
}

type ListTransitiveMemberGroupsOperationOptions struct {
	ConsistencyLevel *odata.ConsistencyLevel
	Count            *bool
	Expand           *odata.Expand
	Filter           *string
	Metadata         *odata.Metadata
	OrderBy          *odata.OrderBy
	RetryFunc        client.RequestRetryFunc
	Search           *string
	Select           *[]string
	Skip             *int64
	Top              *int64
}

func DefaultListTransitiveMemberGroupsOperationOptions() ListTransitiveMemberGroupsOperationOptions {
	return ListTransitiveMemberGroupsOperationOptions{}
}

func (o ListTransitiveMemberGroupsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListTransitiveMemberGroupsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.ConsistencyLevel != nil {
		out.ConsistencyLevel = *o.ConsistencyLevel
	}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListTransitiveMemberGroupsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListTransitiveMemberGroupsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTransitiveMemberGroupsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTransitiveMemberGroups - List group transitive members. Get a list of a group's members. A group can have
// different object types as members. For more information about supported member types for different groups, see Group
// membership. This operation is transitive and returns a flat list of all nested members. An attempt to filter by an
// OData cast that represents an unsupported member type returns a 400 Bad Request error with the
// Request_UnsupportedQuery code.
func (c TransitiveMemberClient) ListTransitiveMemberGroups(ctx context.Context, id beta.GroupId, options ListTransitiveMemberGroupsOperationOptions) (result ListTransitiveMemberGroupsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListTransitiveMemberGroupsCustomPager{},
		Path:          fmt.Sprintf("%s/transitiveMembers/group", id.ID()),
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]beta.Group `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTransitiveMemberGroupsComplete retrieves all the results into a single object
func (c TransitiveMemberClient) ListTransitiveMemberGroupsComplete(ctx context.Context, id beta.GroupId, options ListTransitiveMemberGroupsOperationOptions) (ListTransitiveMemberGroupsCompleteResult, error) {
	return c.ListTransitiveMemberGroupsCompleteMatchingPredicate(ctx, id, options, GroupOperationPredicate{})
}

// ListTransitiveMemberGroupsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TransitiveMemberClient) ListTransitiveMemberGroupsCompleteMatchingPredicate(ctx context.Context, id beta.GroupId, options ListTransitiveMemberGroupsOperationOptions, predicate GroupOperationPredicate) (result ListTransitiveMemberGroupsCompleteResult, err error) {
	items := make([]beta.Group, 0)

	resp, err := c.ListTransitiveMemberGroups(ctx, id, options)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ListTransitiveMemberGroupsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
