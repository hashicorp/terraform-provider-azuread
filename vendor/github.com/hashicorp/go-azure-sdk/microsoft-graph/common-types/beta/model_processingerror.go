package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ClassificationError = ProcessingError{}

type ProcessingError struct {
	ErrorType *ContentProcessingErrorType `json:"errorType,omitempty"`

	// Fields inherited from ClassificationError

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

func (s ProcessingError) ClassificationError() BaseClassificationErrorImpl {
	return BaseClassificationErrorImpl{
		Details:    s.Details,
		Code:       s.Code,
		InnerError: s.InnerError,
		Message:    s.Message,
		ODataId:    s.ODataId,
		ODataType:  s.ODataType,
		Target:     s.Target,
	}
}

func (s ProcessingError) ClassifcationErrorBase() BaseClassifcationErrorBaseImpl {
	return BaseClassifcationErrorBaseImpl{
		Code:       s.Code,
		InnerError: s.InnerError,
		Message:    s.Message,
		ODataId:    s.ODataId,
		ODataType:  s.ODataType,
		Target:     s.Target,
	}
}

var _ json.Marshaler = ProcessingError{}

func (s ProcessingError) MarshalJSON() ([]byte, error) {
	type wrapper ProcessingError
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ProcessingError: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ProcessingError: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.processingError"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ProcessingError: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ProcessingError{}

func (s *ProcessingError) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ErrorType  *ContentProcessingErrorType `json:"errorType,omitempty"`
		Code       nullable.Type[string]       `json:"code,omitempty"`
		InnerError *ClassificationInnerError   `json:"innerError,omitempty"`
		Message    nullable.Type[string]       `json:"message,omitempty"`
		ODataId    *string                     `json:"@odata.id,omitempty"`
		ODataType  *string                     `json:"@odata.type,omitempty"`
		Target     nullable.Type[string]       `json:"target,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ErrorType = decoded.ErrorType
	s.Code = decoded.Code
	s.InnerError = decoded.InnerError
	s.Message = decoded.Message
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Target = decoded.Target

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ProcessingError into map[string]json.RawMessage: %+v", err)
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
				return fmt.Errorf("unmarshaling index %d field 'Details' for 'ProcessingError': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Details = &output
	}

	return nil
}
