package owner

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListOwnerDevicesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Device
}

type ListOwnerDevicesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Device
}

type ListOwnerDevicesOperationOptions struct {
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

func DefaultListOwnerDevicesOperationOptions() ListOwnerDevicesOperationOptions {
	return ListOwnerDevicesOperationOptions{}
}

func (o ListOwnerDevicesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListOwnerDevicesOperationOptions) ToOData() *odata.Query {
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

func (o ListOwnerDevicesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListOwnerDevicesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListOwnerDevicesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListOwnerDevices - Get the items of type microsoft.graph.device in the microsoft.graph.directoryObject collection
func (c OwnerClient) ListOwnerDevices(ctx context.Context, id beta.GroupId, options ListOwnerDevicesOperationOptions) (result ListOwnerDevicesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListOwnerDevicesCustomPager{},
		Path:          fmt.Sprintf("%s/owners/device", id.ID()),
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

// ListOwnerDevicesComplete retrieves all the results into a single object
func (c OwnerClient) ListOwnerDevicesComplete(ctx context.Context, id beta.GroupId, options ListOwnerDevicesOperationOptions) (ListOwnerDevicesCompleteResult, error) {
	return c.ListOwnerDevicesCompleteMatchingPredicate(ctx, id, options, DeviceOperationPredicate{})
}

// ListOwnerDevicesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OwnerClient) ListOwnerDevicesCompleteMatchingPredicate(ctx context.Context, id beta.GroupId, options ListOwnerDevicesOperationOptions, predicate DeviceOperationPredicate) (result ListOwnerDevicesCompleteResult, err error) {
	items := make([]beta.Device, 0)

	resp, err := c.ListOwnerDevices(ctx, id, options)
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

	result = ListOwnerDevicesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
