package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessTransactionSummary struct {
	// The number of transactions that were blocked.
	BlockedCount *int64 `json:"blockedCount,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The total number of transactions.
	TotalCount *int64 `json:"totalCount,omitempty"`

	TrafficType *NetworkaccessTrafficType `json:"trafficType,omitempty"`
}
