package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceManagementConfigurationSettingInstanceTemplate = DeviceManagementConfigurationChoiceSettingCollectionInstanceTemplate{}

type DeviceManagementConfigurationChoiceSettingCollectionInstanceTemplate struct {
	// Linked policy may append values which are not present in the template.
	AllowUnmanagedValues *bool `json:"allowUnmanagedValues,omitempty"`

	// Choice Setting Collection Value Template
	ChoiceSettingCollectionValueTemplate *[]DeviceManagementConfigurationChoiceSettingValueTemplate `json:"choiceSettingCollectionValueTemplate,omitempty"`

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

func (s DeviceManagementConfigurationChoiceSettingCollectionInstanceTemplate) DeviceManagementConfigurationSettingInstanceTemplate() BaseDeviceManagementConfigurationSettingInstanceTemplateImpl {
	return BaseDeviceManagementConfigurationSettingInstanceTemplateImpl{
		IsRequired:                s.IsRequired,
		ODataId:                   s.ODataId,
		ODataType:                 s.ODataType,
		SettingDefinitionId:       s.SettingDefinitionId,
		SettingInstanceTemplateId: s.SettingInstanceTemplateId,
	}
}

var _ json.Marshaler = DeviceManagementConfigurationChoiceSettingCollectionInstanceTemplate{}

func (s DeviceManagementConfigurationChoiceSettingCollectionInstanceTemplate) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementConfigurationChoiceSettingCollectionInstanceTemplate
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementConfigurationChoiceSettingCollectionInstanceTemplate: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementConfigurationChoiceSettingCollectionInstanceTemplate: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementConfigurationChoiceSettingCollectionInstanceTemplate"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementConfigurationChoiceSettingCollectionInstanceTemplate: %+v", err)
	}

	return encoded, nil
}
