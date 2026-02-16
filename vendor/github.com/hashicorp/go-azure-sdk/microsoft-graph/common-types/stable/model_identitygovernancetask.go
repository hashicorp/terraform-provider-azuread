package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = IdentityGovernanceTask{}

type IdentityGovernanceTask struct {
	// Arguments included within the task. For guidance to configure this property, see Configure the arguments for built-in
	// Lifecycle Workflow tasks. Required.
	Arguments []KeyValuePair `json:"arguments"`

	Category *IdentityGovernanceLifecycleTaskCategory `json:"category,omitempty"`

	// A Boolean value that specifies whether, if this task fails, the workflow stops, and subsequent tasks aren't run.
	// Optional.
	ContinueOnError *bool `json:"continueOnError,omitempty"`

	// A string that describes the purpose of the task for administrative use. Optional.
	Description nullable.Type[string] `json:"description,omitempty"`

	// A unique string that identifies the task. Required.Supports $filter(eq, ne) and orderBy.
	DisplayName string `json:"displayName"`

	// An integer that states in what order the task runs in a workflow.Supports $orderby.
	ExecutionSequence *int64 `json:"executionSequence,omitempty"`

	// A Boolean value that denotes whether the task is set to run or not. Optional.Supports $filter(eq, ne) and orderBy.
	IsEnabled *bool `json:"isEnabled,omitempty"`

	// A unique template identifier for the task. For more information about the tasks that Lifecycle Workflows currently
	// supports and their unique identifiers, see Configure the arguments for built-in Lifecycle Workflow tasks.
	// Required.Supports $filter(eq, ne).
	TaskDefinitionId string `json:"taskDefinitionId"`

	// The result of processing the task.
	TaskProcessingResults *[]IdentityGovernanceTaskProcessingResult `json:"taskProcessingResults,omitempty"`

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

func (s IdentityGovernanceTask) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IdentityGovernanceTask{}

func (s IdentityGovernanceTask) MarshalJSON() ([]byte, error) {
	type wrapper IdentityGovernanceTask
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IdentityGovernanceTask: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IdentityGovernanceTask: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.identityGovernance.task"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IdentityGovernanceTask: %+v", err)
	}

	return encoded, nil
}
