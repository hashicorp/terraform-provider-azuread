package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = EdiscoveryCaseSettings{}

type EdiscoveryCaseSettings struct {
	// The OCR (Optical Character Recognition) settings for the case.
	Ocr *EdiscoveryOcrSettings `json:"ocr,omitempty"`

	// The redundancy (near duplicate and email threading) detection settings for the case.
	RedundancyDetection *EdiscoveryRedundancyDetectionSettings `json:"redundancyDetection,omitempty"`

	// The article Modeling (Themes) settings for the case.
	TopicModeling *EdiscoveryTopicModelingSettings `json:"topicModeling,omitempty"`

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

func (s EdiscoveryCaseSettings) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EdiscoveryCaseSettings{}

func (s EdiscoveryCaseSettings) MarshalJSON() ([]byte, error) {
	type wrapper EdiscoveryCaseSettings
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EdiscoveryCaseSettings: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EdiscoveryCaseSettings: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.ediscovery.caseSettings"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EdiscoveryCaseSettings: %+v", err)
	}

	return encoded, nil
}
