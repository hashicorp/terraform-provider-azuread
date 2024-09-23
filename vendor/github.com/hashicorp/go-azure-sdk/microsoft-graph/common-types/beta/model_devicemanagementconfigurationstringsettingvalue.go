package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationStringSettingValue interface {
	DeviceManagementConfigurationSettingValue
	DeviceManagementConfigurationSimpleSettingValue
	DeviceManagementConfigurationStringSettingValue() BaseDeviceManagementConfigurationStringSettingValueImpl
}

var _ DeviceManagementConfigurationStringSettingValue = BaseDeviceManagementConfigurationStringSettingValueImpl{}

type BaseDeviceManagementConfigurationStringSettingValueImpl struct {
	// Value of the string setting.
	Value nullable.Type[string] `json:"value,omitempty"`

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

func (s BaseDeviceManagementConfigurationStringSettingValueImpl) DeviceManagementConfigurationStringSettingValue() BaseDeviceManagementConfigurationStringSettingValueImpl {
	return s
}

func (s BaseDeviceManagementConfigurationStringSettingValueImpl) DeviceManagementConfigurationSimpleSettingValue() BaseDeviceManagementConfigurationSimpleSettingValueImpl {
	return BaseDeviceManagementConfigurationSimpleSettingValueImpl{
		ODataId:                       s.ODataId,
		ODataType:                     s.ODataType,
		SettingValueTemplateReference: s.SettingValueTemplateReference,
	}
}

func (s BaseDeviceManagementConfigurationStringSettingValueImpl) DeviceManagementConfigurationSettingValue() BaseDeviceManagementConfigurationSettingValueImpl {
	return BaseDeviceManagementConfigurationSettingValueImpl{
		ODataId:                       s.ODataId,
		ODataType:                     s.ODataType,
		SettingValueTemplateReference: s.SettingValueTemplateReference,
	}
}

var _ DeviceManagementConfigurationStringSettingValue = RawDeviceManagementConfigurationStringSettingValueImpl{}

// RawDeviceManagementConfigurationStringSettingValueImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDeviceManagementConfigurationStringSettingValueImpl struct {
	deviceManagementConfigurationStringSettingValue BaseDeviceManagementConfigurationStringSettingValueImpl
	Type                                            string
	Values                                          map[string]interface{}
}

func (s RawDeviceManagementConfigurationStringSettingValueImpl) DeviceManagementConfigurationStringSettingValue() BaseDeviceManagementConfigurationStringSettingValueImpl {
	return s.deviceManagementConfigurationStringSettingValue
}

func (s RawDeviceManagementConfigurationStringSettingValueImpl) DeviceManagementConfigurationSimpleSettingValue() BaseDeviceManagementConfigurationSimpleSettingValueImpl {
	return s.deviceManagementConfigurationStringSettingValue.DeviceManagementConfigurationSimpleSettingValue()
}

func (s RawDeviceManagementConfigurationStringSettingValueImpl) DeviceManagementConfigurationSettingValue() BaseDeviceManagementConfigurationSettingValueImpl {
	return s.deviceManagementConfigurationStringSettingValue.DeviceManagementConfigurationSettingValue()
}

var _ json.Marshaler = BaseDeviceManagementConfigurationStringSettingValueImpl{}

func (s BaseDeviceManagementConfigurationStringSettingValueImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseDeviceManagementConfigurationStringSettingValueImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseDeviceManagementConfigurationStringSettingValueImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseDeviceManagementConfigurationStringSettingValueImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementConfigurationStringSettingValue"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseDeviceManagementConfigurationStringSettingValueImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalDeviceManagementConfigurationStringSettingValueImplementation(input []byte) (DeviceManagementConfigurationStringSettingValue, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementConfigurationStringSettingValue into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationReferenceSettingValue") {
		var out DeviceManagementConfigurationReferenceSettingValue
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationReferenceSettingValue: %+v", err)
		}
		return out, nil
	}

	var parent BaseDeviceManagementConfigurationStringSettingValueImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDeviceManagementConfigurationStringSettingValueImpl: %+v", err)
	}

	return RawDeviceManagementConfigurationStringSettingValueImpl{
		deviceManagementConfigurationStringSettingValue: parent,
		Type:   value,
		Values: temp,
	}, nil

}
