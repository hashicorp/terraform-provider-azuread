package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IdentityGovernanceWorkflowExecutionTrigger = IdentityGovernanceAttributeChangeTrigger{}

type IdentityGovernanceAttributeChangeTrigger struct {
	// The trigger attribute being changed that triggers the workflowexecutiontrigger of a workflow.)
	TriggerAttributes *[]IdentityGovernanceTriggerAttribute `json:"triggerAttributes,omitempty"`

	// Fields inherited from IdentityGovernanceWorkflowExecutionTrigger

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s IdentityGovernanceAttributeChangeTrigger) IdentityGovernanceWorkflowExecutionTrigger() BaseIdentityGovernanceWorkflowExecutionTriggerImpl {
	return BaseIdentityGovernanceWorkflowExecutionTriggerImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IdentityGovernanceAttributeChangeTrigger{}

func (s IdentityGovernanceAttributeChangeTrigger) MarshalJSON() ([]byte, error) {
	type wrapper IdentityGovernanceAttributeChangeTrigger
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IdentityGovernanceAttributeChangeTrigger: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IdentityGovernanceAttributeChangeTrigger: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.identityGovernance.attributeChangeTrigger"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IdentityGovernanceAttributeChangeTrigger: %+v", err)
	}

	return encoded, nil
}
