package owner

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

type ListOwnerAppRoleAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AppRoleAssignment
}

type ListOwnerAppRoleAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AppRoleAssignment
}

type ListOwnerAppRoleAssignmentsOperationOptions struct {
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

func DefaultListOwnerAppRoleAssignmentsOperationOptions() ListOwnerAppRoleAssignmentsOperationOptions {
	return ListOwnerAppRoleAssignmentsOperationOptions{}
}

func (o ListOwnerAppRoleAssignmentsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListOwnerAppRoleAssignmentsOperationOptions) ToOData() *odata.Query {
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

func (o ListOwnerAppRoleAssignmentsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListOwnerAppRoleAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListOwnerAppRoleAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListOwnerAppRoleAssignments - Get the items of type microsoft.graph.appRoleAssignment in the
// microsoft.graph.directoryObject collection
func (c OwnerClient) ListOwnerAppRoleAssignments(ctx context.Context, id stable.ApplicationId, options ListOwnerAppRoleAssignmentsOperationOptions) (result ListOwnerAppRoleAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListOwnerAppRoleAssignmentsCustomPager{},
		Path:          fmt.Sprintf("%s/owners/appRoleAssignment", id.ID()),
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
		Values *[]stable.AppRoleAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListOwnerAppRoleAssignmentsComplete retrieves all the results into a single object
func (c OwnerClient) ListOwnerAppRoleAssignmentsComplete(ctx context.Context, id stable.ApplicationId, options ListOwnerAppRoleAssignmentsOperationOptions) (ListOwnerAppRoleAssignmentsCompleteResult, error) {
	return c.ListOwnerAppRoleAssignmentsCompleteMatchingPredicate(ctx, id, options, AppRoleAssignmentOperationPredicate{})
}

// ListOwnerAppRoleAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OwnerClient) ListOwnerAppRoleAssignmentsCompleteMatchingPredicate(ctx context.Context, id stable.ApplicationId, options ListOwnerAppRoleAssignmentsOperationOptions, predicate AppRoleAssignmentOperationPredicate) (result ListOwnerAppRoleAssignmentsCompleteResult, err error) {
	items := make([]stable.AppRoleAssignment, 0)

	resp, err := c.ListOwnerAppRoleAssignments(ctx, id, options)
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

	result = ListOwnerAppRoleAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
