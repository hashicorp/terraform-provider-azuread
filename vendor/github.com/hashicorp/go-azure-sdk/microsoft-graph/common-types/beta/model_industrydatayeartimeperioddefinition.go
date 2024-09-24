package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = IndustryDataYearTimePeriodDefinition{}

type IndustryDataYearTimePeriodDefinition struct {
	// The name of the year. Maximum supported length is 100 characters.
	DisplayName *string `json:"displayName,omitempty"`

	// The last day of the year using ISO 8601 format for date.
	EndDate *string `json:"endDate,omitempty"`

	// The first day of the year using ISO 8601 format for date.
	StartDate *string `json:"startDate,omitempty"`

	Year *IndustryDataYearReferenceValue `json:"year,omitempty"`

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

func (s IndustryDataYearTimePeriodDefinition) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IndustryDataYearTimePeriodDefinition{}

func (s IndustryDataYearTimePeriodDefinition) MarshalJSON() ([]byte, error) {
	type wrapper IndustryDataYearTimePeriodDefinition
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IndustryDataYearTimePeriodDefinition: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IndustryDataYearTimePeriodDefinition: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.industryData.yearTimePeriodDefinition"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IndustryDataYearTimePeriodDefinition: %+v", err)
	}

	return encoded, nil
}
