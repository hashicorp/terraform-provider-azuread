package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DetectedSensitiveContent interface {
	DetectedSensitiveContentBase
	DetectedSensitiveContent() BaseDetectedSensitiveContentImpl
}

var _ DetectedSensitiveContent = BaseDetectedSensitiveContentImpl{}

type BaseDetectedSensitiveContentImpl struct {
	ClassificationAttributes *[]ClassificationAttribute  `json:"classificationAttributes,omitempty"`
	ClassificationMethod     *ClassificationMethod       `json:"classificationMethod,omitempty"`
	Matches                  *[]SensitiveContentLocation `json:"matches,omitempty"`
	Scope                    *SensitiveTypeScope         `json:"scope,omitempty"`
	SensitiveTypeSource      *SensitiveTypeSource        `json:"sensitiveTypeSource,omitempty"`

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

func (s BaseDetectedSensitiveContentImpl) DetectedSensitiveContent() BaseDetectedSensitiveContentImpl {
	return s
}

func (s BaseDetectedSensitiveContentImpl) DetectedSensitiveContentBase() BaseDetectedSensitiveContentBaseImpl {
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

var _ DetectedSensitiveContent = RawDetectedSensitiveContentImpl{}

// RawDetectedSensitiveContentImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDetectedSensitiveContentImpl struct {
	detectedSensitiveContent BaseDetectedSensitiveContentImpl
	Type                     string
	Values                   map[string]interface{}
}

func (s RawDetectedSensitiveContentImpl) DetectedSensitiveContent() BaseDetectedSensitiveContentImpl {
	return s.detectedSensitiveContent
}

func (s RawDetectedSensitiveContentImpl) DetectedSensitiveContentBase() BaseDetectedSensitiveContentBaseImpl {
	return s.detectedSensitiveContent.DetectedSensitiveContentBase()
}

var _ json.Marshaler = BaseDetectedSensitiveContentImpl{}

func (s BaseDetectedSensitiveContentImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseDetectedSensitiveContentImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseDetectedSensitiveContentImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseDetectedSensitiveContentImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.detectedSensitiveContent"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseDetectedSensitiveContentImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalDetectedSensitiveContentImplementation(input []byte) (DetectedSensitiveContent, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DetectedSensitiveContent into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.machineLearningDetectedSensitiveContent") {
		var out MachineLearningDetectedSensitiveContent
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MachineLearningDetectedSensitiveContent: %+v", err)
		}
		return out, nil
	}

	var parent BaseDetectedSensitiveContentImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDetectedSensitiveContentImpl: %+v", err)
	}

	return RawDetectedSensitiveContentImpl{
		detectedSensitiveContent: parent,
		Type:                     value,
		Values:                   temp,
	}, nil

}
