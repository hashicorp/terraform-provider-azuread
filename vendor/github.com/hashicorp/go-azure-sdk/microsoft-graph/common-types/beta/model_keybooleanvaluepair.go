package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ KeyTypedValuePair = KeyBooleanValuePair{}

type KeyBooleanValuePair struct {
	// The Boolean value of the key-value pair.
	Value *bool `json:"value,omitempty"`

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

func (s KeyBooleanValuePair) KeyTypedValuePair() BaseKeyTypedValuePairImpl {
	return BaseKeyTypedValuePairImpl{
		Key:       s.Key,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = KeyBooleanValuePair{}

func (s KeyBooleanValuePair) MarshalJSON() ([]byte, error) {
	type wrapper KeyBooleanValuePair
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling KeyBooleanValuePair: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling KeyBooleanValuePair: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.keyBooleanValuePair"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling KeyBooleanValuePair: %+v", err)
	}

	return encoded, nil
}
