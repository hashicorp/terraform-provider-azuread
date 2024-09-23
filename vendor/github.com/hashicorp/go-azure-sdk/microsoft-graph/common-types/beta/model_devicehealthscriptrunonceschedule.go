package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceHealthScriptTimeSchedule = DeviceHealthScriptRunOnceSchedule{}

type DeviceHealthScriptRunOnceSchedule struct {
	// The date the script is scheduled to run. This collection can contain a maximum of 20 elements.
	Date nullable.Type[string] `json:"date,omitempty"`

	// Fields inherited from DeviceHealthScriptTimeSchedule

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

func (s DeviceHealthScriptRunOnceSchedule) DeviceHealthScriptTimeSchedule() BaseDeviceHealthScriptTimeScheduleImpl {
	return BaseDeviceHealthScriptTimeScheduleImpl{
		Time:      s.Time,
		UseUtc:    s.UseUtc,
		Interval:  s.Interval,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s DeviceHealthScriptRunOnceSchedule) DeviceHealthScriptRunSchedule() BaseDeviceHealthScriptRunScheduleImpl {
	return BaseDeviceHealthScriptRunScheduleImpl{
		Interval:  s.Interval,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceHealthScriptRunOnceSchedule{}

func (s DeviceHealthScriptRunOnceSchedule) MarshalJSON() ([]byte, error) {
	type wrapper DeviceHealthScriptRunOnceSchedule
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceHealthScriptRunOnceSchedule: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceHealthScriptRunOnceSchedule: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceHealthScriptRunOnceSchedule"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceHealthScriptRunOnceSchedule: %+v", err)
	}

	return encoded, nil
}
