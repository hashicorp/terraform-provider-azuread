package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = EducationRubric{}

type EducationRubric struct {
	// The user who created this resource.
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`

	// The timestamp type represents date and time information using ISO 8601 format and is always in UTC. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The description of this rubric.
	Description *EducationItemBody `json:"description,omitempty"`

	// The name of this rubric.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The grading type of this rubric. You can use null for a no-points rubric or educationAssignmentPointsGradeType for a
	// points rubric.
	Grading EducationAssignmentGradeType `json:"grading"`

	// The last user to modify the resource.
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`

	// Moment in time when the resource was last modified. The timestamp type represents date and time information using ISO
	// 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// The collection of levels making up this rubric.
	Levels *[]RubricLevel `json:"levels,omitempty"`

	// The collection of qualities making up this rubric.
	Qualities *[]RubricQuality `json:"qualities,omitempty"`

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

func (s EducationRubric) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EducationRubric{}

func (s EducationRubric) MarshalJSON() ([]byte, error) {
	type wrapper EducationRubric
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EducationRubric: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EducationRubric: %+v", err)
	}

	delete(decoded, "createdBy")
	delete(decoded, "createdDateTime")
	delete(decoded, "lastModifiedBy")
	delete(decoded, "lastModifiedDateTime")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.educationRubric"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EducationRubric: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &EducationRubric{}

func (s *EducationRubric) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CreatedDateTime      nullable.Type[string] `json:"createdDateTime,omitempty"`
		Description          *EducationItemBody    `json:"description,omitempty"`
		DisplayName          nullable.Type[string] `json:"displayName,omitempty"`
		LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`
		Levels               *[]RubricLevel        `json:"levels,omitempty"`
		Qualities            *[]RubricQuality      `json:"qualities,omitempty"`
		Id                   *string               `json:"id,omitempty"`
		ODataId              *string               `json:"@odata.id,omitempty"`
		ODataType            *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CreatedDateTime = decoded.CreatedDateTime
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.Levels = decoded.Levels
	s.Qualities = decoded.Qualities
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling EducationRubric into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'EducationRubric': %+v", err)
		}
		s.CreatedBy = &impl
	}

	if v, ok := temp["grading"]; ok {
		impl, err := UnmarshalEducationAssignmentGradeTypeImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Grading' for 'EducationRubric': %+v", err)
		}
		s.Grading = impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'EducationRubric': %+v", err)
		}
		s.LastModifiedBy = &impl
	}

	return nil
}
