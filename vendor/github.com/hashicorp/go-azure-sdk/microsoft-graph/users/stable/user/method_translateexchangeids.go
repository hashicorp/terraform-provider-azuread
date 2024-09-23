package user

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

type TranslateExchangeIdsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ConvertIdResult
}

type TranslateExchangeIdsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ConvertIdResult
}

type TranslateExchangeIdsOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultTranslateExchangeIdsOperationOptions() TranslateExchangeIdsOperationOptions {
	return TranslateExchangeIdsOperationOptions{}
}

func (o TranslateExchangeIdsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o TranslateExchangeIdsOperationOptions) ToOData() *odata.Query {
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

func (o TranslateExchangeIdsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type TranslateExchangeIdsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *TranslateExchangeIdsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// TranslateExchangeIds - Invoke action translateExchangeIds. Translate identifiers of Outlook-related resources between
// formats.
func (c UserClient) TranslateExchangeIds(ctx context.Context, id stable.UserId, input TranslateExchangeIdsRequest, options TranslateExchangeIdsOperationOptions) (result TranslateExchangeIdsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &TranslateExchangeIdsCustomPager{},
		Path:          fmt.Sprintf("%s/translateExchangeIds", id.ID()),
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
		Values *[]stable.ConvertIdResult `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// TranslateExchangeIdsComplete retrieves all the results into a single object
func (c UserClient) TranslateExchangeIdsComplete(ctx context.Context, id stable.UserId, input TranslateExchangeIdsRequest, options TranslateExchangeIdsOperationOptions) (TranslateExchangeIdsCompleteResult, error) {
	return c.TranslateExchangeIdsCompleteMatchingPredicate(ctx, id, input, options, ConvertIdResultOperationPredicate{})
}

// TranslateExchangeIdsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserClient) TranslateExchangeIdsCompleteMatchingPredicate(ctx context.Context, id stable.UserId, input TranslateExchangeIdsRequest, options TranslateExchangeIdsOperationOptions, predicate ConvertIdResultOperationPredicate) (result TranslateExchangeIdsCompleteResult, err error) {
	items := make([]stable.ConvertIdResult, 0)

	resp, err := c.TranslateExchangeIds(ctx, id, input, options)
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

	result = TranslateExchangeIdsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
