package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessCrossTenantAccess struct {
	// The number of devices that accessed the external tenant.
	DeviceCount *int64 `json:"deviceCount,omitempty"`

	// The timestamp of the most recent access to the external tenant.
	LastAccessDateTime *string `json:"lastAccessDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The tenant ID of the external tenant.
	ResourceTenantId *string `json:"resourceTenantId,omitempty"`

	// The name of the external tenant.
	ResourceTenantName nullable.Type[string] `json:"resourceTenantName,omitempty"`

	// The domain of the external tenant.
	ResourceTenantPrimaryDomain *string `json:"resourceTenantPrimaryDomain,omitempty"`

	UsageStatus *NetworkaccessUsageStatus `json:"usageStatus,omitempty"`

	// The number of users that accessed the external tenant.
	UserCount *int64 `json:"userCount,omitempty"`
}
