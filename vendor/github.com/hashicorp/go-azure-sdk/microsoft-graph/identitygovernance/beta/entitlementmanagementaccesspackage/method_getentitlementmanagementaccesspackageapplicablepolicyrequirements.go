package entitlementmanagementaccesspackage

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

type GetEntitlementManagementAccessPackageApplicablePolicyRequirementsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AccessPackageAssignmentRequestRequirements
}

type GetEntitlementManagementAccessPackageApplicablePolicyRequirementsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AccessPackageAssignmentRequestRequirements
}

type GetEntitlementManagementAccessPackageApplicablePolicyRequirementsOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultGetEntitlementManagementAccessPackageApplicablePolicyRequirementsOperationOptions() GetEntitlementManagementAccessPackageApplicablePolicyRequirementsOperationOptions {
	return GetEntitlementManagementAccessPackageApplicablePolicyRequirementsOperationOptions{}
}

func (o GetEntitlementManagementAccessPackageApplicablePolicyRequirementsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetEntitlementManagementAccessPackageApplicablePolicyRequirementsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o GetEntitlementManagementAccessPackageApplicablePolicyRequirementsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type GetEntitlementManagementAccessPackageApplicablePolicyRequirementsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *GetEntitlementManagementAccessPackageApplicablePolicyRequirementsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GetEntitlementManagementAccessPackageApplicablePolicyRequirements - Invoke action getApplicablePolicyRequirements. In
// Microsoft Entra entitlement management, this action retrieves a list of accessPackageAssignmentRequestRequirements
// objects that the currently signed-in user can use to create an accessPackageAssignmentRequest. Each requirement
// object corresponds to an access package assignment policy that the currently signed-in user is allowed to request an
// assignment for.
func (c EntitlementManagementAccessPackageClient) GetEntitlementManagementAccessPackageApplicablePolicyRequirements(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageId, options GetEntitlementManagementAccessPackageApplicablePolicyRequirementsOperationOptions) (result GetEntitlementManagementAccessPackageApplicablePolicyRequirementsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &GetEntitlementManagementAccessPackageApplicablePolicyRequirementsCustomPager{},
		Path:          fmt.Sprintf("%s/getApplicablePolicyRequirements", id.ID()),
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
		Values *[]beta.AccessPackageAssignmentRequestRequirements `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GetEntitlementManagementAccessPackageApplicablePolicyRequirementsComplete retrieves all the results into a single object
func (c EntitlementManagementAccessPackageClient) GetEntitlementManagementAccessPackageApplicablePolicyRequirementsComplete(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageId, options GetEntitlementManagementAccessPackageApplicablePolicyRequirementsOperationOptions) (GetEntitlementManagementAccessPackageApplicablePolicyRequirementsCompleteResult, error) {
	return c.GetEntitlementManagementAccessPackageApplicablePolicyRequirementsCompleteMatchingPredicate(ctx, id, options, AccessPackageAssignmentRequestRequirementsOperationPredicate{})
}

// GetEntitlementManagementAccessPackageApplicablePolicyRequirementsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementAccessPackageClient) GetEntitlementManagementAccessPackageApplicablePolicyRequirementsCompleteMatchingPredicate(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageId, options GetEntitlementManagementAccessPackageApplicablePolicyRequirementsOperationOptions, predicate AccessPackageAssignmentRequestRequirementsOperationPredicate) (result GetEntitlementManagementAccessPackageApplicablePolicyRequirementsCompleteResult, err error) {
	items := make([]beta.AccessPackageAssignmentRequestRequirements, 0)

	resp, err := c.GetEntitlementManagementAccessPackageApplicablePolicyRequirements(ctx, id, options)
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

	result = GetEntitlementManagementAccessPackageApplicablePolicyRequirementsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
