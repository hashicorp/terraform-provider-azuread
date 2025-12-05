// Copyright IBM Corp. 2019, 2025
// SPDX-License-Identifier: MPL-2.0

package applications

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/application"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
	"github.com/hashicorp/terraform-provider-azuread/internal/sdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/applications/parse"
	applicationsValidate "github.com/hashicorp/terraform-provider-azuread/internal/services/applications/validate"
)

type ApplicationAppRoleModel struct {
	ApplicationId      string   `tfschema:"application_id"`
	RoleId             string   `tfschema:"role_id"`
	AllowedMemberTypes []string `tfschema:"allowed_member_types"`
	Description        string   `tfschema:"description"`
	DisplayName        string   `tfschema:"display_name"`
	Value              string   `tfschema:"value"`
}

var _ sdk.ResourceWithUpdate = ApplicationAppRoleResource{}

type ApplicationAppRoleResource struct{}

func (r ApplicationAppRoleResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return parse.ValidateAppRoleID
}

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
			ValidateFunc: stable.ValidateApplicationID,
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
				Type:         pluginsdk.TypeString,
				ValidateFunc: validation.StringInSlice(possibleValuesForAppRoleAllowedMemberType, false),
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
			client := metadata.Client.Applications.ApplicationClient

			var model ApplicationAppRoleModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			applicationId, err := stable.ParseApplicationID(model.ApplicationId)
			if err != nil {
				return err
			}

			id := parse.NewAppRoleID(applicationId.ApplicationId, model.RoleId)

			tf.LockByName(applicationResourceName, id.ApplicationId)
			defer tf.UnlockByName(applicationResourceName, id.ApplicationId)

			resp, err := client.GetApplication(ctx, *applicationId, application.DefaultGetApplicationOperationOptions())
			if err != nil {
				return fmt.Errorf("retrieving %s: %+v", applicationId, err)
			}

			app := resp.Model
			if app == nil {
				return fmt.Errorf("retrieving %s: model was nil", applicationId)
			}

			newRoles := make([]stable.AppRole, 0)

			// Don't forget any existing roles, since all roles must be updated together
			if app.AppRoles != nil {
				newRoles = *app.AppRoles
			}

			// Check for existing role ID
			for _, role := range newRoles {
				if strings.EqualFold(*role.Id, id.RoleID) {
					return metadata.ResourceRequiresImport(r.ResourceType(), id)
				}
			}

			newRoles = append(newRoles, stable.AppRole{
				Id:                 &model.RoleId,
				IsEnabled:          pointer.To(true),
				AllowedMemberTypes: &model.AllowedMemberTypes,
				Description:        nullable.Value(model.Description),
				DisplayName:        nullable.Value(model.DisplayName),
				Value:              nullable.Value(model.Value),
			})

			properties := stable.Application{
				Id:       &id.ApplicationId,
				AppRoles: &newRoles,
			}

			if _, err = client.UpdateApplication(ctx, *applicationId, properties, application.DefaultUpdateApplicationOperationOptions()); err != nil {
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
			client := metadata.Client.Applications.ApplicationClient

			id, err := parse.ParseAppRoleID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			applicationId := stable.NewApplicationID(id.ApplicationId)

			tf.LockByName(applicationResourceName, id.ApplicationId)
			defer tf.UnlockByName(applicationResourceName, id.ApplicationId)

			resp, err := client.GetApplication(ctx, applicationId, application.DefaultGetApplicationOperationOptions())
			if err != nil {
				if response.WasNotFound(resp.HttpResponse) {
					return metadata.MarkAsGone(id)
				}
				return fmt.Errorf("retrieving %s: %+v", applicationId, err)
			}

			app := resp.Model
			if app == nil {
				return fmt.Errorf("retrieving %s: model was nil", applicationId)
			}

			if app.AppRoles == nil {
				return metadata.MarkAsGone(id)
			}

			// Identify the role by ID
			var role *stable.AppRole
			for _, existingRole := range *app.AppRoles {
				if strings.EqualFold(*existingRole.Id, id.RoleID) {
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
				Description:        role.Description.GetOrZero(),
				DisplayName:        role.DisplayName.GetOrZero(),
				Value:              role.Value.GetOrZero(),
			}

			return metadata.Encode(&state)
		},
	}
}

func (r ApplicationAppRoleResource) Update() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 10 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Applications.ApplicationClient

			id, err := parse.ParseAppRoleID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			var model ApplicationAppRoleModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			// Prepare a new role to replace the existing one
			role := stable.AppRole{
				Id:                 &id.RoleID,
				IsEnabled:          pointer.To(true),
				AllowedMemberTypes: &model.AllowedMemberTypes,
				Description:        nullable.Value(model.Description),
				DisplayName:        nullable.Value(model.DisplayName),
				Value:              nullable.Value(model.Value),
			}

			applicationId := stable.NewApplicationID(id.ApplicationId)

			tf.LockByName(applicationResourceName, id.ApplicationId)
			defer tf.UnlockByName(applicationResourceName, id.ApplicationId)

			resp, err := client.GetApplication(ctx, applicationId, application.DefaultGetApplicationOperationOptions())
			if err != nil {
				return fmt.Errorf("retrieving %s: %+v", applicationId, err)
			}

			app := resp.Model
			if app == nil || app.AppRoles == nil {
				return fmt.Errorf("retrieving %s: appRoles was nil", applicationId)
			}

			// Look for a role to replace, matching by ID
			newRoles := make([]stable.AppRole, 0)
			found := false
			for _, existingRole := range *app.AppRoles {
				if strings.EqualFold(*existingRole.Id, id.RoleID) {
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
			if err = applicationDisableAppRoles(ctx, client, applicationId, &newRoles); err != nil {
				return fmt.Errorf("disabling %s in preparation for update: %+v", id, err)
			}

			properties := stable.Application{
				Id:       &applicationId.ApplicationId,
				AppRoles: &newRoles,
			}

			// Patch the application with the new set of roles
			if _, err = client.UpdateApplication(ctx, applicationId, properties, application.DefaultUpdateApplicationOperationOptions()); err != nil {
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
			client := metadata.Client.Applications.ApplicationClient

			id, err := parse.ParseAppRoleID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			var model ApplicationAppRoleModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			applicationId := stable.NewApplicationID(id.ApplicationId)

			tf.LockByName(applicationResourceName, id.ApplicationId)
			defer tf.UnlockByName(applicationResourceName, id.ApplicationId)

			resp, err := client.GetApplication(ctx, applicationId, application.DefaultGetApplicationOperationOptions())
			if err != nil {
				return fmt.Errorf("retrieving %s: %+v", applicationId, err)
			}

			app := resp.Model
			if app == nil || app.AppRoles == nil {
				return fmt.Errorf("retrieving %s: appRoles was nil", applicationId)
			}

			// Look for a role to remove, matching by ID
			newRoles := make([]stable.AppRole, 0)
			found := false
			for _, existingRole := range *app.AppRoles {
				if strings.EqualFold(*existingRole.Id, id.RoleID) {
					found = true
				} else {
					newRoles = append(newRoles, existingRole)
				}
			}
			if !found {
				return fmt.Errorf("deleting %s: could not identify existing app role", id)
			}

			// Disable the existing role prior to update
			if err = applicationDisableAppRoles(ctx, client, applicationId, &newRoles); err != nil {
				return fmt.Errorf("disabling %s in preparation for deletion: %+v", id, err)
			}

			properties := stable.Application{
				Id:       &applicationId.ApplicationId,
				AppRoles: &newRoles,
			}

			// Patch the application with the new set of roles
			if _, err = client.UpdateApplication(ctx, applicationId, properties, application.DefaultUpdateApplicationOperationOptions()); err != nil {
				return fmt.Errorf("deleting %s: %+v", id, err)
			}

			return nil
		},
	}
}
