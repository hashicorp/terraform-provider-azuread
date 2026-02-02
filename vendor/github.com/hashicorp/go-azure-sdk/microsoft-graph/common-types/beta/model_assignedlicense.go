package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignedLicense struct {
	// A collection of the unique identifiers for plans that have been disabled. IDs are available in servicePlans >
	// servicePlanId in the tenant's subscribedSkus or serviceStatus > servicePlanId in the tenant's companySubscription.
	DisabledPlans *[]string `json:"disabledPlans,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The unique identifier for the SKU. Corresponds to the skuId from subscribedSkus or companySubscription.
	SkuId nullable.Type[string] `json:"skuId,omitempty"`
}
