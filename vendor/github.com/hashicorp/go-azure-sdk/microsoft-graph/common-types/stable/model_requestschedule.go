package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RequestSchedule struct {
	// When the eligible or active assignment expires.
	Expiration *ExpirationPattern `json:"expiration,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The frequency of the eligible or active assignment. This property is currently unsupported in PIM.
	Recurrence *PatternedRecurrence `json:"recurrence,omitempty"`

	// When the eligible or active assignment becomes active.
	StartDateTime nullable.Type[string] `json:"startDateTime,omitempty"`
}
