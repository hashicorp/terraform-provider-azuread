package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyScopeBase interface {
	PolicyScopeBase() BasePolicyScopeBaseImpl
}

var _ PolicyScopeBase = BasePolicyScopeBaseImpl{}

type BasePolicyScopeBaseImpl struct {
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

func (s BasePolicyScopeBaseImpl) PolicyScopeBase() BasePolicyScopeBaseImpl {
	return s
}

var _ PolicyScopeBase = RawPolicyScopeBaseImpl{}

// RawPolicyScopeBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawPolicyScopeBaseImpl struct {
	policyScopeBase BasePolicyScopeBaseImpl
	Type            string
	Values          map[string]interface{}
}

func (s RawPolicyScopeBaseImpl) PolicyScopeBase() BasePolicyScopeBaseImpl {
	return s.policyScopeBase
}

var _ json.Unmarshaler = &BasePolicyScopeBaseImpl{}

func (s *BasePolicyScopeBaseImpl) UnmarshalJSON(bytes []byte) error {
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
		return fmt.Errorf("unmarshaling BasePolicyScopeBaseImpl into map[string]json.RawMessage: %+v", err)
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
				return fmt.Errorf("unmarshaling index %d field 'Locations' for 'BasePolicyScopeBaseImpl': %+v", i, err)
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
				return fmt.Errorf("unmarshaling index %d field 'PolicyActions' for 'BasePolicyScopeBaseImpl': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.PolicyActions = output
	}

	return nil
}

func UnmarshalPolicyScopeBaseImplementation(input []byte) (PolicyScopeBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling PolicyScopeBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.policyTenantScope") {
		var out PolicyTenantScope
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PolicyTenantScope: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.policyUserScope") {
		var out PolicyUserScope
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PolicyUserScope: %+v", err)
		}
		return out, nil
	}

	var parent BasePolicyScopeBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BasePolicyScopeBaseImpl: %+v", err)
	}

	return RawPolicyScopeBaseImpl{
		policyScopeBase: parent,
		Type:            value,
		Values:          temp,
	}, nil

}
