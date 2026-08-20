package directoryroleeligibilityscheduleinstance

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

type ListDirectoryRoleEligibilityScheduleInstancesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.UnifiedRoleEligibilityScheduleInstance
}

type ListDirectoryRoleEligibilityScheduleInstancesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.UnifiedRoleEligibilityScheduleInstance
}

type ListDirectoryRoleEligibilityScheduleInstancesOperationOptions struct {
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

func DefaultListDirectoryRoleEligibilityScheduleInstancesOperationOptions() ListDirectoryRoleEligibilityScheduleInstancesOperationOptions {
	return ListDirectoryRoleEligibilityScheduleInstancesOperationOptions{}
}

func (o ListDirectoryRoleEligibilityScheduleInstancesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDirectoryRoleEligibilityScheduleInstancesOperationOptions) ToOData() *odata.Query {
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

func (o ListDirectoryRoleEligibilityScheduleInstancesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDirectoryRoleEligibilityScheduleInstancesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDirectoryRoleEligibilityScheduleInstancesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDirectoryRoleEligibilityScheduleInstances - List roleEligibilityScheduleInstances. Get the instances of role
// eligibilities.
func (c DirectoryRoleEligibilityScheduleInstanceClient) ListDirectoryRoleEligibilityScheduleInstances(ctx context.Context, options ListDirectoryRoleEligibilityScheduleInstancesOperationOptions) (result ListDirectoryRoleEligibilityScheduleInstancesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDirectoryRoleEligibilityScheduleInstancesCustomPager{},
		Path:          "/roleManagement/directory/roleEligibilityScheduleInstances",
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
		Values *[]stable.UnifiedRoleEligibilityScheduleInstance `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDirectoryRoleEligibilityScheduleInstancesComplete retrieves all the results into a single object
func (c DirectoryRoleEligibilityScheduleInstanceClient) ListDirectoryRoleEligibilityScheduleInstancesComplete(ctx context.Context, options ListDirectoryRoleEligibilityScheduleInstancesOperationOptions) (ListDirectoryRoleEligibilityScheduleInstancesCompleteResult, error) {
	return c.ListDirectoryRoleEligibilityScheduleInstancesCompleteMatchingPredicate(ctx, options, UnifiedRoleEligibilityScheduleInstanceOperationPredicate{})
}

// ListDirectoryRoleEligibilityScheduleInstancesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DirectoryRoleEligibilityScheduleInstanceClient) ListDirectoryRoleEligibilityScheduleInstancesCompleteMatchingPredicate(ctx context.Context, options ListDirectoryRoleEligibilityScheduleInstancesOperationOptions, predicate UnifiedRoleEligibilityScheduleInstanceOperationPredicate) (result ListDirectoryRoleEligibilityScheduleInstancesCompleteResult, err error) {
	items := make([]stable.UnifiedRoleEligibilityScheduleInstance, 0)

	resp, err := c.ListDirectoryRoleEligibilityScheduleInstances(ctx, options)
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

	result = ListDirectoryRoleEligibilityScheduleInstancesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
