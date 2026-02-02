package transitivemember

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListTransitiveMemberDevicesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Device
}

type ListTransitiveMemberDevicesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Device
}

type ListTransitiveMemberDevicesOperationOptions struct {
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

func DefaultListTransitiveMemberDevicesOperationOptions() ListTransitiveMemberDevicesOperationOptions {
	return ListTransitiveMemberDevicesOperationOptions{}
}

func (o ListTransitiveMemberDevicesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListTransitiveMemberDevicesOperationOptions) ToOData() *odata.Query {
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

func (o ListTransitiveMemberDevicesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListTransitiveMemberDevicesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTransitiveMemberDevicesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTransitiveMemberDevices - Get the items of type microsoft.graph.device in the microsoft.graph.directoryObject
// collection
func (c TransitiveMemberClient) ListTransitiveMemberDevices(ctx context.Context, id beta.GroupId, options ListTransitiveMemberDevicesOperationOptions) (result ListTransitiveMemberDevicesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListTransitiveMemberDevicesCustomPager{},
		Path:          fmt.Sprintf("%s/transitiveMembers/device", id.ID()),
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
		Values *[]beta.Device `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTransitiveMemberDevicesComplete retrieves all the results into a single object
func (c TransitiveMemberClient) ListTransitiveMemberDevicesComplete(ctx context.Context, id beta.GroupId, options ListTransitiveMemberDevicesOperationOptions) (ListTransitiveMemberDevicesCompleteResult, error) {
	return c.ListTransitiveMemberDevicesCompleteMatchingPredicate(ctx, id, options, DeviceOperationPredicate{})
}

// ListTransitiveMemberDevicesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TransitiveMemberClient) ListTransitiveMemberDevicesCompleteMatchingPredicate(ctx context.Context, id beta.GroupId, options ListTransitiveMemberDevicesOperationOptions, predicate DeviceOperationPredicate) (result ListTransitiveMemberDevicesCompleteResult, err error) {
	items := make([]beta.Device, 0)

	resp, err := c.ListTransitiveMemberDevices(ctx, id, options)
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

	result = ListTransitiveMemberDevicesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
