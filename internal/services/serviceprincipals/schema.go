// Copyright IBM Corp. 2019, 2026
// SPDX-License-Identifier: MPL-2.0

package serviceprincipals

import "github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"

func schemaAppRolesComputed() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Description: "For more information [official documentation](https://learn.microsoft.com/azure/architecture/multitenant-identity/app-roles).",
		Type:        pluginsdk.TypeList,
		Computed:    true,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"id": {
					Description: "The unique identifier of the app role.",
					Type:        pluginsdk.TypeString,
					Computed:    true,
				},

				"allowed_member_types": {
					Description: "Whether this app role definition can be assigned to users and groups, or to other applications (that are accessing this application in daemon service scenarios).",
					Type:        pluginsdk.TypeList,
					Computed:    true,
					Elem: &pluginsdk.Schema{
						Type: pluginsdk.TypeString,
					},
				},

				"description": {
					Description: "Permission help text that appears in the admin app assignment and consent experiences.",
					Type:        pluginsdk.TypeString,
					Computed:    true,
				},

				"display_name": {
					Description: "Display name for the permission that appears in the admin consent and app assignment experiences.",
					Type:        pluginsdk.TypeString,
					Computed:    true,
				},

				"enabled": {
					Description: "The unique identifier of the app role.",
					Type:        pluginsdk.TypeBool,
					Computed:    true,
				},

				"value": {
					Description: "The value of the roles claim that the application should expect in the authentication and access tokens.",
					Type:        pluginsdk.TypeString,
					Computed:    true,
				},
			},
		},
	}
}

func schemaOauth2PermissionScopesComputed() *pluginsdk.Schema {
	return &pluginsdk.Schema{
		Description: "A collection of OAuth 2.0 delegated permissions exposed by the associated application.",
		Type:        pluginsdk.TypeList,
		Computed:    true,
		Elem: &pluginsdk.Resource{
			Schema: map[string]*pluginsdk.Schema{
				"id": {
					Description: "The unique identifier of the delegated permission.",
					Type:        pluginsdk.TypeString,
					Computed:    true,
				},

				"admin_consent_description": {
					Description: "Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.",
					Type:        pluginsdk.TypeString,
					Computed:    true,
				},

				"admin_consent_display_name": {
					Description: "Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.",
					Type:        pluginsdk.TypeString,
					Computed:    true,
				},

				"enabled": {
					Description: "Whether the permission scope is enabled.",
					Type:        pluginsdk.TypeBool,
					Computed:    true,
				},

				"type": {
					Description: "Whether this delegated permission is considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator is required for consent to the permissions.",
					Type:        pluginsdk.TypeString,
					Computed:    true,
				},

				"user_consent_description": {
					Description: "Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.",
					Type:        pluginsdk.TypeString,
					Computed:    true,
				},

				"user_consent_display_name": {
					Description: "Display name for the delegated permission that appears in the end user consent experience.",
					Type:        pluginsdk.TypeString,
					Computed:    true,
				},

				"value": {
					Description: "The value that is used for the `scp` claim in OAuth 2.0 access tokens.",
					Type:        pluginsdk.TypeString,
					Computed:    true,
				},
			},
		},
	}
}
