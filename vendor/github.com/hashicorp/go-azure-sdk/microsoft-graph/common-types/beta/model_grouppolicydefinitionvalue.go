package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = GroupPolicyDefinitionValue{}

type GroupPolicyDefinitionValue struct {
	// Group Policy Configuration Type
	ConfigurationType *GroupPolicyConfigurationType `json:"configurationType,omitempty"`

	// The date and time the object was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The associated group policy definition with the value.
	Definition *GroupPolicyDefinition `json:"definition,omitempty"`

	// Enables or disables the associated group policy definition.
	Enabled *bool `json:"enabled,omitempty"`

	// The date and time the entity was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// The associated group policy presentation values with the definition value.
	PresentationValues *[]GroupPolicyPresentationValue `json:"presentationValues,omitempty"`

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

func (s GroupPolicyDefinitionValue) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = GroupPolicyDefinitionValue{}

func (s GroupPolicyDefinitionValue) MarshalJSON() ([]byte, error) {
	type wrapper GroupPolicyDefinitionValue
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GroupPolicyDefinitionValue: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GroupPolicyDefinitionValue: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.groupPolicyDefinitionValue"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GroupPolicyDefinitionValue: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &GroupPolicyDefinitionValue{}

func (s *GroupPolicyDefinitionValue) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ConfigurationType    *GroupPolicyConfigurationType `json:"configurationType,omitempty"`
		CreatedDateTime      *string                       `json:"createdDateTime,omitempty"`
		Definition           *GroupPolicyDefinition        `json:"definition,omitempty"`
		Enabled              *bool                         `json:"enabled,omitempty"`
		LastModifiedDateTime *string                       `json:"lastModifiedDateTime,omitempty"`
		Id                   *string                       `json:"id,omitempty"`
		ODataId              *string                       `json:"@odata.id,omitempty"`
		ODataType            *string                       `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ConfigurationType = decoded.ConfigurationType
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Definition = decoded.Definition
	s.Enabled = decoded.Enabled
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling GroupPolicyDefinitionValue into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["presentationValues"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling PresentationValues into list []json.RawMessage: %+v", err)
		}

		output := make([]GroupPolicyPresentationValue, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalGroupPolicyPresentationValueImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'PresentationValues' for 'GroupPolicyDefinitionValue': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.PresentationValues = &output
	}

	return nil
}
