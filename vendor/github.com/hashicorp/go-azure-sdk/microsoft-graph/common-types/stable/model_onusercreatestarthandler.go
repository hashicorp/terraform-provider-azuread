package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnUserCreateStartHandler interface {
	OnUserCreateStartHandler() BaseOnUserCreateStartHandlerImpl
}

var _ OnUserCreateStartHandler = BaseOnUserCreateStartHandlerImpl{}

type BaseOnUserCreateStartHandlerImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseOnUserCreateStartHandlerImpl) OnUserCreateStartHandler() BaseOnUserCreateStartHandlerImpl {
	return s
}

var _ OnUserCreateStartHandler = RawOnUserCreateStartHandlerImpl{}

// RawOnUserCreateStartHandlerImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawOnUserCreateStartHandlerImpl struct {
	onUserCreateStartHandler BaseOnUserCreateStartHandlerImpl
	Type                     string
	Values                   map[string]interface{}
}

func (s RawOnUserCreateStartHandlerImpl) OnUserCreateStartHandler() BaseOnUserCreateStartHandlerImpl {
	return s.onUserCreateStartHandler
}

func UnmarshalOnUserCreateStartHandlerImplementation(input []byte) (OnUserCreateStartHandler, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling OnUserCreateStartHandler into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.onUserCreateStartExternalUsersSelfServiceSignUp") {
		var out OnUserCreateStartExternalUsersSelfServiceSignUp
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnUserCreateStartExternalUsersSelfServiceSignUp: %+v", err)
		}
		return out, nil
	}

	var parent BaseOnUserCreateStartHandlerImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseOnUserCreateStartHandlerImpl: %+v", err)
	}

	return RawOnUserCreateStartHandlerImpl{
		onUserCreateStartHandler: parent,
		Type:                     value,
		Values:                   temp,
	}, nil

}
