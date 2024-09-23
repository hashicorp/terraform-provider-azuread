package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationOutcome interface {
	Entity
	EducationOutcome() BaseEducationOutcomeImpl
}

var _ EducationOutcome = BaseEducationOutcomeImpl{}

type BaseEducationOutcomeImpl struct {
	// The individual who updated the resource.
	LastModifiedBy IdentitySet `json:"lastModifiedBy"`

	// The moment in time when the resource was last modified. The Timestamp type represents date and time information using
	// ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2021 is 2021-01-01T00:00:00Z.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

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

func (s BaseEducationOutcomeImpl) EducationOutcome() BaseEducationOutcomeImpl {
	return s
}

func (s BaseEducationOutcomeImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ EducationOutcome = RawEducationOutcomeImpl{}

// RawEducationOutcomeImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawEducationOutcomeImpl struct {
	educationOutcome BaseEducationOutcomeImpl
	Type             string
	Values           map[string]interface{}
}

func (s RawEducationOutcomeImpl) EducationOutcome() BaseEducationOutcomeImpl {
	return s.educationOutcome
}

func (s RawEducationOutcomeImpl) Entity() BaseEntityImpl {
	return s.educationOutcome.Entity()
}

var _ json.Marshaler = BaseEducationOutcomeImpl{}

func (s BaseEducationOutcomeImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseEducationOutcomeImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseEducationOutcomeImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseEducationOutcomeImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.educationOutcome"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseEducationOutcomeImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseEducationOutcomeImpl{}

func (s *BaseEducationOutcomeImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`
		Id                   *string               `json:"id,omitempty"`
		ODataId              *string               `json:"@odata.id,omitempty"`
		ODataType            *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseEducationOutcomeImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'BaseEducationOutcomeImpl': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	return nil
}

func UnmarshalEducationOutcomeImplementation(input []byte) (EducationOutcome, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling EducationOutcome into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.educationFeedbackOutcome") {
		var out EducationFeedbackOutcome
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationFeedbackOutcome: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.educationFeedbackResourceOutcome") {
		var out EducationFeedbackResourceOutcome
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationFeedbackResourceOutcome: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.educationPointsOutcome") {
		var out EducationPointsOutcome
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationPointsOutcome: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.educationRubricOutcome") {
		var out EducationRubricOutcome
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationRubricOutcome: %+v", err)
		}
		return out, nil
	}

	var parent BaseEducationOutcomeImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseEducationOutcomeImpl: %+v", err)
	}

	return RawEducationOutcomeImpl{
		educationOutcome: parent,
		Type:             value,
		Values:           temp,
	}, nil

}
