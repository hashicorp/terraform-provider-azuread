package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RolePermission struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Resource Actions each containing a set of allowed and not allowed permissions.
	ResourceActions *[]ResourceAction `json:"resourceActions,omitempty"`
}
