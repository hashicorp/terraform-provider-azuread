package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceRunSummary struct {
	// The number of failed workflow runs.
	FailedRuns *int64 `json:"failedRuns,omitempty"`

	// The number of failed tasks of a workflow.
	FailedTasks *int64 `json:"failedTasks,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The number of successful workflow runs.
	SuccessfulRuns *int64 `json:"successfulRuns,omitempty"`

	// The total number of runs for a workflow.
	TotalRuns *int64 `json:"totalRuns,omitempty"`

	// The total number of tasks processed by a workflow.
	TotalTasks *int64 `json:"totalTasks,omitempty"`

	// The total number of users processed by a workflow.
	TotalUsers *int64 `json:"totalUsers,omitempty"`
}
