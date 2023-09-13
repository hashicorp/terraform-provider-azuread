// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serviceprincipals

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-provider-azuread/internal/sdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"
)

type ClientConfigId struct {
	TenantId string
	ClientId string
	ObjectId string
}

func (id ClientConfigId) ID() string {
	return fmt.Sprintf("%s-%s-%s", id.TenantId, id.ClientId, id.ObjectId)
}

func (ClientConfigId) String() string {
	return "Client Config"
}

type ClientConfigDataSourceModel struct {
	ClientId string `tfschema:"client_id"`
	TenantId string `tfschema:"tenant_id"`
	ObjectId string `tfschema:"object_id"`
}

type ClientConfigDataSource struct{}

var _ sdk.DataSource = ClientConfigDataSource{}

func (r ClientConfigDataSource) ResourceType() string {
	return "azurerm_aadb2c_directory"
}

func (r ClientConfigDataSource) ModelObject() interface{} {
	return &ClientConfigDataSourceModel{}
}

func (r ClientConfigDataSource) Arguments() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{}
}

func (r ClientConfigDataSource) Attributes() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"client_id": {
			Description: "The client ID (application ID) linked to the authenticated principal, or the application used for delegated authentication",
			Type:        pluginsdk.TypeString,
			Computed:    true,
		},

		"tenant_id": {
			Description: "The tenant ID of the authenticated principal",
			Type:        pluginsdk.TypeString,
			Computed:    true,
		},

		"object_id": {
			Description: "The object ID of the authenticated principal",
			Type:        pluginsdk.TypeString,
			Computed:    true,
		},
	}
}

func (r ClientConfigDataSource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			state := ClientConfigDataSourceModel{
				TenantId: metadata.Client.TenantID,
				ClientId: metadata.Client.ClientID,
				ObjectId: metadata.Client.ObjectID,
			}

			metadata.SetID(ClientConfigId{
				TenantId: metadata.Client.TenantID,
				ClientId: metadata.Client.ClientID,
				ObjectId: metadata.Client.ObjectID,
			})

			return metadata.Encode(&state)
		},
	}
}
