package privilegedaccessgroupeligibilityscheduleinstance

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

type ListPrivilegedAccessGroupEligibilityScheduleInstancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.PrivilegedAccessGroupEligibilityScheduleInstance
}

type ListPrivilegedAccessGroupEligibilityScheduleInstancesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.PrivilegedAccessGroupEligibilityScheduleInstance
}

type ListPrivilegedAccessGroupEligibilityScheduleInstancesOperationOptions struct {
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

func DefaultListPrivilegedAccessGroupEligibilityScheduleInstancesOperationOptions() ListPrivilegedAccessGroupEligibilityScheduleInstancesOperationOptions {
	return ListPrivilegedAccessGroupEligibilityScheduleInstancesOperationOptions{}
}

func (o ListPrivilegedAccessGroupEligibilityScheduleInstancesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListPrivilegedAccessGroupEligibilityScheduleInstancesOperationOptions) ToOData() *odata.Query {
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

func (o ListPrivilegedAccessGroupEligibilityScheduleInstancesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListPrivilegedAccessGroupEligibilityScheduleInstancesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPrivilegedAccessGroupEligibilityScheduleInstancesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPrivilegedAccessGroupEligibilityScheduleInstances - List eligibilityScheduleInstances. Get a list of the
// privilegedAccessGroupEligibilityScheduleInstance objects and their properties.
func (c PrivilegedAccessGroupEligibilityScheduleInstanceClient) ListPrivilegedAccessGroupEligibilityScheduleInstances(ctx context.Context, options ListPrivilegedAccessGroupEligibilityScheduleInstancesOperationOptions) (result ListPrivilegedAccessGroupEligibilityScheduleInstancesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListPrivilegedAccessGroupEligibilityScheduleInstancesCustomPager{},
		Path:          "/identityGovernance/privilegedAccess/group/eligibilityScheduleInstances",
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
		Values *[]stable.PrivilegedAccessGroupEligibilityScheduleInstance `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPrivilegedAccessGroupEligibilityScheduleInstancesComplete retrieves all the results into a single object
func (c PrivilegedAccessGroupEligibilityScheduleInstanceClient) ListPrivilegedAccessGroupEligibilityScheduleInstancesComplete(ctx context.Context, options ListPrivilegedAccessGroupEligibilityScheduleInstancesOperationOptions) (ListPrivilegedAccessGroupEligibilityScheduleInstancesCompleteResult, error) {
	return c.ListPrivilegedAccessGroupEligibilityScheduleInstancesCompleteMatchingPredicate(ctx, options, PrivilegedAccessGroupEligibilityScheduleInstanceOperationPredicate{})
}

// ListPrivilegedAccessGroupEligibilityScheduleInstancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PrivilegedAccessGroupEligibilityScheduleInstanceClient) ListPrivilegedAccessGroupEligibilityScheduleInstancesCompleteMatchingPredicate(ctx context.Context, options ListPrivilegedAccessGroupEligibilityScheduleInstancesOperationOptions, predicate PrivilegedAccessGroupEligibilityScheduleInstanceOperationPredicate) (result ListPrivilegedAccessGroupEligibilityScheduleInstancesCompleteResult, err error) {
	items := make([]stable.PrivilegedAccessGroupEligibilityScheduleInstance, 0)

	resp, err := c.ListPrivilegedAccessGroupEligibilityScheduleInstances(ctx, options)
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

	result = ListPrivilegedAccessGroupEligibilityScheduleInstancesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
