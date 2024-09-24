package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementSettingDefinition interface {
	Entity
	DeviceManagementSettingDefinition() BaseDeviceManagementSettingDefinitionImpl
}

var _ DeviceManagementSettingDefinition = BaseDeviceManagementSettingDefinitionImpl{}

type BaseDeviceManagementSettingDefinitionImpl struct {
	// Collection of constraints for the setting value
	Constraints *[]DeviceManagementConstraint `json:"constraints,omitempty"`

	// Collection of dependencies on other settings
	Dependencies *[]DeviceManagementSettingDependency `json:"dependencies,omitempty"`

	// The setting's description
	Description nullable.Type[string] `json:"description,omitempty"`

	// The setting's display name
	DisplayName *string `json:"displayName,omitempty"`

	// Url to setting documentation
	DocumentationUrl nullable.Type[string] `json:"documentationUrl,omitempty"`

	// subtitle of the setting header for more details about the category/section
	HeaderSubtitle nullable.Type[string] `json:"headerSubtitle,omitempty"`

	// title of the setting header represents a category/section of a setting/settings
	HeaderTitle nullable.Type[string] `json:"headerTitle,omitempty"`

	// If the setting is top level, it can be configured without the need to be wrapped in a collection or complex setting
	IsTopLevel *bool `json:"isTopLevel,omitempty"`

	// Keywords associated with the setting
	Keywords *[]string `json:"keywords,omitempty"`

	// Placeholder text as an example of valid input
	PlaceholderText nullable.Type[string] `json:"placeholderText,omitempty"`

	ValueType *DeviceManangementIntentValueType `json:"valueType,omitempty"`

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

func (s BaseDeviceManagementSettingDefinitionImpl) DeviceManagementSettingDefinition() BaseDeviceManagementSettingDefinitionImpl {
	return s
}

func (s BaseDeviceManagementSettingDefinitionImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ DeviceManagementSettingDefinition = RawDeviceManagementSettingDefinitionImpl{}

// RawDeviceManagementSettingDefinitionImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDeviceManagementSettingDefinitionImpl struct {
	deviceManagementSettingDefinition BaseDeviceManagementSettingDefinitionImpl
	Type                              string
	Values                            map[string]interface{}
}

func (s RawDeviceManagementSettingDefinitionImpl) DeviceManagementSettingDefinition() BaseDeviceManagementSettingDefinitionImpl {
	return s.deviceManagementSettingDefinition
}

func (s RawDeviceManagementSettingDefinitionImpl) Entity() BaseEntityImpl {
	return s.deviceManagementSettingDefinition.Entity()
}

var _ json.Marshaler = BaseDeviceManagementSettingDefinitionImpl{}

func (s BaseDeviceManagementSettingDefinitionImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseDeviceManagementSettingDefinitionImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseDeviceManagementSettingDefinitionImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseDeviceManagementSettingDefinitionImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementSettingDefinition"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseDeviceManagementSettingDefinitionImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseDeviceManagementSettingDefinitionImpl{}

func (s *BaseDeviceManagementSettingDefinitionImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Dependencies     *[]DeviceManagementSettingDependency `json:"dependencies,omitempty"`
		Description      nullable.Type[string]                `json:"description,omitempty"`
		DisplayName      *string                              `json:"displayName,omitempty"`
		DocumentationUrl nullable.Type[string]                `json:"documentationUrl,omitempty"`
		HeaderSubtitle   nullable.Type[string]                `json:"headerSubtitle,omitempty"`
		HeaderTitle      nullable.Type[string]                `json:"headerTitle,omitempty"`
		IsTopLevel       *bool                                `json:"isTopLevel,omitempty"`
		Keywords         *[]string                            `json:"keywords,omitempty"`
		PlaceholderText  nullable.Type[string]                `json:"placeholderText,omitempty"`
		ValueType        *DeviceManangementIntentValueType    `json:"valueType,omitempty"`
		Id               *string                              `json:"id,omitempty"`
		ODataId          *string                              `json:"@odata.id,omitempty"`
		ODataType        *string                              `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Dependencies = decoded.Dependencies
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.DocumentationUrl = decoded.DocumentationUrl
	s.HeaderSubtitle = decoded.HeaderSubtitle
	s.HeaderTitle = decoded.HeaderTitle
	s.IsTopLevel = decoded.IsTopLevel
	s.Keywords = decoded.Keywords
	s.PlaceholderText = decoded.PlaceholderText
	s.ValueType = decoded.ValueType
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseDeviceManagementSettingDefinitionImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["constraints"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Constraints into list []json.RawMessage: %+v", err)
		}

		output := make([]DeviceManagementConstraint, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDeviceManagementConstraintImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Constraints' for 'BaseDeviceManagementSettingDefinitionImpl': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Constraints = &output
	}

	return nil
}

func UnmarshalDeviceManagementSettingDefinitionImplementation(input []byte) (DeviceManagementSettingDefinition, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementSettingDefinition into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementAbstractComplexSettingDefinition") {
		var out DeviceManagementAbstractComplexSettingDefinition
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementAbstractComplexSettingDefinition: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementCollectionSettingDefinition") {
		var out DeviceManagementCollectionSettingDefinition
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementCollectionSettingDefinition: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.deviceManagementComplexSettingDefinition") {
		var out DeviceManagementComplexSettingDefinition
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DeviceManagementComplexSettingDefinition: %+v", err)
		}
		return out, nil
	}

	var parent BaseDeviceManagementSettingDefinitionImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDeviceManagementSettingDefinitionImpl: %+v", err)
	}

	return RawDeviceManagementSettingDefinitionImpl{
		deviceManagementSettingDefinition: parent,
		Type:                              value,
		Values:                            temp,
	}, nil

}
