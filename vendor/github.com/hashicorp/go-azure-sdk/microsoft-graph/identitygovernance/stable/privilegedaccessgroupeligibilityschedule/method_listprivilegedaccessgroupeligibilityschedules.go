package privilegedaccessgroupeligibilityschedule

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

type ListPrivilegedAccessGroupEligibilitySchedulesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.PrivilegedAccessGroupEligibilitySchedule
}

type ListPrivilegedAccessGroupEligibilitySchedulesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.PrivilegedAccessGroupEligibilitySchedule
}

type ListPrivilegedAccessGroupEligibilitySchedulesOperationOptions struct {
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

func DefaultListPrivilegedAccessGroupEligibilitySchedulesOperationOptions() ListPrivilegedAccessGroupEligibilitySchedulesOperationOptions {
	return ListPrivilegedAccessGroupEligibilitySchedulesOperationOptions{}
}

func (o ListPrivilegedAccessGroupEligibilitySchedulesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListPrivilegedAccessGroupEligibilitySchedulesOperationOptions) ToOData() *odata.Query {
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

func (o ListPrivilegedAccessGroupEligibilitySchedulesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListPrivilegedAccessGroupEligibilitySchedulesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPrivilegedAccessGroupEligibilitySchedulesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPrivilegedAccessGroupEligibilitySchedules - List eligibilitySchedules. Get a list of the
// privilegedAccessGroupEligibilitySchedule objects and their properties.
func (c PrivilegedAccessGroupEligibilityScheduleClient) ListPrivilegedAccessGroupEligibilitySchedules(ctx context.Context, options ListPrivilegedAccessGroupEligibilitySchedulesOperationOptions) (result ListPrivilegedAccessGroupEligibilitySchedulesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListPrivilegedAccessGroupEligibilitySchedulesCustomPager{},
		Path:          "/identityGovernance/privilegedAccess/group/eligibilitySchedules",
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
		Values *[]stable.PrivilegedAccessGroupEligibilitySchedule `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPrivilegedAccessGroupEligibilitySchedulesComplete retrieves all the results into a single object
func (c PrivilegedAccessGroupEligibilityScheduleClient) ListPrivilegedAccessGroupEligibilitySchedulesComplete(ctx context.Context, options ListPrivilegedAccessGroupEligibilitySchedulesOperationOptions) (ListPrivilegedAccessGroupEligibilitySchedulesCompleteResult, error) {
	return c.ListPrivilegedAccessGroupEligibilitySchedulesCompleteMatchingPredicate(ctx, options, PrivilegedAccessGroupEligibilityScheduleOperationPredicate{})
}

// ListPrivilegedAccessGroupEligibilitySchedulesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PrivilegedAccessGroupEligibilityScheduleClient) ListPrivilegedAccessGroupEligibilitySchedulesCompleteMatchingPredicate(ctx context.Context, options ListPrivilegedAccessGroupEligibilitySchedulesOperationOptions, predicate PrivilegedAccessGroupEligibilityScheduleOperationPredicate) (result ListPrivilegedAccessGroupEligibilitySchedulesCompleteResult, err error) {
	items := make([]stable.PrivilegedAccessGroupEligibilitySchedule, 0)

	resp, err := c.ListPrivilegedAccessGroupEligibilitySchedules(ctx, options)
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

	result = ListPrivilegedAccessGroupEligibilitySchedulesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
