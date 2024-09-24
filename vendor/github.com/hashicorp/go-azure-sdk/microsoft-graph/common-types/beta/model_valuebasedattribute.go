package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CustomClaimAttributeBase = ValueBasedAttribute{}

type ValueBasedAttribute struct {
	// The static value to be used an the attribute.
	Value *string `json:"value,omitempty"`

	// Fields inherited from CustomClaimAttributeBase

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s ValueBasedAttribute) CustomClaimAttributeBase() BaseCustomClaimAttributeBaseImpl {
	return BaseCustomClaimAttributeBaseImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ValueBasedAttribute{}

func (s ValueBasedAttribute) MarshalJSON() ([]byte, error) {
	type wrapper ValueBasedAttribute
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ValueBasedAttribute: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ValueBasedAttribute: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.valueBasedAttribute"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ValueBasedAttribute: %+v", err)
	}

	return encoded, nil
}
