package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
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

// RawOnAttributeCollectionStartHandlerImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawOnAttributeCollectionStartHandlerImpl struct {
	onAttributeCollectionStartHandler BaseOnAttributeCollectionStartHandlerImpl
	Type                              string
	Values                            map[string]interface{}
}

func (s RawOnAttributeCollectionStartHandlerImpl) OnAttributeCollectionStartHandler() BaseOnAttributeCollectionStartHandlerImpl {
	return s.onAttributeCollectionStartHandler
}

func (s RawOnAttributeCollectionStartHandlerImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
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
