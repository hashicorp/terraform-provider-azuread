package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GenericError interface {
	GenericError() BaseGenericErrorImpl
}

var _ GenericError = BaseGenericErrorImpl{}

type BaseGenericErrorImpl struct {
	// The error code.
	Code nullable.Type[string] `json:"code,omitempty"`

	// The error message.
	Message nullable.Type[string] `json:"message,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseGenericErrorImpl) GenericError() BaseGenericErrorImpl {
	return s
}

var _ GenericError = RawGenericErrorImpl{}

// RawGenericErrorImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawGenericErrorImpl struct {
	genericError BaseGenericErrorImpl
	Type         string
	Values       map[string]interface{}
}

func (s RawGenericErrorImpl) GenericError() BaseGenericErrorImpl {
	return s.genericError
}

func UnmarshalGenericErrorImplementation(input []byte) (GenericError, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling GenericError into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.accessReviewError") {
		var out AccessReviewError
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessReviewError: %+v", err)
		}
		return out, nil
	}

	var parent BaseGenericErrorImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseGenericErrorImpl: %+v", err)
	}

	return RawGenericErrorImpl{
		genericError: parent,
		Type:         value,
		Values:       temp,
	}, nil

}
