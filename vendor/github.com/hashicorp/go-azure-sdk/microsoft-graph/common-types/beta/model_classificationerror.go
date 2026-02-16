package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClassificationError interface {
	ClassifcationErrorBase
	ClassificationError() BaseClassificationErrorImpl
}

var _ ClassificationError = BaseClassificationErrorImpl{}

type BaseClassificationErrorImpl struct {
	// A collection of more specific errors contributing to the overall error.
	Details *[]ClassifcationErrorBase `json:"details,omitempty"`

	// Fields inherited from ClassifcationErrorBase

	// A service-defined error code string.
	Code nullable.Type[string] `json:"code,omitempty"`

	// Contains more specific, potentially internal error details.
	InnerError *ClassificationInnerError `json:"innerError,omitempty"`

	// A human-readable representation of the error.
	Message nullable.Type[string] `json:"message,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The target of the error (for example, the specific property or item causing the issue).
	Target nullable.Type[string] `json:"target,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseClassificationErrorImpl) ClassificationError() BaseClassificationErrorImpl {
	return s
}

func (s BaseClassificationErrorImpl) ClassifcationErrorBase() BaseClassifcationErrorBaseImpl {
	return BaseClassifcationErrorBaseImpl{
		Code:       s.Code,
		InnerError: s.InnerError,
		Message:    s.Message,
		ODataId:    s.ODataId,
		ODataType:  s.ODataType,
		Target:     s.Target,
	}
}

var _ ClassificationError = RawClassificationErrorImpl{}

// RawClassificationErrorImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawClassificationErrorImpl struct {
	classificationError BaseClassificationErrorImpl
	Type                string
	Values              map[string]interface{}
}

func (s RawClassificationErrorImpl) ClassificationError() BaseClassificationErrorImpl {
	return s.classificationError
}

func (s RawClassificationErrorImpl) ClassifcationErrorBase() BaseClassifcationErrorBaseImpl {
	return s.classificationError.ClassifcationErrorBase()
}

var _ json.Marshaler = BaseClassificationErrorImpl{}

func (s BaseClassificationErrorImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseClassificationErrorImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseClassificationErrorImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseClassificationErrorImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.classificationError"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseClassificationErrorImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseClassificationErrorImpl{}

func (s *BaseClassificationErrorImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Code       nullable.Type[string]     `json:"code,omitempty"`
		InnerError *ClassificationInnerError `json:"innerError,omitempty"`
		Message    nullable.Type[string]     `json:"message,omitempty"`
		ODataId    *string                   `json:"@odata.id,omitempty"`
		ODataType  *string                   `json:"@odata.type,omitempty"`
		Target     nullable.Type[string]     `json:"target,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Code = decoded.Code
	s.InnerError = decoded.InnerError
	s.Message = decoded.Message
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Target = decoded.Target

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseClassificationErrorImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["details"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Details into list []json.RawMessage: %+v", err)
		}

		output := make([]ClassifcationErrorBase, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalClassifcationErrorBaseImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Details' for 'BaseClassificationErrorImpl': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Details = &output
	}

	return nil
}

func UnmarshalClassificationErrorImplementation(input []byte) (ClassificationError, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ClassificationError into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.processingError") {
		var out ProcessingError
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ProcessingError: %+v", err)
		}
		return out, nil
	}

	var parent BaseClassificationErrorImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseClassificationErrorImpl: %+v", err)
	}

	return RawClassificationErrorImpl{
		classificationError: parent,
		Type:                value,
		Values:              temp,
	}, nil

}
