package aadgraph

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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
		"azuread_application":       ApplicationData(),
		"azuread_domains":           DomainsData(),
		"azuread_client_config":     ClientConfigData(),
		"azuread_group":             GroupData(),
		"azuread_groups":            GroupsData(),
		"azuread_service_principal": ServicePrincipalData(),
		"azuread_user":              UserData(),
		"azuread_users":             UsersData(),
	}
}

// SupportedResources returns the supported Resources supported by this Service
func (r Registration) SupportedResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"azuread_application":                   ApplicationResource(),
		"azuread_application_certificate":       ApplicationCertificateResource(),
		"azuread_application_password":          ApplicationPasswordResource(),
		"azuread_group":                         GroupResource(),
		"azuread_group_member":                  GroupMemberResource(),
		"azuread_service_principal":             ServicePrincipalResource(),
		"azuread_service_principal_certificate": ServicePrincipalCertificateResource(),
		"azuread_service_principal_password":    ServicePrincipalPasswordResource(),
		"azuread_user":                          UserResource(),
	}
}
