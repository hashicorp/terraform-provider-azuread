package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LicenseAssignmentState struct {
	// Indicates whether the license is directly-assigned or inherited from a group. If directly-assigned, this field is
	// null; if inherited through a group membership, this field contains the ID of the group. Read-Only.
	AssignedByGroup nullable.Type[string] `json:"assignedByGroup,omitempty"`

	// The service plans that are disabled in this assignment. Read-Only.
	DisabledPlans *[]string `json:"disabledPlans,omitempty"`

	// License assignment failure error. If the license is assigned successfully, this field will be Null. Read-Only. The
	// possible values are CountViolation, MutuallyExclusiveViolation, DependencyViolation,
	// ProhibitedInUsageLocationViolation, UniquenessViolation, and Other. For more information on how to identify and
	// resolve license assignment errors, see here.
	Error nullable.Type[string] `json:"error,omitempty"`

	// The timestamp when the state of the license assignment was last updated.
	LastUpdatedDateTime nullable.Type[string] `json:"lastUpdatedDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The unique identifier for the SKU. Read-Only.
	SkuId nullable.Type[string] `json:"skuId,omitempty"`

	// Indicate the current state of this assignment. Read-Only. The possible values are Active, ActiveWithError, Disabled,
	// and Error.
	State nullable.Type[string] `json:"state,omitempty"`
}
