package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReconciliationCounter struct {
	CorrelatedObjectCount nullable.Type[int64] `json:"correlatedObjectCount,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	SourceObjectCount       nullable.Type[int64] `json:"sourceObjectCount,omitempty"`
	TargetObjectCount       nullable.Type[int64] `json:"targetObjectCount,omitempty"`
	UncorrelatedObjectCount nullable.Type[int64] `json:"uncorrelatedObjectCount,omitempty"`
}
