package owner

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

type ListOwnerServicePrincipalsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ServicePrincipal
}

type ListOwnerServicePrincipalsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ServicePrincipal
}

type ListOwnerServicePrincipalsOperationOptions struct {
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

func DefaultListOwnerServicePrincipalsOperationOptions() ListOwnerServicePrincipalsOperationOptions {
	return ListOwnerServicePrincipalsOperationOptions{}
}

func (o ListOwnerServicePrincipalsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListOwnerServicePrincipalsOperationOptions) ToOData() *odata.Query {
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

func (o ListOwnerServicePrincipalsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListOwnerServicePrincipalsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListOwnerServicePrincipalsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListOwnerServicePrincipals - Get the items of type microsoft.graph.servicePrincipal in the
// microsoft.graph.directoryObject collection
func (c OwnerClient) ListOwnerServicePrincipals(ctx context.Context, id stable.ApplicationId, options ListOwnerServicePrincipalsOperationOptions) (result ListOwnerServicePrincipalsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListOwnerServicePrincipalsCustomPager{},
		Path:          fmt.Sprintf("%s/owners/servicePrincipal", id.ID()),
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
		Values *[]stable.ServicePrincipal `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListOwnerServicePrincipalsComplete retrieves all the results into a single object
func (c OwnerClient) ListOwnerServicePrincipalsComplete(ctx context.Context, id stable.ApplicationId, options ListOwnerServicePrincipalsOperationOptions) (ListOwnerServicePrincipalsCompleteResult, error) {
	return c.ListOwnerServicePrincipalsCompleteMatchingPredicate(ctx, id, options, ServicePrincipalOperationPredicate{})
}

// ListOwnerServicePrincipalsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OwnerClient) ListOwnerServicePrincipalsCompleteMatchingPredicate(ctx context.Context, id stable.ApplicationId, options ListOwnerServicePrincipalsOperationOptions, predicate ServicePrincipalOperationPredicate) (result ListOwnerServicePrincipalsCompleteResult, err error) {
	items := make([]stable.ServicePrincipal, 0)

	resp, err := c.ListOwnerServicePrincipals(ctx, id, options)
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

	result = ListOwnerServicePrincipalsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
