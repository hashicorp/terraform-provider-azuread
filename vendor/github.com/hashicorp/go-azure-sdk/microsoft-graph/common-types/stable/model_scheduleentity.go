package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScheduleEntity interface {
	ScheduleEntity() BaseScheduleEntityImpl
}

var _ ScheduleEntity = BaseScheduleEntityImpl{}

type BaseScheduleEntityImpl struct {
	EndDateTime nullable.Type[string] `json:"endDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	StartDateTime nullable.Type[string] `json:"startDateTime,omitempty"`
	Theme         *ScheduleEntityTheme  `json:"theme,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseScheduleEntityImpl) ScheduleEntity() BaseScheduleEntityImpl {
	return s
}

var _ ScheduleEntity = RawScheduleEntityImpl{}

// RawScheduleEntityImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawScheduleEntityImpl struct {
	scheduleEntity BaseScheduleEntityImpl
	Type           string
	Values         map[string]interface{}
}

func (s RawScheduleEntityImpl) ScheduleEntity() BaseScheduleEntityImpl {
	return s.scheduleEntity
}

func UnmarshalScheduleEntityImplementation(input []byte) (ScheduleEntity, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ScheduleEntity into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.shiftItem") {
		var out ShiftItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ShiftItem: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.timeOffItem") {
		var out TimeOffItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TimeOffItem: %+v", err)
		}
		return out, nil
	}

	var parent BaseScheduleEntityImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseScheduleEntityImpl: %+v", err)
	}

	return RawScheduleEntityImpl{
		scheduleEntity: parent,
		Type:           value,
		Values:         temp,
	}, nil

}
