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

type ListGetsAvailableExtensionPropertiesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ExtensionProperty
}

type ListGetsAvailableExtensionPropertiesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ExtensionProperty
}

type ListGetsAvailableExtensionPropertiesOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultListGetsAvailableExtensionPropertiesOperationOptions() ListGetsAvailableExtensionPropertiesOperationOptions {
	return ListGetsAvailableExtensionPropertiesOperationOptions{}
}

func (o ListGetsAvailableExtensionPropertiesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListGetsAvailableExtensionPropertiesOperationOptions) ToOData() *odata.Query {
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

func (o ListGetsAvailableExtensionPropertiesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListGetsAvailableExtensionPropertiesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListGetsAvailableExtensionPropertiesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListGetsAvailableExtensionProperties - Invoke action getAvailableExtensionProperties. Return all directory extension
// definitions that have been registered in a directory, including through multi-tenant apps. The following entities
// support extension properties
func (c DirectoryRoleTemplateClient) ListGetsAvailableExtensionProperties(ctx context.Context, input ListGetsAvailableExtensionPropertiesRequest, options ListGetsAvailableExtensionPropertiesOperationOptions) (result ListGetsAvailableExtensionPropertiesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &ListGetsAvailableExtensionPropertiesCustomPager{},
		Path:          "/directoryRoleTemplates/getAvailableExtensionProperties",
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
		Values *[]stable.ExtensionProperty `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGetsAvailableExtensionPropertiesComplete retrieves all the results into a single object
func (c DirectoryRoleTemplateClient) ListGetsAvailableExtensionPropertiesComplete(ctx context.Context, input ListGetsAvailableExtensionPropertiesRequest, options ListGetsAvailableExtensionPropertiesOperationOptions) (ListGetsAvailableExtensionPropertiesCompleteResult, error) {
	return c.ListGetsAvailableExtensionPropertiesCompleteMatchingPredicate(ctx, input, options, ExtensionPropertyOperationPredicate{})
}

// ListGetsAvailableExtensionPropertiesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DirectoryRoleTemplateClient) ListGetsAvailableExtensionPropertiesCompleteMatchingPredicate(ctx context.Context, input ListGetsAvailableExtensionPropertiesRequest, options ListGetsAvailableExtensionPropertiesOperationOptions, predicate ExtensionPropertyOperationPredicate) (result ListGetsAvailableExtensionPropertiesCompleteResult, err error) {
	items := make([]stable.ExtensionProperty, 0)

	resp, err := c.ListGetsAvailableExtensionProperties(ctx, input, options)
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

	result = ListGetsAvailableExtensionPropertiesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
