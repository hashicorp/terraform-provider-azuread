// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serviceprincipals

import (
	"github.com/hashicorp/terraform-provider-azuread/internal/sdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"
)

type Registration struct{}

// Name is the name of this Service
func (r Registration) Name() string {
	return "Service Principals"
}

// AssociatedGitHubLabel is the issue/PR label which can be applied to PRs that include changes to this service package
func (r Registration) AssociatedGitHubLabel() string {
	return "feature/service-principals"
}

// WebsiteCategories returns a list of categories which can be used for the sidebar
func (r Registration) WebsiteCategories() []string {
	return []string{
		"Service Principals",
	}
}

// SupportedDataSources returns the supported Data Sources supported by this Service
func (r Registration) SupportedDataSources() map[string]*pluginsdk.Resource {
	return map[string]*pluginsdk.Resource{
		"azuread_service_principal":  servicePrincipalData(),
		"azuread_service_principals": servicePrincipalsDataSource(),
	}
}

// SupportedResources returns the supported Resources supported by this Service
func (r Registration) SupportedResources() map[string]*pluginsdk.Resource {
	return map[string]*pluginsdk.Resource{
		"azuread_service_principal":                                  servicePrincipalResource(),
		"azuread_service_principal_certificate":                      servicePrincipalCertificateResource(),
		"azuread_service_principal_claims_mapping_policy_assignment": servicePrincipalClaimsMappingPolicyAssignmentResource(),
		"azuread_service_principal_delegated_permission_grant":       servicePrincipalDelegatedPermissionGrantResource(),
		"azuread_service_principal_password":                         servicePrincipalPasswordResource(),
		"azuread_service_principal_token_signing_certificate":        servicePrincipalTokenSigningCertificateResource(),
	}
}

// DataSources returns the typed DataSources supported by this service
func (r Registration) DataSources() []sdk.DataSource {
	return []sdk.DataSource{
		ClientConfigDataSource{},
	}
}

// Resources returns the typed Resources supported by this service
func (r Registration) Resources() []sdk.Resource {
	return []sdk.Resource{}
}
