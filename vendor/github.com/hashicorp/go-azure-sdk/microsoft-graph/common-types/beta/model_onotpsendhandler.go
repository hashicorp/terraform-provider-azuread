package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnOtpSendHandler interface {
	OnOtpSendHandler() BaseOnOtpSendHandlerImpl
}

var _ OnOtpSendHandler = BaseOnOtpSendHandlerImpl{}

type BaseOnOtpSendHandlerImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseOnOtpSendHandlerImpl) OnOtpSendHandler() BaseOnOtpSendHandlerImpl {
	return s
}

var _ OnOtpSendHandler = RawOnOtpSendHandlerImpl{}

// RawOnOtpSendHandlerImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawOnOtpSendHandlerImpl struct {
	onOtpSendHandler BaseOnOtpSendHandlerImpl
	Type             string
	Values           map[string]interface{}
}

func (s RawOnOtpSendHandlerImpl) OnOtpSendHandler() BaseOnOtpSendHandlerImpl {
	return s.onOtpSendHandler
}

func UnmarshalOnOtpSendHandlerImplementation(input []byte) (OnOtpSendHandler, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling OnOtpSendHandler into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.onOtpSendCustomExtensionHandler") {
		var out OnOtpSendCustomExtensionHandler
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnOtpSendCustomExtensionHandler: %+v", err)
		}
		return out, nil
	}

	var parent BaseOnOtpSendHandlerImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseOnOtpSendHandlerImpl: %+v", err)
	}

	return RawOnOtpSendHandlerImpl{
		onOtpSendHandler: parent,
		Type:             value,
		Values:           temp,
	}, nil

}
