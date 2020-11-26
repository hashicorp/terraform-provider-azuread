package aadgraph

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Registration struct{}

// Name is the name of this Service
func (r Registration) Name() string {
	return "AAD Graph"
}

// WebsiteCategories returns a list of categories which can be used for the sidebar
func (r Registration) WebsiteCategories() []string {
	return []string{
		"AAD Graph",
	}
}

// SupportedDataSources returns the supported Data Sources supported by this Service
func (r Registration) SupportedDataSources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"azuread_application":       applicationData(),
		"azuread_domains":           domainsData(),
		"azuread_client_config":     clientConfigData(),
		"azuread_group":             groupData(),
		"azuread_groups":            groupsData(),
		"azuread_service_principal": servicePrincipalData(),
		"azuread_user":              userData(),
		"azuread_users":             usersData(),
	}
}

// SupportedResources returns the supported Resources supported by this Service
func (r Registration) SupportedResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"azuread_application":                   applicationResource(),
		"azuread_application_app_role":          applicationAppRoleResource(),
		"azuread_application_certificate":       applicationCertificateResource(),
		"azuread_application_oauth2_permission": applicationOAuth2PermissionResource(),
		"azuread_application_password":          applicationPasswordResource(),
		"azuread_group":                         groupResource(),
		"azuread_group_member":                  groupMemberResource(),
		"azuread_service_principal":             servicePrincipalResource(),
		"azuread_service_principal_certificate": servicePrincipalCertificateResource(),
		"azuread_service_principal_password":    servicePrincipalPasswordResource(),
		"azuread_user":                          userResource(),
	}
}
