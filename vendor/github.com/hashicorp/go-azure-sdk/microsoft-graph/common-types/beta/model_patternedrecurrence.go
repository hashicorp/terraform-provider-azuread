package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PatternedRecurrence struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The frequency of an event. Do not specify for a one-time access review. For access reviews: Do not specify this
	// property for a one-time access review. Only interval, dayOfMonth, and type (weekly, absoluteMonthly) properties of
	// recurrencePattern are supported.
	Pattern *RecurrencePattern `json:"pattern,omitempty"`

	// The duration of an event.
	Range *RecurrenceRange `json:"range,omitempty"`
}
