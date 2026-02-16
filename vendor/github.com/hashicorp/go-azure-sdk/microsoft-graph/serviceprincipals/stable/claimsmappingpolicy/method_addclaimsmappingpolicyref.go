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

type AddClaimsMappingPolicyRefOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type AddClaimsMappingPolicyRefOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultAddClaimsMappingPolicyRefOperationOptions() AddClaimsMappingPolicyRefOperationOptions {
	return AddClaimsMappingPolicyRefOperationOptions{}
}

func (o AddClaimsMappingPolicyRefOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AddClaimsMappingPolicyRefOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o AddClaimsMappingPolicyRefOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// AddClaimsMappingPolicyRef - Assign claimsMappingPolicy. Assign a claimsMappingPolicy to a servicePrincipal.
func (c ClaimsMappingPolicyClient) AddClaimsMappingPolicyRef(ctx context.Context, id stable.ServicePrincipalId, input stable.ReferenceCreate, options AddClaimsMappingPolicyRefOperationOptions) (result AddClaimsMappingPolicyRefOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/claimsMappingPolicies/$ref", id.ID()),
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
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
