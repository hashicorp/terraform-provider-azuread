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
		"azuread_application":       DataApplication(),
		"azuread_domains":           DataDomains(),
		"azuread_client_config":     DataClientConfig(),
		"azuread_group":             DataGroup(),
		"azuread_groups":            DataGroups(),
		"azuread_service_principal": DataServicePrincipal(),
		"azuread_user":              DataUser(),
		"azuread_users":             DataUsers(),
	}
}

// SupportedResources returns the supported Resources supported by this Service
func (r Registration) SupportedResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"azuread_application":                   ResourceApplication(),
		"azuread_application_certificate":       ResourceApplicationCertificate(),
		"azuread_application_password":          ResourceApplicationPassword(),
		"azuread_group":                         ResourceGroup(),
		"azuread_group_member":                  ResourceGroupMember(),
		"azuread_service_principal":             ResourceServicePrincipal(),
		"azuread_service_principal_certificate": ResourceServicePrincipalCertificate(),
		"azuread_service_principal_password":    ResourceServicePrincipalPassword(),
		"azuread_user":                          ResourceUser(),
	}
}
