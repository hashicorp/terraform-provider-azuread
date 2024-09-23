package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EducationOutcome = EducationRubricOutcome{}

type EducationRubricOutcome struct {
	// A copy of the rubricQualityFeedback property that is made when the grade is released to the student.
	PublishedRubricQualityFeedback *[]RubricQualityFeedbackModel `json:"publishedRubricQualityFeedback,omitempty"`

	// A copy of the rubricQualitySelectedLevels property that is made when the grade is released to the student.
	PublishedRubricQualitySelectedLevels *[]RubricQualitySelectedColumnModel `json:"publishedRubricQualitySelectedLevels,omitempty"`

	// A collection of specific feedback for each quality of this rubric.
	RubricQualityFeedback *[]RubricQualityFeedbackModel `json:"rubricQualityFeedback,omitempty"`

	// The level that the teacher has selected for each quality while grading this assignment.
	RubricQualitySelectedLevels *[]RubricQualitySelectedColumnModel `json:"rubricQualitySelectedLevels,omitempty"`

	// Fields inherited from EducationOutcome

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

func (s EducationRubricOutcome) EducationOutcome() BaseEducationOutcomeImpl {
	return BaseEducationOutcomeImpl{
		LastModifiedBy:       s.LastModifiedBy,
		LastModifiedDateTime: s.LastModifiedDateTime,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s EducationRubricOutcome) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EducationRubricOutcome{}

func (s EducationRubricOutcome) MarshalJSON() ([]byte, error) {
	type wrapper EducationRubricOutcome
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EducationRubricOutcome: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EducationRubricOutcome: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.educationRubricOutcome"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EducationRubricOutcome: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &EducationRubricOutcome{}

func (s *EducationRubricOutcome) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		PublishedRubricQualityFeedback       *[]RubricQualityFeedbackModel       `json:"publishedRubricQualityFeedback,omitempty"`
		PublishedRubricQualitySelectedLevels *[]RubricQualitySelectedColumnModel `json:"publishedRubricQualitySelectedLevels,omitempty"`
		RubricQualityFeedback                *[]RubricQualityFeedbackModel       `json:"rubricQualityFeedback,omitempty"`
		RubricQualitySelectedLevels          *[]RubricQualitySelectedColumnModel `json:"rubricQualitySelectedLevels,omitempty"`
		LastModifiedDateTime                 nullable.Type[string]               `json:"lastModifiedDateTime,omitempty"`
		Id                                   *string                             `json:"id,omitempty"`
		ODataId                              *string                             `json:"@odata.id,omitempty"`
		ODataType                            *string                             `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.PublishedRubricQualityFeedback = decoded.PublishedRubricQualityFeedback
	s.PublishedRubricQualitySelectedLevels = decoded.PublishedRubricQualitySelectedLevels
	s.RubricQualityFeedback = decoded.RubricQualityFeedback
	s.RubricQualitySelectedLevels = decoded.RubricQualitySelectedLevels
	s.Id = decoded.Id
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling EducationRubricOutcome into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'EducationRubricOutcome': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	return nil
}
