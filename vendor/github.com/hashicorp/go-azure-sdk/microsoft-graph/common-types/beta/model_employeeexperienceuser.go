package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = EmployeeExperienceUser{}

type EmployeeExperienceUser struct {
	// Represents the collection of Viva Engage roles assigned to a user.
	AssignedRoles *[]EngagementRole `json:"assignedRoles,omitempty"`

	LearningCourseActivities *[]LearningCourseActivity `json:"learningCourseActivities,omitempty"`

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

func (s EmployeeExperienceUser) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EmployeeExperienceUser{}

func (s EmployeeExperienceUser) MarshalJSON() ([]byte, error) {
	type wrapper EmployeeExperienceUser
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EmployeeExperienceUser: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EmployeeExperienceUser: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.employeeExperienceUser"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EmployeeExperienceUser: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &EmployeeExperienceUser{}

func (s *EmployeeExperienceUser) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AssignedRoles *[]EngagementRole `json:"assignedRoles,omitempty"`
		Id            *string           `json:"id,omitempty"`
		ODataId       *string           `json:"@odata.id,omitempty"`
		ODataType     *string           `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AssignedRoles = decoded.AssignedRoles
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling EmployeeExperienceUser into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["learningCourseActivities"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling LearningCourseActivities into list []json.RawMessage: %+v", err)
		}

		output := make([]LearningCourseActivity, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalLearningCourseActivityImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'LearningCourseActivities' for 'EmployeeExperienceUser': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.LearningCourseActivities = &output
	}

	return nil
}
