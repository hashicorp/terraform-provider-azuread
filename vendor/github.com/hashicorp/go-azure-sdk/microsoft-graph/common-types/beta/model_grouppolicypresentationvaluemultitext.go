package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ GroupPolicyPresentationValue = GroupPolicyPresentationValueMultiText{}

type GroupPolicyPresentationValueMultiText struct {
	// A collection of non-empty strings for the associated presentation.
	Values *[]string `json:"values,omitempty"`

	// Fields inherited from GroupPolicyPresentationValue

	// The date and time the object was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The group policy definition value associated with the presentation value.
	DefinitionValue *GroupPolicyDefinitionValue `json:"definitionValue,omitempty"`

	// The date and time the object was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// The group policy presentation associated with the presentation value.
	Presentation *GroupPolicyPresentation `json:"presentation,omitempty"`

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

func (s GroupPolicyPresentationValueMultiText) GroupPolicyPresentationValue() BaseGroupPolicyPresentationValueImpl {
	return BaseGroupPolicyPresentationValueImpl{
		CreatedDateTime:      s.CreatedDateTime,
		DefinitionValue:      s.DefinitionValue,
		LastModifiedDateTime: s.LastModifiedDateTime,
		Presentation:         s.Presentation,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s GroupPolicyPresentationValueMultiText) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = GroupPolicyPresentationValueMultiText{}

func (s GroupPolicyPresentationValueMultiText) MarshalJSON() ([]byte, error) {
	type wrapper GroupPolicyPresentationValueMultiText
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GroupPolicyPresentationValueMultiText: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GroupPolicyPresentationValueMultiText: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.groupPolicyPresentationValueMultiText"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GroupPolicyPresentationValueMultiText: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &GroupPolicyPresentationValueMultiText{}

func (s *GroupPolicyPresentationValueMultiText) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Values               *[]string                   `json:"values,omitempty"`
		CreatedDateTime      *string                     `json:"createdDateTime,omitempty"`
		DefinitionValue      *GroupPolicyDefinitionValue `json:"definitionValue,omitempty"`
		LastModifiedDateTime *string                     `json:"lastModifiedDateTime,omitempty"`
		Id                   *string                     `json:"id,omitempty"`
		ODataId              *string                     `json:"@odata.id,omitempty"`
		ODataType            *string                     `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Values = decoded.Values
	s.CreatedDateTime = decoded.CreatedDateTime
	s.DefinitionValue = decoded.DefinitionValue
	s.Id = decoded.Id
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling GroupPolicyPresentationValueMultiText into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["presentation"]; ok {
		impl, err := UnmarshalGroupPolicyPresentationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Presentation' for 'GroupPolicyPresentationValueMultiText': %+v", err)
		}
		s.Presentation = &impl
	}

	return nil
}
