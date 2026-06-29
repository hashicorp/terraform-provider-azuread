package authenticationstrengthpolicycombinationconfiguration

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

type ListAuthenticationStrengthPolicyCombinationConfigurationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.AuthenticationCombinationConfiguration
}

type ListAuthenticationStrengthPolicyCombinationConfigurationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.AuthenticationCombinationConfiguration
}

type ListAuthenticationStrengthPolicyCombinationConfigurationsOperationOptions struct {
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

func DefaultListAuthenticationStrengthPolicyCombinationConfigurationsOperationOptions() ListAuthenticationStrengthPolicyCombinationConfigurationsOperationOptions {
	return ListAuthenticationStrengthPolicyCombinationConfigurationsOperationOptions{}
}

func (o ListAuthenticationStrengthPolicyCombinationConfigurationsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListAuthenticationStrengthPolicyCombinationConfigurationsOperationOptions) ToOData() *odata.Query {
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

func (o ListAuthenticationStrengthPolicyCombinationConfigurationsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListAuthenticationStrengthPolicyCombinationConfigurationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAuthenticationStrengthPolicyCombinationConfigurationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAuthenticationStrengthPolicyCombinationConfigurations - Get combinationConfigurations from policies. Settings
// that may be used to require specific types or instances of an authentication method to be used when authenticating
// with a specified combination of authentication methods.
func (c AuthenticationStrengthPolicyCombinationConfigurationClient) ListAuthenticationStrengthPolicyCombinationConfigurations(ctx context.Context, id stable.PolicyAuthenticationStrengthPolicyId, options ListAuthenticationStrengthPolicyCombinationConfigurationsOperationOptions) (result ListAuthenticationStrengthPolicyCombinationConfigurationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListAuthenticationStrengthPolicyCombinationConfigurationsCustomPager{},
		Path:          fmt.Sprintf("%s/combinationConfigurations", id.ID()),
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

	temp := make([]stable.AuthenticationCombinationConfiguration, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := stable.UnmarshalAuthenticationCombinationConfigurationImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for stable.AuthenticationCombinationConfiguration (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListAuthenticationStrengthPolicyCombinationConfigurationsComplete retrieves all the results into a single object
func (c AuthenticationStrengthPolicyCombinationConfigurationClient) ListAuthenticationStrengthPolicyCombinationConfigurationsComplete(ctx context.Context, id stable.PolicyAuthenticationStrengthPolicyId, options ListAuthenticationStrengthPolicyCombinationConfigurationsOperationOptions) (ListAuthenticationStrengthPolicyCombinationConfigurationsCompleteResult, error) {
	return c.ListAuthenticationStrengthPolicyCombinationConfigurationsCompleteMatchingPredicate(ctx, id, options, AuthenticationCombinationConfigurationOperationPredicate{})
}

// ListAuthenticationStrengthPolicyCombinationConfigurationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AuthenticationStrengthPolicyCombinationConfigurationClient) ListAuthenticationStrengthPolicyCombinationConfigurationsCompleteMatchingPredicate(ctx context.Context, id stable.PolicyAuthenticationStrengthPolicyId, options ListAuthenticationStrengthPolicyCombinationConfigurationsOperationOptions, predicate AuthenticationCombinationConfigurationOperationPredicate) (result ListAuthenticationStrengthPolicyCombinationConfigurationsCompleteResult, err error) {
	items := make([]stable.AuthenticationCombinationConfiguration, 0)

	resp, err := c.ListAuthenticationStrengthPolicyCombinationConfigurations(ctx, id, options)
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

	result = ListAuthenticationStrengthPolicyCombinationConfigurationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
