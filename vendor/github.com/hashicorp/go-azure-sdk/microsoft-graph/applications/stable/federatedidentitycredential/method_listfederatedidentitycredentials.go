package federatedidentitycredential

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

type ListFederatedIdentityCredentialsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.FederatedIdentityCredential
}

type ListFederatedIdentityCredentialsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.FederatedIdentityCredential
}

type ListFederatedIdentityCredentialsOperationOptions struct {
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

func DefaultListFederatedIdentityCredentialsOperationOptions() ListFederatedIdentityCredentialsOperationOptions {
	return ListFederatedIdentityCredentialsOperationOptions{}
}

func (o ListFederatedIdentityCredentialsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListFederatedIdentityCredentialsOperationOptions) ToOData() *odata.Query {
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

func (o ListFederatedIdentityCredentialsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListFederatedIdentityCredentialsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListFederatedIdentityCredentialsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListFederatedIdentityCredentials - List federatedIdentityCredentials. Get a list of the federatedIdentityCredential
// objects and their properties.
func (c FederatedIdentityCredentialClient) ListFederatedIdentityCredentials(ctx context.Context, id stable.ApplicationId, options ListFederatedIdentityCredentialsOperationOptions) (result ListFederatedIdentityCredentialsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListFederatedIdentityCredentialsCustomPager{},
		Path:          fmt.Sprintf("%s/federatedIdentityCredentials", id.ID()),
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
		Values *[]stable.FederatedIdentityCredential `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListFederatedIdentityCredentialsComplete retrieves all the results into a single object
func (c FederatedIdentityCredentialClient) ListFederatedIdentityCredentialsComplete(ctx context.Context, id stable.ApplicationId, options ListFederatedIdentityCredentialsOperationOptions) (ListFederatedIdentityCredentialsCompleteResult, error) {
	return c.ListFederatedIdentityCredentialsCompleteMatchingPredicate(ctx, id, options, FederatedIdentityCredentialOperationPredicate{})
}

// ListFederatedIdentityCredentialsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c FederatedIdentityCredentialClient) ListFederatedIdentityCredentialsCompleteMatchingPredicate(ctx context.Context, id stable.ApplicationId, options ListFederatedIdentityCredentialsOperationOptions, predicate FederatedIdentityCredentialOperationPredicate) (result ListFederatedIdentityCredentialsCompleteResult, err error) {
	items := make([]stable.FederatedIdentityCredential, 0)

	resp, err := c.ListFederatedIdentityCredentials(ctx, id, options)
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

	result = ListFederatedIdentityCredentialsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
