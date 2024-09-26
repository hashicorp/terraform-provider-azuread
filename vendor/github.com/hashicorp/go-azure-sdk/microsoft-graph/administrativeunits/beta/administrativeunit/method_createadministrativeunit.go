package administrativeunit

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateAdministrativeUnitOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.AdministrativeUnit
}

type CreateAdministrativeUnitOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateAdministrativeUnitOperationOptions() CreateAdministrativeUnitOperationOptions {
	return CreateAdministrativeUnitOperationOptions{}
}

func (o CreateAdministrativeUnitOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateAdministrativeUnitOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateAdministrativeUnitOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateAdministrativeUnit - Create administrativeUnit. Use this API to create a new administrativeUnit.
func (c AdministrativeUnitClient) CreateAdministrativeUnit(ctx context.Context, input beta.AdministrativeUnit, options CreateAdministrativeUnitOperationOptions) (result CreateAdministrativeUnitOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/administrativeUnits",
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

	var model beta.AdministrativeUnit
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
