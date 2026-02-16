package stable

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceTopWorkflowsInsightsSummary struct {
	// Count of failed runs for workflow.
	FailedRuns *int64 `json:"failedRuns,omitempty"`

	// Count of failed users who were processed.
	FailedUsers *int64 `json:"failedUsers,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Count of successful runs of the workflow.
	SuccessfulRuns *int64 `json:"successfulRuns,omitempty"`

	// Count of successful users processed by the workflow.
	SuccessfulUsers *int64 `json:"successfulUsers,omitempty"`

	// Count of total runs of workflow.
	TotalRuns *int64 `json:"totalRuns,omitempty"`

	// Total number of users processed by the workflow.
	TotalUsers *int64 `json:"totalUsers,omitempty"`

	WorkflowCategory *IdentityGovernanceLifecycleWorkflowCategory `json:"workflowCategory,omitempty"`

	// The name of the workflow.
	WorkflowDisplayName *string `json:"workflowDisplayName,omitempty"`

	// The workflow ID.
	WorkflowId *string `json:"workflowId,omitempty"`

	// The version of the workflow that was a top workflow ran.
	WorkflowVersion *int64 `json:"workflowVersion,omitempty"`
}
