package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = IdentityGovernanceUserProcessingResult{}

type IdentityGovernanceUserProcessingResult struct {
	// The date time that the workflow execution for a user completed. Value is null if the workflow hasn't
	// completed.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
	CompletedDateTime nullable.Type[string] `json:"completedDateTime,omitempty"`

	// The number of tasks that failed in the workflow execution.
	FailedTasksCount *int64 `json:"failedTasksCount,omitempty"`

	ProcessingStatus *IdentityGovernanceLifecycleWorkflowProcessingStatus `json:"processingStatus,omitempty"`

	// The date time that the workflow is scheduled to be executed for a user.Supports $filter(lt, le, gt, ge, eq, ne) and
	// $orderby.
	ScheduledDateTime *string `json:"scheduledDateTime,omitempty"`

	// The date time that the workflow execution started. Value is null if the workflow execution has not started.Supports
	// $filter(lt, le, gt, ge, eq, ne) and $orderby.
	StartedDateTime nullable.Type[string] `json:"startedDateTime,omitempty"`

	Subject *User `json:"subject,omitempty"`

	// The associated individual task execution.
	TaskProcessingResults *[]IdentityGovernanceTaskProcessingResult `json:"taskProcessingResults,omitempty"`

	// The total number of tasks that in the workflow execution.
	TotalTasksCount *int64 `json:"totalTasksCount,omitempty"`

	// The total number of unprocessed tasks for the workflow.
	TotalUnprocessedTasksCount *int64 `json:"totalUnprocessedTasksCount,omitempty"`

	WorkflowExecutionType *IdentityGovernanceWorkflowExecutionType `json:"workflowExecutionType,omitempty"`

	// The version of the workflow that was executed.
	WorkflowVersion *int64 `json:"workflowVersion,omitempty"`

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

func (s IdentityGovernanceUserProcessingResult) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IdentityGovernanceUserProcessingResult{}

func (s IdentityGovernanceUserProcessingResult) MarshalJSON() ([]byte, error) {
	type wrapper IdentityGovernanceUserProcessingResult
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IdentityGovernanceUserProcessingResult: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IdentityGovernanceUserProcessingResult: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.identityGovernance.userProcessingResult"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IdentityGovernanceUserProcessingResult: %+v", err)
	}

	return encoded, nil
}
