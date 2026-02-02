package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnAttributeCollectionStartHandler interface {
	OnAttributeCollectionStartHandler() BaseOnAttributeCollectionStartHandlerImpl
}

var _ OnAttributeCollectionStartHandler = BaseOnAttributeCollectionStartHandlerImpl{}

type BaseOnAttributeCollectionStartHandlerImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseOnAttributeCollectionStartHandlerImpl) OnAttributeCollectionStartHandler() BaseOnAttributeCollectionStartHandlerImpl {
	return s
}

var _ OnAttributeCollectionStartHandler = RawOnAttributeCollectionStartHandlerImpl{}

// RawOnAttributeCollectionStartHandlerImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawOnAttributeCollectionStartHandlerImpl struct {
	onAttributeCollectionStartHandler BaseOnAttributeCollectionStartHandlerImpl
	Type                              string
	Values                            map[string]interface{}
}

func (s RawOnAttributeCollectionStartHandlerImpl) OnAttributeCollectionStartHandler() BaseOnAttributeCollectionStartHandlerImpl {
	return s.onAttributeCollectionStartHandler
}

func UnmarshalOnAttributeCollectionStartHandlerImplementation(input []byte) (OnAttributeCollectionStartHandler, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling OnAttributeCollectionStartHandler into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.onAttributeCollectionStartCustomExtensionHandler") {
		var out OnAttributeCollectionStartCustomExtensionHandler
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnAttributeCollectionStartCustomExtensionHandler: %+v", err)
		}
		return out, nil
	}

	var parent BaseOnAttributeCollectionStartHandlerImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseOnAttributeCollectionStartHandlerImpl: %+v", err)
	}

	return RawOnAttributeCollectionStartHandlerImpl{
		onAttributeCollectionStartHandler: parent,
		Type:                              value,
		Values:                            temp,
	}, nil

}
