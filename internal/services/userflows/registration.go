package userflows

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Registration struct{}

// Name is the name of this Service
func (r Registration) Name() string {
	return "User Flows"
}

// WebsiteCategories returns a list of categories which can be used for the sidebar
func (r Registration) WebsiteCategories() []string {
	return []string{
		"User Flows",
	}
}

// SupportedDataSources returns the supported Data Sources supported by this Service
func (r Registration) SupportedDataSources() map[string]*schema.Resource {
	return nil
}

// SupportedResources returns the supported Resources supported by this Service
func (r Registration) SupportedResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"azuread_b2c_user_flow":       b2cUserflowResource(),
		"azuread_user_flow_attribute": userFlowAttributeResource(),
	}
}
