package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UserExperienceAnalyticsAppHealthDeviceModelPerformance{}

type UserExperienceAnalyticsAppHealthDeviceModelPerformance struct {
	// The number of active devices for the model. Valid values 0 to 2147483647. Supports: $filter, $select, $OrderBy.
	// Read-only. Valid values -2147483648 to 2147483647
	ActiveDeviceCount *int64 `json:"activeDeviceCount,omitempty"`

	// The manufacturer name of the device. Supports: $select, $OrderBy. Read-only.
	DeviceManufacturer nullable.Type[string] `json:"deviceManufacturer,omitempty"`

	// The model name of the device. Supports: $select, $OrderBy. Read-only.
	DeviceModel nullable.Type[string] `json:"deviceModel,omitempty"`

	HealthStatus *UserExperienceAnalyticsHealthState `json:"healthStatus,omitempty"`

	// The mean time to failure for the application in minutes. Valid values 0 to 2147483647. Supports: $filter, $select,
	// $OrderBy. Read-only. Valid values -2147483648 to 2147483647
	MeanTimeToFailureInMinutes *int64 `json:"meanTimeToFailureInMinutes,omitempty"`

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

func (s UserExperienceAnalyticsAppHealthDeviceModelPerformance) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserExperienceAnalyticsAppHealthDeviceModelPerformance{}

func (s UserExperienceAnalyticsAppHealthDeviceModelPerformance) MarshalJSON() ([]byte, error) {
	type wrapper UserExperienceAnalyticsAppHealthDeviceModelPerformance
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserExperienceAnalyticsAppHealthDeviceModelPerformance: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserExperienceAnalyticsAppHealthDeviceModelPerformance: %+v", err)
	}

	delete(decoded, "activeDeviceCount")
	delete(decoded, "deviceManufacturer")
	delete(decoded, "deviceModel")
	delete(decoded, "meanTimeToFailureInMinutes")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userExperienceAnalyticsAppHealthDeviceModelPerformance"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserExperienceAnalyticsAppHealthDeviceModelPerformance: %+v", err)
	}

	return encoded, nil
}
