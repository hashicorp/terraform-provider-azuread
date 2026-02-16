package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceManagementConfigurationSetting{}

type DeviceManagementConfigurationSetting struct {
	// List of related Setting Definitions. This property is read-only.
	SettingDefinitions *[]DeviceManagementConfigurationSettingDefinition `json:"settingDefinitions,omitempty"`

	// Setting instance within policy
	SettingInstance DeviceManagementConfigurationSettingInstance `json:"settingInstance"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s DeviceManagementConfigurationSetting) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceManagementConfigurationSetting{}

func (s DeviceManagementConfigurationSetting) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementConfigurationSetting
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementConfigurationSetting: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementConfigurationSetting: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementConfigurationSetting"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementConfigurationSetting: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &DeviceManagementConfigurationSetting{}

func (s *DeviceManagementConfigurationSetting) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Id        *string `json:"id,omitempty"`
		ODataId   *string `json:"@odata.id,omitempty"`
		ODataType *string `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DeviceManagementConfigurationSetting into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["settingDefinitions"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling SettingDefinitions into list []json.RawMessage: %+v", err)
		}

		output := make([]DeviceManagementConfigurationSettingDefinition, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDeviceManagementConfigurationSettingDefinitionImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'SettingDefinitions' for 'DeviceManagementConfigurationSetting': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.SettingDefinitions = &output
	}

	if v, ok := temp["settingInstance"]; ok {
		impl, err := UnmarshalDeviceManagementConfigurationSettingInstanceImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'SettingInstance' for 'DeviceManagementConfigurationSetting': %+v", err)
		}
		s.SettingInstance = impl
	}

	return nil
}
