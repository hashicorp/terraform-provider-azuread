package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationIntegerSettingValueDefaultTemplate interface {
	DeviceManagementConfigurationIntegerSettingValueDefaultTemplate() BaseDeviceManagementConfigurationIntegerSettingValueDefaultTemplateImpl
}

var _ DeviceManagementConfigurationIntegerSettingValueDefaultTemplate = BaseDeviceManagementConfigurationIntegerSettingValueDefaultTemplateImpl{}

type BaseDeviceManagementConfigurationIntegerSettingValueDefaultTemplateImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseDeviceManagementConfigurationIntegerSettingValueDefaultTemplateImpl) DeviceManagementConfigurationIntegerSettingValueDefaultTemplate() BaseDeviceManagementConfigurationIntegerSettingValueDefaultTemplateImpl {
	return s
}

var _ DeviceManagementConfigurationIntegerSettingValueDefaultTemplate = RawDeviceManagementConfigurationIntegerSettingValueDefaultTemplateImpl{}

// RawDeviceManagementConfigurationIntegerSettingValueDefaultTemplateImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawDeviceManagementConfigurationIntegerSettingValueDefaultTemplateImpl struct {
	deviceManagementConfigurationIntegerSettingValueDefaultTemplate BaseDeviceManagementConfigurationIntegerSettingValueDefaultTemplateImpl
	Type                                                            string
	Values                                                          map[string]interface{}
}

func (s RawDeviceManagementConfigurationIntegerSettingValueDefaultTemplateImpl) DeviceManagementConfigurationIntegerSettingValueDefaultTemplate() BaseDeviceManagementConfigurationIntegerSettingValueDefaultTemplateImpl {
	return s.deviceManagementConfigurationIntegerSettingValueDefaultTemplate
}

func (s RawDeviceManagementConfigurationIntegerSettingValueDefaultTemplateImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalDeviceManagementConfigurationIntegerSettingValueDefaultTemplateImplementation(input []byte) (DeviceManagementConfigurationIntegerSettingValueDefaultTemplate, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementConfigurationIntegerSettingValueDefaultTemplate into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementConfigurationIntegerSettingValueConstantDefaultTemplate") {
		var out DeviceManagementConfigurationIntegerSettingValueConstantDefaultTemplate
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementConfigurationIntegerSettingValueConstantDefaultTemplate: %+v", err)
		}
		return out, nil
	}

	var parent BaseDeviceManagementConfigurationIntegerSettingValueDefaultTemplateImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDeviceManagementConfigurationIntegerSettingValueDefaultTemplateImpl: %+v", err)
	}

	return RawDeviceManagementConfigurationIntegerSettingValueDefaultTemplateImpl{
		deviceManagementConfigurationIntegerSettingValueDefaultTemplate: parent,
		Type:   value,
		Values: temp,
	}, nil

}
