package entitlementmanagementaccesspackage

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MoveEntitlementManagementAccessPackageToCatalogOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type MoveEntitlementManagementAccessPackageToCatalogOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultMoveEntitlementManagementAccessPackageToCatalogOperationOptions() MoveEntitlementManagementAccessPackageToCatalogOperationOptions {
	return MoveEntitlementManagementAccessPackageToCatalogOperationOptions{}
}

func (o MoveEntitlementManagementAccessPackageToCatalogOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o MoveEntitlementManagementAccessPackageToCatalogOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o MoveEntitlementManagementAccessPackageToCatalogOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// MoveEntitlementManagementAccessPackageToCatalog - Invoke action moveToCatalog. In Microsoft Entra entitlement
// management, this action moves the accessPackage to a specified target accessPackageCatalog. The resources in the
// access package must be present in the target catalog.
func (c EntitlementManagementAccessPackageClient) MoveEntitlementManagementAccessPackageToCatalog(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageId, input MoveEntitlementManagementAccessPackageToCatalogRequest, options MoveEntitlementManagementAccessPackageToCatalogOperationOptions) (result MoveEntitlementManagementAccessPackageToCatalogOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/moveToCatalog", id.ID()),
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
