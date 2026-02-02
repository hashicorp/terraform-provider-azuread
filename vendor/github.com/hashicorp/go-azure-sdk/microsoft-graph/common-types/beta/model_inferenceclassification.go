package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = InferenceClassification{}

type InferenceClassification struct {
	// A set of overrides for a user to always classify messages from specific senders in certain ways: focused, or other.
	// Read-only. Nullable.
	Overrides *[]InferenceClassificationOverride `json:"overrides,omitempty"`

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

func (s InferenceClassification) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = InferenceClassification{}

func (s InferenceClassification) MarshalJSON() ([]byte, error) {
	type wrapper InferenceClassification
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling InferenceClassification: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling InferenceClassification: %+v", err)
	}

	delete(decoded, "overrides")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.inferenceClassification"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling InferenceClassification: %+v", err)
	}

	return encoded, nil
}
