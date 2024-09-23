package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EducationSynchronizationCustomizationsBase = EducationSynchronizationCustomizations{}

type EducationSynchronizationCustomizations struct {
	// Customizations for School entities.
	School *EducationSynchronizationCustomization `json:"school,omitempty"`

	// Customizations for Section entities.
	Section *EducationSynchronizationCustomization `json:"section,omitempty"`

	// Customizations for Student entities.
	Student *EducationSynchronizationCustomization `json:"student,omitempty"`

	// Customizations for Student Enrollments.
	StudentEnrollment *EducationSynchronizationCustomization `json:"studentEnrollment,omitempty"`

	// Customizations for Teacher entities.
	Teacher *EducationSynchronizationCustomization `json:"teacher,omitempty"`

	// Customizations for Teacher Rosters.
	TeacherRoster *EducationSynchronizationCustomization `json:"teacherRoster,omitempty"`

	// Fields inherited from EducationSynchronizationCustomizationsBase

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s EducationSynchronizationCustomizations) EducationSynchronizationCustomizationsBase() BaseEducationSynchronizationCustomizationsBaseImpl {
	return BaseEducationSynchronizationCustomizationsBaseImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EducationSynchronizationCustomizations{}

func (s EducationSynchronizationCustomizations) MarshalJSON() ([]byte, error) {
	type wrapper EducationSynchronizationCustomizations
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EducationSynchronizationCustomizations: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EducationSynchronizationCustomizations: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.educationSynchronizationCustomizations"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EducationSynchronizationCustomizations: %+v", err)
	}

	return encoded, nil
}
