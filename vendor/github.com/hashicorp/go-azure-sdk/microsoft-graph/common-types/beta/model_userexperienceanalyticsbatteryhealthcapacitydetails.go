package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UserExperienceAnalyticsBatteryHealthCapacityDetails{}

type UserExperienceAnalyticsBatteryHealthCapacityDetails struct {
	// Number of active devices within the tenant. Valid values 0 to 2147483647
	ActiveDevices *int64 `json:"activeDevices,omitempty"`

	// Number of devices whose battery maximum capacity is greater than 50% but lesser than 80%. Valid values 0 to
	// 2147483647
	BatteryCapacityFair *int64 `json:"batteryCapacityFair,omitempty"`

	// Number of devices whose battery maximum capacity is greater than 80%. Valid values 0 to 2147483647
	BatteryCapacityGood *int64 `json:"batteryCapacityGood,omitempty"`

	// Number of devices whose battery maximum capacity is lesser than 50%. Valid values 0 to 2147483647
	BatteryCapacityPoor *int64 `json:"batteryCapacityPoor,omitempty"`

	// Recorded date time of this capacity details instance.
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

func (s UserExperienceAnalyticsBatteryHealthCapacityDetails) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserExperienceAnalyticsBatteryHealthCapacityDetails{}

func (s UserExperienceAnalyticsBatteryHealthCapacityDetails) MarshalJSON() ([]byte, error) {
	type wrapper UserExperienceAnalyticsBatteryHealthCapacityDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserExperienceAnalyticsBatteryHealthCapacityDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserExperienceAnalyticsBatteryHealthCapacityDetails: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userExperienceAnalyticsBatteryHealthCapacityDetails"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserExperienceAnalyticsBatteryHealthCapacityDetails: %+v", err)
	}

	return encoded, nil
}
