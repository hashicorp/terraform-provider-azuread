package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomClaimAttributeBase interface {
	CustomClaimAttributeBase() BaseCustomClaimAttributeBaseImpl
}

var _ CustomClaimAttributeBase = BaseCustomClaimAttributeBaseImpl{}

type BaseCustomClaimAttributeBaseImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseCustomClaimAttributeBaseImpl) CustomClaimAttributeBase() BaseCustomClaimAttributeBaseImpl {
	return s
}

var _ CustomClaimAttributeBase = RawCustomClaimAttributeBaseImpl{}

// RawCustomClaimAttributeBaseImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawCustomClaimAttributeBaseImpl struct {
	customClaimAttributeBase BaseCustomClaimAttributeBaseImpl
	Type                     string
	Values                   map[string]interface{}
}

func (s RawCustomClaimAttributeBaseImpl) CustomClaimAttributeBase() BaseCustomClaimAttributeBaseImpl {
	return s.customClaimAttributeBase
}

func (s RawCustomClaimAttributeBaseImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalCustomClaimAttributeBaseImplementation(input []byte) (CustomClaimAttributeBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling CustomClaimAttributeBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.sourcedAttribute") {
		var out SourcedAttribute
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SourcedAttribute: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.valueBasedAttribute") {
		var out ValueBasedAttribute
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ValueBasedAttribute: %+v", err)
		}
		return out, nil
	}

	var parent BaseCustomClaimAttributeBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseCustomClaimAttributeBaseImpl: %+v", err)
	}

	return RawCustomClaimAttributeBaseImpl{
		customClaimAttributeBase: parent,
		Type:                     value,
		Values:                   temp,
	}, nil

}
