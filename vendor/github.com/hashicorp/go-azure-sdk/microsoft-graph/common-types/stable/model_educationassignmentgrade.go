package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationAssignmentGrade interface {
	EducationAssignmentGrade() BaseEducationAssignmentGradeImpl
}

var _ EducationAssignmentGrade = BaseEducationAssignmentGradeImpl{}

type BaseEducationAssignmentGradeImpl struct {
	// User who did the grading.
	GradedBy IdentitySet `json:"gradedBy"`

	// Moment in time when the grade was applied to this submission object. The Timestamp type represents date and time
	// information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z
	GradedDateTime nullable.Type[string] `json:"gradedDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseEducationAssignmentGradeImpl) EducationAssignmentGrade() BaseEducationAssignmentGradeImpl {
	return s
}

var _ EducationAssignmentGrade = RawEducationAssignmentGradeImpl{}

// RawEducationAssignmentGradeImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawEducationAssignmentGradeImpl struct {
	educationAssignmentGrade BaseEducationAssignmentGradeImpl
	Type                     string
	Values                   map[string]interface{}
}

func (s RawEducationAssignmentGradeImpl) EducationAssignmentGrade() BaseEducationAssignmentGradeImpl {
	return s.educationAssignmentGrade
}

var _ json.Unmarshaler = &BaseEducationAssignmentGradeImpl{}

func (s *BaseEducationAssignmentGradeImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		GradedDateTime nullable.Type[string] `json:"gradedDateTime,omitempty"`
		ODataId        *string               `json:"@odata.id,omitempty"`
		ODataType      *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.GradedDateTime = decoded.GradedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseEducationAssignmentGradeImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["gradedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'GradedBy' for 'BaseEducationAssignmentGradeImpl': %+v", err)
		}
		s.GradedBy = impl
	}

	return nil
}

func UnmarshalEducationAssignmentGradeImplementation(input []byte) (EducationAssignmentGrade, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling EducationAssignmentGrade into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.educationAssignmentPointsGrade") {
		var out EducationAssignmentPointsGrade
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationAssignmentPointsGrade: %+v", err)
		}
		return out, nil
	}

	var parent BaseEducationAssignmentGradeImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseEducationAssignmentGradeImpl: %+v", err)
	}

	return RawEducationAssignmentGradeImpl{
		educationAssignmentGrade: parent,
		Type:                     value,
		Values:                   temp,
	}, nil

}
