package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigManagerPolicySummary struct {
	// The number of devices evaluated to be compliant by the policy.
	CompliantDeviceCount *int64 `json:"compliantDeviceCount,omitempty"`

	// The number of devices that have have been remediated by the policy.
	EnforcedDeviceCount *int64 `json:"enforcedDeviceCount,omitempty"`

	// The number of devices that failed to be evaluated by the policy.
	FailedDeviceCount *int64 `json:"failedDeviceCount,omitempty"`

	// The number of devices evaluated to be noncompliant by the policy.
	NonCompliantDeviceCount *int64 `json:"nonCompliantDeviceCount,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The number of devices that have acknowledged the policy but are pending evaluation.
	PendingDeviceCount *int64 `json:"pendingDeviceCount,omitempty"`

	// The number of devices targeted by the policy.
	TargetedDeviceCount *int64 `json:"targetedDeviceCount,omitempty"`
}
