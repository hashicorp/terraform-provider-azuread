package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = MeetingRegistrationQuestion{}

type MeetingRegistrationQuestion struct {
	// Answer input type of the custom registration question.
	AnswerInputType *AnswerInputType `json:"answerInputType,omitempty"`

	// Answer options when answerInputType is radioButton.
	AnswerOptions *[]string `json:"answerOptions,omitempty"`

	// Display name of the custom registration question.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Indicates whether the question is required. Default value is false.
	IsRequired nullable.Type[bool] `json:"isRequired,omitempty"`

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

func (s MeetingRegistrationQuestion) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MeetingRegistrationQuestion{}

func (s MeetingRegistrationQuestion) MarshalJSON() ([]byte, error) {
	type wrapper MeetingRegistrationQuestion
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MeetingRegistrationQuestion: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MeetingRegistrationQuestion: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.meetingRegistrationQuestion"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MeetingRegistrationQuestion: %+v", err)
	}

	return encoded, nil
}
