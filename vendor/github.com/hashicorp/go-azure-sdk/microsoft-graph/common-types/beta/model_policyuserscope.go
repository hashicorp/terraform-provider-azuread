package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PolicyScopeBase = PolicyUserScope{}

type PolicyUserScope struct {

	// Fields inherited from PolicyScopeBase

	Activities    *UserActivityTypes `json:"activities,omitempty"`
	ExecutionMode *ExecutionMode     `json:"executionMode,omitempty"`

	// The locations (like domains or URLs) to be protected. Required.
	Locations []PolicyLocation `json:"locations"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The enforcement actions to take if the policy conditions are met within this scope. Required.
	PolicyActions []DlpActionInfo `json:"policyActions"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s PolicyUserScope) PolicyScopeBase() BasePolicyScopeBaseImpl {
	return BasePolicyScopeBaseImpl{
		Activities:    s.Activities,
		ExecutionMode: s.ExecutionMode,
		Locations:     s.Locations,
		ODataId:       s.ODataId,
		ODataType:     s.ODataType,
		PolicyActions: s.PolicyActions,
	}
}

var _ json.Marshaler = PolicyUserScope{}

func (s PolicyUserScope) MarshalJSON() ([]byte, error) {
	type wrapper PolicyUserScope
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PolicyUserScope: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PolicyUserScope: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.policyUserScope"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PolicyUserScope: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &PolicyUserScope{}

func (s *PolicyUserScope) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Activities    *UserActivityTypes `json:"activities,omitempty"`
		ExecutionMode *ExecutionMode     `json:"executionMode,omitempty"`
		ODataId       *string            `json:"@odata.id,omitempty"`
		ODataType     *string            `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Activities = decoded.Activities
	s.ExecutionMode = decoded.ExecutionMode
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling PolicyUserScope into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["locations"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Locations into list []json.RawMessage: %+v", err)
		}

		output := make([]PolicyLocation, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalPolicyLocationImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Locations' for 'PolicyUserScope': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Locations = output
	}

	if v, ok := temp["policyActions"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling PolicyActions into list []json.RawMessage: %+v", err)
		}

		output := make([]DlpActionInfo, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDlpActionInfoImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'PolicyActions' for 'PolicyUserScope': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.PolicyActions = output
	}

	return nil
}
