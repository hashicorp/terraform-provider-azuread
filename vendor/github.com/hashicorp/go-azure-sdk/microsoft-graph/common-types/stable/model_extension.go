package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Extension interface {
	Entity
	Extension() BaseExtensionImpl
}

var _ Extension = BaseExtensionImpl{}

type BaseExtensionImpl struct {

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

func (s BaseExtensionImpl) Extension() BaseExtensionImpl {
	return s
}

func (s BaseExtensionImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ Extension = RawExtensionImpl{}

// RawExtensionImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawExtensionImpl struct {
	extension BaseExtensionImpl
	Type      string
	Values    map[string]interface{}
}

func (s RawExtensionImpl) Extension() BaseExtensionImpl {
	return s.extension
}

func (s RawExtensionImpl) Entity() BaseEntityImpl {
	return s.extension.Entity()
}

var _ json.Marshaler = BaseExtensionImpl{}

func (s BaseExtensionImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseExtensionImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseExtensionImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseExtensionImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.extension"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseExtensionImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalExtensionImplementation(input []byte) (Extension, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling Extension into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.openTypeExtension") {
		var out OpenTypeExtension
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OpenTypeExtension: %+v", err)
		}
		return out, nil
	}

	var parent BaseExtensionImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseExtensionImpl: %+v", err)
	}

	return RawExtensionImpl{
		extension: parent,
		Type:      value,
		Values:    temp,
	}, nil

}
