package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceManagementConfigurationIntegerSettingValueDefaultTemplate = DeviceManagementConfigurationIntegerSettingValueConstantDefaultTemplate{}

type DeviceManagementConfigurationIntegerSettingValueConstantDefaultTemplate struct {
	// Default Constant Value. Valid values -2147483648 to 2147483647
	ConstantValue *int64 `json:"constantValue,omitempty"`

	// Fields inherited from DeviceManagementConfigurationIntegerSettingValueDefaultTemplate

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s DeviceManagementConfigurationIntegerSettingValueConstantDefaultTemplate) DeviceManagementConfigurationIntegerSettingValueDefaultTemplate() BaseDeviceManagementConfigurationIntegerSettingValueDefaultTemplateImpl {
	return BaseDeviceManagementConfigurationIntegerSettingValueDefaultTemplateImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceManagementConfigurationIntegerSettingValueConstantDefaultTemplate{}

func (s DeviceManagementConfigurationIntegerSettingValueConstantDefaultTemplate) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementConfigurationIntegerSettingValueConstantDefaultTemplate
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementConfigurationIntegerSettingValueConstantDefaultTemplate: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementConfigurationIntegerSettingValueConstantDefaultTemplate: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementConfigurationIntegerSettingValueConstantDefaultTemplate"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementConfigurationIntegerSettingValueConstantDefaultTemplate: %+v", err)
	}

	return encoded, nil
}
