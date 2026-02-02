package entitlementmanagementaccesspackagecatalog

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

type ListEntitlementManagementAccessPackageCatalogsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AccessPackageCatalog
}

type ListEntitlementManagementAccessPackageCatalogsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AccessPackageCatalog
}

type ListEntitlementManagementAccessPackageCatalogsOperationOptions struct {
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

func DefaultListEntitlementManagementAccessPackageCatalogsOperationOptions() ListEntitlementManagementAccessPackageCatalogsOperationOptions {
	return ListEntitlementManagementAccessPackageCatalogsOperationOptions{}
}

func (o ListEntitlementManagementAccessPackageCatalogsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementAccessPackageCatalogsOperationOptions) ToOData() *odata.Query {
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

func (o ListEntitlementManagementAccessPackageCatalogsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEntitlementManagementAccessPackageCatalogsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementAccessPackageCatalogsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementAccessPackageCatalogs - List accessPackageCatalogs. Retrieve a list of accessPackageCatalog
// objects.
func (c EntitlementManagementAccessPackageCatalogClient) ListEntitlementManagementAccessPackageCatalogs(ctx context.Context, options ListEntitlementManagementAccessPackageCatalogsOperationOptions) (result ListEntitlementManagementAccessPackageCatalogsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementAccessPackageCatalogsCustomPager{},
		Path:          "/identityGovernance/entitlementManagement/accessPackageCatalogs",
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
		Values *[]beta.AccessPackageCatalog `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementAccessPackageCatalogsComplete retrieves all the results into a single object
func (c EntitlementManagementAccessPackageCatalogClient) ListEntitlementManagementAccessPackageCatalogsComplete(ctx context.Context, options ListEntitlementManagementAccessPackageCatalogsOperationOptions) (ListEntitlementManagementAccessPackageCatalogsCompleteResult, error) {
	return c.ListEntitlementManagementAccessPackageCatalogsCompleteMatchingPredicate(ctx, options, AccessPackageCatalogOperationPredicate{})
}

// ListEntitlementManagementAccessPackageCatalogsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementAccessPackageCatalogClient) ListEntitlementManagementAccessPackageCatalogsCompleteMatchingPredicate(ctx context.Context, options ListEntitlementManagementAccessPackageCatalogsOperationOptions, predicate AccessPackageCatalogOperationPredicate) (result ListEntitlementManagementAccessPackageCatalogsCompleteResult, err error) {
	items := make([]beta.AccessPackageCatalog, 0)

	resp, err := c.ListEntitlementManagementAccessPackageCatalogs(ctx, options)
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

	result = ListEntitlementManagementAccessPackageCatalogsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
