package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = EducationAssignmentSettings{}

type EducationAssignmentSettings struct {
	DefaultGradingScheme *EducationGradingScheme `json:"defaultGradingScheme,omitempty"`

	// When set, enables users to weight assignments differently when computing a class average grade.
	GradingCategories *[]EducationGradingCategory `json:"gradingCategories,omitempty"`

	GradingSchemes *[]EducationGradingScheme `json:"gradingSchemes,omitempty"`

	// Indicates whether turn-in celebration animation is shown. If true, the animation doesn't show. The default value is
	// false.
	SubmissionAnimationDisabled nullable.Type[bool] `json:"submissionAnimationDisabled,omitempty"`

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

func (s EducationAssignmentSettings) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EducationAssignmentSettings{}

func (s EducationAssignmentSettings) MarshalJSON() ([]byte, error) {
	type wrapper EducationAssignmentSettings
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EducationAssignmentSettings: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EducationAssignmentSettings: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.educationAssignmentSettings"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EducationAssignmentSettings: %+v", err)
	}

	return encoded, nil
}
