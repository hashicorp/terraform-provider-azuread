package serviceprincipal

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

type CreatePasswordSingleSignOnCredentialsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.PasswordSingleSignOnCredentialSet
}

type CreatePasswordSingleSignOnCredentialsOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreatePasswordSingleSignOnCredentialsOperationOptions() CreatePasswordSingleSignOnCredentialsOperationOptions {
	return CreatePasswordSingleSignOnCredentialsOperationOptions{}
}

func (o CreatePasswordSingleSignOnCredentialsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreatePasswordSingleSignOnCredentialsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreatePasswordSingleSignOnCredentialsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreatePasswordSingleSignOnCredentials - Invoke action createPasswordSingleSignOnCredentials. Create single sign-on
// credentials using a password for a user or group.
func (c ServicePrincipalClient) CreatePasswordSingleSignOnCredentials(ctx context.Context, id beta.ServicePrincipalId, input CreatePasswordSingleSignOnCredentialsRequest, options CreatePasswordSingleSignOnCredentialsOperationOptions) (result CreatePasswordSingleSignOnCredentialsOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/createPasswordSingleSignOnCredentials", id.ID()),
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

	var model beta.PasswordSingleSignOnCredentialSet
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
