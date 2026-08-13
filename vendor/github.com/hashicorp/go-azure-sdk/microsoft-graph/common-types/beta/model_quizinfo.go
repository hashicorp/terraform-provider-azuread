package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
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

// RawQuizInfoImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawQuizInfoImpl struct {
	quizInfo BaseQuizInfoImpl
	Type     string
	Values   map[string]interface{}
}

func (s RawQuizInfoImpl) QuizInfo() BaseQuizInfoImpl {
	return s.quizInfo
}

func (s RawQuizInfoImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
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
