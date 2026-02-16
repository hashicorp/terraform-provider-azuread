package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataIndustryDataActivityStatistics interface {
	IndustryDataIndustryDataActivityStatistics() BaseIndustryDataIndustryDataActivityStatisticsImpl
}

var _ IndustryDataIndustryDataActivityStatistics = BaseIndustryDataIndustryDataActivityStatisticsImpl{}

type BaseIndustryDataIndustryDataActivityStatisticsImpl struct {
	// The identifier for the activity that is being reported on.
	ActivityId *string `json:"activityId,omitempty"`

	// The display name of the underlying flow.
	DisplayName *string `json:"displayName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Status *IndustryDataIndustryDataActivityStatus `json:"status,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseIndustryDataIndustryDataActivityStatisticsImpl) IndustryDataIndustryDataActivityStatistics() BaseIndustryDataIndustryDataActivityStatisticsImpl {
	return s
}

var _ IndustryDataIndustryDataActivityStatistics = RawIndustryDataIndustryDataActivityStatisticsImpl{}

// RawIndustryDataIndustryDataActivityStatisticsImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawIndustryDataIndustryDataActivityStatisticsImpl struct {
	industryDataIndustryDataActivityStatistics BaseIndustryDataIndustryDataActivityStatisticsImpl
	Type                                       string
	Values                                     map[string]interface{}
}

func (s RawIndustryDataIndustryDataActivityStatisticsImpl) IndustryDataIndustryDataActivityStatistics() BaseIndustryDataIndustryDataActivityStatisticsImpl {
	return s.industryDataIndustryDataActivityStatistics
}

var _ json.Marshaler = BaseIndustryDataIndustryDataActivityStatisticsImpl{}

func (s BaseIndustryDataIndustryDataActivityStatisticsImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseIndustryDataIndustryDataActivityStatisticsImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseIndustryDataIndustryDataActivityStatisticsImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseIndustryDataIndustryDataActivityStatisticsImpl: %+v", err)
	}

	delete(decoded, "activityId")
	delete(decoded, "displayName")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseIndustryDataIndustryDataActivityStatisticsImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalIndustryDataIndustryDataActivityStatisticsImplementation(input []byte) (IndustryDataIndustryDataActivityStatistics, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling IndustryDataIndustryDataActivityStatistics into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.inboundActivityResults") {
		var out IndustryDataInboundActivityResults
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataInboundActivityResults: %+v", err)
		}
		return out, nil
	}

	var parent BaseIndustryDataIndustryDataActivityStatisticsImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseIndustryDataIndustryDataActivityStatisticsImpl: %+v", err)
	}

	return RawIndustryDataIndustryDataActivityStatisticsImpl{
		industryDataIndustryDataActivityStatistics: parent,
		Type:   value,
		Values: temp,
	}, nil

}
