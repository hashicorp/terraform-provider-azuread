package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AccessPackageQuestion = AccessPackageMultipleChoiceQuestion{}

type AccessPackageMultipleChoiceQuestion struct {
	// Indicates whether requestor can select multiple choices as their answer.
	AllowsMultipleSelection nullable.Type[bool] `json:"allowsMultipleSelection,omitempty"`

	// List of answer choices.
	Choices *[]AccessPackageAnswerChoice `json:"choices,omitempty"`

	// Fields inherited from AccessPackageQuestion

	// ID of the question.
	Id nullable.Type[string] `json:"id,omitempty"`

	// Specifies whether the requestor is allowed to edit answers to questions.
	IsAnswerEditable nullable.Type[bool] `json:"isAnswerEditable,omitempty"`

	// Whether the requestor is required to supply an answer or not.
	IsRequired nullable.Type[bool] `json:"isRequired,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Relative position of this question when displaying a list of questions to the requestor.
	Sequence nullable.Type[int64] `json:"sequence,omitempty"`

	// The text of the question to show to the requestor.
	Text *AccessPackageLocalizedContent `json:"text,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s AccessPackageMultipleChoiceQuestion) AccessPackageQuestion() BaseAccessPackageQuestionImpl {
	return BaseAccessPackageQuestionImpl{
		Id:               s.Id,
		IsAnswerEditable: s.IsAnswerEditable,
		IsRequired:       s.IsRequired,
		ODataId:          s.ODataId,
		ODataType:        s.ODataType,
		Sequence:         s.Sequence,
		Text:             s.Text,
	}
}

var _ json.Marshaler = AccessPackageMultipleChoiceQuestion{}

func (s AccessPackageMultipleChoiceQuestion) MarshalJSON() ([]byte, error) {
	type wrapper AccessPackageMultipleChoiceQuestion
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AccessPackageMultipleChoiceQuestion: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessPackageMultipleChoiceQuestion: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.accessPackageMultipleChoiceQuestion"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AccessPackageMultipleChoiceQuestion: %+v", err)
	}

	return encoded, nil
}
