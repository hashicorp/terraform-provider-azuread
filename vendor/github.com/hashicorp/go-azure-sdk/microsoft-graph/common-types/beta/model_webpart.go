package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebPart interface {
	Entity
	WebPart() BaseWebPartImpl
}

var _ WebPart = BaseWebPartImpl{}

type BaseWebPartImpl struct {

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

func (s BaseWebPartImpl) WebPart() BaseWebPartImpl {
	return s
}

func (s BaseWebPartImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ WebPart = RawWebPartImpl{}

// RawWebPartImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawWebPartImpl struct {
	webPart BaseWebPartImpl
	Type    string
	Values  map[string]interface{}
}

func (s RawWebPartImpl) WebPart() BaseWebPartImpl {
	return s.webPart
}

func (s RawWebPartImpl) Entity() BaseEntityImpl {
	return s.webPart.Entity()
}

var _ json.Marshaler = BaseWebPartImpl{}

func (s BaseWebPartImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseWebPartImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseWebPartImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseWebPartImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.webPart"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseWebPartImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalWebPartImplementation(input []byte) (WebPart, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling WebPart into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.standardWebPart") {
		var out StandardWebPart
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into StandardWebPart: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.textWebPart") {
		var out TextWebPart
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TextWebPart: %+v", err)
		}
		return out, nil
	}

	var parent BaseWebPartImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseWebPartImpl: %+v", err)
	}

	return RawWebPartImpl{
		webPart: parent,
		Type:    value,
		Values:  temp,
	}, nil

}
