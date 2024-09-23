package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = IdentityGovernanceTaskReport{}

type IdentityGovernanceTaskReport struct {
	// The date time that the associated run completed. Value is null if the run has not completed.Supports $filter(lt, le,
	// gt, ge, eq, ne) and $orderby.
	CompletedDateTime nullable.Type[string] `json:"completedDateTime,omitempty"`

	// The number of users in the run execution for which the associated task failed.Supports $filter(lt, le, gt, ge, eq,
	// ne) and $orderby.
	FailedUsersCount *int64 `json:"failedUsersCount,omitempty"`

	// The date and time that the task report was last updated.
	LastUpdatedDateTime *string `json:"lastUpdatedDateTime,omitempty"`

	ProcessingStatus *IdentityGovernanceLifecycleWorkflowProcessingStatus `json:"processingStatus,omitempty"`

	// The unique identifier of the associated run.
	RunId nullable.Type[string] `json:"runId,omitempty"`

	// The date time that the associated run started. Value is null if the run has not started.
	StartedDateTime nullable.Type[string] `json:"startedDateTime,omitempty"`

	// The number of users in the run execution for which the associated task succeeded.Supports $filter(lt, le, gt, ge, eq,
	// ne) and $orderby.
	SuccessfulUsersCount *int64 `json:"successfulUsersCount,omitempty"`

	Task           *IdentityGovernanceTask           `json:"task,omitempty"`
	TaskDefinition *IdentityGovernanceTaskDefinition `json:"taskDefinition,omitempty"`

	// The related lifecycle workflow taskProcessingResults.
	TaskProcessingResults *[]IdentityGovernanceTaskProcessingResult `json:"taskProcessingResults,omitempty"`

	// The total number of users in the run execution for which the associated task was scheduled to execute.Supports
	// $filter(lt, le, gt, ge, eq, ne) and $orderby.
	TotalUsersCount *int64 `json:"totalUsersCount,omitempty"`

	// The number of users in the run execution for which the associated task is queued, in progress, or canceled.Supports
	// $filter(lt, le, gt, ge, eq, ne) and $orderby.
	UnprocessedUsersCount *int64 `json:"unprocessedUsersCount,omitempty"`

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

func (s IdentityGovernanceTaskReport) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IdentityGovernanceTaskReport{}

func (s IdentityGovernanceTaskReport) MarshalJSON() ([]byte, error) {
	type wrapper IdentityGovernanceTaskReport
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IdentityGovernanceTaskReport: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IdentityGovernanceTaskReport: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.identityGovernance.taskReport"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IdentityGovernanceTaskReport: %+v", err)
	}

	return encoded, nil
}
