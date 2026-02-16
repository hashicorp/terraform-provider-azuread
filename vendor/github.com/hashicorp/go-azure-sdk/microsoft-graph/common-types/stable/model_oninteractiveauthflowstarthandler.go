package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnInteractiveAuthFlowStartHandler interface {
	OnInteractiveAuthFlowStartHandler() BaseOnInteractiveAuthFlowStartHandlerImpl
}

var _ OnInteractiveAuthFlowStartHandler = BaseOnInteractiveAuthFlowStartHandlerImpl{}

type BaseOnInteractiveAuthFlowStartHandlerImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseOnInteractiveAuthFlowStartHandlerImpl) OnInteractiveAuthFlowStartHandler() BaseOnInteractiveAuthFlowStartHandlerImpl {
	return s
}

var _ OnInteractiveAuthFlowStartHandler = RawOnInteractiveAuthFlowStartHandlerImpl{}

// RawOnInteractiveAuthFlowStartHandlerImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawOnInteractiveAuthFlowStartHandlerImpl struct {
	onInteractiveAuthFlowStartHandler BaseOnInteractiveAuthFlowStartHandlerImpl
	Type                              string
	Values                            map[string]interface{}
}

func (s RawOnInteractiveAuthFlowStartHandlerImpl) OnInteractiveAuthFlowStartHandler() BaseOnInteractiveAuthFlowStartHandlerImpl {
	return s.onInteractiveAuthFlowStartHandler
}

func UnmarshalOnInteractiveAuthFlowStartHandlerImplementation(input []byte) (OnInteractiveAuthFlowStartHandler, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling OnInteractiveAuthFlowStartHandler into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.onInteractiveAuthFlowStartExternalUsersSelfServiceSignUp") {
		var out OnInteractiveAuthFlowStartExternalUsersSelfServiceSignUp
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnInteractiveAuthFlowStartExternalUsersSelfServiceSignUp: %+v", err)
		}
		return out, nil
	}

	var parent BaseOnInteractiveAuthFlowStartHandlerImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseOnInteractiveAuthFlowStartHandlerImpl: %+v", err)
	}

	return RawOnInteractiveAuthFlowStartHandlerImpl{
		onInteractiveAuthFlowStartHandler: parent,
		Type:                              value,
		Values:                            temp,
	}, nil

}
