package entitlementmanagementaccesspackagecatalogaccesspackageresource

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

type ListEntitlementManagementAccessPackageCatalogResourcesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AccessPackageResource
}

type ListEntitlementManagementAccessPackageCatalogResourcesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AccessPackageResource
}

type ListEntitlementManagementAccessPackageCatalogResourcesOperationOptions struct {
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

func DefaultListEntitlementManagementAccessPackageCatalogResourcesOperationOptions() ListEntitlementManagementAccessPackageCatalogResourcesOperationOptions {
	return ListEntitlementManagementAccessPackageCatalogResourcesOperationOptions{}
}

func (o ListEntitlementManagementAccessPackageCatalogResourcesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementAccessPackageCatalogResourcesOperationOptions) ToOData() *odata.Query {
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

func (o ListEntitlementManagementAccessPackageCatalogResourcesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEntitlementManagementAccessPackageCatalogResourcesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementAccessPackageCatalogResourcesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementAccessPackageCatalogResources - List accessPackageResources. Retrieve a list of
// accessPackageResource objects in an accessPackageCatalog. To request to add or remove an accessPackageResource, use
// create accessPackageResourceRequest.
func (c EntitlementManagementAccessPackageCatalogAccessPackageResourceClient) ListEntitlementManagementAccessPackageCatalogResources(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageCatalogId, options ListEntitlementManagementAccessPackageCatalogResourcesOperationOptions) (result ListEntitlementManagementAccessPackageCatalogResourcesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementAccessPackageCatalogResourcesCustomPager{},
		Path:          fmt.Sprintf("%s/accessPackageResources", id.ID()),
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
		Values *[]beta.AccessPackageResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementAccessPackageCatalogResourcesComplete retrieves all the results into a single object
func (c EntitlementManagementAccessPackageCatalogAccessPackageResourceClient) ListEntitlementManagementAccessPackageCatalogResourcesComplete(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageCatalogId, options ListEntitlementManagementAccessPackageCatalogResourcesOperationOptions) (ListEntitlementManagementAccessPackageCatalogResourcesCompleteResult, error) {
	return c.ListEntitlementManagementAccessPackageCatalogResourcesCompleteMatchingPredicate(ctx, id, options, AccessPackageResourceOperationPredicate{})
}

// ListEntitlementManagementAccessPackageCatalogResourcesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementAccessPackageCatalogAccessPackageResourceClient) ListEntitlementManagementAccessPackageCatalogResourcesCompleteMatchingPredicate(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageCatalogId, options ListEntitlementManagementAccessPackageCatalogResourcesOperationOptions, predicate AccessPackageResourceOperationPredicate) (result ListEntitlementManagementAccessPackageCatalogResourcesCompleteResult, err error) {
	items := make([]beta.AccessPackageResource, 0)

	resp, err := c.ListEntitlementManagementAccessPackageCatalogResources(ctx, id, options)
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

	result = ListEntitlementManagementAccessPackageCatalogResourcesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
