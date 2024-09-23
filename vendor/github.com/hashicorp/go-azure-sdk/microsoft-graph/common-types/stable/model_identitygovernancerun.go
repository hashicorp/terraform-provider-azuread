package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = IdentityGovernanceRun{}

type IdentityGovernanceRun struct {
	// The date time that the run completed. Value is null if the workflow hasn't completed.Supports $filter(lt, le, gt, ge,
	// eq, ne) and $orderby.
	CompletedDateTime nullable.Type[string] `json:"completedDateTime,omitempty"`

	// The number of tasks that failed in the run execution.
	FailedTasksCount *int64 `json:"failedTasksCount,omitempty"`

	// The number of users that failed in the run execution.
	FailedUsersCount *int64 `json:"failedUsersCount,omitempty"`

	// The datetime that the run was last updated.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
	LastUpdatedDateTime nullable.Type[string] `json:"lastUpdatedDateTime,omitempty"`

	ProcessingStatus *IdentityGovernanceLifecycleWorkflowProcessingStatus `json:"processingStatus,omitempty"`

	// The date time that the run is scheduled to be executed for a workflow.Supports $filter(lt, le, gt, ge, eq, ne) and
	// $orderby.
	ScheduledDateTime *string `json:"scheduledDateTime,omitempty"`

	// The date time that the run execution started.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
	StartedDateTime nullable.Type[string] `json:"startedDateTime,omitempty"`

	// The number of successfully completed users in the run.
	SuccessfulUsersCount *int64 `json:"successfulUsersCount,omitempty"`

	// The related taskProcessingResults.
	TaskProcessingResults *[]IdentityGovernanceTaskProcessingResult `json:"taskProcessingResults,omitempty"`

	TotalTasksCount *int64 `json:"totalTasksCount,omitempty"`

	// The total number of unprocessed tasks in the run execution.
	TotalUnprocessedTasksCount *int64 `json:"totalUnprocessedTasksCount,omitempty"`

	// The total number of users in the workflow execution.
	TotalUsersCount *int64 `json:"totalUsersCount,omitempty"`

	// The associated individual user execution.
	UserProcessingResults *[]IdentityGovernanceUserProcessingResult `json:"userProcessingResults,omitempty"`

	WorkflowExecutionType *IdentityGovernanceWorkflowExecutionType `json:"workflowExecutionType,omitempty"`

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

func (s IdentityGovernanceRun) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IdentityGovernanceRun{}

func (s IdentityGovernanceRun) MarshalJSON() ([]byte, error) {
	type wrapper IdentityGovernanceRun
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IdentityGovernanceRun: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IdentityGovernanceRun: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.identityGovernance.run"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IdentityGovernanceRun: %+v", err)
	}

	return encoded, nil
}
