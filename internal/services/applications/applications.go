// Copyright IBM Corp. 2023, 2026
// SPDX-License-Identifier: MPL-2.0

package applications

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"reflect"
	"slices"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/application"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/applications"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/credentials"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
)

func applicationUpdateRetryFunc() client.RequestRetryFunc {
	return func(resp *http.Response, o *odata.OData) (bool, error) {
		// inconsistent state on this resource can result in a 409 Conflict being returned, in this exception case we retry as per the service design, see  https://github.com/hashicorp/terraform-provider-azuread/issues/1764#issuecomment-3282278691 for more information.
		if response.WasNotFound(resp) || response.WasConflict(resp) {
			return true, nil
		} else if response.WasBadRequest(resp) && o != nil && o.Error != nil {
			return o.Error.Match("Permission (scope or role) cannot be deleted or updated unless disabled first"), nil
		}

		return false, nil
	}
}

func applicationAppRoleChanged(existingRole stable.AppRole, newRole stable.AppRole) bool {
	if !reflect.DeepEqual(existingRole.AllowedMemberTypes, newRole.AllowedMemberTypes) {
		return true
	}
	if !reflect.DeepEqual(existingRole.Description, newRole.Description) {
		return true
	}
	if !reflect.DeepEqual(existingRole.DisplayName, newRole.DisplayName) {
		return true
	}

	if reflect.DeepEqual(existingRole.Value, newRole.Value) {
		return false
	}

	// We consider unset/null to be equivalent to the zero value
	if existingRole.Value.GetOrZero() == newRole.Value.GetOrZero() {
		return false
	}

	return true
}

// applicationPermissions holds the app role and OAuth2 permission scope collections
// involved in an application update. A nil collection is one the caller does not
// manage, and is left out of the request so that Microsoft Graph leaves it untouched.
type applicationPermissions struct {
	AppRoles               *[]stable.AppRole
	OAuth2PermissionScopes *[]stable.PermissionScope
}

// applyTo sets the permission collections on an application update payload.
func (p applicationPermissions) applyTo(properties *stable.Application) {
	properties.AppRoles = p.AppRoles

	if p.OAuth2PermissionScopes != nil {
		if properties.Api == nil {
			properties.Api = &stable.ApiApplication{}
		}
		properties.Api.OAuth2PermissionScopes = p.OAuth2PermissionScopes
	}
}

// applicationDisableChangedPermissions disables any app roles and OAuth2 permission
// scopes that are being changed or removed, which Microsoft Graph requires before a
// permission can be modified, then waits for the application manifest to reflect the
// change.
//
// Pass nil for a collection the caller does not manage. The returned collections are
// those the caller must send in its follow-up request: an unmanaged collection is
// returned unchanged when it holds the counterpart of a shared permission, so that the
// counterpart is re-enabled in the same request that applies the change.
func applicationDisableChangedPermissions(ctx context.Context, client *application.ApplicationClient, applicationId stable.ApplicationId, newRoles *[]stable.AppRole, newScopes *[]stable.PermissionScope) (applicationPermissions, error) {
	permissions := applicationPermissions{AppRoles: newRoles, OAuth2PermissionScopes: newScopes}

	resp, err := client.GetApplication(ctx, applicationId, application.DefaultGetApplicationOperationOptions())
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return permissions, fmt.Errorf("%s was not found", applicationId)
		}

		return permissions, fmt.Errorf("retrieving %s: %+v", applicationId, err)
	}

	app := resp.Model
	if app == nil {
		return permissions, fmt.Errorf("retrieving %s: model was nil", applicationId)
	}

	var existingRoles []stable.AppRole
	if app.AppRoles != nil {
		existingRoles = *app.AppRoles
	}

	var existingScopes []stable.PermissionScope
	if app.Api != nil && app.Api.OAuth2PermissionScopes != nil {
		existingScopes = *app.Api.OAuth2PermissionScopes
	}

	roleIds := make(map[string]struct{})
	if newRoles != nil {
		if roleIds, err = appRoleIdsToDisable(existingRoles, *newRoles); err != nil {
			return permissions, err
		}
	}

	scopeIds := make(map[string]struct{})
	if newScopes != nil {
		if scopeIds, err = permissionScopeIdsToDisable(existingScopes, *newScopes); err != nil {
			return permissions, err
		}
	}

	// An app role and an OAuth2 permission scope may share an ID, in which case Microsoft
	// Graph requires them to agree on their common properties. Neither can be disabled
	// without the other, so disabling one also disables its counterpart, even when the
	// caller does not manage that collection.
	desiredRoles, desiredScopes := existingRoles, existingScopes
	if newRoles != nil {
		desiredRoles = *newRoles
	}
	if newScopes != nil {
		desiredScopes = *newScopes
	}

	sharedIds := applicationSharedPermissionIds(existingRoles, existingScopes)
	disablingShared := make(map[string]struct{})
	for id := range sharedIds {
		_, disablingRole := roleIds[id]
		_, disablingScope := scopeIds[id]
		if disablingRole || disablingScope {
			disablingShared[id] = struct{}{}
		}
	}

	if len(disablingShared) > 0 {
		// Graph would reject the caller's follow-up request, by which point this one has
		// already disabled both sides, so fail before changing anything.
		if err = applicationValidateSharedPermissions(desiredRoles, desiredScopes); err != nil {
			return permissions, err
		}

		for id := range disablingShared {
			roleIds[id] = struct{}{}
			scopeIds[id] = struct{}{}
		}

		permissions.AppRoles = &desiredRoles
		permissions.OAuth2PermissionScopes = &desiredScopes
	}

	disabledRoles := appRolesDisabling(existingRoles, roleIds)
	disabledScopes := permissionScopesDisabling(existingScopes, scopeIds)

	// Shortcut: don't update if there is nothing to disable
	if disabledRoles == nil && disabledScopes == nil {
		return permissions, nil
	}

	properties := stable.Application{
		Id:       app.Id,
		AppRoles: disabledRoles,
	}
	if disabledScopes != nil {
		properties.Api = &stable.ApiApplication{
			OAuth2PermissionScopes: disabledScopes,
		}
	}

	if _, err = client.UpdateApplication(ctx, applicationId, properties, application.DefaultUpdateApplicationOperationOptions()); err != nil {
		return permissions, fmt.Errorf("disabling App Roles and OAuth2 Permission Scopes for %s: %+v", applicationId, err)
	}

	if err = applicationWaitForPermissionsDisabled(ctx, client, applicationId, disabledRoles, disabledScopes); err != nil {
		return permissions, err
	}

	return permissions, nil
}

// applicationSharedPermissionIds returns the IDs that are used by both an app role and
// an OAuth2 permission scope. Microsoft Graph permits an application to expose the same
// permission as both, but requires the two entries to agree on their common properties.
func applicationSharedPermissionIds(roles []stable.AppRole, scopes []stable.PermissionScope) map[string]struct{} {
	roleIds := make(map[string]struct{}, len(roles))
	for _, role := range roles {
		if role.Id != nil {
			roleIds[*role.Id] = struct{}{}
		}
	}

	shared := make(map[string]struct{})
	for _, scope := range scopes {
		if scope.Id == nil {
			continue
		}

		if _, isRole := roleIds[*scope.Id]; isRole {
			shared[*scope.Id] = struct{}{}
		}
	}

	return shared
}

// applicationValidateSharedPermissions returns an error when an app role and an OAuth2
// permission scope share an ID but disagree on the properties Microsoft Graph requires
// to match. Origin is read-only and never sent, so it is not compared.
//
// The azuread_application resource rejects this at plan time, but the app role and
// permission scope resources cannot see each other's configuration, so a mismatch
// between them is only detectable once both collections have been retrieved.
func applicationValidateSharedPermissions(roles []stable.AppRole, scopes []stable.PermissionScope) error {
	rolesById := make(map[string]stable.AppRole, len(roles))
	for _, role := range roles {
		if role.Id != nil {
			rolesById[*role.Id] = role
		}
	}

	for _, scope := range scopes {
		if scope.Id == nil {
			continue
		}

		role, shared := rolesById[*scope.Id]
		if !shared {
			continue
		}

		if role.Description.GetOrZero() == scope.AdminConsentDescription.GetOrZero() &&
			role.DisplayName.GetOrZero() == scope.AdminConsentDisplayName.GetOrZero() &&
			role.Value.GetOrZero() == scope.Value.GetOrZero() &&
			pointer.From(role.IsEnabled) == pointer.From(scope.IsEnabled) {
			continue
		}

		return fmt.Errorf("the following values must match for the 'oauth2Permissions' and 'appRoles' properties with identifier %q: (description, adminConsentDescription), (displayName, adminConsentDisplayName), (isEnabled, isEnabled), (value, value). An app role and a permission scope sharing an ID must be updated together, which the azuread_application_app_role and azuread_application_permission_scope resources cannot do; declare both permissions in the app_role and api.oauth2_permission_scope blocks of azuread_application instead", *scope.Id)
	}

	return nil
}

// appRoleIdsToDisable returns the IDs of the app roles that newRoles changes or removes,
// and which must therefore be disabled before the change can be made.
func appRoleIdsToDisable(existingRoles, newRoles []stable.AppRole) (map[string]struct{}, error) {
	newRolesById := make(map[string]stable.AppRole, len(newRoles))
	for _, newRole := range newRoles {
		if newRole.Id == nil || *newRole.Id == "" {
			return nil, fmt.Errorf("new role provided with nil or empty ID")
		}

		newRolesById[*newRole.Id] = newRole
	}

	ids := make(map[string]struct{})
	for _, existingRole := range existingRoles {
		if existingRole.Id == nil || existingRole.IsEnabled == nil || !*existingRole.IsEnabled {
			continue
		}

		// A role that is absent from newRoles is being removed, and must also be disabled first
		if newRole, found := newRolesById[*existingRole.Id]; found && !applicationAppRoleChanged(existingRole, newRole) {
			continue
		}

		ids[*existingRole.Id] = struct{}{}
	}

	return ids, nil
}

// permissionScopeIdsToDisable returns the IDs of the OAuth2 permission scopes that
// newScopes changes or removes, and which must therefore be disabled before the change
// can be made.
func permissionScopeIdsToDisable(existingScopes, newScopes []stable.PermissionScope) (map[string]struct{}, error) {
	newScopesById := make(map[string]stable.PermissionScope, len(newScopes))
	for _, newScope := range newScopes {
		if newScope.Id == nil || *newScope.Id == "" {
			return nil, fmt.Errorf("new scope provided with nil or empty ID")
		}

		newScopesById[*newScope.Id] = newScope
	}

	ids := make(map[string]struct{})
	for _, existingScope := range existingScopes {
		if existingScope.Id == nil || existingScope.IsEnabled == nil || !*existingScope.IsEnabled {
			continue
		}

		// A scope that is absent from newScopes is being removed, and must also be disabled first
		if newScope, found := newScopesById[*existingScope.Id]; found && reflect.DeepEqual(existingScope, newScope) {
			continue
		}

		ids[*existingScope.Id] = struct{}{}
	}

	return ids, nil
}

// appRolesDisabling returns a copy of roles in which the given permission IDs are
// disabled, or nil when none of them appears in the collection. roles is not modified.
func appRolesDisabling(roles []stable.AppRole, ids map[string]struct{}) *[]stable.AppRole {
	if len(ids) == 0 {
		return nil
	}

	disabled := slices.Clone(roles)
	anyDisabled := false

	for i, role := range disabled {
		if role.Id == nil {
			continue
		}
		if _, disable := ids[*role.Id]; !disable {
			continue
		}

		disabled[i].IsEnabled = pointer.To(false)
		anyDisabled = true
	}

	if !anyDisabled {
		return nil
	}

	return &disabled
}

// permissionScopesDisabling returns a copy of scopes in which the given permission IDs
// are disabled, or nil when none of them appears in the collection. scopes is not
// modified.
func permissionScopesDisabling(scopes []stable.PermissionScope, ids map[string]struct{}) *[]stable.PermissionScope {
	if len(ids) == 0 {
		return nil
	}

	disabled := slices.Clone(scopes)
	anyDisabled := false

	for i, scope := range disabled {
		if scope.Id == nil {
			continue
		}
		if _, disable := ids[*scope.Id]; !disable {
			continue
		}

		disabled[i].IsEnabled = pointer.To(false)
		anyDisabled = true
	}

	if !anyDisabled {
		return nil
	}

	return &disabled
}

// applicationWaitForPermissionsDisabled waits until the application manifest reports
// every permission that was sent as disabled as being disabled. Collections that were
// not sent are not inspected.
func applicationWaitForPermissionsDisabled(ctx context.Context, client *application.ApplicationClient, applicationId stable.ApplicationId, sentRoles *[]stable.AppRole, sentScopes *[]stable.PermissionScope) error {
	deadline, ok := ctx.Deadline()
	if !ok {
		return fmt.Errorf("context has no deadline")
	}

	disabledRoleIds := make(map[string]struct{})
	if sentRoles != nil {
		for _, role := range *sentRoles {
			if role.Id != nil && role.IsEnabled != nil && !*role.IsEnabled {
				disabledRoleIds[*role.Id] = struct{}{}
			}
		}
	}

	disabledScopeIds := make(map[string]struct{})
	if sentScopes != nil {
		for _, scope := range *sentScopes {
			if scope.Id != nil && scope.IsEnabled != nil && !*scope.IsEnabled {
				disabledScopeIds[*scope.Id] = struct{}{}
			}
		}
	}

	_, err := (&pluginsdk.StateChangeConf{ //nolint:staticcheck
		Pending:    []string{"Waiting"},
		Target:     []string{"Disabled"},
		Timeout:    time.Until(deadline),
		MinTimeout: 1 * time.Second,
		Refresh: func() (interface{}, string, error) {
			resp, err := client.GetApplication(ctx, applicationId, application.DefaultGetApplicationOperationOptions())
			if err != nil {
				return nil, "Error", fmt.Errorf("retrieving %s: %+v", applicationId, err)
			}

			app := resp.Model
			if app == nil {
				return nil, "Error", fmt.Errorf("reading permissions for %s: model was nil", applicationId)
			}

			if app.AppRoles != nil {
				for _, actualRole := range *app.AppRoles {
					if actualRole.Id == nil || actualRole.IsEnabled == nil || !*actualRole.IsEnabled {
						continue
					}
					if _, expectedDisabled := disabledRoleIds[*actualRole.Id]; expectedDisabled {
						return app, "Waiting", nil
					}
				}
			}

			if app.Api != nil && app.Api.OAuth2PermissionScopes != nil {
				for _, actualScope := range *app.Api.OAuth2PermissionScopes {
					if actualScope.Id == nil || actualScope.IsEnabled == nil || !*actualScope.IsEnabled {
						continue
					}
					if _, expectedDisabled := disabledScopeIds[*actualScope.Id]; expectedDisabled {
						return app, "Waiting", nil
					}
				}
			}

			return app, "Disabled", nil
		},
	}).WaitForStateContext(ctx)
	if err != nil {
		return fmt.Errorf("waiting for App Roles and OAuth2 Permission Scopes to be disabled for %s: %+v", applicationId, err)
	}

	return nil
}

func applicationFindByName(ctx context.Context, client *application.ApplicationClient, displayName string) (*[]stable.Application, error) {
	options := application.ListApplicationsOperationOptions{
		Filter: pointer.To(fmt.Sprintf("displayName eq '%s'", displayName)),
	}
	resp, err := client.ListApplications(ctx, options)
	if err != nil {
		return nil, fmt.Errorf("unable to list Applications with filter %q: %+v", *options.Filter, err)
	}

	result := make([]stable.Application, 0)
	if apps := resp.Model; apps != nil {
		for _, app := range *apps {
			if app.DisplayName.GetOrZero() == displayName {
				result = append(result, app)
			}
		}
	}

	return &result, nil
}

func applicationParseLogoImage(encodedImage string) (string, []byte, error) {
	imageData, err := base64.StdEncoding.DecodeString(strings.TrimSpace(encodedImage))
	if err != nil {
		return "", nil, err
	}
	contentType := http.DetectContentType(imageData)
	if !strings.HasPrefix(contentType, "image/") {
		return "", nil, fmt.Errorf("unrecognised MIME type detected: %q", contentType)
	}
	return contentType, imageData, nil
}

func applicationValidateRolesScopes(appRoles, oauth2Permissions []interface{}) error {
	type appPermission struct {
		id          string
		displayName string
		description string
		enabled     bool
		value       string
	}
	var appPermissions []appPermission

	for _, roleRaw := range appRoles {
		if roleRaw == nil {
			continue
		}
		role := roleRaw.(map[string]interface{})
		permission := appPermission{
			id:          role["id"].(string),
			displayName: role["display_name"].(string),
			description: role["description"].(string),
			enabled:     role["enabled"].(bool),
			value:       role["value"].(string),
		}
		if pluginsdk.ValueIsNotEmptyOrUnknown(permission.id) && pluginsdk.ValueIsNotEmptyOrUnknown(permission.value) {
			appPermissions = append(appPermissions, permission)
		}
	}

	for _, scopeRaw := range oauth2Permissions {
		if scopeRaw == nil {
			continue
		}
		scope := scopeRaw.(map[string]interface{})
		permission := appPermission{
			id:          scope["id"].(string),
			displayName: scope["admin_consent_display_name"].(string),
			description: scope["admin_consent_description"].(string),
			enabled:     scope["enabled"].(bool),
			value:       scope["value"].(string),
		}
		if pluginsdk.ValueIsNotEmptyOrUnknown(permission.id) && pluginsdk.ValueIsNotEmptyOrUnknown(permission.value) {
			appPermissions = append(appPermissions, permission)
		}
	}

	encounteredPermissions := make([]appPermission, 0)
	for _, ap := range appPermissions {
		for _, ep := range encounteredPermissions {
			if ap.id == ep.id && ap.value != ep.value {
				return fmt.Errorf("validation failed: duplicate ID found: %q", ap.id)
			}
			if ap.value == ep.value && ap.id != ep.id {
				return fmt.Errorf("validation failed: duplicate value found: %q", ap.value)
			}
			if ap.value == ep.value && ap.id == ep.id && !reflect.DeepEqual(ap, ep) {
				return fmt.Errorf(`validation failed: The following values must match for the
				'oauth2Permissions' and 'appRoles' properties with identifier '%q': (description, adminConsentDescription),
				(displayName, adminConsentDisplayName),(isEnabled,isEnabled),(origin, origin),(value, value).
				Ensure that you are intending to have entries with the same identifier, and if so, are updating them together`, ap.id)
			}
		}
		encounteredPermissions = append(encounteredPermissions, ap)
	}

	return nil
}

func expandApplicationApi(input []interface{}) (result *stable.ApiApplication) {
	result = &stable.ApiApplication{
		AcceptMappedClaims:          nullable.Value(false),
		KnownClientApplications:     &[]string{},
		OAuth2PermissionScopes:      &[]stable.PermissionScope{},
		RequestedAccessTokenVersion: nullable.Value(int64(1)),
	}

	if len(input) == 0 || input[0] == nil {
		return
	}

	in := input[0].(map[string]interface{})
	if v, ok := in["mapped_claims_enabled"]; ok {
		result.AcceptMappedClaims.Set(v.(bool))
	}
	if v, ok := in["known_client_applications"]; ok {
		result.KnownClientApplications = tf.ExpandStringSlicePtr(v.(*pluginsdk.Set).List())
	}
	result.OAuth2PermissionScopes = expandApplicationOAuth2PermissionScope(in["oauth2_permission_scope"].(*pluginsdk.Set).List())
	if v, ok := in["requested_access_token_version"]; ok {
		result.RequestedAccessTokenVersion.Set(int64(v.(int)))
	}

	return
}

func expandApplicationPasswordCredentials(input []interface{}) (*[]stable.PasswordCredential, error) {
	if len(input) == 0 {
		return nil, nil
	}

	result := make([]stable.PasswordCredential, 0)

	for _, password := range input {
		if password == nil {
			continue
		}

		credential, err := credentials.PasswordCredential(password.(map[string]interface{}))
		if err != nil {
			return nil, err
		}
		result = append(result, *credential)
	}

	return &result, nil
}

func expandApplicationAppRoles(input []interface{}) *[]stable.AppRole {
	result := make([]stable.AppRole, 0)

	if len(input) == 0 {
		return &result
	}

	for _, appRoleRaw := range input {
		if appRoleRaw == nil {
			continue
		}
		appRole := appRoleRaw.(map[string]interface{})

		allowedMemberTypes := make([]string, 0)
		for _, allowedMemberType := range appRole["allowed_member_types"].(*pluginsdk.Set).List() {
			allowedMemberTypes = append(allowedMemberTypes, allowedMemberType.(string))
		}

		newAppRole := stable.AppRole{
			Id:                 pointer.To(appRole["id"].(string)),
			AllowedMemberTypes: &allowedMemberTypes,
			Description:        nullable.Value(appRole["description"].(string)),
			DisplayName:        nullable.Value(appRole["display_name"].(string)),
			IsEnabled:          pointer.To(appRole["enabled"].(bool)),
		}

		if v, ok := appRole["value"]; ok {
			newAppRole.Value = nullable.Value(v.(string))
		}

		result = append(result, newAppRole)
	}

	return &result
}

func expandApplicationGroupMembershipClaims(in []interface{}) nullable.Type[string] {
	if len(in) == 0 {
		return nullable.NoZero("")
	}

	ret := make([]string, 0)
	for _, claimRaw := range in {
		ret = append(ret, strings.TrimSpace(claimRaw.(string)))
	}

	return nullable.NoZero(strings.Join(ret, ","))
}

func expandApplicationImplicitGrantSettings(input []interface{}) *stable.ImplicitGrantSettings {
	var enableAccessTokenIssuance, enableIdTokenIssuance bool

	if len(input) > 0 && input[0] != nil {
		in := input[0].(map[string]interface{})
		enableAccessTokenIssuance = in["access_token_issuance_enabled"].(bool)
		enableIdTokenIssuance = in["id_token_issuance_enabled"].(bool)
	}

	return &stable.ImplicitGrantSettings{
		EnableAccessTokenIssuance: nullable.Value(enableAccessTokenIssuance),
		EnableIdTokenIssuance:     nullable.Value(enableIdTokenIssuance),
	}
}

func expandApplicationOAuth2PermissionScope(in []interface{}) *[]stable.PermissionScope {
	result := make([]stable.PermissionScope, 0)

	for _, raw := range in {
		if raw == nil {
			continue
		}
		oauth2Permissions := raw.(map[string]interface{})

		result = append(result, stable.PermissionScope{
			AdminConsentDescription: nullable.Value(oauth2Permissions["admin_consent_description"].(string)),
			AdminConsentDisplayName: nullable.Value(oauth2Permissions["admin_consent_display_name"].(string)),
			Id:                      pointer.To(oauth2Permissions["id"].(string)),
			IsEnabled:               pointer.To(oauth2Permissions["enabled"].(bool)),
			Type:                    nullable.Value(oauth2Permissions["type"].(string)),
			UserConsentDescription:  nullable.Value(oauth2Permissions["user_consent_description"].(string)),
			UserConsentDisplayName:  nullable.Value(oauth2Permissions["user_consent_display_name"].(string)),
			Value:                   nullable.Value(oauth2Permissions["value"].(string)),
		})
	}

	return &result
}

func expandApplicationOptionalClaims(in []interface{}) *stable.OptionalClaims {
	result := stable.OptionalClaims{}

	if len(in) == 0 || in[0] == nil {
		return &result
	}

	optionalClaims := in[0].(map[string]interface{})

	result.AccessToken = expandApplicationOptionalClaim(optionalClaims["access_token"].([]interface{}))
	result.IdToken = expandApplicationOptionalClaim(optionalClaims["id_token"].([]interface{}))
	result.Saml2Token = expandApplicationOptionalClaim(optionalClaims["saml2_token"].([]interface{}))

	return &result
}

func expandApplicationOptionalClaim(in []interface{}) *[]stable.OptionalClaim {
	result := make([]stable.OptionalClaim, 0)

	for _, optionalClaimRaw := range in {
		if optionalClaimRaw == nil {
			continue
		}
		optionalClaim := optionalClaimRaw.(map[string]interface{})

		additionalProps := make([]string, 0)
		if props, ok := optionalClaim["additional_properties"]; ok && props != nil {
			for _, prop := range props.([]interface{}) {
				additionalProps = append(additionalProps, prop.(string))
			}
		}

		newClaim := stable.OptionalClaim{
			Name:                 pointer.To(optionalClaim["name"].(string)),
			Essential:            pointer.To(optionalClaim["essential"].(bool)),
			AdditionalProperties: &additionalProps,
		}

		if source, ok := optionalClaim["source"].(string); ok && source != "" {
			newClaim.Source = nullable.Value(source)
		}

		result = append(result, newClaim)
	}

	return &result
}

func expandApplicationPublicClient(input []interface{}) (result *stable.PublicClientApplication) {
	result = &stable.PublicClientApplication{
		RedirectUris: &[]string{},
	}

	if len(input) == 0 || input[0] == nil {
		return
	}

	in := input[0].(map[string]interface{})
	result.RedirectUris = tf.ExpandStringSlicePtr(in["redirect_uris"].(*pluginsdk.Set).List())

	return
}

func expandApplicationRequiredResourceAccess(in []interface{}) *[]stable.RequiredResourceAccess {
	result := make([]stable.RequiredResourceAccess, 0)

	for _, raw := range in {
		if raw == nil {
			continue
		}
		requiredResourceAccess := raw.(map[string]interface{})

		result = append(result, stable.RequiredResourceAccess{
			ResourceAppId: pointer.To(requiredResourceAccess["resource_app_id"].(string)),
			ResourceAccess: expandApplicationResourceAccess(
				requiredResourceAccess["resource_access"].([]interface{}),
			),
		})
	}

	return &result
}

func expandApplicationResourceAccess(in []interface{}) *[]stable.ResourceAccess {
	result := make([]stable.ResourceAccess, 0)

	for _, resourceAccessRaw := range in {
		if resourceAccessRaw == nil {
			continue
		}
		resourceAccess := resourceAccessRaw.(map[string]interface{})

		result = append(result, stable.ResourceAccess{
			Id:   pointer.To(resourceAccess["id"].(string)),
			Type: nullable.Value(resourceAccess["type"].(string)),
		})
	}

	return &result
}

func expandApplicationSpa(input []interface{}) (result *stable.SpaApplication) {
	result = &stable.SpaApplication{
		RedirectUris: &[]string{},
	}

	if len(input) == 0 || input[0] == nil {
		return
	}

	in := input[0].(map[string]interface{})
	result.RedirectUris = tf.ExpandStringSlicePtr(in["redirect_uris"].(*pluginsdk.Set).List())

	return
}

func expandApplicationWeb(input []interface{}) (result *stable.WebApplication) {
	result = &stable.WebApplication{
		HomePageUrl:           nullable.NoZero(""),
		ImplicitGrantSettings: expandApplicationImplicitGrantSettings(nil),
		LogoutUrl:             nullable.NoZero(""),
		RedirectUris:          &[]string{},
	}

	if len(input) == 0 || input[0] == nil {
		return
	}

	in := input[0].(map[string]interface{})
	result.HomePageUrl = nullable.NoZero(in["homepage_url"].(string))
	result.LogoutUrl = nullable.NoZero(in["logout_url"].(string))
	result.ImplicitGrantSettings = expandApplicationImplicitGrantSettings(in["implicit_grant"].([]interface{}))
	result.RedirectUris = tf.ExpandStringSlicePtr(in["redirect_uris"].(*pluginsdk.Set).List())

	return
}

func flattenApplicationApi(in *stable.ApiApplication, dataSource bool) []map[string]interface{} {
	if in == nil {
		return []map[string]interface{}{}
	}

	mappedClaims := in.AcceptMappedClaims.GetOrZero()

	scopesKey := "oauth2_permission_scope"
	if dataSource {
		scopesKey = "oauth2_permission_scopes"
	}

	accessTokenVersion := 1
	if !in.RequestedAccessTokenVersion.IsNull() {
		accessTokenVersion = int(in.RequestedAccessTokenVersion.GetOrZero())
	}

	return []map[string]interface{}{{
		"known_client_applications":      tf.FlattenStringSlicePtr(in.KnownClientApplications),
		"mapped_claims_enabled":          mappedClaims,
		scopesKey:                        applications.FlattenOAuth2PermissionScopes(in.OAuth2PermissionScopes),
		"requested_access_token_version": accessTokenVersion,
	}}
}

func flattenApplicationGroupMembershipClaims(in nullable.Type[string]) []interface{} {
	if in.IsNull() {
		return []interface{}{}
	}

	ret := make([]interface{}, 0)
	for claim := range strings.SplitSeq(in.GetOrZero(), ",") {
		ret = append(ret, strings.TrimSpace(claim))
	}

	return ret
}

func flattenApplicationImplicitGrant(in *stable.ImplicitGrantSettings) []map[string]interface{} {
	if in == nil {
		return []map[string]interface{}{}
	}

	return []map[string]interface{}{{
		"access_token_issuance_enabled": in.EnableAccessTokenIssuance.GetOrZero(),
		"id_token_issuance_enabled":     in.EnableIdTokenIssuance.GetOrZero(),
	}}
}

func flattenApplicationOptionalClaims(in *stable.OptionalClaims) []map[string]interface{} {
	if in == nil {
		return []map[string]interface{}{}
	}

	return []map[string]interface{}{{
		"access_token": flattenApplicationOptionalClaim(in.AccessToken),
		"id_token":     flattenApplicationOptionalClaim(in.IdToken),
		"saml2_token":  flattenApplicationOptionalClaim(in.Saml2Token),
	}}
}

func flattenApplicationOptionalClaim(in *[]stable.OptionalClaim) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	optionalClaims := make([]interface{}, 0)
	for _, claim := range *in {
		optionalClaim := map[string]interface{}{
			"name":                  claim.Name,
			"essential":             claim.Essential,
			"source":                "",
			"additional_properties": []string{},
		}

		if !claim.Source.IsNull() {
			optionalClaim["source"] = claim.Source.GetOrZero()
		}

		if claim.AdditionalProperties != nil && len(*claim.AdditionalProperties) > 0 {
			optionalClaim["additional_properties"] = *claim.AdditionalProperties
		}

		optionalClaims = append(optionalClaims, optionalClaim)
	}

	return optionalClaims
}

func flattenApplicationPublicClient(in *stable.PublicClientApplication) []map[string]interface{} {
	if in == nil {
		return []map[string]interface{}{}
	}

	return []map[string]interface{}{{
		"redirect_uris": tf.FlattenStringSlicePtr(in.RedirectUris),
	}}
}

func flattenApplicationRequiredResourceAccess(in *[]stable.RequiredResourceAccess) []map[string]interface{} {
	if in == nil {
		return []map[string]interface{}{}
	}

	result := make([]map[string]interface{}, 0)
	for _, requiredResourceAccess := range *in {
		resourceAppId := ""
		if requiredResourceAccess.ResourceAppId != nil {
			resourceAppId = *requiredResourceAccess.ResourceAppId
		}

		result = append(result, map[string]interface{}{
			"resource_app_id": resourceAppId,
			"resource_access": flattenApplicationResourceAccess(requiredResourceAccess.ResourceAccess),
		})
	}

	return result
}

func flattenApplicationResourceAccess(in *[]stable.ResourceAccess) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	accesses := make([]interface{}, 0)
	for _, resourceAccess := range *in {
		access := make(map[string]interface{})
		if resourceAccess.Id != nil {
			access["id"] = *resourceAccess.Id
		}
		access["type"] = resourceAccess.Type.GetOrZero()
		accesses = append(accesses, access)
	}

	return accesses
}

func flattenApplicationSpa(in *stable.SpaApplication) []map[string]interface{} {
	if in == nil {
		return []map[string]interface{}{}
	}

	return []map[string]interface{}{{
		"redirect_uris": tf.FlattenStringSlicePtr(in.RedirectUris),
	}}
}

func flattenApplicationPasswordCredentials(input *[]stable.PasswordCredential) []map[string]interface{} {
	output := make([]map[string]interface{}, 0)

	if input == nil {
		return output
	}

	for _, in := range *input {
		output = append(output, map[string]interface{}{
			"key_id":       in.KeyId.GetOrZero(),
			"display_name": in.DisplayName.GetOrZero(),
			"start_date":   in.StartDateTime.GetOrZero(),
			"end_date":     in.EndDateTime.GetOrZero(),
			"value":        in.SecretText.GetOrZero(),
		})
	}

	return output
}

func flattenApplicationWeb(in *stable.WebApplication) []map[string]interface{} {
	if in == nil {
		return []map[string]interface{}{}
	}

	return []map[string]interface{}{{
		"homepage_url":   in.HomePageUrl.GetOrZero(),
		"logout_url":     in.LogoutUrl.GetOrZero(),
		"redirect_uris":  tf.FlattenStringSlicePtr(in.RedirectUris),
		"implicit_grant": flattenApplicationImplicitGrant(in.ImplicitGrantSettings),
	}}
}
