package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationProgress struct {
	// The numerator of a progress ratio; the number of units of changes already processed.
	CompletedUnits *int64 `json:"completedUnits,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The time of a progress observation as an offset in minutes from UTC.
	ProgressObservationDateTime *string `json:"progressObservationDateTime,omitempty"`

	// The denominator of a progress ratio; a number of units of changes to be processed to accomplish synchronization.
	TotalUnits *int64 `json:"totalUnits,omitempty"`

	// An optional description of the units.
	Units nullable.Type[string] `json:"units,omitempty"`
}
