package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AggregationOption struct {
	BucketDefinition *BucketAggregationDefinition `json:"bucketDefinition,omitempty"`

	// Computes aggregation on the field while the field exists in the current entity type. Required.
	Field string `json:"field"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The number of searchBucket resources to be returned. This isn't required when the range is provided manually in the
	// search request. The minimum accepted size is 1, and the maximum is 65535. Optional.
	Size nullable.Type[int64] `json:"size,omitempty"`
}
