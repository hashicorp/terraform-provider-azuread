package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceManagementConfigurationStringSettingValueDefaultTemplate = DeviceManagementConfigurationStringSettingValueConstantDefaultTemplate{}

type DeviceManagementConfigurationStringSettingValueConstantDefaultTemplate struct {
	// Default Constant Value
	ConstantValue *string `json:"constantValue,omitempty"`

	// Fields inherited from DeviceManagementConfigurationStringSettingValueDefaultTemplate

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s DeviceManagementConfigurationStringSettingValueConstantDefaultTemplate) DeviceManagementConfigurationStringSettingValueDefaultTemplate() BaseDeviceManagementConfigurationStringSettingValueDefaultTemplateImpl {
	return BaseDeviceManagementConfigurationStringSettingValueDefaultTemplateImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceManagementConfigurationStringSettingValueConstantDefaultTemplate{}

func (s DeviceManagementConfigurationStringSettingValueConstantDefaultTemplate) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementConfigurationStringSettingValueConstantDefaultTemplate
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementConfigurationStringSettingValueConstantDefaultTemplate: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementConfigurationStringSettingValueConstantDefaultTemplate: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementConfigurationStringSettingValueConstantDefaultTemplate"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementConfigurationStringSettingValueConstantDefaultTemplate: %+v", err)
	}

	return encoded, nil
}
