package administrativeunitmember

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

type ListAdministrativeUnitMemberUsersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.User
}

type ListAdministrativeUnitMemberUsersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.User
}

type ListAdministrativeUnitMemberUsersOperationOptions struct {
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

func DefaultListAdministrativeUnitMemberUsersOperationOptions() ListAdministrativeUnitMemberUsersOperationOptions {
	return ListAdministrativeUnitMemberUsersOperationOptions{}
}

func (o ListAdministrativeUnitMemberUsersOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListAdministrativeUnitMemberUsersOperationOptions) ToOData() *odata.Query {
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

func (o ListAdministrativeUnitMemberUsersOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListAdministrativeUnitMemberUsersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAdministrativeUnitMemberUsersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAdministrativeUnitMemberUsers - Get the items of type microsoft.graph.user in the microsoft.graph.directoryObject
// collection
func (c AdministrativeUnitMemberClient) ListAdministrativeUnitMemberUsers(ctx context.Context, id stable.DirectoryAdministrativeUnitId, options ListAdministrativeUnitMemberUsersOperationOptions) (result ListAdministrativeUnitMemberUsersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListAdministrativeUnitMemberUsersCustomPager{},
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

// ListAdministrativeUnitMemberUsersComplete retrieves all the results into a single object
func (c AdministrativeUnitMemberClient) ListAdministrativeUnitMemberUsersComplete(ctx context.Context, id stable.DirectoryAdministrativeUnitId, options ListAdministrativeUnitMemberUsersOperationOptions) (ListAdministrativeUnitMemberUsersCompleteResult, error) {
	return c.ListAdministrativeUnitMemberUsersCompleteMatchingPredicate(ctx, id, options, UserOperationPredicate{})
}

// ListAdministrativeUnitMemberUsersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AdministrativeUnitMemberClient) ListAdministrativeUnitMemberUsersCompleteMatchingPredicate(ctx context.Context, id stable.DirectoryAdministrativeUnitId, options ListAdministrativeUnitMemberUsersOperationOptions, predicate UserOperationPredicate) (result ListAdministrativeUnitMemberUsersCompleteResult, err error) {
	items := make([]stable.User, 0)

	resp, err := c.ListAdministrativeUnitMemberUsers(ctx, id, options)
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

	result = ListAdministrativeUnitMemberUsersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
