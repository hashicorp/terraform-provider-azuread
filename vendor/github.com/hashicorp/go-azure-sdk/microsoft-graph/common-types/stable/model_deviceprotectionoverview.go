package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceProtectionOverview struct {
	// Indicates number of devices reporting as clean
	CleanDeviceCount *int64 `json:"cleanDeviceCount,omitempty"`

	// Indicates number of devices with critical failures
	CriticalFailuresDeviceCount *int64 `json:"criticalFailuresDeviceCount,omitempty"`

	// Indicates number of devices with inactive threat agent
	InactiveThreatAgentDeviceCount *int64 `json:"inactiveThreatAgentDeviceCount,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Indicates number of devices pending full scan
	PendingFullScanDeviceCount *int64 `json:"pendingFullScanDeviceCount,omitempty"`

	// Indicates number of devices with pending manual steps
	PendingManualStepsDeviceCount *int64 `json:"pendingManualStepsDeviceCount,omitempty"`

	// Indicates number of pending offline scan devices
	PendingOfflineScanDeviceCount *int64 `json:"pendingOfflineScanDeviceCount,omitempty"`

	// Indicates the number of devices that have a pending full scan. Valid values -2147483648 to 2147483647
	PendingQuickScanDeviceCount *int64 `json:"pendingQuickScanDeviceCount,omitempty"`

	// Indicates number of devices pending restart
	PendingRestartDeviceCount *int64 `json:"pendingRestartDeviceCount,omitempty"`

	// Indicates number of devices with an old signature
	PendingSignatureUpdateDeviceCount *int64 `json:"pendingSignatureUpdateDeviceCount,omitempty"`

	// Total device count.
	TotalReportedDeviceCount *int64 `json:"totalReportedDeviceCount,omitempty"`

	// Indicates number of devices with threat agent state as unknown
	UnknownStateThreatAgentDeviceCount *int64 `json:"unknownStateThreatAgentDeviceCount,omitempty"`
}
