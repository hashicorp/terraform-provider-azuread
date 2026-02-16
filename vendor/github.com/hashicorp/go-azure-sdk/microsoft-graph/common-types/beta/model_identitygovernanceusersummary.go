package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceUserSummary struct {
	// The number of failed tasks for users in a user summary.
	FailedTasks *int64 `json:"failedTasks,omitempty"`

	// The number of failed users in a user summary.
	FailedUsers *int64 `json:"failedUsers,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The number of successful users in a user summary.
	SuccessfulUsers *int64 `json:"successfulUsers,omitempty"`

	// The total tasks of users in a user summary.
	TotalTasks *int64 `json:"totalTasks,omitempty"`

	// The total number of users in a user summary
	TotalUsers *int64 `json:"totalUsers,omitempty"`
}
