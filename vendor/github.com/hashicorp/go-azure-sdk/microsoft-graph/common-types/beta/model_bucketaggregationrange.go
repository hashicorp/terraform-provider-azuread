package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BucketAggregationRange struct {
	// Defines the lower bound from which to compute the aggregation. This can be a numeric value or a string representation
	// of a date using the YYYY-MM-DDTHH:mm:ss.sssZ format. Required.
	From string `json:"from"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Defines the upper bound up to which to compute the aggregation. This can be a numeric value or a string
	// representation of a date using the YYYY-MM-DDTHH:mm:ss.sssZ format. Required.
	To string `json:"to"`
}
