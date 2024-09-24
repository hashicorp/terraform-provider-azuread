// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applications

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-helpers/lang/response"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/owner"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
)

func GetOwner(ctx context.Context, client *owner.OwnerClient, id stable.ApplicationIdOwnerId) (stable.DirectoryObject, error) {
	applicationId := stable.NewApplicationID(id.ApplicationId)

	options := owner.ListOwnersOperationOptions{
		Filter: pointer.To(fmt.Sprintf("id eq '%s'", id.DirectoryObjectId)),
	}

	resp, err := client.ListOwners(ctx, applicationId, options)
	if err != nil {
		if response.WasNotFound(resp.HttpResponse) {
			return nil, nil
		}
		return nil, fmt.Errorf("unable to list Owners with filter %q: %+v", *options.Filter, err)
	}

	if resp.Model != nil {
		for _, o := range *resp.Model {
			if o.DirectoryObject().Id != nil && strings.EqualFold(*o.DirectoryObject().Id, id.DirectoryObjectId) {
				return o, nil
			}
		}
	}

	return nil, nil
}

func ExpandFeatures(in []interface{}) []string {
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

func FlattenAppRoleIDs(in *[]stable.AppRole) map[string]string {
	result := make(map[string]string)
	if in != nil {
		for _, role := range *in {
			if value := role.Value.GetOrZero(); value != "" && role.Id != nil {
				result[value] = *role.Id
			}
		}
	}
	return result
}

func FlattenAppRoles(in *[]stable.AppRole) (result []map[string]interface{}) {
	if in == nil {
		return //nolint:nakedret
	}

	for _, role := range *in {
		roleId := ""
		if role.Id != nil {
			roleId = *role.Id
		}

		allowedMemberTypes := make([]interface{}, 0)
		if v := role.AllowedMemberTypes; v != nil {
			for _, m := range *v {
				allowedMemberTypes = append(allowedMemberTypes, m)
			}
		}

		enabled := false
		if role.IsEnabled != nil && *role.IsEnabled {
			enabled = true
		}

		result = append(result, map[string]interface{}{
			"id":                   roleId,
			"allowed_member_types": allowedMemberTypes,
			"description":          role.Description.GetOrZero(),
			"display_name":         role.DisplayName.GetOrZero(),
			"enabled":              enabled,
			"value":                role.Value.GetOrZero(),
		})
	}

	return //nolint:nakedret
}

func FlattenFeatures(tags *[]string, deprecated bool) []interface{} {
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
		switch {
		case strings.EqualFold(tag, "WindowsAzureActiveDirectoryCustomSingleSignOnApplication"):
			result["custom_single_sign_on"] = true
		case strings.EqualFold(tag, "WindowsAzureActiveDirectoryIntegratedApp"):
			result["enterprise"] = true
		case strings.EqualFold(tag, "WindowsAzureActiveDirectoryGalleryApplicationNonPrimaryV1"):
			result["gallery"] = true
		case strings.EqualFold(tag, "HideApp"):
			result["hide"] = true
		}
	}

	return []interface{}{result}
}

func FlattenOAuth2PermissionScopeIDs(in *[]stable.PermissionScope) map[string]string {
	result := make(map[string]string)
	if in != nil {
		for _, scope := range *in {
			if value := scope.Value.GetOrZero(); value != "" && scope.Id != nil {
				result[value] = *scope.Id
			}
		}
	}
	return result
}

func FlattenOAuth2PermissionScopes(in *[]stable.PermissionScope) (result []map[string]interface{}) {
	if in == nil {
		return //nolint:nakedret
	}

	for _, p := range *in {
		result = append(result, map[string]interface{}{
			"admin_consent_description":  p.AdminConsentDescription.GetOrZero(),
			"admin_consent_display_name": p.AdminConsentDisplayName.GetOrZero(),
			"id":                         pointer.From(p.Id),
			"enabled":                    pointer.From(p.IsEnabled),
			"type":                       p.Type.GetOrZero(),
			"user_consent_description":   p.UserConsentDescription.GetOrZero(),
			"user_consent_display_name":  p.UserConsentDisplayName.GetOrZero(),
			"value":                      p.Value.GetOrZero(),
		})
	}

	return //nolint:nakedret
}
