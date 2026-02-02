package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UserExperienceAnalyticsAppHealthDevicePerformanceDetails{}

type UserExperienceAnalyticsAppHealthDevicePerformanceDetails struct {
	// The friendly name of the application for which the event occurred. Possible values are: outlook.exe, excel.exe.
	// Supports: $select, $OrderBy. Read-only.
	AppDisplayName nullable.Type[string] `json:"appDisplayName,omitempty"`

	// The publisher of the application. Supports: $select, $OrderBy. Read-only.
	AppPublisher nullable.Type[string] `json:"appPublisher,omitempty"`

	// The version of the application. Possible values are: 1.0.0.1, 75.65.23.9. Supports: $select, $OrderBy. Read-only.
	AppVersion nullable.Type[string] `json:"appVersion,omitempty"`

	// The name of the device. Supports: $select, $OrderBy. Read-only.
	DeviceDisplayName nullable.Type[string] `json:"deviceDisplayName,omitempty"`

	// The Intune device id of the device. Supports: $select, $OrderBy. Read-only.
	DeviceId nullable.Type[string] `json:"deviceId,omitempty"`

	// The time the event occurred. The value cannot be modified and is automatically populated when the statistics are
	// computed. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time.
	// For example, midnight UTC on Jan 1, 2022 would look like this: '2022-01-01T00:00:00Z'. Returned by default.
	// Read-only.
	EventDateTime *string `json:"eventDateTime,omitempty"`

	// The type of the event. Supports: $select, $OrderBy. Read-only.
	EventType nullable.Type[string] `json:"eventType,omitempty"`

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

func (s UserExperienceAnalyticsAppHealthDevicePerformanceDetails) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserExperienceAnalyticsAppHealthDevicePerformanceDetails{}

func (s UserExperienceAnalyticsAppHealthDevicePerformanceDetails) MarshalJSON() ([]byte, error) {
	type wrapper UserExperienceAnalyticsAppHealthDevicePerformanceDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserExperienceAnalyticsAppHealthDevicePerformanceDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserExperienceAnalyticsAppHealthDevicePerformanceDetails: %+v", err)
	}

	delete(decoded, "appDisplayName")
	delete(decoded, "appPublisher")
	delete(decoded, "appVersion")
	delete(decoded, "deviceDisplayName")
	delete(decoded, "deviceId")
	delete(decoded, "eventDateTime")
	delete(decoded, "eventType")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userExperienceAnalyticsAppHealthDevicePerformanceDetails"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserExperienceAnalyticsAppHealthDevicePerformanceDetails: %+v", err)
	}

	return encoded, nil
}
