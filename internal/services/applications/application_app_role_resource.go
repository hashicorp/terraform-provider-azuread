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
	applicationsValidate "github.com/hashicorp/terraform-provider-azuread/internal/services/applications/validate"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
	"github.com/manicminer/hamilton/msgraph"
)

type ApplicationAppRoleModel struct {
	ApplicationId      string   `tfschema:"application_id"`
	RoleId             string   `tfschema:"role_id"`
	AllowedMemberTypes []string `tfschema:"allowed_member_types"`
	Description        string   `tfschema:"description"`
	DisplayName        string   `tfschema:"display_name"`
	Value              string   `tfschema:"value"`
}

type ApplicationAppRoleResource struct{}

func (r ApplicationAppRoleResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return parse.ValidateAppRoleID
}

var _ sdk.ResourceWithUpdate = ApplicationAppRoleResource{}

func (r ApplicationAppRoleResource) ResourceType() string {
	return "azuread_application_app_role"
}

func (r ApplicationAppRoleResource) ModelObject() interface{} {
	return &ApplicationAppRoleModel{}
}

func (r ApplicationAppRoleResource) Arguments() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"application_id": {
			Description:  "The resource ID of the application to which this app role should be applied",
			Type:         pluginsdk.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: parse.ValidateApplicationID,
		},

		"role_id": {
			Description:  "The unique identifier of the app role",
			Type:         pluginsdk.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.IsUUID,
		},

		"allowed_member_types": {
			Description: "Specifies whether this app role definition can be assigned to users and groups by setting to `User`, or to other applications (that are accessing this application in a standalone scenario) by setting to `Application`, or to both",
			Type:        pluginsdk.TypeSet,
			Required:    true,
			MinItems:    1,
			Elem: &pluginsdk.Schema{
				Type: pluginsdk.TypeString,
				ValidateFunc: validation.StringInSlice(
					[]string{
						msgraph.AppRoleAllowedMemberTypeApplication,
						msgraph.AppRoleAllowedMemberTypeUser,
					}, false,
				),
			},
		},

		"description": {
			Description:  "Description of the app role that appears when the role is being assigned and, if the role functions as an application permissions, during the consent experiences",
			Type:         pluginsdk.TypeString,
			Required:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},

		"display_name": {
			Description:  "Display name for the app role that appears during app role assignment and in consent experiences",
			Type:         pluginsdk.TypeString,
			Required:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},

		"value": {
			Description:      "The value that is used for the `roles` claim in ID tokens and OAuth access tokens that are authenticating an assigned service or user principal",
			Type:             pluginsdk.TypeString,
			Optional:         true,
			ValidateDiagFunc: applicationsValidate.RoleScopeClaimValue,
		},
	}
}

func (r ApplicationAppRoleResource) Attributes() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{}
}

func (r ApplicationAppRoleResource) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 10 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Applications.ApplicationsClient
			client.BaseClient.DisableRetries = true
			defer func() { client.BaseClient.DisableRetries = false }()

			var model ApplicationAppRoleModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			applicationId, err := parse.ParseApplicationID(model.ApplicationId)
			if err != nil {
				return err
			}

			id := parse.NewAppRoleID(applicationId.ApplicationId, model.RoleId)

			tf.LockByName(applicationResourceName, id.ApplicationId)
			defer tf.UnlockByName(applicationResourceName, id.ApplicationId)

			result, _, err := client.Get(ctx, applicationId.ApplicationId, odata.Query{})
			if err != nil {
				return fmt.Errorf("retrieving %s: %+v", applicationId, err)
			}
			if result == nil {
				return fmt.Errorf("retrieving %s: result was nil", applicationId)
			}

			newRoles := make([]msgraph.AppRole, 0)

			// Don't forget any existing roles, since all roles must be updated together
			if result.AppRoles != nil {
				newRoles = *result.AppRoles
			}

			// Check for existing role ID
			for _, role := range newRoles {
				if strings.EqualFold(*role.ID, id.RoleID) {
					return metadata.ResourceRequiresImport(r.ResourceType(), id)
				}
			}

			newRoles = append(newRoles, msgraph.AppRole{
				ID:                 &model.RoleId,
				IsEnabled:          pointer.To(true),
				AllowedMemberTypes: &model.AllowedMemberTypes,
				Description:        &model.Description,
				DisplayName:        &model.DisplayName,
				Value:              &model.Value,
			})

			properties := msgraph.Application{
				DirectoryObject: msgraph.DirectoryObject{
					Id: &id.ApplicationId,
				},
				AppRoles: &newRoles,
			}

			if _, err = client.Update(ctx, properties); err != nil {
				return fmt.Errorf("creating %s: %+v", id, err)
			}

			metadata.SetID(id)
			return nil
		},
	}
}

func (r ApplicationAppRoleResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Applications.ApplicationsClient
			client.BaseClient.DisableRetries = true
			defer func() { client.BaseClient.DisableRetries = false }()

			id, err := parse.ParseAppRoleID(metadata.ResourceData.Id())
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
			if result.AppRoles == nil {
				return metadata.MarkAsGone(id)
			}

			// Identify the role by ID
			var role *msgraph.AppRole
			for _, existingRole := range *result.AppRoles {
				if strings.EqualFold(*existingRole.ID, id.RoleID) {
					role = &existingRole
					break
				}
			}

			if role == nil {
				return metadata.MarkAsGone(id)
			}

			state := ApplicationAppRoleModel{
				ApplicationId:      applicationId.ID(),
				RoleId:             id.RoleID,
				AllowedMemberTypes: pointer.From(role.AllowedMemberTypes),
				Description:        pointer.From(role.Description),
				DisplayName:        pointer.From(role.DisplayName),
				Value:              pointer.From(role.Value),
			}

			return metadata.Encode(&state)
		},
	}
}

func (r ApplicationAppRoleResource) Update() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 10 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Applications.ApplicationsClient

			id, err := parse.ParseAppRoleID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			var model ApplicationAppRoleModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			// Prepare a new role to replace the existing one
			role := msgraph.AppRole{
				ID:                 &id.RoleID,
				IsEnabled:          pointer.To(true),
				AllowedMemberTypes: &model.AllowedMemberTypes,
				Description:        &model.Description,
				DisplayName:        &model.DisplayName,
				Value:              &model.Value,
			}

			tf.LockByName(applicationResourceName, id.ApplicationId)
			defer tf.UnlockByName(applicationResourceName, id.ApplicationId)

			applicationId := parse.NewApplicationID(id.ApplicationId)
			result, _, err := client.Get(ctx, applicationId.ApplicationId, odata.Query{})
			if err != nil {
				return fmt.Errorf("retrieving %s: %+v", applicationId, err)
			}
			if result == nil || result.AppRoles == nil {
				return fmt.Errorf("retrieving %s: appRoles was nil", applicationId)
			}

			// Look for a role to replace, matching by ID
			newRoles := make([]msgraph.AppRole, 0)
			found := false
			for _, existingRole := range *result.AppRoles {
				if strings.EqualFold(*existingRole.ID, id.RoleID) {
					newRoles = append(newRoles, role)
					found = true
				} else {
					newRoles = append(newRoles, existingRole)
				}
			}
			if !found {
				return fmt.Errorf("updating %s: could not identify existing app role", id)
			}

			// Disable the existing role prior to update
			if err = applicationDisableAppRoles(ctx, client, result, &newRoles); err != nil {
				return fmt.Errorf("disabling %s in preparation for update: %+v", id, err)
			}

			properties := msgraph.Application{
				DirectoryObject: msgraph.DirectoryObject{
					Id: &applicationId.ApplicationId,
				},
				AppRoles: &newRoles,
			}

			// Patch the application with the new set of roles
			_, err = client.Update(ctx, properties)
			if err != nil {
				return fmt.Errorf("updating %s: %+v", id, err)
			}

			return nil
		},
	}
}

func (r ApplicationAppRoleResource) Delete() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Applications.ApplicationsClient
			client.BaseClient.DisableRetries = true
			defer func() { client.BaseClient.DisableRetries = false }()

			id, err := parse.ParseAppRoleID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			var model ApplicationAppRoleModel
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
			if result == nil || result.AppRoles == nil {
				return fmt.Errorf("retrieving %s: api.oauth2AppRoles was nil", applicationId)
			}

			// Look for a role to remove, matching by ID
			newRoles := make([]msgraph.AppRole, 0)
			found := false
			for _, existingRole := range *result.AppRoles {
				if strings.EqualFold(*existingRole.ID, id.RoleID) {
					found = true
				} else {
					newRoles = append(newRoles, existingRole)
				}
			}
			if !found {
				return fmt.Errorf("deleting %s: could not identify existing app role", id)
			}

			// Disable the existing role prior to update
			if err = applicationDisableAppRoles(ctx, client, result, &newRoles); err != nil {
				return fmt.Errorf("disabling %s in preparation for deletion: %+v", id, err)
			}

			properties := msgraph.Application{
				DirectoryObject: msgraph.DirectoryObject{
					Id: &applicationId.ApplicationId,
				},
				AppRoles: &newRoles,
			}

			// Patch the application with the new set of roles
			_, err = client.Update(ctx, properties)
			if err != nil {
				return fmt.Errorf("deleting %s: %+v", id, err)
			}

			return nil
		},
	}
}
