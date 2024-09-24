package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessDeviceUsageSummary struct {
	// The number of distinct device IDs between the discovery pivot time and the end of the reporting period.
	ActiveDeviceCount *int64 `json:"activeDeviceCount,omitempty"`

	// The discovery pivot time and the end of the reporting period, but were seen between the start of the reporting period
	// and the discovery pivot time.
	InactiveDeviceCount *int64 `json:"inactiveDeviceCount,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The total number of distinct device IDs that were seen during the reporting period.
	TotalDeviceCount *int64 `json:"totalDeviceCount,omitempty"`
}
