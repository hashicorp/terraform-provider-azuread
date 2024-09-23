package directoryroletemplate

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

type ListDirectoryRoleTemplatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DirectoryRoleTemplate
}

type ListDirectoryRoleTemplatesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DirectoryRoleTemplate
}

type ListDirectoryRoleTemplatesOperationOptions struct {
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

func DefaultListDirectoryRoleTemplatesOperationOptions() ListDirectoryRoleTemplatesOperationOptions {
	return ListDirectoryRoleTemplatesOperationOptions{}
}

func (o ListDirectoryRoleTemplatesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDirectoryRoleTemplatesOperationOptions) ToOData() *odata.Query {
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

func (o ListDirectoryRoleTemplatesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDirectoryRoleTemplatesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDirectoryRoleTemplatesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDirectoryRoleTemplates - List directoryRoleTemplates. Retrieve a list of directoryRoleTemplate objects.
func (c DirectoryRoleTemplateClient) ListDirectoryRoleTemplates(ctx context.Context, options ListDirectoryRoleTemplatesOperationOptions) (result ListDirectoryRoleTemplatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDirectoryRoleTemplatesCustomPager{},
		Path:          "/directoryRoleTemplates",
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
		Values *[]stable.DirectoryRoleTemplate `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDirectoryRoleTemplatesComplete retrieves all the results into a single object
func (c DirectoryRoleTemplateClient) ListDirectoryRoleTemplatesComplete(ctx context.Context, options ListDirectoryRoleTemplatesOperationOptions) (ListDirectoryRoleTemplatesCompleteResult, error) {
	return c.ListDirectoryRoleTemplatesCompleteMatchingPredicate(ctx, options, DirectoryRoleTemplateOperationPredicate{})
}

// ListDirectoryRoleTemplatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DirectoryRoleTemplateClient) ListDirectoryRoleTemplatesCompleteMatchingPredicate(ctx context.Context, options ListDirectoryRoleTemplatesOperationOptions, predicate DirectoryRoleTemplateOperationPredicate) (result ListDirectoryRoleTemplatesCompleteResult, err error) {
	items := make([]stable.DirectoryRoleTemplate, 0)

	resp, err := c.ListDirectoryRoleTemplates(ctx, options)
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

	result = ListDirectoryRoleTemplatesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
