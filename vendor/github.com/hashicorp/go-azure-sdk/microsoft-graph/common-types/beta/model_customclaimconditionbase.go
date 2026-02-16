package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomClaimConditionBase interface {
	CustomClaimConditionBase() BaseCustomClaimConditionBaseImpl
}

var _ CustomClaimConditionBase = BaseCustomClaimConditionBaseImpl{}

type BaseCustomClaimConditionBaseImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseCustomClaimConditionBaseImpl) CustomClaimConditionBase() BaseCustomClaimConditionBaseImpl {
	return s
}

var _ CustomClaimConditionBase = RawCustomClaimConditionBaseImpl{}

// RawCustomClaimConditionBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawCustomClaimConditionBaseImpl struct {
	customClaimConditionBase BaseCustomClaimConditionBaseImpl
	Type                     string
	Values                   map[string]interface{}
}

func (s RawCustomClaimConditionBaseImpl) CustomClaimConditionBase() BaseCustomClaimConditionBaseImpl {
	return s.customClaimConditionBase
}

func UnmarshalCustomClaimConditionBaseImplementation(input []byte) (CustomClaimConditionBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling CustomClaimConditionBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.customClaimCondition") {
		var out CustomClaimCondition
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CustomClaimCondition: %+v", err)
		}
		return out, nil
	}

	var parent BaseCustomClaimConditionBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseCustomClaimConditionBaseImpl: %+v", err)
	}

	return RawCustomClaimConditionBaseImpl{
		customClaimConditionBase: parent,
		Type:                     value,
		Values:                   temp,
	}, nil

}
