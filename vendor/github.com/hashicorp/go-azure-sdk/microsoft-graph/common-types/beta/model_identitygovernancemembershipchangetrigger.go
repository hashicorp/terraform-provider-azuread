package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IdentityGovernanceWorkflowExecutionTrigger = IdentityGovernanceMembershipChangeTrigger{}

type IdentityGovernanceMembershipChangeTrigger struct {
	ChangeType *IdentityGovernanceMembershipChangeType `json:"changeType,omitempty"`

	// Fields inherited from IdentityGovernanceWorkflowExecutionTrigger

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s IdentityGovernanceMembershipChangeTrigger) IdentityGovernanceWorkflowExecutionTrigger() BaseIdentityGovernanceWorkflowExecutionTriggerImpl {
	return BaseIdentityGovernanceWorkflowExecutionTriggerImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IdentityGovernanceMembershipChangeTrigger{}

func (s IdentityGovernanceMembershipChangeTrigger) MarshalJSON() ([]byte, error) {
	type wrapper IdentityGovernanceMembershipChangeTrigger
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IdentityGovernanceMembershipChangeTrigger: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IdentityGovernanceMembershipChangeTrigger: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.identityGovernance.membershipChangeTrigger"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IdentityGovernanceMembershipChangeTrigger: %+v", err)
	}

	return encoded, nil
}
