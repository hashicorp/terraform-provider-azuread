package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceManagementConfigurationSettingInstance = DeviceManagementConfigurationSettingGroupCollectionInstance{}

type DeviceManagementConfigurationSettingGroupCollectionInstance struct {

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

func (s DeviceManagementConfigurationSettingGroupCollectionInstance) DeviceManagementConfigurationSettingInstance() BaseDeviceManagementConfigurationSettingInstanceImpl {
	return BaseDeviceManagementConfigurationSettingInstanceImpl{
		ODataId:                          s.ODataId,
		ODataType:                        s.ODataType,
		SettingDefinitionId:              s.SettingDefinitionId,
		SettingInstanceTemplateReference: s.SettingInstanceTemplateReference,
	}
}

var _ json.Marshaler = DeviceManagementConfigurationSettingGroupCollectionInstance{}

func (s DeviceManagementConfigurationSettingGroupCollectionInstance) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementConfigurationSettingGroupCollectionInstance
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementConfigurationSettingGroupCollectionInstance: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementConfigurationSettingGroupCollectionInstance: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementConfigurationSettingGroupCollectionInstance"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementConfigurationSettingGroupCollectionInstance: %+v", err)
	}

	return encoded, nil
}
