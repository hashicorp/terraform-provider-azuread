package transitivemember

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

type ListTransitiveMemberOrgContactsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.OrgContact
}

type ListTransitiveMemberOrgContactsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.OrgContact
}

type ListTransitiveMemberOrgContactsOperationOptions struct {
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

func DefaultListTransitiveMemberOrgContactsOperationOptions() ListTransitiveMemberOrgContactsOperationOptions {
	return ListTransitiveMemberOrgContactsOperationOptions{}
}

func (o ListTransitiveMemberOrgContactsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListTransitiveMemberOrgContactsOperationOptions) ToOData() *odata.Query {
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

func (o ListTransitiveMemberOrgContactsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListTransitiveMemberOrgContactsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTransitiveMemberOrgContactsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTransitiveMemberOrgContacts - Get the items of type microsoft.graph.orgContact in the
// microsoft.graph.directoryObject collection
func (c TransitiveMemberClient) ListTransitiveMemberOrgContacts(ctx context.Context, id beta.GroupId, options ListTransitiveMemberOrgContactsOperationOptions) (result ListTransitiveMemberOrgContactsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListTransitiveMemberOrgContactsCustomPager{},
		Path:          fmt.Sprintf("%s/transitiveMembers/orgContact", id.ID()),
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
		Values *[]beta.OrgContact `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTransitiveMemberOrgContactsComplete retrieves all the results into a single object
func (c TransitiveMemberClient) ListTransitiveMemberOrgContactsComplete(ctx context.Context, id beta.GroupId, options ListTransitiveMemberOrgContactsOperationOptions) (ListTransitiveMemberOrgContactsCompleteResult, error) {
	return c.ListTransitiveMemberOrgContactsCompleteMatchingPredicate(ctx, id, options, OrgContactOperationPredicate{})
}

// ListTransitiveMemberOrgContactsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TransitiveMemberClient) ListTransitiveMemberOrgContactsCompleteMatchingPredicate(ctx context.Context, id beta.GroupId, options ListTransitiveMemberOrgContactsOperationOptions, predicate OrgContactOperationPredicate) (result ListTransitiveMemberOrgContactsCompleteResult, err error) {
	items := make([]beta.OrgContact, 0)

	resp, err := c.ListTransitiveMemberOrgContacts(ctx, id, options)
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

	result = ListTransitiveMemberOrgContactsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
