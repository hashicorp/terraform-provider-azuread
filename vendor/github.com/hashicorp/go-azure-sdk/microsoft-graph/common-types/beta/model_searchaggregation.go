package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SearchAggregation struct {
	// Defines the actual buckets of the computed aggregation.
	Buckets *[]SearchBucket `json:"buckets,omitempty"`

	// Defines on which field the aggregation was computed on.
	Field nullable.Type[string] `json:"field,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
