package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsQualityUpdateCatalogItemPolicyDetail struct {
	// Enum to describe policy's approval status for catalogitems
	ApprovalStatus *WindowsQualityUpdateApprovalStatus `json:"approvalStatus,omitempty"`

	// Catalog item id for this approval intend
	CatalogItemId *string `json:"catalogItemId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Policy Id for this approval intend
	PolicyId *string `json:"policyId,omitempty"`
}
