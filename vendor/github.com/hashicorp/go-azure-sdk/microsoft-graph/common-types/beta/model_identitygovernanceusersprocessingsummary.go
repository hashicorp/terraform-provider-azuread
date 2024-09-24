package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceUsersProcessingSummary struct {
	FailedTasks *int64 `json:"failedTasks,omitempty"`
	FailedUsers *int64 `json:"failedUsers,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	SuccessfulUsers *int64 `json:"successfulUsers,omitempty"`
	TotalTasks      *int64 `json:"totalTasks,omitempty"`
	TotalUsers      *int64 `json:"totalUsers,omitempty"`
}
