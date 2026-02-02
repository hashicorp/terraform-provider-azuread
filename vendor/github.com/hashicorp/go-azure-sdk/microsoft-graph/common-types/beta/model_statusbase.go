package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StatusBase interface {
	StatusBase() BaseStatusBaseImpl
}

var _ StatusBase = BaseStatusBaseImpl{}

type BaseStatusBaseImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Possible values are: success, warning, failure, skipped, unknownFutureValue. Supports $filter (eq, contains).
	Status *ProvisioningResult `json:"status,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseStatusBaseImpl) StatusBase() BaseStatusBaseImpl {
	return s
}

var _ StatusBase = RawStatusBaseImpl{}

// RawStatusBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawStatusBaseImpl struct {
	statusBase BaseStatusBaseImpl
	Type       string
	Values     map[string]interface{}
}

func (s RawStatusBaseImpl) StatusBase() BaseStatusBaseImpl {
	return s.statusBase
}

func UnmarshalStatusBaseImplementation(input []byte) (StatusBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling StatusBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.statusDetails") {
		var out StatusDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into StatusDetails: %+v", err)
		}
		return out, nil
	}

	var parent BaseStatusBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseStatusBaseImpl: %+v", err)
	}

	return RawStatusBaseImpl{
		statusBase: parent,
		Type:       value,
		Values:     temp,
	}, nil

}
