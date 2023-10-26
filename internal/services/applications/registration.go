// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applications

import (
	"github.com/hashicorp/terraform-provider-azuread/internal/sdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"
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
func (r Registration) SupportedDataSources() map[string]*pluginsdk.Resource {
	return map[string]*pluginsdk.Resource{
		"azuread_application":                   applicationDataSource(),
		"azuread_application_published_app_ids": applicationPublishedAppIdsDataSource(),
		"azuread_application_template":          applicationTemplateDataSource(),
	}
}

// SupportedResources returns the supported Resources supported by this Service
func (r Registration) SupportedResources() map[string]*pluginsdk.Resource {
	return map[string]*pluginsdk.Resource{
		"azuread_application":                               applicationResource(),
		"azuread_application_certificate":                   applicationCertificateResource(),
		"azuread_application_federated_identity_credential": applicationFederatedIdentityCredentialResource(),
		"azuread_application_password":                      applicationPasswordResource(),
		"azuread_application_pre_authorized":                applicationPreAuthorizedResource(),
	}
}

// DataSources returns the typed DataSources supported by this service
func (r Registration) DataSources() []sdk.DataSource {
	return []sdk.DataSource{}
}

// Resources returns the typed Resources supported by this service
func (r Registration) Resources() []sdk.Resource {
	return []sdk.Resource{
		ApplicationApiAccessResource{},
		ApplicationAppRoleResource{},
		ApplicationFallbackPublicClientResource{},
		ApplicationFromTemplateResource{},
		ApplicationIdentifierUriResource{},
		ApplicationKnownClientsResource{},
		ApplicationOptionalClaimsResource{},
		ApplicationOwnerResource{},
		ApplicationPermissionScopeResource{},
		ApplicationRedirectUrisResource{},
		ApplicationRegistrationResource{},
	}
}
