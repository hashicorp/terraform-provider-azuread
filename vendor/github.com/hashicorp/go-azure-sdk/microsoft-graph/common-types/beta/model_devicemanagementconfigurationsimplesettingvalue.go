package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSimpleSettingValue interface {
	DeviceManagementConfigurationSettingValue
	DeviceManagementConfigurationSimpleSettingValue() BaseDeviceManagementConfigurationSimpleSettingValueImpl
}

var _ DeviceManagementConfigurationSimpleSettingValue = BaseDeviceManagementConfigurationSimpleSettingValueImpl{}

type BaseDeviceManagementConfigurationSimpleSettingValueImpl struct {

	// Fields inherited from DeviceManagementConfigurationSettingValue

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Setting value template reference
	SettingValueTemplateReference *DeviceManagementConfigurationSettingValueTemplateReference `json:"settingValueTemplateReference,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseDeviceManagementConfigurationSimpleSettingValueImpl) DeviceManagementConfigurationSimpleSettingValue() BaseDeviceManagementConfigurationSimpleSettingValueImpl {
	return s
}

func (s BaseDeviceManagementConfigurationSimpleSettingValueImpl) DeviceManagementConfigurationSettingValue() BaseDeviceManagementConfigurationSettingValueImpl {
	return BaseDeviceManagementConfigurationSettingValueImpl{
		ODataId:                       s.ODataId,
		ODataType:                     s.ODataType,
		SettingValueTemplateReference: s.SettingValueTemplateReference,
	}
}

var _ DeviceManagementConfigurationSimpleSettingValue = RawDeviceManagementConfigurationSimpleSettingValueImpl{}

// RawDeviceManagementConfigurationSimpleSettingValueImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDeviceManagementConfigurationSimpleSettingValueImpl struct {
	deviceManagementConfigurationSimpleSettingValue BaseDeviceManagementConfigurationSimpleSettingValueImpl
	Type                                            string
	Values                                          map[string]interface{}
}

func (s RawDeviceManagementConfigurationSimpleSettingValueImpl) DeviceManagementConfigurationSimpleSettingValue() BaseDeviceManagementConfigurationSimpleSettingValueImpl {
	return s.deviceManagementConfigurationSimpleSettingValue
}

func (s RawDeviceManagementConfigurationSimpleSettingValueImpl) DeviceManagementConfigurationSettingValue() BaseDeviceManagementConfigurationSettingValueImpl {
	return s.deviceManagementConfigurationSimpleSettingValue.DeviceManagementConfigurationSettingValue()
}

var _ json.Marshaler = BaseDeviceManagementConfigurationSimpleSettingValueImpl{}

func (s BaseDeviceManagementConfigurationSimpleSettingValueImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseDeviceManagementConfigurationSimpleSettingValueImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseDeviceManagementConfigurationSimpleSettingValueImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseDeviceManagementConfigurationSimpleSettingValueImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementConfigurationSimpleSettingValue"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseDeviceManagementConfigurationSimpleSettingValueImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalDeviceManagementConfigurationSimpleSettingValueImplementation(input []byte) (DeviceManagementConfigurationSimpleSettingValue, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementConfigurationSimpleSettingValue into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationIntegerSettingValue") {
		var out DeviceManagementConfigurationIntegerSettingValue
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationIntegerSettingValue: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationSecretSettingValue") {
		var out DeviceManagementConfigurationSecretSettingValue
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationSecretSettingValue: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationStringSettingValue") {
		var out DeviceManagementConfigurationStringSettingValue
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationStringSettingValue: %+v", err)
		}
		return out, nil
	}

	var parent BaseDeviceManagementConfigurationSimpleSettingValueImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDeviceManagementConfigurationSimpleSettingValueImpl: %+v", err)
	}

	return RawDeviceManagementConfigurationSimpleSettingValueImpl{
		deviceManagementConfigurationSimpleSettingValue: parent,
		Type:   value,
		Values: temp,
	}, nil

}
