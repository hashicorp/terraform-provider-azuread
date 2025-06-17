package me

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CheckMemberObjectsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]string
}

type CheckMemberObjectsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []string
}

type CheckMemberObjectsOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultCheckMemberObjectsOperationOptions() CheckMemberObjectsOperationOptions {
	return CheckMemberObjectsOperationOptions{}
}

func (o CheckMemberObjectsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CheckMemberObjectsOperationOptions) ToOData() *odata.Query {
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

func (o CheckMemberObjectsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type CheckMemberObjectsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *CheckMemberObjectsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// CheckMemberObjects - Invoke action checkMemberObjects. Check for membership in a list of group IDs, administrative
// unit IDs, or directory role IDs, for the IDs of the specified - user - group - service principal - organizational
// contact - device - directory object. This method is transitive.
func (c MeClient) CheckMemberObjects(ctx context.Context, input CheckMemberObjectsRequest, options CheckMemberObjectsOperationOptions) (result CheckMemberObjectsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &CheckMemberObjectsCustomPager{},
		Path:          "/me/checkMemberObjects",
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
		Values *[]string `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// CheckMemberObjectsComplete retrieves all the results into a single object
func (c MeClient) CheckMemberObjectsComplete(ctx context.Context, input CheckMemberObjectsRequest, options CheckMemberObjectsOperationOptions) (result CheckMemberObjectsCompleteResult, err error) {
	items := make([]string, 0)

	resp, err := c.CheckMemberObjects(ctx, input, options)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			items = append(items, v)
		}
	}

	result = CheckMemberObjectsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
