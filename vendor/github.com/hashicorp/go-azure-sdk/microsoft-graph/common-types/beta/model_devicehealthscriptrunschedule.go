package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceHealthScriptRunSchedule interface {
	DeviceHealthScriptRunSchedule() BaseDeviceHealthScriptRunScheduleImpl
}

var _ DeviceHealthScriptRunSchedule = BaseDeviceHealthScriptRunScheduleImpl{}

type BaseDeviceHealthScriptRunScheduleImpl struct {
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

func (s BaseDeviceHealthScriptRunScheduleImpl) DeviceHealthScriptRunSchedule() BaseDeviceHealthScriptRunScheduleImpl {
	return s
}

var _ DeviceHealthScriptRunSchedule = RawDeviceHealthScriptRunScheduleImpl{}

// RawDeviceHealthScriptRunScheduleImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDeviceHealthScriptRunScheduleImpl struct {
	deviceHealthScriptRunSchedule BaseDeviceHealthScriptRunScheduleImpl
	Type                          string
	Values                        map[string]interface{}
}

func (s RawDeviceHealthScriptRunScheduleImpl) DeviceHealthScriptRunSchedule() BaseDeviceHealthScriptRunScheduleImpl {
	return s.deviceHealthScriptRunSchedule
}

func UnmarshalDeviceHealthScriptRunScheduleImplementation(input []byte) (DeviceHealthScriptRunSchedule, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceHealthScriptRunSchedule into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceHealthScriptHourlySchedule") {
		var out DeviceHealthScriptHourlySchedule
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceHealthScriptHourlySchedule: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceHealthScriptTimeSchedule") {
		var out DeviceHealthScriptTimeSchedule
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceHealthScriptTimeSchedule: %+v", err)
		}
		return out, nil
	}

	var parent BaseDeviceHealthScriptRunScheduleImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDeviceHealthScriptRunScheduleImpl: %+v", err)
	}

	return RawDeviceHealthScriptRunScheduleImpl{
		deviceHealthScriptRunSchedule: parent,
		Type:                          value,
		Values:                        temp,
	}, nil

}
