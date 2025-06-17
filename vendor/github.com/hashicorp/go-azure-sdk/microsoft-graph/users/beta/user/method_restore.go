package user

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestoreOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.DirectoryObject
}

type RestoreOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultRestoreOperationOptions() RestoreOperationOptions {
	return RestoreOperationOptions{}
}

func (o RestoreOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RestoreOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o RestoreOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// Restore - Invoke action restore. Restore a recently deleted directory object from deleted items. The following types
// are supported: - administrativeUnit - application - certificateBasedAuthPki - certificateAuthorityDetail -
// externalUserProfile - group - pendingExternalUserProfile - servicePrincipal - user If an item was accidentally
// deleted, you can fully restore the item. This isn't applicable to security groups, which are deleted permanently.
// Also, restoring an application doesn't restore the associated service principal automatically. You must call this API
// to explicitly restore the deleted service principal. A recently deleted item remains available for up to 30 days.
// After 30 days, the item is permanently deleted.
func (c UserClient) Restore(ctx context.Context, id beta.UserId, options RestoreOperationOptions) (result RestoreOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/restore", id.ID()),
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := beta.UnmarshalDirectoryObjectImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
