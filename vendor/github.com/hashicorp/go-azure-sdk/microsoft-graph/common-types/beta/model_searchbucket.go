package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SearchBucket struct {
	// A token containing the encoded filter to aggregate search matches by the specific key value. To use the filter, pass
	// the token as part of the aggregationFilter property in a searchRequest object, in the format
	// '{field}:/'{aggregationFilterToken}/''. See an example.
	AggregationFilterToken nullable.Type[string] `json:"aggregationFilterToken,omitempty"`

	// The approximate number of search matches that share the same value specified in the key property. Note that this
	// number is not the exact number of matches.
	Count nullable.Type[int64] `json:"count,omitempty"`

	// The discrete value of the field that an aggregation was computed on.
	Key nullable.Type[string] `json:"key,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
