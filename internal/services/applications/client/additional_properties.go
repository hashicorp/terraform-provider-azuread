// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"context"
	"net/http"

	applicationBeta "github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/application"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
)

type ApplicationAdditionalProperties struct {
	IsDisabled                *bool `json:"isDisabled,omitempty"`
	OAuth2RequirePostResponse *bool `json:"oauth2RequirePostResponse,omitempty"`
}

func (c *Client) GetApplicationAdditionalProperties(ctx context.Context, id beta.ApplicationId) (*ApplicationAdditionalProperties, error) {
	options := applicationBeta.GetApplicationOperationOptions{
		Select: &[]string{"isDisabled", "oauth2RequirePostResponse"},
	}
	opts := client.RequestOptions{
		ContentType:         "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{http.StatusOK},
		HttpMethod:          http.MethodGet,
		OptionsObject:       options,
		Path:                id.ID(),
	}

	req, err := c.ApplicationClientBeta.Client.NewRequest(ctx, opts)
	if err != nil {
		return nil, err
	}

	resp, err := req.Execute(ctx)
	if err != nil {
		return nil, err
	}

	model := &ApplicationAdditionalProperties{}
	if err := resp.Unmarshal(model); err != nil {
		return nil, err
	}

	return model, nil
}

func (c *Client) UpdateApplicationDisabled(ctx context.Context, id beta.ApplicationId, disabled bool) error {
	options := applicationBeta.DefaultUpdateApplicationOperationOptions()
	opts := client.RequestOptions{
		ContentType:         "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{http.StatusAccepted, http.StatusNoContent, http.StatusOK},
		HttpMethod:          http.MethodPatch,
		OptionsObject:       options,
		Path:                id.ID(),
	}

	req, err := c.ApplicationClientBeta.Client.NewRequest(ctx, opts)
	if err != nil {
		return err
	}
	if err := req.Marshal(ApplicationAdditionalProperties{IsDisabled: &disabled}); err != nil {
		return err
	}

	_, err = req.Execute(ctx)
	return err
}
