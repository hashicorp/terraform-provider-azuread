package msgraph

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/manicminer/hamilton/msgraph"

	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
)

func ApplicationFindByName(ctx context.Context, client *msgraph.ApplicationsClient, displayName string) (*msgraph.Application, error) {
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

func ApplicationFlattenApi(in *msgraph.ApplicationApi, dataSource bool) []map[string]interface{} {
	if in == nil {
		return []map[string]interface{}{}
	}

	api := make(map[string]interface{})

	oauth2PermissionScopes := ApplicationFlattenOAuth2PermissionScopes(in.OAuth2PermissionScopes)

	if dataSource {
		api["oauth2_permission_scopes"] = oauth2PermissionScopes
	} else {
		api["oauth2_permission_scope"] = oauth2PermissionScopes
	}

	return []map[string]interface{}{api}
}

func ApplicationFlattenAppRoles(in *[]msgraph.AppRole) []map[string]interface{} {
	if in == nil {
		return []map[string]interface{}{}
	}

	appRoles := make([]map[string]interface{}, 0)
	for _, role := range *in {
		roleId := ""
		if role.ID != nil {
			roleId = *role.ID
		}
		allowedMemberTypes := make([]interface{}, 0)
		if v := role.AllowedMemberTypes; v != nil {
			for _, m := range *v {
				allowedMemberTypes = append(allowedMemberTypes, m)
			}
		}
		description := ""
		if role.Description != nil {
			description = *role.Description
		}
		displayName := ""
		if role.DisplayName != nil {
			displayName = *role.DisplayName
		}
		enabled := false
		if role.IsEnabled != nil && *role.IsEnabled {
			enabled = true
		}
		value := ""
		if role.Value != nil {
			value = *role.Value
		}
		appRoles = append(appRoles, map[string]interface{}{
			"id":                   roleId,
			"allowed_member_types": allowedMemberTypes,
			"description":          description,
			"display_name":         displayName,
			"enabled":              enabled,
			"is_enabled":           enabled, // TODO: remove in v2.0
			"value":                value,
		})
	}

	return appRoles
}

func ApplicationFlattenGroupMembershipClaims(in *[]msgraph.GroupMembershipClaim) *string {
	if in == nil {
		return nil
	}
	result := make([]string, 0)
	for _, c := range *in {
		result = append(result, string(c))
	}
	// TODO: v2.0 this property should be a set/list
	ret := strings.Join(result, ",")
	return &ret
}

func ApplicationFlattenImplicitGrant(in *msgraph.ImplicitGrantSettings) []map[string]interface{} {
	if in == nil {
		return []map[string]interface{}{}
	}

	implicitGrant := map[string]interface{}{
		"access_token_issuance_enabled": in.EnableAccessTokenIssuance != nil && *in.EnableAccessTokenIssuance,
	}

	return []map[string]interface{}{implicitGrant}
}

func ApplicationFlattenOAuth2PermissionScopes(in *[]msgraph.PermissionScope) []map[string]interface{} {
	if in == nil {
		return []map[string]interface{}{}
	}

	result := make([]map[string]interface{}, 0)
	for _, p := range *in {
		adminConsentDescription := ""
		if v := p.AdminConsentDescription; v != nil {
			adminConsentDescription = *v
		}

		adminConsentDisplayName := ""
		if v := p.AdminConsentDisplayName; v != nil {
			adminConsentDisplayName = *v
		}

		id := ""
		if v := p.ID; v != nil {
			id = *v
		}

		userConsentDescription := ""
		if v := p.UserConsentDescription; v != nil {
			userConsentDescription = *v
		}

		userConsentDisplayName := ""
		if v := p.UserConsentDisplayName; v != nil {
			userConsentDisplayName = *v
		}

		value := ""
		if v := p.Value; v != nil {
			value = *v
		}

		enabled := p.IsEnabled != nil && *p.IsEnabled

		result = append(result, map[string]interface{}{
			"admin_consent_description":  adminConsentDescription,
			"admin_consent_display_name": adminConsentDisplayName,
			"id":                         id,
			"enabled":                    enabled,
			"type":                       p.Type,
			"user_consent_description":   userConsentDescription,
			"user_consent_display_name":  userConsentDisplayName,
			"value":                      value,
		})
	}

	return result
}

func ApplicationFlattenOAuth2Permissions(in *[]msgraph.PermissionScope) []map[string]interface{} {
	// TODO: v2.0 remove this func
	oauth2Permissions := ApplicationFlattenOAuth2PermissionScopes(in)

	if len(oauth2Permissions) == 0 {
		return []map[string]interface{}{}
	}

	result := make([]map[string]interface{}, 0)
	for _, p := range oauth2Permissions {
		result = append(result, map[string]interface{}{
			"admin_consent_description":  p["admin_consent_description"],
			"admin_consent_display_name": p["admin_consent_display_name"],
			"id":                         p["id"],
			"is_enabled":                 p["enabled"],
			"type":                       p["type"],
			"user_consent_description":   p["user_consent_description"],
			"user_consent_display_name":  p["user_consent_display_name"],
			"value":                      p["value"],
		})
	}

	return result
}

func ApplicationFlattenWeb(in *msgraph.ApplicationWeb) []map[string]interface{} {
	if in == nil {
		return []map[string]interface{}{}
	}

	homepageUrl := ""
	if in.HomePageUrl != nil {
		homepageUrl = *in.HomePageUrl
	}

	logoutUrl := ""
	if in.LogoutUrl != nil {
		logoutUrl = *in.LogoutUrl
	}

	var redirectUris []string
	if in.RedirectUris != nil {
		redirectUris = *in.RedirectUris
	}

	api := map[string]interface{}{
		"homepage_url":   homepageUrl,
		"logout_url":     logoutUrl,
		"redirect_uris":  redirectUris,
		"implicit_grant": ApplicationFlattenImplicitGrant(in.ImplicitGrantSettings),
	}

	return []map[string]interface{}{api}
}

func ApplicationSetAppRoles(ctx context.Context, client *msgraph.ApplicationsClient, application *msgraph.Application, newRoles *[]msgraph.AppRole) error {
	if application.ID == nil {
		return fmt.Errorf("cannot use Application model with nil ID")
	}

	if newRoles == nil {
		newRoles = &[]msgraph.AppRole{}
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
		properties := msgraph.Application{
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
	properties := msgraph.Application{
		ID:       application.ID,
		AppRoles: newRoles,
	}

	if _, err := client.Update(ctx, properties); err != nil {
		return fmt.Errorf("setting App Roles for Application with object ID %q: %+v", *application.ID, err)
	}

	return nil
}

func ApplicationSetOAuth2PermissionScopes(ctx context.Context, client *msgraph.ApplicationsClient, application *msgraph.Application, newScopes *[]msgraph.PermissionScope) error {
	if application.ID == nil {
		return fmt.Errorf("Cannot use Application model with nil ID")
	}

	if newScopes == nil {
		newScopes = &[]msgraph.PermissionScope{}
	}

	// OAuth2 Permission Scopes must be disabled before they can be edited or removed.
	// Since we cannot match them by ID, we have to disable all the scopes, and replace them in one pass.
	// TODO: v2.0 don't do this! we should be able to find the updated ones by ID and disable them selectively as we update them
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
		properties := msgraph.Application{
			ID: application.ID,
			Api: &msgraph.ApplicationApi{
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
	properties := msgraph.Application{
		ID: application.ID,
		Api: &msgraph.ApplicationApi{
			OAuth2PermissionScopes: newScopes,
		},
	}

	if _, err := client.Update(ctx, properties); err != nil {
		return fmt.Errorf("setting OAuth2 Permission Scopes for Application with object ID %q: %+v", *application.ID, err)
	}

	return nil
}

func ApplicationSetOwners(ctx context.Context, client *msgraph.ApplicationsClient, application *msgraph.Application, desiredOwners []string) error {
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

	if ownersToAdd != nil {
		for _, m := range ownersToAdd {
			application.AppendOwner(client.BaseClient.Endpoint, client.BaseClient.ApiVersion, m)
		}

		if _, err := client.AddOwners(ctx, application); err != nil {
			return fmt.Errorf("adding owners to Application with object ID %q: %+v", *application.ID, err)
		}
	}

	if ownersForRemoval != nil {
		if _, err = client.RemoveOwners(ctx, *application.ID, &ownersForRemoval); err != nil {
			return fmt.Errorf("removing owner from Application with object ID %q: %+v", *application.ID, err)
		}
	}

	return nil
}

func AppRoleFindById(app *msgraph.Application, roleId string) (*msgraph.AppRole, error) {
	if app == nil || app.AppRoles == nil {
		return nil, nil
	}

	if roleId == "" {
		return nil, fmt.Errorf("specified role ID is empty")
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

func OAuth2PermissionFindById(app *msgraph.Application, scopeId string) (*msgraph.PermissionScope, error) {
	if app == nil || app.Api == nil || app.Api.OAuth2PermissionScopes == nil {
		return nil, nil
	}

	if scopeId == "" {
		return nil, fmt.Errorf("specified scope ID is empty")
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
