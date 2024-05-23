// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applications

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"
	"github.com/manicminer/hamilton/msgraph"
)

const applicationResourceName = "azuread_application"

func applicationAppRoleChanged(existing msgraph.AppRole, new msgraph.AppRole) bool {
	if !reflect.DeepEqual(existing.AllowedMemberTypes, new.AllowedMemberTypes) {
		return true
	}
	if !reflect.DeepEqual(existing.Description, new.Description) {
		return true
	}
	if !reflect.DeepEqual(existing.DisplayName, new.DisplayName) {
		return true
	}

	// The following order is important; we must check for nil, and we consider nil and "" to be equivalent Values
	if reflect.DeepEqual(existing.Value, new.Value) {
		return false
	}
	if existing.Value == nil && new.Value != nil && *new.Value == "" {
		return false
	}
	if existing.Value != nil && *existing.Value == "" && new.Value == nil {
		return false
	}

	return true
}

func applicationDisableAppRoles(ctx context.Context, client *msgraph.ApplicationsClient, application *msgraph.Application, newRoles *[]msgraph.AppRole) error {
	if application.ID() == nil {
		return fmt.Errorf("cannot use Application model with nil ID")
	}

	if newRoles == nil {
		newRoles = &[]msgraph.AppRole{}
	}

	app, status, err := client.Get(ctx, *application.ID(), odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return fmt.Errorf("application with ID %q was not found", *application.ID())
		}

		return fmt.Errorf("retrieving Application with object ID %q: %+v", *application.ID(), err)
	}

	var existingRoles []msgraph.AppRole
	if app.AppRoles != nil {
		existingRoles = *app.AppRoles
	}

	// Shortcut: don't update if no changes to be made
	if reflect.DeepEqual(existingRoles, *newRoles) {
		return nil
	}

	// Identify any roles to be changed
	var disable bool
	for _, new := range *newRoles {
		if new.ID == nil || *new.ID == "" {
			return fmt.Errorf("new role provided with nil or empty ID")
		}
		for i, existing := range existingRoles {
			if existing.ID != nil && *existing.ID == *new.ID {
				if existing.IsEnabled != nil && *existing.IsEnabled && applicationAppRoleChanged(existing, new) {
					*existingRoles[i].IsEnabled = false
					disable = true
				}
				break
			}
		}
	}

	// Identify any roles to be removed
	for i, existing := range existingRoles {
		found := false
		for _, new := range *newRoles {
			if existing.ID != nil && *new.ID == *existing.ID {
				found = true
				break
			}
		}
		if !found {
			*existingRoles[i].IsEnabled = false
			disable = true
		}
	}

	if disable {
		// Disable any changed or removed roles
		properties := msgraph.Application{
			DirectoryObject: msgraph.DirectoryObject{
				Id: application.ID(),
			},
			AppRoles: &existingRoles,
		}
		if _, err := client.Update(ctx, properties); err != nil {
			return fmt.Errorf("disabling App Roles for Application with object ID %q: %+v", *application.ID(), err)
		}

		// Wait for application manifest to reflect the disabled roles
		deadline, ok := ctx.Deadline()
		if !ok {
			return fmt.Errorf("context has no deadline")
		}
		timeout := time.Until(deadline)
		_, err = (&pluginsdk.StateChangeConf{ //nolint:staticcheck
			Pending:    []string{"Waiting"},
			Target:     []string{"Disabled"},
			Timeout:    timeout,
			MinTimeout: 1 * time.Second,
			Refresh: func() (interface{}, string, error) {
				app, _, err := client.Get(ctx, *application.ID(), odata.Query{})
				if err != nil {
					return nil, "Error", fmt.Errorf("retrieving Application with object ID %q: %+v", *application.ID(), err)
				}
				if app == nil || app.AppRoles == nil {
					return nil, "Error", fmt.Errorf("reading roles for Application with object ID %q: %+v", *application.ID(), err)
				}
				actualRoles := *app.AppRoles
				for _, expectedRole := range existingRoles {
					if expectedRole.IsEnabled != nil && !*expectedRole.IsEnabled {
						for _, actualRole := range actualRoles {
							if expectedRole.ID != nil && actualRole.ID != nil && *expectedRole.ID == *actualRole.ID {
								if actualRole.IsEnabled != nil && *actualRole.IsEnabled {
									return actualRoles, "Waiting", nil
								}
								break
							}
						}
					}
				}
				return actualRoles, "Disabled", nil
			},
		}).WaitForStateContext(ctx)
		if err != nil {
			return fmt.Errorf("waiting for App Roles to be disabled for Application with object ID %q: %+v", *application.ID(), err)
		}
	}

	return nil
}

func applicationDisableOauth2PermissionScopes(ctx context.Context, client *msgraph.ApplicationsClient, application *msgraph.Application, newScopes *[]msgraph.PermissionScope) error {
	if application.ID() == nil {
		return fmt.Errorf("cannot use Application model with nil ID")
	}

	if newScopes == nil {
		newScopes = &[]msgraph.PermissionScope{}
	}

	app, status, err := client.Get(ctx, *application.ID(), odata.Query{})
	if err != nil {
		if status == http.StatusNotFound {
			return fmt.Errorf("application with ID %q was not found", *application.ID())
		}

		return fmt.Errorf("retrieving Application with object ID %q: %+v", *application.ID(), err)
	}

	var existingScopes []msgraph.PermissionScope
	if app.Api != nil && app.Api.OAuth2PermissionScopes != nil {
		existingScopes = *app.Api.OAuth2PermissionScopes
	}

	// Don't update if no changes to be made
	if reflect.DeepEqual(existingScopes, *newScopes) {
		return nil
	}

	// Identify any scopes to be changed
	var disable bool
	for _, new := range *newScopes {
		if new.ID == nil || *new.ID == "" {
			return fmt.Errorf("new scope provided with nil or empty ID")
		}
		for i, existing := range existingScopes {
			if existing.ID != nil && *existing.ID == *new.ID {
				if existing.IsEnabled != nil && *existing.IsEnabled && !reflect.DeepEqual(existing, new) {
					*existingScopes[i].IsEnabled = false
					disable = true
				}
				break
			}
		}
	}

	// Identify any scopes to be removed
	for i, existing := range existingScopes {
		found := false
		for _, new := range *newScopes {
			if existing.ID != nil && *new.ID == *existing.ID {
				found = true
				break
			}
		}
		if !found {
			*existingScopes[i].IsEnabled = false
			disable = true
		}
	}

	if disable {
		// Disable any changed or removed scopes
		properties := msgraph.Application{
			DirectoryObject: msgraph.DirectoryObject{
				Id: application.ID(),
			},
			Api: &msgraph.ApplicationApi{
				OAuth2PermissionScopes: &existingScopes,
			},
		}
		if _, err := client.Update(ctx, properties); err != nil {
			return fmt.Errorf("disabling OAuth2 Permission Scopes for Application with object ID %q: %+v", *application.ID(), err)
		}

		// Wait for application manifest to reflect the disabled scopes
		deadline, ok := ctx.Deadline()
		if !ok {
			return fmt.Errorf("context has no deadline")
		}
		timeout := time.Until(deadline)
		_, err = (&pluginsdk.StateChangeConf{ //nolint:staticcheck
			Pending:    []string{"Waiting"},
			Target:     []string{"Disabled"},
			Timeout:    timeout,
			MinTimeout: 1 * time.Second,
			Refresh: func() (interface{}, string, error) {
				app, _, err := client.Get(ctx, *application.ID(), odata.Query{})
				if err != nil {
					return nil, "Error", fmt.Errorf("retrieving Application with object ID %q: %+v", *application.ID(), err)
				}
				if app == nil || app.Api == nil || app.Api.OAuth2PermissionScopes == nil {
					return nil, "Error", fmt.Errorf("reading scopes for Application with object ID %q: %+v", *application.ID(), err)
				}
				actualScopes := *app.Api.OAuth2PermissionScopes
				for _, expectedScope := range existingScopes {
					if expectedScope.IsEnabled != nil && !*expectedScope.IsEnabled {
						for _, actualScope := range actualScopes {
							if expectedScope.ID != nil && actualScope.ID != nil && *expectedScope.ID == *actualScope.ID {
								if actualScope.IsEnabled != nil && *actualScope.IsEnabled {
									return actualScopes, "Waiting", nil
								}
								break
							}
						}
					}
				}
				return actualScopes, "Disabled", nil
			},
		}).WaitForStateContext(ctx)
		if err != nil {
			return fmt.Errorf("waiting for OAuth2 Permission Scopes to be disabled for Application with object ID %q: %+v", *application.ID(), err)
		}
	}

	return nil
}

func applicationFindByName(ctx context.Context, client *msgraph.ApplicationsClient, displayName string) (*[]msgraph.Application, error) {
	query := odata.Query{
		Filter: fmt.Sprintf("displayName eq '%s'", displayName),
	}
	apps, _, err := client.List(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("unable to list Applications with filter %q: %+v", query.Filter, err)
	}

	result := make([]msgraph.Application, 0)
	if apps != nil {
		for _, app := range *apps {
			if app.DisplayName != nil && *app.DisplayName == displayName {
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

func expandApplicationApi(input []interface{}) (result *msgraph.ApplicationApi) {
	result = &msgraph.ApplicationApi{
		AcceptMappedClaims:          pointer.To(false),
		KnownClientApplications:     &[]string{},
		OAuth2PermissionScopes:      &[]msgraph.PermissionScope{},
		RequestedAccessTokenVersion: pointer.To(int32(1)),
	}

	if len(input) == 0 || input[0] == nil {
		return
	}

	in := input[0].(map[string]interface{})
	if v, ok := in["mapped_claims_enabled"]; ok {
		result.AcceptMappedClaims = pointer.To(v.(bool))
	}
	if v, ok := in["known_client_applications"]; ok {
		result.KnownClientApplications = tf.ExpandStringSlicePtr(v.(*pluginsdk.Set).List())
	}
	result.OAuth2PermissionScopes = expandApplicationOAuth2PermissionScope(in["oauth2_permission_scope"].(*pluginsdk.Set).List())
	if v, ok := in["requested_access_token_version"]; ok {
		result.RequestedAccessTokenVersion = pointer.To(int32(v.(int)))
	}

	return
}

type CredentialError struct {
	str  string
	attr string
}

func (e CredentialError) Attr() string {
	return e.attr
}

func (e CredentialError) Error() string {
	return e.str
}

func expandApplicationPasswordCredentials(input []interface{}) (*[]msgraph.PasswordCredential, error) {
	result := make([]msgraph.PasswordCredential, 0)

	if len(input) == 0 {
		return &result, nil
	}

	for _, appPasswordRaw := range input {
		if appPasswordRaw == nil {
			continue
		}

		appPassword := appPasswordRaw.(map[string]interface{})

		credential := msgraph.PasswordCredential{
			DisplayName: pointer.To(appPassword["display_name"].(string)),
		}

		if v, ok := appPassword["start_date"]; ok && v.(string) != "" {
			startDate, err := time.Parse(time.RFC3339, v.(string))
			if err != nil {
				return nil, CredentialError{str: fmt.Sprintf("Unable to parse the provided start date %q: %+v", v, err), attr: "start_date"}
			}
			credential.StartDateTime = &startDate
		}

		var endDate *time.Time
		if v, ok := appPassword["end_date"]; ok && v.(string) != "" {
			var err error
			expiry, err := time.Parse(time.RFC3339, v.(string))
			if err != nil {
				return nil, CredentialError{str: fmt.Sprintf("Unable to parse the provided end date %q: %+v", v, err), attr: "end_date"}
			}
			endDate = &expiry
		} else if v, ok := appPassword["end_date_relative"]; ok && v.(string) != "" {
			d, err := time.ParseDuration(v.(string))
			if err != nil {
				return nil, CredentialError{str: fmt.Sprintf("Unable to parse `end_date_relative` (%q) as a duration", v), attr: "end_date_relative"}
			}

			if credential.StartDateTime == nil {
				expiry := time.Now().Add(d)
				endDate = &expiry
			} else {
				expiry := credential.StartDateTime.Add(d)
				endDate = &expiry
			}
		}
		if endDate != nil {
			credential.EndDateTime = endDate
		}

		result = append(result, credential)
	}

	return &result, nil

}

func expandApplicationAppRoles(input []interface{}) *[]msgraph.AppRole {
	result := make([]msgraph.AppRole, 0)

	if len(input) == 0 {
		return &result
	}

	for _, appRoleRaw := range input {
		if appRoleRaw == nil {
			continue
		}
		appRole := appRoleRaw.(map[string]interface{})

		var allowedMemberTypes []msgraph.AppRoleAllowedMemberType
		for _, allowedMemberType := range appRole["allowed_member_types"].(*pluginsdk.Set).List() {
			allowedMemberTypes = append(allowedMemberTypes, allowedMemberType.(string))
		}

		newAppRole := msgraph.AppRole{
			ID:                 pointer.To(appRole["id"].(string)),
			AllowedMemberTypes: &allowedMemberTypes,
			Description:        pointer.To(appRole["description"].(string)),
			DisplayName:        pointer.To(appRole["display_name"].(string)),
			IsEnabled:          pointer.To(appRole["enabled"].(bool)),
		}

		if v, ok := appRole["value"]; ok {
			newAppRole.Value = pointer.To(v.(string))
		}

		result = append(result, newAppRole)
	}

	return &result
}

func expandApplicationGroupMembershipClaims(in []interface{}) *[]msgraph.GroupMembershipClaim {
	result := make([]msgraph.GroupMembershipClaim, 0)
	if len(in) == 0 {
		return &result
	}
	for _, claimRaw := range in {
		result = append(result, claimRaw.(string))
	}
	return &result
}

func expandApplicationImplicitGrantSettings(input []interface{}) *msgraph.ImplicitGrantSettings {
	var enableAccessTokenIssuance, enableIdTokenIssuance bool

	if len(input) > 0 && input[0] != nil {
		in := input[0].(map[string]interface{})
		enableAccessTokenIssuance = in["access_token_issuance_enabled"].(bool)
		enableIdTokenIssuance = in["id_token_issuance_enabled"].(bool)
	}

	return &msgraph.ImplicitGrantSettings{
		EnableAccessTokenIssuance: pointer.To(enableAccessTokenIssuance),
		EnableIdTokenIssuance:     pointer.To(enableIdTokenIssuance),
	}
}

func expandApplicationOAuth2PermissionScope(in []interface{}) *[]msgraph.PermissionScope {
	result := make([]msgraph.PermissionScope, 0)

	for _, raw := range in {
		if raw == nil {
			continue
		}
		oauth2Permissions := raw.(map[string]interface{})

		result = append(result,
			msgraph.PermissionScope{
				AdminConsentDescription: pointer.To(oauth2Permissions["admin_consent_description"].(string)),
				AdminConsentDisplayName: pointer.To(oauth2Permissions["admin_consent_display_name"].(string)),
				ID:                      pointer.To(oauth2Permissions["id"].(string)),
				IsEnabled:               pointer.To(oauth2Permissions["enabled"].(bool)),
				Type:                    oauth2Permissions["type"].(string),
				UserConsentDescription:  pointer.To(oauth2Permissions["user_consent_description"].(string)),
				UserConsentDisplayName:  pointer.To(oauth2Permissions["user_consent_display_name"].(string)),
				Value:                   pointer.To(oauth2Permissions["value"].(string)),
			},
		)
	}

	return &result
}

func expandApplicationOptionalClaims(in []interface{}) *msgraph.OptionalClaims {
	result := msgraph.OptionalClaims{}

	if len(in) == 0 || in[0] == nil {
		return &result
	}

	optionalClaims := in[0].(map[string]interface{})

	result.AccessToken = expandApplicationOptionalClaim(optionalClaims["access_token"].([]interface{}))
	result.IdToken = expandApplicationOptionalClaim(optionalClaims["id_token"].([]interface{}))
	result.Saml2Token = expandApplicationOptionalClaim(optionalClaims["saml2_token"].([]interface{}))

	return &result
}

func expandApplicationOptionalClaim(in []interface{}) *[]msgraph.OptionalClaim {
	result := make([]msgraph.OptionalClaim, 0)

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

		newClaim := msgraph.OptionalClaim{
			Name:                 pointer.To(optionalClaim["name"].(string)),
			Essential:            pointer.To(optionalClaim["essential"].(bool)),
			AdditionalProperties: &additionalProps,
		}

		if source, ok := optionalClaim["source"].(string); ok && source != "" {
			newClaim.Source = &source
		}

		result = append(result, newClaim)
	}

	return &result
}

func expandApplicationPublicClient(input []interface{}) (result *msgraph.PublicClient) {
	result = &msgraph.PublicClient{
		RedirectUris: &[]string{},
	}

	if len(input) == 0 || input[0] == nil {
		return
	}

	in := input[0].(map[string]interface{})
	result.RedirectUris = tf.ExpandStringSlicePtr(in["redirect_uris"].(*pluginsdk.Set).List())

	return
}

func expandApplicationRequiredResourceAccess(in []interface{}) *[]msgraph.RequiredResourceAccess {
	result := make([]msgraph.RequiredResourceAccess, 0)

	for _, raw := range in {
		if raw == nil {
			continue
		}
		requiredResourceAccess := raw.(map[string]interface{})

		result = append(result, msgraph.RequiredResourceAccess{
			ResourceAppId: pointer.To(requiredResourceAccess["resource_app_id"].(string)),
			ResourceAccess: expandApplicationResourceAccess(
				requiredResourceAccess["resource_access"].([]interface{}),
			),
		})
	}

	return &result
}

func expandApplicationResourceAccess(in []interface{}) *[]msgraph.ResourceAccess {
	result := make([]msgraph.ResourceAccess, 0)

	for _, resourceAccessRaw := range in {
		if resourceAccessRaw == nil {
			continue
		}
		resourceAccess := resourceAccessRaw.(map[string]interface{})

		result = append(result, msgraph.ResourceAccess{
			ID:   pointer.To(resourceAccess["id"].(string)),
			Type: resourceAccess["type"].(string),
		})
	}

	return &result
}

func expandApplicationSpa(input []interface{}) (result *msgraph.ApplicationSpa) {
	result = &msgraph.ApplicationSpa{
		RedirectUris: &[]string{},
	}

	if len(input) == 0 || input[0] == nil {
		return
	}

	in := input[0].(map[string]interface{})
	result.RedirectUris = tf.ExpandStringSlicePtr(in["redirect_uris"].(*pluginsdk.Set).List())

	return
}

func expandApplicationWeb(input []interface{}) (result *msgraph.ApplicationWeb) {
	result = &msgraph.ApplicationWeb{
		HomePageUrl:           tf.NullableString(""),
		ImplicitGrantSettings: expandApplicationImplicitGrantSettings(nil),
		LogoutUrl:             tf.NullableString(""),
		RedirectUris:          &[]string{},
	}

	if len(input) == 0 || input[0] == nil {
		return
	}

	in := input[0].(map[string]interface{})
	result.HomePageUrl = tf.NullableString(in["homepage_url"].(string))
	result.LogoutUrl = tf.NullableString(in["logout_url"].(string))
	result.ImplicitGrantSettings = expandApplicationImplicitGrantSettings(in["implicit_grant"].([]interface{}))
	result.RedirectUris = tf.ExpandStringSlicePtr(in["redirect_uris"].(*pluginsdk.Set).List())

	return
}

func flattenApplicationApi(in *msgraph.ApplicationApi, dataSource bool) []map[string]interface{} {
	if in == nil {
		return []map[string]interface{}{}
	}

	mappedClaims := false
	if in.AcceptMappedClaims != nil {
		mappedClaims = *in.AcceptMappedClaims
	}

	scopesKey := "oauth2_permission_scope"
	if dataSource {
		scopesKey = "oauth2_permission_scopes"
	}

	accessTokenVersion := 1
	if in.RequestedAccessTokenVersion != nil {
		accessTokenVersion = int(*in.RequestedAccessTokenVersion)
	}

	return []map[string]interface{}{{
		"known_client_applications":      tf.FlattenStringSlicePtr(in.KnownClientApplications),
		"mapped_claims_enabled":          mappedClaims,
		scopesKey:                        flattenApplicationOAuth2PermissionScopes(in.OAuth2PermissionScopes),
		"requested_access_token_version": accessTokenVersion,
	}}
}

func flattenApplicationAppRoleIDs(in *[]msgraph.AppRole) map[string]string {
	return helpers.ApplicationFlattenAppRoleIDs(in)
}

func flattenApplicationAppRoles(in *[]msgraph.AppRole) []map[string]interface{} {
	return helpers.ApplicationFlattenAppRoles(in)
}

func flattenApplicationImplicitGrant(in *msgraph.ImplicitGrantSettings) []map[string]interface{} {
	if in == nil {
		return []map[string]interface{}{}
	}

	accessToken := false
	if in.EnableAccessTokenIssuance != nil {
		accessToken = *in.EnableAccessTokenIssuance
	}
	idToken := false
	if in.EnableIdTokenIssuance != nil {
		idToken = *in.EnableIdTokenIssuance
	}

	return []map[string]interface{}{{
		"access_token_issuance_enabled": accessToken,
		"id_token_issuance_enabled":     idToken,
	}}
}

func flattenApplicationOAuth2PermissionScopeIDs(in *[]msgraph.PermissionScope) map[string]string {
	return helpers.ApplicationFlattenOAuth2PermissionScopeIDs(in)
}

func flattenApplicationOAuth2PermissionScopes(in *[]msgraph.PermissionScope) []map[string]interface{} {
	return helpers.ApplicationFlattenOAuth2PermissionScopes(in)
}

func flattenApplicationOptionalClaims(in *msgraph.OptionalClaims) []map[string]interface{} {
	if in == nil {
		return []map[string]interface{}{}
	}

	return []map[string]interface{}{{
		"access_token": flattenApplicationOptionalClaim(in.AccessToken),
		"id_token":     flattenApplicationOptionalClaim(in.IdToken),
		"saml2_token":  flattenApplicationOptionalClaim(in.Saml2Token),
	}}
}

func flattenApplicationOptionalClaim(in *[]msgraph.OptionalClaim) []interface{} {
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

		if claim.Source != nil {
			optionalClaim["source"] = *claim.Source
		}

		if claim.AdditionalProperties != nil && len(*claim.AdditionalProperties) > 0 {
			optionalClaim["additional_properties"] = *claim.AdditionalProperties
		}

		optionalClaims = append(optionalClaims, optionalClaim)
	}

	return optionalClaims
}

func flattenApplicationPublicClient(in *msgraph.PublicClient) []map[string]interface{} {
	if in == nil {
		return []map[string]interface{}{}
	}

	return []map[string]interface{}{{
		"redirect_uris": tf.FlattenStringSlicePtr(in.RedirectUris),
	}}
}

func flattenApplicationRequiredResourceAccess(in *[]msgraph.RequiredResourceAccess) []map[string]interface{} {
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

func flattenApplicationResourceAccess(in *[]msgraph.ResourceAccess) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	accesses := make([]interface{}, 0)
	for _, resourceAccess := range *in {
		access := make(map[string]interface{})
		if resourceAccess.ID != nil {
			access["id"] = *resourceAccess.ID
		}
		access["type"] = resourceAccess.Type
		accesses = append(accesses, access)
	}

	return accesses
}

func flattenApplicationSpa(in *msgraph.ApplicationSpa) []map[string]interface{} {
	if in == nil {
		return []map[string]interface{}{}
	}

	return []map[string]interface{}{{
		"redirect_uris": tf.FlattenStringSlicePtr(in.RedirectUris),
	}}
}

func flattenApplicationPasswordCredentials(in *[]msgraph.PasswordCredential, d []interface{}) []map[string]interface{} {
	result := []map[string]interface{}{}
	if in == nil {
		return result
	}

	for _, credential := range *in {
		addKey := true

		// skip credential, display name is required for inline app_password
		if credential.DisplayName == nil {
			continue
		}

		keyId := ""
		if credential.KeyId != nil {
			keyId = *credential.KeyId
		}

		startDate := ""
		if v := credential.StartDateTime; v != nil {
			startDate = v.Format(time.RFC3339)
		}

		endDate := ""
		if v := credential.EndDateTime; v != nil {
			endDate = v.Format(time.RFC3339)
		}

		credentialData := map[string]interface{}{
			"display_name": credential.DisplayName,
			"key_id":       &keyId,
			"start_date":   &startDate,
			"end_date":     &endDate,
		}

		if credential.SecretText != nil {
			credentialData["value"] = credential.SecretText
		}

		for _, appPasswordState := range d {
			if appPasswordState == nil {
				continue
			}

			appPassword := appPasswordState.(map[string]interface{})

			if v, ok := appPassword["value"]; ok {
				if credentialData["value"] == nil {
					credentialData["value"] = v.(string)
				}
			}

			if v, ok := appPassword["rotate_when_changed"]; ok {
				if d, ok := v.(map[string]interface{}); ok {
					if r, ok := d["rotation"].(string); ok {
						if len(r) > 0 {
							credentialData["rotate_when_changed"] = appPassword["rotate_when_changed"]
						}
					}
				}
			}

			// compare state key with secret that is readed.
			if appPassword["display_name"].(string) != *credential.DisplayName {
				addKey = false
			}
		}

		if addKey {
			result = append(result, credentialData)
		}
	}

	return result
}

func flattenApplicationWeb(in *msgraph.ApplicationWeb) []map[string]interface{} {
	if in == nil {
		return []map[string]interface{}{}
	}

	homepageUrl := ""
	if in.HomePageUrl != nil {
		homepageUrl = string(*in.HomePageUrl)
	}
	logoutUrl := ""
	if in.LogoutUrl != nil {
		logoutUrl = string(*in.LogoutUrl)
	}

	return []map[string]interface{}{{
		"homepage_url":   homepageUrl,
		"logout_url":     logoutUrl,
		"redirect_uris":  tf.FlattenStringSlicePtr(in.RedirectUris),
		"implicit_grant": flattenApplicationImplicitGrant(in.ImplicitGrantSettings),
	}}
}
