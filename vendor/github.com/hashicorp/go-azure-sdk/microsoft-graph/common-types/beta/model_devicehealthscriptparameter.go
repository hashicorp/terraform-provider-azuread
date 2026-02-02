package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceHealthScriptParameter interface {
	DeviceHealthScriptParameter() BaseDeviceHealthScriptParameterImpl
}

var _ DeviceHealthScriptParameter = BaseDeviceHealthScriptParameterImpl{}

type BaseDeviceHealthScriptParameterImpl struct {
	// Whether Apply DefaultValue When Not Assigned
	ApplyDefaultValueWhenNotAssigned *bool `json:"applyDefaultValueWhenNotAssigned,omitempty"`

	// The description of the param
	Description nullable.Type[string] `json:"description,omitempty"`

	// Whether the param is required
	IsRequired *bool `json:"isRequired,omitempty"`

	// The name of the param
	Name nullable.Type[string] `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseDeviceHealthScriptParameterImpl) DeviceHealthScriptParameter() BaseDeviceHealthScriptParameterImpl {
	return s
}

var _ DeviceHealthScriptParameter = RawDeviceHealthScriptParameterImpl{}

// RawDeviceHealthScriptParameterImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDeviceHealthScriptParameterImpl struct {
	deviceHealthScriptParameter BaseDeviceHealthScriptParameterImpl
	Type                        string
	Values                      map[string]interface{}
}

func (s RawDeviceHealthScriptParameterImpl) DeviceHealthScriptParameter() BaseDeviceHealthScriptParameterImpl {
	return s.deviceHealthScriptParameter
}

func UnmarshalDeviceHealthScriptParameterImplementation(input []byte) (DeviceHealthScriptParameter, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceHealthScriptParameter into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceHealthScriptBooleanParameter") {
		var out DeviceHealthScriptBooleanParameter
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceHealthScriptBooleanParameter: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceHealthScriptIntegerParameter") {
		var out DeviceHealthScriptIntegerParameter
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceHealthScriptIntegerParameter: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceHealthScriptStringParameter") {
		var out DeviceHealthScriptStringParameter
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceHealthScriptStringParameter: %+v", err)
		}
		return out, nil
	}

	var parent BaseDeviceHealthScriptParameterImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDeviceHealthScriptParameterImpl: %+v", err)
	}

	return RawDeviceHealthScriptParameterImpl{
		deviceHealthScriptParameter: parent,
		Type:                        value,
		Values:                      temp,
	}, nil

}
