package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UserExperienceAnalyticsAnomalyDevice{}

type UserExperienceAnalyticsAnomalyDevice struct {
	// The unique identifier of the anomaly.
	AnomalyId nullable.Type[string] `json:"anomalyId,omitempty"`

	// Indicates the first occurance date and time for the anomaly on the device.
	AnomalyOnDeviceFirstOccurrenceDateTime *string `json:"anomalyOnDeviceFirstOccurrenceDateTime,omitempty"`

	// Indicates the latest occurance date and time for the anomaly on the device.
	AnomalyOnDeviceLatestOccurrenceDateTime *string `json:"anomalyOnDeviceLatestOccurrenceDateTime,omitempty"`

	// The unique identifier of the correlation group.
	CorrelationGroupId nullable.Type[string] `json:"correlationGroupId,omitempty"`

	// The unique identifier of the device.
	DeviceId nullable.Type[string] `json:"deviceId,omitempty"`

	// The manufacturer name of the device.
	DeviceManufacturer nullable.Type[string] `json:"deviceManufacturer,omitempty"`

	// The model name of the device.
	DeviceModel nullable.Type[string] `json:"deviceModel,omitempty"`

	// The name of the device.
	DeviceName nullable.Type[string] `json:"deviceName,omitempty"`

	// Indicates the status of the device in the correlation group. Eg: Device status can be anomalous, affected, at risk.
	DeviceStatus *UserExperienceAnalyticsDeviceStatus `json:"deviceStatus,omitempty"`

	// The name of the OS installed on the device.
	OsName nullable.Type[string] `json:"osName,omitempty"`

	// The OS version installed on the device.
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

func (s UserExperienceAnalyticsAnomalyDevice) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserExperienceAnalyticsAnomalyDevice{}

func (s UserExperienceAnalyticsAnomalyDevice) MarshalJSON() ([]byte, error) {
	type wrapper UserExperienceAnalyticsAnomalyDevice
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserExperienceAnalyticsAnomalyDevice: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserExperienceAnalyticsAnomalyDevice: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userExperienceAnalyticsAnomalyDevice"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserExperienceAnalyticsAnomalyDevice: %+v", err)
	}

	return encoded, nil
}
