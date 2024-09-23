package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ApprovalWorkflowProvider{}

type ApprovalWorkflowProvider struct {
	BusinessFlows                               *[]BusinessFlow             `json:"businessFlows,omitempty"`
	BusinessFlowsWithRequestsAwaitingMyDecision *[]BusinessFlow             `json:"businessFlowsWithRequestsAwaitingMyDecision,omitempty"`
	DisplayName                                 *string                     `json:"displayName,omitempty"`
	PolicyTemplates                             *[]GovernancePolicyTemplate `json:"policyTemplates,omitempty"`

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

func (s ApprovalWorkflowProvider) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ApprovalWorkflowProvider{}

func (s ApprovalWorkflowProvider) MarshalJSON() ([]byte, error) {
	type wrapper ApprovalWorkflowProvider
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ApprovalWorkflowProvider: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ApprovalWorkflowProvider: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.approvalWorkflowProvider"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ApprovalWorkflowProvider: %+v", err)
	}

	return encoded, nil
}
