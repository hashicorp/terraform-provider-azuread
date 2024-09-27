package administrativeunitscopedrolemember

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateAdministrativeUnitScopedRoleMemberOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.ScopedRoleMembership
}

type CreateAdministrativeUnitScopedRoleMemberOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateAdministrativeUnitScopedRoleMemberOperationOptions() CreateAdministrativeUnitScopedRoleMemberOperationOptions {
	return CreateAdministrativeUnitScopedRoleMemberOperationOptions{}
}

func (o CreateAdministrativeUnitScopedRoleMemberOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateAdministrativeUnitScopedRoleMemberOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateAdministrativeUnitScopedRoleMemberOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateAdministrativeUnitScopedRoleMember - Add a scopedRoleMember. Assign a Microsoft Entra role with administrative
// unit scope. For a list of roles that can be assigned with administrative unit scope, see Assign Microsoft Entra roles
// with administrative unit scope.
func (c AdministrativeUnitScopedRoleMemberClient) CreateAdministrativeUnitScopedRoleMember(ctx context.Context, id stable.DirectoryAdministrativeUnitId, input stable.ScopedRoleMembership, options CreateAdministrativeUnitScopedRoleMemberOperationOptions) (result CreateAdministrativeUnitScopedRoleMemberOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/scopedRoleMembers", id.ID()),
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

	var model stable.ScopedRoleMembership
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
