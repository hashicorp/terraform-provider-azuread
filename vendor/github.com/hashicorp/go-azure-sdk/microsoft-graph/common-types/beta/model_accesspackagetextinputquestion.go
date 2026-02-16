package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AccessPackageQuestion = AccessPackageTextInputQuestion{}

type AccessPackageTextInputQuestion struct {
	// Indicates whether the answer will be in single or multiple line format.
	IsSingleLineQuestion nullable.Type[bool] `json:"isSingleLineQuestion,omitempty"`

	// The regex pattern that the corresponding text answer must follow.
	RegexPattern nullable.Type[string] `json:"regexPattern,omitempty"`

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

func (s AccessPackageTextInputQuestion) AccessPackageQuestion() BaseAccessPackageQuestionImpl {
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

var _ json.Marshaler = AccessPackageTextInputQuestion{}

func (s AccessPackageTextInputQuestion) MarshalJSON() ([]byte, error) {
	type wrapper AccessPackageTextInputQuestion
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AccessPackageTextInputQuestion: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessPackageTextInputQuestion: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.accessPackageTextInputQuestion"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AccessPackageTextInputQuestion: %+v", err)
	}

	return encoded, nil
}
