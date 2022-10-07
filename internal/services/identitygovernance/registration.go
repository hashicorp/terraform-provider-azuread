package identitygovernance

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Registration struct{}

// Name is the name of this Service
func (r Registration) Name() string {
	return "Identity Governance"
}

// WebsiteCategories returns a list of categories which can be used for the sidebar
func (r Registration) WebsiteCategories() []string {
	return []string{
		"Identity Governance",
	}
}

// SupportedDataSources returns the supported Data Sources supported by this Service
func (r Registration) SupportedDataSources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"azuread_access_package_catalog": accessPackageCatalogDataSource(),
		"azuread_access_package":         accessPackageDataSource(),
	}
}

// SupportedResources returns the supported Resources supported by this Service
func (r Registration) SupportedResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"azuread_access_package_catalog":                      accessPackageCatalogResource(),
		"azuread_access_package":                              accessPackageResource(),
		"azuread_access_package_assignment_policy":            accessPackageAssignmentPolicyResource(),
		"azuread_access_package_resource_catalog_association": accessPackageResourceCatalogAssociationResource(),
		"azuread_access_package_resource_package_association": accessPackageResourcePackageAssociationResource(),
	}
}
