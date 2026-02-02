package userflowattribute

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListUserFlowAttributesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.IdentityUserFlowAttribute
}

type ListUserFlowAttributesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.IdentityUserFlowAttribute
}

type ListUserFlowAttributesOperationOptions struct {
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

func DefaultListUserFlowAttributesOperationOptions() ListUserFlowAttributesOperationOptions {
	return ListUserFlowAttributesOperationOptions{}
}

func (o ListUserFlowAttributesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListUserFlowAttributesOperationOptions) ToOData() *odata.Query {
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

func (o ListUserFlowAttributesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListUserFlowAttributesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserFlowAttributesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserFlowAttributes - List identityUserFlowAttributes. Retrieve a list of identityUserFlowAttribute objects.
func (c UserFlowAttributeClient) ListUserFlowAttributes(ctx context.Context, options ListUserFlowAttributesOperationOptions) (result ListUserFlowAttributesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListUserFlowAttributesCustomPager{},
		Path:          "/identity/userFlowAttributes",
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
		Values *[]json.RawMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	temp := make([]stable.IdentityUserFlowAttribute, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := stable.UnmarshalIdentityUserFlowAttributeImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for stable.IdentityUserFlowAttribute (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListUserFlowAttributesComplete retrieves all the results into a single object
func (c UserFlowAttributeClient) ListUserFlowAttributesComplete(ctx context.Context, options ListUserFlowAttributesOperationOptions) (ListUserFlowAttributesCompleteResult, error) {
	return c.ListUserFlowAttributesCompleteMatchingPredicate(ctx, options, IdentityUserFlowAttributeOperationPredicate{})
}

// ListUserFlowAttributesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserFlowAttributeClient) ListUserFlowAttributesCompleteMatchingPredicate(ctx context.Context, options ListUserFlowAttributesOperationOptions, predicate IdentityUserFlowAttributeOperationPredicate) (result ListUserFlowAttributesCompleteResult, err error) {
	items := make([]stable.IdentityUserFlowAttribute, 0)

	resp, err := c.ListUserFlowAttributes(ctx, options)
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

	result = ListUserFlowAttributesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
