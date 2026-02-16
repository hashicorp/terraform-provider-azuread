package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataIndustryDataRunStatistics struct {
	// The collection of statistics for each activity included in this run.
	ActivityStatistics *[]IndustryDataIndustryDataActivityStatistics `json:"activityStatistics,omitempty"`

	// The aggregate statistics for all inbound flows.
	InboundTotals *IndustryDataAggregatedInboundStatistics `json:"inboundTotals,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The ID of the underlying run for the statistics.
	RunId *string `json:"runId,omitempty"`

	Status *IndustryDataIndustryDataRunStatus `json:"status,omitempty"`
}

var _ json.Marshaler = IndustryDataIndustryDataRunStatistics{}

func (s IndustryDataIndustryDataRunStatistics) MarshalJSON() ([]byte, error) {
	type wrapper IndustryDataIndustryDataRunStatistics
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IndustryDataIndustryDataRunStatistics: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IndustryDataIndustryDataRunStatistics: %+v", err)
	}

	delete(decoded, "activityStatistics")
	delete(decoded, "inboundTotals")
	delete(decoded, "runId")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IndustryDataIndustryDataRunStatistics: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &IndustryDataIndustryDataRunStatistics{}

func (s *IndustryDataIndustryDataRunStatistics) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		InboundTotals *IndustryDataAggregatedInboundStatistics `json:"inboundTotals,omitempty"`
		ODataId       *string                                  `json:"@odata.id,omitempty"`
		ODataType     *string                                  `json:"@odata.type,omitempty"`
		RunId         *string                                  `json:"runId,omitempty"`
		Status        *IndustryDataIndustryDataRunStatus       `json:"status,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.InboundTotals = decoded.InboundTotals
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.RunId = decoded.RunId
	s.Status = decoded.Status

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling IndustryDataIndustryDataRunStatistics into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["activityStatistics"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ActivityStatistics into list []json.RawMessage: %+v", err)
		}

		output := make([]IndustryDataIndustryDataActivityStatistics, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalIndustryDataIndustryDataActivityStatisticsImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ActivityStatistics' for 'IndustryDataIndustryDataRunStatistics': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ActivityStatistics = &output
	}

	return nil
}
