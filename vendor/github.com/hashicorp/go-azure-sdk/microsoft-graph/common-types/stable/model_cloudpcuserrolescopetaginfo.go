package stable

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCUserRoleScopeTagInfo struct {
	// Scope tag display name.
	DisplayName *string `json:"displayName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Scope tag identity.
	RoleScopeTagId *string `json:"roleScopeTagId,omitempty"`
}
