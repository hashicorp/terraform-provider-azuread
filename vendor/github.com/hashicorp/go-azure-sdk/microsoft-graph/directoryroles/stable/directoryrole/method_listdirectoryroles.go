package directoryrole

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListDirectoryRolesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DirectoryRole
}

type ListDirectoryRolesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DirectoryRole
}

type ListDirectoryRolesOperationOptions struct {
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

func DefaultListDirectoryRolesOperationOptions() ListDirectoryRolesOperationOptions {
	return ListDirectoryRolesOperationOptions{}
}

func (o ListDirectoryRolesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDirectoryRolesOperationOptions) ToOData() *odata.Query {
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

func (o ListDirectoryRolesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDirectoryRolesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDirectoryRolesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDirectoryRoles - List directoryRoles. List the directory roles that are activated in the tenant. This operation
// only returns roles that have been activated. A role becomes activated when an admin activates the role using the
// Activate directoryRole API. Not all built-in roles are initially activated. When assigning a role using the Microsoft
// Entra admin center, the role activation step is implicitly done on the admin's behalf. To get the full list of roles
// that are available in Microsoft Entra ID, use List directoryRoleTemplates.
func (c DirectoryRoleClient) ListDirectoryRoles(ctx context.Context, options ListDirectoryRolesOperationOptions) (result ListDirectoryRolesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDirectoryRolesCustomPager{},
		Path:          "/directoryRoles",
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
		Values *[]stable.DirectoryRole `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDirectoryRolesComplete retrieves all the results into a single object
func (c DirectoryRoleClient) ListDirectoryRolesComplete(ctx context.Context, options ListDirectoryRolesOperationOptions) (ListDirectoryRolesCompleteResult, error) {
	return c.ListDirectoryRolesCompleteMatchingPredicate(ctx, options, DirectoryRoleOperationPredicate{})
}

// ListDirectoryRolesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DirectoryRoleClient) ListDirectoryRolesCompleteMatchingPredicate(ctx context.Context, options ListDirectoryRolesOperationOptions, predicate DirectoryRoleOperationPredicate) (result ListDirectoryRolesCompleteResult, err error) {
	items := make([]stable.DirectoryRole, 0)

	resp, err := c.ListDirectoryRoles(ctx, options)
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

	result = ListDirectoryRolesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
