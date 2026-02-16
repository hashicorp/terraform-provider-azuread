package directoryroledefinition

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

type ListDirectoryRoleDefinitionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.UnifiedRoleDefinition
}

type ListDirectoryRoleDefinitionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.UnifiedRoleDefinition
}

type ListDirectoryRoleDefinitionsOperationOptions struct {
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

func DefaultListDirectoryRoleDefinitionsOperationOptions() ListDirectoryRoleDefinitionsOperationOptions {
	return ListDirectoryRoleDefinitionsOperationOptions{}
}

func (o ListDirectoryRoleDefinitionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDirectoryRoleDefinitionsOperationOptions) ToOData() *odata.Query {
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

func (o ListDirectoryRoleDefinitionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDirectoryRoleDefinitionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDirectoryRoleDefinitionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDirectoryRoleDefinitions - List roleDefinitions. Get a list of unifiedRoleDefinition objects for the provider.
// The following RBAC providers are currently supported: - directory (Microsoft Entra ID) - entitlement management
// (Microsoft Entra Entitlement Management)
func (c DirectoryRoleDefinitionClient) ListDirectoryRoleDefinitions(ctx context.Context, options ListDirectoryRoleDefinitionsOperationOptions) (result ListDirectoryRoleDefinitionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDirectoryRoleDefinitionsCustomPager{},
		Path:          "/roleManagement/directory/roleDefinitions",
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
		Values *[]stable.UnifiedRoleDefinition `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDirectoryRoleDefinitionsComplete retrieves all the results into a single object
func (c DirectoryRoleDefinitionClient) ListDirectoryRoleDefinitionsComplete(ctx context.Context, options ListDirectoryRoleDefinitionsOperationOptions) (ListDirectoryRoleDefinitionsCompleteResult, error) {
	return c.ListDirectoryRoleDefinitionsCompleteMatchingPredicate(ctx, options, UnifiedRoleDefinitionOperationPredicate{})
}

// ListDirectoryRoleDefinitionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DirectoryRoleDefinitionClient) ListDirectoryRoleDefinitionsCompleteMatchingPredicate(ctx context.Context, options ListDirectoryRoleDefinitionsOperationOptions, predicate UnifiedRoleDefinitionOperationPredicate) (result ListDirectoryRoleDefinitionsCompleteResult, err error) {
	items := make([]stable.UnifiedRoleDefinition, 0)

	resp, err := c.ListDirectoryRoleDefinitions(ctx, options)
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

	result = ListDirectoryRoleDefinitionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
