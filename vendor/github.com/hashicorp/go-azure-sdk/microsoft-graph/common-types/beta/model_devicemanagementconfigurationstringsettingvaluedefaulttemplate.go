package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationStringSettingValueDefaultTemplate interface {
	DeviceManagementConfigurationStringSettingValueDefaultTemplate() BaseDeviceManagementConfigurationStringSettingValueDefaultTemplateImpl
}

var _ DeviceManagementConfigurationStringSettingValueDefaultTemplate = BaseDeviceManagementConfigurationStringSettingValueDefaultTemplateImpl{}

type BaseDeviceManagementConfigurationStringSettingValueDefaultTemplateImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseDeviceManagementConfigurationStringSettingValueDefaultTemplateImpl) DeviceManagementConfigurationStringSettingValueDefaultTemplate() BaseDeviceManagementConfigurationStringSettingValueDefaultTemplateImpl {
	return s
}

var _ DeviceManagementConfigurationStringSettingValueDefaultTemplate = RawDeviceManagementConfigurationStringSettingValueDefaultTemplateImpl{}

// RawDeviceManagementConfigurationStringSettingValueDefaultTemplateImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDeviceManagementConfigurationStringSettingValueDefaultTemplateImpl struct {
	deviceManagementConfigurationStringSettingValueDefaultTemplate BaseDeviceManagementConfigurationStringSettingValueDefaultTemplateImpl
	Type                                                           string
	Values                                                         map[string]interface{}
}

func (s RawDeviceManagementConfigurationStringSettingValueDefaultTemplateImpl) DeviceManagementConfigurationStringSettingValueDefaultTemplate() BaseDeviceManagementConfigurationStringSettingValueDefaultTemplateImpl {
	return s.deviceManagementConfigurationStringSettingValueDefaultTemplate
}

func UnmarshalDeviceManagementConfigurationStringSettingValueDefaultTemplateImplementation(input []byte) (DeviceManagementConfigurationStringSettingValueDefaultTemplate, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementConfigurationStringSettingValueDefaultTemplate into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationStringSettingValueConstantDefaultTemplate") {
		var out DeviceManagementConfigurationStringSettingValueConstantDefaultTemplate
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationStringSettingValueConstantDefaultTemplate: %+v", err)
		}
		return out, nil
	}

	var parent BaseDeviceManagementConfigurationStringSettingValueDefaultTemplateImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDeviceManagementConfigurationStringSettingValueDefaultTemplateImpl: %+v", err)
	}

	return RawDeviceManagementConfigurationStringSettingValueDefaultTemplateImpl{
		deviceManagementConfigurationStringSettingValueDefaultTemplate: parent,
		Type:   value,
		Values: temp,
	}, nil

}
