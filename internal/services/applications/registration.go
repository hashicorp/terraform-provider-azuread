package applications

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Registration struct{}

// Name is the name of this Service
func (r Registration) Name() string {
	return "Applications"
}

// WebsiteCategories returns a list of categories which can be used for the sidebar
func (r Registration) WebsiteCategories() []string {
	return []string{
		"Applications",
	}
}

// SupportedDataSources returns the supported Data Sources supported by this Service
func (r Registration) SupportedDataSources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"azuread_application": applicationDataSource(),
	}
}

// SupportedResources returns the supported Resources supported by this Service
func (r Registration) SupportedResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"azuread_application":                         applicationResource(),
		"azuread_application_app_role":                applicationAppRoleResource(),
		"azuread_application_certificate":             applicationCertificateResource(),
		"azuread_application_oauth2_permission":       applicationOAuth2PermissionResource(), // TODO: v2.0 remove this resource
		"azuread_application_oauth2_permission_scope": applicationOAuth2PermissionScopeResource(),
		"azuread_application_password":                applicationPasswordResource(),
	}
}
