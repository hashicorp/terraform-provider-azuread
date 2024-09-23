package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UserExperienceAnalyticsAppHealthDevicePerformance{}

type UserExperienceAnalyticsAppHealthDevicePerformance struct {
	// The number of application crashes for the device. Valid values 0 to 2147483647. Supports: $filter, $select, $OrderBy.
	// Read-only. Valid values -2147483648 to 2147483647
	AppCrashCount *int64 `json:"appCrashCount,omitempty"`

	// The number of application hangs for the device. Valid values 0 to 2147483647. Supports: $select, $OrderBy. Read-only.
	// Valid values -2147483648 to 2147483647
	AppHangCount *int64 `json:"appHangCount,omitempty"`

	// The number of distinct application crashes for the device. Valid values 0 to 2147483647. Supports: $select, $OrderBy.
	// Read-only. Valid values -2147483648 to 2147483647
	CrashedAppCount *int64 `json:"crashedAppCount,omitempty"`

	// The name of the device. Supports: $select, $OrderBy. Read-only.
	DeviceDisplayName nullable.Type[string] `json:"deviceDisplayName,omitempty"`

	// The Intune device id of the device. Supports: $select, $OrderBy. Read-only.
	DeviceId nullable.Type[string] `json:"deviceId,omitempty"`

	// The manufacturer name of the device. Supports: $select, $OrderBy. Read-only.
	DeviceManufacturer nullable.Type[string] `json:"deviceManufacturer,omitempty"`

	// The model name of the device. Supports: $select, $OrderBy. Read-only.
	DeviceModel nullable.Type[string] `json:"deviceModel,omitempty"`

	HealthStatus *UserExperienceAnalyticsHealthState `json:"healthStatus,omitempty"`

	// The mean time to failure for the application in minutes. Valid values 0 to 2147483647. Supports: $filter, $select,
	// $OrderBy. Read-only. Valid values -2147483648 to 2147483647
	MeanTimeToFailureInMinutes *int64 `json:"meanTimeToFailureInMinutes,omitempty"`

	// The date and time when the statistics were last computed. The value cannot be modified and is automatically populated
	// when the statistics are computed. The Timestamp type represents date and time information using ISO 8601 format and
	// is always in UTC time. For example, midnight UTC on Jan 1, 2022 would look like this: '2022-01-01T00:00:00Z'.
	// Returned by default. Read-only.
	ProcessedDateTime *string `json:"processedDateTime,omitempty"`

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

func (s UserExperienceAnalyticsAppHealthDevicePerformance) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserExperienceAnalyticsAppHealthDevicePerformance{}

func (s UserExperienceAnalyticsAppHealthDevicePerformance) MarshalJSON() ([]byte, error) {
	type wrapper UserExperienceAnalyticsAppHealthDevicePerformance
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserExperienceAnalyticsAppHealthDevicePerformance: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserExperienceAnalyticsAppHealthDevicePerformance: %+v", err)
	}

	delete(decoded, "appCrashCount")
	delete(decoded, "appHangCount")
	delete(decoded, "crashedAppCount")
	delete(decoded, "deviceDisplayName")
	delete(decoded, "deviceId")
	delete(decoded, "deviceManufacturer")
	delete(decoded, "deviceModel")
	delete(decoded, "meanTimeToFailureInMinutes")
	delete(decoded, "processedDateTime")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userExperienceAnalyticsAppHealthDevicePerformance"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserExperienceAnalyticsAppHealthDevicePerformance: %+v", err)
	}

	return encoded, nil
}
