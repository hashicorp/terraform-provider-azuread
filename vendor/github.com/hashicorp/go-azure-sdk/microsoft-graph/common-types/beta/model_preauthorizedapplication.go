package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PreAuthorizedApplication struct {
	// The unique identifier for the client application.
	AppId *string `json:"appId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The unique identifier for the scopes the client application is granted.
	PermissionIds *[]string `json:"permissionIds,omitempty"`
}
