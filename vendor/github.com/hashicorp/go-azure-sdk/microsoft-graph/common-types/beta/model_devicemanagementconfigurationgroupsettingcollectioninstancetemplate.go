package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceManagementConfigurationSettingInstanceTemplate = DeviceManagementConfigurationGroupSettingCollectionInstanceTemplate{}

type DeviceManagementConfigurationGroupSettingCollectionInstanceTemplate struct {
	// Linked policy may append values which are not present in the template.
	AllowUnmanagedValues *bool `json:"allowUnmanagedValues,omitempty"`

	// Group Setting Collection Value Template
	GroupSettingCollectionValueTemplate *[]DeviceManagementConfigurationGroupSettingValueTemplate `json:"groupSettingCollectionValueTemplate,omitempty"`

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

func (s DeviceManagementConfigurationGroupSettingCollectionInstanceTemplate) DeviceManagementConfigurationSettingInstanceTemplate() BaseDeviceManagementConfigurationSettingInstanceTemplateImpl {
	return BaseDeviceManagementConfigurationSettingInstanceTemplateImpl{
		IsRequired:                s.IsRequired,
		ODataId:                   s.ODataId,
		ODataType:                 s.ODataType,
		SettingDefinitionId:       s.SettingDefinitionId,
		SettingInstanceTemplateId: s.SettingInstanceTemplateId,
	}
}

var _ json.Marshaler = DeviceManagementConfigurationGroupSettingCollectionInstanceTemplate{}

func (s DeviceManagementConfigurationGroupSettingCollectionInstanceTemplate) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementConfigurationGroupSettingCollectionInstanceTemplate
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementConfigurationGroupSettingCollectionInstanceTemplate: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementConfigurationGroupSettingCollectionInstanceTemplate: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementConfigurationGroupSettingCollectionInstanceTemplate"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementConfigurationGroupSettingCollectionInstanceTemplate: %+v", err)
	}

	return encoded, nil
}
