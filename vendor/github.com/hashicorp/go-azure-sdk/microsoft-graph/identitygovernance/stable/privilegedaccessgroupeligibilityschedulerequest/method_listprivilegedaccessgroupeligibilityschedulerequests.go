package privilegedaccessgroupeligibilityschedulerequest

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

type ListPrivilegedAccessGroupEligibilityScheduleRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.PrivilegedAccessGroupEligibilityScheduleRequest
}

type ListPrivilegedAccessGroupEligibilityScheduleRequestsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.PrivilegedAccessGroupEligibilityScheduleRequest
}

type ListPrivilegedAccessGroupEligibilityScheduleRequestsOperationOptions struct {
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

func DefaultListPrivilegedAccessGroupEligibilityScheduleRequestsOperationOptions() ListPrivilegedAccessGroupEligibilityScheduleRequestsOperationOptions {
	return ListPrivilegedAccessGroupEligibilityScheduleRequestsOperationOptions{}
}

func (o ListPrivilegedAccessGroupEligibilityScheduleRequestsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListPrivilegedAccessGroupEligibilityScheduleRequestsOperationOptions) ToOData() *odata.Query {
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

func (o ListPrivilegedAccessGroupEligibilityScheduleRequestsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListPrivilegedAccessGroupEligibilityScheduleRequestsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPrivilegedAccessGroupEligibilityScheduleRequestsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPrivilegedAccessGroupEligibilityScheduleRequests - List eligibilityScheduleRequests. Get a list of the
// privilegedAccessGroupEligibilityScheduleRequest objects and their properties.
func (c PrivilegedAccessGroupEligibilityScheduleRequestClient) ListPrivilegedAccessGroupEligibilityScheduleRequests(ctx context.Context, options ListPrivilegedAccessGroupEligibilityScheduleRequestsOperationOptions) (result ListPrivilegedAccessGroupEligibilityScheduleRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListPrivilegedAccessGroupEligibilityScheduleRequestsCustomPager{},
		Path:          "/identityGovernance/privilegedAccess/group/eligibilityScheduleRequests",
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
		Values *[]stable.PrivilegedAccessGroupEligibilityScheduleRequest `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPrivilegedAccessGroupEligibilityScheduleRequestsComplete retrieves all the results into a single object
func (c PrivilegedAccessGroupEligibilityScheduleRequestClient) ListPrivilegedAccessGroupEligibilityScheduleRequestsComplete(ctx context.Context, options ListPrivilegedAccessGroupEligibilityScheduleRequestsOperationOptions) (ListPrivilegedAccessGroupEligibilityScheduleRequestsCompleteResult, error) {
	return c.ListPrivilegedAccessGroupEligibilityScheduleRequestsCompleteMatchingPredicate(ctx, options, PrivilegedAccessGroupEligibilityScheduleRequestOperationPredicate{})
}

// ListPrivilegedAccessGroupEligibilityScheduleRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PrivilegedAccessGroupEligibilityScheduleRequestClient) ListPrivilegedAccessGroupEligibilityScheduleRequestsCompleteMatchingPredicate(ctx context.Context, options ListPrivilegedAccessGroupEligibilityScheduleRequestsOperationOptions, predicate PrivilegedAccessGroupEligibilityScheduleRequestOperationPredicate) (result ListPrivilegedAccessGroupEligibilityScheduleRequestsCompleteResult, err error) {
	items := make([]stable.PrivilegedAccessGroupEligibilityScheduleRequest, 0)

	resp, err := c.ListPrivilegedAccessGroupEligibilityScheduleRequests(ctx, options)
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

	result = ListPrivilegedAccessGroupEligibilityScheduleRequestsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
