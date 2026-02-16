package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessDevice struct {
	// A unique device ID.
	DeviceId *string `json:"deviceId,omitempty"`

	// The display name for the device.
	DisplayName *string `json:"displayName,omitempty"`

	// A value that indicates whether or not the device is compliant.
	IsCompliant *bool `json:"isCompliant,omitempty"`

	// The most recent access time for the device.
	LastAccessDateTime *string `json:"lastAccessDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The operating system on the device.
	OperatingSystem *string `json:"operatingSystem,omitempty"`

	TrafficType *NetworkaccessTrafficType `json:"trafficType,omitempty"`
}
