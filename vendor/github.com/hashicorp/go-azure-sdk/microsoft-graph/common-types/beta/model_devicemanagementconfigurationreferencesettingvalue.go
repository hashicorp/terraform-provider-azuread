package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceManagementConfigurationStringSettingValue = DeviceManagementConfigurationReferenceSettingValue{}

type DeviceManagementConfigurationReferenceSettingValue struct {
	// A note that admin can use to put some contextual information
	Note nullable.Type[string] `json:"note,omitempty"`

	// Fields inherited from DeviceManagementConfigurationStringSettingValue

	// Value of the string setting.
	Value nullable.Type[string] `json:"value,omitempty"`

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

func (s DeviceManagementConfigurationReferenceSettingValue) DeviceManagementConfigurationStringSettingValue() BaseDeviceManagementConfigurationStringSettingValueImpl {
	return BaseDeviceManagementConfigurationStringSettingValueImpl{
		Value:                         s.Value,
		ODataId:                       s.ODataId,
		ODataType:                     s.ODataType,
		SettingValueTemplateReference: s.SettingValueTemplateReference,
	}
}

func (s DeviceManagementConfigurationReferenceSettingValue) DeviceManagementConfigurationSimpleSettingValue() BaseDeviceManagementConfigurationSimpleSettingValueImpl {
	return BaseDeviceManagementConfigurationSimpleSettingValueImpl{
		ODataId:                       s.ODataId,
		ODataType:                     s.ODataType,
		SettingValueTemplateReference: s.SettingValueTemplateReference,
	}
}

func (s DeviceManagementConfigurationReferenceSettingValue) DeviceManagementConfigurationSettingValue() BaseDeviceManagementConfigurationSettingValueImpl {
	return BaseDeviceManagementConfigurationSettingValueImpl{
		ODataId:                       s.ODataId,
		ODataType:                     s.ODataType,
		SettingValueTemplateReference: s.SettingValueTemplateReference,
	}
}

var _ json.Marshaler = DeviceManagementConfigurationReferenceSettingValue{}

func (s DeviceManagementConfigurationReferenceSettingValue) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementConfigurationReferenceSettingValue
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementConfigurationReferenceSettingValue: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementConfigurationReferenceSettingValue: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementConfigurationReferenceSettingValue"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementConfigurationReferenceSettingValue: %+v", err)
	}

	return encoded, nil
}
