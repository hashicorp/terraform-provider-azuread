package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceManagementConfigurationSettingInstanceTemplate = DeviceManagementConfigurationGroupSettingInstanceTemplate{}

type DeviceManagementConfigurationGroupSettingInstanceTemplate struct {
	// Group Setting Value Template
	GroupSettingValueTemplate *DeviceManagementConfigurationGroupSettingValueTemplate `json:"groupSettingValueTemplate,omitempty"`

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

func (s DeviceManagementConfigurationGroupSettingInstanceTemplate) DeviceManagementConfigurationSettingInstanceTemplate() BaseDeviceManagementConfigurationSettingInstanceTemplateImpl {
	return BaseDeviceManagementConfigurationSettingInstanceTemplateImpl{
		IsRequired:                s.IsRequired,
		ODataId:                   s.ODataId,
		ODataType:                 s.ODataType,
		SettingDefinitionId:       s.SettingDefinitionId,
		SettingInstanceTemplateId: s.SettingInstanceTemplateId,
	}
}

var _ json.Marshaler = DeviceManagementConfigurationGroupSettingInstanceTemplate{}

func (s DeviceManagementConfigurationGroupSettingInstanceTemplate) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementConfigurationGroupSettingInstanceTemplate
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementConfigurationGroupSettingInstanceTemplate: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementConfigurationGroupSettingInstanceTemplate: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementConfigurationGroupSettingInstanceTemplate"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementConfigurationGroupSettingInstanceTemplate: %+v", err)
	}

	return encoded, nil
}
