package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CompanyPortalBlockedAction struct {
	// Action on a device that can be executed in the Company Portal
	Action *CompanyPortalAction `json:"action,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Owner type of device.
	OwnerType *OwnerType `json:"ownerType,omitempty"`

	// Supported platform types.
	Platform *DevicePlatformType `json:"platform,omitempty"`
}
