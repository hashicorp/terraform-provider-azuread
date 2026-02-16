package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationChoiceSettingValueDefaultTemplate interface {
	DeviceManagementConfigurationChoiceSettingValueDefaultTemplate() BaseDeviceManagementConfigurationChoiceSettingValueDefaultTemplateImpl
}

var _ DeviceManagementConfigurationChoiceSettingValueDefaultTemplate = BaseDeviceManagementConfigurationChoiceSettingValueDefaultTemplateImpl{}

type BaseDeviceManagementConfigurationChoiceSettingValueDefaultTemplateImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseDeviceManagementConfigurationChoiceSettingValueDefaultTemplateImpl) DeviceManagementConfigurationChoiceSettingValueDefaultTemplate() BaseDeviceManagementConfigurationChoiceSettingValueDefaultTemplateImpl {
	return s
}

var _ DeviceManagementConfigurationChoiceSettingValueDefaultTemplate = RawDeviceManagementConfigurationChoiceSettingValueDefaultTemplateImpl{}

// RawDeviceManagementConfigurationChoiceSettingValueDefaultTemplateImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDeviceManagementConfigurationChoiceSettingValueDefaultTemplateImpl struct {
	deviceManagementConfigurationChoiceSettingValueDefaultTemplate BaseDeviceManagementConfigurationChoiceSettingValueDefaultTemplateImpl
	Type                                                           string
	Values                                                         map[string]interface{}
}

func (s RawDeviceManagementConfigurationChoiceSettingValueDefaultTemplateImpl) DeviceManagementConfigurationChoiceSettingValueDefaultTemplate() BaseDeviceManagementConfigurationChoiceSettingValueDefaultTemplateImpl {
	return s.deviceManagementConfigurationChoiceSettingValueDefaultTemplate
}

func UnmarshalDeviceManagementConfigurationChoiceSettingValueDefaultTemplateImplementation(input []byte) (DeviceManagementConfigurationChoiceSettingValueDefaultTemplate, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementConfigurationChoiceSettingValueDefaultTemplate into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationChoiceSettingValueConstantDefaultTemplate") {
		var out DeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplate
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplate: %+v", err)
		}
		return out, nil
	}

	var parent BaseDeviceManagementConfigurationChoiceSettingValueDefaultTemplateImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDeviceManagementConfigurationChoiceSettingValueDefaultTemplateImpl: %+v", err)
	}

	return RawDeviceManagementConfigurationChoiceSettingValueDefaultTemplateImpl{
		deviceManagementConfigurationChoiceSettingValueDefaultTemplate: parent,
		Type:   value,
		Values: temp,
	}, nil

}
