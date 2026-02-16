package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSettingValue interface {
	DeviceManagementConfigurationSettingValue() BaseDeviceManagementConfigurationSettingValueImpl
}

var _ DeviceManagementConfigurationSettingValue = BaseDeviceManagementConfigurationSettingValueImpl{}

type BaseDeviceManagementConfigurationSettingValueImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Setting value template reference
	SettingValueTemplateReference *DeviceManagementConfigurationSettingValueTemplateReference `json:"settingValueTemplateReference,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseDeviceManagementConfigurationSettingValueImpl) DeviceManagementConfigurationSettingValue() BaseDeviceManagementConfigurationSettingValueImpl {
	return s
}

var _ DeviceManagementConfigurationSettingValue = RawDeviceManagementConfigurationSettingValueImpl{}

// RawDeviceManagementConfigurationSettingValueImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDeviceManagementConfigurationSettingValueImpl struct {
	deviceManagementConfigurationSettingValue BaseDeviceManagementConfigurationSettingValueImpl
	Type                                      string
	Values                                    map[string]interface{}
}

func (s RawDeviceManagementConfigurationSettingValueImpl) DeviceManagementConfigurationSettingValue() BaseDeviceManagementConfigurationSettingValueImpl {
	return s.deviceManagementConfigurationSettingValue
}

func UnmarshalDeviceManagementConfigurationSettingValueImplementation(input []byte) (DeviceManagementConfigurationSettingValue, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementConfigurationSettingValue into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationChoiceSettingValue") {
		var out DeviceManagementConfigurationChoiceSettingValue
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationChoiceSettingValue: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationGroupSettingValue") {
		var out DeviceManagementConfigurationGroupSettingValue
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationGroupSettingValue: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationSimpleSettingValue") {
		var out DeviceManagementConfigurationSimpleSettingValue
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationSimpleSettingValue: %+v", err)
		}
		return out, nil
	}

	var parent BaseDeviceManagementConfigurationSettingValueImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDeviceManagementConfigurationSettingValueImpl: %+v", err)
	}

	return RawDeviceManagementConfigurationSettingValueImpl{
		deviceManagementConfigurationSettingValue: parent,
		Type:   value,
		Values: temp,
	}, nil

}
