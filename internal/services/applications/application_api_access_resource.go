// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applications

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-provider-azuread/internal/sdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/applications/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
	"github.com/manicminer/hamilton/msgraph"
)

type ApplicationApiAccessModel struct {
	ApplicationId string   `tfschema:"application_id"`
	ApiClientId   string   `tfschema:"api_client_id"`
	RoleIds       []string `tfschema:"role_ids"`
	ScopeIds      []string `tfschema:"scope_ids"`
}

type ApplicationApiAccessResource struct{}

func (r ApplicationApiAccessResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return parse.ValidateApiAccessID
}

var _ sdk.ResourceWithUpdate = ApplicationApiAccessResource{}

func (r ApplicationApiAccessResource) ResourceType() string {
	return "azuread_application_api_access"
}

func (r ApplicationApiAccessResource) ModelObject() interface{} {
	return &ApplicationApiAccessModel{}
}

func (r ApplicationApiAccessResource) Arguments() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"application_id": {
			Description:  "The resource ID of the application to which this API access is granted",
			Type:         pluginsdk.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: parse.ValidateApplicationID,
		},

		"api_client_id": {
			Description:  "The client ID of the API to which access is being granted",
			Type:         pluginsdk.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.IsUUID,
		},

		"role_ids": {
			Description:  "A set of role IDs to be granted to the application, as published by the API",
			Type:         pluginsdk.TypeSet,
			Optional:     true,
			AtLeastOneOf: []string{"role_ids", "scope_ids"},
			Elem: &pluginsdk.Schema{
				Type:         pluginsdk.TypeString,
				ValidateFunc: validation.IsUUID,
			},
		},

		"scope_ids": {
			Description:  "A set of scope IDs to be granted to the application, as published by the API",
			Type:         pluginsdk.TypeSet,
			Optional:     true,
			AtLeastOneOf: []string{"role_ids", "scope_ids"},
			Elem: &pluginsdk.Schema{
				Type:         pluginsdk.TypeString,
				ValidateFunc: validation.IsUUID,
			},
		},
	}
}

func (r ApplicationApiAccessResource) Attributes() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{}
}

func (r ApplicationApiAccessResource) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 10 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Applications.ApplicationsClient
			client.BaseClient.DisableRetries = true
			defer func() { client.BaseClient.DisableRetries = false }()

			var model ApplicationApiAccessModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			applicationId, err := parse.ParseApplicationID(model.ApplicationId)
			if err != nil {
				return err
			}

			id := parse.NewApiAccessID(applicationId.ApplicationId, model.ApiClientId)

			tf.LockByName(applicationResourceName, id.ApplicationId)
			defer tf.UnlockByName(applicationResourceName, id.ApplicationId)

			result, _, err := client.Get(ctx, applicationId.ApplicationId, odata.Query{})
			if err != nil {
				return fmt.Errorf("retrieving %s: %+v", applicationId, err)
			}
			if result == nil {
				return fmt.Errorf("retrieving %s: result was nil", applicationId)
			}

			newApis := make([]msgraph.RequiredResourceAccess, 0)

			// Don't forget any existing APIs, since they must all be updated together
			if result.RequiredResourceAccess != nil {
				newApis = *result.RequiredResourceAccess
			}

			// Check for existing API
			for _, api := range newApis {
				if strings.EqualFold(*api.ResourceAppId, id.ApiClientId) {
					return metadata.ResourceRequiresImport(r.ResourceType(), id)
				}
			}

			permissions := make([]msgraph.ResourceAccess, 0)
			for _, roleId := range model.RoleIds {
				permissions = append(permissions, msgraph.ResourceAccess{
					ID:   pointer.To(roleId),
					Type: msgraph.ResourceAccessTypeRole,
				})
			}
			for _, scopeId := range model.ScopeIds {
				permissions = append(permissions, msgraph.ResourceAccess{
					ID:   pointer.To(scopeId),
					Type: msgraph.ResourceAccessTypeScope,
				})
			}

			newApis = append(newApis, msgraph.RequiredResourceAccess{
				ResourceAppId:  &model.ApiClientId,
				ResourceAccess: &permissions,
			})

			properties := msgraph.Application{
				DirectoryObject: msgraph.DirectoryObject{
					Id: &id.ApplicationId,
				},
				RequiredResourceAccess: &newApis,
			}

			if _, err = client.Update(ctx, properties); err != nil {
				return fmt.Errorf("creating %s: %+v", id, err)
			}

			metadata.SetID(id)
			return nil
		},
	}
}

func (r ApplicationApiAccessResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Applications.ApplicationsClient
			client.BaseClient.DisableRetries = true
			defer func() { client.BaseClient.DisableRetries = false }()

			id, err := parse.ParseApiAccessID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			applicationId := parse.NewApplicationID(id.ApplicationId)

			tf.LockByName(applicationResourceName, id.ApplicationId)
			defer tf.UnlockByName(applicationResourceName, id.ApplicationId)

			result, status, err := client.Get(ctx, id.ApplicationId, odata.Query{})
			if err != nil {
				if status == http.StatusNotFound {
					return metadata.MarkAsGone(id)
				}
				return fmt.Errorf("retrieving %s: %+v", id, err)
			}
			if result == nil {
				return fmt.Errorf("retrieving %s: result was nil", id)
			}
			if result.RequiredResourceAccess == nil {
				return metadata.MarkAsGone(id)
			}

			// Identify the API
			var api *msgraph.RequiredResourceAccess
			for _, existingApi := range *result.RequiredResourceAccess {
				if strings.EqualFold(*existingApi.ResourceAppId, id.ApiClientId) {
					api = &existingApi
					break
				}
			}

			if api == nil {
				return metadata.MarkAsGone(id)
			}
			if api.ResourceAccess == nil {
				return fmt.Errorf("retrieving %s: resourceAccess was nil", id)
			}

			roleIds := make([]string, 0)
			scopeIds := make([]string, 0)
			for _, permission := range *api.ResourceAccess {
				switch permission.Type {
				case msgraph.ResourceAccessTypeRole:
					roleIds = append(roleIds, pointer.From(permission.ID))
				case msgraph.ResourceAccessTypeScope:
					scopeIds = append(scopeIds, pointer.From(permission.ID))
				}
			}

			state := ApplicationApiAccessModel{
				ApplicationId: applicationId.ID(),
				ApiClientId:   pointer.From(api.ResourceAppId),
				RoleIds:       roleIds,
				ScopeIds:      scopeIds,
			}

			return metadata.Encode(&state)
		},
	}
}

func (r ApplicationApiAccessResource) Update() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 10 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Applications.ApplicationsClient

			id, err := parse.ParseApiAccessID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			var model ApplicationApiAccessModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			// Prepare a new API to replace the existing one
			permissions := make([]msgraph.ResourceAccess, 0)
			for _, roleId := range model.RoleIds {
				permissions = append(permissions, msgraph.ResourceAccess{
					ID:   pointer.To(roleId),
					Type: msgraph.ResourceAccessTypeRole,
				})
			}
			for _, scopeId := range model.ScopeIds {
				permissions = append(permissions, msgraph.ResourceAccess{
					ID:   pointer.To(scopeId),
					Type: msgraph.ResourceAccessTypeScope,
				})
			}
			api := msgraph.RequiredResourceAccess{
				ResourceAppId:  &model.ApiClientId,
				ResourceAccess: &permissions,
			}

			tf.LockByName(applicationResourceName, id.ApplicationId)
			defer tf.UnlockByName(applicationResourceName, id.ApplicationId)

			applicationId := parse.NewApplicationID(id.ApplicationId)
			result, _, err := client.Get(ctx, applicationId.ApplicationId, odata.Query{})
			if err != nil {
				return fmt.Errorf("retrieving %s: %+v", applicationId, err)
			}
			if result == nil || result.RequiredResourceAccess == nil {
				return fmt.Errorf("retrieving %s: requiredResourceAccess was nil", applicationId)
			}

			// Look for an API to replace
			newApis := make([]msgraph.RequiredResourceAccess, 0)
			found := false
			for _, existingApi := range *result.RequiredResourceAccess {
				if strings.EqualFold(*existingApi.ResourceAppId, id.ApiClientId) {
					newApis = append(newApis, api)
					found = true
				} else {
					newApis = append(newApis, existingApi)
				}
			}
			if !found {
				return fmt.Errorf("updating %s: could not identify existing API", id)
			}

			properties := msgraph.Application{
				DirectoryObject: msgraph.DirectoryObject{
					Id: &applicationId.ApplicationId,
				},
				RequiredResourceAccess: &newApis,
			}

			// Patch the application with the new set of APIs
			_, err = client.Update(ctx, properties)
			if err != nil {
				return fmt.Errorf("updating %s: %+v", id, err)
			}

			return nil
		},
	}
}

func (r ApplicationApiAccessResource) Delete() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Applications.ApplicationsClient
			client.BaseClient.DisableRetries = true
			defer func() { client.BaseClient.DisableRetries = false }()

			id, err := parse.ParseApiAccessID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			var model ApplicationApiAccessModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			tf.LockByName(applicationResourceName, id.ApplicationId)
			defer tf.UnlockByName(applicationResourceName, id.ApplicationId)

			applicationId := parse.NewApplicationID(id.ApplicationId)
			result, _, err := client.Get(ctx, applicationId.ApplicationId, odata.Query{})
			if err != nil {
				return fmt.Errorf("retrieving %s: %+v", applicationId, err)
			}
			if result == nil || result.RequiredResourceAccess == nil {
				return fmt.Errorf("retrieving %s: requiredResourceAccess was nil", applicationId)
			}

			// Look for an API to remove
			newApis := make([]msgraph.RequiredResourceAccess, 0)
			found := false
			for _, existingApi := range *result.RequiredResourceAccess {
				if strings.EqualFold(*existingApi.ResourceAppId, id.ApiClientId) {
					found = true
				} else {
					newApis = append(newApis, existingApi)
				}
			}
			if !found {
				return fmt.Errorf("deleting %s: could not identify existing API", id)
			}

			properties := msgraph.Application{
				DirectoryObject: msgraph.DirectoryObject{
					Id: &applicationId.ApplicationId,
				},
				RequiredResourceAccess: &newApis,
			}

			// Patch the application with the new set of APIs
			_, err = client.Update(ctx, properties)
			if err != nil {
				return fmt.Errorf("deleting %s: %+v", id, err)
			}

			return nil
		},
	}
}
