package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId{}

type UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId struct {
	// The number of crashes for the app. Valid values -2147483648 to 2147483647
	AppCrashCount *int64 `json:"appCrashCount,omitempty"`

	// The friendly name of the application.
	AppDisplayName nullable.Type[string] `json:"appDisplayName,omitempty"`

	// The name of the application.
	AppName nullable.Type[string] `json:"appName,omitempty"`

	// The publisher of the application.
	AppPublisher nullable.Type[string] `json:"appPublisher,omitempty"`

	// The version of the application.
	AppVersion nullable.Type[string] `json:"appVersion,omitempty"`

	// The name of the device. Supports: $select, $OrderBy. Read-only.
	DeviceDisplayName nullable.Type[string] `json:"deviceDisplayName,omitempty"`

	// The Intune device id of the device. Supports: $select, $OrderBy. Read-only.
	DeviceId nullable.Type[string] `json:"deviceId,omitempty"`

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

func (s UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId{}

func (s UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId) MarshalJSON() ([]byte, error) {
	type wrapper UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId: %+v", err)
	}

	delete(decoded, "deviceDisplayName")
	delete(decoded, "deviceId")
	delete(decoded, "processedDateTime")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId: %+v", err)
	}

	return encoded, nil
}
