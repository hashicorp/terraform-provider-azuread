package entitlementmanagementroledefinition

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

type ListEntitlementManagementRoleDefinitionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UnifiedRoleDefinition
}

type ListEntitlementManagementRoleDefinitionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UnifiedRoleDefinition
}

type ListEntitlementManagementRoleDefinitionsOperationOptions struct {
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

func DefaultListEntitlementManagementRoleDefinitionsOperationOptions() ListEntitlementManagementRoleDefinitionsOperationOptions {
	return ListEntitlementManagementRoleDefinitionsOperationOptions{}
}

func (o ListEntitlementManagementRoleDefinitionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementRoleDefinitionsOperationOptions) ToOData() *odata.Query {
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

func (o ListEntitlementManagementRoleDefinitionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEntitlementManagementRoleDefinitionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementRoleDefinitionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementRoleDefinitions - List roleDefinitions. Get a list of unifiedRoleDefinition objects for an
// RBAC provider. The following RBAC providers are currently supported: - Cloud PC - device management (Intune) -
// directory (Microsoft Entra ID) - entitlement management (Microsoft Entra ID) - Exchange Online
func (c EntitlementManagementRoleDefinitionClient) ListEntitlementManagementRoleDefinitions(ctx context.Context, options ListEntitlementManagementRoleDefinitionsOperationOptions) (result ListEntitlementManagementRoleDefinitionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementRoleDefinitionsCustomPager{},
		Path:          "/roleManagement/entitlementManagement/roleDefinitions",
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
		Values *[]beta.UnifiedRoleDefinition `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementRoleDefinitionsComplete retrieves all the results into a single object
func (c EntitlementManagementRoleDefinitionClient) ListEntitlementManagementRoleDefinitionsComplete(ctx context.Context, options ListEntitlementManagementRoleDefinitionsOperationOptions) (ListEntitlementManagementRoleDefinitionsCompleteResult, error) {
	return c.ListEntitlementManagementRoleDefinitionsCompleteMatchingPredicate(ctx, options, UnifiedRoleDefinitionOperationPredicate{})
}

// ListEntitlementManagementRoleDefinitionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementRoleDefinitionClient) ListEntitlementManagementRoleDefinitionsCompleteMatchingPredicate(ctx context.Context, options ListEntitlementManagementRoleDefinitionsOperationOptions, predicate UnifiedRoleDefinitionOperationPredicate) (result ListEntitlementManagementRoleDefinitionsCompleteResult, err error) {
	items := make([]beta.UnifiedRoleDefinition, 0)

	resp, err := c.ListEntitlementManagementRoleDefinitions(ctx, options)
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

	result = ListEntitlementManagementRoleDefinitionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
