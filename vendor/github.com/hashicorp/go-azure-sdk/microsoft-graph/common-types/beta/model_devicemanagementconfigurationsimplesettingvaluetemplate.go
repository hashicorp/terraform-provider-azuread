package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSimpleSettingValueTemplate interface {
	DeviceManagementConfigurationSimpleSettingValueTemplate() BaseDeviceManagementConfigurationSimpleSettingValueTemplateImpl
}

var _ DeviceManagementConfigurationSimpleSettingValueTemplate = BaseDeviceManagementConfigurationSimpleSettingValueTemplateImpl{}

type BaseDeviceManagementConfigurationSimpleSettingValueTemplateImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Setting Value Template Id
	SettingValueTemplateId *string `json:"settingValueTemplateId,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseDeviceManagementConfigurationSimpleSettingValueTemplateImpl) DeviceManagementConfigurationSimpleSettingValueTemplate() BaseDeviceManagementConfigurationSimpleSettingValueTemplateImpl {
	return s
}

var _ DeviceManagementConfigurationSimpleSettingValueTemplate = RawDeviceManagementConfigurationSimpleSettingValueTemplateImpl{}

// RawDeviceManagementConfigurationSimpleSettingValueTemplateImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDeviceManagementConfigurationSimpleSettingValueTemplateImpl struct {
	deviceManagementConfigurationSimpleSettingValueTemplate BaseDeviceManagementConfigurationSimpleSettingValueTemplateImpl
	Type                                                    string
	Values                                                  map[string]interface{}
}

func (s RawDeviceManagementConfigurationSimpleSettingValueTemplateImpl) DeviceManagementConfigurationSimpleSettingValueTemplate() BaseDeviceManagementConfigurationSimpleSettingValueTemplateImpl {
	return s.deviceManagementConfigurationSimpleSettingValueTemplate
}

func UnmarshalDeviceManagementConfigurationSimpleSettingValueTemplateImplementation(input []byte) (DeviceManagementConfigurationSimpleSettingValueTemplate, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementConfigurationSimpleSettingValueTemplate into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationIntegerSettingValueTemplate") {
		var out DeviceManagementConfigurationIntegerSettingValueTemplate
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationIntegerSettingValueTemplate: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationStringSettingValueTemplate") {
		var out DeviceManagementConfigurationStringSettingValueTemplate
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationStringSettingValueTemplate: %+v", err)
		}
		return out, nil
	}

	var parent BaseDeviceManagementConfigurationSimpleSettingValueTemplateImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDeviceManagementConfigurationSimpleSettingValueTemplateImpl: %+v", err)
	}

	return RawDeviceManagementConfigurationSimpleSettingValueTemplateImpl{
		deviceManagementConfigurationSimpleSettingValueTemplate: parent,
		Type:   value,
		Values: temp,
	}, nil

}
