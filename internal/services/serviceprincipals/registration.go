// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serviceprincipals

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Registration struct{}

// Name is the name of this Service
func (r Registration) Name() string {
	return "Service Principals"
}

// WebsiteCategories returns a list of categories which can be used for the sidebar
func (r Registration) WebsiteCategories() []string {
	return []string{
		"Service Principals",
	}
}

// SupportedDataSources returns the supported Data Sources supported by this Service
func (r Registration) SupportedDataSources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"azuread_client_config":      clientConfigDataSource(),
		"azuread_service_principal":  servicePrincipalData(),
		"azuread_service_principals": servicePrincipalsDataSource(),
	}
}

// SupportedResources returns the supported Resources supported by this Service
func (r Registration) SupportedResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"azuread_service_principal":                                  servicePrincipalResource(),
		"azuread_service_principal_certificate":                      servicePrincipalCertificateResource(),
		"azuread_service_principal_claims_mapping_policy_assignment": servicePrincipalClaimsMappingPolicyAssignmentResource(),
		"azuread_service_principal_delegated_permission_grant":       servicePrincipalDelegatedPermissionGrantResource(),
		"azuread_service_principal_password":                         servicePrincipalPasswordResource(),
		"azuread_service_principal_token_signing_certificate":        servicePrincipalTokenSigningCertificateResource(),
		"azuread_synchronization_job":                                synchronizationJobResource(),
		"azuread_synchronization_secret":                             synchronizationSecretResource(),
	}
}
