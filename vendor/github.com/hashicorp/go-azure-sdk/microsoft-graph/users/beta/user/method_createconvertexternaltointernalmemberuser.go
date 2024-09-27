package user

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateConvertExternalToInternalMemberUserOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.ConversionUserDetails
}

type CreateConvertExternalToInternalMemberUserOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateConvertExternalToInternalMemberUserOperationOptions() CreateConvertExternalToInternalMemberUserOperationOptions {
	return CreateConvertExternalToInternalMemberUserOperationOptions{}
}

func (o CreateConvertExternalToInternalMemberUserOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateConvertExternalToInternalMemberUserOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateConvertExternalToInternalMemberUserOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateConvertExternalToInternalMemberUser - Invoke action convertExternalToInternalMemberUser. Convert an externally
// authenticated user into an internal user. The user is able to sign into the host tenant as an internal user and
// access resources as a member. For more information about this conversion, see Convert external users to internal
// users.
func (c UserClient) CreateConvertExternalToInternalMemberUser(ctx context.Context, id beta.UserId, input CreateConvertExternalToInternalMemberUserRequest, options CreateConvertExternalToInternalMemberUserOperationOptions) (result CreateConvertExternalToInternalMemberUserOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/convertExternalToInternalMemberUser", id.ID()),
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

	var model beta.ConversionUserDetails
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
