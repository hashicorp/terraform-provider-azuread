package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnAttributeCollectionSubmitHandler interface {
	OnAttributeCollectionSubmitHandler() BaseOnAttributeCollectionSubmitHandlerImpl
}

var _ OnAttributeCollectionSubmitHandler = BaseOnAttributeCollectionSubmitHandlerImpl{}

type BaseOnAttributeCollectionSubmitHandlerImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseOnAttributeCollectionSubmitHandlerImpl) OnAttributeCollectionSubmitHandler() BaseOnAttributeCollectionSubmitHandlerImpl {
	return s
}

var _ OnAttributeCollectionSubmitHandler = RawOnAttributeCollectionSubmitHandlerImpl{}

// RawOnAttributeCollectionSubmitHandlerImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawOnAttributeCollectionSubmitHandlerImpl struct {
	onAttributeCollectionSubmitHandler BaseOnAttributeCollectionSubmitHandlerImpl
	Type                               string
	Values                             map[string]interface{}
}

func (s RawOnAttributeCollectionSubmitHandlerImpl) OnAttributeCollectionSubmitHandler() BaseOnAttributeCollectionSubmitHandlerImpl {
	return s.onAttributeCollectionSubmitHandler
}

func (s RawOnAttributeCollectionSubmitHandlerImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalOnAttributeCollectionSubmitHandlerImplementation(input []byte) (OnAttributeCollectionSubmitHandler, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling OnAttributeCollectionSubmitHandler into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.onAttributeCollectionSubmitCustomExtensionHandler") {
		var out OnAttributeCollectionSubmitCustomExtensionHandler
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnAttributeCollectionSubmitCustomExtensionHandler: %+v", err)
		}
		return out, nil
	}

	var parent BaseOnAttributeCollectionSubmitHandlerImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseOnAttributeCollectionSubmitHandlerImpl: %+v", err)
	}

	return RawOnAttributeCollectionSubmitHandlerImpl{
		onAttributeCollectionSubmitHandler: parent,
		Type:                               value,
		Values:                             temp,
	}, nil

}
