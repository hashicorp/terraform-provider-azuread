// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package conditionalaccess

import (
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"
)

type Registration struct{}

// Name is the name of this Service
func (r Registration) Name() string {
	return "Conditional Access"
}

// WebsiteCategories returns a list of categories which can be used for the sidebar
func (r Registration) WebsiteCategories() []string {
	return []string{
		"Conditional Access",
	}
}

// SupportedDataSources returns the supported Data Sources supported by this Service
func (r Registration) SupportedDataSources() map[string]*pluginsdk.Resource {
	return map[string]*pluginsdk.Resource{
		"azuread_named_location": namedLocationDataSource(),
	}
}

// SupportedResources returns the supported Resources supported by this Service
func (r Registration) SupportedResources() map[string]*pluginsdk.Resource {
	return map[string]*pluginsdk.Resource{
		"azuread_named_location":            namedLocationResource(),
		"azuread_conditional_access_policy": conditionalAccessPolicyResource(),
	}
}
