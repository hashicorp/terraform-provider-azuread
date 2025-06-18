package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceWorkflowExecutionTrigger interface {
	IdentityGovernanceWorkflowExecutionTrigger() BaseIdentityGovernanceWorkflowExecutionTriggerImpl
}

var _ IdentityGovernanceWorkflowExecutionTrigger = BaseIdentityGovernanceWorkflowExecutionTriggerImpl{}

type BaseIdentityGovernanceWorkflowExecutionTriggerImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseIdentityGovernanceWorkflowExecutionTriggerImpl) IdentityGovernanceWorkflowExecutionTrigger() BaseIdentityGovernanceWorkflowExecutionTriggerImpl {
	return s
}

var _ IdentityGovernanceWorkflowExecutionTrigger = RawIdentityGovernanceWorkflowExecutionTriggerImpl{}

// RawIdentityGovernanceWorkflowExecutionTriggerImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawIdentityGovernanceWorkflowExecutionTriggerImpl struct {
	identityGovernanceWorkflowExecutionTrigger BaseIdentityGovernanceWorkflowExecutionTriggerImpl
	Type                                       string
	Values                                     map[string]interface{}
}

func (s RawIdentityGovernanceWorkflowExecutionTriggerImpl) IdentityGovernanceWorkflowExecutionTrigger() BaseIdentityGovernanceWorkflowExecutionTriggerImpl {
	return s.identityGovernanceWorkflowExecutionTrigger
}

func UnmarshalIdentityGovernanceWorkflowExecutionTriggerImplementation(input []byte) (IdentityGovernanceWorkflowExecutionTrigger, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling IdentityGovernanceWorkflowExecutionTrigger into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.identityGovernance.attributeChangeTrigger") {
		var out IdentityGovernanceAttributeChangeTrigger
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityGovernanceAttributeChangeTrigger: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identityGovernance.membershipChangeTrigger") {
		var out IdentityGovernanceMembershipChangeTrigger
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityGovernanceMembershipChangeTrigger: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.identityGovernance.timeBasedAttributeTrigger") {
		var out IdentityGovernanceTimeBasedAttributeTrigger
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IdentityGovernanceTimeBasedAttributeTrigger: %+v", err)
		}
		return out, nil
	}

	var parent BaseIdentityGovernanceWorkflowExecutionTriggerImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseIdentityGovernanceWorkflowExecutionTriggerImpl: %+v", err)
	}

	return RawIdentityGovernanceWorkflowExecutionTriggerImpl{
		identityGovernanceWorkflowExecutionTrigger: parent,
		Type:   value,
		Values: temp,
	}, nil

}
