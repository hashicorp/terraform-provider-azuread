package synchronizationjob

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

type ListSynchronizationJobsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.SynchronizationJob
}

type ListSynchronizationJobsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.SynchronizationJob
}

type ListSynchronizationJobsOperationOptions struct {
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

func DefaultListSynchronizationJobsOperationOptions() ListSynchronizationJobsOperationOptions {
	return ListSynchronizationJobsOperationOptions{}
}

func (o ListSynchronizationJobsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSynchronizationJobsOperationOptions) ToOData() *odata.Query {
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

func (o ListSynchronizationJobsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListSynchronizationJobsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSynchronizationJobsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSynchronizationJobs - List synchronization jobs. List existing jobs for a given application instance (service
// principal).
func (c SynchronizationJobClient) ListSynchronizationJobs(ctx context.Context, id stable.ServicePrincipalId, options ListSynchronizationJobsOperationOptions) (result ListSynchronizationJobsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListSynchronizationJobsCustomPager{},
		Path:          fmt.Sprintf("%s/synchronization/jobs", id.ID()),
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
		Values *[]stable.SynchronizationJob `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSynchronizationJobsComplete retrieves all the results into a single object
func (c SynchronizationJobClient) ListSynchronizationJobsComplete(ctx context.Context, id stable.ServicePrincipalId, options ListSynchronizationJobsOperationOptions) (ListSynchronizationJobsCompleteResult, error) {
	return c.ListSynchronizationJobsCompleteMatchingPredicate(ctx, id, options, SynchronizationJobOperationPredicate{})
}

// ListSynchronizationJobsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SynchronizationJobClient) ListSynchronizationJobsCompleteMatchingPredicate(ctx context.Context, id stable.ServicePrincipalId, options ListSynchronizationJobsOperationOptions, predicate SynchronizationJobOperationPredicate) (result ListSynchronizationJobsCompleteResult, err error) {
	items := make([]stable.SynchronizationJob, 0)

	resp, err := c.ListSynchronizationJobs(ctx, id, options)
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

	result = ListSynchronizationJobsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
