package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationSegment interface {
	Entity
	ApplicationSegment() BaseApplicationSegmentImpl
}

var _ ApplicationSegment = BaseApplicationSegmentImpl{}

type BaseApplicationSegmentImpl struct {

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

func (s BaseApplicationSegmentImpl) ApplicationSegment() BaseApplicationSegmentImpl {
	return s
}

func (s BaseApplicationSegmentImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ ApplicationSegment = RawApplicationSegmentImpl{}

// RawApplicationSegmentImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawApplicationSegmentImpl struct {
	applicationSegment BaseApplicationSegmentImpl
	Type               string
	Values             map[string]interface{}
}

func (s RawApplicationSegmentImpl) ApplicationSegment() BaseApplicationSegmentImpl {
	return s.applicationSegment
}

func (s RawApplicationSegmentImpl) Entity() BaseEntityImpl {
	return s.applicationSegment.Entity()
}

var _ json.Marshaler = BaseApplicationSegmentImpl{}

func (s BaseApplicationSegmentImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseApplicationSegmentImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseApplicationSegmentImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseApplicationSegmentImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.applicationSegment"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseApplicationSegmentImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalApplicationSegmentImplementation(input []byte) (ApplicationSegment, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ApplicationSegment into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.ipApplicationSegment") {
		var out IPApplicationSegment
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IPApplicationSegment: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.webApplicationSegment") {
		var out WebApplicationSegment
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WebApplicationSegment: %+v", err)
		}
		return out, nil
	}

	var parent BaseApplicationSegmentImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseApplicationSegmentImpl: %+v", err)
	}

	return RawApplicationSegmentImpl{
		applicationSegment: parent,
		Type:               value,
		Values:             temp,
	}, nil

}
