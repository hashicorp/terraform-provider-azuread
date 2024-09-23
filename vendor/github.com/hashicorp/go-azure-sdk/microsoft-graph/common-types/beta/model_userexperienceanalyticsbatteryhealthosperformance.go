package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UserExperienceAnalyticsBatteryHealthOsPerformance{}

type UserExperienceAnalyticsBatteryHealthOsPerformance struct {
	// Number of active devices for that os version. Valid values 0 to 2147483647
	ActiveDevices *int64 `json:"activeDevices,omitempty"`

	// The mean of the battery age for all devices running a particular operating system version in a tenant. Unit in days.
	// Valid values 0 to 2147483647
	AverageBatteryAgeInDays *int64 `json:"averageBatteryAgeInDays,omitempty"`

	// The mean of the estimated runtimes on full charge for all devices running a particular operating system version. Unit
	// in minutes. Valid values 0 to 2147483647
	AverageEstimatedRuntimeInMinutes *int64 `json:"averageEstimatedRuntimeInMinutes,omitempty"`

	// The mean of the maximum capacity for all devices running a particular operating system version. Maximum capacity
	// measures the full charge vs. design capacity for a device’s batteries.. Valid values 0 to 2147483647
	AverageMaxCapacityPercentage *int64 `json:"averageMaxCapacityPercentage,omitempty"`

	// The mean of number of times the battery has been discharged an amount that equals 100% of its capacity for all
	// devices running a particular operating system version in a tenant. Valid values 0 to 2147483647
	MeanFullBatteryDrainCount *int64 `json:"meanFullBatteryDrainCount,omitempty"`

	// The median of the estimated runtimes on full charge for all devices running a particular operating system version.
	// Unit in minutes. Valid values 0 to 2147483647
	MedianEstimatedRuntimeInMinutes *int64 `json:"medianEstimatedRuntimeInMinutes,omitempty"`

	// The median of number of times the battery has been discharged an amount that equals 100% of its capacity for all
	// devices running a particular operating system version in a tenant. Valid values 0 to 2147483647
	MedianFullBatteryDrainCount *int64 `json:"medianFullBatteryDrainCount,omitempty"`

	// The median of the maximum capacity for all devices running a particular operating system version. Maximum capacity
	// measures the full charge vs. design capacity for a device’s batteries.. Valid values 0 to 2147483647
	MedianMaxCapacityPercentage *int64 `json:"medianMaxCapacityPercentage,omitempty"`

	// A weighted average of battery health score across all devices running a particular operating system version. Values
	// range from 0-100. Valid values 0 to 2147483647
	OsBatteryHealthScore *int64 `json:"osBatteryHealthScore,omitempty"`

	// Build number of the operating system.
	OsBuildNumber nullable.Type[string] `json:"osBuildNumber,omitempty"`

	OsHealthStatus *UserExperienceAnalyticsHealthState `json:"osHealthStatus,omitempty"`

	// Version of the operating system.
	OsVersion nullable.Type[string] `json:"osVersion,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s UserExperienceAnalyticsBatteryHealthOsPerformance) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserExperienceAnalyticsBatteryHealthOsPerformance{}

func (s UserExperienceAnalyticsBatteryHealthOsPerformance) MarshalJSON() ([]byte, error) {
	type wrapper UserExperienceAnalyticsBatteryHealthOsPerformance
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserExperienceAnalyticsBatteryHealthOsPerformance: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserExperienceAnalyticsBatteryHealthOsPerformance: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userExperienceAnalyticsBatteryHealthOsPerformance"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserExperienceAnalyticsBatteryHealthOsPerformance: %+v", err)
	}

	return encoded, nil
}
