package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsKioskForceUpdateSchedule struct {
	// Day of month. Valid values 1 to 31
	DayofMonth *int64 `json:"dayofMonth,omitempty"`

	DayofWeek *DayOfWeek `json:"dayofWeek,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Possible values for App update on Windows10 recurrence.
	Recurrence *Windows10AppsUpdateRecurrence `json:"recurrence,omitempty"`

	// If true, runs the task immediately if StartDateTime is in the past, else, runs at the next recurrence.
	RunImmediatelyIfAfterStartDateTime *bool `json:"runImmediatelyIfAfterStartDateTime,omitempty"`

	// The start time for the force restart.
	StartDateTime *string `json:"startDateTime,omitempty"`
}
