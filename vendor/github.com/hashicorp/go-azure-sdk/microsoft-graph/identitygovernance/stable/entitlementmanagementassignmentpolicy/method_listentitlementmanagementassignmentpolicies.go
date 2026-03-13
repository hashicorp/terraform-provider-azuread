package entitlementmanagementassignmentpolicy

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

type ListEntitlementManagementAssignmentPoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AccessPackageAssignmentPolicy
}

type ListEntitlementManagementAssignmentPoliciesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AccessPackageAssignmentPolicy
}

type ListEntitlementManagementAssignmentPoliciesOperationOptions struct {
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

func DefaultListEntitlementManagementAssignmentPoliciesOperationOptions() ListEntitlementManagementAssignmentPoliciesOperationOptions {
	return ListEntitlementManagementAssignmentPoliciesOperationOptions{}
}

func (o ListEntitlementManagementAssignmentPoliciesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementAssignmentPoliciesOperationOptions) ToOData() *odata.Query {
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

func (o ListEntitlementManagementAssignmentPoliciesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEntitlementManagementAssignmentPoliciesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementAssignmentPoliciesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementAssignmentPolicies - List assignmentPolicies. Retrieve a list of
// accessPackageAssignmentPolicy objects in Microsoft Entra entitlement management. If the delegated user is in a
// directory role, the resulting list includes all the assignment policies that the caller has access to read, across
// all catalogs and access packages. If the delegated user is an access package manager or catalog owner, they should
// instead retrieve the policies for the access packages they can read with list accessPackages by including
// $expand=assignmentPolicies as a query parameter.
func (c EntitlementManagementAssignmentPolicyClient) ListEntitlementManagementAssignmentPolicies(ctx context.Context, options ListEntitlementManagementAssignmentPoliciesOperationOptions) (result ListEntitlementManagementAssignmentPoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementAssignmentPoliciesCustomPager{},
		Path:          "/identityGovernance/entitlementManagement/assignmentPolicies",
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
		Values *[]stable.AccessPackageAssignmentPolicy `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementAssignmentPoliciesComplete retrieves all the results into a single object
func (c EntitlementManagementAssignmentPolicyClient) ListEntitlementManagementAssignmentPoliciesComplete(ctx context.Context, options ListEntitlementManagementAssignmentPoliciesOperationOptions) (ListEntitlementManagementAssignmentPoliciesCompleteResult, error) {
	return c.ListEntitlementManagementAssignmentPoliciesCompleteMatchingPredicate(ctx, options, AccessPackageAssignmentPolicyOperationPredicate{})
}

// ListEntitlementManagementAssignmentPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementAssignmentPolicyClient) ListEntitlementManagementAssignmentPoliciesCompleteMatchingPredicate(ctx context.Context, options ListEntitlementManagementAssignmentPoliciesOperationOptions, predicate AccessPackageAssignmentPolicyOperationPredicate) (result ListEntitlementManagementAssignmentPoliciesCompleteResult, err error) {
	items := make([]stable.AccessPackageAssignmentPolicy, 0)

	resp, err := c.ListEntitlementManagementAssignmentPolicies(ctx, options)
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

	result = ListEntitlementManagementAssignmentPoliciesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
