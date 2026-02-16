package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageQuestion interface {
	AccessPackageQuestion() BaseAccessPackageQuestionImpl
}

var _ AccessPackageQuestion = BaseAccessPackageQuestionImpl{}

type BaseAccessPackageQuestionImpl struct {
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

func (s BaseAccessPackageQuestionImpl) AccessPackageQuestion() BaseAccessPackageQuestionImpl {
	return s
}

var _ AccessPackageQuestion = RawAccessPackageQuestionImpl{}

// RawAccessPackageQuestionImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAccessPackageQuestionImpl struct {
	accessPackageQuestion BaseAccessPackageQuestionImpl
	Type                  string
	Values                map[string]interface{}
}

func (s RawAccessPackageQuestionImpl) AccessPackageQuestion() BaseAccessPackageQuestionImpl {
	return s.accessPackageQuestion
}

func UnmarshalAccessPackageQuestionImplementation(input []byte) (AccessPackageQuestion, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessPackageQuestion into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.accessPackageMultipleChoiceQuestion") {
		var out AccessPackageMultipleChoiceQuestion
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessPackageMultipleChoiceQuestion: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.accessPackageTextInputQuestion") {
		var out AccessPackageTextInputQuestion
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessPackageTextInputQuestion: %+v", err)
		}
		return out, nil
	}

	var parent BaseAccessPackageQuestionImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAccessPackageQuestionImpl: %+v", err)
	}

	return RawAccessPackageQuestionImpl{
		accessPackageQuestion: parent,
		Type:                  value,
		Values:                temp,
	}, nil

}
