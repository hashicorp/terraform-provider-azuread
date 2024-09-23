package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceWorkflowsInsightsSummary struct {
	// Count of failed workflow runs processed in the tenant.
	FailedRuns *int64 `json:"failedRuns,omitempty"`

	// Count of failed tasks processed in the tenant.
	FailedTasks *int64 `json:"failedTasks,omitempty"`

	// Count of failed users processed by workflows in the tenant.
	FailedUsers *int64 `json:"failedUsers,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Count of successful workflow runs processed in the tenant.
	SuccessfulRuns *int64 `json:"successfulRuns,omitempty"`

	// Count of successful tasks processed in the tenant.
	SuccessfulTasks *int64 `json:"successfulTasks,omitempty"`

	// Count of successful users processed by workflows in the tenant.
	SuccessfulUsers *int64 `json:"successfulUsers,omitempty"`

	// Count of total workflows processed in the tenant.
	TotalRuns *int64 `json:"totalRuns,omitempty"`

	// Count of total tasks processed by workflows in the tenant.
	TotalTasks *int64 `json:"totalTasks,omitempty"`

	// Count of total users processed by workflows in the tenant.
	TotalUsers *int64 `json:"totalUsers,omitempty"`
}
