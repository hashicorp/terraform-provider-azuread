package directoryroleeligibilityschedule

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

type ListDirectoryRoleEligibilitySchedulesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.UnifiedRoleEligibilitySchedule
}

type ListDirectoryRoleEligibilitySchedulesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.UnifiedRoleEligibilitySchedule
}

type ListDirectoryRoleEligibilitySchedulesOperationOptions struct {
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

func DefaultListDirectoryRoleEligibilitySchedulesOperationOptions() ListDirectoryRoleEligibilitySchedulesOperationOptions {
	return ListDirectoryRoleEligibilitySchedulesOperationOptions{}
}

func (o ListDirectoryRoleEligibilitySchedulesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDirectoryRoleEligibilitySchedulesOperationOptions) ToOData() *odata.Query {
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

func (o ListDirectoryRoleEligibilitySchedulesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDirectoryRoleEligibilitySchedulesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDirectoryRoleEligibilitySchedulesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDirectoryRoleEligibilitySchedules - List roleEligibilitySchedules. Get the unifiedRoleEligibilitySchedule
// resources from the roleEligibilitySchedules navigation property.
func (c DirectoryRoleEligibilityScheduleClient) ListDirectoryRoleEligibilitySchedules(ctx context.Context, options ListDirectoryRoleEligibilitySchedulesOperationOptions) (result ListDirectoryRoleEligibilitySchedulesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDirectoryRoleEligibilitySchedulesCustomPager{},
		Path:          "/roleManagement/directory/roleEligibilitySchedules",
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
		Values *[]stable.UnifiedRoleEligibilitySchedule `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDirectoryRoleEligibilitySchedulesComplete retrieves all the results into a single object
func (c DirectoryRoleEligibilityScheduleClient) ListDirectoryRoleEligibilitySchedulesComplete(ctx context.Context, options ListDirectoryRoleEligibilitySchedulesOperationOptions) (ListDirectoryRoleEligibilitySchedulesCompleteResult, error) {
	return c.ListDirectoryRoleEligibilitySchedulesCompleteMatchingPredicate(ctx, options, UnifiedRoleEligibilityScheduleOperationPredicate{})
}

// ListDirectoryRoleEligibilitySchedulesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DirectoryRoleEligibilityScheduleClient) ListDirectoryRoleEligibilitySchedulesCompleteMatchingPredicate(ctx context.Context, options ListDirectoryRoleEligibilitySchedulesOperationOptions, predicate UnifiedRoleEligibilityScheduleOperationPredicate) (result ListDirectoryRoleEligibilitySchedulesCompleteResult, err error) {
	items := make([]stable.UnifiedRoleEligibilitySchedule, 0)

	resp, err := c.ListDirectoryRoleEligibilitySchedules(ctx, options)
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

	result = ListDirectoryRoleEligibilitySchedulesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
