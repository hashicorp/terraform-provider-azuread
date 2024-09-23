package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceManagementSettingDefinition = DeviceManagementCollectionSettingDefinition{}

type DeviceManagementCollectionSettingDefinition struct {
	// The Setting Definition ID that describes what each element of the collection looks like
	ElementDefinitionId nullable.Type[string] `json:"elementDefinitionId,omitempty"`

	// Fields inherited from DeviceManagementSettingDefinition

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

func (s DeviceManagementCollectionSettingDefinition) DeviceManagementSettingDefinition() BaseDeviceManagementSettingDefinitionImpl {
	return BaseDeviceManagementSettingDefinitionImpl{
		Constraints:      s.Constraints,
		Dependencies:     s.Dependencies,
		Description:      s.Description,
		DisplayName:      s.DisplayName,
		DocumentationUrl: s.DocumentationUrl,
		HeaderSubtitle:   s.HeaderSubtitle,
		HeaderTitle:      s.HeaderTitle,
		IsTopLevel:       s.IsTopLevel,
		Keywords:         s.Keywords,
		PlaceholderText:  s.PlaceholderText,
		ValueType:        s.ValueType,
		Id:               s.Id,
		ODataId:          s.ODataId,
		ODataType:        s.ODataType,
	}
}

func (s DeviceManagementCollectionSettingDefinition) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceManagementCollectionSettingDefinition{}

func (s DeviceManagementCollectionSettingDefinition) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementCollectionSettingDefinition
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementCollectionSettingDefinition: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementCollectionSettingDefinition: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementCollectionSettingDefinition"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementCollectionSettingDefinition: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &DeviceManagementCollectionSettingDefinition{}

func (s *DeviceManagementCollectionSettingDefinition) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ElementDefinitionId nullable.Type[string]                `json:"elementDefinitionId,omitempty"`
		Dependencies        *[]DeviceManagementSettingDependency `json:"dependencies,omitempty"`
		Description         nullable.Type[string]                `json:"description,omitempty"`
		DisplayName         *string                              `json:"displayName,omitempty"`
		DocumentationUrl    nullable.Type[string]                `json:"documentationUrl,omitempty"`
		HeaderSubtitle      nullable.Type[string]                `json:"headerSubtitle,omitempty"`
		HeaderTitle         nullable.Type[string]                `json:"headerTitle,omitempty"`
		IsTopLevel          *bool                                `json:"isTopLevel,omitempty"`
		Keywords            *[]string                            `json:"keywords,omitempty"`
		PlaceholderText     nullable.Type[string]                `json:"placeholderText,omitempty"`
		ValueType           *DeviceManangementIntentValueType    `json:"valueType,omitempty"`
		Id                  *string                              `json:"id,omitempty"`
		ODataId             *string                              `json:"@odata.id,omitempty"`
		ODataType           *string                              `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ElementDefinitionId = decoded.ElementDefinitionId
	s.Dependencies = decoded.Dependencies
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.DocumentationUrl = decoded.DocumentationUrl
	s.HeaderSubtitle = decoded.HeaderSubtitle
	s.HeaderTitle = decoded.HeaderTitle
	s.Id = decoded.Id
	s.IsTopLevel = decoded.IsTopLevel
	s.Keywords = decoded.Keywords
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.PlaceholderText = decoded.PlaceholderText
	s.ValueType = decoded.ValueType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DeviceManagementCollectionSettingDefinition into map[string]json.RawMessage: %+v", err)
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
				return fmt.Errorf("unmarshaling index %d field 'Constraints' for 'DeviceManagementCollectionSettingDefinition': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Constraints = &output
	}

	return nil
}
