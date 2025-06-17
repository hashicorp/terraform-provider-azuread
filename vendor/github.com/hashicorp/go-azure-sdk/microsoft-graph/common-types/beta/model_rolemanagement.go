package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagement struct {
	CloudPC  *RbacApplicationMultiple `json:"cloudPC,omitempty"`
	Defender *RbacApplicationMultiple `json:"defender,omitempty"`

	// The RbacApplication for Device Management
	DeviceManagement *RbacApplicationMultiple `json:"deviceManagement,omitempty"`

	Directory      *RbacApplication   `json:"directory,omitempty"`
	EnterpriseApps *[]RbacApplication `json:"enterpriseApps,omitempty"`

	// The RbacApplication for Entitlement Management
	EntitlementManagement *RbacApplication `json:"entitlementManagement,omitempty"`

	Exchange *UnifiedRbacApplication `json:"exchange,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
