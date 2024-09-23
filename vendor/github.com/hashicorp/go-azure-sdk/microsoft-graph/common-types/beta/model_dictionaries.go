package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Dictionaries interface {
	Dictionaries() BaseDictionariesImpl
}

var _ Dictionaries = BaseDictionariesImpl{}

type BaseDictionariesImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseDictionariesImpl) Dictionaries() BaseDictionariesImpl {
	return s
}

var _ Dictionaries = RawDictionariesImpl{}

// RawDictionariesImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDictionariesImpl struct {
	dictionaries BaseDictionariesImpl
	Type         string
	Values       map[string]interface{}
}

func (s RawDictionariesImpl) Dictionaries() BaseDictionariesImpl {
	return s.dictionaries
}

func UnmarshalDictionariesImplementation(input []byte) (Dictionaries, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling Dictionaries into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.searchResourceMetadataDictionary") {
		var out SearchResourceMetadataDictionary
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SearchResourceMetadataDictionary: %+v", err)
		}
		return out, nil
	}

	var parent BaseDictionariesImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDictionariesImpl: %+v", err)
	}

	return RawDictionariesImpl{
		dictionaries: parent,
		Type:         value,
		Values:       temp,
	}, nil

}
