package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = InferenceClassificationOverride{}

type InferenceClassificationOverride struct {
	// Specifies how incoming messages from a specific sender should always be classified as. Possible values are: focused,
	// other.
	ClassifyAs *InferenceClassificationType `json:"classifyAs,omitempty"`

	// The email address information of the sender for whom the override is created.
	SenderEmailAddress EmailAddress `json:"senderEmailAddress"`

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

var _ json.Unmarshaler = &InferenceClassificationOverride{}

func (s *InferenceClassificationOverride) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ClassifyAs *InferenceClassificationType `json:"classifyAs,omitempty"`
		Id         *string                      `json:"id,omitempty"`
		ODataId    *string                      `json:"@odata.id,omitempty"`
		ODataType  *string                      `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ClassifyAs = decoded.ClassifyAs
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling InferenceClassificationOverride into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["senderEmailAddress"]; ok {
		impl, err := UnmarshalEmailAddressImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'SenderEmailAddress' for 'InferenceClassificationOverride': %+v", err)
		}
		s.SenderEmailAddress = impl
	}

	return nil
}
