package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AccessPackageAnswer = AccessPackageAnswerString{}

type AccessPackageAnswerString struct {
	// The value stored on the requestor's user profile, if this answer is configured to be stored as a specific attribute.
	Value nullable.Type[string] `json:"value,omitempty"`

	// Fields inherited from AccessPackageAnswer

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

func (s AccessPackageAnswerString) AccessPackageAnswer() BaseAccessPackageAnswerImpl {
	return BaseAccessPackageAnswerImpl{
		AnsweredQuestion: s.AnsweredQuestion,
		DisplayValue:     s.DisplayValue,
		ODataId:          s.ODataId,
		ODataType:        s.ODataType,
	}
}

var _ json.Marshaler = AccessPackageAnswerString{}

func (s AccessPackageAnswerString) MarshalJSON() ([]byte, error) {
	type wrapper AccessPackageAnswerString
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AccessPackageAnswerString: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessPackageAnswerString: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.accessPackageAnswerString"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AccessPackageAnswerString: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &AccessPackageAnswerString{}

func (s *AccessPackageAnswerString) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Value        nullable.Type[string] `json:"value,omitempty"`
		DisplayValue nullable.Type[string] `json:"displayValue,omitempty"`
		ODataId      *string               `json:"@odata.id,omitempty"`
		ODataType    *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Value = decoded.Value
	s.DisplayValue = decoded.DisplayValue
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AccessPackageAnswerString into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["answeredQuestion"]; ok {
		impl, err := UnmarshalAccessPackageQuestionImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'AnsweredQuestion' for 'AccessPackageAnswerString': %+v", err)
		}
		s.AnsweredQuestion = &impl
	}

	return nil
}
