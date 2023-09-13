// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package directoryroles

import "github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"

type Registration struct{}

// Name is the name of this Service
func (r Registration) Name() string {
	return "Directory Roles"
}

// WebsiteCategories returns a list of categories which can be used for the sidebar
func (r Registration) WebsiteCategories() []string {
	return []string{
		"Directory Roles",
	}
}

// SupportedDataSources returns the supported Data Sources supported by this Service
func (r Registration) SupportedDataSources() map[string]*pluginsdk.Resource {
	return map[string]*pluginsdk.Resource{
		"azuread_directory_roles":          directoryRolesDataSource(),
		"azuread_directory_role_templates": directoryRoleTemplatesDataSource(),
	}
}

// SupportedResources returns the supported Resources supported by this Service
func (r Registration) SupportedResources() map[string]*pluginsdk.Resource {
	return map[string]*pluginsdk.Resource{
		"azuread_custom_directory_role":     customDirectoryRoleResource(),
		"azuread_directory_role":            directoryRoleResource(),
		"azuread_directory_role_assignment": directoryRoleAssignmentResource(),
		"azuread_directory_role_member":     directoryRoleMemberResource(),
	}
}
