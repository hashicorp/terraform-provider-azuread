package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceManagementConfigurationSimpleSettingValue = DeviceManagementConfigurationIntegerSettingValue{}

type DeviceManagementConfigurationIntegerSettingValue struct {
	// Value of the integer setting.
	Value *int64 `json:"value,omitempty"`

	// Fields inherited from DeviceManagementConfigurationSettingValue

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Setting value template reference
	SettingValueTemplateReference *DeviceManagementConfigurationSettingValueTemplateReference `json:"settingValueTemplateReference,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s DeviceManagementConfigurationIntegerSettingValue) DeviceManagementConfigurationSimpleSettingValue() BaseDeviceManagementConfigurationSimpleSettingValueImpl {
	return BaseDeviceManagementConfigurationSimpleSettingValueImpl{
		ODataId:                       s.ODataId,
		ODataType:                     s.ODataType,
		SettingValueTemplateReference: s.SettingValueTemplateReference,
	}
}

func (s DeviceManagementConfigurationIntegerSettingValue) DeviceManagementConfigurationSettingValue() BaseDeviceManagementConfigurationSettingValueImpl {
	return BaseDeviceManagementConfigurationSettingValueImpl{
		ODataId:                       s.ODataId,
		ODataType:                     s.ODataType,
		SettingValueTemplateReference: s.SettingValueTemplateReference,
	}
}

var _ json.Marshaler = DeviceManagementConfigurationIntegerSettingValue{}

func (s DeviceManagementConfigurationIntegerSettingValue) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementConfigurationIntegerSettingValue
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementConfigurationIntegerSettingValue: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementConfigurationIntegerSettingValue: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementConfigurationIntegerSettingValue"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementConfigurationIntegerSettingValue: %+v", err)
	}

	return encoded, nil
}
