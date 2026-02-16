package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSettingValueDefinition interface {
	DeviceManagementConfigurationSettingValueDefinition() BaseDeviceManagementConfigurationSettingValueDefinitionImpl
}

var _ DeviceManagementConfigurationSettingValueDefinition = BaseDeviceManagementConfigurationSettingValueDefinitionImpl{}

type BaseDeviceManagementConfigurationSettingValueDefinitionImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseDeviceManagementConfigurationSettingValueDefinitionImpl) DeviceManagementConfigurationSettingValueDefinition() BaseDeviceManagementConfigurationSettingValueDefinitionImpl {
	return s
}

var _ DeviceManagementConfigurationSettingValueDefinition = RawDeviceManagementConfigurationSettingValueDefinitionImpl{}

// RawDeviceManagementConfigurationSettingValueDefinitionImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDeviceManagementConfigurationSettingValueDefinitionImpl struct {
	deviceManagementConfigurationSettingValueDefinition BaseDeviceManagementConfigurationSettingValueDefinitionImpl
	Type                                                string
	Values                                              map[string]interface{}
}

func (s RawDeviceManagementConfigurationSettingValueDefinitionImpl) DeviceManagementConfigurationSettingValueDefinition() BaseDeviceManagementConfigurationSettingValueDefinitionImpl {
	return s.deviceManagementConfigurationSettingValueDefinition
}

func UnmarshalDeviceManagementConfigurationSettingValueDefinitionImplementation(input []byte) (DeviceManagementConfigurationSettingValueDefinition, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementConfigurationSettingValueDefinition into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationIntegerSettingValueDefinition") {
		var out DeviceManagementConfigurationIntegerSettingValueDefinition
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationIntegerSettingValueDefinition: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationStringSettingValueDefinition") {
		var out DeviceManagementConfigurationStringSettingValueDefinition
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationStringSettingValueDefinition: %+v", err)
		}
		return out, nil
	}

	var parent BaseDeviceManagementConfigurationSettingValueDefinitionImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDeviceManagementConfigurationSettingValueDefinitionImpl: %+v", err)
	}

	return RawDeviceManagementConfigurationSettingValueDefinitionImpl{
		deviceManagementConfigurationSettingValueDefinition: parent,
		Type:   value,
		Values: temp,
	}, nil

}
