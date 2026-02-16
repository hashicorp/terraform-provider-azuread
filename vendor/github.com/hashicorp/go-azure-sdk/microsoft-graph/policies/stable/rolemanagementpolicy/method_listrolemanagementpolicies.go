package rolemanagementpolicy

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

type ListRoleManagementPoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.UnifiedRoleManagementPolicy
}

type ListRoleManagementPoliciesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.UnifiedRoleManagementPolicy
}

type ListRoleManagementPoliciesOperationOptions struct {
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

func DefaultListRoleManagementPoliciesOperationOptions() ListRoleManagementPoliciesOperationOptions {
	return ListRoleManagementPoliciesOperationOptions{}
}

func (o ListRoleManagementPoliciesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListRoleManagementPoliciesOperationOptions) ToOData() *odata.Query {
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

func (o ListRoleManagementPoliciesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListRoleManagementPoliciesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListRoleManagementPoliciesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListRoleManagementPolicies - List roleManagementPolicies. Get the details of the policies in PIM that can be applied
// to Microsoft Entra roles or group membership or ownership. To retrieve policies that apply to Azure RBAC, use the
// Azure REST PIM API for role management policies.
func (c RoleManagementPolicyClient) ListRoleManagementPolicies(ctx context.Context, options ListRoleManagementPoliciesOperationOptions) (result ListRoleManagementPoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListRoleManagementPoliciesCustomPager{},
		Path:          "/policies/roleManagementPolicies",
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
		Values *[]stable.UnifiedRoleManagementPolicy `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListRoleManagementPoliciesComplete retrieves all the results into a single object
func (c RoleManagementPolicyClient) ListRoleManagementPoliciesComplete(ctx context.Context, options ListRoleManagementPoliciesOperationOptions) (ListRoleManagementPoliciesCompleteResult, error) {
	return c.ListRoleManagementPoliciesCompleteMatchingPredicate(ctx, options, UnifiedRoleManagementPolicyOperationPredicate{})
}

// ListRoleManagementPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleManagementPolicyClient) ListRoleManagementPoliciesCompleteMatchingPredicate(ctx context.Context, options ListRoleManagementPoliciesOperationOptions, predicate UnifiedRoleManagementPolicyOperationPredicate) (result ListRoleManagementPoliciesCompleteResult, err error) {
	items := make([]stable.UnifiedRoleManagementPolicy, 0)

	resp, err := c.ListRoleManagementPolicies(ctx, options)
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

	result = ListRoleManagementPoliciesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
