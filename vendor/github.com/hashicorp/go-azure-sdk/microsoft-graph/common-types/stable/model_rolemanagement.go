package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleManagement struct {
	Directory *RbacApplication `json:"directory,omitempty"`

	// Container for roles and assignments for entitlement management resources.
	EntitlementManagement *RbacApplication `json:"entitlementManagement,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
