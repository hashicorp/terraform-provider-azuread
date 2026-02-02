package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceManagementConfigurationSettingInstanceTemplate = DeviceManagementConfigurationSimpleSettingCollectionInstanceTemplate{}

type DeviceManagementConfigurationSimpleSettingCollectionInstanceTemplate struct {
	// Linked policy may append values which are not present in the template.
	AllowUnmanagedValues *bool `json:"allowUnmanagedValues,omitempty"`

	// Simple Setting Collection Value Template
	SimpleSettingCollectionValueTemplate *[]DeviceManagementConfigurationSimpleSettingValueTemplate `json:"simpleSettingCollectionValueTemplate,omitempty"`

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

func (s DeviceManagementConfigurationSimpleSettingCollectionInstanceTemplate) DeviceManagementConfigurationSettingInstanceTemplate() BaseDeviceManagementConfigurationSettingInstanceTemplateImpl {
	return BaseDeviceManagementConfigurationSettingInstanceTemplateImpl{
		IsRequired:                s.IsRequired,
		ODataId:                   s.ODataId,
		ODataType:                 s.ODataType,
		SettingDefinitionId:       s.SettingDefinitionId,
		SettingInstanceTemplateId: s.SettingInstanceTemplateId,
	}
}

var _ json.Marshaler = DeviceManagementConfigurationSimpleSettingCollectionInstanceTemplate{}

func (s DeviceManagementConfigurationSimpleSettingCollectionInstanceTemplate) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementConfigurationSimpleSettingCollectionInstanceTemplate
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementConfigurationSimpleSettingCollectionInstanceTemplate: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementConfigurationSimpleSettingCollectionInstanceTemplate: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementConfigurationSimpleSettingCollectionInstanceTemplate"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementConfigurationSimpleSettingCollectionInstanceTemplate: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &DeviceManagementConfigurationSimpleSettingCollectionInstanceTemplate{}

func (s *DeviceManagementConfigurationSimpleSettingCollectionInstanceTemplate) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AllowUnmanagedValues      *bool   `json:"allowUnmanagedValues,omitempty"`
		IsRequired                *bool   `json:"isRequired,omitempty"`
		ODataId                   *string `json:"@odata.id,omitempty"`
		ODataType                 *string `json:"@odata.type,omitempty"`
		SettingDefinitionId       *string `json:"settingDefinitionId,omitempty"`
		SettingInstanceTemplateId *string `json:"settingInstanceTemplateId,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AllowUnmanagedValues = decoded.AllowUnmanagedValues
	s.IsRequired = decoded.IsRequired
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.SettingDefinitionId = decoded.SettingDefinitionId
	s.SettingInstanceTemplateId = decoded.SettingInstanceTemplateId

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DeviceManagementConfigurationSimpleSettingCollectionInstanceTemplate into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["simpleSettingCollectionValueTemplate"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling SimpleSettingCollectionValueTemplate into list []json.RawMessage: %+v", err)
		}

		output := make([]DeviceManagementConfigurationSimpleSettingValueTemplate, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDeviceManagementConfigurationSimpleSettingValueTemplateImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'SimpleSettingCollectionValueTemplate' for 'DeviceManagementConfigurationSimpleSettingCollectionInstanceTemplate': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.SimpleSettingCollectionValueTemplate = &output
	}

	return nil
}
