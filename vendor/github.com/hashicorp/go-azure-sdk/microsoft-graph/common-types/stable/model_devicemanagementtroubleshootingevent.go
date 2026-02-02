package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementTroubleshootingEvent interface {
	Entity
	DeviceManagementTroubleshootingEvent() BaseDeviceManagementTroubleshootingEventImpl
}

var _ DeviceManagementTroubleshootingEvent = BaseDeviceManagementTroubleshootingEventImpl{}

type BaseDeviceManagementTroubleshootingEventImpl struct {
	// Id used for tracing the failure in the service.
	CorrelationId nullable.Type[string] `json:"correlationId,omitempty"`

	// Time when the event occurred .
	EventDateTime *string `json:"eventDateTime,omitempty"`

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

func (s BaseDeviceManagementTroubleshootingEventImpl) DeviceManagementTroubleshootingEvent() BaseDeviceManagementTroubleshootingEventImpl {
	return s
}

func (s BaseDeviceManagementTroubleshootingEventImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ DeviceManagementTroubleshootingEvent = RawDeviceManagementTroubleshootingEventImpl{}

// RawDeviceManagementTroubleshootingEventImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDeviceManagementTroubleshootingEventImpl struct {
	deviceManagementTroubleshootingEvent BaseDeviceManagementTroubleshootingEventImpl
	Type                                 string
	Values                               map[string]interface{}
}

func (s RawDeviceManagementTroubleshootingEventImpl) DeviceManagementTroubleshootingEvent() BaseDeviceManagementTroubleshootingEventImpl {
	return s.deviceManagementTroubleshootingEvent
}

func (s RawDeviceManagementTroubleshootingEventImpl) Entity() BaseEntityImpl {
	return s.deviceManagementTroubleshootingEvent.Entity()
}

var _ json.Marshaler = BaseDeviceManagementTroubleshootingEventImpl{}

func (s BaseDeviceManagementTroubleshootingEventImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseDeviceManagementTroubleshootingEventImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseDeviceManagementTroubleshootingEventImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseDeviceManagementTroubleshootingEventImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementTroubleshootingEvent"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseDeviceManagementTroubleshootingEventImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalDeviceManagementTroubleshootingEventImplementation(input []byte) (DeviceManagementTroubleshootingEvent, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementTroubleshootingEvent into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.enrollmentTroubleshootingEvent") {
		var out EnrollmentTroubleshootingEvent
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EnrollmentTroubleshootingEvent: %+v", err)
		}
		return out, nil
	}

	var parent BaseDeviceManagementTroubleshootingEventImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDeviceManagementTroubleshootingEventImpl: %+v", err)
	}

	return RawDeviceManagementTroubleshootingEventImpl{
		deviceManagementTroubleshootingEvent: parent,
		Type:                                 value,
		Values:                               temp,
	}, nil

}
