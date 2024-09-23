package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RubricLevel struct {
	// The description of this rubric level.
	Description *EducationItemBody `json:"description,omitempty"`

	// The name of this rubric level.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Null if this is a no-points rubric; educationAssignmentPointsGradeType if it's a points rubric.
	Grading EducationAssignmentGradeType `json:"grading"`

	// The ID of this resource.
	LevelId nullable.Type[string] `json:"levelId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

var _ json.Unmarshaler = &RubricLevel{}

func (s *RubricLevel) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Description *EducationItemBody    `json:"description,omitempty"`
		DisplayName nullable.Type[string] `json:"displayName,omitempty"`
		LevelId     nullable.Type[string] `json:"levelId,omitempty"`
		ODataId     *string               `json:"@odata.id,omitempty"`
		ODataType   *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.LevelId = decoded.LevelId
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling RubricLevel into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["grading"]; ok {
		impl, err := UnmarshalEducationAssignmentGradeTypeImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Grading' for 'RubricLevel': %+v", err)
		}
		s.Grading = impl
	}

	return nil
}
