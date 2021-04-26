package directory_roles

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Registration struct{}

// Name is the name of this Service
func (r Registration) Name() string {
	return "DirectoryRoles"
}

// WebsiteCategories returns a list of categories which can be used for the sidebar
func (r Registration) WebsiteCategories() []string {
	return []string{
		"DirectoryRoles",
	}
}

// SupportedDataSources returns the supported Data Sources supported by this Service
func (r Registration) SupportedDataSources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		//"azuread_directory_role":  directoryRoleDataSource(),
		//"azuread_directory_roles": directoryRolesDataSource(),
	}
}

// SupportedResources returns the supported Resources supported by this Service
func (r Registration) SupportedResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"azuread_directory_role":        directoryRoleResource(),
		"azuread_directory_role_member": directoryRoleMemberResource(),
	}
}
