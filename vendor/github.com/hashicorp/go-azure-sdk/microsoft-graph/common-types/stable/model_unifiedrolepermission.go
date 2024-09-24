package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnifiedRolePermission struct {
	// Set of tasks that can be performed on a resource. Required.
	AllowedResourceActions []string `json:"allowedResourceActions"`

	// Optional constraints that must be met for the permission to be effective. Not supported for custom roles.
	Condition nullable.Type[string] `json:"condition,omitempty"`

	// Set of tasks that may not be performed on a resource. Not yet supported.
	ExcludedResourceActions *[]string `json:"excludedResourceActions,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
