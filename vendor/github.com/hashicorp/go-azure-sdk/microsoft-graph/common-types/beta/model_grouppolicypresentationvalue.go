package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyPresentationValue interface {
	Entity
	GroupPolicyPresentationValue() BaseGroupPolicyPresentationValueImpl
}

var _ GroupPolicyPresentationValue = BaseGroupPolicyPresentationValueImpl{}

type BaseGroupPolicyPresentationValueImpl struct {
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

func (s BaseGroupPolicyPresentationValueImpl) GroupPolicyPresentationValue() BaseGroupPolicyPresentationValueImpl {
	return s
}

func (s BaseGroupPolicyPresentationValueImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ GroupPolicyPresentationValue = RawGroupPolicyPresentationValueImpl{}

// RawGroupPolicyPresentationValueImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawGroupPolicyPresentationValueImpl struct {
	groupPolicyPresentationValue BaseGroupPolicyPresentationValueImpl
	Type                         string
	Values                       map[string]interface{}
}

func (s RawGroupPolicyPresentationValueImpl) GroupPolicyPresentationValue() BaseGroupPolicyPresentationValueImpl {
	return s.groupPolicyPresentationValue
}

func (s RawGroupPolicyPresentationValueImpl) Entity() BaseEntityImpl {
	return s.groupPolicyPresentationValue.Entity()
}

var _ json.Marshaler = BaseGroupPolicyPresentationValueImpl{}

func (s BaseGroupPolicyPresentationValueImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseGroupPolicyPresentationValueImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseGroupPolicyPresentationValueImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseGroupPolicyPresentationValueImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.groupPolicyPresentationValue"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseGroupPolicyPresentationValueImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseGroupPolicyPresentationValueImpl{}

func (s *BaseGroupPolicyPresentationValueImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
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

	s.CreatedDateTime = decoded.CreatedDateTime
	s.DefinitionValue = decoded.DefinitionValue
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseGroupPolicyPresentationValueImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["presentation"]; ok {
		impl, err := UnmarshalGroupPolicyPresentationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Presentation' for 'BaseGroupPolicyPresentationValueImpl': %+v", err)
		}
		s.Presentation = &impl
	}

	return nil
}

func UnmarshalGroupPolicyPresentationValueImplementation(input []byte) (GroupPolicyPresentationValue, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling GroupPolicyPresentationValue into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.groupPolicyPresentationValueBoolean") {
		var out GroupPolicyPresentationValueBoolean
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupPolicyPresentationValueBoolean: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupPolicyPresentationValueDecimal") {
		var out GroupPolicyPresentationValueDecimal
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupPolicyPresentationValueDecimal: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupPolicyPresentationValueList") {
		var out GroupPolicyPresentationValueList
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupPolicyPresentationValueList: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupPolicyPresentationValueLongDecimal") {
		var out GroupPolicyPresentationValueLongDecimal
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupPolicyPresentationValueLongDecimal: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupPolicyPresentationValueMultiText") {
		var out GroupPolicyPresentationValueMultiText
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupPolicyPresentationValueMultiText: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupPolicyPresentationValueText") {
		var out GroupPolicyPresentationValueText
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupPolicyPresentationValueText: %+v", err)
		}
		return out, nil
	}

	var parent BaseGroupPolicyPresentationValueImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseGroupPolicyPresentationValueImpl: %+v", err)
	}

	return RawGroupPolicyPresentationValueImpl{
		groupPolicyPresentationValue: parent,
		Type:                         value,
		Values:                       temp,
	}, nil

}
