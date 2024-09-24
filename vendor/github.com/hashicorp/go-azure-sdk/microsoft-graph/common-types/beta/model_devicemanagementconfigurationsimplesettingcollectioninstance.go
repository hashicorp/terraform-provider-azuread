package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceManagementConfigurationSettingInstance = DeviceManagementConfigurationSimpleSettingCollectionInstance{}

type DeviceManagementConfigurationSimpleSettingCollectionInstance struct {
	// Simple setting collection instance value
	SimpleSettingCollectionValue *[]DeviceManagementConfigurationSimpleSettingValue `json:"simpleSettingCollectionValue,omitempty"`

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

func (s DeviceManagementConfigurationSimpleSettingCollectionInstance) DeviceManagementConfigurationSettingInstance() BaseDeviceManagementConfigurationSettingInstanceImpl {
	return BaseDeviceManagementConfigurationSettingInstanceImpl{
		ODataId:                          s.ODataId,
		ODataType:                        s.ODataType,
		SettingDefinitionId:              s.SettingDefinitionId,
		SettingInstanceTemplateReference: s.SettingInstanceTemplateReference,
	}
}

var _ json.Marshaler = DeviceManagementConfigurationSimpleSettingCollectionInstance{}

func (s DeviceManagementConfigurationSimpleSettingCollectionInstance) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementConfigurationSimpleSettingCollectionInstance
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementConfigurationSimpleSettingCollectionInstance: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementConfigurationSimpleSettingCollectionInstance: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementConfigurationSimpleSettingCollectionInstance"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementConfigurationSimpleSettingCollectionInstance: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &DeviceManagementConfigurationSimpleSettingCollectionInstance{}

func (s *DeviceManagementConfigurationSimpleSettingCollectionInstance) UnmarshalJSON(bytes []byte) error {
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
		return fmt.Errorf("unmarshaling DeviceManagementConfigurationSimpleSettingCollectionInstance into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["simpleSettingCollectionValue"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling SimpleSettingCollectionValue into list []json.RawMessage: %+v", err)
		}

		output := make([]DeviceManagementConfigurationSimpleSettingValue, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDeviceManagementConfigurationSimpleSettingValueImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'SimpleSettingCollectionValue' for 'DeviceManagementConfigurationSimpleSettingCollectionInstance': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.SimpleSettingCollectionValue = &output
	}

	return nil
}
