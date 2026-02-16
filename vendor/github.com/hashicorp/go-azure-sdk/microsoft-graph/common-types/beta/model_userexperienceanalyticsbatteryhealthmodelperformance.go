package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UserExperienceAnalyticsBatteryHealthModelPerformance{}

type UserExperienceAnalyticsBatteryHealthModelPerformance struct {
	// Number of active devices for that model. Valid values 0 to 2147483647
	ActiveDevices *int64 `json:"activeDevices,omitempty"`

	// The mean of the battery age for all devices of a given model in a tenant. Unit in days. Valid values 0 to 2147483647
	AverageBatteryAgeInDays *int64 `json:"averageBatteryAgeInDays,omitempty"`

	// The mean of the estimated runtimes on full charge for all devices of a given model. Unit in minutes. Valid values 0
	// to 2147483647
	AverageEstimatedRuntimeInMinutes *int64 `json:"averageEstimatedRuntimeInMinutes,omitempty"`

	// The mean of the maximum capacity for all devices of a given model. Maximum capacity measures the full charge vs.
	// design capacity for a device’s batteries.. Valid values 0 to 2147483647
	AverageMaxCapacityPercentage *int64 `json:"averageMaxCapacityPercentage,omitempty"`

	// The manufacturer name of the device.
	DeviceManufacturerName nullable.Type[string] `json:"deviceManufacturerName,omitempty"`

	// The model name of the device.
	DeviceModelName nullable.Type[string] `json:"deviceModelName,omitempty"`

	// Name of the device manufacturer. Deprecated in favor of DeviceManufacturerName.
	Manufacturer nullable.Type[string] `json:"manufacturer,omitempty"`

	// The mean of number of times the battery has been discharged an amount that equals 100% of its capacity for all
	// devices of a given model in a tenant. Valid values 0 to 2147483647
	MeanFullBatteryDrainCount *int64 `json:"meanFullBatteryDrainCount,omitempty"`

	// The median of the estimated runtimes on full charge for all devices of a given model. Unit in minutes. Valid values 0
	// to 2147483647
	MedianEstimatedRuntimeInMinutes *int64 `json:"medianEstimatedRuntimeInMinutes,omitempty"`

	// The median of number of times the battery has been discharged an amount that equals 100% of its capacity for all
	// devices of a given model in a tenant. Valid values 0 to 2147483647
	MedianFullBatteryDrainCount *int64 `json:"medianFullBatteryDrainCount,omitempty"`

	// The median of the maximum capacity for all devices of a given model. Maximum capacity measures the full charge vs.
	// design capacity for a device’s batteries.. Valid values 0 to 2147483647
	MedianMaxCapacityPercentage *int64 `json:"medianMaxCapacityPercentage,omitempty"`

	// The model name of the device. Deprecated in favor of DeviceModelName.
	Model nullable.Type[string] `json:"model,omitempty"`

	// A weighted average of a model’s maximum capacity score and runtime estimate score. Values range from 0-100. Valid
	// values 0 to 2147483647
	ModelBatteryHealthScore *int64 `json:"modelBatteryHealthScore,omitempty"`

	ModelHealthStatus *UserExperienceAnalyticsHealthState `json:"modelHealthStatus,omitempty"`

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

func (s UserExperienceAnalyticsBatteryHealthModelPerformance) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserExperienceAnalyticsBatteryHealthModelPerformance{}

func (s UserExperienceAnalyticsBatteryHealthModelPerformance) MarshalJSON() ([]byte, error) {
	type wrapper UserExperienceAnalyticsBatteryHealthModelPerformance
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserExperienceAnalyticsBatteryHealthModelPerformance: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserExperienceAnalyticsBatteryHealthModelPerformance: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userExperienceAnalyticsBatteryHealthModelPerformance"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserExperienceAnalyticsBatteryHealthModelPerformance: %+v", err)
	}

	return encoded, nil
}
