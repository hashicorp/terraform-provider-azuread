package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceManagementConfigurationSettingInstance = DeviceManagementConfigurationGroupSettingCollectionInstance{}

type DeviceManagementConfigurationGroupSettingCollectionInstance struct {
	// A collection of GroupSetting values
	GroupSettingCollectionValue *[]DeviceManagementConfigurationGroupSettingValue `json:"groupSettingCollectionValue,omitempty"`

	// Fields inherited from DeviceManagementConfigurationSettingInstance

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Setting Definition Id
	SettingDefinitionId *string `json:"settingDefinitionId,omitempty"`

	// Setting Instance Template Reference
	SettingInstanceTemplateReference *DeviceManagementConfigurationSettingInstanceTemplateReference `json:"settingInstanceTemplateReference,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s DeviceManagementConfigurationGroupSettingCollectionInstance) DeviceManagementConfigurationSettingInstance() BaseDeviceManagementConfigurationSettingInstanceImpl {
	return BaseDeviceManagementConfigurationSettingInstanceImpl{
		ODataId:                          s.ODataId,
		ODataType:                        s.ODataType,
		SettingDefinitionId:              s.SettingDefinitionId,
		SettingInstanceTemplateReference: s.SettingInstanceTemplateReference,
	}
}

var _ json.Marshaler = DeviceManagementConfigurationGroupSettingCollectionInstance{}

func (s DeviceManagementConfigurationGroupSettingCollectionInstance) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementConfigurationGroupSettingCollectionInstance
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementConfigurationGroupSettingCollectionInstance: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementConfigurationGroupSettingCollectionInstance: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementConfigurationGroupSettingCollectionInstance"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementConfigurationGroupSettingCollectionInstance: %+v", err)
	}

	return encoded, nil
}
