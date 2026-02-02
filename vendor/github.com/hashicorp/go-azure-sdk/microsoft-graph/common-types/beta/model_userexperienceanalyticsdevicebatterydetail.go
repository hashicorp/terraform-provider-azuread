package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsDeviceBatteryDetail struct {
	// Uniquely identifies the batteries in a single device.
	BatteryId nullable.Type[string] `json:"batteryId,omitempty"`

	// Number of times the battery has been discharged an amount that equals 100% of its capacity, but not necessarily by
	// discharging it from 100% to 0%. Valid values 0 to 2147483647
	FullBatteryDrainCount *int64 `json:"fullBatteryDrainCount,omitempty"`

	// Ratio of current capacity and design capacity of the battery. Unit in percentage and values range from 0-100. Valid
	// values 0 to 2147483647
	MaxCapacityPercentage *int64 `json:"maxCapacityPercentage,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
