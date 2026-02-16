package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomSecurityAttributeExemption interface {
	Entity
	CustomSecurityAttributeExemption() BaseCustomSecurityAttributeExemptionImpl
}

var _ CustomSecurityAttributeExemption = BaseCustomSecurityAttributeExemptionImpl{}

type BaseCustomSecurityAttributeExemptionImpl struct {
	// The possible values are: equals, unknownFutureValue. If equals, the customSecurityAttributeExemption value is
	// compared to match the custom security attribute value for the exemption to be applied. The comparison is case
	// sensitive.
	Operator *CustomSecurityAttributeComparisonOperator `json:"operator,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseCustomSecurityAttributeExemptionImpl) CustomSecurityAttributeExemption() BaseCustomSecurityAttributeExemptionImpl {
	return s
}

func (s BaseCustomSecurityAttributeExemptionImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ CustomSecurityAttributeExemption = RawCustomSecurityAttributeExemptionImpl{}

// RawCustomSecurityAttributeExemptionImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawCustomSecurityAttributeExemptionImpl struct {
	customSecurityAttributeExemption BaseCustomSecurityAttributeExemptionImpl
	Type                             string
	Values                           map[string]interface{}
}

func (s RawCustomSecurityAttributeExemptionImpl) CustomSecurityAttributeExemption() BaseCustomSecurityAttributeExemptionImpl {
	return s.customSecurityAttributeExemption
}

func (s RawCustomSecurityAttributeExemptionImpl) Entity() BaseEntityImpl {
	return s.customSecurityAttributeExemption.Entity()
}

var _ json.Marshaler = BaseCustomSecurityAttributeExemptionImpl{}

func (s BaseCustomSecurityAttributeExemptionImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseCustomSecurityAttributeExemptionImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseCustomSecurityAttributeExemptionImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseCustomSecurityAttributeExemptionImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.customSecurityAttributeExemption"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseCustomSecurityAttributeExemptionImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalCustomSecurityAttributeExemptionImplementation(input []byte) (CustomSecurityAttributeExemption, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling CustomSecurityAttributeExemption into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.customSecurityAttributeStringValueExemption") {
		var out CustomSecurityAttributeStringValueExemption
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CustomSecurityAttributeStringValueExemption: %+v", err)
		}
		return out, nil
	}

	var parent BaseCustomSecurityAttributeExemptionImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseCustomSecurityAttributeExemptionImpl: %+v", err)
	}

	return RawCustomSecurityAttributeExemptionImpl{
		customSecurityAttributeExemption: parent,
		Type:                             value,
		Values:                           temp,
	}, nil

}
