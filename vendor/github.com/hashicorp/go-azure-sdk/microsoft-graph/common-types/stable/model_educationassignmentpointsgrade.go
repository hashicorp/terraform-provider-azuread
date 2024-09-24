package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EducationAssignmentGrade = EducationAssignmentPointsGrade{}

type EducationAssignmentPointsGrade struct {

	// Fields inherited from EducationAssignmentGrade

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

func (s EducationAssignmentPointsGrade) EducationAssignmentGrade() BaseEducationAssignmentGradeImpl {
	return BaseEducationAssignmentGradeImpl{
		GradedBy:       s.GradedBy,
		GradedDateTime: s.GradedDateTime,
		ODataId:        s.ODataId,
		ODataType:      s.ODataType,
	}
}

var _ json.Marshaler = EducationAssignmentPointsGrade{}

func (s EducationAssignmentPointsGrade) MarshalJSON() ([]byte, error) {
	type wrapper EducationAssignmentPointsGrade
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EducationAssignmentPointsGrade: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EducationAssignmentPointsGrade: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.educationAssignmentPointsGrade"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EducationAssignmentPointsGrade: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &EducationAssignmentPointsGrade{}

func (s *EducationAssignmentPointsGrade) UnmarshalJSON(bytes []byte) error {
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
		return fmt.Errorf("unmarshaling EducationAssignmentPointsGrade into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["gradedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'GradedBy' for 'EducationAssignmentPointsGrade': %+v", err)
		}
		s.GradedBy = impl
	}

	return nil
}
