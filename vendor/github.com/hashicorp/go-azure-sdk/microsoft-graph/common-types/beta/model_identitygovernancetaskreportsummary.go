package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceTaskReportSummary struct {
	// The number of failed tasks in a report.
	FailedTasks *int64 `json:"failedTasks,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The total number of successful tasks in a report.
	SuccessfulTasks *int64 `json:"successfulTasks,omitempty"`

	// The total number of tasks in a report.
	TotalTasks *int64 `json:"totalTasks,omitempty"`

	// The number of unprocessed tasks in a report.
	UnprocessedTasks *int64 `json:"unprocessedTasks,omitempty"`
}
