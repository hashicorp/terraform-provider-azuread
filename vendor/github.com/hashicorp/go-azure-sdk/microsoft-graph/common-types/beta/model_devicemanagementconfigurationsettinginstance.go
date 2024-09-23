package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSettingInstance interface {
	DeviceManagementConfigurationSettingInstance() BaseDeviceManagementConfigurationSettingInstanceImpl
}

var _ DeviceManagementConfigurationSettingInstance = BaseDeviceManagementConfigurationSettingInstanceImpl{}

type BaseDeviceManagementConfigurationSettingInstanceImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Setting Definition Id
	SettingDefinitionId *string `json:"settingDefinitionId,omitempty"`

	// Setting Instance Template Reference
	SettingInstanceTemplateReference *DeviceManagementConfigurationSettingInstanceTemplateReference `json:"settingInstanceTemplateReference,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseDeviceManagementConfigurationSettingInstanceImpl) DeviceManagementConfigurationSettingInstance() BaseDeviceManagementConfigurationSettingInstanceImpl {
	return s
}

var _ DeviceManagementConfigurationSettingInstance = RawDeviceManagementConfigurationSettingInstanceImpl{}

// RawDeviceManagementConfigurationSettingInstanceImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDeviceManagementConfigurationSettingInstanceImpl struct {
	deviceManagementConfigurationSettingInstance BaseDeviceManagementConfigurationSettingInstanceImpl
	Type                                         string
	Values                                       map[string]interface{}
}

func (s RawDeviceManagementConfigurationSettingInstanceImpl) DeviceManagementConfigurationSettingInstance() BaseDeviceManagementConfigurationSettingInstanceImpl {
	return s.deviceManagementConfigurationSettingInstance
}

func UnmarshalDeviceManagementConfigurationSettingInstanceImplementation(input []byte) (DeviceManagementConfigurationSettingInstance, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementConfigurationSettingInstance into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationChoiceSettingCollectionInstance") {
		var out DeviceManagementConfigurationChoiceSettingCollectionInstance
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationChoiceSettingCollectionInstance: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationChoiceSettingInstance") {
		var out DeviceManagementConfigurationChoiceSettingInstance
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationChoiceSettingInstance: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationGroupSettingCollectionInstance") {
		var out DeviceManagementConfigurationGroupSettingCollectionInstance
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationGroupSettingCollectionInstance: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationGroupSettingInstance") {
		var out DeviceManagementConfigurationGroupSettingInstance
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationGroupSettingInstance: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationSettingGroupCollectionInstance") {
		var out DeviceManagementConfigurationSettingGroupCollectionInstance
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationSettingGroupCollectionInstance: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationSettingGroupInstance") {
		var out DeviceManagementConfigurationSettingGroupInstance
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationSettingGroupInstance: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationSimpleSettingCollectionInstance") {
		var out DeviceManagementConfigurationSimpleSettingCollectionInstance
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationSimpleSettingCollectionInstance: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationSimpleSettingInstance") {
		var out DeviceManagementConfigurationSimpleSettingInstance
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationSimpleSettingInstance: %+v", err)
		}
		return out, nil
	}

	var parent BaseDeviceManagementConfigurationSettingInstanceImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDeviceManagementConfigurationSettingInstanceImpl: %+v", err)
	}

	return RawDeviceManagementConfigurationSettingInstanceImpl{
		deviceManagementConfigurationSettingInstance: parent,
		Type:   value,
		Values: temp,
	}, nil

}
