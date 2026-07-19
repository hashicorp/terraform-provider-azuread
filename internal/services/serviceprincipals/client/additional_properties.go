// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	serviceprincipalBeta "github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/serviceprincipal"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
)

type ServicePrincipalAdditionalProperties struct {
	IsDisabled      *bool   `json:"isDisabled,omitempty"`
	SamlMetadataUrl *string `json:"samlMetadataUrl,omitempty"`
}

func (c *Client) GetServicePrincipalAdditionalProperties(ctx context.Context, id beta.ServicePrincipalId) (*ServicePrincipalAdditionalProperties, error) {
	options := serviceprincipalBeta.GetServicePrincipalOperationOptions{
		Select: &[]string{"isDisabled", "samlMetadataUrl"},
	}
	opts := client.RequestOptions{
		ContentType:         "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{http.StatusOK},
		HttpMethod:          http.MethodGet,
		OptionsObject:       options,
		Path:                id.ID(),
	}

	req, err := c.ServicePrincipalClientBeta.Client.NewRequest(ctx, opts)
	if err != nil {
		return nil, err
	}

	resp, err := req.Execute(ctx)
	if err != nil {
		return nil, err
	}

	model := &ServicePrincipalAdditionalProperties{}
	if err := resp.Unmarshal(model); err != nil {
		return nil, err
	}

	return model, nil
}

func (c *Client) UpdateServicePrincipalDisabled(ctx context.Context, id beta.ServicePrincipalId, disabled bool) error {
	options := serviceprincipalBeta.DefaultUpdateServicePrincipalOperationOptions()
	opts := client.RequestOptions{
		ContentType:         "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{http.StatusAccepted, http.StatusNoContent, http.StatusOK},
		HttpMethod:          http.MethodPatch,
		OptionsObject:       options,
		Path:                id.ID(),
	}

	req, err := c.ServicePrincipalClientBeta.Client.NewRequest(ctx, opts)
	if err != nil {
		return err
	}
	if err := req.Marshal(ServicePrincipalAdditionalProperties{IsDisabled: &disabled}); err != nil {
		return err
	}

	_, err = req.Execute(ctx)
	return err
}
