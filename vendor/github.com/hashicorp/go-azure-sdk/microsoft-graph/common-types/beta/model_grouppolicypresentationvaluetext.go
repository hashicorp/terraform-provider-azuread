package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ GroupPolicyPresentationValue = GroupPolicyPresentationValueText{}

type GroupPolicyPresentationValueText struct {
	// A string value for the associated presentation.
	Value nullable.Type[string] `json:"value,omitempty"`

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

func (s GroupPolicyPresentationValueText) GroupPolicyPresentationValue() BaseGroupPolicyPresentationValueImpl {
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

func (s GroupPolicyPresentationValueText) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = GroupPolicyPresentationValueText{}

func (s GroupPolicyPresentationValueText) MarshalJSON() ([]byte, error) {
	type wrapper GroupPolicyPresentationValueText
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GroupPolicyPresentationValueText: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GroupPolicyPresentationValueText: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.groupPolicyPresentationValueText"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GroupPolicyPresentationValueText: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &GroupPolicyPresentationValueText{}

func (s *GroupPolicyPresentationValueText) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Value                nullable.Type[string]       `json:"value,omitempty"`
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

	s.Value = decoded.Value
	s.CreatedDateTime = decoded.CreatedDateTime
	s.DefinitionValue = decoded.DefinitionValue
	s.Id = decoded.Id
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling GroupPolicyPresentationValueText into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["presentation"]; ok {
		impl, err := UnmarshalGroupPolicyPresentationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Presentation' for 'GroupPolicyPresentationValueText': %+v", err)
		}
		s.Presentation = &impl
	}

	return nil
}
