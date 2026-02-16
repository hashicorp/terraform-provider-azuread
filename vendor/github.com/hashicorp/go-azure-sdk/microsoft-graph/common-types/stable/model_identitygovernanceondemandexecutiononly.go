package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IdentityGovernanceWorkflowExecutionConditions = IdentityGovernanceOnDemandExecutionOnly{}

type IdentityGovernanceOnDemandExecutionOnly struct {

	// Fields inherited from IdentityGovernanceWorkflowExecutionConditions

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s IdentityGovernanceOnDemandExecutionOnly) IdentityGovernanceWorkflowExecutionConditions() BaseIdentityGovernanceWorkflowExecutionConditionsImpl {
	return BaseIdentityGovernanceWorkflowExecutionConditionsImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IdentityGovernanceOnDemandExecutionOnly{}

func (s IdentityGovernanceOnDemandExecutionOnly) MarshalJSON() ([]byte, error) {
	type wrapper IdentityGovernanceOnDemandExecutionOnly
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IdentityGovernanceOnDemandExecutionOnly: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IdentityGovernanceOnDemandExecutionOnly: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.identityGovernance.onDemandExecutionOnly"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IdentityGovernanceOnDemandExecutionOnly: %+v", err)
	}

	return encoded, nil
}
