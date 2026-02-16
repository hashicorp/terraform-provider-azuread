package stable

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceExchangeAccessStateSummary struct {
	// Total count of devices with Exchange Access State: Allowed.
	AllowedDeviceCount *int64 `json:"allowedDeviceCount,omitempty"`

	// Total count of devices with Exchange Access State: Blocked.
	BlockedDeviceCount *int64 `json:"blockedDeviceCount,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Total count of devices with Exchange Access State: Quarantined.
	QuarantinedDeviceCount *int64 `json:"quarantinedDeviceCount,omitempty"`

	// Total count of devices for which no Exchange Access State could be found.
	UnavailableDeviceCount *int64 `json:"unavailableDeviceCount,omitempty"`

	// Total count of devices with Exchange Access State: Unknown.
	UnknownDeviceCount *int64 `json:"unknownDeviceCount,omitempty"`
}
