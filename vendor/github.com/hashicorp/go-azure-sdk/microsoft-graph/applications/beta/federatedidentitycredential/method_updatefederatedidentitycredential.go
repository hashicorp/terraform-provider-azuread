package federatedidentitycredential

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateFederatedIdentityCredentialOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateFederatedIdentityCredentialOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateFederatedIdentityCredentialOperationOptions() UpdateFederatedIdentityCredentialOperationOptions {
	return UpdateFederatedIdentityCredentialOperationOptions{}
}

func (o UpdateFederatedIdentityCredentialOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateFederatedIdentityCredentialOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateFederatedIdentityCredentialOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateFederatedIdentityCredential - Upsert federatedIdentityCredential. Create a new federatedIdentityCredential
// object for an application if it doesn't exist, or update the properties of an existing federatedIdentityCredential
// object. By configuring a trust relationship between your Microsoft Entra application registration and the identity
// provider for your compute platform, you can use tokens issued by that platform to authenticate with Microsoft
// identity platform and call APIs in the Microsoft ecosystem. Maximum of 20 objects can be added to an application.
func (c FederatedIdentityCredentialClient) UpdateFederatedIdentityCredential(ctx context.Context, id beta.ApplicationIdFederatedIdentityCredentialId, input beta.FederatedIdentityCredential, options UpdateFederatedIdentityCredentialOperationOptions) (result UpdateFederatedIdentityCredentialOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
