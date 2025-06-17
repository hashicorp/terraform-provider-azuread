package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HealthMonitoringEnrichment struct {
	// A collection of resource impact summaries that gives a high level view of the kind of resources that were impacted
	// and to what degree.
	Impacts *[]HealthMonitoringResourceImpactSummary `json:"impacts,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	State *HealthMonitoringEnrichmentState `json:"state,omitempty"`

	// A collection of supportingData locations that can be queried for debugging the alert.
	SupportingData *HealthMonitoringSupportingData `json:"supportingData,omitempty"`
}

var _ json.Unmarshaler = &HealthMonitoringEnrichment{}

func (s *HealthMonitoringEnrichment) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ODataId        *string                          `json:"@odata.id,omitempty"`
		ODataType      *string                          `json:"@odata.type,omitempty"`
		State          *HealthMonitoringEnrichmentState `json:"state,omitempty"`
		SupportingData *HealthMonitoringSupportingData  `json:"supportingData,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.State = decoded.State
	s.SupportingData = decoded.SupportingData

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling HealthMonitoringEnrichment into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["impacts"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Impacts into list []json.RawMessage: %+v", err)
		}

		output := make([]HealthMonitoringResourceImpactSummary, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalHealthMonitoringResourceImpactSummaryImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Impacts' for 'HealthMonitoringEnrichment': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Impacts = &output
	}

	return nil
}
