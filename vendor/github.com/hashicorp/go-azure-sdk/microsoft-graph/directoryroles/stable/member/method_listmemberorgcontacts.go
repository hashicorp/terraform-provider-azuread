package member

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

type ListMemberOrgContactsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.OrgContact
}

type ListMemberOrgContactsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.OrgContact
}

type ListMemberOrgContactsOperationOptions struct {
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

func DefaultListMemberOrgContactsOperationOptions() ListMemberOrgContactsOperationOptions {
	return ListMemberOrgContactsOperationOptions{}
}

func (o ListMemberOrgContactsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListMemberOrgContactsOperationOptions) ToOData() *odata.Query {
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

func (o ListMemberOrgContactsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListMemberOrgContactsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMemberOrgContactsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMemberOrgContacts - Get the items of type microsoft.graph.orgContact in the microsoft.graph.directoryObject
// collection
func (c MemberClient) ListMemberOrgContacts(ctx context.Context, id stable.DirectoryRoleId, options ListMemberOrgContactsOperationOptions) (result ListMemberOrgContactsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListMemberOrgContactsCustomPager{},
		Path:          fmt.Sprintf("%s/members/orgContact", id.ID()),
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
		Values *[]stable.OrgContact `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMemberOrgContactsComplete retrieves all the results into a single object
func (c MemberClient) ListMemberOrgContactsComplete(ctx context.Context, id stable.DirectoryRoleId, options ListMemberOrgContactsOperationOptions) (ListMemberOrgContactsCompleteResult, error) {
	return c.ListMemberOrgContactsCompleteMatchingPredicate(ctx, id, options, OrgContactOperationPredicate{})
}

// ListMemberOrgContactsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MemberClient) ListMemberOrgContactsCompleteMatchingPredicate(ctx context.Context, id stable.DirectoryRoleId, options ListMemberOrgContactsOperationOptions, predicate OrgContactOperationPredicate) (result ListMemberOrgContactsCompleteResult, err error) {
	items := make([]stable.OrgContact, 0)

	resp, err := c.ListMemberOrgContacts(ctx, id, options)
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

	result = ListMemberOrgContactsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
