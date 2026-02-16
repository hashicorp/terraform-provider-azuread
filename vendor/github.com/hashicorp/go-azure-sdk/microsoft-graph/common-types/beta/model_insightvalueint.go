package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ UserExperienceAnalyticsInsightValue = InsightValueInt{}

type InsightValueInt struct {
	// The int value of the user experience analytics insight.
	Value *int64 `json:"value,omitempty"`

	// Fields inherited from UserExperienceAnalyticsInsightValue

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s InsightValueInt) UserExperienceAnalyticsInsightValue() BaseUserExperienceAnalyticsInsightValueImpl {
	return BaseUserExperienceAnalyticsInsightValueImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = InsightValueInt{}

func (s InsightValueInt) MarshalJSON() ([]byte, error) {
	type wrapper InsightValueInt
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling InsightValueInt: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling InsightValueInt: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.insightValueInt"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling InsightValueInt: %+v", err)
	}

	return encoded, nil
}
