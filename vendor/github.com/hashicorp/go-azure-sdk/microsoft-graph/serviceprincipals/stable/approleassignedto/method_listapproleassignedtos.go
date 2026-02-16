package approleassignedto

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

type ListAppRoleAssignedTosOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AppRoleAssignment
}

type ListAppRoleAssignedTosCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AppRoleAssignment
}

type ListAppRoleAssignedTosOperationOptions struct {
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

func DefaultListAppRoleAssignedTosOperationOptions() ListAppRoleAssignedTosOperationOptions {
	return ListAppRoleAssignedTosOperationOptions{}
}

func (o ListAppRoleAssignedTosOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListAppRoleAssignedTosOperationOptions) ToOData() *odata.Query {
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

func (o ListAppRoleAssignedTosOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListAppRoleAssignedTosCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAppRoleAssignedTosCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAppRoleAssignedTos - Get appRoleAssignment. Read the properties and relationships of an appRoleAssignment object.
func (c AppRoleAssignedToClient) ListAppRoleAssignedTos(ctx context.Context, id stable.ServicePrincipalId, options ListAppRoleAssignedTosOperationOptions) (result ListAppRoleAssignedTosOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListAppRoleAssignedTosCustomPager{},
		Path:          fmt.Sprintf("%s/appRoleAssignedTo", id.ID()),
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
		Values *[]stable.AppRoleAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAppRoleAssignedTosComplete retrieves all the results into a single object
func (c AppRoleAssignedToClient) ListAppRoleAssignedTosComplete(ctx context.Context, id stable.ServicePrincipalId, options ListAppRoleAssignedTosOperationOptions) (ListAppRoleAssignedTosCompleteResult, error) {
	return c.ListAppRoleAssignedTosCompleteMatchingPredicate(ctx, id, options, AppRoleAssignmentOperationPredicate{})
}

// ListAppRoleAssignedTosCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AppRoleAssignedToClient) ListAppRoleAssignedTosCompleteMatchingPredicate(ctx context.Context, id stable.ServicePrincipalId, options ListAppRoleAssignedTosOperationOptions, predicate AppRoleAssignmentOperationPredicate) (result ListAppRoleAssignedTosCompleteResult, err error) {
	items := make([]stable.AppRoleAssignment, 0)

	resp, err := c.ListAppRoleAssignedTos(ctx, id, options)
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

	result = ListAppRoleAssignedTosCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
