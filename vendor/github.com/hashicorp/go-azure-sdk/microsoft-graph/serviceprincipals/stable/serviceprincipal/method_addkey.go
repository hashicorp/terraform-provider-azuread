package serviceprincipal

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

type AddKeyOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.KeyCredential
}

type AddKeyOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultAddKeyOperationOptions() AddKeyOperationOptions {
	return AddKeyOperationOptions{}
}

func (o AddKeyOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AddKeyOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o AddKeyOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// AddKey - Invoke action addKey. Adds a key credential to a servicePrincipal. This method along with removeKey can be
// used by a servicePrincipal to automate rolling its expiring keys. As part of the request validation for this method,
// a proof of possession of an existing key is verified before the action can be performed. ServicePrincipals that
// don’t have any existing valid certificates (i.e.: no certificates have been added yet, or all certificates have
// expired), won’t be able to use this service action. Update servicePrincipal can be used to perform an update
// instead.
func (c ServicePrincipalClient) AddKey(ctx context.Context, id stable.ServicePrincipalId, input AddKeyRequest, options AddKeyOperationOptions) (result AddKeyOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/addKey", id.ID()),
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

	var model stable.KeyCredential
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
