package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApprovalSettings struct {
	// One of SingleStage, Serial, Parallel, NoApproval (default). NoApproval is used when isApprovalRequired is false.
	ApprovalMode nullable.Type[string] `json:"approvalMode,omitempty"`

	// If approval is required, the one or two elements of this collection define each of the stages of approval. An empty
	// array if no approval is required.
	ApprovalStages *[]UnifiedApprovalStage `json:"approvalStages,omitempty"`

	// Indicates whether approval is required for requests in this policy.
	IsApprovalRequired nullable.Type[bool] `json:"isApprovalRequired,omitempty"`

	// Indicates whether approval is required for a user to extend their assignment.
	IsApprovalRequiredForExtension nullable.Type[bool] `json:"isApprovalRequiredForExtension,omitempty"`

	// Indicates whether the requestor is required to supply a justification in their request.
	IsRequestorJustificationRequired nullable.Type[bool] `json:"isRequestorJustificationRequired,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
