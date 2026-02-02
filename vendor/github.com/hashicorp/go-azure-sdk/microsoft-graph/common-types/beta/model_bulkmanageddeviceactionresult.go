package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BulkManagedDeviceActionResult struct {
	// Failed devices
	FailedDeviceIds *[]string `json:"failedDeviceIds,omitempty"`

	// Not found devices
	NotFoundDeviceIds *[]string `json:"notFoundDeviceIds,omitempty"`

	// Not supported devices
	NotSupportedDeviceIds *[]string `json:"notSupportedDeviceIds,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Successful devices
	SuccessfulDeviceIds *[]string `json:"successfulDeviceIds,omitempty"`
}
