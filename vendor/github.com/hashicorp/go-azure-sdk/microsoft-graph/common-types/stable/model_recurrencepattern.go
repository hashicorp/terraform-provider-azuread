package stable

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecurrencePattern struct {
	// The day of the month on which the event occurs. Required if type is absoluteMonthly or absoluteYearly.
	DayOfMonth *int64 `json:"dayOfMonth,omitempty"`

	// A collection of the days of the week on which the event occurs. The possible values are: sunday, monday, tuesday,
	// wednesday, thursday, friday, saturday. If type is relativeMonthly or relativeYearly, and daysOfWeek specifies more
	// than one day, the event falls on the first day that satisfies the pattern. Required if type is weekly,
	// relativeMonthly, or relativeYearly.
	DaysOfWeek *[]DayOfWeek `json:"daysOfWeek,omitempty"`

	// The first day of the week. The possible values are: sunday, monday, tuesday, wednesday, thursday, friday, saturday.
	// Default is sunday. Required if type is weekly.
	FirstDayOfWeek *DayOfWeek `json:"firstDayOfWeek,omitempty"`

	// Specifies on which instance of the allowed days specified in daysOfWeek the event occurs, counted from the first
	// instance in the month. The possible values are: first, second, third, fourth, last. Default is first. Optional and
	// used if type is relativeMonthly or relativeYearly.
	Index *WeekIndex `json:"index,omitempty"`

	// The number of units between occurrences, where units can be in days, weeks, months, or years, depending on the type.
	// Required.
	Interval int64 `json:"interval"`

	// The month in which the event occurs. This is a number from 1 to 12.
	Month *int64 `json:"month,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The recurrence pattern type: daily, weekly, absoluteMonthly, relativeMonthly, absoluteYearly, relativeYearly.
	// Required. For more information, see values of type property.
	Type RecurrencePatternType `json:"type"`
}
