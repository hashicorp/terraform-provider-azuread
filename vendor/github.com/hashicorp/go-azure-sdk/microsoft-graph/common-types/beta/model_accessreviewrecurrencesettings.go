package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewRecurrenceSettings struct {
	// The duration in days for recurrence.
	DurationInDays *int64 `json:"durationInDays,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The count of recurrences, if the value of recurrenceEndType is occurrences, or 0 otherwise.
	RecurrenceCount *int64 `json:"recurrenceCount,omitempty"`

	// How the recurrence ends. Possible values: never, endBy, occurrences, or recurrenceCount. If it's never, then there's
	// no explicit end of the recurrence series. If it's endBy, then the recurrence ends at a certain date. If it's
	// occurrences, then the series ends after recurrenceCount instances of the review have completed.
	RecurrenceEndType nullable.Type[string] `json:"recurrenceEndType,omitempty"`

	// The recurrence interval. Possible values: onetime, weekly, monthly, quarterly, halfyearly or annual.
	RecurrenceType nullable.Type[string] `json:"recurrenceType,omitempty"`
}
