package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EducationOutcome = EducationFeedbackOutcome{}

type EducationFeedbackOutcome struct {
	// Teacher's written feedback to the student.
	Feedback *EducationFeedback `json:"feedback,omitempty"`

	// A copy of the feedback property that is made when the grade is released to the student.
	PublishedFeedback *EducationFeedback `json:"publishedFeedback,omitempty"`

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

func (s EducationFeedbackOutcome) EducationOutcome() BaseEducationOutcomeImpl {
	return BaseEducationOutcomeImpl{
		LastModifiedBy:       s.LastModifiedBy,
		LastModifiedDateTime: s.LastModifiedDateTime,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s EducationFeedbackOutcome) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EducationFeedbackOutcome{}

func (s EducationFeedbackOutcome) MarshalJSON() ([]byte, error) {
	type wrapper EducationFeedbackOutcome
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EducationFeedbackOutcome: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EducationFeedbackOutcome: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.educationFeedbackOutcome"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EducationFeedbackOutcome: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &EducationFeedbackOutcome{}

func (s *EducationFeedbackOutcome) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Feedback             *EducationFeedback    `json:"feedback,omitempty"`
		PublishedFeedback    *EducationFeedback    `json:"publishedFeedback,omitempty"`
		LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`
		Id                   *string               `json:"id,omitempty"`
		ODataId              *string               `json:"@odata.id,omitempty"`
		ODataType            *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Feedback = decoded.Feedback
	s.PublishedFeedback = decoded.PublishedFeedback
	s.Id = decoded.Id
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling EducationFeedbackOutcome into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'EducationFeedbackOutcome': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	return nil
}
