package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UserInsightsRoot{}

type UserInsightsRoot struct {
	// Summaries of daily user activities on apps registered in your tenant that is configured for Microsoft Entra External
	// ID for customers.
	Daily *DailyUserInsightMetricsRoot `json:"daily,omitempty"`

	// Summaries of monthly user activities on apps registered in your tenant that is configured for Microsoft Entra
	// External ID for customers.
	Monthly *MonthlyUserInsightMetricsRoot `json:"monthly,omitempty"`

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

func (s UserInsightsRoot) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserInsightsRoot{}

func (s UserInsightsRoot) MarshalJSON() ([]byte, error) {
	type wrapper UserInsightsRoot
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserInsightsRoot: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserInsightsRoot: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userInsightsRoot"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserInsightsRoot: %+v", err)
	}

	return encoded, nil
}
