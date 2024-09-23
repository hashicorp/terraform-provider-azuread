package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnAttributeCollectionHandler interface {
	OnAttributeCollectionHandler() BaseOnAttributeCollectionHandlerImpl
}

var _ OnAttributeCollectionHandler = BaseOnAttributeCollectionHandlerImpl{}

type BaseOnAttributeCollectionHandlerImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseOnAttributeCollectionHandlerImpl) OnAttributeCollectionHandler() BaseOnAttributeCollectionHandlerImpl {
	return s
}

var _ OnAttributeCollectionHandler = RawOnAttributeCollectionHandlerImpl{}

// RawOnAttributeCollectionHandlerImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawOnAttributeCollectionHandlerImpl struct {
	onAttributeCollectionHandler BaseOnAttributeCollectionHandlerImpl
	Type                         string
	Values                       map[string]interface{}
}

func (s RawOnAttributeCollectionHandlerImpl) OnAttributeCollectionHandler() BaseOnAttributeCollectionHandlerImpl {
	return s.onAttributeCollectionHandler
}

func UnmarshalOnAttributeCollectionHandlerImplementation(input []byte) (OnAttributeCollectionHandler, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling OnAttributeCollectionHandler into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.onAttributeCollectionExternalUsersSelfServiceSignUp") {
		var out OnAttributeCollectionExternalUsersSelfServiceSignUp
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnAttributeCollectionExternalUsersSelfServiceSignUp: %+v", err)
		}
		return out, nil
	}

	var parent BaseOnAttributeCollectionHandlerImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseOnAttributeCollectionHandlerImpl: %+v", err)
	}

	return RawOnAttributeCollectionHandlerImpl{
		onAttributeCollectionHandler: parent,
		Type:                         value,
		Values:                       temp,
	}, nil

}
