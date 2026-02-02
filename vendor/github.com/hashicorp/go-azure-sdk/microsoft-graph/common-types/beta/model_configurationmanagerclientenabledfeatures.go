package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigurationManagerClientEnabledFeatures struct {
	// Whether compliance policy is managed by Intune
	CompliancePolicy *bool `json:"compliancePolicy,omitempty"`

	// Whether device configuration is managed by Intune
	DeviceConfiguration *bool `json:"deviceConfiguration,omitempty"`

	// Whether Endpoint Protection is managed by Intune
	EndpointProtection *bool `json:"endpointProtection,omitempty"`

	// Whether inventory is managed by Intune
	Inventory *bool `json:"inventory,omitempty"`

	// Whether modern application is managed by Intune
	ModernApps *bool `json:"modernApps,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Whether Office application is managed by Intune
	OfficeApps *bool `json:"officeApps,omitempty"`

	// Whether resource access is managed by Intune
	ResourceAccess *bool `json:"resourceAccess,omitempty"`

	// Whether Windows Update for Business is managed by Intune
	WindowsUpdateForBusiness *bool `json:"windowsUpdateForBusiness,omitempty"`
}
