// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package customsecurityattributes

import (
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/pluginsdk"
)

type Registration struct{}

func (r Registration) Name() string {
	return "Custom Security Attributes"
}

func (r Registration) AssociatedGitHubLabel() string {
	return "feature/custom-security-attributes"
}

func (r Registration) WebsiteCategories() []string {
	return []string{"Custom Security Attributes"}
}

func (r Registration) SupportedDataSources() map[string]*pluginsdk.Resource {
	return map[string]*pluginsdk.Resource{}
}

func (r Registration) SupportedResources() map[string]*pluginsdk.Resource {
	return map[string]*pluginsdk.Resource{
		"azuread_custom_security_attribute_assignment": customSecurityAttributeAssignmentResource(),
	}
}
