package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceManagementConfigurationSimpleSettingValueTemplate = DeviceManagementConfigurationStringSettingValueTemplate{}

type DeviceManagementConfigurationStringSettingValueTemplate struct {
	// String Setting Value Default Template.
	DefaultValue DeviceManagementConfigurationStringSettingValueDefaultTemplate `json:"defaultValue"`

	// Fields inherited from DeviceManagementConfigurationSimpleSettingValueTemplate

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Setting Value Template Id
	SettingValueTemplateId *string `json:"settingValueTemplateId,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s DeviceManagementConfigurationStringSettingValueTemplate) DeviceManagementConfigurationSimpleSettingValueTemplate() BaseDeviceManagementConfigurationSimpleSettingValueTemplateImpl {
	return BaseDeviceManagementConfigurationSimpleSettingValueTemplateImpl{
		ODataId:                s.ODataId,
		ODataType:              s.ODataType,
		SettingValueTemplateId: s.SettingValueTemplateId,
	}
}

var _ json.Marshaler = DeviceManagementConfigurationStringSettingValueTemplate{}

func (s DeviceManagementConfigurationStringSettingValueTemplate) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementConfigurationStringSettingValueTemplate
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementConfigurationStringSettingValueTemplate: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementConfigurationStringSettingValueTemplate: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementConfigurationStringSettingValueTemplate"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementConfigurationStringSettingValueTemplate: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &DeviceManagementConfigurationStringSettingValueTemplate{}

func (s *DeviceManagementConfigurationStringSettingValueTemplate) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ODataId                *string `json:"@odata.id,omitempty"`
		ODataType              *string `json:"@odata.type,omitempty"`
		SettingValueTemplateId *string `json:"settingValueTemplateId,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.SettingValueTemplateId = decoded.SettingValueTemplateId

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DeviceManagementConfigurationStringSettingValueTemplate into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["defaultValue"]; ok {
		impl, err := UnmarshalDeviceManagementConfigurationStringSettingValueDefaultTemplateImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'DefaultValue' for 'DeviceManagementConfigurationStringSettingValueTemplate': %+v", err)
		}
		s.DefaultValue = impl
	}

	return nil
}
