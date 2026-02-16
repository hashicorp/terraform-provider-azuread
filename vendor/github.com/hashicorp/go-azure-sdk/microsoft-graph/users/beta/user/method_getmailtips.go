package user

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

type GetMailTipsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.MailTips
}

type GetMailTipsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.MailTips
}

type GetMailTipsOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultGetMailTipsOperationOptions() GetMailTipsOperationOptions {
	return GetMailTipsOperationOptions{}
}

func (o GetMailTipsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetMailTipsOperationOptions) ToOData() *odata.Query {
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

func (o GetMailTipsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type GetMailTipsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *GetMailTipsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GetMailTips - Invoke action getMailTips. Get the MailTips of one or more recipients as available to the signed-in
// user. Note that by making a POST call to the getMailTips action, you can request specific types of MailTips to be
// returned for more than one recipient at one time. The requested MailTips are returned in a mailTips collection.
func (c UserClient) GetMailTips(ctx context.Context, id beta.UserId, input GetMailTipsRequest, options GetMailTipsOperationOptions) (result GetMailTipsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &GetMailTipsCustomPager{},
		Path:          fmt.Sprintf("%s/getMailTips", id.ID()),
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
		Values *[]beta.MailTips `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GetMailTipsComplete retrieves all the results into a single object
func (c UserClient) GetMailTipsComplete(ctx context.Context, id beta.UserId, input GetMailTipsRequest, options GetMailTipsOperationOptions) (GetMailTipsCompleteResult, error) {
	return c.GetMailTipsCompleteMatchingPredicate(ctx, id, input, options, MailTipsOperationPredicate{})
}

// GetMailTipsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserClient) GetMailTipsCompleteMatchingPredicate(ctx context.Context, id beta.UserId, input GetMailTipsRequest, options GetMailTipsOperationOptions, predicate MailTipsOperationPredicate) (result GetMailTipsCompleteResult, err error) {
	items := make([]beta.MailTips, 0)

	resp, err := c.GetMailTips(ctx, id, input, options)
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

	result = GetMailTipsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
