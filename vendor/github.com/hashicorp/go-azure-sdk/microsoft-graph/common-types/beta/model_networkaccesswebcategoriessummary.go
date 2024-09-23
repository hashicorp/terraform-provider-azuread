package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessWebCategoriesSummary struct {
	Action *NetworkaccessFilteringPolicyAction `json:"action,omitempty"`

	// The number of unique devices that were seen.
	DeviceCount *int64 `json:"deviceCount,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The number of transactions that were seen.
	TransactionCount *int64 `json:"transactionCount,omitempty"`

	// The number of unique Microsoft Entra ID users that were seen.
	UserCount *int64 `json:"userCount,omitempty"`

	WebCategory *NetworkaccessWebCategory `json:"webCategory,omitempty"`
}
