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

type ListTransitiveMemberServicePrincipalsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ServicePrincipal
}

type ListTransitiveMemberServicePrincipalsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ServicePrincipal
}

type ListTransitiveMemberServicePrincipalsOperationOptions struct {
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

func DefaultListTransitiveMemberServicePrincipalsOperationOptions() ListTransitiveMemberServicePrincipalsOperationOptions {
	return ListTransitiveMemberServicePrincipalsOperationOptions{}
}

func (o ListTransitiveMemberServicePrincipalsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListTransitiveMemberServicePrincipalsOperationOptions) ToOData() *odata.Query {
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

func (o ListTransitiveMemberServicePrincipalsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListTransitiveMemberServicePrincipalsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTransitiveMemberServicePrincipalsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTransitiveMemberServicePrincipals - Get the items of type microsoft.graph.servicePrincipal in the
// microsoft.graph.directoryObject collection
func (c TransitiveMemberClient) ListTransitiveMemberServicePrincipals(ctx context.Context, id beta.GroupId, options ListTransitiveMemberServicePrincipalsOperationOptions) (result ListTransitiveMemberServicePrincipalsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListTransitiveMemberServicePrincipalsCustomPager{},
		Path:          fmt.Sprintf("%s/transitiveMembers/servicePrincipal", id.ID()),
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
		Values *[]beta.ServicePrincipal `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTransitiveMemberServicePrincipalsComplete retrieves all the results into a single object
func (c TransitiveMemberClient) ListTransitiveMemberServicePrincipalsComplete(ctx context.Context, id beta.GroupId, options ListTransitiveMemberServicePrincipalsOperationOptions) (ListTransitiveMemberServicePrincipalsCompleteResult, error) {
	return c.ListTransitiveMemberServicePrincipalsCompleteMatchingPredicate(ctx, id, options, ServicePrincipalOperationPredicate{})
}

// ListTransitiveMemberServicePrincipalsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TransitiveMemberClient) ListTransitiveMemberServicePrincipalsCompleteMatchingPredicate(ctx context.Context, id beta.GroupId, options ListTransitiveMemberServicePrincipalsOperationOptions, predicate ServicePrincipalOperationPredicate) (result ListTransitiveMemberServicePrincipalsCompleteResult, err error) {
	items := make([]beta.ServicePrincipal, 0)

	resp, err := c.ListTransitiveMemberServicePrincipals(ctx, id, options)
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

	result = ListTransitiveMemberServicePrincipalsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
