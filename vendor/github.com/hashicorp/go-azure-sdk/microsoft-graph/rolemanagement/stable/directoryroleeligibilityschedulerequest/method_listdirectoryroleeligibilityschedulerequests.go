package directoryroleeligibilityschedulerequest

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

type ListDirectoryRoleEligibilityScheduleRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.UnifiedRoleEligibilityScheduleRequest
}

type ListDirectoryRoleEligibilityScheduleRequestsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.UnifiedRoleEligibilityScheduleRequest
}

type ListDirectoryRoleEligibilityScheduleRequestsOperationOptions struct {
	Count     *bool
	Expand    *odata.Expand
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Select    *[]string
	Skip      *int64
	Top       *int64
}

func DefaultListDirectoryRoleEligibilityScheduleRequestsOperationOptions() ListDirectoryRoleEligibilityScheduleRequestsOperationOptions {
	return ListDirectoryRoleEligibilityScheduleRequestsOperationOptions{}
}

func (o ListDirectoryRoleEligibilityScheduleRequestsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDirectoryRoleEligibilityScheduleRequestsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
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

func (o ListDirectoryRoleEligibilityScheduleRequestsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDirectoryRoleEligibilityScheduleRequestsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDirectoryRoleEligibilityScheduleRequestsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDirectoryRoleEligibilityScheduleRequests - List roleEligibilityScheduleRequests. In PIM, retrieve the requests
// for role eligibilities for principals made through the unifiedRoleEligibilityScheduleRequest object.
func (c DirectoryRoleEligibilityScheduleRequestClient) ListDirectoryRoleEligibilityScheduleRequests(ctx context.Context, options ListDirectoryRoleEligibilityScheduleRequestsOperationOptions) (result ListDirectoryRoleEligibilityScheduleRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDirectoryRoleEligibilityScheduleRequestsCustomPager{},
		Path:          "/roleManagement/directory/roleEligibilityScheduleRequests",
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
		Values *[]stable.UnifiedRoleEligibilityScheduleRequest `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDirectoryRoleEligibilityScheduleRequestsComplete retrieves all the results into a single object
func (c DirectoryRoleEligibilityScheduleRequestClient) ListDirectoryRoleEligibilityScheduleRequestsComplete(ctx context.Context, options ListDirectoryRoleEligibilityScheduleRequestsOperationOptions) (ListDirectoryRoleEligibilityScheduleRequestsCompleteResult, error) {
	return c.ListDirectoryRoleEligibilityScheduleRequestsCompleteMatchingPredicate(ctx, options, UnifiedRoleEligibilityScheduleRequestOperationPredicate{})
}

// ListDirectoryRoleEligibilityScheduleRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DirectoryRoleEligibilityScheduleRequestClient) ListDirectoryRoleEligibilityScheduleRequestsCompleteMatchingPredicate(ctx context.Context, options ListDirectoryRoleEligibilityScheduleRequestsOperationOptions, predicate UnifiedRoleEligibilityScheduleRequestOperationPredicate) (result ListDirectoryRoleEligibilityScheduleRequestsCompleteResult, err error) {
	items := make([]stable.UnifiedRoleEligibilityScheduleRequest, 0)

	resp, err := c.ListDirectoryRoleEligibilityScheduleRequests(ctx, options)
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

	result = ListDirectoryRoleEligibilityScheduleRequestsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
