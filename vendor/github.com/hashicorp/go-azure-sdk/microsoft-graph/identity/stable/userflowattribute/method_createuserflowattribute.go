package userflowattribute

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

type CreateUserFlowAttributeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        stable.IdentityUserFlowAttribute
}

type CreateUserFlowAttributeOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateUserFlowAttributeOperationOptions() CreateUserFlowAttributeOperationOptions {
	return CreateUserFlowAttributeOperationOptions{}
}

func (o CreateUserFlowAttributeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateUserFlowAttributeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateUserFlowAttributeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateUserFlowAttribute - Create identityUserFlowAttribute. Create a new custom identityUserFlowAttribute object.
func (c UserFlowAttributeClient) CreateUserFlowAttribute(ctx context.Context, input stable.IdentityUserFlowAttribute, options CreateUserFlowAttributeOperationOptions) (result CreateUserFlowAttributeOperationResponse, err error) {
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
		Path:          "/identity/userFlowAttributes",
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
