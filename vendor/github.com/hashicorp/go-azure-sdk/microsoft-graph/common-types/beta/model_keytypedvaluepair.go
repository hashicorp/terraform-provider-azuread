package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KeyTypedValuePair interface {
	KeyTypedValuePair() BaseKeyTypedValuePairImpl
}

var _ KeyTypedValuePair = BaseKeyTypedValuePairImpl{}

type BaseKeyTypedValuePairImpl struct {
	// The string key of the key-value pair.
	Key *string `json:"key,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseKeyTypedValuePairImpl) KeyTypedValuePair() BaseKeyTypedValuePairImpl {
	return s
}

var _ KeyTypedValuePair = RawKeyTypedValuePairImpl{}

// RawKeyTypedValuePairImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawKeyTypedValuePairImpl struct {
	keyTypedValuePair BaseKeyTypedValuePairImpl
	Type              string
	Values            map[string]interface{}
}

func (s RawKeyTypedValuePairImpl) KeyTypedValuePair() BaseKeyTypedValuePairImpl {
	return s.keyTypedValuePair
}

func UnmarshalKeyTypedValuePairImplementation(input []byte) (KeyTypedValuePair, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling KeyTypedValuePair into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.keyBooleanValuePair") {
		var out KeyBooleanValuePair
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into KeyBooleanValuePair: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.keyIntegerValuePair") {
		var out KeyIntegerValuePair
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into KeyIntegerValuePair: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.keyRealValuePair") {
		var out KeyRealValuePair
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into KeyRealValuePair: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.keyStringValuePair") {
		var out KeyStringValuePair
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into KeyStringValuePair: %+v", err)
		}
		return out, nil
	}

	var parent BaseKeyTypedValuePairImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseKeyTypedValuePairImpl: %+v", err)
	}

	return RawKeyTypedValuePairImpl{
		keyTypedValuePair: parent,
		Type:              value,
		Values:            temp,
	}, nil

}
