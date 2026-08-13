package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceWorkflowExecutionConditions interface {
	IdentityGovernanceWorkflowExecutionConditions() BaseIdentityGovernanceWorkflowExecutionConditionsImpl
}

var _ IdentityGovernanceWorkflowExecutionConditions = BaseIdentityGovernanceWorkflowExecutionConditionsImpl{}

type BaseIdentityGovernanceWorkflowExecutionConditionsImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseIdentityGovernanceWorkflowExecutionConditionsImpl) IdentityGovernanceWorkflowExecutionConditions() BaseIdentityGovernanceWorkflowExecutionConditionsImpl {
	return s
}

var _ IdentityGovernanceWorkflowExecutionConditions = RawIdentityGovernanceWorkflowExecutionConditionsImpl{}

// RawIdentityGovernanceWorkflowExecutionConditionsImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawIdentityGovernanceWorkflowExecutionConditionsImpl struct {
	identityGovernanceWorkflowExecutionConditions BaseIdentityGovernanceWorkflowExecutionConditionsImpl
	Type                                          string
	Values                                        map[string]interface{}
}

func (s RawIdentityGovernanceWorkflowExecutionConditionsImpl) IdentityGovernanceWorkflowExecutionConditions() BaseIdentityGovernanceWorkflowExecutionConditionsImpl {
	return s.identityGovernanceWorkflowExecutionConditions
}

func (s RawIdentityGovernanceWorkflowExecutionConditionsImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalIdentityGovernanceWorkflowExecutionConditionsImplementation(input []byte) (IdentityGovernanceWorkflowExecutionConditions, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling IdentityGovernanceWorkflowExecutionConditions into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.identityGovernance.onDemandExecutionOnly") {
		var out IdentityGovernanceOnDemandExecutionOnly
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityGovernanceOnDemandExecutionOnly: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identityGovernance.triggerAndScopeBasedConditions") {
		var out IdentityGovernanceTriggerAndScopeBasedConditions
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityGovernanceTriggerAndScopeBasedConditions: %+v", err)
		}
		return out, nil
	}

	var parent BaseIdentityGovernanceWorkflowExecutionConditionsImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseIdentityGovernanceWorkflowExecutionConditionsImpl: %+v", err)
	}

	return RawIdentityGovernanceWorkflowExecutionConditionsImpl{
		identityGovernanceWorkflowExecutionConditions: parent,
		Type:   value,
		Values: temp,
	}, nil

}
