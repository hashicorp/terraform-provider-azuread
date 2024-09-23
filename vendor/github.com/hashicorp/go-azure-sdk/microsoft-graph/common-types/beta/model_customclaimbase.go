package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomClaimBase interface {
	CustomClaimBase() BaseCustomClaimBaseImpl
}

var _ CustomClaimBase = BaseCustomClaimBaseImpl{}

type BaseCustomClaimBaseImpl struct {
	// One or more configurations that describe how the claim is sourced and under what conditions.
	Configurations *[]CustomClaimConfiguration `json:"configurations,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseCustomClaimBaseImpl) CustomClaimBase() BaseCustomClaimBaseImpl {
	return s
}

var _ CustomClaimBase = RawCustomClaimBaseImpl{}

// RawCustomClaimBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawCustomClaimBaseImpl struct {
	customClaimBase BaseCustomClaimBaseImpl
	Type            string
	Values          map[string]interface{}
}

func (s RawCustomClaimBaseImpl) CustomClaimBase() BaseCustomClaimBaseImpl {
	return s.customClaimBase
}

func UnmarshalCustomClaimBaseImplementation(input []byte) (CustomClaimBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling CustomClaimBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.customClaim") {
		var out CustomClaim
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CustomClaim: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.samlNameIdClaim") {
		var out SamlNameIdClaim
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SamlNameIdClaim: %+v", err)
		}
		return out, nil
	}

	var parent BaseCustomClaimBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseCustomClaimBaseImpl: %+v", err)
	}

	return RawCustomClaimBaseImpl{
		customClaimBase: parent,
		Type:            value,
		Values:          temp,
	}, nil

}
