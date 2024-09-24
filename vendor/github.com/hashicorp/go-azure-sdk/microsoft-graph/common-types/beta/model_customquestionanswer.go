package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CustomQuestionAnswer{}

type CustomQuestionAnswer struct {
	// Display name of the custom registration question. Read-only.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// ID the custom registration question. Read-only.
	QuestionId nullable.Type[string] `json:"questionId,omitempty"`

	// Answer to the custom registration question.
	Value nullable.Type[string] `json:"value,omitempty"`

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

func (s CustomQuestionAnswer) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CustomQuestionAnswer{}

func (s CustomQuestionAnswer) MarshalJSON() ([]byte, error) {
	type wrapper CustomQuestionAnswer
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CustomQuestionAnswer: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CustomQuestionAnswer: %+v", err)
	}

	delete(decoded, "displayName")
	delete(decoded, "questionId")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.customQuestionAnswer"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CustomQuestionAnswer: %+v", err)
	}

	return encoded, nil
}
