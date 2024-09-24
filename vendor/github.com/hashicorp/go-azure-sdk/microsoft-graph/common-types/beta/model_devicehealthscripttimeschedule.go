package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceHealthScriptTimeSchedule interface {
	DeviceHealthScriptRunSchedule
	DeviceHealthScriptTimeSchedule() BaseDeviceHealthScriptTimeScheduleImpl
}

var _ DeviceHealthScriptTimeSchedule = BaseDeviceHealthScriptTimeScheduleImpl{}

type BaseDeviceHealthScriptTimeScheduleImpl struct {
	// At what time the script is scheduled to run. This collection can contain a maximum of 20 elements.
	Time nullable.Type[string] `json:"time,omitempty"`

	// Indicate if the time is Utc or client local time.
	UseUtc *bool `json:"useUtc,omitempty"`

	// Fields inherited from DeviceHealthScriptRunSchedule

	// The x value of every x hours for hourly schedule, every x days for Daily Schedule, every x weeks for weekly schedule,
	// every x months for Monthly Schedule. Valid values 1 to 23
	Interval *int64 `json:"interval,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseDeviceHealthScriptTimeScheduleImpl) DeviceHealthScriptTimeSchedule() BaseDeviceHealthScriptTimeScheduleImpl {
	return s
}

func (s BaseDeviceHealthScriptTimeScheduleImpl) DeviceHealthScriptRunSchedule() BaseDeviceHealthScriptRunScheduleImpl {
	return BaseDeviceHealthScriptRunScheduleImpl{
		Interval:  s.Interval,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ DeviceHealthScriptTimeSchedule = RawDeviceHealthScriptTimeScheduleImpl{}

// RawDeviceHealthScriptTimeScheduleImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDeviceHealthScriptTimeScheduleImpl struct {
	deviceHealthScriptTimeSchedule BaseDeviceHealthScriptTimeScheduleImpl
	Type                           string
	Values                         map[string]interface{}
}

func (s RawDeviceHealthScriptTimeScheduleImpl) DeviceHealthScriptTimeSchedule() BaseDeviceHealthScriptTimeScheduleImpl {
	return s.deviceHealthScriptTimeSchedule
}

func (s RawDeviceHealthScriptTimeScheduleImpl) DeviceHealthScriptRunSchedule() BaseDeviceHealthScriptRunScheduleImpl {
	return s.deviceHealthScriptTimeSchedule.DeviceHealthScriptRunSchedule()
}

var _ json.Marshaler = BaseDeviceHealthScriptTimeScheduleImpl{}

func (s BaseDeviceHealthScriptTimeScheduleImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseDeviceHealthScriptTimeScheduleImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseDeviceHealthScriptTimeScheduleImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseDeviceHealthScriptTimeScheduleImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceHealthScriptTimeSchedule"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseDeviceHealthScriptTimeScheduleImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalDeviceHealthScriptTimeScheduleImplementation(input []byte) (DeviceHealthScriptTimeSchedule, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceHealthScriptTimeSchedule into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceHealthScriptDailySchedule") {
		var out DeviceHealthScriptDailySchedule
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceHealthScriptDailySchedule: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceHealthScriptRunOnceSchedule") {
		var out DeviceHealthScriptRunOnceSchedule
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceHealthScriptRunOnceSchedule: %+v", err)
		}
		return out, nil
	}

	var parent BaseDeviceHealthScriptTimeScheduleImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDeviceHealthScriptTimeScheduleImpl: %+v", err)
	}

	return RawDeviceHealthScriptTimeScheduleImpl{
		deviceHealthScriptTimeSchedule: parent,
		Type:                           value,
		Values:                         temp,
	}, nil

}
