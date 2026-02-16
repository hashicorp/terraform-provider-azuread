package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesApplicableContentDeviceMatch struct {
	// Collection of vendors who recommend the content.
	DeviceId *string `json:"deviceId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Collection of vendors who recommend the content.
	RecommendedBy *[]string `json:"recommendedBy,omitempty"`
}
