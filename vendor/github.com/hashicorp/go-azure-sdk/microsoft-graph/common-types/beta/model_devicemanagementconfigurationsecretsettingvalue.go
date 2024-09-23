package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceManagementConfigurationSimpleSettingValue = DeviceManagementConfigurationSecretSettingValue{}

type DeviceManagementConfigurationSecretSettingValue struct {
	// Value of the secret setting.
	Value nullable.Type[string] `json:"value,omitempty"`

	// type tracking the encryption state of a secret setting value
	ValueState *DeviceManagementConfigurationSecretSettingValueState `json:"valueState,omitempty"`

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

func (s DeviceManagementConfigurationSecretSettingValue) DeviceManagementConfigurationSimpleSettingValue() BaseDeviceManagementConfigurationSimpleSettingValueImpl {
	return BaseDeviceManagementConfigurationSimpleSettingValueImpl{
		ODataId:                       s.ODataId,
		ODataType:                     s.ODataType,
		SettingValueTemplateReference: s.SettingValueTemplateReference,
	}
}

func (s DeviceManagementConfigurationSecretSettingValue) DeviceManagementConfigurationSettingValue() BaseDeviceManagementConfigurationSettingValueImpl {
	return BaseDeviceManagementConfigurationSettingValueImpl{
		ODataId:                       s.ODataId,
		ODataType:                     s.ODataType,
		SettingValueTemplateReference: s.SettingValueTemplateReference,
	}
}

var _ json.Marshaler = DeviceManagementConfigurationSecretSettingValue{}

func (s DeviceManagementConfigurationSecretSettingValue) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementConfigurationSecretSettingValue
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementConfigurationSecretSettingValue: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementConfigurationSecretSettingValue: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementConfigurationSecretSettingValue"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementConfigurationSecretSettingValue: %+v", err)
	}

	return encoded, nil
}
