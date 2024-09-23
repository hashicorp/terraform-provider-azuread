package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = FilterOperatorSchema{}

type FilterOperatorSchema struct {
	Arity                     *ScopeOperatorType                      `json:"arity,omitempty"`
	MultivaluedComparisonType *ScopeOperatorMultiValuedComparisonType `json:"multivaluedComparisonType,omitempty"`

	// Attribute types supported by the operator. Possible values are: Boolean, Binary, Reference, Integer, String.
	SupportedAttributeTypes *[]AttributeType `json:"supportedAttributeTypes,omitempty"`

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

func (s FilterOperatorSchema) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = FilterOperatorSchema{}

func (s FilterOperatorSchema) MarshalJSON() ([]byte, error) {
	type wrapper FilterOperatorSchema
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling FilterOperatorSchema: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling FilterOperatorSchema: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.filterOperatorSchema"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling FilterOperatorSchema: %+v", err)
	}

	return encoded, nil
}
