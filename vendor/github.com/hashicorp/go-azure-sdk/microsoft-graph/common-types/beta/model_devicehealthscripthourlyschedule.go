package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceHealthScriptRunSchedule = DeviceHealthScriptHourlySchedule{}

type DeviceHealthScriptHourlySchedule struct {

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

func (s DeviceHealthScriptHourlySchedule) DeviceHealthScriptRunSchedule() BaseDeviceHealthScriptRunScheduleImpl {
	return BaseDeviceHealthScriptRunScheduleImpl{
		Interval:  s.Interval,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceHealthScriptHourlySchedule{}

func (s DeviceHealthScriptHourlySchedule) MarshalJSON() ([]byte, error) {
	type wrapper DeviceHealthScriptHourlySchedule
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceHealthScriptHourlySchedule: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceHealthScriptHourlySchedule: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceHealthScriptHourlySchedule"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceHealthScriptHourlySchedule: %+v", err)
	}

	return encoded, nil
}
