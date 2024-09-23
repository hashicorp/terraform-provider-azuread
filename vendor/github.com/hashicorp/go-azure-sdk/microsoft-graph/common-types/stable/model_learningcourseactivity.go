package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LearningCourseActivity interface {
	Entity
	LearningCourseActivity() BaseLearningCourseActivityImpl
}

var _ LearningCourseActivity = BaseLearningCourseActivityImpl{}

type BaseLearningCourseActivityImpl struct {
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

func (s BaseLearningCourseActivityImpl) LearningCourseActivity() BaseLearningCourseActivityImpl {
	return s
}

func (s BaseLearningCourseActivityImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ LearningCourseActivity = RawLearningCourseActivityImpl{}

// RawLearningCourseActivityImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawLearningCourseActivityImpl struct {
	learningCourseActivity BaseLearningCourseActivityImpl
	Type                   string
	Values                 map[string]interface{}
}

func (s RawLearningCourseActivityImpl) LearningCourseActivity() BaseLearningCourseActivityImpl {
	return s.learningCourseActivity
}

func (s RawLearningCourseActivityImpl) Entity() BaseEntityImpl {
	return s.learningCourseActivity.Entity()
}

var _ json.Marshaler = BaseLearningCourseActivityImpl{}

func (s BaseLearningCourseActivityImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseLearningCourseActivityImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseLearningCourseActivityImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseLearningCourseActivityImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.learningCourseActivity"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseLearningCourseActivityImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalLearningCourseActivityImplementation(input []byte) (LearningCourseActivity, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling LearningCourseActivity into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.learningAssignment") {
		var out LearningAssignment
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LearningAssignment: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.learningSelfInitiatedCourse") {
		var out LearningSelfInitiatedCourse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LearningSelfInitiatedCourse: %+v", err)
		}
		return out, nil
	}

	var parent BaseLearningCourseActivityImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseLearningCourseActivityImpl: %+v", err)
	}

	return RawLearningCourseActivityImpl{
		learningCourseActivity: parent,
		Type:                   value,
		Values:                 temp,
	}, nil

}
