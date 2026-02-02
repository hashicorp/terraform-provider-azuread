package userflowattribute

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetUserFlowAttributeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        stable.IdentityUserFlowAttribute
}

type GetUserFlowAttributeOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetUserFlowAttributeOperationOptions() GetUserFlowAttributeOperationOptions {
	return GetUserFlowAttributeOperationOptions{}
}

func (o GetUserFlowAttributeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetUserFlowAttributeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetUserFlowAttributeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetUserFlowAttribute - Get identityUserFlowAttribute. Retrieve the properties and relationships of a
// identityUserFlowAttribute object.
func (c UserFlowAttributeClient) GetUserFlowAttribute(ctx context.Context, id stable.IdentityUserFlowAttributeId, options GetUserFlowAttributeOperationOptions) (result GetUserFlowAttributeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := stable.UnmarshalIdentityUserFlowAttributeImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
