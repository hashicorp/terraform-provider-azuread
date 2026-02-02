package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceAccess struct {
	// The unique identifier of an app role or delegated permission exposed by the resource application. For delegated
	// permissions, this should match the id property of one of the delegated permissions in the oauth2PermissionScopes
	// collection of the resource application's service principal. For app roles (application permissions), this should
	// match the id property of an app role in the appRoles collection of the resource application's service principal.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Specifies whether the id property references a delegated permission or an app role (application permission). The
	// possible values are: Scope (for delegated permissions) or Role (for app roles).
	Type nullable.Type[string] `json:"type,omitempty"`
}
