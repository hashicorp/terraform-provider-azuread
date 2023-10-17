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

type ApplicationPermissionScopeModel struct {
	ApplicationId           string `tfschema:"application_id"`
	ScopeId                 string `tfschema:"scope_id"`
	AdminConsentDescription string `tfschema:"admin_consent_description"`
	AdminConsentDisplayName string `tfschema:"admin_consent_display_name"`
	Type                    string `tfschema:"type"`
	UserConsentDescription  string `tfschema:"user_consent_description"`
	UserConsentDisplayName  string `tfschema:"user_consent_display_name"`
	Value                   string `tfschema:"value"`
}

type ApplicationPermissionScopeResource struct{}

func (r ApplicationPermissionScopeResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return parse.ValidatePermissionScopeID
}

var _ sdk.ResourceWithUpdate = ApplicationPermissionScopeResource{}

func (r ApplicationPermissionScopeResource) ResourceType() string {
	return "azuread_application_permission_scope"
}

func (r ApplicationPermissionScopeResource) ModelObject() interface{} {
	return &ApplicationPermissionScopeModel{}
}

func (r ApplicationPermissionScopeResource) Arguments() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"application_id": {
			Description:  "The resource ID of the application to which this permission scope should be applied",
			Type:         pluginsdk.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: parse.ValidateApplicationID,
		},

		"scope_id": {
			Description:  "The unique identifier of the permission scope",
			Type:         pluginsdk.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.IsUUID,
		},

		"admin_consent_description": {
			Description:  "Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users",
			Type:         pluginsdk.TypeString,
			Required:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},

		"admin_consent_display_name": {
			Description:  "Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users",
			Type:         pluginsdk.TypeString,
			Required:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},

		"value": {
			Description:      "The value that is used for the `scp` claim in OAuth access tokens",
			Type:             pluginsdk.TypeString,
			Required:         true,
			ValidateDiagFunc: applicationsValidate.RoleScopeClaimValue,
		},

		"type": {
			Description: "Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions",
			Type:        pluginsdk.TypeString,
			Optional:    true,
			Default:     msgraph.PermissionScopeTypeUser,
			ValidateFunc: validation.StringInSlice([]string{
				msgraph.PermissionScopeTypeAdmin,
				msgraph.PermissionScopeTypeUser,
			}, false),
		},

		"user_consent_description": {
			Description:  "Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf",
			Type:         pluginsdk.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},

		"user_consent_display_name": {
			Description:  "Display name for the delegated permission that appears in the end user consent experience",
			Type:         pluginsdk.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringIsNotEmpty,
		},
	}
}

func (r ApplicationPermissionScopeResource) Attributes() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{}
}

func (r ApplicationPermissionScopeResource) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 10 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Applications.ApplicationsClient
			client.BaseClient.DisableRetries = true
			defer func() { client.BaseClient.DisableRetries = false }()

			var model ApplicationPermissionScopeModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			applicationId, err := parse.ParseApplicationID(model.ApplicationId)
			if err != nil {
				return err
			}

			id := parse.NewPermissionScopeID(applicationId.ApplicationId, model.ScopeId)

			tf.LockByName(applicationResourceName, id.ApplicationId)
			defer tf.UnlockByName(applicationResourceName, id.ApplicationId)

			result, _, err := client.Get(ctx, applicationId.ApplicationId, odata.Query{})
			if err != nil {
				return fmt.Errorf("retrieving %s: %+v", applicationId, err)
			}
			if result == nil {
				return fmt.Errorf("retrieving %s: result was nil", applicationId)
			}

			newScopes := make([]msgraph.PermissionScope, 0)

			// Don't forget any existing scopes, since all scopes must be updated together
			if result.Api != nil && result.Api.OAuth2PermissionScopes != nil {
				newScopes = *result.Api.OAuth2PermissionScopes
			}

			// Check for existing scope ID
			for _, scope := range newScopes {
				if strings.EqualFold(*scope.ID, id.ScopeID) {
					return metadata.ResourceRequiresImport(r.ResourceType(), id)
				}
			}

			newScopes = append(newScopes, msgraph.PermissionScope{
				ID:                      &model.ScopeId,
				IsEnabled:               pointer.To(true),
				AdminConsentDescription: &model.AdminConsentDescription,
				AdminConsentDisplayName: &model.AdminConsentDisplayName,
				Type:                    model.Type,
				UserConsentDescription:  &model.UserConsentDescription,
				UserConsentDisplayName:  &model.UserConsentDisplayName,
				Value:                   &model.Value,
			})

			properties := msgraph.Application{
				DirectoryObject: msgraph.DirectoryObject{
					Id: &id.ApplicationId,
				},
				Api: &msgraph.ApplicationApi{
					OAuth2PermissionScopes: &newScopes,
				},
			}

			if _, err = client.Update(ctx, properties); err != nil {
				return fmt.Errorf("creating %s: %+v", id, err)
			}

			metadata.SetID(id)
			return nil
		},
	}
}

func (r ApplicationPermissionScopeResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Applications.ApplicationsClient
			client.BaseClient.DisableRetries = true
			defer func() { client.BaseClient.DisableRetries = false }()

			id, err := parse.ParsePermissionScopeID(metadata.ResourceData.Id())
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
			if result.Api == nil || result.Api.OAuth2PermissionScopes == nil {
				return metadata.MarkAsGone(id)
			}

			// Identify the scope by ID
			var scope *msgraph.PermissionScope
			for _, existingScope := range *result.Api.OAuth2PermissionScopes {
				if strings.EqualFold(*existingScope.ID, id.ScopeID) {
					scope = &existingScope
					break
				}
			}

			if scope == nil {
				return metadata.MarkAsGone(id)
			}

			state := ApplicationPermissionScopeModel{
				ApplicationId:           applicationId.ID(),
				ScopeId:                 id.ScopeID,
				AdminConsentDescription: pointer.From(scope.AdminConsentDescription),
				AdminConsentDisplayName: pointer.From(scope.AdminConsentDisplayName),
				Type:                    scope.Type,
				UserConsentDescription:  pointer.From(scope.UserConsentDescription),
				UserConsentDisplayName:  pointer.From(scope.UserConsentDisplayName),
				Value:                   pointer.From(scope.Value),
			}

			return metadata.Encode(&state)
		},
	}
}

func (r ApplicationPermissionScopeResource) Update() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 10 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Applications.ApplicationsClient

			id, err := parse.ParsePermissionScopeID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			var model ApplicationPermissionScopeModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			// Prepare a new scope to replace the existing one
			scope := msgraph.PermissionScope{
				ID:                      &id.ScopeID,
				IsEnabled:               pointer.To(true),
				AdminConsentDescription: &model.AdminConsentDescription,
				AdminConsentDisplayName: &model.AdminConsentDisplayName,
				Type:                    model.Type,
				UserConsentDescription:  &model.UserConsentDescription,
				UserConsentDisplayName:  &model.UserConsentDisplayName,
				Value:                   &model.Value,
			}

			tf.LockByName(applicationResourceName, id.ApplicationId)
			defer tf.UnlockByName(applicationResourceName, id.ApplicationId)

			applicationId := parse.NewApplicationID(id.ApplicationId)
			result, _, err := client.Get(ctx, applicationId.ApplicationId, odata.Query{})
			if err != nil {
				return fmt.Errorf("retrieving %s: %+v", applicationId, err)
			}
			if result == nil || result.Api == nil || result.Api.OAuth2PermissionScopes == nil {
				return fmt.Errorf("retrieving %s: api.oauth2PermissionScopes was nil", applicationId)
			}

			// Look for a scope to replace, matching by ID
			newScopes := make([]msgraph.PermissionScope, 0)
			found := false
			for _, existingScope := range *result.Api.OAuth2PermissionScopes {
				if strings.EqualFold(*existingScope.ID, id.ScopeID) {
					newScopes = append(newScopes, scope)
					found = true
				} else {
					newScopes = append(newScopes, existingScope)
				}
			}
			if !found {
				return fmt.Errorf("updating %s: could not identify existing permission scope", id)
			}

			// Disable the existing scope prior to update
			if err = applicationDisableOauth2PermissionScopes(ctx, client, result, &newScopes); err != nil {
				return fmt.Errorf("disabling %s in preparation for update: %+v", id, err)
			}

			properties := msgraph.Application{
				DirectoryObject: msgraph.DirectoryObject{
					Id: &applicationId.ApplicationId,
				},
				Api: &msgraph.ApplicationApi{
					OAuth2PermissionScopes: &newScopes,
				},
			}

			// Patch the application with the new set of scopes
			_, err = client.Update(ctx, properties)
			if err != nil {
				return fmt.Errorf("updating %s: %+v", id, err)
			}

			return nil
		},
	}
}

func (r ApplicationPermissionScopeResource) Delete() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Applications.ApplicationsClient
			client.BaseClient.DisableRetries = true
			defer func() { client.BaseClient.DisableRetries = false }()

			id, err := parse.ParsePermissionScopeID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			var model ApplicationPermissionScopeModel
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
			if result == nil || result.Api == nil || result.Api.OAuth2PermissionScopes == nil {
				return fmt.Errorf("retrieving %s: api.oauth2PermissionScopes was nil", applicationId)
			}

			// Look for a scope to remove, matching by ID
			newScopes := make([]msgraph.PermissionScope, 0)
			found := false
			for _, existingScope := range *result.Api.OAuth2PermissionScopes {
				if strings.EqualFold(*existingScope.ID, id.ScopeID) {
					found = true
				} else {
					newScopes = append(newScopes, existingScope)
				}
			}
			if !found {
				return fmt.Errorf("deleting %s: could not identify existing permission scope", id)
			}

			// Disable the existing scope prior to update
			if err = applicationDisableOauth2PermissionScopes(ctx, client, result, &newScopes); err != nil {
				return fmt.Errorf("disabling %s in preparation for deletion: %+v", id, err)
			}

			properties := msgraph.Application{
				DirectoryObject: msgraph.DirectoryObject{
					Id: &applicationId.ApplicationId,
				},
				Api: &msgraph.ApplicationApi{
					OAuth2PermissionScopes: &newScopes,
				},
			}

			// Patch the application with the new set of scopes
			_, err = client.Update(ctx, properties)
			if err != nil {
				return fmt.Errorf("deleting %s: %+v", id, err)
			}

			return nil
		},
	}
}
