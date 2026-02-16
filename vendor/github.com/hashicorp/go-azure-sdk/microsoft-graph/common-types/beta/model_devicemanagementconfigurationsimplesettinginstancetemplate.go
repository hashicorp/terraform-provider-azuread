package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceManagementConfigurationSettingInstanceTemplate = DeviceManagementConfigurationSimpleSettingInstanceTemplate{}

type DeviceManagementConfigurationSimpleSettingInstanceTemplate struct {
	// Simple Setting Value Template
	SimpleSettingValueTemplate DeviceManagementConfigurationSimpleSettingValueTemplate `json:"simpleSettingValueTemplate"`

	// Fields inherited from DeviceManagementConfigurationSettingInstanceTemplate

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

func (s DeviceManagementConfigurationSimpleSettingInstanceTemplate) DeviceManagementConfigurationSettingInstanceTemplate() BaseDeviceManagementConfigurationSettingInstanceTemplateImpl {
	return BaseDeviceManagementConfigurationSettingInstanceTemplateImpl{
		IsRequired:                s.IsRequired,
		ODataId:                   s.ODataId,
		ODataType:                 s.ODataType,
		SettingDefinitionId:       s.SettingDefinitionId,
		SettingInstanceTemplateId: s.SettingInstanceTemplateId,
	}
}

var _ json.Marshaler = DeviceManagementConfigurationSimpleSettingInstanceTemplate{}

func (s DeviceManagementConfigurationSimpleSettingInstanceTemplate) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementConfigurationSimpleSettingInstanceTemplate
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementConfigurationSimpleSettingInstanceTemplate: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementConfigurationSimpleSettingInstanceTemplate: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementConfigurationSimpleSettingInstanceTemplate"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementConfigurationSimpleSettingInstanceTemplate: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &DeviceManagementConfigurationSimpleSettingInstanceTemplate{}

func (s *DeviceManagementConfigurationSimpleSettingInstanceTemplate) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		IsRequired                *bool   `json:"isRequired,omitempty"`
		ODataId                   *string `json:"@odata.id,omitempty"`
		ODataType                 *string `json:"@odata.type,omitempty"`
		SettingDefinitionId       *string `json:"settingDefinitionId,omitempty"`
		SettingInstanceTemplateId *string `json:"settingInstanceTemplateId,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.IsRequired = decoded.IsRequired
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.SettingDefinitionId = decoded.SettingDefinitionId
	s.SettingInstanceTemplateId = decoded.SettingInstanceTemplateId

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DeviceManagementConfigurationSimpleSettingInstanceTemplate into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["simpleSettingValueTemplate"]; ok {
		impl, err := UnmarshalDeviceManagementConfigurationSimpleSettingValueTemplateImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'SimpleSettingValueTemplate' for 'DeviceManagementConfigurationSimpleSettingInstanceTemplate': %+v", err)
		}
		s.SimpleSettingValueTemplate = impl
	}

	return nil
}
