package group

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

type CheckGrantedPermissionsForAppsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ResourceSpecificPermissionGrant
}

type CheckGrantedPermissionsForAppsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ResourceSpecificPermissionGrant
}

type CheckGrantedPermissionsForAppsOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultCheckGrantedPermissionsForAppsOperationOptions() CheckGrantedPermissionsForAppsOperationOptions {
	return CheckGrantedPermissionsForAppsOperationOptions{}
}

func (o CheckGrantedPermissionsForAppsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CheckGrantedPermissionsForAppsOperationOptions) ToOData() *odata.Query {
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

func (o CheckGrantedPermissionsForAppsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type CheckGrantedPermissionsForAppsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *CheckGrantedPermissionsForAppsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// CheckGrantedPermissionsForApps - Invoke action checkGrantedPermissionsForApp
func (c GroupClient) CheckGrantedPermissionsForApps(ctx context.Context, id beta.GroupId, options CheckGrantedPermissionsForAppsOperationOptions) (result CheckGrantedPermissionsForAppsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &CheckGrantedPermissionsForAppsCustomPager{},
		Path:          fmt.Sprintf("%s/checkGrantedPermissionsForApp", id.ID()),
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
		Values *[]beta.ResourceSpecificPermissionGrant `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// CheckGrantedPermissionsForAppsComplete retrieves all the results into a single object
func (c GroupClient) CheckGrantedPermissionsForAppsComplete(ctx context.Context, id beta.GroupId, options CheckGrantedPermissionsForAppsOperationOptions) (CheckGrantedPermissionsForAppsCompleteResult, error) {
	return c.CheckGrantedPermissionsForAppsCompleteMatchingPredicate(ctx, id, options, ResourceSpecificPermissionGrantOperationPredicate{})
}

// CheckGrantedPermissionsForAppsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupClient) CheckGrantedPermissionsForAppsCompleteMatchingPredicate(ctx context.Context, id beta.GroupId, options CheckGrantedPermissionsForAppsOperationOptions, predicate ResourceSpecificPermissionGrantOperationPredicate) (result CheckGrantedPermissionsForAppsCompleteResult, err error) {
	items := make([]beta.ResourceSpecificPermissionGrant, 0)

	resp, err := c.CheckGrantedPermissionsForApps(ctx, id, options)
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

	result = CheckGrantedPermissionsForAppsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
