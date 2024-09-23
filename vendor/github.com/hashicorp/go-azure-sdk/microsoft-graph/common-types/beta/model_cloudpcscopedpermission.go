package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCScopedPermission struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The operations allowed on scoped resources for the authenticated user. Example permission is
	// Microsoft.CloudPC/ProvisioningPolicies/Create.
	Permission nullable.Type[string] `json:"permission,omitempty"`

	// The scope IDs of corresponding permission. Currently, it's Intune scope tag ID.
	ScopeIds *[]string `json:"scopeIds,omitempty"`
}
