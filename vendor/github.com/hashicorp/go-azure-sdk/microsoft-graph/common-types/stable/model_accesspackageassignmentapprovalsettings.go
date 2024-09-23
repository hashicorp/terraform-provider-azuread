package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAssignmentApprovalSettings struct {
	// If false, then approval isn't required for new requests in this policy.
	IsApprovalRequiredForAdd nullable.Type[bool] `json:"isApprovalRequiredForAdd,omitempty"`

	// If false, then approval isn't required for updates to requests in this policy.
	IsApprovalRequiredForUpdate nullable.Type[bool] `json:"isApprovalRequiredForUpdate,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// If approval is required, the one, two or three elements of this collection define each of the stages of approval. An
	// empty array is present if no approval is required.
	Stages *[]AccessPackageApprovalStage `json:"stages,omitempty"`
}
