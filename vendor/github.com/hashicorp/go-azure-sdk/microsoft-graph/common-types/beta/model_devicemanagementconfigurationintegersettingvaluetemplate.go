package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceManagementConfigurationSimpleSettingValueTemplate = DeviceManagementConfigurationIntegerSettingValueTemplate{}

type DeviceManagementConfigurationIntegerSettingValueTemplate struct {
	// Integer Setting Value Default Template.
	DefaultValue DeviceManagementConfigurationIntegerSettingValueDefaultTemplate `json:"defaultValue"`

	// Recommended value definition.
	RecommendedValueDefinition *DeviceManagementConfigurationIntegerSettingValueDefinitionTemplate `json:"recommendedValueDefinition,omitempty"`

	// Required value definition.
	RequiredValueDefinition *DeviceManagementConfigurationIntegerSettingValueDefinitionTemplate `json:"requiredValueDefinition,omitempty"`

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

func (s DeviceManagementConfigurationIntegerSettingValueTemplate) DeviceManagementConfigurationSimpleSettingValueTemplate() BaseDeviceManagementConfigurationSimpleSettingValueTemplateImpl {
	return BaseDeviceManagementConfigurationSimpleSettingValueTemplateImpl{
		ODataId:                s.ODataId,
		ODataType:              s.ODataType,
		SettingValueTemplateId: s.SettingValueTemplateId,
	}
}

var _ json.Marshaler = DeviceManagementConfigurationIntegerSettingValueTemplate{}

func (s DeviceManagementConfigurationIntegerSettingValueTemplate) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementConfigurationIntegerSettingValueTemplate
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementConfigurationIntegerSettingValueTemplate: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementConfigurationIntegerSettingValueTemplate: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementConfigurationIntegerSettingValueTemplate"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementConfigurationIntegerSettingValueTemplate: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &DeviceManagementConfigurationIntegerSettingValueTemplate{}

func (s *DeviceManagementConfigurationIntegerSettingValueTemplate) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		RecommendedValueDefinition *DeviceManagementConfigurationIntegerSettingValueDefinitionTemplate `json:"recommendedValueDefinition,omitempty"`
		RequiredValueDefinition    *DeviceManagementConfigurationIntegerSettingValueDefinitionTemplate `json:"requiredValueDefinition,omitempty"`
		ODataId                    *string                                                             `json:"@odata.id,omitempty"`
		ODataType                  *string                                                             `json:"@odata.type,omitempty"`
		SettingValueTemplateId     *string                                                             `json:"settingValueTemplateId,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.RecommendedValueDefinition = decoded.RecommendedValueDefinition
	s.RequiredValueDefinition = decoded.RequiredValueDefinition
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.SettingValueTemplateId = decoded.SettingValueTemplateId

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DeviceManagementConfigurationIntegerSettingValueTemplate into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["defaultValue"]; ok {
		impl, err := UnmarshalDeviceManagementConfigurationIntegerSettingValueDefaultTemplateImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'DefaultValue' for 'DeviceManagementConfigurationIntegerSettingValueTemplate': %+v", err)
		}
		s.DefaultValue = impl
	}

	return nil
}
