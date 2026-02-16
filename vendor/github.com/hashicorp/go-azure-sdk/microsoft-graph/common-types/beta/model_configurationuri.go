package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigurationUri struct {
	// The single sign-on mode that the URI is configured for. Possible values are: saml, password.
	AppliesToSingleSignOnMode *string `json:"appliesToSingleSignOnMode,omitempty"`

	// The various formats that the URI should follow.
	Examples *[]string `json:"examples,omitempty"`

	// Indicates whether this URI is required for the single sign-on configuration.
	IsRequired *bool `json:"isRequired,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Usage *UriUsageType `json:"usage,omitempty"`

	// The suggested values for the URI. Developers may need to customize these values for their tenant.
	Values *[]string `json:"values,omitempty"`
}
