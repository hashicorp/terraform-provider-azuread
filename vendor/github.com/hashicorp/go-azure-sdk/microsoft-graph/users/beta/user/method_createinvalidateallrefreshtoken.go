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

type CreateInvalidateAllRefreshTokenOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *CreateInvalidateAllRefreshTokenResult
}

type CreateInvalidateAllRefreshTokenOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateInvalidateAllRefreshTokenOperationOptions() CreateInvalidateAllRefreshTokenOperationOptions {
	return CreateInvalidateAllRefreshTokenOperationOptions{}
}

func (o CreateInvalidateAllRefreshTokenOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateInvalidateAllRefreshTokenOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateInvalidateAllRefreshTokenOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateInvalidateAllRefreshToken - Invoke action invalidateAllRefreshTokens. Invalidates all of the user's refresh
// tokens issued to applications (as well as session cookies in a user's browser), by resetting the
// refreshTokensValidFromDateTime user property to the current date-time. Typically, this operation is performed (by the
// user or an administrator) if the user has a lost or stolen device. This operation would prevent access to any of the
// organization's data accessed through applications on the device without the user first being required to sign in
// again. In fact, this operation would force the user to sign in again for all applications that they have previously
// consented to, independent of device. For developers, if the application attempts to redeem a delegated access token
// for this user by using an invalidated refresh token, the application will get an error. If this happens, the
// application will need to acquire a new refresh token by making a request to the authorize endpoint, which will force
// the user to sign in.
func (c UserClient) CreateInvalidateAllRefreshToken(ctx context.Context, id beta.UserId, options CreateInvalidateAllRefreshTokenOperationOptions) (result CreateInvalidateAllRefreshTokenOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/invalidateAllRefreshTokens", id.ID()),
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

	var model CreateInvalidateAllRefreshTokenResult
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
