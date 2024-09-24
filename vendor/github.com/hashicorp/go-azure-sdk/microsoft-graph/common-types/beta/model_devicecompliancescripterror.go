package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceScriptError interface {
	DeviceComplianceScriptError() BaseDeviceComplianceScriptErrorImpl
}

var _ DeviceComplianceScriptError = BaseDeviceComplianceScriptErrorImpl{}

type BaseDeviceComplianceScriptErrorImpl struct {
	// Error code for rule validation.
	Code *Code `json:"code,omitempty"`

	// Error code for rule validation.
	DeviceComplianceScriptRulesValidationError *DeviceComplianceScriptRulesValidationError `json:"deviceComplianceScriptRulesValidationError,omitempty"`

	// Error message.
	Message nullable.Type[string] `json:"message,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseDeviceComplianceScriptErrorImpl) DeviceComplianceScriptError() BaseDeviceComplianceScriptErrorImpl {
	return s
}

var _ DeviceComplianceScriptError = RawDeviceComplianceScriptErrorImpl{}

// RawDeviceComplianceScriptErrorImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDeviceComplianceScriptErrorImpl struct {
	deviceComplianceScriptError BaseDeviceComplianceScriptErrorImpl
	Type                        string
	Values                      map[string]interface{}
}

func (s RawDeviceComplianceScriptErrorImpl) DeviceComplianceScriptError() BaseDeviceComplianceScriptErrorImpl {
	return s.deviceComplianceScriptError
}

func UnmarshalDeviceComplianceScriptErrorImplementation(input []byte) (DeviceComplianceScriptError, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceComplianceScriptError into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceComplianceScriptRuleError") {
		var out DeviceComplianceScriptRuleError
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceComplianceScriptRuleError: %+v", err)
		}
		return out, nil
	}

	var parent BaseDeviceComplianceScriptErrorImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDeviceComplianceScriptErrorImpl: %+v", err)
	}

	return RawDeviceComplianceScriptErrorImpl{
		deviceComplianceScriptError: parent,
		Type:                        value,
		Values:                      temp,
	}, nil

}
