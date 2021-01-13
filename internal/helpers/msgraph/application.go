package msgraph

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/manicminer/hamilton/clients"
	"github.com/manicminer/hamilton/models"

	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
)

func ApplicationFindByName(ctx context.Context, client *clients.ApplicationsClient, displayName string) (*models.Application, error) {
	filter := fmt.Sprintf("displayName eq '%s'", displayName)
	result, _, err := client.List(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("unable to list Applications with filter %q: %+v", filter, err)
	}

	if result != nil {
		for _, app := range *result {
			if app.DisplayName != nil && *app.DisplayName == displayName {
				return &app, nil
			}
		}
	}

	return nil, nil
}

func ApplicationFlattenAppRoles(in *[]models.AppRole) []map[string]interface{} {
	if in == nil {
		return []map[string]interface{}{}
	}

	appRoles := make([]map[string]interface{}, 0, len(*in))
	for _, role := range *in {
		appRole := make(map[string]interface{})
		if v := role.ID; v != nil {
			appRole["id"] = v
		}
		if v := role.AllowedMemberTypes; v != nil {
			memberTypes := make([]interface{}, 0, len(*v))
			for _, m := range *v {
				memberTypes = append(memberTypes, m)
			}
			appRole["allowed_member_types"] = memberTypes
		}
		if v := role.Description; v != nil {
			appRole["description"] = v
		}
		if v := role.DisplayName; v != nil {
			appRole["display_name"] = v
		}
		if v := role.IsEnabled; v != nil {
			appRole["is_enabled"] = v
		}
		if v := role.Value; v != nil {
			appRole["value"] = v
		}
		appRoles = append(appRoles, appRole)
	}

	return appRoles
}

func ApplicationFlattenOAuth2Permissions(in *[]models.PermissionScope) []map[string]interface{} {
	if in == nil {
		return []map[string]interface{}{}
	}

	result := make([]map[string]interface{}, 0, len(*in))
	for _, p := range *in {
		permission := map[string]interface{}{
			"admin_consent_description":  "",
			"admin_consent_display_name": "",
			"id":                         "",
			"is_enabled":                 false,
			"type":                       "",
			"user_consent_description":   "",
			"user_consent_display_name":  "",
			"value":                      "",
		}

		if v := p.AdminConsentDescription; v != nil {
			permission["admin_consent_description"] = v
		}

		if v := p.AdminConsentDisplayName; v != nil {
			permission["admin_consent_display_name"] = v
		}

		if v := p.ID; v != nil {
			permission["id"] = v
		}

		if v := p.IsEnabled; v != nil {
			permission["is_enabled"] = *v
		}

		if v := p.Type; v != nil {
			permission["type"] = v
		}

		if v := p.UserConsentDescription; v != nil {
			permission["user_consent_description"] = v
		}

		if v := p.UserConsentDisplayName; v != nil {
			permission["user_consent_display_name"] = v
		}

		if v := p.Value; v != nil {
			permission["value"] = v
		}

		result = append(result, permission)
	}

	return result
}

func ApplicationSetAppRoles(ctx context.Context, client *clients.ApplicationsClient, application *models.Application, newRoles *[]models.AppRole) error {
	if application.ID == nil {
		return fmt.Errorf("Cannot use Application model with nil ID")
	}

	if newRoles == nil {
		newRoles = &[]models.AppRole{}
	}

	// Roles must be disabled before they can be edited or removed.
	// Since we cannot match them by ID, we have to disable all the roles, and replace them in one pass.
	app, status, err := client.Get(ctx, *application.ID)
	if err != nil {
		if status == http.StatusNotFound {
			return fmt.Errorf("application with ID %q was not found", *application.ID)
		}

		return fmt.Errorf("retrieving Application with object ID %q: %+v", *application.ID, err)
	}

	// don't update if no changes to be made
	if app.AppRoles != nil && reflect.DeepEqual(*app.AppRoles, *newRoles) {
		return nil
	}

	// first disable any existing roles
	if app.AppRoles != nil && len(*app.AppRoles) > 0 {
		properties := models.Application{
			ID:       application.ID,
			AppRoles: app.AppRoles,
		}

		for _, role := range *properties.AppRoles {
			*role.IsEnabled = false
		}

		if _, err := client.Update(ctx, properties); err != nil {
			return fmt.Errorf("disabling App Roles for Application with object ID %q: %+v", *application.ID, err)
		}
	}

	// then set the new roles
	properties := models.Application{
		ID:       application.ID,
		AppRoles: newRoles,
	}

	if _, err := client.Update(ctx, properties); err != nil {
		return fmt.Errorf("setting App Roles for Application with object ID %q: %+v", *application.ID, err)
	}

	return nil
}

func ApplicationSetOAuth2PermissionScopes(ctx context.Context, client *clients.ApplicationsClient, application *models.Application, newScopes *[]models.PermissionScope) error {
	if application.ID == nil {
		return fmt.Errorf("Cannot use Application model with nil ID")
	}

	if newScopes == nil {
		newScopes = &[]models.PermissionScope{}
	}

	// OAuth2 Permission Scopes must be disabled before they can be edited or removed.
	// Since we cannot match them by ID, we have to disable all the scopes, and replace them in one pass.
	app, status, err := client.Get(ctx, *application.ID)
	if err != nil {
		if status == http.StatusNotFound {
			return fmt.Errorf("application with ID %q was not found", *application.ID)
		}

		return fmt.Errorf("retrieving Application with object ID %q: %+v", *application.ID, err)
	}

	// don't update if no changes to be made
	if app.Api != nil && app.Api.OAuth2PermissionScopes != nil && reflect.DeepEqual(*app.Api.OAuth2PermissionScopes, *newScopes) {
		return nil
	}

	// first disable any existing scopes
	if app.Api != nil && app.Api.OAuth2PermissionScopes != nil && len(*app.Api.OAuth2PermissionScopes) > 0 {
		properties := models.Application{
			ID: application.ID,
			Api: &models.ApplicationApi{
				OAuth2PermissionScopes: app.Api.OAuth2PermissionScopes,
			},
		}

		for _, scope := range *properties.Api.OAuth2PermissionScopes {
			*scope.IsEnabled = false
		}

		if _, err := client.Update(ctx, properties); err != nil {
			return fmt.Errorf("disabling OAuth2 Permission Scopes for Application with object ID %q: %+v", *application.ID, err)
		}
	}

	// then set the new scopes
	properties := models.Application{
		ID: application.ID,
		Api: &models.ApplicationApi{
			OAuth2PermissionScopes: newScopes,
		},
	}

	if _, err := client.Update(ctx, properties); err != nil {
		return fmt.Errorf("setting OAuth2 Permission Scopes for Application with object ID %q: %+v", *application.ID, err)
	}

	return nil
}

func ApplicationSetOwners(ctx context.Context, client *clients.ApplicationsClient, application *models.Application, desiredOwners []string) error {
	if application.ID == nil {
		return fmt.Errorf("Cannot use Application model with nil ID")
	}

	owners, _, err := client.ListOwners(ctx, *application.ID)
	if err != nil {
		return fmt.Errorf("retrieving owners for Application with object ID %q: %+v", *application.ID, err)
	}

	existingOwners := *owners
	ownersForRemoval := utils.Difference(existingOwners, desiredOwners)
	ownersToAdd := utils.Difference(desiredOwners, existingOwners)

	if ownersForRemoval != nil {
		if _, err = client.RemoveOwners(ctx, *application.ID, &ownersForRemoval); err != nil {
			return fmt.Errorf("removing owner from Application with object ID %q: %+v", *application.ID, err)
		}
	}

	if ownersToAdd != nil {
		for _, m := range ownersToAdd {
			application.AppendOwner(client.BaseClient.Endpoint, client.BaseClient.ApiVersion, m)
		}

		if _, err := client.AddOwners(ctx, application); err != nil {
			return err
		}
	}
	return nil
}

func AppRoleFindById(app *models.Application, roleId string) (*models.AppRole, error) {
	if app == nil || app.AppRoles == nil {
		return nil, nil
	}

	if roleId == "" {
		return nil, errors.New("specified role ID is blank")
	}

	for _, r := range *app.AppRoles {
		if r.ID == nil {
			continue
		}
		if *r.ID == roleId {
			return &r, nil
		}
	}

	return nil, nil
}

func OAuth2PermissionFindById(app *models.Application, scopeId string) (*models.PermissionScope, error) {
	if app == nil || app.Api == nil || app.Api.OAuth2PermissionScopes == nil {
		return nil, nil
	}

	if scopeId == "" {
		return nil, errors.New("specified scope ID is blank")
	}

	for _, s := range *app.Api.OAuth2PermissionScopes {
		if s.ID == nil {
			continue
		}
		if *s.ID == scopeId {
			return &s, nil
		}
	}

	return nil, nil
}
