package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ShiftAvailability struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Specifies the pattern for recurrence
	Recurrence *PatternedRecurrence `json:"recurrence,omitempty"`

	// The time slot(s) preferred by the user.
	TimeSlots *[]TimeRange `json:"timeSlots,omitempty"`

	// Specifies the time zone for the indicated time.
	TimeZone nullable.Type[string] `json:"timeZone,omitempty"`
}
