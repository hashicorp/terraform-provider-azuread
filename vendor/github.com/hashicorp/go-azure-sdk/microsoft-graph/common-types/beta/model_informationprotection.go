package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = InformationProtection{}

type InformationProtection struct {
	Bitlocker                  *Bitlocker                   `json:"bitlocker,omitempty"`
	DataLossPreventionPolicies *[]DataLossPreventionPolicy  `json:"dataLossPreventionPolicies,omitempty"`
	Policy                     *InformationProtectionPolicy `json:"policy,omitempty"`
	SensitivityLabels          *[]SensitivityLabel          `json:"sensitivityLabels,omitempty"`
	SensitivityPolicySettings  *SensitivityPolicySettings   `json:"sensitivityPolicySettings,omitempty"`
	ThreatAssessmentRequests   *[]ThreatAssessmentRequest   `json:"threatAssessmentRequests,omitempty"`

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

func (s InformationProtection) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = InformationProtection{}

func (s InformationProtection) MarshalJSON() ([]byte, error) {
	type wrapper InformationProtection
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling InformationProtection: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling InformationProtection: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.informationProtection"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling InformationProtection: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &InformationProtection{}

func (s *InformationProtection) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Bitlocker                  *Bitlocker                   `json:"bitlocker,omitempty"`
		DataLossPreventionPolicies *[]DataLossPreventionPolicy  `json:"dataLossPreventionPolicies,omitempty"`
		Policy                     *InformationProtectionPolicy `json:"policy,omitempty"`
		SensitivityLabels          *[]SensitivityLabel          `json:"sensitivityLabels,omitempty"`
		SensitivityPolicySettings  *SensitivityPolicySettings   `json:"sensitivityPolicySettings,omitempty"`
		Id                         *string                      `json:"id,omitempty"`
		ODataId                    *string                      `json:"@odata.id,omitempty"`
		ODataType                  *string                      `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Bitlocker = decoded.Bitlocker
	s.DataLossPreventionPolicies = decoded.DataLossPreventionPolicies
	s.Policy = decoded.Policy
	s.SensitivityLabels = decoded.SensitivityLabels
	s.SensitivityPolicySettings = decoded.SensitivityPolicySettings
	s.Id = decoded.Id
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
