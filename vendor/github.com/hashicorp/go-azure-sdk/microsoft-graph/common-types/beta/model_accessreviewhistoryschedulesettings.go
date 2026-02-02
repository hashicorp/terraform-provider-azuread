package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewHistoryScheduleSettings struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Recurrence *PatternedRecurrence `json:"recurrence,omitempty"`

	// A duration string in ISO 8601 duration format specifying the lookback period of the generated review history data.
	// For example, if a history definition is scheduled to run on the first of every month, the reportRange is P1M. In this
	// case, on the first of every month, access review history data is collected containing only the previous month's
	// review data. Note: Only years, months, and days ISO 8601 properties are supported. Required.
	ReportRange string `json:"reportRange"`
}
