package helpers

import (
	"strings"

	"github.com/manicminer/hamilton/msgraph"
)

func ApplicationExpandFeatures(in []interface{}) []string {
	out := make([]string, 0)

	if len(in) == 0 || in[0] == nil {
		return out
	}

	features := in[0].(map[string]interface{})

	if v, ok := features["custom_single_sign_on"]; ok && v.(bool) {
		out = append(out, "WindowsAzureActiveDirectoryCustomSingleSignOnApplication")
	} else if v, ok := features["custom_single_sign_on_app"]; ok && v.(bool) {
		out = append(out, "WindowsAzureActiveDirectoryCustomSingleSignOnApplication")
	}

	if v, ok := features["enterprise"]; ok && v.(bool) {
		out = append(out, "WindowsAzureActiveDirectoryIntegratedApp")
	} else if v, ok := features["enterprise_application"]; ok && v.(bool) { // TODO: remove in v3.0
		out = append(out, "WindowsAzureActiveDirectoryIntegratedApp")
	}

	if v, ok := features["gallery"]; ok && v.(bool) {
		out = append(out, "WindowsAzureActiveDirectoryGalleryApplicationNonPrimaryV1")
	} else if v, ok := features["gallery_application"]; ok && v.(bool) { // TODO: remove in v3.0
		out = append(out, "WindowsAzureActiveDirectoryGalleryApplicationNonPrimaryV1")
	}

	if v, ok := features["hide"]; ok && v.(bool) {
		out = append(out, "HideApp")
	} else if v, ok := features["visible_to_users"]; ok && !v.(bool) { // TODO: remove in v3.0
		out = append(out, "HideApp")
	}

	return out
}

func ApplicationFlattenAppRoleIDs(in *[]msgraph.AppRole) map[string]string {
	result := make(map[string]string)
	if in != nil {
		for _, role := range *in {
			if role.Value != nil && *role.Value != "" && role.ID != nil {
				result[*role.Value] = *role.ID
			}
		}
	}
	return result
}

func ApplicationFlattenAppRoles(in *[]msgraph.AppRole) (result []map[string]interface{}) {
	if in == nil {
		return
	}

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
		result = append(result, map[string]interface{}{
			"id":                   roleId,
			"allowed_member_types": allowedMemberTypes,
			"description":          description,
			"display_name":         displayName,
			"enabled":              enabled,
			"value":                value,
		})
	}

	return //nolint:nakedret
}

func ApplicationFlattenFeatures(tags *[]string, deprecated bool) []interface{} {
	// TODO: remove this in v3.0
	if deprecated {
		result := map[string]bool{
			"custom_single_sign_on_app": false,
			"enterprise_application":    false,
			"gallery_application":       false,
			"visible_to_users":          true,
		}

		if tags == nil || len(*tags) == 0 {
			return []interface{}{result}
		}

		for _, tag := range *tags {
			if strings.EqualFold(tag, "WindowsAzureActiveDirectoryCustomSingleSignOnApplication") {
				result["custom_single_sign_on_app"] = true
			}
			if strings.EqualFold(tag, "WindowsAzureActiveDirectoryIntegratedApp") {
				result["enterprise_application"] = true
			}
			if strings.EqualFold(tag, "WindowsAzureActiveDirectoryGalleryApplicationNonPrimaryV1") {
				result["gallery_application"] = true
			}
			if strings.EqualFold(tag, "HideApp") {
				result["visible_to_users"] = false
			}
		}

		return []interface{}{result}
	}

	result := map[string]bool{
		"custom_single_sign_on": false,
		"enterprise":            false,
		"gallery":               false,
		"hide":                  false,
	}

	if tags == nil || len(*tags) == 0 {
		return []interface{}{result}
	}

	for _, tag := range *tags {
		if strings.EqualFold(tag, "WindowsAzureActiveDirectoryCustomSingleSignOnApplication") {
			result["custom_single_sign_on"] = true
		}
		if strings.EqualFold(tag, "WindowsAzureActiveDirectoryIntegratedApp") {
			result["enterprise"] = true
		}
		if strings.EqualFold(tag, "WindowsAzureActiveDirectoryGalleryApplicationNonPrimaryV1") {
			result["gallery"] = true
		}
		if strings.EqualFold(tag, "HideApp") {
			result["hide"] = true
		}
	}

	return []interface{}{result}
}

func ApplicationFlattenOAuth2PermissionScopeIDs(in *[]msgraph.PermissionScope) map[string]string {
	result := make(map[string]string)
	if in != nil {
		for _, scope := range *in {
			if scope.Value != nil && *scope.Value != "" && scope.ID != nil {
				result[*scope.Value] = *scope.ID
			}
		}
	}
	return result
}

func ApplicationFlattenOAuth2PermissionScopes(in *[]msgraph.PermissionScope) (result []map[string]interface{}) {
	if in == nil {
		return
	}

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

	return //nolint:nakedret
}
