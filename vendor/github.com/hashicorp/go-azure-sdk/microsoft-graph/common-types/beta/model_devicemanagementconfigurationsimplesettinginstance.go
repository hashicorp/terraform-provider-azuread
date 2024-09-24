package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceManagementConfigurationSettingInstance = DeviceManagementConfigurationSimpleSettingInstance{}

type DeviceManagementConfigurationSimpleSettingInstance struct {
	SimpleSettingValue *DeviceManagementConfigurationSimpleSettingValue `json:"simpleSettingValue,omitempty"`

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

func (s DeviceManagementConfigurationSimpleSettingInstance) DeviceManagementConfigurationSettingInstance() BaseDeviceManagementConfigurationSettingInstanceImpl {
	return BaseDeviceManagementConfigurationSettingInstanceImpl{
		ODataId:                          s.ODataId,
		ODataType:                        s.ODataType,
		SettingDefinitionId:              s.SettingDefinitionId,
		SettingInstanceTemplateReference: s.SettingInstanceTemplateReference,
	}
}

var _ json.Marshaler = DeviceManagementConfigurationSimpleSettingInstance{}

func (s DeviceManagementConfigurationSimpleSettingInstance) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementConfigurationSimpleSettingInstance
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementConfigurationSimpleSettingInstance: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementConfigurationSimpleSettingInstance: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementConfigurationSimpleSettingInstance"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementConfigurationSimpleSettingInstance: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &DeviceManagementConfigurationSimpleSettingInstance{}

func (s *DeviceManagementConfigurationSimpleSettingInstance) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ODataId                          *string                                                        `json:"@odata.id,omitempty"`
		ODataType                        *string                                                        `json:"@odata.type,omitempty"`
		SettingDefinitionId              *string                                                        `json:"settingDefinitionId,omitempty"`
		SettingInstanceTemplateReference *DeviceManagementConfigurationSettingInstanceTemplateReference `json:"settingInstanceTemplateReference,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.SettingDefinitionId = decoded.SettingDefinitionId
	s.SettingInstanceTemplateReference = decoded.SettingInstanceTemplateReference

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DeviceManagementConfigurationSimpleSettingInstance into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["simpleSettingValue"]; ok {
		impl, err := UnmarshalDeviceManagementConfigurationSimpleSettingValueImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'SimpleSettingValue' for 'DeviceManagementConfigurationSimpleSettingInstance': %+v", err)
		}
		s.SimpleSettingValue = &impl
	}

	return nil
}
