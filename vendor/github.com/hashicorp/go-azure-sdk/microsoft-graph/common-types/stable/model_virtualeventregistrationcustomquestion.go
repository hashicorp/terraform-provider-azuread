package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ VirtualEventRegistrationQuestionBase = VirtualEventRegistrationCustomQuestion{}

type VirtualEventRegistrationCustomQuestion struct {
	// Answer choices when answerInputType is singleChoice or multiChoice.
	AnswerChoices *[]string `json:"answerChoices,omitempty"`

	// Input type of the registration question answer. Possible values are text, multilineText, singleChoice, multiChoice,
	// boolean, and unknownFutureValue.
	AnswerInputType *VirtualEventRegistrationQuestionAnswerInputType `json:"answerInputType,omitempty"`

	// Fields inherited from VirtualEventRegistrationQuestionBase

	// Display name of the registration question.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Indicates whether an answer to the question is required. The default value is false.
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

func (s VirtualEventRegistrationCustomQuestion) VirtualEventRegistrationQuestionBase() BaseVirtualEventRegistrationQuestionBaseImpl {
	return BaseVirtualEventRegistrationQuestionBaseImpl{
		DisplayName: s.DisplayName,
		IsRequired:  s.IsRequired,
		Id:          s.Id,
		ODataId:     s.ODataId,
		ODataType:   s.ODataType,
	}
}

func (s VirtualEventRegistrationCustomQuestion) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = VirtualEventRegistrationCustomQuestion{}

func (s VirtualEventRegistrationCustomQuestion) MarshalJSON() ([]byte, error) {
	type wrapper VirtualEventRegistrationCustomQuestion
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling VirtualEventRegistrationCustomQuestion: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling VirtualEventRegistrationCustomQuestion: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.virtualEventRegistrationCustomQuestion"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling VirtualEventRegistrationCustomQuestion: %+v", err)
	}

	return encoded, nil
}
