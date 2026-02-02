package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApprovalItemViewPoint struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Collection of roles associated with the requesting user for the approval item. If the owner of the approval item is
	// making the request, the collection of roles includes the role owner. If the requesting user was assigned as an
	// approver, the collection includes the role approver.
	Roles *[]ApproverRole `json:"roles,omitempty"`
}
