package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UserExperienceAnalyticsBatteryHealthRuntimeDetails{}

type UserExperienceAnalyticsBatteryHealthRuntimeDetails struct {
	// Number of active devices within the tenant. Valid values 0 to 2147483647
	ActiveDevices *int64 `json:"activeDevices,omitempty"`

	// Number of devices whose active runtime is greater than 3 hours but lesser than 5 hours. Valid values 0 to 2147483647
	BatteryRuntimeFair *int64 `json:"batteryRuntimeFair,omitempty"`

	// Number of devices whose active runtime is greater than 5 hours. Valid values 0 to 2147483647
	BatteryRuntimeGood *int64 `json:"batteryRuntimeGood,omitempty"`

	// Number of devices whose active runtime is lesser than 3 hours. Valid values 0 to 2147483647
	BatteryRuntimePoor *int64 `json:"batteryRuntimePoor,omitempty"`

	// Recorded date time of this runtime details instance.
	LastRefreshedDateTime *string `json:"lastRefreshedDateTime,omitempty"`

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

func (s UserExperienceAnalyticsBatteryHealthRuntimeDetails) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserExperienceAnalyticsBatteryHealthRuntimeDetails{}

func (s UserExperienceAnalyticsBatteryHealthRuntimeDetails) MarshalJSON() ([]byte, error) {
	type wrapper UserExperienceAnalyticsBatteryHealthRuntimeDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserExperienceAnalyticsBatteryHealthRuntimeDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserExperienceAnalyticsBatteryHealthRuntimeDetails: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userExperienceAnalyticsBatteryHealthRuntimeDetails"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserExperienceAnalyticsBatteryHealthRuntimeDetails: %+v", err)
	}

	return encoded, nil
}
