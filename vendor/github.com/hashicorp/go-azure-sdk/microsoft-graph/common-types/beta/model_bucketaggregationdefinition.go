package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BucketAggregationDefinition struct {
	// True to specify the sort order as descending. The default is false, with the sort order as ascending. Optional.
	IsDescending nullable.Type[bool] `json:"isDescending,omitempty"`

	// The minimum number of items that should be present in the aggregation to be returned in a bucket. Optional.
	MinimumCount nullable.Type[int64] `json:"minimumCount,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// A filter to define a matching criteria. The key should start with the specified prefix to be returned in the
	// response. Optional.
	PrefixFilter nullable.Type[string] `json:"prefixFilter,omitempty"`

	// Specifies the manual ranges to compute the aggregations. This is only valid for nonstring refiners of date or numeric
	// type. Optional.
	Ranges *[]BucketAggregationRange `json:"ranges,omitempty"`

	SortBy *BucketAggregationSortProperty `json:"sortBy,omitempty"`
}
