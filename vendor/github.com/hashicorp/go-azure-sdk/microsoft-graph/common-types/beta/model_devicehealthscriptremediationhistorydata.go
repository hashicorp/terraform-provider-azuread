package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceHealthScriptRemediationHistoryData struct {
	// The date on which devices were remediated by the device health script.
	Date *string `json:"date,omitempty"`

	// The number of devices for which the detection script found an issue.
	DetectFailedDeviceCount *int64 `json:"detectFailedDeviceCount,omitempty"`

	// The number of devices that were found to have no issue by the device health script.
	NoIssueDeviceCount *int64 `json:"noIssueDeviceCount,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The number of devices remediated by the device health script.
	RemediatedDeviceCount *int64 `json:"remediatedDeviceCount,omitempty"`
}
