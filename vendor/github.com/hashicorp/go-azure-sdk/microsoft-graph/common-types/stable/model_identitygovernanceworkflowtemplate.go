package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = IdentityGovernanceWorkflowTemplate{}

type IdentityGovernanceWorkflowTemplate struct {
	Category *IdentityGovernanceLifecycleWorkflowCategory `json:"category,omitempty"`

	// The description of the workflowTemplate.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The display name of the workflowTemplate.Supports $filter(eq, ne) and $orderby.
	DisplayName *string `json:"displayName,omitempty"`

	// Conditions describing when to execute the workflow and the criteria to identify in-scope subject set.
	ExecutionConditions IdentityGovernanceWorkflowExecutionConditions `json:"executionConditions"`

	// Represents the configured tasks to execute and their execution sequence within a workflow. This relationship is
	// expanded by default.
	Tasks *[]IdentityGovernanceTask `json:"tasks,omitempty"`

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

func (s IdentityGovernanceWorkflowTemplate) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IdentityGovernanceWorkflowTemplate{}

func (s IdentityGovernanceWorkflowTemplate) MarshalJSON() ([]byte, error) {
	type wrapper IdentityGovernanceWorkflowTemplate
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IdentityGovernanceWorkflowTemplate: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IdentityGovernanceWorkflowTemplate: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.identityGovernance.workflowTemplate"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IdentityGovernanceWorkflowTemplate: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &IdentityGovernanceWorkflowTemplate{}

func (s *IdentityGovernanceWorkflowTemplate) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Category    *IdentityGovernanceLifecycleWorkflowCategory `json:"category,omitempty"`
		Description nullable.Type[string]                        `json:"description,omitempty"`
		DisplayName *string                                      `json:"displayName,omitempty"`
		Tasks       *[]IdentityGovernanceTask                    `json:"tasks,omitempty"`
		Id          *string                                      `json:"id,omitempty"`
		ODataId     *string                                      `json:"@odata.id,omitempty"`
		ODataType   *string                                      `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Category = decoded.Category
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.Tasks = decoded.Tasks
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling IdentityGovernanceWorkflowTemplate into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["executionConditions"]; ok {
		impl, err := UnmarshalIdentityGovernanceWorkflowExecutionConditionsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ExecutionConditions' for 'IdentityGovernanceWorkflowTemplate': %+v", err)
		}
		s.ExecutionConditions = impl
	}

	return nil
}
