package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GovernanceRoleAssignmentRequestStatus struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The status of the role assignment request. The value can be InProgress or Closed.
	Status nullable.Type[string] `json:"status,omitempty"`

	// The details of the status of the role assignment request. It represents the evaluation results of different rules.
	StatusDetails *[]KeyValue `json:"statusDetails,omitempty"`

	// The sub status of the role assignment request. The values can be Accepted, PendingEvaluation, Granted, Denied,
	// PendingProvisioning, Provisioned, PendingRevocation, Revoked, Canceled, Failed, PendingApprovalProvisioning,
	// PendingApproval, FailedAsResourceIsLocked, PendingAdminDecision, AdminApproved, AdminDenied, TimedOut, and
	// ProvisioningStarted.
	SubStatus nullable.Type[string] `json:"subStatus,omitempty"`
}
