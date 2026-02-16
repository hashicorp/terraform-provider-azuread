package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationQuarantine struct {
	// Date and time when the quarantine was last evaluated and imposed. The Timestamp type represents date and time
	// information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z.
	CurrentBegan *string `json:"currentBegan,omitempty"`

	// Describes the error(s) that occurred when putting the synchronization job into quarantine.
	Error *SynchronizationError `json:"error,omitempty"`

	// Date and time when the next attempt to re-evaluate the quarantine will be made. The Timestamp type represents date
	// and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z.
	NextAttempt *string `json:"nextAttempt,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Reason *QuarantineReason `json:"reason,omitempty"`

	// Date and time when the quarantine was first imposed in this series (a series starts when a quarantine is first
	// imposed, and is reset as soon as the quarantine is lifted). The Timestamp type represents date and time information
	// using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	SeriesBegan *string `json:"seriesBegan,omitempty"`

	// Number of times in this series the quarantine was re-evaluated and left in effect (a series starts when quarantine is
	// first imposed, and is reset as soon as quarantine is lifted).
	SeriesCount *int64 `json:"seriesCount,omitempty"`
}
