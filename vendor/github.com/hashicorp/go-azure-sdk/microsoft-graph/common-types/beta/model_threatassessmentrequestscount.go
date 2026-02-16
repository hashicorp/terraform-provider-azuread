package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ThreatAssessmentRequestsCount struct {
	Count           nullable.Type[int64]  `json:"count,omitempty"`
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	PivotValue nullable.Type[string] `json:"pivotValue,omitempty"`
}
