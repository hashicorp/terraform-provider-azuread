package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCBulkRemoteActionResult struct {
	// A list of all the Intune managed device IDs that completed the bulk action with a failure.
	FailedDeviceIds *[]string `json:"failedDeviceIds,omitempty"`

	// A list of all the Intune managed device IDs that were not found when the bulk action was attempted.
	NotFoundDeviceIds *[]string `json:"notFoundDeviceIds,omitempty"`

	// A list of all the Intune managed device IDs that were identified as unsupported for the bulk action.
	NotSupportedDeviceIds *[]string `json:"notSupportedDeviceIds,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// A list of all the Intune managed device IDs that completed the bulk action successfully.
	SuccessfulDeviceIds *[]string `json:"successfulDeviceIds,omitempty"`
}
