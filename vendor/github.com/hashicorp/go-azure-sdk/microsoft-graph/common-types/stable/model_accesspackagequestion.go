package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageQuestion interface {
	Entity
	AccessPackageQuestion() BaseAccessPackageQuestionImpl
}

var _ AccessPackageQuestion = BaseAccessPackageQuestionImpl{}

type BaseAccessPackageQuestionImpl struct {
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

func (s BaseAccessPackageQuestionImpl) AccessPackageQuestion() BaseAccessPackageQuestionImpl {
	return s
}

func (s BaseAccessPackageQuestionImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
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

func (s RawAccessPackageQuestionImpl) Entity() BaseEntityImpl {
	return s.accessPackageQuestion.Entity()
}

var _ json.Marshaler = BaseAccessPackageQuestionImpl{}

func (s BaseAccessPackageQuestionImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseAccessPackageQuestionImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseAccessPackageQuestionImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseAccessPackageQuestionImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.accessPackageQuestion"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseAccessPackageQuestionImpl: %+v", err)
	}

	return encoded, nil
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
