package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InformationProtection struct {
	Bitlocker *Bitlocker `json:"bitlocker,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	ThreatAssessmentRequests *[]ThreatAssessmentRequest `json:"threatAssessmentRequests,omitempty"`
}

var _ json.Unmarshaler = &InformationProtection{}

func (s *InformationProtection) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Bitlocker *Bitlocker `json:"bitlocker,omitempty"`
		ODataId   *string    `json:"@odata.id,omitempty"`
		ODataType *string    `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Bitlocker = decoded.Bitlocker
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling InformationProtection into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["threatAssessmentRequests"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ThreatAssessmentRequests into list []json.RawMessage: %+v", err)
		}

		output := make([]ThreatAssessmentRequest, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalThreatAssessmentRequestImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ThreatAssessmentRequests' for 'InformationProtection': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ThreatAssessmentRequests = &output
	}

	return nil
}
