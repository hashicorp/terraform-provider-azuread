package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ KeyTypedValuePair = KeyStringValuePair{}

type KeyStringValuePair struct {
	// The string value of the key-value pair.
	Value *string `json:"value,omitempty"`

	// Fields inherited from KeyTypedValuePair

	// The string key of the key-value pair.
	Key *string `json:"key,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s KeyStringValuePair) KeyTypedValuePair() BaseKeyTypedValuePairImpl {
	return BaseKeyTypedValuePairImpl{
		Key:       s.Key,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = KeyStringValuePair{}

func (s KeyStringValuePair) MarshalJSON() ([]byte, error) {
	type wrapper KeyStringValuePair
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling KeyStringValuePair: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling KeyStringValuePair: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.keyStringValuePair"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling KeyStringValuePair: %+v", err)
	}

	return encoded, nil
}
