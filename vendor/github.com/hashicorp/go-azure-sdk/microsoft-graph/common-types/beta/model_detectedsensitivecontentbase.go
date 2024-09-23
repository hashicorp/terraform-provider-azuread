package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DetectedSensitiveContentBase interface {
	DetectedSensitiveContentBase() BaseDetectedSensitiveContentBaseImpl
}

var _ DetectedSensitiveContentBase = BaseDetectedSensitiveContentBaseImpl{}

type BaseDetectedSensitiveContentBaseImpl struct {
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

func (s BaseDetectedSensitiveContentBaseImpl) DetectedSensitiveContentBase() BaseDetectedSensitiveContentBaseImpl {
	return s
}

var _ DetectedSensitiveContentBase = RawDetectedSensitiveContentBaseImpl{}

// RawDetectedSensitiveContentBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDetectedSensitiveContentBaseImpl struct {
	detectedSensitiveContentBase BaseDetectedSensitiveContentBaseImpl
	Type                         string
	Values                       map[string]interface{}
}

func (s RawDetectedSensitiveContentBaseImpl) DetectedSensitiveContentBase() BaseDetectedSensitiveContentBaseImpl {
	return s.detectedSensitiveContentBase
}

func UnmarshalDetectedSensitiveContentBaseImplementation(input []byte) (DetectedSensitiveContentBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DetectedSensitiveContentBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.detectedSensitiveContent") {
		var out DetectedSensitiveContent
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DetectedSensitiveContent: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.exactMatchDetectedSensitiveContent") {
		var out ExactMatchDetectedSensitiveContent
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExactMatchDetectedSensitiveContent: %+v", err)
		}
		return out, nil
	}

	var parent BaseDetectedSensitiveContentBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDetectedSensitiveContentBaseImpl: %+v", err)
	}

	return RawDetectedSensitiveContentBaseImpl{
		detectedSensitiveContentBase: parent,
		Type:                         value,
		Values:                       temp,
	}, nil

}
