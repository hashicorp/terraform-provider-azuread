package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerTaskRecurrence struct {
	// The taskId of the next task in this series. This value is assigned at the time the next task in the series is
	// created, and is null prior to that time.
	NextInSeriesTaskId nullable.Type[string] `json:"nextInSeriesTaskId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The 1-based index of this task within the recurrence series. The first task in a series has the value 1, the next
	// task in the series has the value 2, and so on.
	OccurrenceId *int64 `json:"occurrenceId,omitempty"`

	// The taskId of the previous task in this series. null for the first task in a series since it has no predecessor. All
	// subsequent tasks in the series have a value that corresponds to their predecessors.
	PreviousInSeriesTaskId nullable.Type[string] `json:"previousInSeriesTaskId,omitempty"`

	// The date and time when this recurrence series begin. For the first task in a series (occurrenceId = 1) this value is
	// copied from schedule.patternStartDateTime. For subsequent tasks in the series (occurrenceId >= 2) this value is
	// copied from the previous task and never changes; it preserves the start date of the recurring series. The Timestamp
	// type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC
	// on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	RecurrenceStartDateTime *string `json:"recurrenceStartDateTime,omitempty"`

	// The schedule for recurrence. Clients define and edit recurrence by specifying the schedule. If nextInSeriesTaskId
	// isn't assigned, clients may terminate the series by assigning null to this property.
	Schedule *PlannerRecurrenceSchedule `json:"schedule,omitempty"`

	// The recurrence series this task belongs to. A GUID-based value that serves as the unique identifier for a series.
	SeriesId *string `json:"seriesId,omitempty"`
}
