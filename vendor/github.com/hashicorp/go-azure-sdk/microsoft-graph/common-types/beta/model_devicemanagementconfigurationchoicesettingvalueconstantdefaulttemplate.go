package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceManagementConfigurationChoiceSettingValueDefaultTemplate = DeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplate{}

type DeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplate struct {
	// Option Children
	Children *[]DeviceManagementConfigurationSettingInstanceTemplate `json:"children,omitempty"`

	// Default Constant Value
	SettingDefinitionOptionId nullable.Type[string] `json:"settingDefinitionOptionId,omitempty"`

	// Fields inherited from DeviceManagementConfigurationChoiceSettingValueDefaultTemplate

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s DeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplate) DeviceManagementConfigurationChoiceSettingValueDefaultTemplate() BaseDeviceManagementConfigurationChoiceSettingValueDefaultTemplateImpl {
	return BaseDeviceManagementConfigurationChoiceSettingValueDefaultTemplateImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplate{}

func (s DeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplate) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplate
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplate: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplate: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementConfigurationChoiceSettingValueConstantDefaultTemplate"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplate: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &DeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplate{}

func (s *DeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplate) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		SettingDefinitionOptionId nullable.Type[string] `json:"settingDefinitionOptionId,omitempty"`
		ODataId                   *string               `json:"@odata.id,omitempty"`
		ODataType                 *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.SettingDefinitionOptionId = decoded.SettingDefinitionOptionId
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplate into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["children"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Children into list []json.RawMessage: %+v", err)
		}

		output := make([]DeviceManagementConfigurationSettingInstanceTemplate, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDeviceManagementConfigurationSettingInstanceTemplateImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Children' for 'DeviceManagementConfigurationChoiceSettingValueConstantDefaultTemplate': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Children = &output
	}

	return nil
}
