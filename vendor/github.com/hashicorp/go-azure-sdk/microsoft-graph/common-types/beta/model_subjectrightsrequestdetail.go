package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubjectRightsRequestDetail struct {
	// Count of items that are excluded from the request.
	ExcludedItemCount nullable.Type[int64] `json:"excludedItemCount,omitempty"`

	// Count of items per insight.
	InsightCounts *[]KeyValuePair `json:"insightCounts,omitempty"`

	// Count of items found.
	ItemCount nullable.Type[int64] `json:"itemCount,omitempty"`

	// Count of item that need review.
	ItemNeedReview nullable.Type[int64] `json:"itemNeedReview,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Count of items per product, such as Exchange, SharePoint, OneDrive, and Teams.
	ProductItemCounts *[]KeyValuePair `json:"productItemCounts,omitempty"`

	// Count of items signed off by the administrator.
	SignedOffItemCount nullable.Type[int64] `json:"signedOffItemCount,omitempty"`

	// Total item size in bytes.
	TotalItemSize nullable.Type[int64] `json:"totalItemSize,omitempty"`
}
