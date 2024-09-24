package entitlementmanagementaccesspackageassignmentpolicy

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

type ListEntitlementManagementAccessPackageAssignmentPoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AccessPackageAssignmentPolicy
}

type ListEntitlementManagementAccessPackageAssignmentPoliciesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AccessPackageAssignmentPolicy
}

type ListEntitlementManagementAccessPackageAssignmentPoliciesOperationOptions struct {
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

func DefaultListEntitlementManagementAccessPackageAssignmentPoliciesOperationOptions() ListEntitlementManagementAccessPackageAssignmentPoliciesOperationOptions {
	return ListEntitlementManagementAccessPackageAssignmentPoliciesOperationOptions{}
}

func (o ListEntitlementManagementAccessPackageAssignmentPoliciesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementAccessPackageAssignmentPoliciesOperationOptions) ToOData() *odata.Query {
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

func (o ListEntitlementManagementAccessPackageAssignmentPoliciesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEntitlementManagementAccessPackageAssignmentPoliciesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementAccessPackageAssignmentPoliciesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementAccessPackageAssignmentPolicies - List accessPackageAssignmentPolicies. In Microsoft Entra
// entitlement management, retrieve a list of accessPackageAssignmentPolicy objects. If the delegated user is in a
// directory role, the resulting list includes all the assignment policies that the caller has access to read, across
// all catalogs and access packages. If the delegated user is an access package manager or catalog owner, they should
// instead retrieve the policies for the access packages they can read with list accessPackages by including
// $expand=accessPackageAssignmentPolicies in the query.
func (c EntitlementManagementAccessPackageAssignmentPolicyClient) ListEntitlementManagementAccessPackageAssignmentPolicies(ctx context.Context, options ListEntitlementManagementAccessPackageAssignmentPoliciesOperationOptions) (result ListEntitlementManagementAccessPackageAssignmentPoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementAccessPackageAssignmentPoliciesCustomPager{},
		Path:          "/identityGovernance/entitlementManagement/accessPackageAssignmentPolicies",
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
		Values *[]beta.AccessPackageAssignmentPolicy `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementAccessPackageAssignmentPoliciesComplete retrieves all the results into a single object
func (c EntitlementManagementAccessPackageAssignmentPolicyClient) ListEntitlementManagementAccessPackageAssignmentPoliciesComplete(ctx context.Context, options ListEntitlementManagementAccessPackageAssignmentPoliciesOperationOptions) (ListEntitlementManagementAccessPackageAssignmentPoliciesCompleteResult, error) {
	return c.ListEntitlementManagementAccessPackageAssignmentPoliciesCompleteMatchingPredicate(ctx, options, AccessPackageAssignmentPolicyOperationPredicate{})
}

// ListEntitlementManagementAccessPackageAssignmentPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementAccessPackageAssignmentPolicyClient) ListEntitlementManagementAccessPackageAssignmentPoliciesCompleteMatchingPredicate(ctx context.Context, options ListEntitlementManagementAccessPackageAssignmentPoliciesOperationOptions, predicate AccessPackageAssignmentPolicyOperationPredicate) (result ListEntitlementManagementAccessPackageAssignmentPoliciesCompleteResult, err error) {
	items := make([]beta.AccessPackageAssignmentPolicy, 0)

	resp, err := c.ListEntitlementManagementAccessPackageAssignmentPolicies(ctx, options)
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

	result = ListEntitlementManagementAccessPackageAssignmentPoliciesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
