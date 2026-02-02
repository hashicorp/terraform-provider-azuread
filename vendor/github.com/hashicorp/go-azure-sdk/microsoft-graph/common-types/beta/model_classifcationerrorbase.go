package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClassifcationErrorBase interface {
	ClassifcationErrorBase() BaseClassifcationErrorBaseImpl
}

var _ ClassifcationErrorBase = BaseClassifcationErrorBaseImpl{}

type BaseClassifcationErrorBaseImpl struct {
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

func (s BaseClassifcationErrorBaseImpl) ClassifcationErrorBase() BaseClassifcationErrorBaseImpl {
	return s
}

var _ ClassifcationErrorBase = RawClassifcationErrorBaseImpl{}

// RawClassifcationErrorBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawClassifcationErrorBaseImpl struct {
	classifcationErrorBase BaseClassifcationErrorBaseImpl
	Type                   string
	Values                 map[string]interface{}
}

func (s RawClassifcationErrorBaseImpl) ClassifcationErrorBase() BaseClassifcationErrorBaseImpl {
	return s.classifcationErrorBase
}

func UnmarshalClassifcationErrorBaseImplementation(input []byte) (ClassifcationErrorBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ClassifcationErrorBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.classificationError") {
		var out ClassificationError
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ClassificationError: %+v", err)
		}
		return out, nil
	}

	var parent BaseClassifcationErrorBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseClassifcationErrorBaseImpl: %+v", err)
	}

	return RawClassifcationErrorBaseImpl{
		classifcationErrorBase: parent,
		Type:                   value,
		Values:                 temp,
	}, nil

}
