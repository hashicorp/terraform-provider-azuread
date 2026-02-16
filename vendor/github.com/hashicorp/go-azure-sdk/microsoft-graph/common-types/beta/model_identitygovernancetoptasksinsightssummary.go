package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceTopTasksInsightsSummary struct {
	// Count of failed runs of the task.
	FailedTasks *int64 `json:"failedTasks,omitempty"`

	// Count of failed users who were processed by the task.
	FailedUsers *int64 `json:"failedUsers,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Count of successful runs of the task.
	SuccessfulTasks *int64 `json:"successfulTasks,omitempty"`

	// Count of successful users processed by the task.
	SuccessfulUsers *int64 `json:"successfulUsers,omitempty"`

	// The name of the task.
	TaskDefinitionDisplayName *string `json:"taskDefinitionDisplayName,omitempty"`

	// The task ID.
	TaskDefinitionId *string `json:"taskDefinitionId,omitempty"`

	// Count of total runs of the task.
	TotalTasks *int64 `json:"totalTasks,omitempty"`

	// Count of total users processed by the task.
	TotalUsers *int64 `json:"totalUsers,omitempty"`
}
