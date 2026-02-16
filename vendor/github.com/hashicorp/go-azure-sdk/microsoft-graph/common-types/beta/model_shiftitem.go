package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ShiftItem interface {
	ScheduleEntity
	ShiftItem() BaseShiftItemImpl
}

var _ ShiftItem = BaseShiftItemImpl{}

type BaseShiftItemImpl struct {
	// An incremental part of a shift which can cover details of when and where an employee is during their shift. For
	// example, an assignment or a scheduled break or lunch. Required.
	Activities []ShiftActivity `json:"activities"`

	// The shift label of the shiftItem.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The shift notes for the shiftItem.
	Notes nullable.Type[string] `json:"notes,omitempty"`

	// Fields inherited from ScheduleEntity

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

func (s BaseShiftItemImpl) ShiftItem() BaseShiftItemImpl {
	return s
}

func (s BaseShiftItemImpl) ScheduleEntity() BaseScheduleEntityImpl {
	return BaseScheduleEntityImpl{
		EndDateTime:   s.EndDateTime,
		ODataId:       s.ODataId,
		ODataType:     s.ODataType,
		StartDateTime: s.StartDateTime,
		Theme:         s.Theme,
	}
}

var _ ShiftItem = RawShiftItemImpl{}

// RawShiftItemImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawShiftItemImpl struct {
	shiftItem BaseShiftItemImpl
	Type      string
	Values    map[string]interface{}
}

func (s RawShiftItemImpl) ShiftItem() BaseShiftItemImpl {
	return s.shiftItem
}

func (s RawShiftItemImpl) ScheduleEntity() BaseScheduleEntityImpl {
	return s.shiftItem.ScheduleEntity()
}

var _ json.Marshaler = BaseShiftItemImpl{}

func (s BaseShiftItemImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseShiftItemImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseShiftItemImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseShiftItemImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.shiftItem"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseShiftItemImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalShiftItemImplementation(input []byte) (ShiftItem, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ShiftItem into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.openShiftItem") {
		var out OpenShiftItem
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OpenShiftItem: %+v", err)
		}
		return out, nil
	}

	var parent BaseShiftItemImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseShiftItemImpl: %+v", err)
	}

	return RawShiftItemImpl{
		shiftItem: parent,
		Type:      value,
		Values:    temp,
	}, nil

}
