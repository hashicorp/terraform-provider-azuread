package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ LearningCourseActivity = LearningSelfInitiatedCourse{}

type LearningSelfInitiatedCourse struct {
	// The date and time on which the learner started the self-initiated course. Optional.
	StartedDateTime nullable.Type[string] `json:"startedDateTime,omitempty"`

	// Fields inherited from LearningCourseActivity

	// Date and time when the assignment was completed. Optional.
	CompletedDateTime nullable.Type[string] `json:"completedDateTime,omitempty"`

	// The percentage completion value of the course activity. Optional.
	CompletionPercentage nullable.Type[int64] `json:"completionPercentage,omitempty"`

	ExternalcourseActivityId nullable.Type[string] `json:"externalcourseActivityId,omitempty"`

	// The user ID of the learner to whom the activity is assigned. Required.
	LearnerUserId string `json:"learnerUserId"`

	// The ID of the learning content created in Viva Learning. Required.
	LearningContentId string `json:"learningContentId"`

	// The registration ID of the provider. Required.
	LearningProviderId nullable.Type[string] `json:"learningProviderId,omitempty"`

	// The status of the course activity. Possible values are: notStarted, inProgress, completed. Required.
	Status CourseStatus `json:"status"`

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

func (s LearningSelfInitiatedCourse) LearningCourseActivity() BaseLearningCourseActivityImpl {
	return BaseLearningCourseActivityImpl{
		CompletedDateTime:        s.CompletedDateTime,
		CompletionPercentage:     s.CompletionPercentage,
		ExternalcourseActivityId: s.ExternalcourseActivityId,
		LearnerUserId:            s.LearnerUserId,
		LearningContentId:        s.LearningContentId,
		LearningProviderId:       s.LearningProviderId,
		Status:                   s.Status,
		Id:                       s.Id,
		ODataId:                  s.ODataId,
		ODataType:                s.ODataType,
	}
}

func (s LearningSelfInitiatedCourse) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = LearningSelfInitiatedCourse{}

func (s LearningSelfInitiatedCourse) MarshalJSON() ([]byte, error) {
	type wrapper LearningSelfInitiatedCourse
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling LearningSelfInitiatedCourse: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling LearningSelfInitiatedCourse: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.learningSelfInitiatedCourse"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling LearningSelfInitiatedCourse: %+v", err)
	}

	return encoded, nil
}
