package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DetectedSensitiveContentBase = ExactMatchDetectedSensitiveContent{}

type ExactMatchDetectedSensitiveContent struct {
	Matches *[]SensitiveContentLocation `json:"matches,omitempty"`

	// Fields inherited from DetectedSensitiveContentBase

	Confidence  nullable.Type[int64]  `json:"confidence,omitempty"`
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`
	Id          nullable.Type[string] `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	RecommendedConfidence nullable.Type[int64] `json:"recommendedConfidence,omitempty"`
	UniqueCount           nullable.Type[int64] `json:"uniqueCount,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s ExactMatchDetectedSensitiveContent) DetectedSensitiveContentBase() BaseDetectedSensitiveContentBaseImpl {
	return BaseDetectedSensitiveContentBaseImpl{
		Confidence:            s.Confidence,
		DisplayName:           s.DisplayName,
		Id:                    s.Id,
		ODataId:               s.ODataId,
		ODataType:             s.ODataType,
		RecommendedConfidence: s.RecommendedConfidence,
		UniqueCount:           s.UniqueCount,
	}
}

var _ json.Marshaler = ExactMatchDetectedSensitiveContent{}

func (s ExactMatchDetectedSensitiveContent) MarshalJSON() ([]byte, error) {
	type wrapper ExactMatchDetectedSensitiveContent
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ExactMatchDetectedSensitiveContent: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ExactMatchDetectedSensitiveContent: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.exactMatchDetectedSensitiveContent"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ExactMatchDetectedSensitiveContent: %+v", err)
	}

	return encoded, nil
}
