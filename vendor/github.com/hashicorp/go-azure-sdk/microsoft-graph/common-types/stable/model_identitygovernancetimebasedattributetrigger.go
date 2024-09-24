package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IdentityGovernanceWorkflowExecutionTrigger = IdentityGovernanceTimeBasedAttributeTrigger{}

type IdentityGovernanceTimeBasedAttributeTrigger struct {
	// How many days before or after the time-based attribute specified the workflow should trigger. For example, if the
	// attribute is employeeHireDate and offsetInDays is -1, then the workflow should trigger one day before the employee
	// hire date. The value can range between -180 and 180 days.
	OffsetInDays *int64 `json:"offsetInDays,omitempty"`

	TimeBasedAttribute *IdentityGovernanceWorkflowTriggerTimeBasedAttribute `json:"timeBasedAttribute,omitempty"`

	// Fields inherited from IdentityGovernanceWorkflowExecutionTrigger

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s IdentityGovernanceTimeBasedAttributeTrigger) IdentityGovernanceWorkflowExecutionTrigger() BaseIdentityGovernanceWorkflowExecutionTriggerImpl {
	return BaseIdentityGovernanceWorkflowExecutionTriggerImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IdentityGovernanceTimeBasedAttributeTrigger{}

func (s IdentityGovernanceTimeBasedAttributeTrigger) MarshalJSON() ([]byte, error) {
	type wrapper IdentityGovernanceTimeBasedAttributeTrigger
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IdentityGovernanceTimeBasedAttributeTrigger: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IdentityGovernanceTimeBasedAttributeTrigger: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.identityGovernance.timeBasedAttributeTrigger"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IdentityGovernanceTimeBasedAttributeTrigger: %+v", err)
	}

	return encoded, nil
}
