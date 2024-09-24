package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSettingInstanceTemplate interface {
	DeviceManagementConfigurationSettingInstanceTemplate() BaseDeviceManagementConfigurationSettingInstanceTemplateImpl
}

var _ DeviceManagementConfigurationSettingInstanceTemplate = BaseDeviceManagementConfigurationSettingInstanceTemplateImpl{}

type BaseDeviceManagementConfigurationSettingInstanceTemplateImpl struct {
	// Indicates if a policy must specify this setting.
	IsRequired *bool `json:"isRequired,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Setting Definition Id
	SettingDefinitionId *string `json:"settingDefinitionId,omitempty"`

	// Setting Instance Template Id
	SettingInstanceTemplateId *string `json:"settingInstanceTemplateId,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseDeviceManagementConfigurationSettingInstanceTemplateImpl) DeviceManagementConfigurationSettingInstanceTemplate() BaseDeviceManagementConfigurationSettingInstanceTemplateImpl {
	return s
}

var _ DeviceManagementConfigurationSettingInstanceTemplate = RawDeviceManagementConfigurationSettingInstanceTemplateImpl{}

// RawDeviceManagementConfigurationSettingInstanceTemplateImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDeviceManagementConfigurationSettingInstanceTemplateImpl struct {
	deviceManagementConfigurationSettingInstanceTemplate BaseDeviceManagementConfigurationSettingInstanceTemplateImpl
	Type                                                 string
	Values                                               map[string]interface{}
}

func (s RawDeviceManagementConfigurationSettingInstanceTemplateImpl) DeviceManagementConfigurationSettingInstanceTemplate() BaseDeviceManagementConfigurationSettingInstanceTemplateImpl {
	return s.deviceManagementConfigurationSettingInstanceTemplate
}

func UnmarshalDeviceManagementConfigurationSettingInstanceTemplateImplementation(input []byte) (DeviceManagementConfigurationSettingInstanceTemplate, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementConfigurationSettingInstanceTemplate into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationChoiceSettingCollectionInstanceTemplate") {
		var out DeviceManagementConfigurationChoiceSettingCollectionInstanceTemplate
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationChoiceSettingCollectionInstanceTemplate: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationChoiceSettingInstanceTemplate") {
		var out DeviceManagementConfigurationChoiceSettingInstanceTemplate
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationChoiceSettingInstanceTemplate: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationGroupSettingCollectionInstanceTemplate") {
		var out DeviceManagementConfigurationGroupSettingCollectionInstanceTemplate
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationGroupSettingCollectionInstanceTemplate: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationGroupSettingInstanceTemplate") {
		var out DeviceManagementConfigurationGroupSettingInstanceTemplate
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationGroupSettingInstanceTemplate: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationSimpleSettingCollectionInstanceTemplate") {
		var out DeviceManagementConfigurationSimpleSettingCollectionInstanceTemplate
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationSimpleSettingCollectionInstanceTemplate: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationSimpleSettingInstanceTemplate") {
		var out DeviceManagementConfigurationSimpleSettingInstanceTemplate
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationSimpleSettingInstanceTemplate: %+v", err)
		}
		return out, nil
	}

	var parent BaseDeviceManagementConfigurationSettingInstanceTemplateImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDeviceManagementConfigurationSettingInstanceTemplateImpl: %+v", err)
	}

	return RawDeviceManagementConfigurationSettingInstanceTemplateImpl{
		deviceManagementConfigurationSettingInstanceTemplate: parent,
		Type:   value,
		Values: temp,
	}, nil

}
