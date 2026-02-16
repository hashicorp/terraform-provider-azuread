package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCBulkActionSummary struct {
	// The number of Cloud PCs where the action failed.
	FailedCount *int64 `json:"failedCount,omitempty"`

	// The number of Cloud PCs where the action is in progress.
	InProgressCount *int64 `json:"inProgressCount,omitempty"`

	// The number of Cloud PCs where the action isn't supported.
	NotSupportedCount *int64 `json:"notSupportedCount,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The number of Cloud PCs where the action is pending.
	PendingCount *int64 `json:"pendingCount,omitempty"`

	// The number of Cloud PCs where the action is successful.
	SuccessfulCount *int64 `json:"successfulCount,omitempty"`
}
