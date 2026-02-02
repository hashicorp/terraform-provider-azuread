package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CustomSecurityAttributeExemption = CustomSecurityAttributeStringValueExemption{}

type CustomSecurityAttributeStringValueExemption struct {
	// Value representing custom security attribute value to compare against while evaluating the exemption.
	Value nullable.Type[string] `json:"value,omitempty"`

	// Fields inherited from CustomSecurityAttributeExemption

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

func (s CustomSecurityAttributeStringValueExemption) CustomSecurityAttributeExemption() BaseCustomSecurityAttributeExemptionImpl {
	return BaseCustomSecurityAttributeExemptionImpl{
		Operator:  s.Operator,
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s CustomSecurityAttributeStringValueExemption) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CustomSecurityAttributeStringValueExemption{}

func (s CustomSecurityAttributeStringValueExemption) MarshalJSON() ([]byte, error) {
	type wrapper CustomSecurityAttributeStringValueExemption
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CustomSecurityAttributeStringValueExemption: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CustomSecurityAttributeStringValueExemption: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.customSecurityAttributeStringValueExemption"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CustomSecurityAttributeStringValueExemption: %+v", err)
	}

	return encoded, nil
}
