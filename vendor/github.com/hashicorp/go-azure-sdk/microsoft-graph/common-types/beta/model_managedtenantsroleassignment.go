package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsRoleAssignment struct {
	// The type of the admin relationship(s) associated with the role assignment. Possible values are: none,
	// delegatedAdminPrivileges, unknownFutureValue, granularDelegatedAdminPrivileges,
	// delegatedAndGranularDelegetedAdminPrivileges. Use the Prefer: include-unknown-enum-members request header to get the
	// following values from this evolvable enum: granularDelegatedAdminPrivileges ,
	// delegatedAndGranularDelegetedAdminPrivileges.
	AssignmentType *ManagedTenantsDelegatedPrivilegeStatus `json:"assignmentType,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The collection of roles assigned.
	Roles *[]ManagedTenantsRoleDefinition `json:"roles,omitempty"`
}
