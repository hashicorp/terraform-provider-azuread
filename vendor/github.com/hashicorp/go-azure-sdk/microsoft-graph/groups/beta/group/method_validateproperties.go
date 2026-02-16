package group

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

type ValidatePropertiesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type ValidatePropertiesOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultValidatePropertiesOperationOptions() ValidatePropertiesOperationOptions {
	return ValidatePropertiesOperationOptions{}
}

func (o ValidatePropertiesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ValidatePropertiesOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o ValidatePropertiesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// ValidateProperties - Invoke action validateProperties. Validate if a Microsoft 365 group's display name or mail
// nickname complies with naming policies. Clients can use the API to determine if a display name or mail nickname is
// valid before trying to update a Microsoft 365 group. For validating properties before creating a group, use the
// validateProperties function for directory objects. The following validations are performed for the display name and
// mail nickname properties: This API returns with the first failure encountered. If one or more properties fail
// multiple validations, only the property with the first validation failure is returned. However, you can validate both
// the mail nickname and the display name and receive a collection of validation errors if you are only validating the
// prefix and suffix naming policy.
func (c GroupClient) ValidateProperties(ctx context.Context, id beta.GroupId, input ValidatePropertiesRequest, options ValidatePropertiesOperationOptions) (result ValidatePropertiesOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/validateProperties", id.ID()),
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
