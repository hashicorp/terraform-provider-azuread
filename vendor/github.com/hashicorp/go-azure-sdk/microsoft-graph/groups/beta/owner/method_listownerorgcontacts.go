package owner

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

type ListOwnerOrgContactsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.OrgContact
}

type ListOwnerOrgContactsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.OrgContact
}

type ListOwnerOrgContactsOperationOptions struct {
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

func DefaultListOwnerOrgContactsOperationOptions() ListOwnerOrgContactsOperationOptions {
	return ListOwnerOrgContactsOperationOptions{}
}

func (o ListOwnerOrgContactsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListOwnerOrgContactsOperationOptions) ToOData() *odata.Query {
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

func (o ListOwnerOrgContactsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListOwnerOrgContactsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListOwnerOrgContactsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListOwnerOrgContacts - Get the items of type microsoft.graph.orgContact in the microsoft.graph.directoryObject
// collection
func (c OwnerClient) ListOwnerOrgContacts(ctx context.Context, id beta.GroupId, options ListOwnerOrgContactsOperationOptions) (result ListOwnerOrgContactsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListOwnerOrgContactsCustomPager{},
		Path:          fmt.Sprintf("%s/owners/orgContact", id.ID()),
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

// ListOwnerOrgContactsComplete retrieves all the results into a single object
func (c OwnerClient) ListOwnerOrgContactsComplete(ctx context.Context, id beta.GroupId, options ListOwnerOrgContactsOperationOptions) (ListOwnerOrgContactsCompleteResult, error) {
	return c.ListOwnerOrgContactsCompleteMatchingPredicate(ctx, id, options, OrgContactOperationPredicate{})
}

// ListOwnerOrgContactsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OwnerClient) ListOwnerOrgContactsCompleteMatchingPredicate(ctx context.Context, id beta.GroupId, options ListOwnerOrgContactsOperationOptions, predicate OrgContactOperationPredicate) (result ListOwnerOrgContactsCompleteResult, err error) {
	items := make([]beta.OrgContact, 0)

	resp, err := c.ListOwnerOrgContacts(ctx, id, options)
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

	result = ListOwnerOrgContactsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
