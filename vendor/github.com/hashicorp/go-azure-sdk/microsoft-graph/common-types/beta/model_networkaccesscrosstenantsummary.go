package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessCrossTenantSummary struct {
	// The total number of authentication sessions between startDateTime and endDateTime.
	AuthTransactionCount *int64 `json:"authTransactionCount,omitempty"`

	// The number of unique devices that performed cross-tenant access.
	DeviceCount *int64 `json:"deviceCount,omitempty"`

	// The number of unique tenants that were accessed between endDateTime and discoveryPivotDateTime, but weren't accessed
	// between discoveryPivotDateTime and startDateTime.
	NewTenantCount *int64 `json:"newTenantCount,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The number of tenants that are rarely used.
	RarelyUsedTenantCount *int64 `json:"rarelyUsedTenantCount,omitempty"`

	// The number of unique tenants that were accessed, not including the device's tenant.
	TenantCount *int64 `json:"tenantCount,omitempty"`

	// The number of unique users that performed cross-tenant access.
	UserCount *int64 `json:"userCount,omitempty"`
}
