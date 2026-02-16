package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UserExperienceAnalyticsBatteryHealthDevicePerformance{}

type UserExperienceAnalyticsBatteryHealthDevicePerformance struct {
	// Estimated battery age. Unit in days. Valid values 0 to 2147483647
	BatteryAgeInDays *int64 `json:"batteryAgeInDays,omitempty"`

	// Properties (maxCapacity and cycleCount) related to all batteries of the device.
	DeviceBatteriesDetails *[]UserExperienceAnalyticsDeviceBatteryDetail `json:"deviceBatteriesDetails,omitempty"`

	// Number of batteries in a user device. Valid values 0 to 2147483647
	DeviceBatteryCount *int64 `json:"deviceBatteryCount,omitempty"`

	// A weighted average of a device’s maximum capacity score and runtime estimate score. Values range from 0-100. Valid
	// values 0 to 2147483647
	DeviceBatteryHealthScore *int64 `json:"deviceBatteryHealthScore,omitempty"`

	// Tags for computed information on how battery on the device is behaving. E.g. newbattery, batterycapacityred,
	// designcapacityzero, etc.
	DeviceBatteryTags *[]string `json:"deviceBatteryTags,omitempty"`

	// The unique identifier of the device, Intune DeviceID.
	DeviceId nullable.Type[string] `json:"deviceId,omitempty"`

	// The manufacturer name of the device.
	DeviceManufacturerName nullable.Type[string] `json:"deviceManufacturerName,omitempty"`

	// The model name of the device.
	DeviceModelName nullable.Type[string] `json:"deviceModelName,omitempty"`

	// Device friendly name.
	DeviceName nullable.Type[string] `json:"deviceName,omitempty"`

	// The estimated runtime of the device when the battery is fully charged. Unit in minutes. Valid values 0 to 2147483647
	EstimatedRuntimeInMinutes *int64 `json:"estimatedRuntimeInMinutes,omitempty"`

	// Number of times the battery has been discharged an amount that equals 100% of its capacity, but not necessarily by
	// discharging it from 100% to 0%. Valid values 0 to 2147483647
	FullBatteryDrainCount *int64 `json:"fullBatteryDrainCount,omitempty"`

	HealthStatus *UserExperienceAnalyticsHealthState `json:"healthStatus,omitempty"`

	// The manufacturer name of the device. Deprecated in favor of DeviceManufacturerName.
	Manufacturer nullable.Type[string] `json:"manufacturer,omitempty"`

	// Ratio of current capacity and design capacity of the battery with the lowest capacity. Unit in percentage and values
	// range from 0-100. Valid values 0 to 2147483647
	MaxCapacityPercentage *int64 `json:"maxCapacityPercentage,omitempty"`

	// The model name of the device. Deprecated in favor of DeviceModelName.
	Model nullable.Type[string] `json:"model,omitempty"`

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

func (s UserExperienceAnalyticsBatteryHealthDevicePerformance) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserExperienceAnalyticsBatteryHealthDevicePerformance{}

func (s UserExperienceAnalyticsBatteryHealthDevicePerformance) MarshalJSON() ([]byte, error) {
	type wrapper UserExperienceAnalyticsBatteryHealthDevicePerformance
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserExperienceAnalyticsBatteryHealthDevicePerformance: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserExperienceAnalyticsBatteryHealthDevicePerformance: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userExperienceAnalyticsBatteryHealthDevicePerformance"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserExperienceAnalyticsBatteryHealthDevicePerformance: %+v", err)
	}

	return encoded, nil
}
