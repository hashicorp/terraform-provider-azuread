package federatedidentitycredential

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

type CreateFederatedIdentityCredentialOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.FederatedIdentityCredential
}

type CreateFederatedIdentityCredentialOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateFederatedIdentityCredentialOperationOptions() CreateFederatedIdentityCredentialOperationOptions {
	return CreateFederatedIdentityCredentialOperationOptions{}
}

func (o CreateFederatedIdentityCredentialOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateFederatedIdentityCredentialOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateFederatedIdentityCredentialOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateFederatedIdentityCredential - Create federatedIdentityCredential. Create a new federatedIdentityCredential
// object for an application. By configuring a trust relationship between your Microsoft Entra application registration
// and the identity provider for your compute platform, you can use tokens issued by that platform to authenticate with
// Microsoft identity platform and call APIs in the Microsoft ecosystem. Maximum of 20 objects can be added to an
// application.
func (c FederatedIdentityCredentialClient) CreateFederatedIdentityCredential(ctx context.Context, id beta.ApplicationId, input beta.FederatedIdentityCredential, options CreateFederatedIdentityCredentialOperationOptions) (result CreateFederatedIdentityCredentialOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/federatedIdentityCredentials", id.ID()),
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

	var model beta.FederatedIdentityCredential
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
