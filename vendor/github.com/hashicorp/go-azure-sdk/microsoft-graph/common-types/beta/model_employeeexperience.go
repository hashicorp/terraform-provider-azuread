package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = EmployeeExperience{}

type EmployeeExperience struct {
	// A collection of communities in Viva Engage.
	Communities *[]Community `json:"communities,omitempty"`

	// A collection of long-running, asynchronous operations related to Viva Engage.
	EngagementAsyncOperations *[]EngagementAsyncOperation `json:"engagementAsyncOperations,omitempty"`

	// Represents a collection of goals in a Viva Goals organization.
	Goals *Goals `json:"goals,omitempty"`

	LearningCourseActivities *[]LearningCourseActivity `json:"learningCourseActivities,omitempty"`

	// A collection of learning providers.
	LearningProviders *[]LearningProvider `json:"learningProviders,omitempty"`

	// A collection of roles in Viva Engage.
	Roles *[]EngagementRole `json:"roles,omitempty"`

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

func (s EmployeeExperience) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EmployeeExperience{}

func (s EmployeeExperience) MarshalJSON() ([]byte, error) {
	type wrapper EmployeeExperience
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EmployeeExperience: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EmployeeExperience: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.employeeExperience"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EmployeeExperience: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &EmployeeExperience{}

func (s *EmployeeExperience) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Communities               *[]Community                `json:"communities,omitempty"`
		EngagementAsyncOperations *[]EngagementAsyncOperation `json:"engagementAsyncOperations,omitempty"`
		Goals                     *Goals                      `json:"goals,omitempty"`
		LearningProviders         *[]LearningProvider         `json:"learningProviders,omitempty"`
		Roles                     *[]EngagementRole           `json:"roles,omitempty"`
		Id                        *string                     `json:"id,omitempty"`
		ODataId                   *string                     `json:"@odata.id,omitempty"`
		ODataType                 *string                     `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Communities = decoded.Communities
	s.EngagementAsyncOperations = decoded.EngagementAsyncOperations
	s.Goals = decoded.Goals
	s.LearningProviders = decoded.LearningProviders
	s.Roles = decoded.Roles
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling EmployeeExperience into map[string]json.RawMessage: %+v", err)
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
				return fmt.Errorf("unmarshaling index %d field 'LearningCourseActivities' for 'EmployeeExperience': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.LearningCourseActivities = &output
	}

	return nil
}
