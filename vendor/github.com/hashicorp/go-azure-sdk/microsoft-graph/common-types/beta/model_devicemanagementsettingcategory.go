package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementSettingCategory interface {
	Entity
	DeviceManagementSettingCategory() BaseDeviceManagementSettingCategoryImpl
}

var _ DeviceManagementSettingCategory = BaseDeviceManagementSettingCategoryImpl{}

type BaseDeviceManagementSettingCategoryImpl struct {
	// The category name
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The category contains top level required setting
	HasRequiredSetting nullable.Type[bool] `json:"hasRequiredSetting,omitempty"`

	// The setting definitions this category contains
	SettingDefinitions *[]DeviceManagementSettingDefinition `json:"settingDefinitions,omitempty"`

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

func (s BaseDeviceManagementSettingCategoryImpl) DeviceManagementSettingCategory() BaseDeviceManagementSettingCategoryImpl {
	return s
}

func (s BaseDeviceManagementSettingCategoryImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ DeviceManagementSettingCategory = RawDeviceManagementSettingCategoryImpl{}

// RawDeviceManagementSettingCategoryImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDeviceManagementSettingCategoryImpl struct {
	deviceManagementSettingCategory BaseDeviceManagementSettingCategoryImpl
	Type                            string
	Values                          map[string]interface{}
}

func (s RawDeviceManagementSettingCategoryImpl) DeviceManagementSettingCategory() BaseDeviceManagementSettingCategoryImpl {
	return s.deviceManagementSettingCategory
}

func (s RawDeviceManagementSettingCategoryImpl) Entity() BaseEntityImpl {
	return s.deviceManagementSettingCategory.Entity()
}

var _ json.Marshaler = BaseDeviceManagementSettingCategoryImpl{}

func (s BaseDeviceManagementSettingCategoryImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseDeviceManagementSettingCategoryImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseDeviceManagementSettingCategoryImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseDeviceManagementSettingCategoryImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementSettingCategory"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseDeviceManagementSettingCategoryImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseDeviceManagementSettingCategoryImpl{}

func (s *BaseDeviceManagementSettingCategoryImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DisplayName        nullable.Type[string] `json:"displayName,omitempty"`
		HasRequiredSetting nullable.Type[bool]   `json:"hasRequiredSetting,omitempty"`
		Id                 *string               `json:"id,omitempty"`
		ODataId            *string               `json:"@odata.id,omitempty"`
		ODataType          *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DisplayName = decoded.DisplayName
	s.HasRequiredSetting = decoded.HasRequiredSetting
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseDeviceManagementSettingCategoryImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["settingDefinitions"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling SettingDefinitions into list []json.RawMessage: %+v", err)
		}

		output := make([]DeviceManagementSettingDefinition, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDeviceManagementSettingDefinitionImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'SettingDefinitions' for 'BaseDeviceManagementSettingCategoryImpl': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.SettingDefinitions = &output
	}

	return nil
}

func UnmarshalDeviceManagementSettingCategoryImplementation(input []byte) (DeviceManagementSettingCategory, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementSettingCategory into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementIntentSettingCategory") {
		var out DeviceManagementIntentSettingCategory
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementIntentSettingCategory: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementTemplateSettingCategory") {
		var out DeviceManagementTemplateSettingCategory
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementTemplateSettingCategory: %+v", err)
		}
		return out, nil
	}

	var parent BaseDeviceManagementSettingCategoryImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDeviceManagementSettingCategoryImpl: %+v", err)
	}

	return RawDeviceManagementSettingCategoryImpl{
		deviceManagementSettingCategory: parent,
		Type:                            value,
		Values:                          temp,
	}, nil

}
