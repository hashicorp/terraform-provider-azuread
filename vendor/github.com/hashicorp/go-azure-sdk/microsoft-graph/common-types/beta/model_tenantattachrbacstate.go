package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TenantAttachRBACState struct {
	// Indicates whether the tenant is enabled for Tenant Attach with role management. TRUE if enabled, FALSE if the Tenant
	// Attach with rolemanagement is disabled.
	Enabled *bool `json:"enabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
