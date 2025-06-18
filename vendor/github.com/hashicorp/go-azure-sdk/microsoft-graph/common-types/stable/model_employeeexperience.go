package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmployeeExperience struct {
	// A collection of communities in Viva Engage.
	Communities *[]Community `json:"communities,omitempty"`

	// A collection of long-running, asynchronous operations related to Viva Engage.
	EngagementAsyncOperations *[]EngagementAsyncOperation `json:"engagementAsyncOperations,omitempty"`

	LearningCourseActivities *[]LearningCourseActivity `json:"learningCourseActivities,omitempty"`

	// A collection of learning providers.
	LearningProviders *[]LearningProvider `json:"learningProviders,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

var _ json.Unmarshaler = &EmployeeExperience{}

func (s *EmployeeExperience) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Communities               *[]Community                `json:"communities,omitempty"`
		EngagementAsyncOperations *[]EngagementAsyncOperation `json:"engagementAsyncOperations,omitempty"`
		LearningProviders         *[]LearningProvider         `json:"learningProviders,omitempty"`
		ODataId                   *string                     `json:"@odata.id,omitempty"`
		ODataType                 *string                     `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Communities = decoded.Communities
	s.EngagementAsyncOperations = decoded.EngagementAsyncOperations
	s.LearningProviders = decoded.LearningProviders
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
