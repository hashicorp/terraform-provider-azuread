package member

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

type ListMemberUsersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.User
}

type ListMemberUsersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.User
}

type ListMemberUsersOperationOptions struct {
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

func DefaultListMemberUsersOperationOptions() ListMemberUsersOperationOptions {
	return ListMemberUsersOperationOptions{}
}

func (o ListMemberUsersOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListMemberUsersOperationOptions) ToOData() *odata.Query {
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

func (o ListMemberUsersOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListMemberUsersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMemberUsersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMemberUsers - Get the items of type microsoft.graph.user in the microsoft.graph.directoryObject collection
func (c MemberClient) ListMemberUsers(ctx context.Context, id stable.DirectoryRoleId, options ListMemberUsersOperationOptions) (result ListMemberUsersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListMemberUsersCustomPager{},
		Path:          fmt.Sprintf("%s/members/user", id.ID()),
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
		Values *[]stable.User `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMemberUsersComplete retrieves all the results into a single object
func (c MemberClient) ListMemberUsersComplete(ctx context.Context, id stable.DirectoryRoleId, options ListMemberUsersOperationOptions) (ListMemberUsersCompleteResult, error) {
	return c.ListMemberUsersCompleteMatchingPredicate(ctx, id, options, UserOperationPredicate{})
}

// ListMemberUsersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MemberClient) ListMemberUsersCompleteMatchingPredicate(ctx context.Context, id stable.DirectoryRoleId, options ListMemberUsersOperationOptions, predicate UserOperationPredicate) (result ListMemberUsersCompleteResult, err error) {
	items := make([]stable.User, 0)

	resp, err := c.ListMemberUsers(ctx, id, options)
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

	result = ListMemberUsersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
