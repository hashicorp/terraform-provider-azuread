package conditionalaccessnamedlocation

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetConditionalAccessNamedLocationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        stable.NamedLocation
}

type GetConditionalAccessNamedLocationOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetConditionalAccessNamedLocationOperationOptions() GetConditionalAccessNamedLocationOperationOptions {
	return GetConditionalAccessNamedLocationOperationOptions{}
}

func (o GetConditionalAccessNamedLocationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetConditionalAccessNamedLocationOperationOptions) ToOData() *odata.Query {
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

func (o GetConditionalAccessNamedLocationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetConditionalAccessNamedLocation - Get ipNamedLocation. Retrieve the properties and relationships of an
// ipNamedLocation object.
func (c ConditionalAccessNamedLocationClient) GetConditionalAccessNamedLocation(ctx context.Context, id stable.IdentityConditionalAccessNamedLocationId, options GetConditionalAccessNamedLocationOperationOptions) (result GetConditionalAccessNamedLocationOperationResponse, err error) {
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
	model, err := stable.UnmarshalNamedLocationImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
