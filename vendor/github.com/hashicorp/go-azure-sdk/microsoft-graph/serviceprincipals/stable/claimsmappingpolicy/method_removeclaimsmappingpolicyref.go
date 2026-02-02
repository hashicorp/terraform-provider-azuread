package claimsmappingpolicy

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoveClaimsMappingPolicyRefOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type RemoveClaimsMappingPolicyRefOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultRemoveClaimsMappingPolicyRefOperationOptions() RemoveClaimsMappingPolicyRefOperationOptions {
	return RemoveClaimsMappingPolicyRefOperationOptions{}
}

func (o RemoveClaimsMappingPolicyRefOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o RemoveClaimsMappingPolicyRefOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o RemoveClaimsMappingPolicyRefOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// RemoveClaimsMappingPolicyRef - Remove claimsMappingPolicy. Remove a claimsMappingPolicy from a servicePrincipal.
func (c ClaimsMappingPolicyClient) RemoveClaimsMappingPolicyRef(ctx context.Context, id stable.ServicePrincipalIdClaimsMappingPolicyId, options RemoveClaimsMappingPolicyRefOperationOptions) (result RemoveClaimsMappingPolicyRefOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/$ref", id.ID()),
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	return
}
