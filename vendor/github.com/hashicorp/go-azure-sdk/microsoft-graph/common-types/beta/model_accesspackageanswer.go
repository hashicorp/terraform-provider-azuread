package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAnswer interface {
	AccessPackageAnswer() BaseAccessPackageAnswerImpl
}

var _ AccessPackageAnswer = BaseAccessPackageAnswerImpl{}

type BaseAccessPackageAnswerImpl struct {
	// The question the answer is for. Required and Read-only.
	AnsweredQuestion *AccessPackageQuestion `json:"answeredQuestion,omitempty"`

	// The display value of the answer. Required.
	DisplayValue nullable.Type[string] `json:"displayValue,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseAccessPackageAnswerImpl) AccessPackageAnswer() BaseAccessPackageAnswerImpl {
	return s
}

var _ AccessPackageAnswer = RawAccessPackageAnswerImpl{}

// RawAccessPackageAnswerImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAccessPackageAnswerImpl struct {
	accessPackageAnswer BaseAccessPackageAnswerImpl
	Type                string
	Values              map[string]interface{}
}

func (s RawAccessPackageAnswerImpl) AccessPackageAnswer() BaseAccessPackageAnswerImpl {
	return s.accessPackageAnswer
}

var _ json.Marshaler = BaseAccessPackageAnswerImpl{}

func (s BaseAccessPackageAnswerImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseAccessPackageAnswerImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseAccessPackageAnswerImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseAccessPackageAnswerImpl: %+v", err)
	}

	delete(decoded, "answeredQuestion")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseAccessPackageAnswerImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseAccessPackageAnswerImpl{}

func (s *BaseAccessPackageAnswerImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DisplayValue nullable.Type[string] `json:"displayValue,omitempty"`
		ODataId      *string               `json:"@odata.id,omitempty"`
		ODataType    *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DisplayValue = decoded.DisplayValue
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseAccessPackageAnswerImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["answeredQuestion"]; ok {
		impl, err := UnmarshalAccessPackageQuestionImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'AnsweredQuestion' for 'BaseAccessPackageAnswerImpl': %+v", err)
		}
		s.AnsweredQuestion = &impl
	}

	return nil
}

func UnmarshalAccessPackageAnswerImplementation(input []byte) (AccessPackageAnswer, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessPackageAnswer into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.accessPackageAnswerString") {
		var out AccessPackageAnswerString
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessPackageAnswerString: %+v", err)
		}
		return out, nil
	}

	var parent BaseAccessPackageAnswerImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAccessPackageAnswerImpl: %+v", err)
	}

	return RawAccessPackageAnswerImpl{
		accessPackageAnswer: parent,
		Type:                value,
		Values:              temp,
	}, nil

}
