package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EducationAssignmentGradeType = EducationAssignmentPointsGradeType{}

type EducationAssignmentPointsGradeType struct {

	// Fields inherited from EducationAssignmentGradeType

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s EducationAssignmentPointsGradeType) EducationAssignmentGradeType() BaseEducationAssignmentGradeTypeImpl {
	return BaseEducationAssignmentGradeTypeImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EducationAssignmentPointsGradeType{}

func (s EducationAssignmentPointsGradeType) MarshalJSON() ([]byte, error) {
	type wrapper EducationAssignmentPointsGradeType
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EducationAssignmentPointsGradeType: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EducationAssignmentPointsGradeType: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.educationAssignmentPointsGradeType"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EducationAssignmentPointsGradeType: %+v", err)
	}

	return encoded, nil
}
