package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ClassifcationErrorBase = ClassificationError{}

type ClassificationError struct {
	Details *[]ClassifcationErrorBase `json:"details,omitempty"`

	// Fields inherited from ClassifcationErrorBase

	Code       nullable.Type[string]     `json:"code,omitempty"`
	InnerError *ClassificationInnerError `json:"innerError,omitempty"`
	Message    nullable.Type[string]     `json:"message,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Target nullable.Type[string] `json:"target,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s ClassificationError) ClassifcationErrorBase() BaseClassifcationErrorBaseImpl {
	return BaseClassifcationErrorBaseImpl{
		Code:       s.Code,
		InnerError: s.InnerError,
		Message:    s.Message,
		ODataId:    s.ODataId,
		ODataType:  s.ODataType,
		Target:     s.Target,
	}
}

var _ json.Marshaler = ClassificationError{}

func (s ClassificationError) MarshalJSON() ([]byte, error) {
	type wrapper ClassificationError
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ClassificationError: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ClassificationError: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.classificationError"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ClassificationError: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ClassificationError{}

func (s *ClassificationError) UnmarshalJSON(bytes []byte) error {
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
		return fmt.Errorf("unmarshaling ClassificationError into map[string]json.RawMessage: %+v", err)
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
				return fmt.Errorf("unmarshaling index %d field 'Details' for 'ClassificationError': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Details = &output
	}

	return nil
}
