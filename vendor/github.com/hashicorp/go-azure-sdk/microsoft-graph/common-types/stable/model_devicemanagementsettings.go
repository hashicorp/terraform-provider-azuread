package stable

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementSettings struct {
	// The number of days a device is allowed to go without checking in to remain compliant.
	DeviceComplianceCheckinThresholdDays *int64 `json:"deviceComplianceCheckinThresholdDays,omitempty"`

	// Is feature enabled or not for scheduled action for rule.
	IsScheduledActionEnabled *bool `json:"isScheduledActionEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Device should be noncompliant when there is no compliance policy targeted when this is true
	SecureByDefault *bool `json:"secureByDefault,omitempty"`
}
