package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceWorkflowsInsightsByCategory struct {
	// Failed 'Joiner' workflows processed in a tenant.
	FailedJoinerRuns *int64 `json:"failedJoinerRuns,omitempty"`

	// Failed 'Leaver' workflows processed in a tenant.
	FailedLeaverRuns *int64 `json:"failedLeaverRuns,omitempty"`

	// Failed 'Mover' workflows processed in a tenant.
	FailedMoverRuns *int64 `json:"failedMoverRuns,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Successful 'Joiner' workflows processed in a tenant.
	SuccessfulJoinerRuns *int64 `json:"successfulJoinerRuns,omitempty"`

	// Successful 'Leaver' workflows processed in a tenant.
	SuccessfulLeaverRuns *int64 `json:"successfulLeaverRuns,omitempty"`

	// Successful 'Mover' workflows processed in a tenant.
	SuccessfulMoverRuns *int64 `json:"successfulMoverRuns,omitempty"`

	// Total 'Joiner' workflows processed in a tenant.
	TotalJoinerRuns *int64 `json:"totalJoinerRuns,omitempty"`

	// Total 'Leaver' workflows processed in a tenant.
	TotalLeaverRuns *int64 `json:"totalLeaverRuns,omitempty"`

	// Total 'Mover' workflows processed in a tenant.
	TotalMoverRuns *int64 `json:"totalMoverRuns,omitempty"`
}
