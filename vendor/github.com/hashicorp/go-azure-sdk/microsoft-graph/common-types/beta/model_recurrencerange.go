package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecurrenceRange struct {
	// The date to stop applying the recurrence pattern. Depending on the recurrence pattern of the event, the last
	// occurrence of the meeting may not be this date. Required if type is endDate.
	EndDate nullable.Type[string] `json:"endDate,omitempty"`

	// The number of times to repeat the event. Required and must be positive if type is numbered.
	NumberOfOccurrences *int64 `json:"numberOfOccurrences,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Time zone for the startDate and endDate properties. Optional. If not specified, the time zone of the event is used.
	RecurrenceTimeZone nullable.Type[string] `json:"recurrenceTimeZone,omitempty"`

	// The date to start applying the recurrence pattern. The first occurrence of the meeting may be this date or later,
	// depending on the recurrence pattern of the event. Must be the same value as the start property of the recurring
	// event. Required.
	StartDate nullable.Type[string] `json:"startDate,omitempty"`

	// The recurrence range. Possible values are: endDate, noEnd, numbered. Required.
	Type RecurrenceRangeType `json:"type"`
}
