package serviceprincipal

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

type CheckMemberGroupsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]string
}

type CheckMemberGroupsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []string
}

type CheckMemberGroupsOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Skip      *int64
	Top       *int64
}

func DefaultCheckMemberGroupsOperationOptions() CheckMemberGroupsOperationOptions {
	return CheckMemberGroupsOperationOptions{}
}

func (o CheckMemberGroupsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CheckMemberGroupsOperationOptions) ToOData() *odata.Query {
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

func (o CheckMemberGroupsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type CheckMemberGroupsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *CheckMemberGroupsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// CheckMemberGroups - Invoke action checkMemberGroups. Check for membership in a specified list of group IDs, and
// return from that list the IDs of groups where a specified object is a member. The specified object can be of one of
// the following types: - user - group - service principal - organizational contact - device - directory object This
// function is transitive. You can check up to a maximum of 20 groups per request. This function supports all groups
// provisioned in Microsoft Entra ID. Because Microsoft 365 groups cannot contain other groups, membership in a
// Microsoft 365 group is always direct.
func (c ServicePrincipalClient) CheckMemberGroups(ctx context.Context, id beta.ServicePrincipalId, input CheckMemberGroupsRequest, options CheckMemberGroupsOperationOptions) (result CheckMemberGroupsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &CheckMemberGroupsCustomPager{},
		Path:          fmt.Sprintf("%s/checkMemberGroups", id.ID()),
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

// CheckMemberGroupsComplete retrieves all the results into a single object
func (c ServicePrincipalClient) CheckMemberGroupsComplete(ctx context.Context, id beta.ServicePrincipalId, input CheckMemberGroupsRequest, options CheckMemberGroupsOperationOptions) (result CheckMemberGroupsCompleteResult, err error) {
	items := make([]string, 0)

	resp, err := c.CheckMemberGroups(ctx, id, input, options)
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

	result = CheckMemberGroupsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
