package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AccessPackageQuestion = AccessPackageTextInputQuestion{}

type AccessPackageTextInputQuestion struct {
	// Indicates whether the answer is in single or multiple line format.
	IsSingleLineQuestion nullable.Type[bool] `json:"isSingleLineQuestion,omitempty"`

	// The regular expression pattern that any answer to this question must match.
	RegexPattern nullable.Type[string] `json:"regexPattern,omitempty"`

	// Fields inherited from AccessPackageQuestion

	// Specifies whether the requestor is allowed to edit answers to questions for an assignment by posting an update to
	// accessPackageAssignmentRequest.
	IsAnswerEditable nullable.Type[bool] `json:"isAnswerEditable,omitempty"`

	// Whether the requestor is required to supply an answer or not.
	IsRequired nullable.Type[bool] `json:"isRequired,omitempty"`

	// The text of the question represented in a format for a specific locale.
	Localizations *[]AccessPackageLocalizedText `json:"localizations,omitempty"`

	// Relative position of this question when displaying a list of questions to the requestor.
	Sequence nullable.Type[int64] `json:"sequence,omitempty"`

	// The text of the question to show to the requestor.
	Text nullable.Type[string] `json:"text,omitempty"`

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

func (s AccessPackageTextInputQuestion) AccessPackageQuestion() BaseAccessPackageQuestionImpl {
	return BaseAccessPackageQuestionImpl{
		IsAnswerEditable: s.IsAnswerEditable,
		IsRequired:       s.IsRequired,
		Localizations:    s.Localizations,
		Sequence:         s.Sequence,
		Text:             s.Text,
		Id:               s.Id,
		ODataId:          s.ODataId,
		ODataType:        s.ODataType,
	}
}

func (s AccessPackageTextInputQuestion) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
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
