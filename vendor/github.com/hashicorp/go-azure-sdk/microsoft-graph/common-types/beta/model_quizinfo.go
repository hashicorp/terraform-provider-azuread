package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QuizInfo interface {
	QuizInfo() BaseQuizInfoImpl
}

var _ QuizInfo = BaseQuizInfoImpl{}

type BaseQuizInfoImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseQuizInfoImpl) QuizInfo() BaseQuizInfoImpl {
	return s
}

var _ QuizInfo = RawQuizInfoImpl{}

// RawQuizInfoImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawQuizInfoImpl struct {
	quizInfo BaseQuizInfoImpl
	Type     string
	Values   map[string]interface{}
}

func (s RawQuizInfoImpl) QuizInfo() BaseQuizInfoImpl {
	return s.quizInfo
}

func UnmarshalQuizInfoImplementation(input []byte) (QuizInfo, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling QuizInfo into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.matrixChoiceGroupQuizInfo") {
		var out MatrixChoiceGroupQuizInfo
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MatrixChoiceGroupQuizInfo: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.npsQuizInfo") {
		var out NpsQuizInfo
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NpsQuizInfo: %+v", err)
		}
		return out, nil
	}

	var parent BaseQuizInfoImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseQuizInfoImpl: %+v", err)
	}

	return RawQuizInfoImpl{
		quizInfo: parent,
		Type:     value,
		Values:   temp,
	}, nil

}
