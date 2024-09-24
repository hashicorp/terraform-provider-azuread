package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamsAppResourceSpecificPermission struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The type of resource-specific permission.
	PermissionType *TeamsAppResourceSpecificPermissionType `json:"permissionType,omitempty"`

	// The name of the resource-specific permission.
	PermissionValue nullable.Type[string] `json:"permissionValue,omitempty"`
}
