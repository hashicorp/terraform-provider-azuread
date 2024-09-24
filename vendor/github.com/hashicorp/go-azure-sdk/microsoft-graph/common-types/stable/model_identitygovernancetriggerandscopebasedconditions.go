package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IdentityGovernanceWorkflowExecutionConditions = IdentityGovernanceTriggerAndScopeBasedConditions{}

type IdentityGovernanceTriggerAndScopeBasedConditions struct {
	// Defines who the workflow runs for.
	Scope SubjectSet `json:"scope"`

	// What triggers a workflow to run.
	Trigger IdentityGovernanceWorkflowExecutionTrigger `json:"trigger"`

	// Fields inherited from IdentityGovernanceWorkflowExecutionConditions

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s IdentityGovernanceTriggerAndScopeBasedConditions) IdentityGovernanceWorkflowExecutionConditions() BaseIdentityGovernanceWorkflowExecutionConditionsImpl {
	return BaseIdentityGovernanceWorkflowExecutionConditionsImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IdentityGovernanceTriggerAndScopeBasedConditions{}

func (s IdentityGovernanceTriggerAndScopeBasedConditions) MarshalJSON() ([]byte, error) {
	type wrapper IdentityGovernanceTriggerAndScopeBasedConditions
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IdentityGovernanceTriggerAndScopeBasedConditions: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IdentityGovernanceTriggerAndScopeBasedConditions: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.identityGovernance.triggerAndScopeBasedConditions"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IdentityGovernanceTriggerAndScopeBasedConditions: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &IdentityGovernanceTriggerAndScopeBasedConditions{}

func (s *IdentityGovernanceTriggerAndScopeBasedConditions) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ODataId   *string `json:"@odata.id,omitempty"`
		ODataType *string `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling IdentityGovernanceTriggerAndScopeBasedConditions into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["scope"]; ok {
		impl, err := UnmarshalSubjectSetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Scope' for 'IdentityGovernanceTriggerAndScopeBasedConditions': %+v", err)
		}
		s.Scope = impl
	}

	if v, ok := temp["trigger"]; ok {
		impl, err := UnmarshalIdentityGovernanceWorkflowExecutionTriggerImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Trigger' for 'IdentityGovernanceTriggerAndScopeBasedConditions': %+v", err)
		}
		s.Trigger = impl
	}

	return nil
}
