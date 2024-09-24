package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = InferenceClassificationOverride{}

type InferenceClassificationOverride struct {
	// Specifies how incoming messages from a specific sender should always be classified as. The possible values are:
	// focused, other.
	ClassifyAs *InferenceClassificationType `json:"classifyAs,omitempty"`

	// The email address information of the sender for whom the override is created.
	SenderEmailAddress *EmailAddress `json:"senderEmailAddress,omitempty"`

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

func (s InferenceClassificationOverride) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = InferenceClassificationOverride{}

func (s InferenceClassificationOverride) MarshalJSON() ([]byte, error) {
	type wrapper InferenceClassificationOverride
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling InferenceClassificationOverride: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling InferenceClassificationOverride: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.inferenceClassificationOverride"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling InferenceClassificationOverride: %+v", err)
	}

	return encoded, nil
}
