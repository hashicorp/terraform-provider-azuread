package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerSystemUpdateFreezePeriod struct {
	// The day of the end date of the freeze period. Valid values 1 to 31
	EndDay *int64 `json:"endDay,omitempty"`

	// The month of the end date of the freeze period. Valid values 1 to 12
	EndMonth *int64 `json:"endMonth,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The day of the start date of the freeze period. Valid values 1 to 31
	StartDay *int64 `json:"startDay,omitempty"`

	// The month of the start date of the freeze period. Valid values 1 to 12
	StartMonth *int64 `json:"startMonth,omitempty"`
}
